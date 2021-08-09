package models

import (
	"blockshop/common"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type GoodsCat struct {
	BaseModel
	Id           int64     `orm:"pk;column(id);auto;size(11)" description:"商品类别ID" json:"id"`
    CatLevel     int8      `orm:"default(1)" json:"cat_level"` // 分类级别
    FatherCatId  int64     `json:"father_cat_id"`              // 父级分类 ID
	Name         string    `orm:"column(name);size(512);index" description:"分类名称" json:"name"`
	Icon         string    `orm:"column(icon);size(150);default(/static/upload/default/user-default-60x60.png)" description:"分类Icon" json:"icon"`
	IsShow       int8      `orm:"column(is_show);default(0)" description:"是否显示" json:"is_show"`   // 0 显示 1 不显示
}

func (this *GoodsCat) TableName() string {
	return common.TableName("goods_cat")
}

func (this *GoodsCat) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *GoodsCat)SearchField() []string{
	return []string{"name"}
}

func (this *GoodsCat) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *GoodsCat) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *GoodsCat) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *GoodsCat) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

func GetGdsCatList() ([]*GoodsCat) {
	gcat_list := make([]*GoodsCat, 0)
	_, err := orm.NewOrm().QueryTable(GoodsCat{}).OrderBy("-id").All(&gcat_list)
	if err != nil {
		return nil
	}
	return gcat_list
}