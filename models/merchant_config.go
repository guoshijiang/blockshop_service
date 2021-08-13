package models


import (
	"blockshop/common"
	"github.com/astaxie/beego/orm"
)

type MerchantConfig struct {
	BaseModel
	Id             int64         `orm:"pk;column(id);auto;size(11)" description:"配置ID" json:"id"`
	BtcAmount      float64       `orm:"pk;column(id);auto;size(11)" description:"开通商家需要的BTC数量" json:"id"`
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
