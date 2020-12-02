package service

type TransferStatus int8

const (
	// 转账失败
	TransferFailure TransferStatus = -1
	// 余额不足
	TransferSufficientFunds TransferStatus = -2
	// 转账成功
	TransferSuccess TransferStatus = 1
)

// 转账类型0 创建账户 1 入账 -1 支出
type TransferType int8

const (
	// 账户创建
	AccountCreated TransferType = 100
	// 自己向个人账户中充值
	AccountStoreValue TransferType = 1
	// 通过转账支出
	TransferOutgoing TransferType = -1
	// 通过转账收入
	TransferIncoming TransferType = 2
	// 转账退款
	EnvelopExpiredRefund TransferType = 3
)

type AccountChangeType int8

const (
	// 用于资金列表详情时标识
	FlagAccountCreated AccountChangeType = 1
	// 资金收入
	AccountIncoming AccountChangeType = 2
	// 资金支出
	AccountOutgoing AccountChangeType = -1
)
