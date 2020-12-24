package files

import (
	"bytes"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"learn-go/web/service"
)

type FileInfo struct {
	ID uint64 `gorm:"primary_key"`
	//图片高度
	ImageHeight int
	//图片宽度
	ImageWidth int
	//文件组
	FileKey string
	// 设备名称
	DeviceName string
	// 设备索引
	DeviceIndex int
	// 文件序号
	FileIndex int `gorm:"not null"`
	// 文件类型
	FileType int `gorm:"not null"`
	// 标定文件信息
	Content string
	// 文件编号
	FileId uint64 `gorm:"not null;"`
	// 视频长度
	VideoLength int
	// 数据集编号
	DatasetId int `gorm:"not null"`
}

func (FileInfo) TableName() string {
	return "ouroboros_file_info"
}

func (do *FileInfo) fromDTO(fileId uint64, dto service.FileDTO) error {
	do.DatasetId = dto.DatasetId
	do.FileType = dto.FileType
	do.FileIndex = dto.FileIndex
	do.FileId = fileId
	do.DeviceIndex = dto.DeviceIndex
	do.DeviceName = dto.DeviceName
	im, _, err := image.Decode(bytes.NewReader(dto.Bytes))
	if err != nil {
		return err
	}
	do.ImageHeight = im.Bounds().Dy()
	do.ImageWidth = im.Bounds().Dx()
	return nil
}
