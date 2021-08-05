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
  self.WhereField = append(self.WhereField,[]string{"name"}...)
  var data []*models.GoodsType
  o := orm.NewOrm().QueryTable(new(models.GoodsType))
  _, err := self.PaginateAndScopeWhere(o, listRows, params).All(&data)
  if err != nil {
    return nil, self.Pagination
  } else {
    return data, self.Pagination
  }
}

func (*GoodsTypeService) Create(form *form_validate.GoodsTypeForm) int {
  data := models.GoodsType{
    Name: form.Name,
    Icon: form.Icon,
    IsShow: form.IsShow,
  }
  id, err := orm.NewOrm().Insert(&data)
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
    data.Name = form.Name
    if len(form.Icon) > 0 {
      data.Icon = form.Icon
    }
    data.IsShow = form.IsShow
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

func (*GoodsTypeService) GetGoodsTypes() []*models.GoodsType{
  var data []*models.GoodsType
  _,err := orm.NewOrm().QueryTable(new(models.GoodsType)).Filter("is_show__eq",0).All(&data)
  if err != nil {
    return nil
  }
  return data
}