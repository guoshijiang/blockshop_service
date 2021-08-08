package models

import (
	"blockshop/common"
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


func (this *ForumCat) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
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

func (this *ForumCat) GetCatLevel(page int, page_size int, cat_id int64, forum_cat_level int8) ([]*ForumCat, int64, error) {
	offset := (page - 1) * page_size
	forum_cat_list := make([]*ForumCat, 0)
	query_good := orm.NewOrm().QueryTable(ForumCat{}).Filter("forum_cat_level", forum_cat_level)
	if cat_id >= 1 {
		query_good = query_good.Filter("father_cat_id", cat_id)
	}
	total, _ := query_good.Count()
	_, err := query_good.Limit(page_size, offset).All(&forum_cat_list)
	if err != nil {
		return nil, 0, errors.New("查询数据库失败")
	}
	return forum_cat_list, total, nil
}

func (this *ForumCat) GetCatByFatherId(father_id int64) ([]*ForumCat, int64, error) {
	forum_cat_list := make([]*ForumCat, 0)
	var forum_cat ForumCat
	query := orm.NewOrm().QueryTable(ForumCat{}).Filter("father_cat_id", father_id).OrderBy("-id")
	_, err := query.All(&forum_cat_list)
	err = query.One(&forum_cat)
	if err != nil {
		return nil, 0, errors.New("查询数据库失败")
	}
	return forum_cat_list, forum_cat.Id, nil
}

func GetTopicCatList(tc_name string, is_tc int) ([]*ForumCat, int64, error) {
	tc_list := make([]*ForumCat, 0)
	if is_tc == 1 {  // 类别
		_, err := orm.NewOrm().QueryTable(ForumCat{}).Filter("forum_cat_level", 1).Filter("name__contains", tc_name).OrderBy("-id").All(&tc_list)
		if err != nil {
			return nil, 0, errors.New("查询数据库失败")
		}
	} else if is_tc == 2 {
		_, err := orm.NewOrm().QueryTable(ForumCat{}).Filter("forum_cat_level", 2).Filter("name__contains", tc_name).OrderBy("-id").All(&tc_list)
		if err != nil {
			return nil, 0, errors.New("查询数据库失败")
		}
	}
	return tc_list, 0, nil
}

func CreateOrGetFcat(tc_name string, father_cat_id int64, forum_cat_level int8) (int64,  error) {
	var forum_cat ForumCat
	if err := orm.NewOrm().QueryTable(ForumCat{}).Filter("name", tc_name).Filter("forum_cat_level", forum_cat_level).RelatedSel().One(&forum_cat); err != nil {
		create_forum_cat := ForumCat {
			FatherCatId: father_cat_id,
			ForumCatLevel: forum_cat_level,
			Name: tc_name,
			Introduce: "",
			Icon: "",
		}
		err, id := create_forum_cat.Insert()
		if err != nil {
			return 0, errors.New("创建论坛失败")
		}
		return  id, nil
	}
	return forum_cat.Id, nil
}