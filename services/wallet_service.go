package services

import (
  beego_pagination "blockshop/common/utils/beego-pagination"
  "blockshop/models"
  "github.com/astaxie/beego/orm"
  "net/url"
)

type WalletService struct {
  BaseService
}

func (Self *WalletService) GetUserData (listRows int,params url.Values) ([]*models.UserWalletList,beego_pagination.Pagination) {
  var data []*models.UserWalletList
  var total int64
  om := orm.NewOrm()
  inner := "from user_wallet as t0 inner join user as t1 on t1.id = t0.user_id inner join asset as t2 on t2.id = t0.asset_id where t0.id > 0 "
  sql := "select t0.*,t1.user_name,t2.name asset_name " + inner
  sql1 := "select count(*) total " + inner

  //搜索、查询字段赋值
  Self.SearchField = append(Self.SearchField, new(models.WalletRecord).SearchField()...)
  where,param := Self.ScopeWhereRaw(params)

  if err := om.Raw(sql1+where,param).QueryRow(&total);err != nil {
    return nil,beego_pagination.Pagination{}
  }
  Self.Pagination.Total = int(total)
  Self.PaginateRaw(listRows,params)
  param = append(param,listRows*(Self.Pagination.CurrentPage-1),listRows)
  if _,err := om.Raw(sql+where+" order by created_at desc limit ?,?",param).QueryRows(&data);err != nil {
    return nil,beego_pagination.Pagination{}
  }
  return data,Self.Pagination
}

func (Self *WalletService) GetRecordData(listRows int,params url.Values) ([]*models.WalletRecordList,beego_pagination.Pagination){
  var data []*models.WalletRecordList
  var total int64
  om := orm.NewOrm()
  inner := "from wallet_record as t0 inner join user as t1 on t1.id = t0.user_id inner join asset as t2 on t2.id = t0.asset_id inner join admin_user t3 on t3.id = t0.admin where t0.id > 0 "
  sql := "select t0.*,t1.user_name,t2.name asset_name,t3.username admin_name " + inner
  sql1 := "select count(*) total " + inner

  //搜索、查询字段赋值
  Self.SearchField = append(Self.SearchField, new(models.WalletRecord).SearchField()...)
  where,param := Self.ScopeWhereRaw(params)

  if err := om.Raw(sql1+where,param).QueryRow(&total);err != nil {
    return nil,beego_pagination.Pagination{}
  }
  Self.Pagination.Total = int(total)
  Self.PaginateRaw(listRows,params)
  param = append(param,listRows*(Self.Pagination.CurrentPage-1),listRows)
  if _,err := om.Raw(sql+where+" order by created_at desc limit ?,?",param).QueryRows(&data);err != nil {
    return nil,beego_pagination.Pagination{}
  }
  return data,Self.Pagination
}

