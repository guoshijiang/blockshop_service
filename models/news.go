package models

import (
	"blockshop/common"
	"blockshop/types"
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

func (new *News) GetNewsById() (*News, int, string) {
	var news News
	err := news.Query().Filter("id", new.Id).One(&news)
	if err != nil {
		return nil, types.QueryNewsFail, "获取新闻详情失败"
	}
	return &news, types.ReturnSuccess, "获取新闻详情成功"
}

func GetNewsList(page, page_size int64) ([]*News, int, string) {
	var nl []*News
	filter := orm.NewOrm().QueryTable(&News{})
	total, err := filter.Count()
	if err != nil {
		return nil, types.QueryNewsFail, "获取新闻记录数量失败"
	}
	_, err = filter.Limit(page_size, page_size*(page-1)).OrderBy("-id").All(&nl)
	if err != nil {
		return nil, types.QueryNewsFail, "获取新闻列表失败"
	}
	return nl, int(total), "获取新闻列表成功"
}

func (new *News) NewsGetList(page, pageSize int, condition *orm.Condition) ([]*News, int64) {
	offset := (page - 1) * pageSize
	list := make([]*News, 0)
	query := orm.NewOrm().QueryTable(new.TableName())
	query = query.SetCond(condition)
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize,offset).All(&list)
	return list, total
}

func (new *News) NewsUpdate(fields ...string) error {
	if _, err := orm.NewOrm().Update(new, fields...); err != nil {
		return err
	}
	return nil
}

