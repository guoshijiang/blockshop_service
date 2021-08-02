package services

import (
  beego_pagination "blockshop/common/utils/beego-pagination"
  "blockshop/form_validate"
  "blockshop/models"
  "github.com/astaxie/beego/orm"
  "net/url"
)

type MessageService struct {
  BaseService
}

func (Self *MessageService) GetPaginateDataList(listRows int, params url.Values) ([]*models.MessageData, beego_pagination.Pagination) {
  _,err := orm.NewOrm().Raw("SET sql_mode = 'STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION'").Exec()
  if err != nil {
    println("err ----",err)
  }
  var data []*models.MessageData
  var total int64
  om := orm.NewOrm()
  inner := "from (select message.*,user.user_name from message inner join user on user.id= message.send_user_id order by created_at desc) t0 "
  sql := "select * " + inner
  sql1 := "select count(*) total " + inner

  //搜索、查询字段赋值
  Self.SearchField = append(Self.SearchField, new(models.Message).SearchField()...)
  where,param := Self.ScopeWhereRaw(params)
  Self.PaginateRaw(listRows,params)

  if err := om.Raw(sql1+where,param).QueryRow(&total);err != nil {
    return nil,beego_pagination.Pagination{}
  }
  Self.Pagination.Total = int(total)
  param = append(param,listRows*(Self.Pagination.CurrentPage-1),listRows)
  if _,err := om.Raw(sql+where+" group by send_user_id limit ?,?",param).QueryRows(&data);err != nil {
    return nil,beego_pagination.Pagination{}
  }
  return data,Self.Pagination
}

func (Self *MessageService) GetPaginateDataHistory(listRows int, params url.Values) ([]*models.MessageData, beego_pagination.Pagination) {
  var data []*models.MessageData
  var total int64
  om := orm.NewOrm()

  //搜索、查询字段赋值
  Self.SearchField = append(Self.SearchField, new(models.Message).SearchField()...)
  _,param := Self.ScopeWhereRaw(params)
  Self.PaginateRaw(listRows,params)
  sql_total := "select sum(t.c) from (select count(*) as c from message where send_user_id = ? and target_user_id = 0 union select count(*) as c from message where send_user_id = 0 and target_user_id = ?) t; "
  param = append(param,[]interface{}{params.Get("send_user_id"),params.Get("send_user_id")}...)
  if err := om.Raw(sql_total,param).QueryRow(&total);err != nil {
    return nil,beego_pagination.Pagination{}
  }
  Self.Pagination.Total = int(total)
  sql_raw := "select * from (select * from message where send_user_id = ? and target_user_id = 0 union select * from message where  send_user_id = 0 and target_user_id = ?) t order by t.created_at desc"
  param = append(param,listRows*(Self.Pagination.CurrentPage-1),listRows)
  if _,err := om.Raw(sql_raw+" limit ?,?",param).QueryRows(&data);err != nil {
    return nil,beego_pagination.Pagination{}
  }
  return data,Self.Pagination
}

func (*MessageService) Create(form *form_validate.MessageForm) int {
  cate := models.Message{
    SendUserId: 0,
    TargetUserId: form.TargetUserId,
    MsgType: form.MsgType,
   // MsgTilte: form.MsgTilte,
    MsgContent: form.MsgContent,
  }
  id, err := orm.NewOrm().Insert(&cate)
  if err == nil {
    return int(id)
  } else {
    return 0
  }
}