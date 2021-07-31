package services

import (
  beego_pagination "blockshop/common/utils/beego-pagination"
  "blockshop/form_validate"
  "blockshop/models"
  "github.com/astaxie/beego/orm"
  "net/url"
)

type GoodsTypeService struct {
  BaseService
}


func (self *GoodsTypeService) GetPaginateData(listRows int, params url.Values) ([]*models.GoodsType, beego_pagination.Pagination) {
  //搜索、查询字段赋值
  self.SearchField = append(self.SearchField, new(models.GoodsType).SearchField()...)

  self.WhereField = append(self.WhereField,[]string{"goods_id"}...)
  var cate []*models.GoodsType
  o := orm.NewOrm().QueryTable(new(models.GoodsType))
  _, err := self.PaginateAndScopeWhere(o, listRows, params).All(&cate)
  if err != nil {
    return nil, self.Pagination
  } else {
    return cate, self.Pagination
  }
}

func (*GoodsTypeService) Create(form *form_validate.GoodsTypeForm) int {
  cate := models.GoodsType{
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

func (*GoodsTypeService) GetById(id int64) *models.GoodsType {
  o := orm.NewOrm()
  data := models.GoodsType{Id: id}
  err := o.Read(&data)
  if err != nil {
    return nil
  }
  return &data
}


func (*GoodsTypeService) Update(form *form_validate.GoodsTypeForm) int{
  o := orm.NewOrm()
  data := models.GoodsType{Id: form.Id}
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

func (*GoodsTypeService) Del(ids []int) int{
  count, err := orm.NewOrm().QueryTable(new(models.GoodsType)).Filter("id__in", ids).Delete()
  if err == nil {
    return int(count)
  } else {
    return 0
  }
}