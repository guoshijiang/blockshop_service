package api

import (
	"blockshop/models"
	"blockshop/types"
	"blockshop/types/collect"
	"encoding/json"
	"github.com/astaxie/beego"
	"strings"
)

type BlackListController struct {
	beego.Controller
}

// AddBlackList @Title AddBlackList finished
// @Description 添加黑名单 AddBlackList
// @Success 200 status bool, data interface{}, msg string
// @router /add_blacklist [post]
func (this *BlackListController) AddBlackList() {
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
	blacklist_c := collect.BlackListReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &blacklist_c); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
}


// GetBlackListList @Title GetBlackListList finished
// @Description 黑名单列表 GetBlackListList
// @Success 200 status bool, data interface{}, msg string
// @router /get_blacklist_list [post]
func (this *BlackListController) GetBlackListList() {
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


// RemoveBlackList @Title RemoveBlackList finished
// @Description 移除黑名单 RemoveBlackList
// @Success 200 status bool, data interface{}, msg string
// @router /remove_blacklist [post]
func (this *BlackListController) RemoveBlackList() {
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

