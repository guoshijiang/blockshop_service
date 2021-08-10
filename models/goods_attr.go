package models

import (
	"blockshop/common"
	"blockshop/types"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type GoodsAttr struct {
	BaseModel
	Id             int64     `orm:"pk;column(id);auto;size(11)" description:"商品属性ID" json:"id"`
	GoodsId        int64     `orm:"column(goods_id)" description:"商品ID" json:"goods_id"`                         // 商品ID
	TypeKey        string    `orm:"column(type_key);size(512)" description:"属性名称" json:"type_key"`   // 如颜色，输入商品的人可以自定义
	TypeVale       string    `orm:"column(type_vale);size(1024)" description:"属性文字" json:"type_vale"` // 入库数据格式 ["白色", "蓝色", "黄色"]
	IsShow         int8      `orm:"column(is_show);default(1)" description:"是否显示" json:"is_show"`   // 0 不显示 1 显示
}


func (this *GoodsAttr) TableName() string {
	return common.TableName("goods_attr")
}

func (this *GoodsAttr) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *GoodsAttr) SearchField() []string {
	return []string{"type_name"}
}

func (this *GoodsAttr) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *GoodsAttr) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *GoodsAttr) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *GoodsAttr) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}

func GetGoodsAttrList(goods_id int64) ([]*GoodsAttr, int64, error) {
	var type_list []*GoodsAttr
	if _, err := orm.NewOrm().QueryTable(GoodsAttr{}).Filter("GoodsId", goods_id).All(&type_list); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return type_list, types.ReturnSuccess, nil
}