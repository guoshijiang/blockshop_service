package models

import (
	"blockshop/common"
	"github.com/astaxie/beego/orm"
)

type ForumReply struct {
	BaseModel
	Id             int64         `orm:"column(user_name)" description:"ID" json:"id"`
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
