package accounts

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AccountLogDao struct {
	Db *gorm.DB
}

// 插入交易流水
func (dao *AccountLogDao) Insert(log AccountLog) (bool, error) {
	db := dao.Db
	if err := db.Create(&log).Error; err != nil {
		logrus.Errorf("create account log error. errorMessage: %s", err)
		return false, err
	}
	return true, nil
}

// 根据流水编号查询
func (dao *AccountLogDao) GetByLogNo(logNo string) *AccountLog {
	db := dao.Db
	accountLog := &AccountLog{LogNo: logNo}
	if err := db.Find(accountLog).Error; err != nil {
		logrus.Errorf("create account log error. errorMessage: %s", err)
		return nil
	}
	return accountLog
}

// 根据交易编号查询
func (dao *AccountLogDao) GetByTradeNo(tradeNo string) *AccountLog {
	db := dao.Db
	accountLog := &AccountLog{TradeNo: tradeNo}
	if err := db.Find(accountLog).Error; err != nil {
		logrus.Errorf("create account log error. errorMessage: %s", err)
		return nil
	}
	return accountLog
}
