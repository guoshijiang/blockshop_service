package models

import (
	"blockshop/common"
	"blockshop/types"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type HelpDesk struct {
	BaseModel
	Id        int64     `orm:"pk;column(id);auto;size(11)" description:"问题ID" json:"id"`
	UserId    int64     `orm:"column(user_id)" description:"商家对应的用户ID" json:"user_id"`
	Author    string    `orm:"column(author)" description:"反馈人姓名" json:"author"`
	IsHandle  int8      `orm:"column(is_handle);default(0);index" description:"是否处理" json:"is_handle"` // 0:未处理 1:已处理
	Contract  string    `orm:"column(contract);size(256)" description:"联系方式" json:"contract"`
	HdTitle   string    `orm:"column(hd_title);size(256)" description:"问题标题" json:"hd_title"`
	HdDetail  string    `orm:"column(hd_detail);type(text)" description:"问题详细" json:"hd_detail"`
}

func (this *HelpDesk) TableName() string {
	return common.TableName("help_desk")
}

func (this *HelpDesk) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *HelpDesk) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *HelpDesk) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *HelpDesk) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}


func GetHelpDeskList(page, pageSize int) ([]*HelpDesk, int64, error) {
	offset := (page - 1) * pageSize
	hd_list := make([]*HelpDesk, 0)
	query := orm.NewOrm().QueryTable(HelpDesk{}).OrderBy("-id")
	total, _ := query.Count()
	_, err := query.Limit(pageSize, offset).All(&hd_list)
	if err != nil {
		return nil, types.SystemDbErr, errors.New("查询数据库失败")
	}
	return hd_list, total, nil
}

func GetHelpDeskDetail(id int64) (*HelpDesk, int, error) {
	hd_detail := HelpDesk{}
	if err := orm.NewOrm().QueryTable(HelpDesk{}).Filter("Id", id).RelatedSel().One(&hd_detail); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &hd_detail, types.ReturnSuccess, nil
}

