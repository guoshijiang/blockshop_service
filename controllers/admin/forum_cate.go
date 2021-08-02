package controllers

import (
  "blockshop/form_validate"
  "blockshop/global"
  "blockshop/global/response"
  "blockshop/services"
  "github.com/gookit/validate"
  "log"
  "strconv"
  "strings"
)

type ForumCateController struct {
  baseController
}

func (Self *ForumCateController) Index() {
  var serv services.ForumCateService
  data, pagination := serv.GetPaginateData(admin["per_page"].(int), gQueryParams)
  Self.Data["data"] = data
  Self.Data["paginate"] = pagination

  Self.Layout = "public/base.html"
  Self.TplName = "forum_cate/index.html"
}

func (Self *ForumCateController) Add() {
  Self.Data["cats"] = new(services.ForumCateService).GetCats()
  Self.Layout = "public/base.html"
  Self.TplName = "forum_cate/add.html"
}


func (Self *ForumCateController) Create() {
  var form form_validate.ForumCateForm
  var serv services.ForumCateService
  if err := Self.ParseForm(&form); err != nil {
    response.ErrorWithMessage(err.Error(), Self.Ctx)
  }
  v := validate.Struct(form)
  if !v.Validate() {
    response.ErrorWithMessage(v.Errors.One(), Self.Ctx)
  }

  imgPath, err := new(services.UploadService).Upload(Self.Ctx, "icon")
  if err != nil {
    log.Println("upload--err",err)
  }
  form.Icon = imgPath

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

func (Self *ForumCateController) Edit() {
  id, _ := Self.GetInt64("id", -1)
  if id <= 0 {
    response.ErrorWithMessage("Param is error.", Self.Ctx)
  }
  var serv services.ForumCateService

  cate := serv.GetRowById(id)
  if cate == nil {
    response.ErrorWithMessage("Not Found Info By Id.", Self.Ctx)
  }

  Self.Data["cats"] = new(services.ForumCateService).GetCats()
  Self.Data["data"] = cate
  Self.Layout = "public/base.html"
  Self.TplName = "forum_cate/edit.html"
}

func (Self *ForumCateController) Update(){
  var form form_validate.ForumCateForm
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
  var serv services.ForumCateService
  if serv.IsExistName(strings.TrimSpace(form.Name), form.Id) {
    response.ErrorWithMessage("分类已经存在", Self.Ctx)
  }

  imgPath, err := new(services.UploadService).Upload(Self.Ctx, "icon")
  if err != nil {
    log.Println("upload--err",err)
  }
  form.Icon = imgPath
  if serv.Update(&form) > 0 {
    response.Success(Self.Ctx)
  } else {
    response.Error(Self.Ctx)
  }
}

func (Self *ForumCateController) Del() {
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
  var serv services.ForumCateService
  if serv.Del(idArr) > 0 {
    response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, Self.Ctx)
  } else {
    response.Error(Self.Ctx)
  }
}