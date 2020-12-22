package enums

type RedEnvelopeType int8

const (
	Normal = iota + 1
	TryLuck
)

type OrderStatus int8

const (
	OrderCreate  OrderStatus = 1
	OrderSuccess OrderStatus = 2
	OrderExpired OrderStatus = 3
	OrderFailure OrderStatus = -1
)

type OrderType int8

const (
	// 创建
	Create OrderType = 1
	// 退款
	Refund OrderType = -1
)
