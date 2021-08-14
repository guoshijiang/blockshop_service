package models

import (
	"blockshop/common"
	"blockshop/types"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type ForumReply struct {
	BaseModel
	Id             int64         `orm:"pk;column(id);auto;size(11)" description:"论坛回复ID" json:"id"`
	ForumId        int64         `orm:"column(forum_id)" description:"论坛ID" json:"forum_id"`
	FatherReplyId  int64         `orm:"column(father_reply_id)" description:"上级回复ID" json:"father_reply_id"`
	UserId         int64         `orm:"column(user_id)" description:"用户ID" json:"user_id"`
	Content        string        `orm:"column(content);type(text)" description:"回复内容" json:"content"`
	Views          int64         `orm:"column(abstract);default(0)" description:"回复浏览次数" json:"views"`
	Likes          int64         `orm:"column(likes);default(0)" description:"回复点赞次数" json:"likes"`
	UnLikes        int64         `orm:"column(un_likes);default(0)" description:"回复踩次数" json:"un_likes"`
	Answers        int64         `orm:"column(answers);default(0)" description:"回复次数" json:"answers"`
	IsCheck        int8          `orm:"column(is_check);default(0);index" description:"是否审核" json:"is_check"`  // 0:未审核 1:已审核
}

func (this *ForumReply) TableName() string {
	return common.TableName("forum_reply")
}

func (this *ForumReply) SearchField() []string {
	return []string{"user_name"}
}

func (this *ForumReply) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *User) ForumReply() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *ForumReply) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *ForumReply) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}

func GetFormLastestConment(form_id int64) (*ForumReply, int, error) {
	var form_reply_dtl ForumReply
	if err := orm.NewOrm().QueryTable(ForumReply{}).Filter("forum_id", form_id).OrderBy("-created_at").One(&form_reply_dtl); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &form_reply_dtl, types.ReturnSuccess, nil
}

func GetTotalReplyNum(form_id int64) int64 {
	total, _  := orm.NewOrm().QueryTable(&ForumReply{}).Filter("forum_id", form_id).Count()
	return total
}

func GetForumCommentList(page int64, page_size int64, forum_id int64) ([]*ForumReply, int, error) {
	offset := (page - 1) * page_size
	forum_reply_list := make([]*ForumReply, 0)
	forum_reply := orm.NewOrm().QueryTable(ForumReply{}).Filter("forum_id", forum_id)
	total, _ := forum_reply.Count()
	_, err := forum_reply.Limit(page_size, offset).All(&forum_reply_list)
	if err != nil {
		return nil, 0, errors.New("查询数据库失败")
	}
	return forum_reply_list, int(total), nil
}

func GetForumReplyList(forum_id int64) ([]*ForumReply, int, error) {
	forum_reply_list := make([]*ForumReply, 0)
	forum_reply := orm.NewOrm().QueryTable(ForumReply{}).Filter("father_reply_id", forum_id)
	total, _ := forum_reply.Count()
	_, err := forum_reply.All(&forum_reply_list)
	if err != nil {
		return nil, 0, errors.New("查询数据库失败")
	}
	return forum_reply_list, int(total), nil
}

func CreateForumCmtReply(user_id int64, forumt_id int64, father_reply_id int64, content string) (code int, msg string) {
	create_cmt_replu := ForumReply {
		ForumId:forumt_id,
		FatherReplyId: father_reply_id,
		UserId: user_id,
		Content: content,
	}
	err, _ := create_cmt_replu.Insert()
	if err != nil {
		return types.SystemDbErr, "创建评论失败"
	}
	return types.ReturnSuccess,  "创建评论成功"
}

func CommnetReplyLike(id int64, is_like int) (int) {
	var forum_reply ForumReply
	if err := orm.NewOrm().QueryTable(&ForumReply{}).Filter("Id", id).RelatedSel().One(&forum_reply); err != nil {
		return types.SystemDbErr
	}
	if is_like == 0 {
		forum_reply.Likes = forum_reply.Likes + 1
	} else {
		forum_reply.UnLikes = forum_reply.UnLikes + 1
	}

	return types.ReturnSuccess
}