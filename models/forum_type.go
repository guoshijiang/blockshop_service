package models

import (
	"blockshop/common"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type ForumTpye struct {
	BaseModel
	Id           int64     `orm:"pk;column(id);auto;size(11)" description:"论坛类别ID" json:"id"`
	Name         string    `orm:"column(name);size(512);index" description:"论坛分类名称" json:"name"`
	Icon         string    `orm:"column(icon);size(150);default(/static/upload/default/user-default-60x60.png)" description:"论坛分类Icon" json:"icon"`
	IsShow       int8      `orm:"column(is_show);default(0)" description:"是否显示" json:"is_show"`   // 0 显示 1 不显示
}

func (this *ForumTpye) TableName() string {
	return common.TableName("forum_type")
}

func (this *ForumTpye) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *ForumTpye)SearchField() []string{
	return []string{"name"}
}

func (this *ForumTpye) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *ForumTpye) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *ForumTpye) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}
