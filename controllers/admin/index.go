package controllers

import (
  "blockshop/services"
  "fmt"
  "reflect"
)

type IndexController struct {
	baseController
}

func (this *IndexController) Index() {
	this.Data["login_user"] = loginUser
	//后台用户数量
	var adminUserService services.AdminUserService
	this.Data["admin_user_count"] = adminUserService.GetCount()
	//后台角色数量
	var adminRoleService services.AdminRoleService
	this.Data["admin_role_count"] = adminRoleService.GetCount()
	//后台菜单数量
	var adminMenuService services.AdminMenuService
	this.Data["admin_menu_count"] = adminMenuService.GetCount()
	this.Data["admin_log_count"] = 0

	this.TestInterface()

	this.Layout = "public/base.html"
	this.TplName = "index/index.html"
}


func (this *IndexController) TestInterface() {
	var a int
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a).Kind())
}
