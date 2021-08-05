package controllers

import (
  "blockshop/form_validate"
  "blockshop/global"
  "blockshop/global/response"
  "blockshop/services"
  "github.com/gookit/validate"
  "strconv"
)

type GoodsAttrController struct {
  baseController
}

func (Self *GoodsAttrController) Index() {
  goods_id,_ := Self.GetInt64("id")
  var srv services.GoodsAttrService
  gQueryParams.Add("goods_id",strconv.Itoa(int(goods_id)))
  data, pagination := srv.GetPaginateData(admin["per_page"].(int), gQueryParams)
  Self.Data["goods_id"] = goods_id
  Self.Data["data"] = data
  Self.Data["paginate"] = pagination

  Self.Layout = "public/base.html"
  Self.TplName = "goods_attr/index.html"
}

func (Self *GoodsAttrController) Add() {
  goods_id,_ := Self.GetInt64("goods_id")
  Self.Data["goods_id"] = goods_id
  Self.Layout = "public/base.html"
  Self.TplName = "goods_attr/add.html"
}


func (Self *GoodsAttrController) Create() {
  var frm form_validate.GoodsAttrForm
  var srv services.GoodsAttrService
  if err := Self.ParseForm(&frm); err != nil {
    response.ErrorWithMessage(err.Error(), Self.Ctx)
  }
  v := validate.Struct(frm)
  if !v.Validate() {
    response.ErrorWithMessage(v.Errors.One(), Self.Ctx)
  }

  insertId := srv.Create(&frm)
  url := global.URL_BACK

  if frm.IsCreate == 1 {
    url = global.URL_RELOAD
  }
  if insertId > 0 {
    response.SuccessWithMessageAndUrl("添加成功", url, Self.Ctx)
  } else {
    response.Error(Self.Ctx)
  }
}

func (Self *GoodsAttrController) Edit() {
  id, _ := Self.GetInt64("id", -1)
  goods_id,_ := Self.GetInt64("goods_id")
  Self.Data["goods_id"] = goods_id
  if id <= 0 {
    response.ErrorWithMessage("Param is error.", Self.Ctx)
  }
  var srv services.GoodsAttrService

  data := srv.GetById(id)
  if data == nil {
    response.ErrorWithMessage("Not Found Info By Id.", Self.Ctx)
  }

  Self.Data["data"] = data
  Self.Layout = "public/base.html"
  Self.TplName = "goods_attr/edit.html"
}

func (Self *GoodsAttrController) Update(){
  var frm form_validate.GoodsAttrForm
  if err := Self.ParseForm(&frm); err != nil {
    response.ErrorWithMessage(err.Error(), Self.Ctx)
  }

  if frm.Id <= 0 {
    response.ErrorWithMessage("Params is Error.", Self.Ctx)
  }

  v := validate.Struct(frm)

  if !v.Validate() {
    response.ErrorWithMessage(v.Errors.One(), Self.Ctx)
  }

  //商家验重
  var srv services.GoodsAttrService

  if srv.Update(&frm) > 0 {
    response.Success(Self.Ctx)
  } else {
    response.Error(Self.Ctx)
  }
}

func (Self *GoodsAttrController) Del() {
  idStr := Self.GetString("id")
  ids := make([]int, 0)
  var idArr []int

  if idStr == "" {
    Self.Ctx.Input.Bind(&ids, "id")
  } else {
    id, _ := strconv.Atoi(idStr)
    idArr = append(idArr, id)
  }

  if len(ids) > 0 {
    idArr = ids
  }
  if len(idArr) == 0 {
    response.ErrorWithMessage("参数id错误.", Self.Ctx)
  }
  var srv services.GoodsAttrService
  if srv.Del(idArr) > 0 {
    response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, Self.Ctx)
  } else {
    response.Error(Self.Ctx)
  }
}
