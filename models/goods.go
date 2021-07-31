package models

import (
  "blockshop/common"
  "github.com/astaxie/beego/logs"
  "github.com/astaxie/beego/orm"
)

type Goods struct {
  BaseModel
  Id             int64     `orm:"pk;column(id);auto;size(11)" description:"商家ID" json:"id"`
  UserId         int64     `orm:"column(user_id)" description:"商家对应的用户ID" json:"user_id"`
  GoodsCatId     int64     `orm:"column(goods_cat_id)" description:"商品分类ID" json:"goods_cat_id"`
  GoodsMark      string    `orm:"column(goods_mark);size(512);index" description:"商品备注" json:"goods_mark"`
  Serveice       string    `orm:"column(serveice);size(512);index" description:"服务说明" json:"serveice"`
  CalcWay        int8      `orm:"column(calc_way);default(0);index" description:"计量方式" json:"calc_way"` // 0:按件计量 1:按近计量
  MerchantId     int64     `orm:"column(merchant_id);default(0);index" description:"商品所属商家ID" json:"merchant_id"`
  Title          string    `orm:"column(title);size(512)" description:"商品标题" json:"title"`
  Logo           string    `orm:"column(logo);size(150);default(/static/upload/default/user-default-60x60.png)" description:"商品封面" json:"logo"`
  TotalAmount    int64     `orm:"column(total_amount);default(150000)" description:"商品总量" json:"total_amount"`
  LeftAmount     int64     `orm:"column(left_amount);default(150000)" description:"剩余商品总量" json:"left_amount"`
  GoodsPrice     float64   `orm:"column(goods_price);default(1);digits(22);decimals(8)" description:"商品价格" json:"goods_price"`
  GoodsDisPrice  float64   `orm:"column(goods_discount_price);default(1);digits(22);decimals(8)" description:"商品折扣价格" json:"goods_discount_price"`
  GoodsName      string    `orm:"column(goods_name);size(512);index" description:"产品名称" json:"goods_name"`
  GoodsParams    string    `orm:"column(goods_params);type(text)" description:"产品参数" json:"goods_params"`
  GoodsDetail    string    `orm:"column(goods_detail);type(text)" description:"产品详细介绍" json:"goods_detail"`
  Discount       float64   `orm:"column(discount);default(0);index" description:"折扣" json:"discount"`              // 取值 0.1-9.9；0代表不打折
  Sale           int8      `orm:"column(sale);default(0);index" description:"上架下架" json:"sale"`                   // 0:上架 1:下架
  IsDisplay      int8      `orm:"column(is_display);default(0);index" description:"首页展示" json:"is_display"`       // 0:首页不展示, 1:首页展示
  SellNums       int64     `orm:"column(sell_nums);default(0);index" description:"售出数量" json:"sell_nums"`
  IsHot          int8      `orm:"column(is_hot);default(0);index" description:"爆款产品" json:"is_hot"`                // 0:非爆款产品 1:爆款产品
  IsDiscount     int8      `orm:"column(is_discount);default(0);index" description:"打折活动" json:"is_discount"`      // 0:不打折，1:打折活动产品
  LeftTime       int64     `orm:"column(left_time);default(0);index" description:"限时产品剩余时间" json:"left_time"`   // 限时产品剩余时间
  IsLimitTime    int8      `orm:"column(is_limit_time);default(0);index" description:"限时产品" json:"is_limit_time"`  // 0:不是限时产品 1:是限时
}

type Select struct {
  Id			int					`json:"id"`
  Name		string				`json:"name"`
}

func (this *Goods) TableName() string {
  return common.TableName("goods")
}

func (this *Goods) Read(fields ...string) error {
  logs.Info(fields)
  return nil
}

func (*Goods) SearchField() []string {
  return []string{"goods_name"}
}

func (this *Goods) Update(fields ...string) error {
  if _, err := orm.NewOrm().Update(this, fields...); err != nil {
    return err
  }
  return nil
}

func (this *Goods) Delete() error {
  if _, err := orm.NewOrm().Delete(this); err != nil {
    return err
  }
  return nil
}

func (this *Goods) Query() orm.QuerySeter {
  return orm.NewOrm().QueryTable(this)
}

func (this *Goods) Insert() error {
  if _, err := orm.NewOrm().Insert(this); err != nil {
    return err
  }
  return nil
}