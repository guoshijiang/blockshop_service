package models

import (
	"blockshop/common"
)

type Merchant struct {
	BaseModel
	Id             int64     `orm:"pk;column(id);auto;size(11)" description:"商品属性ID"  json:"id"`
	Logo           string    `orm:"column(logo);size(150);default(/static/upload/default/user-default-60x60.png)" description:"商家Logo" json:"logo"`
	MerchantName   string    `orm:"column(merchant_name);size(512);index" description:"商家名称" json:"merchant_name"`
	MerchantIntro  string    `orm:"column(merchant_intro);size(512);index" description:"商家简介" json:"merchant_intro"`
	MerchantDetail string    `orm:"column(merchant_detail);type(text)" description:"商家详情" json:"merchant_detail"`
	ContactUser    string    `orm:"column(contact_user);size(128);index" description:"商家联系人" json:"contact_user"`
	Phone          string    `orm:"column(phone);size(64);index" description:"商家联系电话" json:"phone"`
	WeChat         string    `orm:"column(we_chat);size(64);index" description:"商家联系微信" json:"we_chat"`
	Address        string    `orm:"column(address);size(512);index" description:"店铺地址" json:"address"`
	GoodsNum       int64     `orm:"column(goods_num)" description:"商品总数" json:"goods_num"`
	MerchantWay    int8      `orm:"column(merchant_way)default(0);index" description:"商家类别" json:"merchant_way"`   // 0:自营商家； 1:认证商家  2:普通商家
	SettlePercent  float64   `orm:"column(settle_percent)default(0);digits(22);decimals(8)" description:"结算比例"  json:"settle_percent"`
	ShopLevel      int8      `orm:"column(shop_level)" description:"店铺等级" json:"shop_level"`
	ShopServer     int8      `orm:"column(shop_server)" description:"店铺服务" json:"shop_server"`
}

func (this *Merchant) TableName() string {
	return common.TableName("merchant")
}

