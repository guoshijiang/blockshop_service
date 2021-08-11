package services

import (
  beego_pagination "blockshop/common/utils/beego-pagination"
  "blockshop/form_validate"
  "blockshop/models"
  "github.com/astaxie/beego/orm"
  "net/url"
)

type GoodsCateService struct {
  BaseService
}

func (self *GoodsCateService) GetPaginateData(listRows int, params url.Values) ([]*models.GoodsCat, beego_pagination.Pagination) {
  //搜索、查询字段赋值
  self.SearchField = append(self.SearchField, new(models.GoodsCat).SearchField()...)

  var cate []*models.GoodsCat
  o := orm.NewOrm().QueryTable(new(models.GoodsCat))
  _, err := self.PaginateAndScopeWhere(o, listRows, params).All(&cate)
  if err != nil {
    return nil, self.Pagination
  } else {
    return cate, self.Pagination
  }
}

func (*GoodsCateService) Create(form *form_validate.GoodsCateForm) int {
  cate := models.GoodsCat{
    CatLevel: form.CatLevel,
    FatherCatId: form.FatherCatId,
    Name: form.Name,
    Icon: form.Icon,
    IsShow: form.IsShow,
  }
  id, err := orm.NewOrm().Insert(&cate)
  if err == nil {
    return int(id)
  } else {
    return 0
  }
}

func (*GoodsCateService) GetGoodsById(id int64) *models.GoodsCat {
  o := orm.NewOrm()
  cate := models.GoodsCat{Id: id}
  err := o.Read(&cate)
  if err != nil {
    return nil
  }
  return &cate
}

func (*GoodsCateService) IsExistName(name string, id int64) bool {
  if id == 0 {
    return orm.NewOrm().QueryTable(new(models.GoodsCat)).Filter("name", name).Exist()
  } else {
    return orm.NewOrm().QueryTable(new(models.GoodsCat)).Filter("name", name).Exclude("id", id).Exist()
  }
}

func (*GoodsCateService) Update(form *form_validate.GoodsCateForm) int{
  o := orm.NewOrm()
  cate := models.GoodsCat{Id: form.Id}
  if o.Read(&cate) == nil {
    cate.CatLevel = form.CatLevel
    cate.FatherCatId = form.FatherCatId
    cate.Name = form.Name
    cate.Icon = form.Icon
    cate.IsShow = form.IsShow
    num, err := o.Update(&cate)
    if err == nil {
      return int(num)
    } else {
      return 0
    }
  }
  return 0
}

func (*GoodsCateService) Del(ids []int) int{
  count, err := orm.NewOrm().QueryTable(new(models.GoodsCat)).Filter("id__in", ids).Delete()
  if err == nil {
    return int(count)
  } else {
    return 0
  }
}

func (*GoodsCateService) GetGoodsCats() []*models.GoodsCat{
  var cats []*models.GoodsCat
  _,err := orm.NewOrm().QueryTable(new(models.GoodsCat)).Filter("is_removed__contains",0).All(&cats)
  if err != nil {
    return nil
  }
  return cats
}
