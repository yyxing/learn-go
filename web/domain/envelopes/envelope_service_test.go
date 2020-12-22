package envelopes

import (
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"learn-go/web/core/starter"
	"learn-go/web/enums"
	"learn-go/web/service"
	"learn-go/web/util"
	"testing"
)

func TestEnvelopeService_FindEnvelopes(t *testing.T) {
	Convey("查询可领红包列表", t, func() {
		envelopeService := GetEnvelopeService(starter.DefaultDB())
		envelopePage := envelopeService.FindEnvelopes(util.PageParam{
			All:     false,
			Size:    10,
			Current: 1,
		})
		logrus.Info(envelopePage.Records)
		So(envelopePage, ShouldNotBeNil)
		So(envelopePage.Total, ShouldEqual, 1)
	})
}

func TestEnvelopeService_FindReceiveEnvelopes(t *testing.T) {
	Convey("查询用户领取的红包", t, func() {
		envelopeService := GetEnvelopeService(starter.DefaultDB())
		envelopePage := envelopeService.FindReceiveEnvelopes("333301933932544001", util.PageParam{
			All:     false,
			Size:    10,
			Current: 1,
		})
		logrus.Info(envelopePage.Records)
		So(envelopePage, ShouldNotBeNil)
		So(envelopePage.Total, ShouldEqual, 0)
	})
}

func TestEnvelopeService_FindSendEnvelopes(t *testing.T) {
	Convey("查询用户发的红包详情", t, func() {
		envelopeService := GetEnvelopeService(starter.DefaultDB())
		envelopePage := envelopeService.FindSendEnvelopes("333301933932544001", util.PageParam{
			All:     false,
			Size:    10,
			Current: 1,
		})
		logrus.Info(envelopePage.Records)
		So(envelopePage, ShouldNotBeNil)
		So(envelopePage.Total, ShouldEqual, 1)
	})
}

func TestEnvelopeService_GetOneEnvelope(t *testing.T) {
	Convey("测试查询单个红包详情", t, func() {
		envelopeService := GetEnvelopeService(starter.DefaultDB())
		envelopeDTO, err := envelopeService.GetOneEnvelope("333750832119939073")
		So(envelopeDTO, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})
}

func TestEnvelopeService_ReceiveEnvelope(t *testing.T) {
	Convey("测试收红包", t, func() {
		tx := starter.DefaultDB().Begin()
		_ = starter.Transaction(tx, func() error {
			envelopeService := GetEnvelopeService(tx)
			envelopeItemDTO, err := envelopeService.ReceiveEnvelope(service.RedEnvelopeReceiveDTO{
				EnvelopeNo:      "333842611578077185",
				ReceiveUsername: "BabyLu",
				ReceiveUserId:   "333755084590546945",
				AccountNo:       "6225774590612481",
			})
			logrus.Info(envelopeItemDTO)
			logrus.Error(err)
			//So(envelopeItemDTO, ShouldNotBeNil)
			//So(err, ShouldBeNil)
			//So(envelopeItemDTO.ReceiveAccountNo, ShouldEqual, "6225774590612481")
			return err
		})
	})
}

func TestEnvelopeService_SendEnvelope(t *testing.T) {
	Convey("测试发红包", t, func() {
		envelopeService := GetEnvelopeService(starter.DefaultDB())
		envelopeDTO, err := envelopeService.SendEnvelope(service.RedEnvelopeSendDTO{
			EnvelopeType: enums.TryLuck,
			Username:     "LuLu",
			UserId:       "333301933932544001",
			Blessing:     "恭喜发财，大吉大利",
			Amount:       "66.66",
			Quantity:     6,
			AccountNo:    "6225773932609537",
		})
		So(envelopeDTO, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})
}

func TestEnvelopeService_Refund(t *testing.T) {
	Convey("测试红包退款", t, func() {
		envelopeService := GetEnvelopeService(starter.DefaultDB())
		envelopeDTO, err := envelopeService.Refund("333842611578077185")
		So(envelopeDTO, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})
}
