package models

import (
	"blockshop/common"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type GoodsType struct {
	BaseModel
	Id           int64     `orm:"pk;column(id);auto;size(11)" description:"商品类别ID" json:"id"`
	Name         string    `orm:"column(name);size(512);index" description:"分类名称" json:"name"`
	Icon         string    `orm:"column(icon);size(150);default(/static/upload/default/user-default-60x60.png)" description:"分类Icon" json:"icon"`
	IsShow       int8      `orm:"column(is_show);default(0)" description:"是否显示" json:"is_show"`   // 0 显示 1 不显示
}

func (this *GoodsType) TableName() string {
	return common.TableName("goods_type")
}

func (this *GoodsType) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *GoodsType)SearchField() []string{
	return []string{"name"}
}

func (this *GoodsType) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *GoodsType) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *GoodsType) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *GoodsType) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}

func GetGdsTypeList() ([]*GoodsType) {
	gdtp_list := make([]*GoodsType, 0)
	_, err := orm.NewOrm().QueryTable(GoodsType{}).OrderBy("-id").All(&gdtp_list)
	if err != nil {
		return nil
	}
	return gdtp_list
}

func GetGdsTypeByName(type_name string) *GoodsType {
	var goods_cat GoodsType
	err := orm.NewOrm().QueryTable(GoodsType{}).Filter("name", type_name).One(&goods_cat)
	if err != nil {
		return nil
	}
	return &goods_cat
}

func GetGdsTypeById(id int64) *GoodsType {
	var goods_cat GoodsType
	err := orm.NewOrm().QueryTable(GoodsType{}).Filter("id", id).One(&goods_cat)
	if err != nil {
		return nil
	}
	return &goods_cat
}
