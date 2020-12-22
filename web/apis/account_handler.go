package apis

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"learn-go/web/core/starter"
	"learn-go/web/domain/accounts"
	"learn-go/web/service"
	"net/http"
)

var (
	// db连接
	accountDb *gorm.DB
	// 无需事务的服务
	accountService service.AccountService
)

func registerAccountHandlers(app *iris.Application) {
	accountDb = starter.DefaultDB()
	accountService = accounts.GetAccountService(accountDb)
	accountGroup := app.Party("/v1/account")
	accountGroup.Post("/", createAccount)
	accountGroup.Get("/{accountNo:string}", findAccount)
	accountGroup.Put("/transfer", transfer)
	accountGroup.Put("/store", storeValue)
}

// 创建账户api
func createAccount(ctx context.Context) {
	accountCreatedDTO := service.AccountCreatedDTO{}
	err := ctx.ReadJSON(&accountCreatedDTO)
	if err != nil {
		logrus.Error(err)
		_, _ = ctx.JSON(Fail(http.StatusBadRequest, "数据解析错误", err))
		return
	}
	var accountDTO *service.AccountDTO
	// 开启事务
	tx := accountDb.Begin()
	err = starter.Transaction(tx, func() error {
		accountService := accounts.GetAccountService(tx)
		account, createError := accountService.CreateAccount(accountCreatedDTO)
		accountDTO = account
		return createError
	})
	if err != nil {
		logrus.Error(err)
		_, _ = ctx.JSON(Fail(http.StatusInternalServerError, "服务器错误", err))
		return
	}
	_, _ = ctx.JSON(Success("", accountDTO))
}

// 发起转账
func transfer(ctx context.Context) {
	accountTransferDTO := service.AccountTransferDTO{}
	err := ctx.ReadJSON(&accountTransferDTO)
	if err != nil {
		logrus.Error(err)
		_, _ = ctx.JSON(Fail(http.StatusBadRequest, "数据解析错误", nil))
		return
	}
	tx := accountDb.Begin()
	err = starter.Transaction(tx, func() error {
		accountService := accounts.GetAccountService(tx)
		_, transferError := accountService.Transfer(accountTransferDTO)
		return transferError
	})
	if err != nil {
		logrus.Error(err)
		_, _ = ctx.JSON(Fail(http.StatusInternalServerError, "转账失败", err))
		return
	}
	_, _ = ctx.JSON(Success("转账成功", nil))
}

// 向指定账户储值
func storeValue(ctx context.Context) {
	accountTransferDTO := service.AccountTransferDTO{}
	err := ctx.ReadJSON(&accountTransferDTO)
	if err != nil {
		logrus.Error(err)
		_, _ = ctx.JSON(Fail(http.StatusBadRequest, "数据解析错误", nil))
		return
	}
	tx := accountDb.Begin()
	err = starter.Transaction(tx, func() error {
		accountService := accounts.GetAccountService(tx)
		_, storeError := accountService.StoreValue(accountTransferDTO)
		return storeError
	})
	if err != nil {
		logrus.Error(err)
		_, _ = ctx.JSON(Fail(http.StatusInternalServerError, "充值失败", nil))
		return
	}
	_, _ = ctx.JSON(Success("充值成功", nil))
}

// 查询账户信息
func findAccount(ctx context.Context) {
	accountNo := ctx.Params().GetString("accountNo")
	accountDTO, err := accountService.GetAccountByAccountNo(accountNo)
	if err != nil {
		logrus.Error(err)
		_, _ = ctx.JSON(Fail(http.StatusInternalServerError, "服务器错误", nil))
		return
	}
	_, _ = ctx.JSON(Success("", accountDTO))
}
