package models

import (
	"blockshop/common"
	"blockshop/types/forum"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type ForumCat struct {
	BaseModel
	Id            int64     `orm:"pk;column(id);auto;size(11)" description:"论坛类别ID" json:"id"`
	FatherCatId   int64     `orm:"column(father_cat_id)" description:"上级类别ID" json:"father_cat_id"`
	ForumCatLevel int8      `orm:"column(forum_cat_level)" description:"论坛级别" json:"forum_cat_level"` // 1: 第一级；2: 第二级
	Name          string    `orm:"column(name);size(512);index" description:"论坛分类名称" json:"name"`
	Introduce     string    `orm:"column(introduce);type(text)" description:"论坛类别介绍" json:"introduce"`
	Icon          string    `orm:"column(icon);size(150);default(/static/upload/default/user-default-60x60.png)" description:"论坛分类Icon" json:"icon"`
	IsShow        int8      `orm:"column(is_show);default(0)" description:"是否显示" json:"is_show"`   // 0 显示 1 不显示
}

func (this *ForumCat) TableName() string {
	return common.TableName("forum_cat")
}

func (this *ForumCat) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *ForumCat)SearchField() []string{
	return []string{"name"}
}

func (this *ForumCat) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *ForumCat) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *ForumCat) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *ForumCat) GetCatFirstLevel(req forum.ForumListReq) ([]*ForumCat, int64, error) {
	offset := (req.Page - 1) * req.PageSize
	forum_cat_list := make([]*ForumCat, 0)
	query_good := orm.NewOrm().QueryTable(ForumCat{}).Filter("forum_cat_level", 1)
	total, _ := query_good.Count()
	_, err := query_good.Limit(req.PageSize, offset).All(&forum_cat_list)
	if err != nil {
		return nil, 0, errors.New("查询数据库失败")
	}
	return forum_cat_list, total, nil
}

func (this *ForumCat) GetCatByFatherId(father_id int64) ([]*ForumCat, error) {
	forum_cat_list := make([]*ForumCat, 0)
	_, err := orm.NewOrm().QueryTable(ForumCat{}).Filter("id", father_id).All(&forum_cat_list)
	if err != nil {
		return nil, errors.New("查询数据库失败")
	}
	return forum_cat_list, nil
}