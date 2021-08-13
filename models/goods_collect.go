package models


import (
	"blockshop/common"
	"blockshop/types"
	"blockshop/types/collect"
	"github.com/astaxie/beego/orm"
)

type GoodsCollect struct {
	BaseModel
	Id             int64      `orm:"pk;column(id);auto;size(11)" description:"ID" json:"id"`
	UserId         int64      `orm:"column(user_id);index" description:"用户ID" json:"user_id"`
	CtGdsId        int64      `orm:"column(ct_gds_id);index" description:"收藏的商品ID" json:"ct_gds_id"`
}

func (this *GoodsCollect) TableName() string {
	return common.TableName("goods_collect")
}

func (this *GoodsCollect) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *GoodsCollect) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *GoodsCollect) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *GoodsCollect) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}

func AddGoodsCollect(req_c collect.GoodsCollectReq) (msg string, code int) {
	ok := orm.NewOrm().QueryTable(GoodsCollect{}).Filter("ct_gds_id", req_c.GoodsId).Exist()
	if ok {
		return "该商品已经收藏过了", types.GoodsCollectExist
	}
	bl := GoodsCollect{
		UserId: req_c.UserId,
		CtGdsId: req_c.GoodsId,
	}
	err, _ := bl.Insert()
	if err != nil {
		return "该店铺加入黑名单失败", types.SystemDbErr
	}
	return "", types.ReturnSuccess
}

func GoodsCollectList(page, pageSize int) ([]*GoodsCollect, int64) {
	offset := (page - 1) * pageSize
	list := make([]*GoodsCollect, 0)
	query := orm.NewOrm().QueryTable(GoodsCollect{}).Filter("is_removed", 0)
	total, _ := query.Count()
	_, err := query.OrderBy("-id").Limit(pageSize,offset).All(&list)
	if err != nil {
		return nil, types.SystemDbErr
	}
	return list, total
}

func RemoveGoodsCollect(bl_id int64) (msg string, code int) {
	_, err := orm.NewOrm().QueryTable(GoodsCollect{}).Filter("id", bl_id).Delete()
	if err != nil {
		return "移除收藏商品失败", types.BlackListExist
	}
	return "", types.ReturnSuccess
}