package api


import (
	"encoding/json"
	"blockshop/models"
	"blockshop/types"
	type_comment "blockshop/types/comment"
	"github.com/astaxie/beego"
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
		Star: add_comment.Star,
		Content: add_comment.Content,
		ImgOneId: add_comment.ImgOneId,
		ImgTwoId: add_comment.ImgTwoId,
		ImgThreeId: add_comment.ImgThreeId,
	}
	err, id := cmt.Insert()
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, err.Error(), "添加评论失败")
		this.ServeJSON()
		return
	} else {
		order_detail, _, _ := models.GetGoodsOrderDetail(add_comment.OrderId)
		order_detail.IsComment  = 1
		err = order_detail.Update()
		if err != nil {
			this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "更新评论状态失败")
			this.ServeJSON()
			return
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
	clst, total, err := models.GetGoodsCommentList(clist.Page, clist.PageSize, clist.GoodsId)
	if err !=  nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, nil, err.Error())
		this.ServeJSON()
		return
	}
	var cmt_list []type_comment.CommentListRep
	for _, v := range clst {
		var img_mdl_one, img_mdl_two, img_mdl_three   models.ImageFile
		img_mdl_one.Id = v.ImgOneId
		img_mdl_two.Id = v.ImgTwoId
		img_mdl_three.Id = v.ImgThreeId
		var one_url, two_url, three_url string
		ImgOne_img, _, _ := img_mdl_one.GetImageById(v.ImgOneId)
		if ImgOne_img != nil {
			one_url = ImgOne_img.Url
		} else {
			one_url = ""
		}
		ImgTwo_img, _, _ := img_mdl_one.GetImageById(v.ImgTwoId)
		if ImgTwo_img != nil {
			two_url = ImgTwo_img.Url
		} else {
			two_url = ""
		}
		ImgThree_img, _, _ := img_mdl_one.GetImageById(v.ImgThreeId)
		if ImgThree_img != nil {
			three_url = ImgThree_img.Url
		} else {
			three_url = ""
		}
		image_path := beego.AppConfig.String("img_root_path")
		user_s, _ := models.GetUserById(v.UserId)
		cl := type_comment.CommentListRep{
			CommentId: v.Id,
			UserName: user_s.UserName,
			UserPho: user_s.Avator,
			GoodsId: v.GoodsId,
			UserId: v.UserId,
			Star: v.Star,
			Content: v.Content,
			ImgOne: image_path + one_url,
			ImgTwo: image_path + two_url,
			ImgThree: image_path + three_url,
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

