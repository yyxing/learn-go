package files

import (
	"gorm.io/gorm"
)

type FileDao struct {
	tx *gorm.DB
}

func (dao *FileDao) InsertStorageFile(file *StorageFile) (bool, error) {
	result := dao.tx.Create(&file)
	return result.RowsAffected > 0, result.Error
}

func (dao *FileDao) InsertFileInfo(info *FileInfo) (bool, error) {
	result := dao.tx.Create(&info)
	return result.RowsAffected > 0, result.Error
}

func (dao *FileDao) UpdateUrl(id uint64, url string, thumbnailUrl string) (bool, error) {
	result := dao.tx.Model(&StorageFile{}).Where("id", id).
		UpdateColumns(map[string]interface{}{"thumbnail_url": thumbnailUrl, "url": url})
	return result.RowsAffected > 0, result.Error
}
