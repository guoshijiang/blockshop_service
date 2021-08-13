package models


import (
	"blockshop/common"
	"github.com/astaxie/beego/orm"
)

type MerchantConfig struct {
	BaseModel
	Id             int64         `orm:"pk;column(id);auto;size(11)" description:"配置ID" json:"id"`
	BtcAmount      float64       `orm:"pk;column(btc_amount);default(0)" description:"BTC数量" json:"btc_amount"`
	UsdtAmount     float64       `orm:"pk;column(usdt_amount);default(0)" description:"BTC数量" json:"usdt_amount"`
	ConfigType     int8    		 `orm:"pk;column(config_type);default(0)" description:"BTC数量" json:"config_type"` //0:da
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
