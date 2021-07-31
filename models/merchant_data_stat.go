package models

import (
	"blockshop/common"
	"github.com/astaxie/beego/orm"
)

type MerchantDataStat struct {
	BaseModel
	Id             int64         `orm:"pk;column(id);auto;size(11)" description:"统计ID" json:"id"`
	UserId         int64         `orm:"column(user_id);index" description:"用户ID" json:"user_id"`
}

func (this *MerchantDataStat) TableName() string {
	return common.TableName("user")
}

func (this *MerchantDataStat) SearchField() []string {
	return []string{"user_name"}
}

func (this *MerchantDataStat) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *MerchantDataStat) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *MerchantDataStat) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *MerchantDataStat) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}

