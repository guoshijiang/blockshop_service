package services

import (
  beego_pagination "blockshop/common/utils/beego-pagination"
  "blockshop/form_validate"
  "blockshop/models"
  "github.com/astaxie/beego/orm"
  "net/url"
)

type GoodsAttrService struct {
  BaseService
}


func (self *GoodsAttrService) GetPaginateData(listRows int, params url.Values) ([]*models.GoodsAttr, beego_pagination.Pagination) {
  //搜索、查询字段赋值
  self.SearchField = append(self.SearchField, new(models.GoodsAttr).SearchField()...)
  self.WhereField = append(self.WhereField,[]string{"goods_id"}...)
  var cate []*models.GoodsAttr
  o := orm.NewOrm().QueryTable(new(models.GoodsAttr))
  _, err := self.PaginateAndScopeWhere(o, listRows, params).All(&cate)
  if err != nil {
    return nil, self.Pagination
  } else {
    return cate, self.Pagination
  }
}

func (*GoodsAttrService) Create(form *form_validate.GoodsAttrForm) int {
  cate := models.GoodsAttr{
    GoodsId: form.GoodsId,
    TypeKey: form.TypeKey,
    TypeVale: form.TypeVale,
    IsShow: form.IsShow,
  }
  id, err := orm.NewOrm().Insert(&cate)
  if err == nil {
    return int(id)
  } else {
    return 0
  }
}

func (*GoodsAttrService) GetById(id int64) *models.GoodsAttr {
  o := orm.NewOrm()
  data := models.GoodsAttr{Id: id}
  err := o.Read(&data)
  if err != nil {
    return nil
  }
  return &data
}


func (*GoodsAttrService) Update(form *form_validate.GoodsAttrForm) int{
  o := orm.NewOrm()
  data := models.GoodsAttr{Id: form.Id}
  if o.Read(&data) == nil {
    data.TypeVale = form.TypeVale
    data.TypeKey = form.TypeKey
    data.IsShow = form.IsShow
    data.GoodsId = form.GoodsId
    num, err := o.Update(&data)
    if err == nil {
      return int(num)
    } else {
      return 0
    }
  }
  return 0
}

func (*GoodsAttrService) Del(ids []int) int{
  count, err := orm.NewOrm().QueryTable(new(models.GoodsAttr)).Filter("id__in", ids).Delete()
  if err == nil {
    return int(count)
  } else {
    return 0
  }
}

//func (*GoodsAttrService) GetGoodsTypes() []*models.GoodsType{
//  var data []*models.GoodsType
//  _,err := orm.NewOrm().QueryTable(new(models.GoodsType)).Filter("is_show__contains",0).All(&data)
//  if err != nil {
//    return nil
//  }
//  return data
//}