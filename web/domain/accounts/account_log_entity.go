package accounts

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"learn-go/web/enums"
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
	//对方账户编号 账户ID
	CounterpartyAccountNo string `gorm:"not null"`
	//目标用户编号
	CounterpartyUserId string `gorm:"not null"`
	//目标用户名称
	CounterpartyUsername sql.NullString
	//交易金额,该交易涉及的金额
	Amount decimal.Decimal `gorm:"not null"`
	//交易后余额,该交易后的余额
	Balance decimal.Decimal
	//流水交易类型，100 创建账户，>0 为收入类型，<0 为支出类型，自定义
	TradeType enums.TradeType
	//交易状态：
	Status enums.TradeStatus
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
	do.CounterpartyAccountNo = dto.CounterpartyAccount.AccountNo
	do.CounterpartyUserId = dto.CounterpartyAccount.UserId
	do.CounterpartyUsername = sql.NullString{
		String: dto.CounterpartyAccount.Username,
		Valid:  true,
	}
	do.AccountNo = dto.TradeAccount.AccountNo
	do.Username = sql.NullString{
		String: dto.TradeAccount.Username,
		Valid:  true,
	}
	do.UserId = dto.TradeAccount.UserId
	do.TradeType = dto.TransferType
	do.Desc = dto.TradeDesc
}
