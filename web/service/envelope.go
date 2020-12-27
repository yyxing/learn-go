package service

import (
	"github.com/shopspring/decimal"
	"learn-go/web/enums"
	"learn-go/web/util"
	"time"
)

// 订单服务
type EnvelopeService interface {
	// 发红包
	SendEnvelope(RedEnvelopeSendDTO) (*RedEnvelopeDTO, error)
	// 收红包
	ReceiveEnvelope(RedEnvelopeReceiveDTO) (*RedEnvelopeItemDTO, error)
	// 查询单个红包信息
	GetOneEnvelope(string) (*RedEnvelopeDTO, error)
	// 查询用户已经领取红包列表
	FindReceiveEnvelopes(string, util.PageParam) *util.Page
	// 查询用户发送的红包列表
	FindSendEnvelopes(string, util.PageParam) *util.Page
	// 查询可领红包列表
	FindEnvelopes(util.PageParam) *util.Page
	// 退款
	Refund(string) (*RedEnvelopeDTO, error)
	// 查询所有的过期红包
	FindAllExpiredRedEnvelope() []*RedEnvelopeDTO
}

// 发送红包DTO
type RedEnvelopeSendDTO struct {
	//红包类型：普通红包，碰运气红包
	EnvelopeType enums.RedEnvelopeType ` validate:"required"`
	//用户名称
	Username string `validate:"required"`
	//用户编号, 红包所属用户
	UserId string `validate:"required"`
	//祝福语
	Blessing string
	//红包金额:普通红包指单个红包金额，碰运气红包指总金额
	Amount string `validate:"required,numeric"`
	//红包总数量
	Quantity int `validate:"required,numeric"`
	// 发送红包的账户
	AccountNo string `validate:"required"`
}

// 接收红包DTO
type RedEnvelopeReceiveDTO struct {
	//红包编号,红包唯一标识
	EnvelopeNo string `validate:"required"`
	//红包接收者用户名称
	ReceiveUsername string ` validate:"required"`
	//红包接收者用户编号
	ReceiveUserId string `validate:"required"`
	// 接收者账户编号
	AccountNo string `validate:"required"`
}

// 查询DTO
type RedEnvelopeDTO struct {
	// 红包编号
	EnvelopeNo string `json:"envelopeNo"`
	//所属用户名称
	Username string `json:"username"`
	//所属用户编号, 红包所属用户
	UserId string `json:"userId"`
	// 发送者账号
	AccountNO string `json:"accountNO"`
	//祝福语
	Blessing string `json:"blessing"`
	//红包总金额
	Amount decimal.Decimal `json:"amount"`
	//红包总数量
	Quantity int `json:"quantity"`
	//红包剩余金额
	RemainAmount decimal.Decimal `json:"remainAmount"`
	//红包剩余数量
	RemainQuantity int `json:"remainQuantity"`
	//过期时间
	ExpiredAt time.Time `json:"expiredAt" `
	// 红包类型
	EnvelopeType enums.RedEnvelopeType
	// 已领取红包列表
	ReceiveItems []*RedEnvelopeItemDTO `json:"receiveItems"`
	//创建时间
	CreatedAt time.Time `json:"createdAt"`
}

type RedEnvelopeItemDTO struct {
	// 领取详情编号
	ItemNo string `json:"itemNo"`
	//领取金额
	ReceiveAmount decimal.Decimal `json:"receiveAmount"`
	//领取数量
	ReceiveQuantity int `json:"receiveQuantity"`
	//创建时间
	ReceiveTime time.Time `json:"receiveTime"`
	//领取用户名
	ReceiveUserName string `json:"receiveUserName"`
	//领取用户编号
	ReceiveUserId string `json:"receiveUserId"`
	// 接收红包的账户编号
	ReceiveAccountNo string `json:"receiveAccountNo"`
}
