package models

import (
  "blockshop/common"
  "github.com/astaxie/beego/orm"
)

type UserWallet struct {
  BaseModel
  Id          int64     `orm:"pk;column(id);auto;size(11)" description:"钱包ID" json:"id"`
  UserId      int64     `orm:"column(user_id);index" description:"用户ID" json:"user_id"`
  AssetId     int64     `orm:"column(asset_id);index" description:"资产ID" json:"asset_id"`
  ChainName   string    `orm:"column(chain_name);default(Bitcoin)" description:"链的名称" json:"chain_name"`
  Address     string    `orm:"column(address);size(256)" description:"地址" json:"address"`
  Balance     float64   `orm:"column(balance);default(150);digits(22);decimals(8)" description:"钱包余额" json:"balance"`
}

func (this *UserWallet) TableName() string {
  return common.TableName("user_wallet")
}

func (this *UserWallet) Query() orm.QuerySeter {
  return orm.NewOrm().QueryTable(this)
}

func (this *UserWallet) Insert() error {
  if _, err := orm.NewOrm().Insert(this); err != nil {
    return err
  }
  return nil
}

func (this *UserWallet) Update(fields ...string) error {
  if _, err := orm.NewOrm().Update(this, fields...); err != nil {
    return err
  }
  return nil
}