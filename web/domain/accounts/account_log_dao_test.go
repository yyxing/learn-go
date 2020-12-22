package accounts

import (
	"database/sql"
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"learn-go/web/core/boot"
	"learn-go/web/core/starter"
	"learn-go/web/enums"
	"learn-go/web/util"
	"testing"
)

func TestMain(m *testing.M) {
	_ = boot.TestEnv(
		"C:\\Users\\bodenai\\go\\src\\learn-go\\web\\resource\\application.yml")
	m.Run()
}

func TestAccountLogDao_GetByLogNo(t *testing.T) {
	Convey("根据流水编号查询流水数据", t, func() {
		tx := starter.DefaultDB()
		accountLogDao := AccountLogDao{Db: tx}
		accountLog := accountLogDao.GetByLogNo("1l60DFGAf9hksJum6SYuNTzs9P6")
		So(accountLog, ShouldNotBeNil)
		bytes, _ := util.Marshal(accountLog)
		logrus.Info(string(bytes))
	})
}

func TestAccountLogDao_GetByTradeNo(t *testing.T) {
	Convey("根据交易编号查询流水数据", t, func() {
		tx := starter.DefaultDB()
		accountLogDao := AccountLogDao{Db: tx}
		accountLog := accountLogDao.GetByTradeNo("1l60DHpaFn35wt13J0cn9X1Vdh4")
		So(accountLog, ShouldNotBeNil)
		bytes, _ := util.Marshal(accountLog)
		logrus.Info(string(bytes))
	})
}

func TestAccountLogDao_Insert(t *testing.T) {
	Convey("插入流水数据", t, func() {
		tx := starter.DefaultDB()
		accountLogDao := AccountLogDao{Db: tx}
		result, err := accountLogDao.Insert(AccountLog{
			LogNo:                 ksuid.New().String(),
			TradeNo:               ksuid.New().String(),
			AccountNo:             ksuid.New().String(),
			UserId:                ksuid.New().String(),
			Username:              sql.NullString{String: "Devil", Valid: true},
			CounterpartyAccountNo: ksuid.New().String(),
			CounterpartyUserId:    ksuid.New().String(),
			CounterpartyUsername:  sql.NullString{String: "Lu", Valid: true},
			Amount:                decimal.NewFromFloat(333),
			Balance:               decimal.NewFromFloat(666),
			TradeType:             enums.AccountCreated,
			Status:                enums.TradeSuccess,
			Desc:                  "这是描述",
		})
		So(err, ShouldBeNil)
		So(result, ShouldEqual, true)
	})
}
