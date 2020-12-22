package apis

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/sirupsen/logrus"
	"learn-go/web/core/starter"
	"learn-go/web/domain/envelopes"
	"learn-go/web/service"
	"learn-go/web/util"
	"net/http"
)

var envelopeService service.EnvelopeService

func registerEnvelopeHandlers(app *iris.Application) {
	envelopeGroup := app.Party("/v1/envelope")
	envelopeService = envelopes.GetEnvelopeService(starter.DefaultDB())
	envelopeGroup.Post("/", sendRedEnvelope)
	envelopeGroup.Get("/{envelopeId:string}", getRedEnvelope)
	envelopeGroup.Get("/user/{userId:string}/receive", userReceiveRedEnvelope)
	envelopeGroup.Get("/", findRedEnvelopes)
	envelopeGroup.Get("/user/{userId:string}/send", findSendEnvelopes)
	envelopeGroup.Post("/receive", receiveRedEnvelopes)
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
func getRedEnvelope(ctx context.Context) {
	envelopeId := ctx.Params().GetString("envelopeId")
	envelopeDTO, err := envelopeService.GetOneEnvelope(envelopeId)
	if err != nil {
		logrus.Error(err)
		_, _ = ctx.JSON(Fail(http.StatusInternalServerError, "服务器错误", err))
		return
	}
	_, _ = ctx.JSON(envelopeDTO)
}
func userReceiveRedEnvelope(ctx context.Context) {
	userId := ctx.Params().GetString("userId")
	pageParam := util.PageParam{}
	err := ctx.ReadJSON(&pageParam)
	if err != nil {
		logrus.Error(err)
		_, _ = ctx.JSON(Fail(http.StatusBadRequest, "数据解析错误", nil))
		return
	}
	redEnvelopePage := envelopeService.FindReceiveEnvelopes(userId, pageParam)
	_, _ = ctx.JSON(redEnvelopePage)
}
func findRedEnvelopes(ctx context.Context) {
	pageParam := util.PageParam{}
	err := ctx.ReadJSON(&pageParam)
	if err != nil {
		logrus.Error(err)
		_, _ = ctx.JSON(Fail(http.StatusBadRequest, "数据解析错误", nil))
		return
	}
	redEnvelopePage := envelopeService.FindEnvelopes(pageParam)
	_, _ = ctx.JSON(redEnvelopePage)
}
func findSendEnvelopes(ctx context.Context) {
	userId := ctx.Params().GetString("userId")
	pageParam := util.PageParam{}
	err := ctx.ReadJSON(&pageParam)
	if err != nil {
		logrus.Error(err)
		_, _ = ctx.JSON(Fail(http.StatusBadRequest, "数据解析错误", nil))
		return
	}
	redEnvelopePage := envelopeService.FindSendEnvelopes(userId, pageParam)
	_, _ = ctx.JSON(redEnvelopePage)
}
func receiveRedEnvelopes(ctx context.Context) {
	redEnvelopeReceiveDTO := service.RedEnvelopeReceiveDTO{}
	err := ctx.ReadJSON(&redEnvelopeReceiveDTO)
	if err != nil {
		logrus.Error(err)
		_, _ = ctx.JSON(Fail(http.StatusBadRequest, "数据解析错误", nil))
		return
	}
	redEnvelopeItemDTO, err := envelopeService.ReceiveEnvelope(redEnvelopeReceiveDTO)
	if err != nil {
		logrus.Error(err)
		_, _ = ctx.JSON(Fail(http.StatusInternalServerError, "服务器错误", redEnvelopeItemDTO))
		return
	}
	_, _ = ctx.JSON(redEnvelopeItemDTO)
}
