package accounts

import (
	"database/sql"
	"errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"learn-go/web/enums"
	"learn-go/web/service"
	"time"
)

type Account struct {
	ID           uint64            `gorm:"primary_key"`
	AccountNo    string            `gorm:"unique"`
	AccountName  string            `gorm:"not null"`
	AccountType  enums.AccountType `gorm:"not null"`
	CurrencyCode string            `gorm:"not null"`
	UserId       string            `gorm:"index;not null"`
	Username     sql.NullString
	Balance      decimal.Decimal `gorm:"not null;"`
	Status       enums.AccountStatus
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (Account) TableName() string {
	return "account"
}
func (do *Account) ToDTO() *service.AccountDTO {
	dto := &service.AccountDTO{
		AccountCreatedDTO: service.AccountCreatedDTO{
			Username:     do.Username.String,
			AccountName:  do.AccountName,
			AccountType:  do.AccountType,
			CurrencyCode: do.CurrencyCode,
		},
		UserId:    do.UserId,
		Balance:   do.Balance,
		AccountNo: do.AccountNo,
		CreatedAt: do.CreatedAt,
		UpdatedAt: do.UpdatedAt,
	}
	return dto
}

func (do *Account) FromDTO(dto *service.AccountDTO) error {
	if dto == nil {
		logrus.Error("AccountDTO is nil ")
		return errors.New("AccountDTO is nil")
	}
	amount, err := decimal.NewFromString(dto.Amount)
	if dto.Amount != "" && len(dto.Amount) > 0 && err != nil {
		logrus.Error(err)
		return err
	}
	do.Balance = amount
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
	return nil
}
