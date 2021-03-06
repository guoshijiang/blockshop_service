package controllers

import (
	"blockshop/common/utils"
	"blockshop/form_validate"
	"blockshop/global"
	"blockshop/global/response"
	"blockshop/models"
	"blockshop/services"
	"github.com/adam-hanna/arrayOperations"
	"github.com/gookit/validate"
	"strconv"
	"strings"
)

type AdminMenuController struct {
	baseController
}

//菜单首页
func (this *AdminMenuController) Index() {
	var adminTreeService services.AdminTreeService
	this.Data["data"] = adminTreeService.AdminMenuTree()

	this.Layout = "public/base.html"
	this.TplName = "admin_menu/index.html"
}

//添加菜单界面
func (this *AdminMenuController) Add() {
	var adminTreeService services.AdminTreeService
	parentId, _ := this.GetInt("parent_id", 0)
	parents := adminTreeService.Menu(parentId, 0)

	this.Data["parents"] = parents
	this.Data["log_method"] = new(models.AdminMenu).GetLogMethod()

	this.Layout = "public/base.html"
	this.TplName = "admin_menu/add.html"
}

//添加菜单
func (this *AdminMenuController) Create() {
	var adminMenuService services.AdminMenuService
	adminMenuForm := form_validate.AdminMenuForm{}

	if err := this.ParseForm(&adminMenuForm); err != nil {
		response.ErrorWithMessage(err.Error(), this.Ctx)
	}

	//去除Url前后两侧的空格
	if adminMenuForm.Url != "" {
		adminMenuForm.Url = strings.TrimSpace(adminMenuForm.Url)
	}

	//数据校验
	v := validate.Struct(adminMenuForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), this.Ctx)
	}

	//添加之前url验重
	if adminMenuService.IsExistUrl(adminMenuForm.Url, adminMenuForm.Id) {
		response.ErrorWithMessage("url【"+adminMenuForm.Url+"】已经存在.", this.Ctx)
	}

	//创建
	_, err := adminMenuService.Create(&adminMenuForm)
	if err != nil {
		response.Error(this.Ctx)
	}

	url := global.URL_BACK
	if adminMenuForm.IsCreate == 1 {
		url = global.URL_RELOAD
	}

	response.SuccessWithMessageAndUrl("添加成功", url, this.Ctx)
}

//编辑菜单界面
func (this *AdminMenuController) Edit() {
	id, _ := this.GetInt("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", this.Ctx)
	}

	var (
		adminMenuService services.AdminMenuService
		adminTreeService services.AdminTreeService
	)

	adminMenu := adminMenuService.GetAdminMenuById(id)
	if adminMenu == nil {
		response.ErrorWithMessage("Not Found Info By Id.", this.Ctx)
	}

	parentId := adminMenu.ParentId
	parents := adminTreeService.Menu(parentId, 0)

	this.Data["data"] = adminMenu
	this.Data["parents"] = parents
	this.Data["log_method"] = new(models.AdminMenu).GetLogMethod()

	this.Layout = "public/base.html"
	this.TplName = "admin_menu/edit.html"
}

//菜单更新
func (this *AdminMenuController) Update() {
	var adminMenuService services.AdminMenuService
	adminMenuForm := form_validate.AdminMenuForm{}

	if err := this.ParseForm(&adminMenuForm); err != nil {
		response.ErrorWithMessage(err.Error(), this.Ctx)
	}

	//去除Url前后两侧的空格
	if adminMenuForm.Url != "" {
		adminMenuForm.Url = strings.TrimSpace(adminMenuForm.Url)
	}

	//数据校验
	v := validate.Struct(adminMenuForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), this.Ctx)
	}

	//添加之前url验重
	if adminMenuService.IsExistUrl(adminMenuForm.Url, adminMenuForm.Id) {
		response.ErrorWithMessage("url【"+adminMenuForm.Url+"】已经存在.", this.Ctx)
	}

	count := adminMenuService.Update(&adminMenuForm)

	if count > 0 {
		response.Success(this.Ctx)
	} else {
		response.Error(this.Ctx)
	}
}

//删除
func (this *AdminMenuController) Del() {
	idStr := this.GetString("id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		this.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}

	if len(idArr) == 0 {
		response.ErrorWithMessage("参数id错误.", this.Ctx)
	}

	var adminMenuService services.AdminMenuService
	//判断是否有子菜单
	if adminMenuService.IsChildMenu(idArr) {
		response.ErrorWithMessage("有子菜单不可删除！", this.Ctx)
	}

	noDeletionId := new(models.AdminMenu).NoDeletionId()

	m, b := arrayOperations.Intersect(noDeletionId, idArr)

	if len(noDeletionId) > 0 && len(m.Interface().([]int)) > 0 && b {
		response.ErrorWithMessage("ID为"+strings.Join(utils.IntArrToStringArr(noDeletionId), ",")+"的数据无法删除!", this.Ctx)
	}

	count := adminMenuService.Del(idArr)

	if count > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, this.Ctx)
	} else {
		response.Error(this.Ctx)
	}
}
