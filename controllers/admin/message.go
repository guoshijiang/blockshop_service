package controllers

import (
  "blockshop/form_validate"
  "blockshop/global"
  "blockshop/global/response"
  "blockshop/services"
  "github.com/gookit/validate"
)

type MessageController struct {
  baseController
}

//工单列表-按照发送者显示
func(Self *MessageController) Index() {
  var srv services.MessageService
  data, pagination := srv.GetPaginateDataList(admin["per_page"].(int), gQueryParams)
  Self.Data["data"] = data
  Self.Data["paginate"] = pagination

  Self.Layout = "public/base.html"
  Self.TplName = "message/index.html"
}

//发送历史
func(Self *MessageController) History() {
  var srv services.MessageService
  data, pagination := srv.GetPaginateDataHistory(admin["per_page"].(int), gQueryParams)
  Self.Data["data"] = data
  Self.Data["paginate"] = pagination
  sendUserId,_ := Self.GetInt64("send_user_id",0)
  Self.Data["SendUserId"] = sendUserId

  Self.Layout = "public/base.html"
  Self.TplName = "message/history.html"
}

//发送消息
func(Self *MessageController) Send() {
  var (
    vForm form_validate.MessageForm
    srv services.MessageService
  )
  if err := Self.ParseForm(&vForm); err != nil {
    response.ErrorWithMessage(err.Error(), Self.Ctx)
  }
  v := validate.Struct(vForm)
  if !v.Validate() {
    response.ErrorWithMessage(v.Errors.One(), Self.Ctx)
  }

  if vForm.MsgType == 0 {
    vForm.MsgContent = vForm.MsgText
  } else {
    vForm.MsgContent = vForm.MsgImg
  }
  insertId := srv.Create(&vForm)
  url := global.URL_BACK

  if vForm.IsCreate == 1 {
    url = global.URL_RELOAD
  }
  if insertId > 0 {
    response.SuccessWithMessageAndUrl("添加成功", url, Self.Ctx)
  } else {
    response.Error(Self.Ctx)
  }
}