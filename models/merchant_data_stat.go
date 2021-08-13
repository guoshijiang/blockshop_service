package models

import (
	"blockshop/common"
	"github.com/astaxie/beego/orm"
)

type MerchantDataStat struct {
	BaseModel
	Id             int64         `orm:"pk;column(id);auto;size(11)" description:"统计ID" json:"id"`
	UserId         int64         `orm:"column(user_id);index" description:"用户ID" json:"user_id"`
	Desc           string        `orm:"column(desc);index" description:"统计识别符号" json:"desc"`
	Val            int64         `orm:"column(val);index" description:"统计值" json:"val"`
}

func (this *MerchantDataStat) TableName() string {
	return common.TableName("merchant_data_stat")
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

func (this *MerchantDataStat) InsertMany(data []*MerchantDataStat)(err error,id int64) {
  if id,err = orm.NewOrm().InsertMulti(1,data);err != nil {
    return  err,0
  }
  return nil,id
}
