package models

import (
	"blockshop/common"
	"github.com/astaxie/beego/orm"
)

type Forum struct { // 论坛表
	BaseModel
	Id             int64         `orm:"column(id)" description:"论坛ID" json:"id"`
	UserId         int64         `orm:"column(user_id)" description:"用户ID" json:"user_id"`
	Tilte          string        `orm:"column(tilte);size(128)" description:"论坛标题" json:"tilte"`
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
