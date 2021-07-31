package models

import (
	"blockshop/common"
	"github.com/astaxie/beego/orm"
)

type Message struct {
	BaseModel
	Id              int64      `orm:"pk;column(id);auto;size(11)" description:"工单ID" json:"id"`
	SendUserId      int64      `orm:"column(send_user_id);size(64);index" description:"发送人" json:"send_user_id"` //0:为平台发送;其他数字对应用户ID
	TargetUserId    int64  	   `orm:"column(target_user_id);size(64);index" description:"接收人" json:"target_user_id"`  // 0:为发送给平台;其他数字对应用户ID
	MsgType         int8       `orm:"column(msg_type);default(0);index"  description:"消息类型" json:"msg_type"`      // 0:文字消息；1:图片消息
	MsgContent      string     `orm:"column(content);type(text)" description:"消息内容" json:"content"`
}


func (this *Message) SearchField() []string {
	return []string{"msg_tilte"}
}

func (this *Message) TableName() string {
	return common.TableName("message")
}

func (this *Message) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

func (this *Message) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *Message) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

