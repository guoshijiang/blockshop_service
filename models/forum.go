package models

import (
	"blockshop/common"
	"blockshop/types"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type Forum struct { // 论坛表
	BaseModel
	Id             int64         `orm:"pk;column(id);auto;size(11)" description:"论坛ID" json:"id"`
	UserId         int64         `orm:"column(user_id)" description:"用户ID" json:"user_id"`
	FathrerCatId   int64  		 `orm:"column(fathrer_cat_id)" description:"顶级分类ID" json:"fathrer_cat_id"`
	CatId   	   int64         `orm:"column(cat_id)" description:"类别ID" json:"cat_id"`
	Title          string        `orm:"column(title);size(128)" description:"论坛标题" json:"title"`
	Abstract       string        `orm:"column(abstract);type(text)" description:"论坛摘要" json:"abstract"`
	Content        string        `orm:"column(content);type(text)" description:"论坛内容" json:"content"`
	Views          int64         `orm:"column(views);default(0)" description:"论坛浏览次数" json:"views"`
	Likes          int64         `orm:"column(likes);default(0)" description:"论坛点赞次数" json:"likes"`
	Answers        int64         `orm:"column(answers);default(0)" description:"论坛评论次数" json:"answers"`
	IsCheck        int8          `orm:"column(is_check);default(0);index" description:"是否审核" json:"is_check"`  // 0:未审核 1:已审核
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

// 获取论坛数量
func GetForumByCatId(cat_id int64) int64 {
	total, _  := orm.NewOrm().QueryTable(&Forum{}).Filter("cat_id", cat_id).Count()
	return total
}

func GetForumListByCatId(cat_id int64) ([]*Forum, int, error) {
	forum_list := make([]*Forum, 0)
	query_forum := orm.NewOrm().QueryTable(Forum{}).Filter("cat_id", cat_id)
	total, _ := query_forum.Count()
	_, err := query_forum.All(&forum_list)
	if err != nil {
		return nil, 0, errors.New("查询数据库失败")
	}
	return forum_list, int(total), nil
}

func GetLastestForumByCatId(cat_id int64) (*Forum, int, error) {
	var forum_dtl Forum
	if err := orm.NewOrm().QueryTable(Forum{}).Filter("cat_id", cat_id).OrderBy("-created_at").One(&forum_dtl); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &forum_dtl, types.ReturnSuccess, nil
}

// 是不是父级，0:不是，1:是
func GetForumList(page int64, page_size int64, cat_id int64, is_father int8) ([]*Forum, int, error) {
	offset := (page - 1) * page_size
	forum_list := make([]*Forum, 0)
	query_forum := orm.NewOrm().QueryTable(Forum{}).Filter("is_removed", 0)
	if is_father == 1 {
		query_forum = orm.NewOrm().QueryTable(Forum{}).Filter("fathrer_cat_id", cat_id)
	} else {
		query_forum = orm.NewOrm().QueryTable(Forum{}).Filter("cat_id", cat_id)
	}
	total, _ := query_forum.Count()
	_, err := query_forum.Limit(page_size, offset).All(&forum_list)
	if err != nil {
		return nil, 0, errors.New("查询数据库失败")
	}
	return forum_list, int(total), nil
}

func GetForumDetail(id int64) (*Forum, int, error) {
	var forum Forum
	if err := orm.NewOrm().QueryTable(Forum{}).Filter("Id", id).RelatedSel().One(&forum); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &forum, types.ReturnSuccess, nil
}

func CreateForum(user_id int64, father_cat_id int64, cat_id int64, title string, abstract string, content string) (code int, msg string) {
	create_forum := Forum {
		UserId: user_id,
		FathrerCatId: father_cat_id,
		CatId: cat_id,
		Title: title,
		Abstract: abstract,
		Content: content,
	}
	err, _ := create_forum.Insert()
	if err != nil {
		return types.SystemDbErr, "创建论坛失败"
	}
	return types.ReturnSuccess,  "创建论坛成功"
}

func ForumTopicLike(id int64) (int) {
	var forum Forum
	if err := orm.NewOrm().QueryTable(Forum{}).Filter("Id", id).RelatedSel().One(&forum); err != nil {
		return types.SystemDbErr
	}
	forum.Likes = forum.Likes + 1
	return types.ReturnSuccess
}