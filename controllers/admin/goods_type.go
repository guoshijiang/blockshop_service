package controllers

import (
  "blockshop/form_validate"
  "blockshop/global"
  "blockshop/global/response"
  "blockshop/services"
  "github.com/gookit/validate"
  "strconv"
)

type GoodsTypeController struct {
  baseController
}

func (Self *GoodsTypeController) Index() {
  goods_id,_ := Self.GetInt64("id")
  var srv services.GoodsTypeService
  gQueryParams.Add("goods_id",strconv.Itoa(int(goods_id)))
  data, pagination := srv.GetPaginateData(admin["per_page"].(int), gQueryParams)
  Self.Data["goods_id"] = goods_id
  Self.Data["data"] = data
  Self.Data["paginate"] = pagination

  Self.Layout = "public/base.html"
  Self.TplName = "goods_type/index.html"
}

func (Self *GoodsTypeController) Add() {
  goods_id,_ := Self.GetInt64("goods_id")
  Self.Data["goods_id"] = goods_id
  Self.Layout = "public/base.html"
  Self.TplName = "goods_type/add.html"
}


func (Self *GoodsTypeController) Create() {
  var frm form_validate.GoodsTypeForm
  var srv services.GoodsTypeService
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

func (Self *GoodsTypeController) Edit() {
  id, _ := Self.GetInt64("id", -1)
  goods_id,_ := Self.GetInt64("goods_id")
  Self.Data["goods_id"] = goods_id
  if id <= 0 {
    response.ErrorWithMessage("Param is error.", Self.Ctx)
  }
  var srv services.GoodsTypeService

  data := srv.GetById(id)
  if data == nil {
    response.ErrorWithMessage("Not Found Info By Id.", Self.Ctx)
  }

  Self.Data["data"] = data
  Self.Layout = "public/base.html"
  Self.TplName = "goods_type/edit.html"
}

func (Self *GoodsTypeController) Update(){
  var frm form_validate.GoodsTypeForm
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
  var srv services.GoodsTypeService

  if srv.Update(&frm) > 0 {
    response.Success(Self.Ctx)
  } else {
    response.Error(Self.Ctx)
  }
}

func (Self *GoodsTypeController) Del() {
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
  var srv services.GoodsTypeService
  if srv.Del(idArr) > 0 {
    response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, Self.Ctx)
  } else {
    response.Error(Self.Ctx)
  }
}
