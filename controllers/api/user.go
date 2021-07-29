package api


import (
  "blockshop/models"
  "github.com/astaxie/beego"
)


type UserController struct {
	beego.Controller
}

// @Title Register
// @Description 注册手机号 Register
// @Success 200 status bool, data interface{}, msg string
// @router /register [post]
func (this *UserController) Register() {
	models.GetUserById(1)
	this.Data["json"] = RetResource(true, 200, nil, "success")
	this.ServeJSON()
	return
}