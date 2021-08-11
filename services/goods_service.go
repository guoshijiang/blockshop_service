package services

import (
  beego_pagination "blockshop/common/utils/beego-pagination"
  "blockshop/form_validate"
  "blockshop/models"
  "fmt"
  "github.com/astaxie/beego/orm"
  "net/url"
  "strconv"
)

type GoodsServices struct {
  BaseService
}

func (self *GoodsServices) GetPaginateData(listRows int, params url.Values) ([]*models.Goods, beego_pagination.Pagination) {
  //搜索、查询字段赋值
  self.SearchField = append(self.SearchField, new(models.Goods).SearchField()...)
  if AdminUserVal.MerchantId > 0 {
    params.Set("_merchant_id", strconv.Itoa(AdminUserVal.MerchantId))
  }
  var goods []*models.Goods
  o := orm.NewOrm().QueryTable(new(models.Goods))
  _, err := self.PaginateAndScopeWhere(o, listRows, params).All(&goods)
  if err != nil {
    return nil, self.Pagination
  } else {
    return goods, self.Pagination
  }
}

func (*GoodsServices) Create(form *form_validate.GoodsForm) int {
  if AdminUserVal.MerchantId > 0 {
    form.MerchantId = int64(AdminUserVal.MerchantId)
  }
  goods := models.Goods{
    GoodsName: form.GoodsName,
    GoodsParams: form.GoodsParams,
    GoodsDetail: form.GoodsDetail,
    Discount: form.Discount,
    IsDiscount: form.IsDiscount,
    Sale: form.Sale,
    Title: form.Title,
    IsHot: form.IsHot,
    IsDisplay: form.IsDisplay,
    Logo: form.Logo,
    GoodsMark: form.GoodsMark,
    IsLimitTime: form.IsLimitTime,
    GoodsPrice: form.GoodsPrice,
    GoodsDisPrice: form.GoodsDisPrice,
    Serveice: form.Serveice,
    CalcWay:form.CalcWay,
    TotalAmount: form.TotalAmount,
    MerchantId: form.MerchantId,
    GoodsCatId: form.GoodsCatId,
    OriginStateId: form.OriginStateId,
  }
  id, err := orm.NewOrm().Insert(&goods)

  if err == nil {
    return int(id)
  } else {
    return 0
  }
}

func (*GoodsServices) GetGoodsById(id int64) *models.Goods {
  o := orm.NewOrm()
  var goods models.Goods
  if AdminUserVal.MerchantId > 0 {
    goods = models.Goods{Id: id, MerchantId: int64(AdminUserVal.MerchantId)}
  } else {
    goods = models.Goods{Id:id}
  }
  err := o.Read(&goods)
  if err != nil {
    return nil
  }
  return &goods
}

func (*GoodsServices) GetGoodsImagesById(id int64) []*models.GoodsImage {
  o := orm.NewOrm()
  var data []*models.GoodsImage
  _,err := o.QueryTable(new(models.GoodsImage)).Filter("goods_id__contains",id).All(&data)
  if err != nil {
    return nil
  }
  return data
}

func (*GoodsServices) IsExistName(goods_name string, id int64) bool {
  fmt.Println("goods_name___",goods_name)
  if id == 0 {
    if AdminUserVal.MerchantId  > 0 {
      return orm.NewOrm().QueryTable(new(models.Goods)).Filter("goods_name", goods_name).Filter("merchant_id",AdminUserVal.MerchantId).Exist()
    }
    return orm.NewOrm().QueryTable(new(models.Goods)).Filter("goods_name", goods_name).Exist()
  } else {
    if AdminUserVal.MerchantId > 0 {
      return orm.NewOrm().QueryTable(new(models.Goods)).Filter("goods_name", goods_name).Filter("merchant_id", AdminUserVal.MerchantId).Exclude("id", id).Exist()
    }
    return orm.NewOrm().QueryTable(new(models.Goods)).Filter("goods_name", goods_name).Exclude("id", id).Exist()
  }
}

func (*GoodsServices) Update(form *form_validate.GoodsForm) int{
  o := orm.NewOrm()
  goods := models.Goods{Id: form.Id}
  if o.Read(&goods) == nil {
    goods.GoodsName = form.GoodsName
    goods.Title = form.Title
    goods.GoodsParams = form.GoodsParams
    goods.GoodsDetail = form.GoodsDetail
    if len(form.Logo) > 0 {
      goods.Logo = form.Logo
    }
    goods.Discount  = form.Discount
    goods.Sale  = form.Sale
    goods.IsHot  = form.IsHot
    goods.IsDisplay  = form.IsDisplay
    goods.GoodsMark  = form.GoodsMark
    goods.IsLimitTime  = form.IsLimitTime
    goods.GoodsPrice  = form.GoodsPrice
    goods.GoodsDisPrice  = form.GoodsDisPrice
    goods.Serveice  = form.Serveice
    goods.CalcWay  = form.CalcWay
    goods.TotalAmount  = form.TotalAmount
    goods.GoodsCatId = form.GoodsCatId
    goods.Discount = form.Discount
    goods.IsDiscount = form.IsDiscount
    goods.OriginStateId = form.OriginStateId
    if AdminUserVal.MerchantId > 0 {
      goods.MerchantId = int64(AdminUserVal.MerchantId)
    } else {
      goods.MerchantId = form.MerchantId
    }
    num, err := o.Update(&goods)
    if err == nil {
      return int(num)
    } else {
      fmt.Println("update--err",err)
      return 0
    }
  }
  return 0
}

func (*GoodsServices) Del(ids []int) int{
  var (
    count  int64
    err	error
  )
  if AdminUserVal.MerchantId > 0 {
    count, err = orm.NewOrm().QueryTable(new(models.Goods)).Filter("id__in", ids).Delete()
  } else {
    count, err = orm.NewOrm().QueryTable(new(models.Goods)).Filter("id__in", ids).Delete()
  }
  if err == nil {
    return int(count)
  } else {
    return 0
  }
}


//func (Self *GoodsServices) GetPaginateCommentData(listRows int, params url.Values) ([]*models.GoodsComment, beego_pagination.Pagination) {
//  var data []*models.GoodsComment
//  var total int64
//  om := orm.NewOrm()
//  inner := "from goods_comment as t0 inner join goods as t1 on t1.id = t0.goods_id where t0.goods_id > 0 "
//  sql := "select t0.* " + inner
//  sql1 := "select count(*) total " + inner
//
//  //搜索、查询字段赋值
//  Self.SearchField = append(Self.SearchField, new(models.IntegralRecord).SearchField()...)
//  where,param := Self.ScopeWhereRaw(params)
//
//  if AdminUserVal.MerchantId > 0 {
//    where += " and t1.merchant_id = ? "
//    param = append(param,AdminUserVal.MerchantId)
//  }
//  if err := om.Raw(sql1+where,param).QueryRow(&total);err != nil {
//    return nil,beego_pagination.Pagination{}
//  }
//  Self.Pagination.Total = int(total)
//  Self.PaginateRaw(listRows,params)
//  param = append(param,listRows*(Self.Pagination.CurrentPage-1),listRows)
//  if _,err := om.Raw(sql+where+" order by created_at desc limit ?,?",param).QueryRows(&data);err != nil {
//    return nil,beego_pagination.Pagination{}
//  }
//  return data,Self.Pagination
//}
