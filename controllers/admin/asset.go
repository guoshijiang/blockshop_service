package controllers

import (
  "blockshop/form_validate"
  "blockshop/global"
  "blockshop/global/response"
  "blockshop/services"
  "github.com/gookit/validate"
  "strconv"
)

type AssetController struct {
  baseController
}

func (Self *AssetController) Index() {
  var assetService services.AssetService
  data, pagination := assetService.GetPaginateData(admin["per_page"].(int), gQueryParams)
  Self.Data["data"] = data
  Self.Data["paginate"] = pagination

  Self.Layout = "public/base.html"
  Self.TplName = "asset/index.html"
}

func (Self *AssetController) Add() {
  Self.Layout = "public/base.html"
  Self.TplName = "asset/add.html"
}


func (Self *AssetController) Create() {
  var AssetForm form_validate.AssetForm
  var assetService services.AssetService
  if err := Self.ParseForm(&AssetForm); err != nil {
    response.ErrorWithMessage(err.Error(), Self.Ctx)
  }
  v := validate.Struct(AssetForm)
  if !v.Validate() {
    response.ErrorWithMessage(v.Errors.One(), Self.Ctx)
  }


  insertId := assetService.Create(&AssetForm)
  url := global.URL_BACK



  if AssetForm.IsCreate == 1 {
    url = global.URL_RELOAD
  }
  if insertId > 0 {
    response.SuccessWithMessageAndUrl("添加成功", url, Self.Ctx)
  } else {
    response.Error(Self.Ctx)
  }
}

func (Self *AssetController) Edit() {
  id, _ := Self.GetInt64("id", -1)
  if id <= 0 {
    response.ErrorWithMessage("Param is error.", Self.Ctx)
  }
  var assetService services.AssetService

  asset := assetService.GetRowById(id)
  if asset == nil {
    response.ErrorWithMessage("Not Found Info By Id.", Self.Ctx)
  }

  Self.Data["data"] = asset
  Self.Layout = "public/base.html"
  Self.TplName = "asset/edit.html"
}

func (Self *AssetController) Update(){
  var assetForm form_validate.AssetForm
  if err := Self.ParseForm(&assetForm); err != nil {
    response.ErrorWithMessage(err.Error(), Self.Ctx)
  }

  if assetForm.Id <= 0 {
    response.ErrorWithMessage("Params is Error.", Self.Ctx)
  }

  v := validate.Struct(assetForm)

  if !v.Validate() {
    response.ErrorWithMessage(v.Errors.One(), Self.Ctx)
  }

  //商家验重
  var assetService services.AssetService
  if assetService.Update(&assetForm) > 0 {
    response.Success(Self.Ctx)
  } else {
    response.Error(Self.Ctx)
  }
}

func (Self *AssetController) Del() {
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
  var assetService services.AssetService
  if assetService.Del(idArr) > 0 {
    response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, Self.Ctx)
  } else {
    response.Error(Self.Ctx)
  }
}
