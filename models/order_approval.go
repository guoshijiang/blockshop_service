package models

import (
	"blockshop/common"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type OrderApproval struct {
	BaseModel
	Id            int64      `json:"id"`
	OrderId       int64      `orm:"size(64);index" json:"order_id"`                       // 商品 ID
	UserId	      int64		 `orm:"size(64);index" json:"user_id"`						  // 用户ID
	MerchantId    int64      `orm:"size(64);index" json:"merchant_id"`                    // 商户 ID
	GoodsId       int64      `orm:"size(64);index" json:"goods_id"`                       // 商品 ID
	ProcessId     int64      `orm:"size(64);index" json:"process_id"`                     // 订单进度 ID

}

func (this *OrderApproval) TableName() string {
	return common.TableName("order_process")
}

func (this *OrderApproval) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *OrderApproval) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *OrderApproval) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *OrderApproval) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *OrderApproval) Insert() (error, int64) {
	id, err := orm.NewOrm().Insert(this)
	if err != nil {
		return err, 0
	}
	return nil, id
}
