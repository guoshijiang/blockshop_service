package api

import (
	"blockshop/models"
	"blockshop/types"
	"blockshop/types/collect"
	"encoding/json"
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
	mct_c := collect.MerchantCollectReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &mct_c); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	msg, code := models.AddMerchantCollect(mct_c)
	if code != types.ReturnSuccess {
		this.Data["json"] = RetResource(false, types.SystemDbErr, err, msg)
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = RetResource(true, types.ReturnSuccess, err, "收藏店铺成功")
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
	page_s := types.PageSizeData{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &page_s); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	mct_c_list, total := models.MerchantCollectList(page_s.Page, page_s.PageSize)
	var mctc_list []*collect.MerchantListRep
	for _, value := range mct_c_list {
		var mct_name string
		var date_time string
		mct, _, _ := models.GetMerchantDetail(value.CtMctId)
		if mct != nil {
			mct_name = mct.MerchantName
			date_time = mct.CreatedAt.Format("2006-01-02 15:04:05")
		} else {
			mct_name = ""
			date_time = ""
		}
		bl := &collect.MerchantListRep{
			MctCollectId: value.Id,
			MerchantId: value.CtMctId,
			MerchanName: mct_name,
			DateTime: date_time,
		}
		mctc_list =append(mctc_list, bl)
	}
	data := map[string]interface{}{
		"total": total,
		"data": mctc_list,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取收藏店铺列表成功")
	this.ServeJSON()
	return
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
	mct_c_del := collect.MerchantCollectDelReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &mct_c_del); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	msg, code := models.RemoveMerchantCollect(mct_c_del.MctCollectId)
	if code != types.ReturnSuccess {
		this.Data["json"] = RetResource(false, types.SystemDbErr, err, msg)
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = RetResource(true, types.ReturnSuccess, err, "移除收藏商铺成功")
		this.ServeJSON()
		return
	}
}
