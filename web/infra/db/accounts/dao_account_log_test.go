package accounts

import (
	"database/sql"
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"learn-go/web/core/boot"
	"learn-go/web/core/starter"
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
		accountLogDao := AccountLogDao{tx: tx}
		accountLog := accountLogDao.GetByLogNo("1l60DFGAf9hksJum6SYuNTzs9P6")
		So(accountLog, ShouldNotBeNil)
		bytes, _ := util.Marshal(accountLog)
		logrus.Info(string(bytes))
	})
}

func TestAccountLogDao_GetByTradeNo(t *testing.T) {
	Convey("根据交易编号查询流水数据", t, func() {
		tx := starter.DefaultDB()
		accountLogDao := AccountLogDao{tx: tx}
		accountLog := accountLogDao.GetByTradeNo("1l60DHpaFn35wt13J0cn9X1Vdh4")
		So(accountLog, ShouldNotBeNil)
		bytes, _ := util.Marshal(accountLog)
		logrus.Info(string(bytes))
	})
}

func TestAccountLogDao_Insert(t *testing.T) {
	Convey("插入流水数据", t, func() {
		tx := starter.DefaultDB()
		accountLogDao := AccountLogDao{tx: tx}
		result, err := accountLogDao.Insert(AccountLog{
			LogNo:           ksuid.New().String(),
			TradeNo:         ksuid.New().String(),
			AccountNo:       ksuid.New().String(),
			UserId:          ksuid.New().String(),
			Username:        sql.NullString{String: "Devil", Valid: true},
			TargetAccountNo: ksuid.New().String(),
			TargetUserId:    ksuid.New().String(),
			TargetUsername:  sql.NullString{String: "Lu", Valid: true},
			Amount:          decimal.NewFromFloat(333),
			Balance:         decimal.NewFromFloat(666),
			TransferType:    100,
			ChangeFlag:      100,
			Status:          1,
			Desc:            "这是描述",
		})
		So(err, ShouldBeNil)
		So(result, ShouldEqual, true)
	})
}
