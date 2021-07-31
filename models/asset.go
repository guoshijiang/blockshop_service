package models

import (
	"blockshop/common"
	"github.com/astaxie/beego/orm"
)

type Asset struct {
	BaseModel
	Id        int64    `orm:"pk;column(id);auto;size(11)" description:"币种ID" json:"id"`
	Name      string   `orm:"column(name);unique;size(11);index" description:"币种名称" json:"name"`
	Unit      int64    `orm:"column(name);default(8)" description:"币种精度" json:"unit"`
}

func (this *Asset) TableName() string {
	return common.TableName("asset")
}

func (this *Asset) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

func (this *Asset) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}
