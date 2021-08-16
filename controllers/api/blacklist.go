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
	msg, code := models.AddBlackList(blacklist_c)
	if code != types.ReturnSuccess {
		this.Data["json"] = RetResource(false, types.SystemDbErr, err, msg)
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = RetResource(true, types.ReturnSuccess, err, "加入黑名单成功")
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
	t_user, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	page_s := types.PageSizeData{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &page_s); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	bl_lst, total := models.BlackListList(page_s.Page, page_s.PageSize, t_user.Id)
	var bl_list []*collect.BlackListRep
	for _, value := range bl_lst {
		var mct_name string
		var date_time string
		mct, _, _ := models.GetMerchantDetail(value.BlackMctId)
		if mct != nil {
			mct_name = mct.MerchantName
			date_time = mct.CreatedAt.Format("2006-01-02 15:04:05")
		} else {
			mct_name = ""
			date_time = ""
		}
		bl := &collect.BlackListRep{
			BlId: value.Id,
			MerchantId: value.BlackMctId,
			MerchanName: mct_name,
			DateTime: date_time,
		}
		bl_list =append(bl_list, bl)
	}
	data := map[string]interface{}{
		"total": total,
		"data": bl_list,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取屏蔽商家列表成功")
	this.ServeJSON()
	return
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
	blacklist_c := collect.BlackListDelReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &blacklist_c); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	msg, code := models.RemoveBlackList(blacklist_c.BlackListId)
	if code != types.ReturnSuccess {
		this.Data["json"] = RetResource(false, types.SystemDbErr, err, msg)
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = RetResource(true, types.ReturnSuccess, err, "移除黑名单成功")
		this.ServeJSON()
		return
	}
}

