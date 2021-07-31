package models

import (
  "blockshop/common"
  "github.com/astaxie/beego/logs"
  "github.com/astaxie/beego/orm"
)

type GoodsType struct {
  BaseModel
  Id             int64     `orm:"pk;column(id);auto;size(11)" description:"商品属性ID" json:"id"`
  GoodsId        int64     `orm:"column(goods_id)" description:"商品ID" json:"goods_id"`                         // 商品ID
  TypeKey        string    `orm:"column(type_key);size(512)" description:"属性名称" json:"type_key"`   // 如颜色，输入商品的人可以自定义
  TypeVale       string    `orm:"column(type_vale);size(1024)" description:"属性文字" json:"type_vale"` // 入库数据格式 ["白色", "蓝色", "黄色"]
  IsShow         int8      `orm:"column(is_show);default(1)" description:"是否显示" json:"is_show"`   // 0 不显示 1 显示
}


func (this *GoodsType) TableName() string {
  return common.TableName("goods_type")
}

func (this *GoodsType) Read(fields ...string) error {
  logs.Info(fields)
  return nil
}

func (this *GoodsType) SearchField() []string {
  return []string{"type_name"}
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

func (this *GoodsType) Insert() error {
  if _, err := orm.NewOrm().Insert(this); err != nil {
    return err
  }
  return nil
}