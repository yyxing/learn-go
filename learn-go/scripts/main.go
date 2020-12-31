package main

import (
	"archive/zip"
	"bytes"
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type LabelResult struct {
	ImageId     int                   `bson:"image_id"`
	StageRunId  int                   `bson:"stage_run_id"`
	PackageId   int                   `bson:"package_id"`
	ImageHeight int                   `bson:"image_height"`
	ImageWidth  int                   `bson:"image_width"`
	ImageName   string                `bson:"image_name"`
	CreateAt    time.Time             `bson:"create_at"`
	UpdateAt    time.Time             `bson:"update_at"`
	Content     map[string][]Instance `bson:"content"`
}
type Instance struct {
	ClassId      int    `bson:"classId"`
	FeatureType  int    `bson:"featureType"`
	FeatureValue string `bson:"featureValue"`
}

var (
	minioClient *minio.Client
	mongoClient *qmgo.QmgoClient
	ctx         = context.Background()
)

func init() {
	var err error
	minioClient, err = minio.New("minio:9000", &minio.Options{
		Creds: credentials.NewStaticV4("bodenai-minio-key", "bodenai2019", ""),
	})
	if err != nil {
		panic(err)
	}
	// 设置客户端连接配置
	ctx := context.Background()
	// 连接到MongoDB
	mongoClient, err = qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://root:bodenai2019@ouroboros-monggodb.mongodb:27017",
		Database: "ouroborosLabel", Coll: "label_result"})
	if err != nil {
		panic(err)
	}
}

type CocoJson struct {
	Info        CocoInfo         `json:"info"`
	Licenses    []CocoLicense    `json:"licenses"`
	Images      []CocoImage      `json:"images"`
	Annotations []CocoAnnotation `json:"annotations"`
	Categories  []CocoCategory   `json:"categories"`
}

type CocoInfo struct {
	Year        int       `json:"year"`
	Version     string    `json:"version"`
	Description string    `json:"description"`
	Contributor string    `json:"contributor"`
	Url         string    `json:"url"`
	DateCreated time.Time `json:"date_created"`
}
type CocoLicense struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}
type CocoImage struct {
	Id           int       `json:"id"`
	Width        int       `json:"width"`
	Height       int       `json:"height"`
	FileName     string    `json:"file_name"`
	License      string    `json:"license"`
	FlickrUrl    string    `json:"flickr_url"`
	CocoUrl      string    `json:"coco_url"`
	DateCaptured time.Time `json:"date_captured"`
}
type CocoAnnotation struct {
	Id           int         `json:"id"`
	ImageId      int         `json:"image_id"`
	CategoryId   int         `json:"category_id"`
	Segmentation [][]float64 `json:"segmentation"`
	Area         float64     `json:"area"`
	Bbox         []float64   `json:"bbox"`
	IsCrowd      int         `json:"iscrowd"`
}

type CocoCategory struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	SuperCategory string `json:"supercategory"`
}

// 连接mongodb获取标注结果
func GetLabelResult() []*LabelResult {
	var labelResult []*LabelResult
	filter := bson.D{{}}
	err := mongoClient.Find(ctx, filter).All(&labelResult)
	if err != nil {
		panic(err)
	}
	return labelResult
}

func CreateCocoJson() ([]byte, error) {
	labelResult := GetLabelResult()
	var cocoImages []CocoImage
	var cocoAnnotations []CocoAnnotation
	//var cocoCategories []CocoCategory
	// 生成json文件
	for _, result := range labelResult {
		if result.ImageId <= 0 {
			continue
		}
		cocoImage := CocoImage{
			Id:           result.ImageId,
			Width:        result.ImageWidth,
			Height:       result.ImageHeight,
			FileName:     result.ImageName,
			DateCaptured: time.Now(),
		}
		cocoImages = append(cocoImages, cocoImage)
		instances := result.Content["instances"]
		for _, instance := range instances {
			if instance.FeatureType > 1 {
				continue
			}
			featureValue := instance.FeatureValue
			values := strings.Split(featureValue[1:len(featureValue)-1], ",")
			var segmentation [][]float64
			var bbox []float64
			for _, value := range values {
				v, err := strconv.ParseFloat(value, 64)
				if err != nil {
					log.Fatal(err)
				}
				bbox = append(bbox, v)
			}
			segmentation = append(segmentation, bbox)
			cocoAnnotation := CocoAnnotation{
				Id:           0,
				ImageId:      cocoImage.Id,
				CategoryId:   instance.ClassId,
				Segmentation: segmentation,
				Bbox:         bbox,
				IsCrowd:      0,
			}
			cocoAnnotations = append(cocoAnnotations, cocoAnnotation)
		}
	}
	json := CocoJson{
		Info: CocoInfo{
			Year:        time.Now().Year(),
			Version:     "v1",
			Description: "boden coco export",
			Url:         "https://www.bodenai.com/",
			DateCreated: time.Now(),
		},
		Licenses:    []CocoLicense{},
		Images:      cocoImages,
		Annotations: cocoAnnotations,
		Categories:  []CocoCategory{},
	}
	jsonBytes, err := jsoniter.Marshal(json)
	return jsonBytes, err
}
func Compress(jsonBytes []byte) ([]byte, os.FileInfo, error) {
	// 压缩文件路径
	zipDir := "/coco/export"
	if _, err := os.Stat(zipDir); err != nil {
		err := os.MkdirAll(zipDir, 0777)
		if err != nil {
			return nil, nil, err
		}
	}
	// 创建具体的压缩文件
	zipFile, err := os.Create(path.Join(zipDir, "export.zip"))
	if err != nil {
		return nil, nil, err
	}
	defer zipFile.Close()
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()
	// 获取数据集地址
	datasetId := os.Getenv("DATASET_ID")
	// 获取图片地址
	imagePath := "/data/export/boden-file/" + datasetId
	imageFiles, err := ioutil.ReadDir(imagePath)
	if err != nil {
		return nil, nil, err
	}
	// 生成图片至压缩包中
	imageDir := "/images/"
	for _, image := range imageFiles {
		writer, err := zipWriter.Create(imageDir + image.Name())
		if err != nil {
			return nil, nil, err
		}
		file, err := ioutil.ReadFile(imagePath + image.Name())
		if err != nil {
			return nil, nil, err
		}
		_, _ = writer.Write(file)
	}
	// 生成json文件至压缩包中
	writer, err := zipWriter.Create("/annotations/train.json")
	if err != nil {
		return nil, nil, err
	}
	_, _ = writer.Write(jsonBytes)
	// 执行压缩 生成压缩包
	zipWriter.Flush()

	// 重新读取生成的压缩包 将其传到minio中
	zipFile, err = os.Open(path.Join(zipDir, "export.zip"))
	if err != nil {
		return nil, nil, err
	}
	zipInfo, err := zipFile.Stat()
	if err != nil {
		return nil, nil, err
	}
	zipBytes := make([]byte, zipInfo.Size())
	_, err = zipFile.Read(zipBytes)
	return zipBytes, zipInfo, err
}

func Upload(bucketName string, objectName string, fileBytes []byte, fileSize int64) error {
	_, err := minioClient.PutObject(context.Background(), bucketName, objectName, bytes.NewReader(fileBytes), fileSize, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	if err != nil {
		return err
	}
	return nil
}

type Response struct {
	taskId string
	status compressStatus
	err    string
}

func sendResult(status compressStatus, err error) {
	var errMsg string
	callback := os.Getenv("CALLBACK_URL")
	taskId := os.Getenv("TASK_ID")
	if err != nil {
		errMsg = err.Error()
	}
	response := Response{
		taskId: taskId,
		status: status,
		err:    errMsg,
	}
	data, err := jsoniter.Marshal(response)
	if err != nil {
		panic(err)
	}
	_, err = http.Post(callback, "application/json", bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
}

type compressStatus uint8

const (
	success compressStatus = iota
	fail
)

func main() {
	json, err := CreateCocoJson()
	if err != nil {
		sendResult(fail, err)
		return
	}
	zipBytes, zipInfo, err := Compress(json)
	if err != nil {
		sendResult(fail, err)
		return
	}
	err = Upload("boden-zip", zipInfo.Name(), zipBytes, zipInfo.Size())
	if err != nil {
		sendResult(fail, err)
		panic(err)
	}
	sendResult(success, nil)
}
