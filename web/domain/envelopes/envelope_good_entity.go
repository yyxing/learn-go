package envelopes

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"learn-go/web/enums"
	"learn-go/web/service"
	"learn-go/web/util"
	"time"
)

type RedEnvelopeGood struct {
	ID uint64 `gorm:"primary_key"`
	//红包编号,红包唯一标识
	EnvelopeNo string `gorm:"unique;not null"`
	//红包类型：普通红包，碰运气红包,过期红包
	EnvelopeType enums.RedEnvelopeType `gorm:"not null"`
	//红包发送者用户名称
	Username sql.NullString
	//红包发送者用户编号
	UserId string `gorm:"not null"`
	// 发送者支付账户
	AccountNo string `gorm:"not null"`
	// 祝福语
	Blessing sql.NullString
	// 收到金额
	Amount decimal.Decimal `gorm:"not null"`
	// 红包数量
	Quantity int `gorm:"not null"`
	// 红包剩余金额
	RemainAmount decimal.Decimal `gorm:"not null;"`
	// 红包剩余数量
	RemainQuantity int `gorm:"not null;"`
	// 过期时间
	ExpiredAt time.Time `gorm:"not null;"`
	// 支付状态：1未支付，2支付中，3已支付，-1支付失败
	PayStatus int8 `gorm:"not null"`
	// 红包/订单状态：1 创建、2 发布启用、3过期、-1失效
	OrderStatus enums.OrderStatus `gorm:"not null"`
	// 订单类型：发布单、退款单
	OrderType enums.OrderType `gorm:"not null"`
	// 创建时间
	CreatedAt time.Time
	// 修改时间
	UpdatedAt time.Time
}

func (RedEnvelopeGood) TableName() string {
	return "red_envelope_good"
}

func (do RedEnvelopeGood) toDTO() *service.RedEnvelopeDTO {
	return &service.RedEnvelopeDTO{
		EnvelopeNo:     do.EnvelopeNo,
		Username:       do.Username.String,
		UserId:         do.UserId,
		AccountNO:      do.AccountNo,
		Blessing:       do.Blessing.String,
		Amount:         do.Amount,
		Quantity:       do.Quantity,
		RemainAmount:   do.RemainAmount,
		RemainQuantity: do.RemainQuantity,
		ExpiredAt:      do.ExpiredAt,
		CreatedAt:      do.CreatedAt,
		EnvelopeType:   do.EnvelopeType,
	}
}

func (do *RedEnvelopeGood) fromDTO(dto service.RedEnvelopeSendDTO) error {
	do.Blessing = sql.NullString{
		String: dto.Blessing,
		Valid:  true,
	}
	amount, err := decimal.NewFromString(dto.Amount)
	if err != nil {
		return err
	}
	do.Quantity = dto.Quantity
	do.Amount = amount
	do.UserId = dto.UserId
	do.Username = sql.NullString{
		String: dto.Username,
		Valid:  true,
	}
	do.AccountNo = dto.AccountNo
	do.EnvelopeType = dto.EnvelopeType
	do.EnvelopeNo = util.NextId()
	do.RemainQuantity = dto.Quantity
	do.RemainAmount = amount
	do.OrderType = enums.Create
	do.OrderStatus = enums.OrderCreate
	do.PayStatus = 1
	do.ExpiredAt = time.Now().Add(time.Hour * 24)
	return nil
}
