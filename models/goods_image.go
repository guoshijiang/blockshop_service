package models

import (
	"blockshop/common"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type GoodsImage struct {
	BaseModel
	Id           int64     `orm:"pk;column(id);auto;size(11)'" description:"商品图片ID" json:"id"`
	GoodsId      int64     `orm:"column(goods_id)" description:"商品ID" json:"goods_id"`
	Image        string    `orm:"column(image);size(150);default(/static/upload/default/user-default-60x60.png)" description:"商品图片" json:"image"`
	IsShow       int8      `orm:"column(is_show);default(1)" description:"是否显示" json:"is_show"`   // 0 不显示 1 显示
}

func (this *GoodsImage) TableName() string {
	return common.TableName("goods_image")
}

func (this *GoodsImage) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (*GoodsImage) SearchField() []string {
	return []string{"goods_name"}
}

func (this *GoodsImage) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *GoodsImage) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *GoodsImage) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *GoodsImage) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

