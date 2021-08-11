package services

import (
  beego_pagination "blockshop/common/utils/beego-pagination"
  "blockshop/models"
  "github.com/astaxie/beego/orm"
  "net/url"
  "time"
)

type OrderService struct {
  BaseService
}

func (Self *OrderService) GetPaginateData(listRows int, params url.Values) ([]*models.GoodsOrder, beego_pagination.Pagination) {
  //搜索、查询字段赋值
  Self.SearchField = append(Self.SearchField, new(models.GoodsOrder).SearchField()...)

  var order []*models.GoodsOrder
  o := orm.NewOrm().QueryTable(new(models.GoodsOrder))
  _, err := Self.PaginateAndScopeWhere(o, listRows, params).All(&order)
  if err != nil {
    return nil, Self.Pagination
  } else {
    return order, Self.Pagination
  }
}


func (Self *OrderService) GetPaginateDataRaw(listRows int, params url.Values) ([]*models.GoodsOrderList, beego_pagination.Pagination) {
  var data []*models.GoodsOrderList
  var total int64
  om := orm.NewOrm()
  inner := "from  goods_order as t0 inner join user as t1 on t1.id = t0.user_id where t0.id > 0 "
  sql := "select t0.*,t1.user_name buy_user " + inner
  sql1 := "select count(*) total " + inner

  //搜索、查询字段赋值
  Self.SearchField = append(Self.SearchField, new(models.GoodsOrder).SearchField()...)
  where,param := Self.ScopeWhereRaw(params)

  if err := om.Raw(sql1+where).QueryRow(&total);err != nil {
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

func (Self *OrderService) GetPaginateProcessDataRaw(listRows int, params url.Values) ([]*models.OrderProcessList, beego_pagination.Pagination){
  var data []*models.OrderProcessList
  var total int64
  om := orm.NewOrm()
  inner := "from  order_process as t0 inner join goods_order as t1 on t1.id = t0.order_id inner join user as t2 on t2.id = t0.user_id where t0.id > 0 "
  sql := "select t0.*,t2.user_name,t1.goods_title,t1.order_number,t1.goods_title,t1.pay_amount " + inner
  sql1 := "select count(*) total " + inner

  //搜索、查询字段赋值
  Self.SearchField = append(Self.SearchField, new(models.OrderProcess).SearchField()...)
  where,param := Self.ScopeWhereRaw(params)

  if err := om.Raw(sql1+where).QueryRow(&total);err != nil {
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

func (Self *OrderService) GetPaginateProcessDetailRaw (id int64)(*models.OrderProcessList, error) {
  var data models.OrderProcessList
  om := orm.NewOrm()
  inner := "from  order_process as t0 inner join goods_order as t1 on t1.id = t0.order_id inner join user as t2 on t2.id = t0.user_id where t0.id  = ?"
  sql := "select t0.*,t2.user_name,t1.goods_title,t1.order_number,t1.goods_title,t1.pay_amount " + inner

  if err := om.Raw(sql+" limit 1", id).QueryRow(&data); err != nil {
    return nil,err
  }
  return &data, nil
}

func (*OrderService) Del(ids []int) int{
  var (
    count  int64
    err	error
  )
  if AdminUserVal.MerchantId > 0 {
    count, err = orm.NewOrm().QueryTable(new(models.GoodsOrder)).Filter("id__in", ids).Delete()
  } else {
    count, err = orm.NewOrm().QueryTable(new(models.GoodsOrder)).Filter("id__in", ids).Delete()
  }
  if err == nil {
    return int(count)
  } else {
    return 0
  }
}

func (Self *OrderService) GetOrderById(id int64) *models.GoodsOrder {
  var data  models.GoodsOrder
  if AdminUserVal.MerchantId > 0 {
    if err := orm.NewOrm().QueryTable(new(models.GoodsOrder)).Filter("id__eq", id).Filter("merchant_id", AdminUserVal.MerchantId).One(&data);err != nil {
      return nil
    }
  }else {
    data = models.GoodsOrder{Id: id}
    if err := orm.NewOrm().Read(&data); err != nil{
      return nil
    }
  }
  return &data
}

func (Self *OrderService) UpdateShipNumber(cond *models.GoodsOrder) int{
  o := orm.NewOrm()
  data := models.GoodsOrder{Id: cond.Id}
  if err := o.Read(&data);err == nil {
    data.ShipNumber = cond.ShipNumber
    data.OrderStatus = 4
    data.DealMerchant = AdminUserVal.Username
    data.DealAt = time.Now()
    if _,err := o.Update(&data);err == nil {
      return 1
    }
  }
  return 0
}