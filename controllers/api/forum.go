package api

import (
	"blockshop/models"
	"blockshop/types"
	"blockshop/types/forum"
	"encoding/json"
	"github.com/astaxie/beego"
	"strings"
)

type ForumController struct {
	beego.Controller
}

const (
	LevelOneCatId = 0
	LevelOne      = 1
	LevelTwo 	  = 2
)

// ForumMainTopicList @Title ForumMainTopicList
// @Description 论坛分类列表 ForumMainTopicList
// @Success 200 status bool, data interface{}, msg string
// @router /forum_main_topic_list [post]
func (this *ForumController) ForumMainTopicList() {
	forum_req := forum.ForumListReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &forum_req); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	var forum_cat models.ForumCat
	forum_list, total, err := forum_cat.GetCatLevel(forum_req.Page, forum_req.PageSize, LevelOneCatId, LevelOne)
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, nil, err.Error())
		this.ServeJSON()
		return
	}
	var f_level_one_list []forum.ForumLevelOneRep
	for _, value := range forum_list {
		forum_level_list, lst_id_id, err := forum_cat.GetCatByFatherId(value.Id)
		var total_form_top int64
		var f_level_list []forum.ChildFormRep
		for _, value_fl := range forum_level_list {
			fl := forum.ChildFormRep {
				Id: value_fl.Id,
				Title: value_fl.Name,
				Icon: value_fl.Icon,
			}
			total_form_top += models.GetForumByCatId(value_fl.Id)
			f_level_list = append(f_level_list,  fl)
		}
		var lst_f_reply *forum.LastestForumRep
		lst_forum, _, err := models.GetLastestForumByCatId(lst_id_id)
		if err != nil {
			lst_f_reply = nil
		} else {
			lst_forum_reply, _, err := models.GetFormLastestConment(lst_forum.Id)
			if err != nil {
				lst_f_reply = nil
			} else {
				query_user, err := models.GetUserById(lst_forum_reply.UserId)
				if err != nil {
					this.Data["json"] = RetResource(false, types.UserIsNotExist, err, "用户不存在")
					this.ServeJSON()
					return
				}
				_lst_f_reply := &forum.LastestForumRep {
					UserId: lst_forum_reply.UserId,
					FormId:lst_forum.Id,
					UserName: query_user.UserName,
					UserPhoto: query_user.Avator,
					LstComment: lst_forum_reply.Content,
					DataTime: lst_forum_reply.CreatedAt.Format("2006-01-02 15:04:05"),
				}
				lst_f_reply = _lst_f_reply
			}
		}
		fllop := forum.ForumLevelOneRep {
			Id:        value.Id,
			Title:     value.Name,
			ThemeNum: int64(len(forum_level_list)),
			TopicNum:  total_form_top,
			Abstruct:  value.Introduce,
			ChildForm: f_level_list,
			LastestForm: lst_f_reply,
		}
		f_level_one_list = append(f_level_one_list, fllop)
	}
	data := map[string]interface{}{
		"total": total,
		"form_lst": f_level_one_list,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取论坛父板块成功")
	this.ServeJSON()
	return
}

// ForumChildTopicList @Title ForumChildTopicList
// @Description 论坛子模块接口 ForumChildTopicList
// @Success 200 status bool, data interface{}, msg string
// @router /forum_child_topic_list [post]
func (this *ForumController) ForumChildTopicList() {
	forum_req := forum.ForumChildListReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &forum_req); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	var forum_cat models.ForumCat
	forum_cat_list, total, err := forum_cat.GetCatLevel(forum_req.Page, forum_req.PageSize, forum_req.CatId, LevelTwo)
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, nil, err.Error())
		this.ServeJSON()
		return
	}
	var f_level_child_list []forum.ForumLevelChildRep
	for _, value := range forum_cat_list {
		var total_reply_num int64
		forum_list, _, err := models.GetForumListByCatId(value.Id)
		if err != nil {
			total_reply_num = 0
		} else {
			for _, fm := range forum_list {
				total_reply_num += models.GetTotalReplyNum(fm.Id)
			}
		}
		var lst_f_reply *forum.LastestForumRep
		lst_forum, _, err := models.GetLastestForumByCatId(value.Id)
		if err != nil {
			lst_f_reply = nil
		} else {
			lst_forum_reply, _, err := models.GetFormLastestConment(lst_forum.Id)
			if err != nil {
				lst_f_reply = nil
			} else {
				query_user, err := models.GetUserById(lst_forum_reply.UserId)
				if err != nil {
					this.Data["json"] = RetResource(false, types.UserIsNotExist, err, "用户不存在")
					this.ServeJSON()
					return
				}
				_lst_f_reply := &forum.LastestForumRep {
					UserId: lst_forum_reply.UserId,
					FormId:lst_forum.Id,
					UserName: query_user.UserName,
					UserPhoto: query_user.Avator,
					LstComment: lst_forum_reply.Content,
					DataTime: lst_forum_reply.CreatedAt.Format("2006-01-02 15:04:05"),
				}
				lst_f_reply = _lst_f_reply
			}
		}
		f_level_child := forum.ForumLevelChildRep{
			Id: value.Id,
			Title: value.Name,
			TopicNum: models.GetForumByCatId(value.Id),
			ReplyNum: total_reply_num,
			LastestForm: lst_f_reply,
		}
		f_level_child_list = append(f_level_child_list, f_level_child)
	}
	data := map[string]interface{}{
		"total": total,
		"form_lst": f_level_child_list,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取论坛子模块成功")
	this.ServeJSON()
	return
}

// ForumCTopicList @Title ForumCTopicList
// @Description 论坛帖子列表 ForumCTopicList
// @Success 200 status bool, data interface{}, msg string
// @router /forum_topic_list [post]
func (this *ForumController) ForumCTopicList() {
	forum_topic_req := forum.ForumTopicListReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &forum_topic_req); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	ft_list, total, err := models.GetForumList(int64(forum_topic_req.Page), int64(forum_topic_req.PageSize), forum_topic_req.CatId, forum_topic_req.IsFather)
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "获取论坛帖子数据失败")
		this.ServeJSON()
		return
	}
	var ft_list_rep []*forum.ForumTopicListRep
	for _, value := range ft_list {
		user, _ := models.GetUserById(value.UserId)
		ft := &forum.ForumTopicListRep{
			UserId: value.UserId,
			FormId: value.Id,
			UserName: user.UserName,
			UserPhoto: user.Avator,
			Title: value.Title,
			DataTime: value.CreatedAt.Format("2006-01-02 15:04:05"),
			Views: value.Views,
			Likes: value.Likes,
			UnLikes: value.UnLikes,
			Answers: value.Answers,
		}
		ft_list_rep = append(ft_list_rep, ft)
	}
	data := map[string]interface{}{
		"total": total,
		"form_lst": ft_list_rep,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取论坛帖子列表成功")
	this.ServeJSON()
	return
}

// ForumCTopicDetail @Title ForumCTopicDetail
// @Description 论坛帖子 ForumCTopicDetail
// @Success 200 status bool, data interface{}, msg string
// @router /forum_topic_detail [post]
func (this *ForumController) ForumCTopicDetail() {
	forum_detail_req := forum.ForumTopicDetailReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &forum_detail_req); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	forum_, code, err := models.GetForumDetail(forum_detail_req.ForumId)
	if err != nil {
		this.Data["json"] = RetResource(false, code, nil, "获取论坛帖子数据失败")
		this.ServeJSON()
		return
	}
	user, _ := models.GetUserById(forum_.UserId)
	data := map[string]interface{}{
		"id": forum_.Id,
		"title": forum_.Title,
		"content": forum_.Content,
		"datetime": forum_.CreatedAt.Format("2006-01-02 15:04:05"),
		"user_id": forum_.UserId,
		"user_name": user.UserName,
		"photo": user.Avator,
		"views": forum_.Views,
		"likes": forum_.Likes,
		"answers": forum_.Answers,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取论坛子详情成功")
	this.ServeJSON()
	return
}


// ForumTopicCommentList @Title ForumTopicCommentList
// @Description 论坛帖子评论 ForumTopicCommentList
// @Success 200 status bool, data interface{}, msg string
// @router /forum_topic_comment_list [post]
func (this *ForumController) ForumTopicCommentList() {
	forum_reply_req := forum.ForumTopicDetailReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &forum_reply_req); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	cmt_list, total, err := models.GetForumCommentList(int64(forum_reply_req.Page), int64(forum_reply_req.PageSize), forum_reply_req.ForumId)
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "获取论坛评论数据失败")
		this.ServeJSON()
		return
	}
	var comment_reply_list []*forum.ForumCommentListRep
	for _, cmt := range cmt_list {
		user, _ := models.GetUserById(cmt.UserId)
		var reply_list []forum.ForumReply
		reply_list_dt, _, err := models.GetForumReplyList(cmt.Id)
		if err != nil {
			reply_list = nil
		} else {
			for _, reply := range reply_list_dt {
				user, _ := models.GetUserById(reply.UserId)
				f_rly := forum.ForumReply{
					Id: reply.Id,
					UserName: user.UserName,
					UserPhoto: user.Avator,
					Reply: reply.Content,
					UnLikes: reply.UnLikes,
					Likes: reply.Likes,
					Datetime: reply.CreatedAt.Format("2006-01-02 15:04:05"),
				}
				reply_list = append(reply_list, f_rly)
			}
		}
		comment_reply := &forum.ForumCommentListRep {
			Id: cmt.Id,
			UserName: user.UserName,
			UserPhoto: user.Avator,
			Comment: cmt.Content,
			UnLikes: cmt.UnLikes,
			Likes: cmt.Likes,
			Datetime: cmt.CreatedAt.Format("2006-01-02 15:04:05"),
			Reply: reply_list,
		}
		comment_reply_list = append(comment_reply_list, comment_reply)
	}
	data := map[string]interface{}{
		"total": total,
		"form_lst": comment_reply_list,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取帖子评论成功")
	this.ServeJSON()
	return
}

// GetForumCatTopic @Title GetForumCatTopic
// @Description 获取分类和主题 CreateForumTopic
// @Success 200 status bool, data interface{}, msg string
// @router /get_ftc_list [post]
func (this *ForumController) GetForumCatTopic() {
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
	tc_req := forum.ForumTopicCatReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &tc_req); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	fct_list, _, err := models.GetTopicCatList(tc_req.TcName, tc_req.IsTc)
	var tc_list []*forum.FmTopicCatListRep
	for _, value := range fct_list {
		tcl := &forum.FmTopicCatListRep{
			TcId:value.Id,
			TcName: value.Name,
		}
		tc_list = append(tc_list, tcl)
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, tc_list, "获取类别成功")
	this.ServeJSON()
	return
}


// CreateForumTopic @Title CreateForumTopic
// @Description 发布帖子 CreateForumTopic
// @Success 200 status bool, data interface{}, msg string
// @router /create_forum_topic [post]
func (this *ForumController) CreateForumTopic() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	user_, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	create_frm := forum.CreateForumReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &create_frm); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	father_cat_id, _ := models.CreateOrGetFcat(create_frm.CatName, 0, 1)
	cat_id, _ := models.CreateOrGetFcat(create_frm.TopName, father_cat_id, 2)
	code, msg  := models.CreateForum(user_.Id, father_cat_id, cat_id, create_frm.Title, create_frm.Abstract, create_frm.Content)
	if code == types.ReturnSuccess {
		this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "发布帖子成功")
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = RetResource(false, types.SystemDbErr, nil, msg)
		this.ServeJSON()
		return
	}
}

// ForumTopicCommentReply @Title ForumTopicCommentReply
// @Description 论坛帖子评论和回复 ForumTopicCommentReply
// @Success 200 status bool, data interface{}, msg string
// @router /forum_comment_reply [post]
func (this *ForumController) ForumTopicCommentReply() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	user_, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	create_cmt_reply := forum.CreateCmtReplyReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &create_cmt_reply); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	code, msg  := models.CreateForumCmtReply(user_.Id, create_cmt_reply.ForumId, create_cmt_reply.FatherCmtId, create_cmt_reply.CtmReply)
	if code == types.ReturnSuccess {
		this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "评论成功")
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = RetResource(false, types.SystemDbErr, nil, msg)
		this.ServeJSON()
		return
	}
}

// ForumTopicLike @Title ForumTopicLike
// @Description 帖子点赞 ForumTopicLike
// @Success 200 status bool, data interface{}, msg string
// @router /forum_like [post]
func (this *ForumController) ForumTopicLike() {
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
	forum_topic_like := forum.ForumTopiceUnOrLikeReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &forum_topic_like); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	code := models.ForumTopicLike(forum_topic_like.ForumId, forum_topic_like.IsLike)
	if code == types.ReturnSuccess {
		this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "点赞成功")
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "点赞失败")
		this.ServeJSON()
		return
	}
}

// CommentReplyLike @Title CommentReplyLike
// @Description 评论回复点赞 CommentReplyLike
// @Success 200 status bool, data interface{}, msg string
// @router /comment_reply_like [post]
func (this *ForumController) CommentReplyLike() {
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
	cmtr_like := forum.CommentReplyUnOrLikeReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &cmtr_like); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	code := models.CommnetReplyLike(cmtr_like.CmtReplyId, cmtr_like.IsLike)
	if code == types.ReturnSuccess {
		this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "点赞评论回复成功")
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "点赞评论回复失败")
		this.ServeJSON()
		return
	}
}