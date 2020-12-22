package accounts

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"learn-go/web/core/starter"
	"learn-go/web/enums"
	"learn-go/web/service"
	"learn-go/web/util"
	"sync"
)

var (
	db          *gorm.DB
	once        sync.Once
	serviceImpl *accountService
)

type accountService struct {
	tx *gorm.DB
}

func GetAccountService(tx *gorm.DB) service.AccountService {
	mutex := sync.Mutex{}
	mutex.Lock()
	if serviceImpl == nil {
		once.Do(func() {
			serviceImpl = &accountService{tx: tx}
		})
		return serviceImpl
	}
	mutex.Unlock()
	serviceImpl.tx = tx
	return serviceImpl
}

// 创建账户
func (svc *accountService) CreateAccount(createDTO service.AccountCreatedDTO) (*service.AccountDTO, error) {
	// 数据校验
	err := starter.ParamValidate(createDTO)
	if err != nil {
		return nil, err
	}
	switch createDTO.AccountType {
	case enums.DebitCard:
	case enums.CreditCard:
	case enums.ForeignCurrencyCard:
	default:
		return nil, accountError{"账户类型不合法"}
	}
	// 数据拼接
	account := Account{}
	accountDTO := service.AccountDTO{
		AccountCreatedDTO: createDTO,
		UserId:            util.NextId(),
		AccountNo:         util.GenerateAccountNo(),
		Status:            enums.AccountActivated,
	}
	err = svc.tx.Transaction(func(tx *gorm.DB) error {
		// 开启事务
		accountDao := AccountDao{Db: tx}
		err = account.FromDTO(&accountDTO)
		if err != nil {
			return err
		}
		result, _ := accountDao.Insert(account)
		if !result {
			return accountError{"创建账户失败，请稍后再试！"}
		}
		logDao := AccountLogDao{Db: tx}
		accountLog := AccountLog{
			LogNo:                 util.NextId(),
			TradeNo:               util.NextId(),
			AccountNo:             account.AccountNo,
			UserId:                account.UserId,
			Username:              account.Username,
			CounterpartyAccountNo: account.AccountNo,
			CounterpartyUserId:    account.UserId,
			CounterpartyUsername:  account.Username,
			Amount:                account.Balance,
			Balance:               account.Balance,
			TradeType:             enums.AccountCreated,
			Status:                enums.TradeSuccess,
			Desc:                  "账户创建",
		}
		result, _ = logDao.Insert(accountLog)
		if !result {
			return accountError{"创建账户失败，请稍后再试！"}
		}
		account = *accountDao.getByAccountNo(account.AccountNo)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return account.ToDTO(), nil
}
func transfer(dto service.AccountTransferDTO, tx *gorm.DB) (enums.TransferStatus, error) {
	amount, err := decimal.NewFromString(dto.Amount)
	if err != nil {
		return enums.TransferFailure, err
	}
	if dto.TransferType < 0 {
		amount = amount.Neg()
	}
	accountDao := AccountDao{Db: tx}
	// 账户余额操作
	result, err := accountDao.UpdateBalance(dto.TradeAccount.AccountNo, amount)
	if err != nil {
		return enums.TransferFailure, err
	}
	if !result && dto.TransferType < 0 {
		return enums.TransferSufficientFunds, accountError{"账户余额不足"}
	}
	// 转账后获取账户信息
	account := accountDao.getByAccountNo(dto.TradeAccount.AccountNo)
	if account == nil {
		return enums.TransferFailure, accountError{"账户异常"}
	}
	// 自身流水记录 自身为支出
	accountLog := AccountLog{}
	accountLog.FromDTO(&dto)
	accountLog.TradeNo = util.NextId()
	accountLog.LogNo = util.NextId()
	accountLog.Amount = amount
	accountLog.Balance = account.Balance
	accountLog.Status = enums.TradeSuccess
	accountLogDao := AccountLogDao{Db: tx}
	result, err = accountLogDao.Insert(accountLog)
	if !result {
		return enums.TransferFailure, accountError{"记录交易流水失败"}
	}
	if err != nil {
		return enums.TransferFailure, err
	}
	return enums.TransferSuccess, nil
}

// 转账
func (svc *accountService) Transfer(dto service.AccountTransferDTO) (enums.TransferStatus, error) {
	err := starter.ParamValidate(dto)
	if err != nil {
		return enums.TransferFailure, err
	}
	var status enums.TransferStatus
	err = svc.tx.Transaction(func(tx *gorm.DB) error {
		// 自身扣钱
		// 交易账号为自身 对手账号为收入账号
		transferStatus, transferError := transfer(dto, tx)
		// 目标加钱
		if transferError == nil {
			dto.TransferType = enums.TransferIncoming
			dto.TradeDesc = "转账收入"
			// 交换交易账号为收款人账号 对手账号为支出账号
			dto.TradeAccount, dto.CounterpartyAccount = dto.CounterpartyAccount, dto.TradeAccount
			transferStatus, transferError = transfer(dto, tx)
		}
		status = transferStatus
		return transferError
	})
	return status, err
}

// 充值
func (svc *accountService) StoreValue(dto service.AccountTransferDTO) (enums.TransferStatus, error) {
	dto.CounterpartyAccount = dto.TradeAccount
	dto.TransferType = enums.AccountStoreValue
	err := starter.ParamValidate(dto)
	if err != nil {
		return enums.TransferFailure, err
	}
	var status enums.TransferStatus
	err = svc.tx.Transaction(func(tx *gorm.DB) error {
		transferStatus, transferError := transfer(dto, tx)
		status = transferStatus
		return transferError
	})
	return status, err
}

// 获取账户信息
func (svc *accountService) GetEnvelopeAccountByUserId(dto service.AccountDTO) ([]*service.AccountDTO, error) {
	accountDao := AccountDao{svc.tx}
	accountParam := Account{}
	err := accountParam.FromDTO(&dto)
	if err != nil {
		return nil, err
	}
	accounts := accountDao.Find(accountParam)
	var results = make([]*service.AccountDTO, len(accounts))
	for i, account := range accounts {
		if account == nil {
			return nil, accountError{dto.UserId + "账户异常，请与客服联系"}
		}
		results[i] = account.ToDTO()
	}
	return results, nil
}

// 获取账户信息
func (svc *accountService) GetAccountByAccountNo(accountNo string) (*service.AccountDTO, error) {
	if accountNo == "" || len(accountNo) == 0 {
		return nil, accountError{"账号不能为空"}
	}
	accountDao := AccountDao{Db: svc.tx}
	account := accountDao.getByAccountNo(accountNo)
	if account == nil {
		return nil, accountError{"查询不到指定账户，请与客服联系。"}
	}
	return account.ToDTO(), nil
}
