package accounts

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"learn-go/web/service"
	"time"
)

type AccountLog struct {
	Id uint64 `gorm:"primary_key"`
	//流水编号 全局不重复字符或数字，唯一性标识
	LogNo string `gorm:"unique;not null"`
	//交易单号 全局不重复字符或数字，唯一性标识
	TradeNo string `gorm:"not null"`
	//账户编号 账户ID
	AccountNo string `gorm:"not null"`
	//用户编号
	UserId string `gorm:"not null"`
	//用户名称
	Username sql.NullString
	//账户编号 账户ID
	TargetAccountNo string `gorm:"not null"`
	//目标用户编号
	TargetUserId string `gorm:"not null"`
	//目标用户名称
	TargetUsername sql.NullString
	//交易金额,该交易涉及的金额
	Amount decimal.Decimal `gorm:"not null"`
	//交易后余额,该交易后的余额
	Balance decimal.Decimal
	//流水交易类型，100 创建账户，>0 为收入类型，<0 为支出类型，自定义
	TransferType service.TransferType
	//交易变化标识：-1 出账 1为进账，枚举
	ChangeFlag service.AccountChangeType
	//交易状态：
	Status int
	//交易描述
	Desc string
	//创建时间
	CreatedAt time.Time
	// 修改时间
	UpdatedAt time.Time
}

func (AccountLog) TableName() string {
	return "account_log"
}
func (do *AccountLog) FromDTO(dto *service.AccountTransferDTO) {
	if dto == nil {
		logrus.Error("AccountDTO is nil ")
		return
	}
	amount, err := decimal.NewFromString(dto.Amount)
	if err != nil {
		logrus.Error(err)
		return
	}
	do.Amount = amount
	do.TradeNo = dto.TradeNo
	do.TargetAccountNo = dto.TradeTarget.AccountNo
	do.TargetUserId = dto.TradeTarget.UserId
	do.TargetUsername = sql.NullString{
		String: dto.TradeTarget.Username,
		Valid:  true,
	}
	do.AccountNo = dto.TradeBody.AccountNo
	do.Username = sql.NullString{
		String: dto.TradeBody.Username,
		Valid:  true,
	}
	do.UserId = dto.TradeBody.UserId
	do.TransferType = dto.TransferType
	do.ChangeFlag = dto.ChangeFlag
	do.Desc = dto.TradeDesc
}
