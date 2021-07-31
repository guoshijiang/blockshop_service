package api

import (
  "blockshop/models"
  "blockshop/types"
  "blockshop/types/user"
  "encoding/json"
  "github.com/astaxie/beego"
)


type UserController struct {
  beego.Controller
}

// Register @Title Register
// @Description 注册手机号 Register
// @Success 200 status bool, data interface{}, msg string
// @router /register [post]
func (this *UserController) Register() {
  user_register := user.Register{}
  if err := json.Unmarshal(this.Ctx.Input.RequestBody, &user_register); err != nil {
    this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
    this.ServeJSON()
    return
  }
  if code, err := user_register.ParamCheck(); err != nil {
    this.Data["json"] = RetResource(false, code, nil, err.Error())
    this.ServeJSON()
    return
  }
  code, msg := models.UserRegister(user_register)
  if code == types.ReturnSuccess {
    this.Data["json"] = RetResource(true, code, nil, msg)
  } else {
    this.Data["json"] = RetResource(false, code, nil, msg)
  }
  this.ServeJSON()
  return
}


// Login @Title Login
// @Description 登录 Login
// @Success 200 status bool, data interface{}, msg string
// @router /login [post]
func (this *UserController) Login() {
  user_login := user.Login{}
  if err := json.Unmarshal(this.Ctx.Input.RequestBody, &user_login); err != nil {
    this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
    this.ServeJSON()
    return
  }
  if code, err := user_login.ParamCheck(); err != nil {
    this.Data["json"] = RetResource(false, code, nil, err.Error())
    this.ServeJSON()
    return
  }
  data, code, msg := models.UserLogin(user_login)
  if code == types.ReturnSuccess {
    this.Data["json"] = RetResource(true, code, data, msg)
  } else {
    this.Data["json"] = RetResource(false, code, nil, msg)
  }
  this.ServeJSON()
  return
}