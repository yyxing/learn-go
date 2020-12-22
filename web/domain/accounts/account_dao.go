package accounts

import (
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AccountDao struct {
	Db *gorm.DB
}

// 根据账户唯一编号获取账户信息
func (dao *AccountDao) getByAccountNo(accountNo string) *Account {
	account := Account{AccountNo: accountNo}
	if err := dao.Db.Where(&account).Find(&account).Error; err != nil {
		logrus.Errorf("getByAccountNo error accountNo: %s errorMessage: %s", accountNo, err)
		return nil
	}
	return &account
}

// 根据用户id和账户类型查询账户信息 可能一个用户下有多个账户
func (dao *AccountDao) Find(account Account) []*Account {
	var accounts []*Account
	if err := dao.Db.Where(&account).Find(&accounts).Error; err != nil {
		logrus.Errorf("find error. paramMap: %v errorMessage: %s", account, err)
		return nil
	}
	return accounts
}

// 插入数据
func (dao *AccountDao) Insert(account Account) (bool, error) {
	if err := dao.Db.Create(&account).Error; err != nil {
		logrus.Errorf("create account error. errorMessage: %s", err)
		return false, err
	}
	return true, nil
}

// 账户余额修改
func (dao *AccountDao) UpdateBalance(accountNo string, amount decimal.Decimal) (bool, error) {
	result := dao.Db.Model(&Account{}).Where("account_no = ?", accountNo).
		Where("balance >= ?", amount.Mul(decimal.NewFromInt(-1))).
		Update("balance", gorm.Expr("balance + ?", amount))
	return result.RowsAffected > 0, result.Error
}

// 账户状态修改
func (dao *AccountDao) UpdateStatus(accountNo string, status int) (bool, error) {
	result := dao.Db.Model(&Account{}).Where("account_no = ?", accountNo).
		Update("status", status)
	return result.RowsAffected > 0, result.Error
}
