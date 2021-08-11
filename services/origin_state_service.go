package services

import (
  beego_pagination "blockshop/common/utils/beego-pagination"
  "blockshop/form_validate"
  "blockshop/models"
  "github.com/astaxie/beego/orm"
  "net/url"
)

type OriginStateService struct {
  BaseService
}

func (self *OriginStateService) GetPaginateData(listRows int, params url.Values) ([]*models.GoodsOriginState, beego_pagination.Pagination) {
  //搜索、查询字段赋值
  self.SearchField = append(self.SearchField, new(models.GoodsOriginState).SearchField()...)

  var data []*models.GoodsOriginState
  o := orm.NewOrm().QueryTable(new(models.GoodsOriginState))
  _, err := self.PaginateAndScopeWhere(o, listRows, params).All(&data)
  if err != nil {
    return nil, self.Pagination
  } else {
    return data, self.Pagination
  }
}

func (*OriginStateService) Create(form *form_validate.OriginStateForm) int {
  data := models.GoodsOriginState{
    Name: form.Name,
    IsShow: form.IsShow,
  }
  id, err := orm.NewOrm().Insert(&data)

  if err == nil {
    return int(id)
  } else {
    return 0
  }
}

func (*OriginStateService) GetRowById(id int64) *models.GoodsOriginState {
  o := orm.NewOrm()
  data := models.GoodsOriginState{Id: id}
  err := o.Read(&data)
  if err != nil {
    return nil
  }
  return &data
}

func (*OriginStateService) Update(form *form_validate.OriginStateForm) int{
  o := orm.NewOrm()
  data := models.GoodsOriginState{Id: form.Id}
  if o.Read(&data) == nil {
    data.Name = form.Name
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

func (*OriginStateService) Del(ids []int) int{
  count, err := orm.NewOrm().QueryTable(new(models.GoodsOriginState)).Filter("id__in", ids).Delete()
  if err == nil {
    return int(count)
  } else {
    return 0
  }
}

func (*OriginStateService) GetOrigins() []*models.GoodsOriginState{
  var data []*models.GoodsOriginState
  _,err := orm.NewOrm().QueryTable(new(models.GoodsOriginState)).Filter("is_show__eq",0).All(&data)
  if err != nil {
    return nil
  }
  return data
}