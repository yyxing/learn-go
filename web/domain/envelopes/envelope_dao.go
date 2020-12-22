package envelopes

import (
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"learn-go/web/util"
	"time"
)

type EnvelopeDao struct {
	tx *gorm.DB
}

// 发红包
func (dao *EnvelopeDao) Insert(envelopeGood RedEnvelopeGood) (bool, error) {
	if err := dao.tx.Create(&envelopeGood).Error; err != nil {
		logrus.Error(err)
		return false, err
	}
	return true, nil
}

// 查询指定红包信息
func (dao *EnvelopeDao) GetOne(envelopeNo string) *RedEnvelopeGood {
	redEnvelopeGood := RedEnvelopeGood{
		EnvelopeNo: envelopeNo,
	}
	if err := dao.tx.Where(&redEnvelopeGood).Find(&redEnvelopeGood).Error; err != nil {
		logrus.Error(err)
		return nil
	}
	return &redEnvelopeGood
}
func paginate(page *util.Page) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Offset(page.Offset).Limit(page.Size)
	}
}

// 查询指定用户发送的红包
func (dao *EnvelopeDao) FindSendByUserId(userId string, page *util.Page) []*RedEnvelopeGood {
	var sendRedEnvelope []*RedEnvelopeGood
	var totalCount int64
	if err := dao.tx.Scopes(paginate(page)).Where("user_id = ?", userId).Find(&sendRedEnvelope).Error; err != nil {
		logrus.Error(err)
		return nil
	}
	if err := dao.tx.Model(&RedEnvelopeGood{}).Where("user_id", userId).Count(&totalCount).Error; err != nil {
		logrus.Error(err)
		return nil
	}
	page.Total = totalCount
	return sendRedEnvelope
}

// 查询红包的领取记录
func (dao *EnvelopeDao) FindOneRedReceive(envelopeNo string) []*RedEnvelopeItem {
	var receiveRedEnvelopeItems []*RedEnvelopeItem
	if err := dao.tx.Where("envelope_no", envelopeNo).Find(&receiveRedEnvelopeItems).Error; err != nil {
		logrus.Error(err)
		return nil
	}
	return receiveRedEnvelopeItems
}

// 查询用户的红包领取记录
func (dao *EnvelopeDao) FindUserReceive(userId string, page *util.Page) []*RedEnvelopeItem {
	var receiveRedEnvelopeItems []*RedEnvelopeItem
	var totalCount int64
	if err := dao.tx.Scopes(paginate(page)).Where("receive_user_id", userId).
		Find(&receiveRedEnvelopeItems).Error; err != nil {
		logrus.Error(err)
		return nil
	}
	if err := dao.tx.Model(&RedEnvelopeItem{}).Where("receive_user_id", userId).Count(&totalCount).Error; err != nil {
		logrus.Error(err)
		return nil
	}
	page.Total = totalCount
	return receiveRedEnvelopeItems
}

// 查询单个用户的某个红包的领取记录
func (dao *EnvelopeDao) FindReceiveByEnvelopeNo(envelopeNo string, userId string) *RedEnvelopeItem {
	var receiveRedEnvelopeItem RedEnvelopeItem
	if err := dao.tx.Where("envelope_no", envelopeNo).Where("receive_user_id", userId).
		Find(&receiveRedEnvelopeItem).Error; err != nil {
		logrus.Error(err)
		return nil
	}
	return &receiveRedEnvelopeItem
}

// 查询指定用户过期的红包
func (dao *EnvelopeDao) FindExpiredByUserId(userId string, page *util.Page) []*RedEnvelopeGood {
	var expiredRedEnvelope []*RedEnvelopeGood
	var totalCount int64
	if err := dao.tx.Scopes(paginate(page)).Where("user_id = ?", userId).
		Where("expired_at <= ?", time.Now()).Find(&expiredRedEnvelope).Error; err != nil {
		logrus.Error(err)
		return nil
	}
	if err := dao.tx.Model(&RedEnvelopeGood{}).Where("user_id = ?", userId).
		Where("expired_at <= ?", time.Now()).Count(&totalCount).Error; err != nil {
		logrus.Error(err)
		return nil
	}
	page.Total = totalCount
	return expiredRedEnvelope
}

// 分页查询可领取的红包
func (dao *EnvelopeDao) FindRedEnvelopes(page *util.Page) []*RedEnvelopeGood {
	var expiredRedEnvelope []*RedEnvelopeGood
	var totalCount int64
	if err := dao.tx.Scopes(paginate(page)).Where("expired_at > ?", time.Now()).
		Find(&expiredRedEnvelope).Error; err != nil {
		logrus.Error(err)
		return nil
	}
	if err := dao.tx.Model(&RedEnvelopeGood{}).Where("expired_at > ?", time.Now()).Count(&totalCount).Error; err != nil {
		logrus.Error(err)
		return nil
	}
	page.Total = totalCount
	return expiredRedEnvelope
}

// 查询所有过期的红包
func (dao *EnvelopeDao) FindExpired(page *util.Page) []*RedEnvelopeGood {
	var expiredRedEnvelope []*RedEnvelopeGood
	var totalCount int64
	if err := dao.tx.Scopes(paginate(page)).Where("expired_at <= ?", time.Now()).
		Find(&expiredRedEnvelope).Error; err != nil {
		logrus.Error(err)
		return nil
	}
	if err := dao.tx.Where("expired_at <= ?", time.Now()).Count(&totalCount).Error; err != nil {
		logrus.Error(err)
		return nil
	}
	page.Total = totalCount
	return expiredRedEnvelope
}

// 领红包 减少金额和库存
func (dao *EnvelopeDao) ReceiveRedEnvelope(envelopeNo string, amount decimal.Decimal) (bool, error) {
	result := dao.tx.Model(&RedEnvelopeGood{}).Where("envelope_no = ?", envelopeNo).
		Where("remain_amount >= ?", amount).Where("remain_quantity > ?", 0).
		UpdateColumns(map[string]interface{}{
			"remain_amount":   gorm.Expr("remain_amount - ?", amount),
			"remain_quantity": gorm.Expr("remain_quantity - ?", 1),
		})
	return result.RowsAffected > 0, result.Error
}

// 领红包 增加数据库记录
func (dao *EnvelopeDao) InsertReceiveRedEnvelopeLog(redEnvelopeItem RedEnvelopeItem) (bool, error) {
	if err := dao.tx.Create(&redEnvelopeItem).Error; err != nil {
		logrus.Error(err)
		return false, err
	}
	return true, nil
}
