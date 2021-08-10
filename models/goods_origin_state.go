package models

import (
	"blockshop/common"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type GoodsOriginState struct {
	BaseModel
	Id           int64     `orm:"pk;column(id);auto;size(11)" description:"国家ID" json:"id"`
	Name         string    `orm:"column(name);size(512);index" description:"国家名称" json:"name"`
	IsShow       int8      `orm:"column(is_show);default(0)" description:"是否显示" json:"is_show"`   // 0 显示 1 不显示
}

func (this *GoodsOriginState) TableName() string {
	return common.TableName("goods_origin_state")
}

func (this *GoodsOriginState) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *GoodsOriginState)SearchField() []string{
	return []string{"name"}
}

func (this *GoodsOriginState) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *GoodsOriginState) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *GoodsOriginState) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *GoodsOriginState) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}

func GetOriginStateList() ([]*GoodsOriginState) {
	gog_list := make([]*GoodsOriginState, 0)
	_, err := orm.NewOrm().QueryTable(GoodsOriginState{}).OrderBy("-id").All(&gog_list)
	if err != nil {
		return nil
	}
	return gog_list
}


func GetGdsOsByName(os_name string) *GoodsOriginState {
	var goods_ost GoodsOriginState
	err := orm.NewOrm().QueryTable(GoodsOriginState{}).Filter("name", os_name).One(&goods_ost)
	if err != nil {
		return nil
	}
	return &goods_ost
}
