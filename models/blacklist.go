package models

import (
	"blockshop/common"
	"blockshop/types"
	"blockshop/types/collect"
	"github.com/astaxie/beego/orm"
)

type BlackList struct {
	BaseModel
	Id             int64      `orm:"pk;column(id);auto;size(11)" description:"ID" json:"id"`
	UserId         int64      `orm:"column(user_id);index" description:"用户ID" json:"user_id"`
	BlackMctId     int64      `orm:"column(black_mct_id);index" description:"黑名单商家ID" json:"black_mct_id"`
}

func (this *BlackList) TableName() string {
	return common.TableName("blacklist")
}

func (this *BlackList) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *BlackList) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *BlackList) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *BlackList) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}

func AddBlackList(req_c collect.BlackListReq) (msg string, code int) {
	ok := orm.NewOrm().QueryTable(BlackList{}).Filter("black_mct_id", req_c.MerchantId).Exist()
	if ok {
		return "该店铺已经加入黑名单", types.BlackListExist
	}
	bl := BlackList{
		UserId: req_c.UserId,
		BlackMctId: req_c.MerchantId,
	}
	err, _ := bl.Insert()
	if err != nil {
		return "该店铺加入黑名单失败", types.SystemDbErr
	}
	return "", types.ReturnSuccess
}

func lackListList(page, pageSize int) ([]*BlackList, int64) {
	offset := (page - 1) * pageSize
	list := make([]*BlackList, 0)
	query := orm.NewOrm().QueryTable(BlackList{}).Filter("is_removed", 0)
	total, _ := query.Count()
	_, err := query.OrderBy("-id").Limit(pageSize,offset).All(&list)
	if err != nil {
		return nil, types.SystemDbErr
	}
	return list, total
}

func RemoveBlackList(bl_id int64) (msg string, code int) {
	_, err := orm.NewOrm().QueryTable(BlackList{}).Filter("id", bl_id).Delete()
	if err != nil {
		return "移除商铺黑名单失败", types.BlackListExist
	}
	return "", types.ReturnSuccess
}
