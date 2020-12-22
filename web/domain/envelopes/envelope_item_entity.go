package envelopes

import (
	"database/sql"
	"fmt"
	"github.com/shopspring/decimal"
	"learn-go/web/service"
	"learn-go/web/util"
	"time"
)

type RedEnvelopeItem struct {
	ID uint64 `gorm:"primary_key"`
	//红包订单详情编号
	ItemNo string `gorm:"unique;not null"`
	//红包编号,红包唯一标识
	EnvelopeNo string `gorm:"index;not null"`
	//红包接收者用户名称
	ReceiveUsername sql.NullString
	//红包接收者用户编号
	ReceiveUserId string `gorm:"not null"`
	// 收到金额
	Amount decimal.Decimal `gorm:"not null"`
	// 收到数量
	Quantity int `gorm:"not null"`
	// 收到后红包剩余金额
	RemainAmount decimal.Decimal `gorm:"not null;"`
	// 收到后红包剩余数量
	RemainQuantity int `gorm:"not null;"`
	// 红包接收者账户ID
	AccountNo string `gorm:"not null;"`
	// 支付状态：未支付，支付中，已支付，支付失败
	PayStatus int `gorm:"not null"`
	// 订单描述
	OrderDesc sql.NullString
	// 创建时间
	CreatedAt time.Time
	// 修改时间
	UpdatedAt time.Time
}

func (RedEnvelopeItem) TableName() string {
	return "red_envelope_item"
}

func (do RedEnvelopeItem) toDTO() *service.RedEnvelopeItemDTO {
	return &service.RedEnvelopeItemDTO{
		ItemNo:           do.ItemNo,
		ReceiveAmount:    do.Amount,
		ReceiveQuantity:  do.Quantity,
		ReceiveTime:      do.CreatedAt,
		ReceiveUserName:  do.ReceiveUsername.String,
		ReceiveUserId:    do.ReceiveUserId,
		ReceiveAccountNo: do.AccountNo,
	}
}

func (do *RedEnvelopeItem) fromDTO(dto service.RedEnvelopeReceiveDTO) {
	do.ItemNo = util.NextId()
	do.EnvelopeNo = dto.EnvelopeNo
	do.ReceiveUsername = sql.NullString{String: dto.ReceiveUsername, Valid: true}
	do.ReceiveUserId = dto.ReceiveUserId
	do.AccountNo = dto.AccountNo
	do.PayStatus = 1
	do.OrderDesc = sql.NullString{String: fmt.Sprintf("%s收红包啦", dto.ReceiveUsername), Valid: true}
}
