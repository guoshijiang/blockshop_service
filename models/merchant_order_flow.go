package models

import (
	"blockshop/common"
	"github.com/astaxie/beego/orm"
)

type MerchantOrderFlow struct {
	BaseModel
	Id             int64      `orm:"pk;column(id);auto;size(11)" description:"ID" json:"id"`
	MerchantId     int64      `orm:"column(merchant_id);index" description:"商家ID" json:"merchant_id"`
	OrderId        int64      `orm:"column(order_id);index" description:"订单ID" json:"order_id"`
	AssetId        int64 	  `orm:"column(asset_id);index" description:"资产ID" json:"asset_id"`
	CoinAmount     float64    `orm:"column(coin_amount);default(0);digits(22);decimals(8)" description:"币的数量" json:"coin_amount"`
	IsValid        int8       `orm:"column(is_valid);default(0)" description:"是否是有效流水" json:"is_valid"` // 0:有效；1:无效
	IsStat         int8       `orm:"column(is_stat);default(0)" description:"是否已经结算" json:"is_stat"`     // 0:没有；1:已经结算
}

func (this *MerchantOrderFlow) TableName() string {
	return common.TableName("merchant_merchant_flow")
}

func (this *MerchantOrderFlow) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *MerchantOrderFlow) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *MerchantOrderFlow) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *MerchantOrderFlow) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}
