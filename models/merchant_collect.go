package models

import (
	"blockshop/common"
	"github.com/astaxie/beego/orm"
)

type MerchantCollect struct {
	BaseModel
	Id             int64      `orm:"pk;column(id);auto;size(11)" description:"ID" json:"id"`
	UserId         int64      `orm:"column(user_id);index" description:"用户ID" json:"user_id"`
	CtMctId        int64      `orm:"column(ct_mct_id);index" description:"收藏的商家ID" json:"ct_mct_id"`
}

func (this *MerchantCollect) TableName() string {
	return common.TableName("merchant_collect")
}

func (this *MerchantCollect) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *MerchantCollect) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *MerchantCollect) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *MerchantCollect) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}
