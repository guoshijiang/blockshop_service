package order


import (
	"blockshop/types"
	"github.com/pkg/errors"
)

type CreateOrderReq struct {
	GoodsId       int64   `json:"goods_id"`
	AddressId     int64   `json:"address_id"`
	UserId        int64   `json:"user_id"`
	BuyNums       int64   `json:"buy_nums"`
	PayWay        int8    `json:"pay_way"`         // 0:BTC支付，1:USDT支付
	PayCnyPrice   float64 `json:"pay_cny_price"`      // 支付金额
	PayCoinAmount float64 `json:"pay_coin_amount"`
	GoodsAttr     string  `json:"goods_attr"`
}

func (this CreateOrderReq) ParamCheck() (int, error) {
	if this.GoodsId <= 0 {
		return types.ParamLessZero, errors.New("商品ID小于等于 0")
	}
	if this.AddressId <= 0 {
		return types.ParamLessZero, errors.New("您没有选择地址，或者您还没有添加地址，请去选择或者添加")
	}
	if this.UserId <= 0 {
		return types.ParamLessZero, errors.New("用户ID小于等于 0")
	}
	if this.BuyNums <= 0 {
		return types.ParamLessZero, errors.New("购买数量小于等于 0")
	}
	if this.PayWay < 0 || this.PayWay > 3 {
		return types.InvalidVerifyWay, errors.New("无效的付款方式")
	}
	return types.ReturnSuccess, nil
}

type OrderListReq struct {
	types.PageSizeData
	UserId      int64 `json:"user_id"`
	MerchantId  int64 `json:"merchant_id"`
	OrderStatus int8 `json:"order_status"`  // 0: 未支付，1: 支付中，2：支付成功；3：支付失败 4：已发货；5：已经收货; 6:待退款，7：全部
}

func (this OrderListReq) ParamCheck() (int, error) {
	code, err := this.SizeParamCheck()
	if err != nil {
		return code, err
	}
	if this.UserId <= 0 {
		return types.ParamLessZero, errors.New("用户ID小于等于 0")
	}
	return types.ReturnSuccess, nil
}

type OrderDetailReq struct {
	OrderId int64  `json:"order_id"`
	IsCancle int8  `json:"is_cancle"` //0:正常； 1.退换货

}

func (this OrderDetailReq) ParamCheck() (int, error) {
	if this.OrderId <= 0 {
		return types.ParamLessZero, errors.New("订单 ID 小于等于 0")
	}
	return types.ReturnSuccess, nil
}

type ReturnGoodsOrderReq struct {
	OrderId       int64  `json:"order_id"`
	RetGoodsRs    string `json:"ret_goods_rs"`    // 退货原因
	QsDescribe    string `json:"qs_describe"`     // 问题描述
	QsImgOne      string `json:"qs_img_one"`
	QsImgTwo      string `json:"qs_img_two"`
	QsImgThree    string `json:"qs_img_three"`
	IsRecvGoods   int8   `json:"is_recv_goods"`   // 0:未收到货物，1:已经收到货物
	FundRet       int8   `json:"fund_ret"`        // 1.退货 2:换货
}

func (this ReturnGoodsOrderReq) ParamCheck() (int, error) {
	if this.OrderId <= 0 {
		return types.ParamLessZero, errors.New("订单 ID 小于等于 0")
	}
	return types.ReturnSuccess, nil
}

type CancleReturnGoodsOrderReq struct {
	OrderId  int64 `json:"order_id"`
}

func (this CancleReturnGoodsOrderReq) ParamCheck() (int, error) {
	if this.OrderId <= 0 {
		return types.ParamLessZero, errors.New("订单 ID 小于等于 0")
	}
	return types.ReturnSuccess, nil
}


type PayOrderReq struct {
	OrderId         int64   `json:"order_id"`
	PayCoinAmount   float64 `json:"pay_coin_amount"`    // 付款金额或者付款积分
	PayWay          int8    `json:"pay_way"`            // 0:BTC支付，1:USDT支付
}

func (this PayOrderReq) ParamCheck() (int, error) {
	if this.OrderId <= 0 {
		return types.ParamLessZero, errors.New("订单的 ID 不能小于等于 0")
	}
	if this.PayWay <0  || this.PayWay > 3 {
		return types.InvalidVerifyWay, errors.New("无效的支付方式")
	}
	return types.ReturnSuccess, nil
}

type ReturnShipNumber struct {
	OrderId         int64   `json:"order_id"`
	RetShipNumber   string  `json:"ret_ship_number"`
}

func (this ReturnShipNumber) ParamCheck() (int, error) {
	if this.OrderId <= 0 {
		return types.ParamLessZero, errors.New("订单的 ID 不能小于等于 0")
	}
	if this.RetShipNumber == "" {
		return types.InvalidFormatError, errors.New("快递号不能为空")
	}
	return types.ReturnSuccess, nil
}

type ReturnGoodsApprovalReq struct {
	OrderId         int64   `json:"order_id"`
	AdjustContent   string  `json:"adjust_content"`
}


func (this ReturnGoodsApprovalReq) ParamCheck() (int, error) {
	if this.OrderId <= 0 {
		return types.ParamLessZero, errors.New("订单的 ID 不能小于等于 0")
	}
	if this.AdjustContent == "" {
		return types.InvalidFormatError, errors.New("申诉描述不能为空")
	}
	return types.ReturnSuccess, nil
}