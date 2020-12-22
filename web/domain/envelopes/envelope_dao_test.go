package envelopes

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"learn-go/web/core/starter"
	"learn-go/web/enums"
	"learn-go/web/util"
	"testing"
	"time"
)

func TestEnvelopeDao_FindSendByUserId(t *testing.T) {
	Convey("查询用户发送红包", t, func() {
		envelopeDao := EnvelopeDao{tx: starter.DefaultDB()}
		param := util.PageParam{
			All:     false,
			Size:    10,
			Current: 1,
		}
		redEnvelopeGoods := envelopeDao.FindSendByUserId("333301385552461825", param.Page())
		So(redEnvelopeGoods, ShouldNotBeNil)
		So(len(redEnvelopeGoods), ShouldEqual, 0)
	})
}

func TestEnvelopeDao_FindExpired(t *testing.T) {
	Convey("查询过期红包", t, func() {
		envelopeDao := EnvelopeDao{tx: starter.DefaultDB()}
		param := util.PageParam{
			All:     false,
			Size:    1,
			Current: 2,
		}
		redEnvelopeGoods := envelopeDao.FindExpired(param.Page())
		logrus.Info(redEnvelopeGoods[0])
		So(redEnvelopeGoods, ShouldNotBeNil)
		So(len(redEnvelopeGoods), ShouldEqual, 1)
	})
}

func TestEnvelopeDao_FindExpiredByUserId(t *testing.T) {
	Convey("查询用户发送且过期红包", t, func() {
		envelopeDao := EnvelopeDao{tx: starter.DefaultDB()}
		param := util.PageParam{
			All:     false,
			Size:    10,
			Current: 1,
		}
		redEnvelopeGoods := envelopeDao.FindExpiredByUserId("333301933932544001", param.Page())
		logrus.Info(redEnvelopeGoods[0])
		So(redEnvelopeGoods, ShouldNotBeNil)
		So(len(redEnvelopeGoods), ShouldEqual, 1)
	})
}

func TestEnvelopeDao_GetOne(t *testing.T) {
	Convey("查询单个红包", t, func() {
		envelopeDao := EnvelopeDao{tx: starter.DefaultDB()}
		redEnvelopeGood := envelopeDao.GetOne("333569671976452097")
		So(redEnvelopeGood, ShouldNotBeNil)
		So(redEnvelopeGood.Amount, ShouldEqual, decimal.NewFromFloat(66.66))
	})
}

func TestEnvelopeDao_Insert(t *testing.T) {
	Convey("测试发红包", t, func() {
		envelopeDao := EnvelopeDao{tx: starter.DefaultDB()}
		result, err := envelopeDao.Insert(RedEnvelopeGood{
			EnvelopeNo:     util.NextId(),
			EnvelopeType:   enums.TryLuck,
			Username:       sql.NullString{String: "LuLu", Valid: true},
			UserId:         "333301933932544001",
			Blessing:       sql.NullString{String: "恭喜发财，大吉大利", Valid: true},
			Amount:         decimal.NewFromFloat(88.88),
			AccountNo:      "6225773932609537",
			Quantity:       6,
			RemainAmount:   decimal.NewFromFloat(88.88),
			RemainQuantity: 6,
			ExpiredAt:      time.Now().Add(time.Minute * 5),
			PayStatus:      1,
			OrderStatus:    enums.OrderCreate,
			OrderType:      enums.Create,
		})
		So(err, ShouldBeNil)
		So(result, ShouldEqual, true)
	})
}

func TestEnvelopeDao_ReceiveRedEnvelope(t *testing.T) {
	Convey("测试收红包", t, func() {
		envelopeDao := EnvelopeDao{tx: starter.DefaultDB()}
		result, err := envelopeDao.ReceiveRedEnvelope("333569689072435201", decimal.NewFromFloat(44.44))
		So(result, ShouldNotBeNil)
		So(err, ShouldBeNil)
		So(result, ShouldEqual, true)
	})
}
