package models

import (
	"blockshop/common"
	"github.com/astaxie/beego/orm"
)

type Forum struct { // 论坛表
	BaseModel
	Id             int64         `orm:"pk;column(id);auto;size(11)" description:"论坛ID" json:"id"`
	UserId         int64         `orm:"column(user_id)" description:"用户ID" json:"user_id"`
	Title          string        `orm:"column(title);size(128)" description:"论坛标题" json:"title"`
	Abstract       string        `orm:"column(abstract);type(text)" description:"论坛摘要" json:"abstract"`
	Content        string        `orm:"column(content);type(text)" description:"论坛内容" json:"content"`
	Views          int64         `orm:"column(views);default(0)" description:"论坛浏览次数" json:"views"`
	Likes          int64         `orm:"column(likes);default(0)" description:"论坛点赞次数" json:"likes"`
	Answers        int64         `orm:"column(answers);default(0)" description:"论坛评论次数" json:"answers"`
}

func (this *Forum) TableName() string {
	return common.TableName("forum")
}

func (this *Forum) SearchField() []string {
	return []string{"title"}
}

func (this *Forum) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *Forum) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *Forum) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *Forum) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}
