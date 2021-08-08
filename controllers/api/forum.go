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
					DataTime: lst_forum_reply.CreatedAt.String(),
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
					DataTime: lst_forum_reply.CreatedAt.String(),
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
	this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "获取论坛帖子列表成功")
	this.ServeJSON()
	return
}

// ForumCTopicDetail @Title ForumCTopicDetail
// @Description 论坛帖子 ForumCTopicDetail
// @Success 200 status bool, data interface{}, msg string
// @router /forum_topic_detail [post]
func (this *ForumController) ForumCTopicDetail() {
	this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "获取论坛子板块成功")
	this.ServeJSON()
	return
}

// ForumTopicCommentList @Title ForumTopicCommentList
// @Description 论坛帖子评论 ForumTopicCommentList
// @Success 200 status bool, data interface{}, msg string
// @router /forum_topic_comment_list [post]
func (this *ForumController) ForumTopicCommentList() {
	this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "获取帖子评论成功")
	this.ServeJSON()
	return
}

// ForumTopicCommentReply @Title ForumTopicCommentReply
// @Description 论坛帖子评论和回复 ForumTopicCommentReply
// @Success 200 status bool, data interface{}, msg string
// @router /forum_topic_comment_reply [post]
func (this *ForumController) ForumTopicCommentReply() {
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
	this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "发表评论回复成功")
	this.ServeJSON()
	return
}


// CreateForumTopic @Title CreateForumTopic
// @Description 发论坛帖子 CreateForumTopic
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
	_, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "发布帖子成功")
	this.ServeJSON()
	return
}


