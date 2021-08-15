package models


import (
	"blockshop/common"
	"blockshop/types"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	"time"
)

type OrderProcess struct {
	BaseModel
	Id            int64      `json:"id"`
	OrderId       int64      `orm:"column(order_id);size(64);index" description:"订单ID"  json:"order_id"`
	UserId	      int64		 `orm:"column(user_id);size(64);index" description:"用户ID" json:"user_id"`
	MerchantId    int64      `orm:"column(merchant_id);size(64);index" description:"商户ID" json:"merchant_id"`
	AddressId     int64      `orm:"column(address_id);size(64);index" description:"地址ID" json:"address_id"`
	GoodsId       int64      `orm:"column(goods_id);size(64);index" description:"商品ID" json:"goods_id"`
	RetGoodsRs    string     `orm:"column(ret_goods_rs);size(512)" description:"退货原因" json:"ret_goods_rs"`
	RetPayRs	  string	 `orm:"column(ret_pay_rs);size(512);default('')" description:"拒绝原因" json:"ret_pay_rs"`
	QsDescribe    string     `orm:"column(qs_describe);size(512)" description:"问题描述" json:"qs_describe"`
	VectoryId     int64      `orm:"column(vectory_id);default(0);size(64);index" description:"申诉胜出方" json:"vectory_id"`  // 商家是商家ID，用户是用户ID
	FailId        int64      `orm:"column(fail_id);default(0);size(64);index" description:"申诉失败方" json:"fail_id"`  // 商家是商家ID，用户是用户ID
	AcceptRejectRs  string   `orm:"column(reject_rs);size(512)" description:"商家拒绝原因" json:"accept_reject_rs"`
	AdjustContent string     `orm:"column(adjust_content);size(512)" description:"申述描述" json:"adjust_content"`
	QsImgOne      string     `orm:"column(qs_img_one);size(150)" description:"图片1" json:"qs_img_one"`
	QsImgTwo      string     `orm:"column(qs_img_two);size(150)" description:"图片2" json:"qs_img_two"`
	QsImgThree    string  	 `orm:"column(qs_img_three);size(150)" description:"图片3" json:"qs_img_three"`
	// 0:等待卖家确认; 1:卖家已同意; 2:卖家拒绝; 3:买家发起申诉，4:平台处理申诉; 5:等待买家邮寄; 6:等待卖家收货; 7:卖家已经发货; 8:等待买家收货; 9:已完成
	Process       int8       `orm:"column(process);default(0)" description:"订单退换货情况" json:"process"`
	IsRecvGoods   int8       `orm:"column(is_recv_goods);default(0)" description:"是否收到货物" json:"is_recv_goods"` // 0:未收到货物，1:已经收到货物
	FundRet       int8       `orm:"column(fund_ret);default(0)" description:"退换货" json:"fund_ret"` // 1.退货 2:换货
	LeftTime      int64      `orm:"column(left_time);default(604800)" description:"处理时长" json:"left_time"`
	DealTime      time.Time  `orm:"column(deal_time);auto_now_add;type(datetime)" description:"处理时间" json:"deal_time"`
}

func (this *OrderProcess) TableName() string {
	return common.TableName("order_process")
}

func (this *OrderProcess) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *OrderProcess) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *OrderProcess) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *OrderProcess) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *OrderProcess) Insert() (error, int64) {
	id, err := orm.NewOrm().Insert(this)
	if err != nil {
		return err, 0
	}
	return nil, id
}

func (this *OrderProcess) SearchField() []string {
	return []string{"order_num"}
}


func (this *OrderProcess) WaitReturnOrderTotal() int64 {
  total,_ := orm.NewOrm().QueryTable(this).Filter("process__lt",7).Count()
  return  total
}

func GetOrderProcessDetail(id int64) (*OrderProcess, int, error) {
	order_ps := OrderProcess{}
	if err := orm.NewOrm().QueryTable(OrderProcess{}).Filter("OrderId", id).RelatedSel().Limit(1).One(&order_ps); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &order_ps, types.ReturnSuccess, nil
}

func GetOrderProcessDetailById(id int64) (*OrderProcess, int, error) {
	order_ps := OrderProcess{}
	if err := orm.NewOrm().QueryTable(OrderProcess{}).Filter("id", id).RelatedSel().Limit(1).One(&order_ps); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &order_ps, types.ReturnSuccess, nil
}

func RemoveOrderProcess(order_id int64, user_id int64) error {
	_, err := orm.NewOrm().QueryTable(OrderProcess{}).Filter("order_id", order_id).Filter("user_id", user_id).Delete()
	if err != nil {
		return err
	}
	return nil
}

func OrderApproval(order_id int64, adjust_content string) (error, string) {
	order_ps := OrderProcess{}
	if err := orm.NewOrm().QueryTable(OrderProcess{}).Filter("order_id", order_id).One(&order_ps); err != nil {
		return err, "该订单没有发起过申诉"
	}
	if order_ps.Process != 2 {
		return errors.New("卖家没有拒绝该订单换货，不能发起申诉"), "卖家没有拒绝该订单换货，不能发起申诉"
	}
	if order_ps.Process == 3 {
		return errors.New("该订单已经发起过申诉"), "该订单已经发起过申诉"
	}
	order_ps.AdjustContent = adjust_content
	order_ps.Process = 3
	err := order_ps.Update()
	if err != nil {
		return err, "更新申诉描述失败"
	}
	return nil, ""
}

func OrderAcceptOrReject(order_id int64, reson string, is_accept int8) (error, string) {
	order_ps := OrderProcess{}
	if err := orm.NewOrm().QueryTable(OrderProcess{}).Filter("order_id", order_id).One(&order_ps); err != nil {
		return err, "该订单没有发起过申诉"
	}
	//0:接受；1:拒绝
	if is_accept == 0 {
		order_ps.Process = 1
	}
	if is_accept == 1 {
		if order_ps.Process == 2 {
			return errors.New("该订单已经发起拒绝"), "该订单已经发起拒绝"
		}
		order_ps.Process = 2
	}
	order_ps.AcceptRejectRs = reson
	err := order_ps.Update()
	if err != nil {
		return err, "更新申诉描述失败"
	}
	return nil, ""
}