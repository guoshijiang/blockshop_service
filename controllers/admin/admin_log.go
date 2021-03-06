package controllers

import (
	"blockshop/services"
)

type AdminLogController struct {
	baseController
}

func (this *AdminLogController) Index() {
	var (
		adminLogService  services.AdminLogService
		adminUserService services.AdminUserService
	)
	data, pagination := adminLogService.GetPaginateData(admin["per_page"].(int), gQueryParams)

	this.Data["admin_user_list"] = adminUserService.GetAllAdminUser()

	this.Data["data"] = data
	this.Data["paginate"] = pagination
	this.Layout = "public/base.html"
	this.TplName = "admin_log/index.html"
}
