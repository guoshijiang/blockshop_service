package services

import (
  beego_pagination "blockshop/common/utils/beego-pagination"
  "blockshop/form_validate"
  "blockshop/models"
  "github.com/astaxie/beego/orm"
  "net/url"
)

type NewsService struct {
  BaseService
}

func (self *NewsService) GetPaginateData(listRows int, params url.Values) ([]*models.News, beego_pagination.Pagination) {
  //搜索、查询字段赋值
  self.SearchField = append(self.SearchField, new(models.News).SearchField()...)

  var data []*models.News
  o := orm.NewOrm().QueryTable(new(models.News))
  _, err := self.PaginateAndScopeWhere(o, listRows, params).All(&data)
  if err != nil {
    return nil, self.Pagination
  } else {
    return data, self.Pagination
  }
}

func (*NewsService) Create(form *form_validate.NewsForm) int {
  data := models.News{
    Title: form.Title,
    Author: form.Author,
    Abstract: form.Abstract,
    Content: form.Content,
    Image: form.Image,
  }
  id, err := orm.NewOrm().Insert(&data)

  if err == nil {
    return int(id)
  } else {
    return 0
  }
}

func (*NewsService) GetRowById(id int64) *models.News {
  o := orm.NewOrm()
  data := models.News{Id: id}
  err := o.Read(&data)
  if err != nil {
    return nil
  }
  return &data
}

func (*NewsService) IsExistName(name string, id int64) bool {
  if id == 0 {
    return orm.NewOrm().QueryTable(new(models.News)).Filter("name", name).Exist()
  } else {
    return orm.NewOrm().QueryTable(new(models.News)).Filter("name", name).Exclude("id", id).Exist()
  }
}

func (*NewsService) Update(form *form_validate.NewsForm) int{
  o := orm.NewOrm()
  data := models.News{Id: form.Id}
  if o.Read(&data) == nil {
    data.Title = form.Title
    data.Author = form.Author
    data.Abstract = form.Abstract
    data.Content = form.Content
    if len(form.Image) > 0 {
      data.Image = form.Image
    }
    num, err := o.Update(&data)
    if err == nil {
      return int(num)
    } else {
      return 0
    }
  }
  return 0
}

func (*NewsService) Del(ids []int) int{
  count, err := orm.NewOrm().QueryTable(new(models.News)).Filter("id__in", ids).Delete()
  if err == nil {
    return int(count)
  } else {
    return 0
  }
}