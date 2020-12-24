package apis

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"learn-go/web/core/starter"
	"learn-go/web/domain/files"
	"learn-go/web/service"
	"net/http"
)

var (
	// 数据库
	fileDb *gorm.DB
	// 查询service
	fileService service.FileService
)

func registerFileHandlers(app *iris.Application) {
	fileService = files.GetFileService(starter.DefaultDB())
	app.Post("/upload", func(ctx context.Context) {
		fileDTO := service.FileDTO{}
		err := ctx.ReadForm(&fileDTO)
		if err != nil {
			logrus.Error(err)
			_, _ = ctx.JSON(Fail(http.StatusBadRequest, "数据解析错误", err))
			return
		}
		file, info, err := ctx.FormFile("file")
		if err != nil {
			logrus.Error(err)
			_, _ = ctx.JSON(Fail(http.StatusInternalServerError, "服务器错误", err))
			return
		}
		fileDTO.FileSize = info.Size
		fileDTO.FileName = info.Filename
		var fileBytes = make([]byte, info.Size)
		_, err = file.Read(fileBytes)
		fileDTO.Bytes = fileBytes
		if err != nil {
			logrus.Error(err)
			_, _ = ctx.JSON(Fail(http.StatusInternalServerError, "文件读取失败", err))
		}
		err = fileService.UploadFile(fileDTO)
		if err != nil {
			logrus.Error(err)
			_, _ = ctx.JSON(Fail(http.StatusInternalServerError, "服务器错误", fileDTO))
			return
		}
		_, _ = ctx.JSON(Success("", nil))
	})
}
