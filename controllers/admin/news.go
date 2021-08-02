package controllers

import (
  "blockshop/form_validate"
  "blockshop/global"
  "blockshop/global/response"
  "blockshop/services"
  "github.com/gookit/validate"
  "log"
  "strconv"
)

type NewsController struct {
   baseController
}

func (Self *NewsController) Index() {
  var serv services.NewsService
  data, pagination := serv.GetPaginateData(admin["per_page"].(int), gQueryParams)
  Self.Data["data"] = data
  Self.Data["paginate"] = pagination

  Self.Layout = "public/base.html"
  Self.TplName = "news/index.html"
}

func (Self *NewsController) Add() {
  Self.Layout = "public/base.html"
  Self.TplName = "news/add.html"
}


func (Self *NewsController) Create() {
  var form form_validate.NewsForm
  var serv services.NewsService
  if err := Self.ParseForm(&form); err != nil {
    response.ErrorWithMessage(err.Error(), Self.Ctx)
  }
  v := validate.Struct(form)
  if !v.Validate() {
    response.ErrorWithMessage(v.Errors.One(), Self.Ctx)
  }
  //默认
  form.Image = "/static/admin/images/avatar.png"
  //上传LOGO
  imgPath, err := new(services.UploadService).Upload(Self.Ctx, "image")
  if err != nil {
    log.Println("upload--err",err)
  }
  if len(imgPath) > 0 {
    form.Image = imgPath
  }

  insertId := serv.Create(&form)
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

func (Self *NewsController) Edit() {
  id, _ := Self.GetInt64("id", -1)
  if id <= 0 {
    response.ErrorWithMessage("Param is error.", Self.Ctx)
  }
  var serv services.NewsService

  data := serv.GetRowById(id)
  if data == nil {
    response.ErrorWithMessage("Not Found Info By Id.", Self.Ctx)
  }

  Self.Data["data"] = data
  Self.Layout = "public/base.html"
  Self.TplName = "news/edit.html"
}

func (Self *NewsController) Update(){
  var form form_validate.NewsForm
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

  //上传Image
  imgPath, err := new(services.UploadService).Upload(Self.Ctx, "image")
  if err != nil {
    log.Println("upload--err",err)
  }
  form.Image = imgPath

  //商家验重
  var serv services.NewsService
  if serv.Update(&form) > 0 {
    response.Success(Self.Ctx)
  } else {
    response.Error(Self.Ctx)
  }
}

func (Self *NewsController) Del() {
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
  var serv services.NewsService
  if serv.Del(idArr) > 0 {
    response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, Self.Ctx)
  } else {
    response.Error(Self.Ctx)
  }
}
