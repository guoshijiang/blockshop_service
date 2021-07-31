package controllers

import (
  "blockshop/form_validate"
  "blockshop/global"
  "blockshop/global/response"
  "blockshop/models"
  "blockshop/services"
  "github.com/gookit/validate"
  "log"
  "strconv"
  "strings"
)

type GoodsCateController struct {
  baseController
}

func (Self *GoodsCateController) Index() {
  var goodsCateService services.GoodsCateService
  data, pagination := goodsCateService.GetPaginateData(admin["per_page"].(int), gQueryParams)
  Self.Data["data"] = data
  Self.Data["paginate"] = pagination

  Self.Layout = "public/base.html"
  Self.TplName = "goods_cate/index.html"
}

func (Self *GoodsCateController) Add() {
  Self.Data["levels"] = []models.Select{{Id: 1,Name: "一级"},{Id: 2,Name: "二级"}}
  Self.Data["cats"] = new(services.GoodsCateService).GetGoodsCats()
  Self.Layout = "public/base.html"
  Self.TplName = "goods_cate/add.html"
}


func (Self *GoodsCateController) Create() {
  var goodsCateForm form_validate.GoodsCateForm
  var goodsCateService services.GoodsCateService
  if err := Self.ParseForm(&goodsCateForm); err != nil {
    response.ErrorWithMessage(err.Error(), Self.Ctx)
  }
  v := validate.Struct(goodsCateForm)
  if !v.Validate() {
    response.ErrorWithMessage(v.Errors.One(), Self.Ctx)
  }

  imgPath, err := new(services.UploadService).Upload(Self.Ctx, "icon")
  if err != nil {
    log.Println("upload--err",err)
  }
  goodsCateForm.Icon = imgPath

  insertId := goodsCateService.Create(&goodsCateForm)
  url := global.URL_BACK

  if goodsCateForm.IsCreate == 1 {
    url = global.URL_RELOAD
  }
  if insertId > 0 {
    response.SuccessWithMessageAndUrl("添加成功", url, Self.Ctx)
  } else {
    response.Error(Self.Ctx)
  }
}

func (Self *GoodsCateController) Edit() {
  id, _ := Self.GetInt64("id", -1)
  if id <= 0 {
    response.ErrorWithMessage("Param is error.", Self.Ctx)
  }
  var goodsCateService services.GoodsCateService

  cate := goodsCateService.GetGoodsById(id)
  if cate == nil {
    response.ErrorWithMessage("Not Found Info By Id.", Self.Ctx)
  }

  Self.Data["levels"] = []models.Select{{Id: 1,Name: "一级"},{Id: 2,Name: "二级"}}
  Self.Data["cats"] = new(services.GoodsCateService).GetGoodsCats()
  Self.Data["data"] = cate
  Self.Layout = "public/base.html"
  Self.TplName = "goods_cate/edit.html"
}

func (Self *GoodsCateController) Update(){
  var goodsCateForm form_validate.GoodsCateForm
  if err := Self.ParseForm(&goodsCateForm); err != nil {
    response.ErrorWithMessage(err.Error(), Self.Ctx)
  }

  if goodsCateForm.Id <= 0 {
    response.ErrorWithMessage("Params is Error.", Self.Ctx)
  }

  v := validate.Struct(goodsCateForm)

  if !v.Validate() {
    response.ErrorWithMessage(v.Errors.One(), Self.Ctx)
  }

  //商家验重
  var goodsCateService services.GoodsCateService
  if goodsCateService.IsExistName(strings.TrimSpace(goodsCateForm.Name), goodsCateForm.Id) {
    response.ErrorWithMessage("分类已经存在", Self.Ctx)
  }

  imgPath, err := new(services.UploadService).Upload(Self.Ctx, "icon")
  if err != nil {
    log.Println("upload--err",err)
  }
  goodsCateForm.Icon = imgPath
  if goodsCateService.Update(&goodsCateForm) > 0 {
    response.Success(Self.Ctx)
  } else {
    response.Error(Self.Ctx)
  }
}

func (Self *GoodsCateController) Del() {
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
  var goodsCateService services.GoodsCateService
  if goodsCateService.Del(idArr) > 0 {
    response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, Self.Ctx)
  } else {
    response.Error(Self.Ctx)
  }
}