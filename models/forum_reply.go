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