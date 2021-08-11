package controllers

import (
  "blockshop/form_validate"
  "blockshop/global"
  "blockshop/global/response"
  "blockshop/services"
  "github.com/gookit/validate"
  "strconv"
)

type OriginStateController struct {
  baseController
}

func (Self *OriginStateController) Index() {
  var srv services.OriginStateService
  data, pagination := srv.GetPaginateData(admin["per_page"].(int), gQueryParams)
  Self.Data["data"] = data
  Self.Data["paginate"] = pagination

  Self.Layout = "public/base.html"
  Self.TplName = "origin/index.html"
}

func (Self *OriginStateController) Add() {
  Self.Layout = "public/base.html"
  Self.TplName = "origin/add.html"
}


func (Self *OriginStateController) Create() {
  var form form_validate.OriginStateForm
  var srv services.OriginStateService
  if err := Self.ParseForm(&form); err != nil {
    response.ErrorWithMessage(err.Error(), Self.Ctx)
  }
  v := validate.Struct(form)
  if !v.Validate() {
    response.ErrorWithMessage(v.Errors.One(), Self.Ctx)
  }


  insertId := srv.Create(&form)
  url := global.URL_BACK



  if form.IsCreate == 1 {
    url = global.URL_RELOAD
  }
  if insertId > 0 {
    response.SuccessWithMessageAndUrl("添加成功", url, Self.Ctx)
  } else {
    response.Error(Self.Ctx)
  }
}

func (Self *OriginStateController) Edit() {
  id, _ := Self.GetInt64("id", -1)
  if id <= 0 {
    response.ErrorWithMessage("Param is error.", Self.Ctx)
  }
  var srv services.OriginStateService

  asset := srv.GetRowById(id)
  if asset == nil {
    response.ErrorWithMessage("Not Found Info By Id.", Self.Ctx)
  }

  Self.Data["data"] = asset
  Self.Layout = "public/base.html"
  Self.TplName = "origin/edit.html"
}

func (Self *OriginStateController) Update(){
  var form form_validate.OriginStateForm
  if err := Self.ParseForm(&form); err != nil {
    response.ErrorWithMessage(err.Error(), Self.Ctx)
  }

  if form.Id <= 0 {
    response.ErrorWithMessage("Params is Error.", Self.Ctx)
  }

  v := validate.Struct(form)

  if !v.Validate() {
    response.ErrorWithMessage(v.Errors.One(), Self.Ctx)
  }

  //商家验重
  var srv services.OriginStateService
  if srv.Update(&form) > 0 {
    response.Success(Self.Ctx)
  } else {
    response.Error(Self.Ctx)
  }
}

func (Self *OriginStateController) Del() {
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
  var srv services.OriginStateService
  if srv.Del(idArr) > 0 {
    response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, Self.Ctx)
  } else {
    response.Error(Self.Ctx)
  }
}