package api

import (
	"blockshop/models"
	"blockshop/types"
	type_comment "blockshop/types/comment"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
)

type CommentController struct {
	beego.Controller
}

// AddCommet @Title AddCommet finished
// @Description 添加评论 AddCommet
// @Success 200 status bool, data interface{}, msg string
// @router /add_comment [post]
func (this *CommentController) AddCommet() {
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
	var add_comment type_comment.AddCommentReq
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &add_comment); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := add_comment.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	if requestUser.Id != add_comment.UserId {
		this.Data["json"] = RetResource(false, types.UserIsNotExist, err, "Token 和用户不匹配，拒绝添加地址")
		this.ServeJSON()
		return
	}
	cmt := models.GoodsComment{
		GoodsId: add_comment.GoodsId,
		UserId: add_comment.UserId,
		QualityStar: add_comment.QualityStar,
		ServiceStar: add_comment.ServiceStar,
		TradeStar: add_comment.TradeStar,
		Content: add_comment.Content,
		ImgOneUrl: add_comment.ImgOne,
		ImgTwoUrl: add_comment.ImgTwo,
		ImgThreeUrl: add_comment.ImgThree,
		MerchantId: add_comment.MerchantId,
	}
	err, id := cmt.Insert()
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, err.Error(), "添加评论失败")
		this.ServeJSON()
		return
	} else {
		  //统计评论
		  _, err := new(models.MerchantStat).UpdateByMerchant(models.MerchantStateCountRaw{
			MerchantId: add_comment.MerchantId,
			QualityStar: add_comment.QualityStar,
			ServiceStar: add_comment.ServiceStar,
			TradeStar: add_comment.TradeStar,
		})
		if err != nil {
		  this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "评论状态统计失败")
		  this.ServeJSON()
		  return
		}
		order_detail, _, _ := models.GetGoodsOrderDetail(add_comment.OrderId)
		order_detail.IsComment  = 1
		order_detail.OrderStatus = models.OrederFinish
		err = order_detail.Update()
		if err != nil {
			this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "更新评论状态失败")
			this.ServeJSON()
			return
		}
		var order_process models.OrderProcess
		err = orm.NewOrm().QueryTable(models.OrderProcess{}).Filter("order_id", add_comment.OrderId).RelatedSel().One(&order_process)
		if err == nil {
			order_process.Process = models.ProcesFinish
			_ = order_detail.Update()
		}
		this.Data["json"] = RetResource(true, types.ReturnSuccess, map[string]interface{}{"id": id}, "添加评论成功")
		this.ServeJSON()
		return
	}
}

// DelCommet @Title DelCommet finished
// @Description 删除评论 DelCommet
// @Success 200 status bool, data interface{}, msg string
// @router /del_commet [post]
func (this *CommentController) DelCommet() {
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
	var del_comment type_comment.DelCommentReq
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &del_comment); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := del_comment.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	if requestUser.Id != del_comment.UserId {
		this.Data["json"] = RetResource(false, types.UserIsNotExist, err, "Token 和用户不匹配，拒绝添加地址")
		this.ServeJSON()
		return
	}
	var gdc models.GoodsComment
	gdc.Id = del_comment.CommentId
	err = gdc.Delete()
	if err != nil {
		this.Data["json"] = RetResource(false, types.AddressIsEmpty, err, "删除评论失败")
		this.ServeJSON()
		return
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, err, "删除评论成功")
	this.ServeJSON()
	return
}

// GetCommentList @Title GetCommentList finished
// @Description 获取评论列表 GetCommentList
// @Success 200 status bool, data interface{}, msg string
// @router /comment_list [post]
func (this *CommentController) GetCommentList() {
	var clist type_comment.CommentListReq
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &clist); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := clist.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	clst, total, err := models.GetGoodsCommentList(clist.Page, clist.PageSize, clist.GoodsId, clist.MerchantId, clist.CmtStatus)
	if err !=  nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, nil, err.Error())
		this.ServeJSON()
		return
	}
	var cmt_list []type_comment.CommentListRep
	for _, v := range clst {
		//image_path := beego.AppConfig.String("img_root_path")
		user_s, _ := models.GetUserById(v.UserId)
		cl := type_comment.CommentListRep{
			CommentId: v.Id,
			UserName: user_s.UserName,
			UserPho: user_s.Avator,
			GoodsId: v.GoodsId,
			UserId: v.UserId,
			QualityStar: v.QualityStar,
			ServiceStar: v.ServiceStar,
			TradeStar: v.TradeStar,
			Content: v.Content,
			ImgOne: v.ImgOneUrl,
			ImgTwo: v.ImgTwoUrl,
			ImgThree: v.ImgThreeUrl,
			CreateTime: v.CreatedAt,
		}
		cmt_list = append(cmt_list, cl)
	}
	data := map[string]interface{}{
		"total": total,
		"cmt_lst": cmt_list,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取评论列表成功")
	this.ServeJSON()
	return
}

