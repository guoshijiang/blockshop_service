package api

import (
	"blockshop/models"
	"blockshop/types"
	"blockshop/types/address"
	"encoding/json"
	"github.com/astaxie/beego"
	"strings"
)

type UserAddressController struct {
	beego.Controller
}

// AddAddress @Title AddAddress finished
// @Description 添加新地址 AddAddress
// @Success 200 status bool, data interface{}, msg string
// @router /add_address [post]
func (this *UserAddressController) AddAddress() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	requestUser, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	var add_address address.UserAddressAddCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &add_address); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := add_address.UserAddressAddCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	if requestUser.Id != add_address.UserId {
		this.Data["json"] = RetResource(false, types.UserIsNotExist, err, "Token 和用户不匹配，拒绝添加地址")
		this.ServeJSON()
		return
	}
	address_u := models.UserAddress{
		UserId:   add_address.UserId,
		UserName: add_address.UserName,
		Phone:    add_address.Phone,
		Address:  add_address.Address,
		IsSet:    add_address.IsSet,
		Status:   0,
	}
	if err, id := address_u.Insert(); err != nil {
		this.Data["json"] = RetResource(false, types.CreateAddressFail, nil, "创建地址失败，请联系客服处理")
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = RetResource(true, types.ReturnSuccess, map[string]interface{}{"id": id}, "添加地址成功")
		this.ServeJSON()
		return
	}
}

// UpdAddress @Title UpdAddress finished
// @Description 修改地址和手机好码 UpdAddress
// @Success 200 status bool, data interface{}, msg string
// @router /upd_address [post]
func (this *UserAddressController) UpdAddress() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	_, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	var upd_address address.UserUpdAddressCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &upd_address); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := upd_address.UserUpdAddressCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	// 修改地址
	var crfr_addr_m models.UserAddress
	crfr_addr_m.Id = upd_address.AddressId
	crfr_addr, code, msg := crfr_addr_m.GetAddressById()
	if code != types.ReturnSuccess || msg != "" {
		this.Data["json"] = RetResource(false, types.AddressIsEmpty, err, "没有这个地址, 请检查之后提交")
		this.ServeJSON()
		return
	}
	crfr_addr.UserName = upd_address.UserName
	crfr_addr.Phone = upd_address.Phone
	crfr_addr.Address = upd_address.Address
	crfr_addr.IsSet = upd_address.IsSet
	ret_v := crfr_addr.UpdateAddressInfo()
	if ret_v == false {
		this.Data["json"] = RetResource(false, types.UpdateAddressFail, err, "修改地址失败")
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = RetResource(true, types.ReturnSuccess, err, "修改地址成功")
		this.ServeJSON()
		return
	}
}

// PostDelAddress @Title PostDelAddress finished
// @Description 删除地址 PostDelAddress
// @Success 200 status bool, data interface{}, msg string
// @router /del_address [post]
func (this *UserAddressController) PostDelAddress() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	_, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	var del_address address.UserAdddressDelCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &del_address); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := del_address.UserAdddressDelParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	var crfr_addr models.UserAddress
	crfr_addr.Id = del_address.AddressId
	err = crfr_addr.Delete()
	if err != nil {
		this.Data["json"] = RetResource(false, types.AddressIsEmpty, err, "删除地址失败")
		this.ServeJSON()
		return
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, err, "删除地址成功")
	this.ServeJSON()
	return
}

// GetAddressList @Title GetAddressList finished
// @Description 获取地址列表 GetAddressList
// @Success 200 status bool, data interface{}, msg string
// @router /address_list [post]
func (this *UserAddressController) GetAddressList() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	user_token, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	var crfr_addr models.UserAddress
	crfr_addr.UserId = user_token.Id
	address_list, code, msg := crfr_addr.GetUserAddressList()
	if code != types.ReturnSuccess {
		this.Data["json"] = RetResource(false, int(code), nil, msg)
		this.ServeJSON()
		return
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, address_list, "获取地址列表成功")
	this.ServeJSON()
	return
}