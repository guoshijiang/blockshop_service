package models

import (
	"blockshop/common"
)

type AdminRole struct {
	Id          int    `orm:"column(id);auto;size(11)" description:"表ID"`
	Name        string `orm:"column(name);size(50)" description:"名称"`
	Description string `orm:"column(description);size(100)" description:"简介"`
	Url         string `orm:"column(url);size(1000)" description:"权限"`
	Status      int8   `orm:"column(status);size(1);default(1)" description:"是否启用 0：否 1：是"`
}

func (*AdminRole) TableName() string {
	return common.TableName("admin_role")
}

//定义模型的可搜索字段
func (*AdminRole) SearchField() []string {
	return []string{"name", "description"}
}

//禁止删除的数据id
func (*AdminRole) NoDeletionId() []int {
	return []int{1}
}

//定义模型可作为条件的字段
func (*AdminRole) WhereField() []string {
	return []string{}
}

//定义可做为时间范围查询的字段
func (*AdminRole) TimeField() []string {
	return []string{}
}
