package controllers

import (
  "blockshop/form_validate"
  "blockshop/global/response"
  "blockshop/services"
  "github.com/gookit/validate"
)

type ForumController struct {
  baseController
}

/**
 * list
 */
func (Self *ForumController) List() {
  var serv services.ForumService
  data, pagination := serv.GetForumData(admin["per_page"].(int), gQueryParams)
  Self.Data["data"] = data
  Self.Data["paginate"] = pagination

  Self.Layout = "public/base.html"
  Self.TplName = "forum/index.html"
}

/**
 * check
*/
func(Self *ForumController) CheckForum() {
  var serv services.ForumService
  var form form_validate.ForumForm
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

  if serv.CheckForum(&form) > 0 {
    response.Success(Self.Ctx)
  } else {
    response.Error(Self.Ctx)
  }
}

/**
 * reply list
*/
func (Self *ForumController) Reply() {
  var serv services.ForumService
  data, pagination := serv.Reply(admin["per_page"].(int), gQueryParams)
  Self.Data["data"] = data
  Self.Data["paginate"] = pagination

  Self.Layout = "public/base.html"
  Self.TplName = "forum/reply.html"
}

/**
 * reply check
*/
func (Self *ForumController) CheckReplyForum() {
  var serv services.ForumService
  var form form_validate.ForumForm
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

  if serv.CheckReplyForum(&form) > 0 {
    response.Success(Self.Ctx)
  } else {
    response.Error(Self.Ctx)
  }
}

