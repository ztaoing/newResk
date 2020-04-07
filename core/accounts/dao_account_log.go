package accounts

import (
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
)

type AccountLogDao struct {
	runner *dbx.TxRunner //事务型
}

//通过流水编号查询流水记录
func (dao *AccountLogDao) GetOne(lognNo string) *AccountLog {
	a := &AccountLog{LogNo: lognNo}
	ok, err := dao.runner.GetOne(a)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	if !ok {
		return nil
	}
	return a
}

//通过交易编号查询流水记录
func (dao *AccountLogDao) GetByTradeNo(tradeNo string) *AccountLog {
	sql := "select * from account_log where trade_no=?"
	out := &AccountLog{}
	ok, err := dao.runner.Get(a, sql, tradeNo)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	if !ok {
		return nil
	}
	return out
}

//流水记录的写入
func (dao *AccountLogDao) Insert(l *AccountLog) (id int64, err error) {
	res, err := dao.runner.Insert(l)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}