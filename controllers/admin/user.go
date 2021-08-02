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
)

type UserController struct {
  baseController
}

func (Self *UserController) Index() {
  var srv services.UserService
  data, pagination := srv.GetPaginateData(admin["per_page"].(int), gQueryParams)

  Self.Data["data"] = data
  Self.Data["paginate"] = pagination

  Self.Layout = "public/base.html"
  Self.TplName = "user/index.html"
}

func (Self *UserController) Add() {
  Self.Data["user_level_list"] = []models.OptionList{{Id: 0,Name: "V0"},{Id: 1,Name: "V1"},{Id: 2,Name: "V2"},{Id: 3,Name: "V3"}}
  Self.Layout = "public/base.html"
  Self.TplName = "user/add.html"
}


//添加用户
func (Self *UserController) Create() {
 var userForm form_validate.UserForm
 if err := Self.ParseForm(&userForm); err != nil {
   response.ErrorWithMessage(err.Error(), Self.Ctx)
 }

 v := validate.Struct(userForm)

 if !v.Validate() {
   response.ErrorWithMessage(v.Errors.One(), Self.Ctx)
 }

 var srv services.UserService
  //默认
  userForm.Avator = "/static/admin/images/avatar.png"
  //上传LOGO
  imgPath, err := new(services.UploadService).Upload(Self.Ctx, "avatar")
  if err != nil {
    log.Println("upload--err",err)
  }
  if len(imgPath) > 0 {
    userForm.Avator = imgPath
  }


 insertId := srv.Create(&userForm)

 url := global.URL_BACK
 if userForm.IsCreate == 1 {
   url = global.URL_RELOAD
 }

 if insertId > 0 {
   response.SuccessWithMessageAndUrl("添加成功", url, Self.Ctx)
 } else {
   response.Error(Self.Ctx)
 }
}

//用户-修改界面
func (Self *UserController) Edit() {
  id, _ := Self.GetInt64("id", -1)
  if id <= 0 {
    response.ErrorWithMessage("Param is error.", Self.Ctx)
  }

  var userService services.UserService

  user := userService.GetRowById(id)
  if user == nil {
    response.ErrorWithMessage("Not Found Info By Id.", Self.Ctx)
  }

  Self.Data["user_level_list"] = []models.OptionList{{Id: 0,Name: "V0"},{Id: 1,Name: "V1"},{Id: 2,Name: "V2"},{Id: 3,Name: "V3"}}
  Self.Data["data"] = user
  Self.Layout = "public/base.html"
  Self.TplName = "user/edit.html"
}

//用户-修改
func (Self *UserController) Update() {
  var userForm form_validate.UserForm
  if err := Self.ParseForm(&userForm); err != nil {
    response.ErrorWithMessage(err.Error(), Self.Ctx)
  }

  if userForm.Id <= 0 {
    response.ErrorWithMessage("Params is Error.", Self.Ctx)
  }
  var userService services.UserService

  if err := userService.CheckPassword(&userForm,1);err != nil {
    response.ErrorWithMessage("资金密码和原密码相同", Self.Ctx)
  }

  if err := userService.CheckPassword(&userForm,2);err != nil {
    response.ErrorWithMessage("用户密码和原密码相同", Self.Ctx)
  }

  v := validate.Struct(userForm)

  if !v.Validate() {
    response.ErrorWithMessage(v.Errors.One(), Self.Ctx)
  }


  num := userService.Update(&userForm)

  if num > 0 {
    response.Success(Self.Ctx)
  } else {
    response.Error(Self.Ctx)
  }
}

//删除
func (Self *UserController) Del() {
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

  var userService services.UserService
  count := userService.Del(idArr)

  if count > 0 {
    response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, Self.Ctx)
  } else {
    response.Error(Self.Ctx)
  }
}
