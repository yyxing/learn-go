package files

import (
	"crypto/md5"
	"fmt"
	"learn-go/web/service"
	"path"
	"time"
)

type StorageFile struct {
	ID uint64 `gorm:"primary_key"`
	// 缩略图路径
	ThumbnailUrl string
	// 文件路径
	Path string
	// url路径
	Url string `gorm:"not null;"`
	// 后缀
	Suffix string `gorm:"not null;"`
	// md5
	Mask string
	// 文件名称
	Filename string `gorm:"not null"`
	// 文件大小
	FileSize int64
	// 创建时间
	CreateTime time.Time
	// 修改时间
	UpdateTime time.Time
}

func (StorageFile) TableName() string {
	return "ouroboros_storage_file"
}

func (do *StorageFile) fromDTO(dto service.FileDTO) {
	do.FileSize = dto.FileSize
	do.Filename = dto.FileName
	do.Mask = fmt.Sprintf("%x", md5.Sum(dto.Bytes))
	do.Suffix = path.Ext(dto.FileName)[1:]
	do.CreateTime = time.Now()
	do.UpdateTime = time.Now()
}
