package api

import (
	"blockshop/models"
	"blockshop/types"
	"github.com/astaxie/beego"
	"strings"
)

type MerchantCollectController struct {
	beego.Controller
}

// AddMctCollect @Title AddMctCollect finished
// @Description 收藏商家 AddMctCollect
// @Success 200 status bool, data interface{}, msg string
// @router /add_mct_collect [post]
func (this *MerchantCollectController) AddMctCollect() {
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
}


// GetMctCollectList @Title GetMctCollectList finished
// @Description 收藏商铺列表 GetMctCollectList
// @Success 200 status bool, data interface{}, msg string
// @router /get_mct_collect_list [post]
func (this *MerchantCollectController) GetMctCollectList() {
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
}


// RemoveMctCollect @Title RemoveMctCollect finished
// @Description 移除收藏商店 RemoveMctCollect
// @Success 200 status bool, data interface{}, msg string
// @router /remove_mct_collect [post]
func (this *MerchantCollectController) RemoveMctCollect() {
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
}
