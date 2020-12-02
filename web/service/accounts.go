package service

import "time"

type AccountService interface {
	// 创建账户
	CreateAccount(AccountCreatedDTO) (*AccountDTO, error)
	// 转账
	Transfer(AccountTransferDTO) (*TransferStatus, error)
	// 储值
	StoreValue(AccountTransferDTO) (*TransferStatus, error)
	// 获取账户信息
	GetEnvelopeAccountByUserId(string)
}

// 交易的参与者
type TradeParticipator struct {
	AccountNo string
	UserId    string
	Username  string
}

// 转账入参
type AccountTransferDTO struct {
	TradeNo      string
	TradeBody    TradeParticipator
	TradeTarget  TradeParticipator
	Amount       string
	TransferType TransferType
	ChangeFlag   AccountChangeType
	TradeDesc    string
}

// 账户创建参数
type AccountCreatedDTO struct {
	UserId       string
	Username     string
	AccountName  string
	AccountType  int
	CurrencyCode string
	Balance      string
}

// 账户创建成功后的返回值
type AccountDTO struct {
	AccountCreatedDTO
	AccountNo string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
