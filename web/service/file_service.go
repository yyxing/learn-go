package service

// 文件服务
type FileService interface {
	// 上传接口
	UploadFile(dto FileDTO) error
}

type FileDTO struct {
	// 数据集编号
	DatasetId int `gorm:"not null"`
	// 设备名称
	DeviceName string
	// 设备索引
	DeviceIndex int
	// 文件序号
	FileIndex int
	// 文件类型
	FileType int
	// 文件名称
	FileName string `json:"-"`
	// 文件大小
	FileSize int64 `json:"-"`
	// 文件byte数组
	Bytes []byte
}
