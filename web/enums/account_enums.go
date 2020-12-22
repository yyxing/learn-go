package enums

// 交易状态
type AccountType int8

const (
	// 借记卡账户
	DebitCard AccountType = iota + 1
	// 信用卡账户
	CreditCard
	// 外币卡
	ForeignCurrencyCard
)

type AccountStatus int8

const (
	AccountActivated AccountStatus = 1
	AccountFrozen    AccountStatus = -1
)

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
type TradeType int8

const (
	// 账户创建
	AccountCreated TradeType = 100
	// 自己向个人账户中充值
	AccountStoreValue TradeType = 2
	// 通过转账支出
	TransferOutgoing TradeType = -1
	// 通过转账收入
	TransferIncoming TradeType = 1
	// 转账退款
	EnvelopExpiredRefund TradeType = 3
)

// 交易状态
type TradeStatus int8

const (
	// 账户创建
	TradeSuccess TradeStatus = 1
	// 通过转账支出
	TradeFailure TradeStatus = -1
)
