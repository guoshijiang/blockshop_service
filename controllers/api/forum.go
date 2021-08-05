package api

import (
	"blockshop/models"
	"blockshop/types"
	"blockshop/types/forum"
	"encoding/json"
	"github.com/astaxie/beego"
)

type ForumController struct {
	beego.Controller
}

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
	forum_list, total, err := forum_cat.GetCatFirstLevel(forum_req)
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, nil, err.Error())
		this.ServeJSON()
		return
	}
	var f_level_one_list []forum.ForumLevelOneRep
	for _, value := range forum_list {
		forum_level_list, err := forum_cat.GetCatByFatherId(value.Id)
		if err != nil {
			panic(err)
		}
		var f_level_list []forum.ChildFormRep
		for _, value_fl := range forum_level_list {
			fl := forum.ChildFormRep {
				Id: value_fl.Id,
				Title: value_fl.Name,
				Icon: value_fl.Icon,
			}
			f_level_list = append(f_level_list,  fl)
		}
		fllop := forum.ForumLevelOneRep {
			Id:        value.Id,
			Title:     value.Name,
			ThemeNum: int64(len(forum_level_list)),
			TopicNum:  100,
			Abstruct:  value.Introduce,
			ChildForm: f_level_list,
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
	this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "获取论坛子模块成功")
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
	this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "发表评论回复成功")
	this.ServeJSON()
	return
}


// CreateForumTopic @Title CreateForumTopic
// @Description 发论坛帖子 CreateForumTopic
// @Success 200 status bool, data interface{}, msg string
// @router /create_forum_topic [post]
func (this *ForumController) CreateForumTopic() {
	this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "发布帖子成功")
	this.ServeJSON()
	return
}


