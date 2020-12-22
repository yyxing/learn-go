package accounts

import (
	. "github.com/smartystreets/goconvey/convey"
	"learn-go/web/core/starter"
	"learn-go/web/enums"
	"learn-go/web/service"
	"testing"
)

func TestAccountService_CreateAccount(t *testing.T) {
	Convey("创建账户", t, func() {
		accountService := accountService{tx: starter.DefaultDB()}
		accountDto, err := accountService.CreateAccount(service.AccountCreatedDTO{
			Username:     "BabyLu",
			AccountName:  "借记卡",
			AccountType:  1,
			CurrencyCode: "CNY",
			Amount:       "66666",
		})
		So(err, ShouldBeNil)
		So(accountDto, ShouldNotBeNil)
		So(accountDto.Balance, ShouldEqual, "66666")
	})
}

func TestAccountService_GetEnvelopeAccountByUserId(t *testing.T) {
	Convey("根据userId查询账户信息", t, func() {
		accountService := accountService{starter.DefaultDB()}
		accountList, err := accountService.GetEnvelopeAccountByUserId(service.AccountDTO{
			AccountCreatedDTO: service.AccountCreatedDTO{
				Username:    "Devil",
				AccountType: 1,
			},
		})
		So(err, ShouldBeNil)
		So(accountList, ShouldNotBeNil)
	})
}

func TestAccountService_GetAccountByAccountNo(t *testing.T) {
	Convey("根据accountNo查询账户信息", t, func() {
		accountService := accountService{}
		account, err := accountService.GetAccountByAccountNo("6225777003618305")
		So(err, ShouldBeNil)
		So(account, ShouldNotBeNil)
		So("6225777003618305", ShouldEqual, account.AccountNo)
	})
}

func TestAccountService_StoreValue(t *testing.T) {
	Convey("对账户进行充值", t, func() {
		accountService := accountService{tx: starter.DefaultDB()}
		status, err := accountService.StoreValue(service.AccountTransferDTO{
			TradeAccount: service.TradeParticipator{
				AccountNo: "6225773932609537",
				UserId:    "333301933932544001",
				Username:  "LuLu",
			},
			Amount:    "88888",
			TradeDesc: "充值",
		})
		So(err, ShouldBeNil)
		So(status, ShouldNotBeNil)
		So(status, ShouldEqual, enums.TransferSuccess)
	})
}

func TestAccountService_Transfer(t *testing.T) {
	Convey("A账户对B账户进行转账", t, func() {
		accountService := accountService{}
		status, err := accountService.Transfer(service.AccountTransferDTO{
			TradeAccount: service.TradeParticipator{
				AccountNo: "6225777003618305",
				UserId:    "Devil",
				Username:  "Devil",
			},
			CounterpartyAccount: service.TradeParticipator{
				AccountNo: "6225778531903489",
				UserId:    "LuLu",
				Username:  "LuLu",
			},
			Amount:       "888",
			TransferType: enums.TransferOutgoing,
			TradeDesc:    "转账支出",
		})
		accountBody, _ := accountService.GetAccountByAccountNo("6225777003618305")
		accountTarget, _ := accountService.GetAccountByAccountNo("6225778531903489")
		So(err, ShouldBeNil)
		So(status, ShouldNotBeNil)
		So(status, ShouldEqual, enums.TransferSuccess)
		So(accountBody.Balance, ShouldEqual, "666")
		So(accountTarget.Balance, ShouldEqual, "1554")
	})
	Convey("A账户对B账户进行转账", t, func() {
		accountService := accountService{}
		status, err := accountService.Transfer(service.AccountTransferDTO{
			TradeAccount: service.TradeParticipator{
				AccountNo: "6225777003618305",
				UserId:    "Devil",
				Username:  "Devil",
			},
			CounterpartyAccount: service.TradeParticipator{
				AccountNo: "6225778531903489",
				UserId:    "LuLu",
				Username:  "LuLu",
			},
			Amount:       "888",
			TransferType: enums.TransferOutgoing,
			TradeDesc:    "转账支出",
		})
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "账户余额不足")
		So(status, ShouldEqual, enums.TransferSufficientFunds)
	})
}

//func TestTransaction(t *testing.T) {
//	db := starter.DefaultDB()
//	tx := db.Begin()
//	tx.SavePoint("test")
//	envelopeDao := envelopes.EnvelopeDao{Db: tx}
//	_ = starter.DefaultDB().Transaction(func(tx *gorm.DB) error {
//		_, _ = envelopeDao.InsertReceiveRedEnvelopeLog(envelopes.RedEnvelopeItem{
//			ItemNo:          "123",
//			EnvelopeNo:      "123",
//			ReceiveUsername: sql.NullString{},
//			ReceiveUserId:   "123",
//			Amount:          decimal.Decimal{},
//			Quantity:        1,
//			RemainAmount:    decimal.Decimal{},
//			RemainQuantity:    decimal.Decimal{},
//			AccountNo:       "123",
//			PayStatus:       1,
//			OrderDesc:       sql.NullString{},
//		})
//		return nil
//	})
//	accountService := GetAccountService(tx)
//	accountService.Transfer(service.AccountTransferDTO{
//		TradeAccount:        service.TradeParticipator{
//			AccountNo: "123",
//			UserId:    "123",
//			Username:  "123",
//		},
//		CounterpartyAccount: service.TradeParticipator{
//			AccountNo: "46546",
//			UserId:    "456",
//			Username:  "456",
//		},
//		Amount:              "666",
//		TransferType:        1,
//		TradeDesc:           "转账",
//	})
//	tx.RollbackTo("test")
//}
