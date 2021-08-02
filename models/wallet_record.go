package models

import (
	"blockshop/common"
	"github.com/astaxie/beego/orm"
)

type WalletRecord struct {
	BaseModel
	Id          int64     `orm:"pk;column(id);auto;size(11)" description:"记录ID" json:"id"`
	UserId      int64     `orm:"column(user_id);index" description:"用户ID" json:"user_id"`
	AssetId     int64     `orm:"column(asset_id);index" description:"资产ID" json:"asset_id"`
	ChainName   string    `orm:"column(chain_name);default(Bitcoin)" description:"链的名称" json:"chain_name"`
	FromAddress string    `orm:"column(from_address);size(256)" description:"转出地址" json:"from_address"`
	ToAddress   string    `orm:"column(to_address);size(256)" description:"转入地址" json:"to_address"`
	Amount      float64   `orm:"column(amount);default(0);digits(32);decimals(8)" description:"金额" json:"amount"`
	TxHash      string    `orm:"column(tx_hash);size(256)" description:"交易hash" json:"tx_hash"`
	TxFee       float64   `orm:"column(tx_fee);default(0);digits(32);decimals(8)" description:"链上充提手续费" json:"tx_fee"`
	Comment     string    `orm:"column(comment);size(256)" description:"充提备注" json:"comment"`
	Admin       string    `orm:"column(admin);size(256)" description:"处理员" json:"admin"`
	WOrD        int8      `orm:"column(w_or_d);default(0)" description:"充值或提现" json:"w_or_d"` // 0:充值, 1:提现
	Status      int8      `orm:"column(status);default(0)" description:"充提状态" json:"status"`   // 0:审核中(未锁定)；1:交易中 2:已发出 3:成功 4:失败 5:锁定未审核 6:审核通过 7:审核拒绝
}

func (this *WalletRecord) TableName() string {
	return common.TableName("wallet_record")
}

func (this *WalletRecord) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

func (this *WalletRecord) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *WalletRecord) SearchField() []string {
  return []string{"chain_name"}
}
