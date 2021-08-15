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


const (
	OrderStatusNoPay       = 0  // 未支付
	OrderStatusPaySuccess  = 1  // 支付成功
	OrderStatusPayFailure  = 2  // 支付失败
	OrderStatusSendGoods   = 3  // 已发货
	OrderStatusRecvGoods   = 4  // 已经收货
	OrederReturnGoods      = 5  // 退货换货
	OrederReturnSellerRjt  = 6  // 卖家拒绝
	OrederReturnSellerAcpt = 7  // 卖家同意
	OrederBuyerAppoval     = 8  // 买家申诉
	OrederAppovalSuccess   = 9  // 订单申诉成功
	OrederSellerReturnMny  = 10 // 卖家已退款
	OrederFinish           = 11 // 已完成
	PayWayUSDT             = 1
	PayWayBTC              = 2
)

type GoodsOrder struct {
	BaseModel
	Id            int64      `orm:"pk;column(id);auto;size(11)" description:"订单ID" json:"id"`
	GoodsId       int64      `orm:"column(goods_id);size(64);index" description:"商品ID" json:"goods_id"`
	MerchantId    int64      `orm:"column(merchant_id);size(64);index" description:"商户ID" json:"merchant_id"`
	AddressId     int64      `orm:"column(address_id);size(64);index" description:"地址ID" json:"address_id"`
	GoodsAttrs    string     `orm:"column(goods_types);size(512)" description:"商品属性" json:"goods_attrs"`
	GoodsTitle    string     `orm:"column(goods_title);size(64)" description:"商品标题" json:"goods_title"`
	GoodsName     string     `orm:"column(goods_name);size(512);index" description:"产品名称" json:"goods_name"`
	Logo          string     `orm:"column(logo);size(150);default(/static/upload/default/user-default-60x60.png)" description:"商品Logo" json:"logo"`
	UserId        int64      `orm:"column(user_id);size(64);index" description:"购买用户" json:"user_id"`
	BuyNums       int64      `orm:"column(buy_nums);default(0)" description:"购买数量" json:"buy_nums"`
	PayWay        int8       `orm:"column(pay_way);index" description:"支付方式" json:"pay_way"`  // 0:BTC，1:USDT
	PayCnyPrice   float64 	 `orm:"column(pay_cny_price);default(0);digits(22);decimals(8)" description:"支付金额" json:"pay_cny_price"`
	PayCoinAmount float64    `orm:"column(pay_coin_amount);default(0);digits(22);decimals(8)" description:"支付币的数量" json:"pay_coin_amount"`
	OrderNumber   string     `orm:"column(order_number);size(64);index" description:"订单号" json:"order_number"`
	Logistics	  string     `orm:"column(logistics);size(64);index;default('')" description:"物流公司" json:"logistics"`
	ShipNumber    string     `orm:"column(ship_number);size(64);index;default('')" description:"运单号" json:"ship_number"`
	RetShipNumber string     `orm:"column(ret_ship_number);size(64);index;default('')" description:"退货运单号" json:"ret_ship_number"`
	OrderStatus   int8       `orm:"column(order_status);index" description:"订单状态" json:"order_status"`
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

// 0: 未支付，1: 支付中，2：支付成功；3：支付失败 4：已发货；5：已完成
func (this *GoodsOrder) Aggregation(merchant int64) (int64,int64,int64) {
  _,err := orm.NewOrm().Raw("SET sql_mode = 'STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION'").Exec()
  if err != nil {
    println("err ----",err)
  }
  var data []order.StateStatic
  _,err = orm.NewOrm().Raw("select count(order_status) total,order_status state from goods_order where merchant_id = ?",merchant).QueryRows(&data)
  if err != nil {
    return 0,0,0
  }
  var (
    WaidPayOrderNum int64
    WaitSendOrderNum int64
    SendOrderNum int64
  )
  for _,item := range data {
    if item.State == 0 {
      WaidPayOrderNum += 1
    }
    if item.State == 2 {
      WaitSendOrderNum += 1
    }
    if item.State == 4 {
      SendOrderNum += 1
    }
  }
  return WaidPayOrderNum,WaitSendOrderNum,SendOrderNum
}

func PayOrder(order_id int64) (success bool, err error, code int, msg string) {
	db := orm.NewOrm()
	if err := db.Begin(); err != nil {
		err := errors.Wrap(err, "开启支付事物失败")
		return false, err, types.OrderPayException, "开启支付事物失败"
	}
	defer func() {
		if err != nil {
			err = errors.Wrap(db.Rollback(), "回滚事物失败")
		} else {
			err = errors.Wrap(db.Commit(), "提交事物失败")
		}
	}()
	goods_order := GoodsOrder{}
	if err = db.QueryTable(goods_order.TableName()).Filter("id", order_id).One(&goods_order); err != nil {
		err := errors.New("查询订单失败")
		return false, err, types.OrderPayException, "查询订单失败"
	}
	if goods_order.OrderStatus == OrderStatusPaySuccess {
		err := errors.New("订单已经支付")
		return false, err, types.OrderAlreadyPay, "订单已经支付"
	}
	if goods_order.OrderStatus != 0 {
		goods_order.OrderStatus = 0
		goods_order.FailureReason = ""
	}
	goods := Goods{}
	if err = db.QueryTable(goods.TableName()).Filter("id", goods_order.GoodsId).One(&goods); err != nil {
		err = errors.New("查询商品信息失败")
		return false, err, types.OrderPayException, "查询商品信息失败"
	}
	user := User{}
	if err = db.QueryTable(user.TableName()).RelatedSel().Filter("id", goods_order.UserId).One(&user); err != nil {
		err = errors.New("查询购买商品的用户失败")
		return false, err, types.OrderPayException, "查询购买商品的用户失败"
	}
	var pay_asset *Asset
	if goods_order.PayWay == PayWayUSDT {
		supportedPayPrice := float64(goods_order.BuyNums) * goods.GoodsPrice
		if supportedPayPrice != float64(goods_order.PayCnyPrice) {
			err = errors.New("支付方式错误")
			return false, err, types.VerifyPayAmount, "支付方式错误"
		}
		var ast Asset
		ast.Name = "USDT"
		asst, _ := ast.GetAsset()
		pay_asset = asst
		total, code, err := GetUserWalletBalance(asst.Id, goods_order.UserId)
		if err != nil {
			return false, err, code, "获取用户钱包失败"
		}
		if total < supportedPayPrice {
			err = errors.New("账户没有足够的资金，请去充值")
			return false, err, types.AccountAmountNotEnough, "账户没有足够的资金，请去充值"
		}
	} else {
		supportedPayPrice := float64(goods_order.BuyNums) * goods.GoodsPrice
		if supportedPayPrice != float64(goods_order.PayCnyPrice) {
			err = errors.New("支付方式错误")
			return false, err, types.VerifyPayAmount, "支付方式错误"
		}
		var ast Asset
		ast.Name = "BTC"
		asst, _ := ast.GetAsset()
		pay_asset = asst
		total, code, err := GetUserWalletBalance(asst.Id, goods_order.UserId)
		if err != nil {
			return false, err, code, "获取用户钱包失败"
		}
		if total < goods_order.PayCoinAmount {
			err = errors.New("账户没有足够的资金，请去充值")
			return false, err, types.AccountAmountNotEnough, "账户没有足够的资金，请去充值"
		}
	}
	if len(goods_order.FailureReason) > 0 {
		goods_order.OrderStatus = OrderStatusPayFailure
		if _, err = db.Update(&goods_order, "OrderStatus", "FailureReason"); err != nil {
			err = errors.New("更新订单的状态失败")
			return false, err, types.OrderPayException, "更新订单的状态失败"
		}
	}
	success, code, err = UpdateWalletBalance(db, pay_asset.Id, goods_order.UserId, float64(goods_order.PayCoinAmount))
	if err != nil {
		return success, err, code, "更新用户钱包余额失败"
	}
	goods.LeftAmount -= goods_order.BuyNums
	if _, err := db.Update(&Goods{}, "LeftAmount"); err != nil {
		err = errors.New("更新剩余商品个数失败")
		return false, err, types.OrderPayException, "更新剩余商品个数失败"
	}
	now := time.Now()
	goods_order.OrderStatus = OrderStatusPaySuccess
	goods_order.FailureReason = ""
	goods_order.PayAt = &now
	if _, err = db.Update(&goods_order, "OrderStatus", "FailureReason", "PayAt"); err != nil {
		err = errors.New("更新订单状态失败")
		return false, err, types.OrderPayException, "更新订单状态失败"
	}
	mct_order_flow := MerchantOrderFlow{
		MerchantId: goods_order.MerchantId,
		OrderId: goods_order.Id,
		AssetId: pay_asset.Id,
		CoinAmount: float64(goods_order.PayCoinAmount),
	}
	err, _ = mct_order_flow.Insert()
	if err != nil {
		err = errors.New("更新流水失败")
		return false, err, types.OrderPayException, "更新流水失败"
	}
	return true, nil, types.ReturnSuccess, "支付成功"
}


func GetGoodsOrderList(page, pageSize int, user_id, merchant_id int64, status int8) ([]*GoodsOrder, int64, error) {
	offset := (page - 1) * pageSize
	gds_order_list := make([]*GoodsOrder, 0)
	query := orm.NewOrm().QueryTable(GoodsOrder{}).Filter("is_removed", 0).OrderBy("-id")
	if user_id >= 1 {
		query = query.Filter("user_id", user_id)
	}
	if merchant_id >= 1 {
		query = query.Filter("merchant_id", merchant_id)
	}
	if status >= 0  && status <= 11 {
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


func ReturnGoodsOrder(oret order.ReturnGoodsOrderReq) (*GoodsOrder, int, error) {
	var order_dtl GoodsOrder
	if err := orm.NewOrm().QueryTable(GoodsOrder{}).Filter("Id", oret.OrderId).RelatedSel().One(&order_dtl); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	order_dtl.OrderStatus = OrederReturnGoods
	if oret.FundRet == 1 {
		order_dtl.IsCancle = 1
	}
	if oret.FundRet == 2  {
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
		Process: ProcessWaitSellerConfirm,
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

func UpdReturnShipNumber(order_id int64, ship_number string) error {
	var order_dtl GoodsOrder
	if err := orm.NewOrm().QueryTable(GoodsOrder{}).Filter("id", order_id).RelatedSel().One(&order_dtl); err != nil {
		return errors.New("数据库查询失败，请联系客服处理")
	}
	order_dtl.RetShipNumber = ship_number
	err := order_dtl.Update()
	if err != nil {
		return errors.New("数据库查询失败，请联系客服处理")
	}
	var order_pcs OrderProcess
	err = orm.NewOrm().QueryTable(OrderProcess{}).Filter("order_id", order_id).RelatedSel().One(&order_pcs)
	if err == nil {
		order_pcs.Process = ProcesBuyerPost
		err := order_pcs.Update()
		if err != nil {
			return errors.New("数据库查询失败，请联系客服处理")
		}
	}
	return nil
}

func UpdShipNumber(order_id int64, ship_number string) error {
	var order_dtl GoodsOrder
	if err := orm.NewOrm().QueryTable(GoodsOrder{}).Filter("id", order_id).RelatedSel().One(&order_dtl); err != nil {
		return errors.New("数据库查询失败，请联系客服处理")
	}
	order_dtl.ShipNumber = ship_number
	order_dtl.OrderStatus = OrderStatusSendGoods
	err := order_dtl.Update()
	if err != nil {
		return errors.New("数据库查询失败，请联系客服处理")
	}
	var order_pcs OrderProcess
	err = orm.NewOrm().QueryTable(OrderProcess{}).Filter("order_id", order_id).RelatedSel().One(&order_pcs)
	if err == nil {
		order_pcs.Process = ProcesSellerPost
		err := order_pcs.Update()
		if err != nil {
			return errors.New("数据库查询失败，请联系客服处理")
		}
	}
	return nil
}