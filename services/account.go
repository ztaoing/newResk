package services

import (
	"github.com/shopspring/decimal"
	"time"
)

type AccountService interface {
	CreateAccount(dto AccountCreateDTO) (*AccountDTO, error)     //创建账户接口
	Transfer(dto AccountTransferDTO) (TransferedStatus, error)   //转账接口
	StoreValue(dto AccountTransferDTO) (TransferedStatus, error) //储值接口
	GetEnvelopeAccountByUserId(userId string) *AccountDTO        //查询账户接口
}

//CreateAccount 入参
type AccountCreateDTO struct {
	UserId       string //用户id
	Username     string //用户名称
	AccountName  string //账户名称
	AccountType  int    //账户的类型
	CurrencyCode string //货币类型编码：CNY人民币，EUR欧元，USD美元 。。。
	Amount       string //金额 使用string，防止在传递的过程中丢失精度
}

//CreateAccount 出参
type AccountDTO struct {
	AccountCreateDTO
	AccountNo string          //账户编号
	CreateAt  time.Time       //创建时间
	Balance   decimal.Decimal //账户可用余额
	Status    int             //账户状态，账户状态：0账户初始化，1启用，2停用
	CreatedAt time.Time       //创建时间
	UpdatedAt time.Time       //更新时间
}

//Transfer接口的入参
//--转账的入参
type AccountTransferDTO struct {
	TradeNo     string           //交易的编号
	TradeBody   TradePaticipator //交易的主体
	TradeTarget TradePaticipator //交易的目标
	AmountStr   string           //交易金额
	ChangeType  ChangeType       //交易变化的类型
	ChangeFlag  ChangeFlag       //交易变化的标识
	Desc        string           //交易描述
}

//--交易的参与者
type TradePaticipator struct {
	AccountNo string //交易变化
	UserId    string //交易的id
	UserName  string //交易的名称
}
