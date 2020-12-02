package accounts

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AccountLogDao struct {
	tx *gorm.DB
}

// 插入交易流水
func (dao *AccountLogDao) Insert(log AccountLog) (bool, error) {
	tx := dao.tx
	if err := tx.Create(&log).Error; err != nil {
		logrus.Errorf("create account log error. errorMessage: %s", err)
		return false, err
	}
	return true, nil
}

// 根据流水编号查询
func (dao *AccountLogDao) GetByLogNo(logNo string) *AccountLog {
	tx := dao.tx
	accountLog := &AccountLog{LogNo: logNo}
	if err := tx.Find(accountLog).Error; err != nil {
		logrus.Errorf("create account log error. errorMessage: %s", err)
		return nil
	}
	return accountLog
}

// 根据交易编号查询
func (dao *AccountLogDao) GetByTradeNo(tradeNo string) *AccountLog {
	tx := dao.tx
	accountLog := &AccountLog{TradeNo: tradeNo}
	if err := tx.Find(accountLog).Error; err != nil {
		logrus.Errorf("create account log error. errorMessage: %s", err)
		return nil
	}
	return accountLog
}
