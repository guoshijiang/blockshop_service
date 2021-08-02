package services

import (
  beego_pagination "blockshop/common/utils/beego-pagination"
  "blockshop/form_validate"
  "blockshop/models"
  "github.com/astaxie/beego/orm"
  "net/url"
)

type ForumCateService struct {
  BaseService
}

func (self *ForumCateService) GetPaginateData(listRows int, params url.Values) ([]*models.ForumCat, beego_pagination.Pagination) {
  //搜索、查询字段赋值
  self.SearchField = append(self.SearchField, new(models.ForumCat).SearchField()...)

  var data []*models.ForumCat
  o := orm.NewOrm().QueryTable(new(models.ForumCat))
  _, err := self.PaginateAndScopeWhere(o, listRows, params).All(&data)
  if err != nil {
    return nil, self.Pagination
  } else {
    return data, self.Pagination
  }
}

func (*ForumCateService) Create(form *form_validate.ForumCateForm) int {
  cate := models.ForumCat{
    FatherCatId: form.FatherCatId,
    Name: form.Name,
    Icon: form.Icon,
    IsShow: form.IsShow,
    Introduce:form.Introduce,
  }
  id, err := orm.NewOrm().Insert(&cate)
  if err == nil {
    return int(id)
  } else {
    return 0
  }
}

func (*ForumCateService) GetRowById(id int64) *models.ForumCat {
  o := orm.NewOrm()
  data := models.ForumCat{Id: id}
  err := o.Read(&data)
  if err != nil {
    return nil
  }
  return &data
}

func (*ForumCateService) IsExistName(name string, id int64) bool {
  if id == 0 {
    return orm.NewOrm().QueryTable(new(models.ForumCat)).Filter("name", name).Exist()
  } else {
    return orm.NewOrm().QueryTable(new(models.ForumCat)).Filter("name", name).Exclude("id", id).Exist()
  }
}

func (*ForumCateService) Update(form *form_validate.ForumCateForm) int{
  o := orm.NewOrm()
  data:= models.ForumCat{Id: form.Id}
  if o.Read(&data) == nil {
    data.Introduce = form.Introduce
    data.FatherCatId = form.FatherCatId
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

func (*ForumCateService) Del(ids []int) int{
  count, err := orm.NewOrm().QueryTable(new(models.ForumCat)).Filter("id__in", ids).Delete()
  if err == nil {
    return int(count)
  } else {
    return 0
  }
}

func (*ForumCateService) GetCats() []*models.ForumCat{
  var data []*models.ForumCat
  _,err := orm.NewOrm().QueryTable(new(models.ForumCat)).Filter("is_removed__contains",0).All(&data)
  if err != nil {
    return nil
  }
  return data
}




