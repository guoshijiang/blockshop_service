package services

import (
  beego_pagination "blockshop/common/utils/beego-pagination"
  "blockshop/form_validate"
  "blockshop/models"
  "fmt"
  "github.com/astaxie/beego/orm"
  "net/url"
)

type ForumService struct {
  BaseService
}

func (Self *ForumService) GetForumData (listRows int,params url.Values) ([]*models.ForumList,beego_pagination.Pagination) {
  var data []*models.ForumList
  var total int64
  om := orm.NewOrm()
  inner := "from forum as t0 inner join user as t1 on t1.id = t0.user_id inner join forum_cat as t2 on t2.id = t0.cat_id where t0.id > 0 "
  sql := "select t0.*,t1.user_name,t2.name cat_name " + inner
  sql1 := "select count(*) total " + inner

  //搜索、查询字段赋值
  Self.SearchField = append(Self.SearchField, new(models.Forum).SearchField()...)
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

//check
func (Self *ForumService) CheckForum(form *form_validate.ForumForm) int{
  data := models.Forum{
    Id: form.Id,
    IsCheck: 1,
  }
  id, err := orm.NewOrm().Update(&data,"is_check")
  if err == nil {
    return int(id)
  } else {
    return 0
  }
}

/**
 * reply list
*/
func (Self *ForumService) Reply(listRows int,params url.Values) ([]*models.ForumReplyList,beego_pagination.Pagination){
  var data []*models.ForumReplyList
  var total int64
  om := orm.NewOrm()
  inner := "from forum_reply as t0 inner join user as t1 on t1.id = t0.user_id inner join forum as t2 on t2.id = t0.forum_id where t0.id > 0 "
  sql := "select t0.*,t1.user_name " + inner
  sql1 := "select count(*) total " + inner

  fmt.Println("params--",params)
  //搜索、查询字段赋值
  Self.SearchField = append(Self.SearchField, new(models.ForumReply).SearchField()...)
  Self.WhereField = append(Self.WhereField,[]string{"t0.id"}...)
  where,param := Self.ScopeWhereRaw(params)

  fmt.Println("where",where)
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

func (Self *ForumService) CheckReplyForum(form *form_validate.ForumForm) int{
  data := models.ForumReply{
    Id: form.Id,
    IsCheck: 1,
  }
  id, err := orm.NewOrm().Update(&data,"is_check")
  if err == nil {
    return int(id)
  } else {
    return 0
  }
}
