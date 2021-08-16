package api

import (
	"blockshop/models"
	"blockshop/types"
	"blockshop/types/help_desk"
	"blockshop/types/question"
	"encoding/json"
	"github.com/astaxie/beego"
	"strings"
)

type HelpDeskController struct {
	beego.Controller
}


// CreateHelpDesk @Title CreateHelpDesk
// @Description 提交反馈 CreateHelpDesk
// @Success 200 status bool, data interface{}, msg string
// @router /create_help_desk [post]
func (this *HelpDeskController) CreateHelpDesk() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	usr, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	hp_lp := help_desk.CreateHelpDeskReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &hp_lp); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	hd_desk := models.HelpDesk{
		UserId: usr.Id,
		Author: hp_lp.Author,
		Contract: hp_lp.Contract,
		HdTitle: hp_lp.HdTitle,
		HdDetail: hp_lp.HdDetail,
	}
	err, _ = hd_desk.Insert()
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "发送反馈失败")
		this.ServeJSON()
		return
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "创建反馈成功")
	this.ServeJSON()
	return
}


// HdList @Title HdList
// @Description 反馈列表 HdList
// @Success 200 status bool, data interface{}, msg string
// @router /hd_list [post]
func (this *HelpDeskController) HdList() {
	hd_lp := question.QsListCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &hd_lp); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	var qhd_list_ret []help_desk.HdListRep
	hd_list, total, err := models.GetHelpDeskList(hd_lp.Page, hd_lp.PageSize)
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "数据库错误")
		this.ServeJSON()
		return
	}
	for _, v := range hd_list {
		hd_ := help_desk.HdListRep{
			Id:v.Id,
			IsHand: v.IsHandle,
			HdTitle: v.HdTitle,
			HdTime: v.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		qhd_list_ret = append(qhd_list_ret, hd_)
	}
	data := map[string]interface{}{
		"total":     total,
		"qs_list":  qhd_list_ret,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取反馈列表成功")
	this.ServeJSON()
	return
}


// HdDetail @Title HdDetail
// @Description 反馈详情 HdDetail
// @Success 200 status bool, data interface{}, msg string
// @router /hd_detail [post]
func (this *HelpDeskController) HdDetail() {
	hd_dtl_p := help_desk.HdDetailCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &hd_dtl_p); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	hd, code, err := models.GetHelpDeskDetail(hd_dtl_p.HdId)
	if err != nil {
		this.Data["json"] = RetResource(false, code, nil, "数据库错误")
		this.ServeJSON()
		return
	}
	qs_detail := help_desk.HdDetailRep{
		Id: hd.Id,
		IsHand: hd.IsHandle,
		HdTitle: hd.HdTitle,
		HdDetai: hd.HdDetail,
		HdTime: hd.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, qs_detail, "获取反馈详情成功")
	this.ServeJSON()
	return
}
