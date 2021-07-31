package models

import (
	"blockshop/common"
	"github.com/astaxie/beego/orm"
)

type UserDataStat struct {
	BaseModel
	Id             int64         `orm:"pk;column(id);auto;size(11)" description:"统计ID" json:"id"`
	UserId         int64         `orm:"column(user_id);index" description:"用户ID" json:"user_id"`
}

func (this *UserDataStat) TableName() string {
	return common.TableName("user_data_stat")
}

func (this *UserDataStat) SearchField() []string {
	return []string{"user_name"}
}

func (this *UserDataStat) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *UserDataStat) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *UserDataStat) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *UserDataStat) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}

