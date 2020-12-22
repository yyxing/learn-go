package envelopes

import (
	"fmt"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"learn-go/web/core/starter"
	"learn-go/web/domain/accounts"
	"learn-go/web/enums"
	"learn-go/web/service"
	"learn-go/web/util"
	"sync"
	"time"
)

var (
	db          *gorm.DB
	once        sync.Once
	serviceImpl *envelopeService
)

type envelopeService struct {
	tx          *gorm.DB
	envelopeDao *EnvelopeDao
}

func GetEnvelopeService(tx *gorm.DB) service.EnvelopeService {
	mutex := sync.Mutex{}
	mutex.Lock()
	if serviceImpl == nil {
		once.Do(func() {
			serviceImpl = new(envelopeService)
			serviceImpl.envelopeDao = &EnvelopeDao{}
			serviceImpl.tx = tx
			serviceImpl.envelopeDao.tx = tx
		})
		return serviceImpl
	}
	mutex.Unlock()
	serviceImpl.tx = tx
	serviceImpl.envelopeDao.tx = tx
	return serviceImpl
}

// 发红包
func (svc *envelopeService) SendEnvelope(dto service.RedEnvelopeSendDTO) (*service.RedEnvelopeDTO, error) {
	err := starter.ParamValidate(dto)
	if err != nil {
		return nil, err
	}
	redEnvelope := RedEnvelopeGood{}
	err = redEnvelope.fromDTO(dto)
	if err != nil {
		return nil, err
	}
	err = svc.tx.Transaction(func(tx *gorm.DB) error {
		envelopeDao := svc.envelopeDao
		result, err := envelopeDao.Insert(redEnvelope)
		if err != nil {
			return err
		}
		if !result {
			return envelopeError{"网络异常，请稍后再试！"}
		}
		// 转账操作 将用户发的红包转到系统账户中
		accountService := accounts.GetAccountService(svc.tx)
		systemAccount, err := accountService.GetAccountByAccountNo(util.SystemAccountNo)
		if err != nil {
			return err
		}
		_, err = accountService.Transfer(service.AccountTransferDTO{
			TradeAccount: service.TradeParticipator{
				AccountNo: dto.AccountNo,
				UserId:    dto.UserId,
				Username:  dto.Username,
			},
			CounterpartyAccount: service.TradeParticipator{
				AccountNo: systemAccount.AccountNo,
				UserId:    systemAccount.UserId,
				Username:  systemAccount.Username,
			},
			Amount:       dto.Amount,
			TransferType: enums.TransferOutgoing,
			TradeDesc:    fmt.Sprintf("%s发红包啦", dto.Username),
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return svc.GetOneEnvelope(redEnvelope.EnvelopeNo)
}

// 收红包
func (svc *envelopeService) ReceiveEnvelope(dto service.RedEnvelopeReceiveDTO) (*service.RedEnvelopeItemDTO, error) {
	err := starter.ParamValidate(dto)
	if err != nil {
		return nil, err
	}
	redEnvelopeItem := RedEnvelopeItem{}
	redEnvelopeItem.fromDTO(dto)
	envelopeDao := svc.envelopeDao
	// 获取红包详情
	redEnvelopeGood := envelopeDao.GetOne(dto.EnvelopeNo)
	if redEnvelopeGood.RemainAmount.Equal(decimal.NewFromFloat(0)) || redEnvelopeGood.RemainQuantity == 0 {
		return nil, envelopeError{"红包领完啦，下次早点来哦！"}
	}
	err = svc.tx.Transaction(func(tx *gorm.DB) error {

		if redEnvelopeGood == nil {
			return envelopeError{"要领取的红包不存在"}
		}
		if redEnvelopeGood.ExpiredAt.Before(time.Now()) {
			return envelopeError{"要领取的红包已过期"}
		}
		// 红宝金额生成
		receiveAmount := util.GenerateRedEnvelopeAmount(int64(redEnvelopeGood.RemainQuantity),
			redEnvelopeGood.RemainAmount, redEnvelopeGood.EnvelopeType)
		// 扣除金额和数量
		result, err := envelopeDao.ReceiveRedEnvelope(dto.EnvelopeNo, receiveAmount)
		if !result {
			return envelopeError{"网络异常，请稍后再领取！"}
		}
		if err != nil {
			return err
		}
		// 金额转账
		accountService := accounts.GetAccountService(tx)
		systemAccount, err := accountService.GetAccountByAccountNo(util.SystemAccountNo)
		if err != nil {
			return err
		}
		_, err = accountService.Transfer(service.AccountTransferDTO{
			TradeAccount: service.TradeParticipator{
				AccountNo: systemAccount.AccountNo,
				UserId:    systemAccount.UserId,
				Username:  systemAccount.Username,
			},
			CounterpartyAccount: service.TradeParticipator{
				AccountNo: dto.AccountNo,
				UserId:    dto.ReceiveUserId,
				Username:  dto.ReceiveUsername,
			},
			Amount:       receiveAmount.String(),
			TransferType: enums.TransferOutgoing,
			TradeDesc:    fmt.Sprintf("系统向%s发送来自%s的红包", dto.ReceiveUsername, redEnvelopeGood.Username.String),
		})
		redEnvelopeItem.Amount = receiveAmount
		redEnvelopeItem.Quantity = 1
		redEnvelopeItem.RemainAmount = redEnvelopeGood.RemainAmount.Sub(receiveAmount)
		redEnvelopeItem.RemainQuantity = redEnvelopeGood.RemainQuantity - 1
		if err != nil {
			return err
		}
		// 添加红包领取详情
		result, err = envelopeDao.InsertReceiveRedEnvelopeLog(redEnvelopeItem)
		if !result {
			return envelopeError{"网络异常，请稍后再领取！"}
		}
		return err
	})

	return envelopeDao.FindReceiveByEnvelopeNo(dto.EnvelopeNo, dto.ReceiveUserId).toDTO(), err
}

// 查询单个红包详情
func (svc *envelopeService) GetOneEnvelope(envelopeNo string) (*service.RedEnvelopeDTO, error) {
	envelopeDao := svc.envelopeDao
	envelopeGood := envelopeDao.GetOne(envelopeNo)
	if envelopeGood == nil {
		return nil, envelopeError{"不存在该红包信息"}
	}
	receiveRedEnvelopes := envelopeDao.FindOneRedReceive(envelopeNo)
	receiveRedEnvelopeItems := make([]*service.RedEnvelopeItemDTO, len(receiveRedEnvelopes))
	for i, envelope := range receiveRedEnvelopes {
		receiveRedEnvelopeItems[i] = envelope.toDTO()
	}
	envelopeDTO := envelopeGood.toDTO()
	envelopeDTO.ReceiveItems = receiveRedEnvelopeItems
	return envelopeDTO, nil
}

// 分页查询已经领取的红包
func (svc *envelopeService) FindReceiveEnvelopes(userId string, pageParam util.PageParam) *util.Page {
	envelopeDao := svc.envelopeDao
	page := pageParam.Page()
	redEnvelopeItems := envelopeDao.FindUserReceive(userId, page)
	page.Records = redEnvelopeItems
	return page
}

// 分页查询已经发送的红包
func (svc *envelopeService) FindSendEnvelopes(userId string, pageParam util.PageParam) *util.Page {
	envelopeDao := svc.envelopeDao
	page := pageParam.Page()
	redEnvelopeItems := envelopeDao.FindSendByUserId(userId, page)
	page.Records = redEnvelopeItems
	return page
}

// 分页查询可领的红包
func (svc *envelopeService) FindEnvelopes(pageParam util.PageParam) *util.Page {
	envelopeDao := svc.envelopeDao
	page := pageParam.Page()
	redEnvelopeItems := envelopeDao.FindRedEnvelopes(page)
	page.Records = redEnvelopeItems
	return page
}

// 退款流程
func (svc *envelopeService) Refund(envelopeNo string) (*service.RedEnvelopeDTO, error) {
	envelopeDao := svc.envelopeDao
	envelopeGood := envelopeDao.GetOne(envelopeNo)
	// 发起退款流程
	accountService := accounts.GetAccountService(svc.tx)
	// 查找系统账户
	systemAccount, err := accountService.GetAccountByAccountNo(util.SystemAccountNo)
	if err != nil {
		return nil, err
	}
	_, err = accountService.Transfer(service.AccountTransferDTO{
		TradeAccount: service.TradeParticipator{
			AccountNo: systemAccount.AccountNo,
			UserId:    systemAccount.UserId,
			Username:  systemAccount.Username,
		},
		CounterpartyAccount: service.TradeParticipator{
			AccountNo: envelopeGood.AccountNo,
			UserId:    envelopeGood.UserId,
			Username:  envelopeGood.Username.String,
		},
		Amount:       envelopeGood.RemainAmount.String(),
		TransferType: enums.EnvelopExpiredRefund,
		TradeDesc:    fmt.Sprintf("%s的红包退还到%s账户", envelopeGood.Username.String, envelopeGood.AccountNo),
	})
	return envelopeGood.toDTO(), err
}
