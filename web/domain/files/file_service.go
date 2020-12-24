package files

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"learn-go/web/core/starter"
	"learn-go/web/service"
	"learn-go/web/util"
	"net/url"
	"sync"
	"time"
)

var (
	db          *gorm.DB
	once        sync.Once
	serviceImpl *fileService
	minioClient *minio.Client
)

type fileService struct {
	tx      *gorm.DB
	fileDao *FileDao
}

type OSSConfig struct {
	Endpoint            string
	BucketName          string
	AccessKey           string
	SecretKey           string
	BucketNameThumbnail string
}

func initMinio(config OSSConfig) {
	var err error
	minioClient, err = minio.New(config.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
	})
	if err != nil {
		logrus.Error(err)
	}
}

func GetFileService(tx *gorm.DB) service.FileService {
	mutex := sync.Mutex{}
	mutex.Lock()
	if serviceImpl == nil {
		once.Do(func() {
			serviceImpl = new(fileService)
			serviceImpl.fileDao = &FileDao{}
			serviceImpl.tx = tx
			serviceImpl.fileDao.tx = tx
		})
		if minioClient == nil {
			ossConfig := OSSConfig{}
			config := starter.GetConfig()
			err := config.UnmarshalKey("oss", &ossConfig)
			if err != nil {
				logrus.Error(err)
			}
			initMinio(ossConfig)
		}
		return serviceImpl
	}
	mutex.Unlock()
	serviceImpl.tx = tx
	serviceImpl.fileDao.tx = tx
	return serviceImpl
}

func (svc *fileService) UploadFile(dto service.FileDTO) error {
	err := svc.tx.Transaction(func(tx *gorm.DB) error {
		fileDao := svc.fileDao
		// 生成存储文件
		file := &StorageFile{}
		file.fromDTO(dto)
		_, fileError := fileDao.InsertStorageFile(file)
		if fileError != nil {
			return fileError
		}
		// 生成文件信息
		fileInfo := &FileInfo{}
		fileError = fileInfo.fromDTO(file.ID, dto)
		if fileError != nil {
			return fileError
		}
		_, fileError = fileDao.InsertFileInfo(fileInfo)
		if fileError != nil {
			return fileError
		}
		// 上传文件到minio
		objectUrl, err := svc.upload("local", dto.FileName, dto)
		if err != nil {
			return err
		}
		_, fileError = fileDao.UpdateUrl(file.ID, objectUrl, objectUrl)
		return fileError
	})
	return err
}

func (svc *fileService) upload(bucketName string, objectName string, dto service.FileDTO) (string, error) {
	_, err := minioClient.PutObject(context.Background(), bucketName, objectName, bytes.NewReader(dto.Bytes), dto.FileSize, minio.PutObjectOptions{
		ContentType: util.GetFileType(dto.Bytes),
	})
	if err != nil {
		return "", err
	}
	return svc.getUrl(bucketName, objectName)
}
func (svc *fileService) getUrl(bucketName string, objectName string) (string, error) {
	reqParams := make(url.Values)
	objectUrl, err := minioClient.PresignedGetObject(context.Background(), bucketName, objectName,
		time.Hour*24*7, reqParams)
	if err != nil {
		return "", err
	}
	return objectUrl.String(), nil
}
