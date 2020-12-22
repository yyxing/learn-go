package apis

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/sirupsen/logrus"
	"learn-go/web/core/starter"
	"learn-go/web/domain/envelopes"
	"learn-go/web/service"
	"net/http"
)

var envelopeService service.EnvelopeService

func registerEnvelopeHandlers(app *iris.Application) {
	envelopeGroup := app.Party("/v1/envelope")
	envelopeService = envelopes.GetEnvelopeService(starter.DefaultDB())
	envelopeGroup.Post("/", sendRedEnvelope)
}

func sendRedEnvelope(ctx context.Context) {
	envelopeSendDTO := service.RedEnvelopeSendDTO{}
	err := ctx.ReadJSON(&envelopeSendDTO)
	if err != nil {
		logrus.Error(err)
		_, _ = ctx.JSON(Fail(http.StatusBadRequest, "数据解析错误", nil))
		return
	}
	envelopeDTO, err := envelopeService.SendEnvelope(envelopeSendDTO)
	if err != nil {
		logrus.Error(err)
		_, _ = ctx.JSON(Fail(http.StatusInternalServerError, "服务器错误", err))
		return
	}
	_, _ = ctx.JSON(envelopeDTO)
}
