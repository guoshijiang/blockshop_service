package models

import (
	"blockshop/common"
	"blockshop/types"
	"blockshop/types/collect"
	"github.com/astaxie/beego/orm"
)

type MerchantCollect struct {
	BaseModel
	Id             int64      `orm:"pk;column(id);auto;size(11)" description:"ID" json:"id"`
	UserId         int64      `orm:"column(user_id);index" description:"用户ID" json:"user_id"`
	CtMctId        int64      `orm:"column(ct_mct_id);index" description:"收藏的商家ID" json:"ct_mct_id"`
}

func (this *MerchantCollect) TableName() string {
	return common.TableName("merchant_collect")
}

func (this *MerchantCollect) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *MerchantCollect) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *MerchantCollect) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *MerchantCollect) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}


func AddMerchantCollect(req_c collect.MerchantCollectReq) (msg string, code int) {
	ok := orm.NewOrm().QueryTable(MerchantCollect{}).Filter("ct_mct_id", req_c.MerchantId).Filter("user_id", req_c.UserId).Exist()
	if ok {
		return "该店铺已经收藏过了", types.MctCollectExist
	}
	bl := MerchantCollect{
		UserId: req_c.UserId,
		CtMctId: req_c.MerchantId,
	}
	err, _ := bl.Insert()
	if err != nil {
		return "店铺收藏失败", types.SystemDbErr
	}
	return "", types.ReturnSuccess
}

func MerchantCollectList(page, pageSize int, user_id int64) ([]*MerchantCollect, int64) {
	offset := (page - 1) * pageSize
	list := make([]*MerchantCollect, 0)
	query := orm.NewOrm().QueryTable(MerchantCollect{}).Filter("is_removed", 0).Filter("user_id", user_id)
	total, _ := query.Count()
	_, err := query.OrderBy("-id").Limit(pageSize,offset).All(&list)
	if err != nil {
		return nil, types.SystemDbErr
	}
	return list, total
}

func RemoveMerchantCollect(bl_id int64) (msg string, code int) {
	_, err := orm.NewOrm().QueryTable(MerchantCollect{}).Filter("id", bl_id).Delete()
	if err != nil {
		return "移除收藏店铺失败", types.SystemDbErr
	}
	return "", types.ReturnSuccess
}
