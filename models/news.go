package models

import (
	"blockshop/common"
	"github.com/astaxie/beego/orm"
)

type News struct {
	BaseModel
	Id        int64     `orm:"pk;column(id);auto;size(11)" description:"公告ID" json:"id"`
	Title     string    `orm:"column(title);size(256)" description:"公告标题" json:"title"`
	Abstract  string    `orm:"column(abstract);type(text)" description:"公告摘要" json:"abstract"`
	Content   string    `orm:"column(content);type(text)" description:"公告内容" json:"content"`
	Image     string    `orm:"column(image);default(0)" description:"公告封面" json:"image"`
	Author    string    `orm:"column(author);default(blockshop)" description:"公告作者" json:"author"`
	Views     int64     `orm:"column(views);default(0)" description:"公告浏览次数" json:"views"`
	Likes     int64     `orm:"column(likes);default(0)"  description:"公告点赞次数" json:"likes"`
}

func (new *News) TableName() string {
	return common.TableName("news")
}

func (new *News) Insert() error {
	if _, err := orm.NewOrm().Insert(new); err != nil {
		return err
	}
	return nil
}

func (new *News) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new)
}

func (new *News) SearchField() []string {
  return []string{"title"}
}
