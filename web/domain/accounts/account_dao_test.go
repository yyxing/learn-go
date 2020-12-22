package accounts

import (
	"database/sql"
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"learn-go/web/core/starter"
	"learn-go/web/util"
	"testing"
)

var (
	accountNo = "1l5miERsRUT2xB6pM3ZhwcqGDYs"
)

func TestAccountDao_Insert(t *testing.T) {
	Convey("插入数据", t, func() {
		db := starter.DefaultDB()
		accountDao := AccountDao{Db: db}
		result, err := accountDao.Insert(Account{
			AccountNo:    ksuid.New().String(),
			AccountName:  "Devil",
			AccountType:  1,
			CurrencyCode: "CNY",
			UserId:       ksuid.New().String(),
			Username:     sql.NullString{String: "Devil", Valid: true},
			Balance:      decimal.NewFromFloat(666),
			Status:       1,
		})
		So(err, ShouldBeNil)
		So(result, ShouldEqual, true)
	})
}

func TestAccountDao_Find(t *testing.T) {
	Convey("查询数据", t, func() {
		db := starter.DefaultDB()
		accountDao := AccountDao{Db: db}
		account := accountDao.Find(Account{
			AccountNo: accountNo,
		})
		So(account, ShouldNotBeNil)
		So(len(account), ShouldEqual, 1)
		bytes, _ := util.Marshal(account)
		logrus.Info(string(bytes))
	})
}

func TestAccountDao_UpdateStatus(t *testing.T) {
	Convey("修改状态", t, func() {
		db := starter.DefaultDB()
		accountDao := AccountDao{Db: db}
		result, err := accountDao.UpdateStatus(accountNo, 0)
		So(err, ShouldBeNil)
		So(result, ShouldEqual, true)
		account := accountDao.getByAccountNo(accountNo)
		So(account.Status, ShouldEqual, 0)
	})
}

func TestAccountDao_UpdateBalance(t *testing.T) {
	Convey("修改金额", t, func() {
		db := starter.DefaultDB()
		accountDao := AccountDao{Db: db}
		var amount = float64(-733)
		result, err := accountDao.UpdateBalance(accountNo, decimal.NewFromFloat(amount))
		So(err, ShouldBeNil)
		So(result, ShouldEqual, false)
	})
}
