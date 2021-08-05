package models

import (
	"blockshop/common"
	"blockshop/types"
	"blockshop/types/order"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	"time"
)

type GoodsOrder struct {
	BaseModel
	Id            int64      `orm:"pk;column(id);auto;size(11)" description:"订单ID" json:"id"`
	GoodsId       int64      `orm:"column(goods_id);size(64);index" description:"商品ID" json:"goods_id"`
	MerchantId    int64      `orm:"column(merchant_id);size(64);index" description:"商户ID" json:"merchant_id"`
	AddressId     int64      `orm:"column(address_id);size(64);index" description:"地址ID" json:"address_id"`
	GoodsTypes    string     `orm:"column(goods_types);size(512)" description:"商品属性" json:"goods_types"`
	GoodsTitle    string     `orm:"column(goods_title);size(64)" description:"商品标题" json:"goods_title"`
	GoodsName     string     `orm:"column(goods_name);size(512);index" description:"产品名称" json:"goods_name"`
	Logo          string     `orm:"column(logo);size(150);default(/static/upload/default/user-default-60x60.png)" description:"商品Logo" json:"logo"`
	UserId        int64      `orm:"column(user_id);size(64);index" description:"购买用户" json:"user_id"`
	BuyNums       int64      `orm:"column(buy_nums);default(0)" description:"购买数量" json:"buy_nums"`
	PayWay        int8       `orm:"column(pay_way);index" description:"支付方式" json:"pay_way"`  // 0:BTC，1:USDT
	PayAmount     float64    `orm:"column(pay_amount);default(0);digits(22);decimals(8)" description:"支付金额" json:"pay_amount"`
	OrderNumber   string     `orm:"column(order_number);size(64);index" description:"订单号" json:"order_number"`
	Logistics	  string     `orm:"column(logistics);size(64);index;default('')" description:"物流公司" json:"logistics"`
	ShipNumber    string     `orm:"column(ship_number);size(64);index;default('')" description:"运单号" json:"ship_number"`
	OrderStatus   int8       `orm:"column(order_status);index" description:"支付状态" json:"order_status"` // 0: 未支付，1: 支付中，2：支付成功；3：支付失败 4：已发货；5：已完成
	FailureReason string     `orm:"column(failure_reason)" description:"失败原因" json:"failure_reason"`
	PayAt         *time.Time `orm:"column(pay_at);type(datetime);null" description:"支付时间" json:"pay_at"`
	DealMerchant  string     `orm:"column(deal_user);default('')" description:"处理商家" json:"deal_user"`
	DealAt        time.Time  `orm:"column(deal_at);null;type(datetime);" description:"处理时间" json:"deal_at"`
	IsCancle      int8       `orm:"column(is_cancle);default(0);index" description:"退换货状态" json:"is_cancle"`       // 0 正常；1.退货; 2:换货; 3:退货成功; 4:换货成功
	IsComment     int8       `orm:"column(is_comment);default(0);index" description:"是否平价" json:"is_comment"`        // 0 正常；1.已评价
	IsStatic      int8 		 `orm:"column(is_static);default(0);index" description:"是否统计交易流水" json:"is_static"`  // 0 正常；1.已统计
}

func (this *GoodsOrder) TableName() string {
	return common.TableName("goods_order")
}

func (this *GoodsOrder) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *GoodsOrder) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *GoodsOrder) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *GoodsOrder) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *GoodsOrder) Insert() (error, int64) {
	id, err := orm.NewOrm().Insert(this)
	if err != nil {
		return err, 0
	}
	return nil, id
}

func (this *GoodsOrder) SearchField() []string {
	return []string{"order_num"}
}


func GetGoodsOrderList(page, pageSize int, user_id int64, status int8) ([]*GoodsOrder, int64, error) {
	offset := (page - 1) * pageSize
	gds_order_list := make([]*GoodsOrder, 0)
	query := orm.NewOrm().QueryTable(GoodsOrder{}).Filter("UserId", user_id).OrderBy("-id")
	if status >= 0  && status <= 5 {
		query = query.Filter("OrderStatus", status)
	}
	total, _ := query.Count()
	_, err := query.Limit(pageSize, offset).All(&gds_order_list)
	if err != nil {
		return nil, types.SystemDbErr, errors.New("查询数据库失败")
	}
	return gds_order_list, total, nil
}


func GetGoodsOrderDetail(id int64) (*GoodsOrder, int, error) {
	var order_dtl GoodsOrder
	if err := orm.NewOrm().QueryTable(GoodsOrder{}).Filter("Id", id).RelatedSel().One(&order_dtl); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &order_dtl, types.ReturnSuccess, nil
}


// 1.退货,资金返回钱包账号; 2:退货,资金原路返回; 3:换货
func ReturnGoodsOrder(oret order.ReturnGoodsOrderReq) (*GoodsOrder, int, error) {
	var order_dtl GoodsOrder
	if err := orm.NewOrm().QueryTable(GoodsOrder{}).Filter("Id", oret.OrderId).RelatedSel().One(&order_dtl); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	if oret.FundRet == 1 || oret.FundRet == 2 {
		order_dtl.IsCancle = 1
	}
	if oret.FundRet == 3 {
		order_dtl.IsCancle = 2
	}
	err := order_dtl.Update()
	if err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	order_p := OrderProcess{
		OrderId: order_dtl.Id,
		MerchantId: order_dtl.MerchantId,
		UserId: order_dtl.UserId,
		AddressId: order_dtl.AddressId,
		GoodsId: order_dtl.GoodsId,
		RetGoodsRs: oret.RetGoodsRs,
		QsDescribe: oret.QsDescribe,
		QsImgOne: oret.QsImgOne,
		QsImgTwo: oret.QsImgTwo,
		QsImgThree: oret.QsImgThree,
		Process: 0,
		LeftTime: 604800,
		IsRecvGoods: oret.IsRecvGoods,    // 0:未收到货物，1:已经收到货物
		FundRet: oret.FundRet,
	}
	err, _ = order_p.Insert()
	if err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &order_dtl, types.ReturnSuccess, nil
}

