package models

import (
	"blockshop/common"
	"github.com/astaxie/beego/orm"
	"time"
)

type MerchantOpenRecord struct {
	BaseModel
	Id             int64      `orm:"pk;column(id);auto;size(11)" description:"ID" json:"id"`
	UserId         int64      `orm:"column(user_id);index" description:"用户ID" json:"user_id"`
	MerchantId     int64      `orm:"column(merchant_id);index" description:"商家ID" json:"merchant_id"`
	PayCoinAmount  float64    `orm:"column(pay_coin_amount);default(0);digits(22);decimals(8)" description:"支付币的数量" json:"pay_coin_amount"`
	PayWay         int8       `orm:"column(pay_way);index" description:"支付方式" json:"pay_way"`  // 0:BTC，1:USDT
	PayAt          *time.Time `orm:"column(pay_at);type(datetime);null" description:"支付时间" json:"pay_at"`
}

func (this *MerchantOpenRecord) TableName() string {
	return common.TableName("merchant_open_record")
}

func (this *MerchantOpenRecord) SearchField() []string {
	return []string{"user_name"}
}

func (this *MerchantOpenRecord) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *MerchantOpenRecord) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *MerchantOpenRecord) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *MerchantOpenRecord) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}

func (this *MerchantOpenRecord) InsertMany(data []*MerchantOpenRecord)(err error,id int64) {
  if id,err = orm.NewOrm().InsertMulti(1,data);err != nil {
    return  err,0
  }
  return nil,id
}
