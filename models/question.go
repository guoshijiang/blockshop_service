package models

import (
	"blockshop/common"
	"blockshop/types"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type Questions struct {
	BaseModel
	Id        int64     `orm:"pk;column(id);auto;size(11)" description:"问题ID" json:"id"`
	Author    string    `orm:"column(author)" description:"问题ID" json:"author"`
	QsTitle   string    `orm:"column(qs_title);size(256)" description:"问题标题" json:"qs_title"`
	QsDetail  string    `orm:"column(qs_detail);type(text)" description:"常见问题详细" json:"qs_detail"`
}

func (this *Questions) TableName() string {
	return common.TableName("questions")
}

func (this *Questions) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *Questions) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *Questions) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *Questions) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}


func GetQuestionsList(page, pageSize int) ([]*Questions, int64, error) {
	offset := (page - 1) * pageSize
	qs_list := make([]*Questions, 0)
	query := orm.NewOrm().QueryTable(Questions{}).OrderBy("-id")
	total, _ := query.Count()
	_, err := query.Limit(pageSize, offset).All(&qs_list)
	if err != nil {
		return nil, types.SystemDbErr, errors.New("查询数据库失败")
	}
	return qs_list, total, nil
}

func GetQuestionsDetail(id int64) (*Questions, int, error) {
	qs_detail := Questions{}
	if err := orm.NewOrm().QueryTable(Questions{}).Filter("Id", id).RelatedSel().One(&qs_detail); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &qs_detail, types.ReturnSuccess, nil
}
