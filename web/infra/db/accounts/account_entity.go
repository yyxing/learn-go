package accounts

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"learn-go/web/service"
	"time"
)

type Account struct {
	ID           uint64 `gorm:"primary_key"`
	AccountNo    string `gorm:"unique"`
	AccountName  string `gorm:"not null"`
	AccountType  int    `gorm:"not null"`
	CurrencyCode string `gorm:"not null"`
	UserId       string `gorm:"index;not null"`
	Username     sql.NullString
	Balance      decimal.Decimal `gorm:"not null;"`
	Status       int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (Account) TableName() string {
	return "account"
}
func (do *Account) ToDTO() *service.AccountDTO {
	dto := &service.AccountDTO{
		AccountCreatedDTO: service.AccountCreatedDTO{
			UserId:       do.UserId,
			Username:     do.Username.String,
			AccountName:  do.AccountName,
			AccountType:  do.AccountType,
			CurrencyCode: do.CurrencyCode,
			Balance:      do.Balance.String(),
		},
		AccountNo: do.AccountNo,
		CreatedAt: do.CreatedAt,
		UpdatedAt: do.UpdatedAt,
	}
	return dto
}

func (do *Account) FromDTO(dto *service.AccountDTO) {
	if dto == nil {
		logrus.Error("AccountDTO is nil ")
		return
	}
	balance, err := decimal.NewFromString(dto.Balance)
	if err != nil {
		logrus.Error(err)
		return
	}
	do.Balance = balance
	do.CurrencyCode = dto.CurrencyCode
	do.AccountType = dto.AccountType
	do.AccountName = dto.AccountName
	do.AccountNo = dto.AccountNo
	do.Username = sql.NullString{
		String: dto.Username,
		Valid:  true,
	}
	do.UserId = dto.UserId
	do.Status = dto.Status
}
