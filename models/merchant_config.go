package models


import (
	"blockshop/common"
	"github.com/astaxie/beego/orm"
)

type MerchantConfig struct {
	BaseModel
	Id             int64         `orm:"pk;column(id);auto;size(11)" description:"统计ID" json:"id"`
	UserId         int64         `orm:"column(user_id);index" description:"用户ID" json:"user_id"`
}

func (this *MerchantConfig) TableName() string {
	return common.TableName("merchant_config")
}

func (this *MerchantConfig) SearchField() []string {
	return []string{"user_name"}
}

func (this *MerchantConfig) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *MerchantConfig) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}
