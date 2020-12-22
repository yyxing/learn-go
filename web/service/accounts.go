package service

import (
	"github.com/shopspring/decimal"
	"learn-go/web/enums"
	"time"
)

type AccountService interface {
	// 创建账户
	CreateAccount(AccountCreatedDTO) (*AccountDTO, error)
	// 转账
	Transfer(AccountTransferDTO) (enums.TransferStatus, error)
	// 储值
	StoreValue(AccountTransferDTO) (enums.TransferStatus, error)
	// 获取账户信息
	GetEnvelopeAccountByUserId(AccountDTO) ([]*AccountDTO, error)
	// 获取账户信息
	GetAccountByAccountNo(string) (*AccountDTO, error)
}

// 交易的参与者
type TradeParticipator struct {
	AccountNo string `validate:"required"`
	UserId    string `validate:"required"`
	Username  string `validate:"required"`
}

// 转账入参
type AccountTransferDTO struct {
	TradeAccount        TradeParticipator `validate:"required"`         //交易账户
	CounterpartyAccount TradeParticipator `validate:"required"`         //对方账户
	Amount              string            `validate:"required,numeric"` //交易金额,该交易涉及的金额
	TransferType        enums.TradeType   `validate:"required,numeric"` //流水交易类型，0 创建账户，>0 为收入类型，<0 为支出类型，自定义
	TradeDesc           string            //交易描述
}

// 账户创建参数
type AccountCreatedDTO struct {
	Username     string            `validate:"required"` //用户名称
	AccountName  string            `validate:"required"` //账户名称,用来说明账户的简短描述,账户对应的名称或者命名，比如xxx积分、xxx零钱
	AccountType  enums.AccountType //账户类型，用来区分不同类型的账户：积分账户、会员卡账户、钱包账户、红包账户
	CurrencyCode string            //货币类型编码：CNY人民币，EUR欧元，USD美元
	Amount       string            `validate:"required,numeric"` //账户可用余额
}

// 账户创建成功后的返回值
type AccountDTO struct {
	AccountCreatedDTO
	UserId    string              //用户编号, 账户所属用户
	AccountNo string              //账户编号,账户唯一标识
	Status    enums.AccountStatus //账户状态，账户状态：0账户初始化，1启用，2停用
	Balance   decimal.Decimal     //账户状态，账户状态：0账户初始化，1启用，2停用
	CreatedAt time.Time
	UpdatedAt time.Time
}

//账户流水返回信息
type AccountLogDTO struct {
	LogNo           string            //流水编号 全局不重复字符或数字，唯一性标识
	TradeNo         string            //交易单号 全局不重复字符或数字，唯一性标识
	AccountNo       string            //账户编号 账户ID
	TargetAccountNo string            //账户编号 账户ID
	UserId          string            //用户编号
	Username        string            //用户名称
	TargetUserId    string            //目标用户编号
	TargetUsername  string            //目标用户名称
	Amount          decimal.Decimal   //交易金额,该交易涉及的金额
	Balance         decimal.Decimal   //交易后余额,该交易后的余额
	TransferType    enums.TradeType   //流水交易类型，0 创建账户，>0 为收入类型，<0 为支出类型，自定义
	Status          enums.TradeStatus //交易状态：
	Decs            string            //交易描述
	CreatedAt       time.Time         //创建时间
}
