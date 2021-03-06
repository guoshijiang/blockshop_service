package api

import (
	"blockshop/models"
	"blockshop/types"
	"blockshop/types/order"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	uuid "github.com/satori/go.uuid"
	"strconv"
	"strings"
)

type OrderController struct {
	beego.Controller
}

// CreateOrder @Title CreateOrder finished
// @Description 直接创建订单 CreateOrder
// @Success 200 status bool, data interface{}, msg string
// @router /create_order [post]
func (this *OrderController) CreateOrder() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	requestUser, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	var create_order order.CreateOrderReq
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &create_order); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := create_order.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	if requestUser.Id != create_order.UserId {
		this.Data["json"] = RetResource(false, types.UserIsNotExist, err, "Token 和用户不匹配，拒绝添加地址")
		this.ServeJSON()
		return
	}
	gds, _, _ := models.GetGoodsDetail(create_order.GoodsId)
	order_nmb := uuid.NewV4()
	if gds.GoodsPrice * float64(create_order.BuyNums) != create_order.PayCnyPrice {
		this.Data["json"] = RetResource(false, types.InvalidGoodsPirce, err, "无效的商品价格")
		this.ServeJSON()
		return
	}
	var btc_usdt_price float64
	if create_order.PayWay == 0 {
		btc_price := models.GetAssetByName("BTC")
		if btc_price != nil {
			btc_price_b_, _ := strconv.ParseFloat(btc_price.CnyPrice, 64)
			btc_usdt_price = btc_price_b_
		} else {
			this.Data["json"] = RetResource(false, types.GetBtcRateFail, err.Error(), "获取 BTC 费率失败")
			this.ServeJSON()
			return
		}
	}
	if create_order.PayWay == 1{
		usdt_price := models.GetAssetByName("USDT")
		if usdt_price != nil {
			usdt_price_u_, _ := strconv.ParseFloat(usdt_price.CnyPrice, 64)
			btc_usdt_price = usdt_price_u_
		} else {
			this.Data["json"] = RetResource(false, types.GetUsdtRateFail, err.Error(), "获取 BTC 费率失败")
			this.ServeJSON()
			return
		}
	}
	pay_coin_amount := create_order.PayCnyPrice / btc_usdt_price
	cmt := models.GoodsOrder{
		GoodsId: gds.Id,
		MerchantId: gds.MerchantId,
		AddressId: create_order.AddressId,
		GoodsTitle: gds.Title,
		GoodsName: gds.GoodsName,
		GoodsAttrs: create_order.GoodsAttr,
		Logo: gds.Logo,
		UserId: create_order.UserId,
		BuyNums: create_order.BuyNums,
		PayWay: create_order.PayWay,
		PayCnyPrice: create_order.PayCnyPrice,
		PayCoinAmount: pay_coin_amount,
		OrderNumber: order_nmb.String(),
		OrderStatus: models.OrderStatusNoPay,
		FailureReason: "未支付",
	}
	err, id := cmt.Insert()
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, err.Error(), "创建订单失败")
		this.ServeJSON()
		return
	} else {
		data := map[string]interface{}{
			"id": id,
			"pay_cny_amount": create_order.PayCnyPrice,
			"pay_coin_amount": pay_coin_amount,
			"pay_way": create_order.PayWay,
		}
		this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "创建订单成功")
		this.ServeJSON()
		return
	}
}

// PayOrder @Title PayOrder finished
// @Description 单个订单支付 PayOrder
// @Success 200 status bool, data interface{}, msg string
// @router /pay_order [post]
func (this *OrderController) PayOrder () {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	_, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	pay_order := order.PayOrderReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &pay_order); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := pay_order.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	ordr, code, err := models.GetGoodsOrderDetail(pay_order.OrderId)
	if err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	if ordr.PayCoinAmount != pay_order.PayCoinAmount {
		this.Data["json"] = RetResource(false, types.VerifyPayAmount, nil, "支付金额不对")
		this.ServeJSON()
		return
	}
	if ordr.PayWay != pay_order.PayWay {
		this.Data["json"] = RetResource(false, types.PayOrderError, nil, "支付方式不对")
		this.ServeJSON()
		return
	}
	ok, err, code, msg := models.PayOrder(pay_order.OrderId)
	fmt.Println(err)
	if err == nil && ok == true {
		this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "支付成功")
		this.ServeJSON()
		return
	} else {
		if err != nil {
			this.Data["json"] = RetResource(false, code, nil, msg)
			this.ServeJSON()
			return
		}
		this.Data["json"] = RetResource(false, code, nil, msg)
		this.ServeJSON()
		return
	}
}


// OrderList @Title OrderList finished
// @Description 订单列表 OrderList
// @Success 200 status bool, data interface{}, msg string
// @router /order_list [post]
func (this *OrderController) OrderList() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	u_tk, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	var order_lst order.OrderListReq
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &order_lst); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := order_lst.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	ols, total, err := models.GetGoodsOrderList(order_lst.Page, order_lst.PageSize, u_tk.Id, order_lst.MerchantId, order_lst.OrderStatus)
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, err, err.Error())
		this.ServeJSON()
		return
	}
	var olst_ret []order.OrderListRet
	//img_path := beego.AppConfig.String("img_root_path")
	for _, value := range ols {
		m, _, _ := models.GetMerchantDetail(value.MerchantId)
		gds, _, _ := models.GetGoodsDetail(value.GoodsId)
		var goods_last_price float64
		if gds.IsDiscount == 0 {
			goods_last_price = gds.GoodsPrice
		}  else {
			goods_last_price = gds.GoodsDisPrice
		}
		var addr models.UserAddress
		addr.Id = value.AddressId
		rec_user := ""
		rec_pbone := ""
		rec_address := ""
		addrs, _, _ := addr.GetAddressById()
		if addrs != nil {
			rec_user = addrs.UserName
			rec_pbone = addrs.Phone
			rec_address = addrs.Address
		}
		ordr := order.OrderListRet {
			MerchantId: m.Id,
			MerchantName: m.MerchantName,
			MerchantPhone: m.Phone,
			OrderId:value.Id,
			GoodsName: value.GoodsName,
			GoodsTitel: value.GoodsTitle,
			GoodsLogo: value.Logo,
			GoodsPrice: goods_last_price,
			RecUser: rec_user,
			RecPhone: rec_pbone,
			RecAddress: rec_address,
			OrderStatus: value.OrderStatus,
			BuyNums: value.BuyNums,
			PayCnyPrice: value.PayCnyPrice,
			PayCoinAmount: value.PayCoinAmount,
			PayWay: value.PayWay,
			IsCancle: value.IsCancle,
			IsComment: value.IsComment,
			IsDiscount: gds.IsDiscount,
			IsAdmin: gds.IsAdmin,
			CreateTime: gds.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		olst_ret = append(olst_ret, ordr)
	}
	data := map[string]interface{}{
		"total":     total,
		"order_lst": olst_ret,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取订单列表成功")
	this.ServeJSON()
	return
}


// OrderDetail @Title OrderDetail finished
// @Description 订单详情 OrderDetail
// @Success 200 status bool, data interface{}, msg string
// @router /order_detail [post]
func (this *OrderController) OrderDetail() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	_, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	var order_dtl order.OrderDetailReq
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &order_dtl); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := order_dtl.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	ord_dtl, code, err := models.GetGoodsOrderDetail(order_dtl.OrderId)
	if err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	//img_path := beego.AppConfig.String("img_root_path")
	var addr models.UserAddress
	addr.Id = ord_dtl.AddressId
	addrs, _, _ := addr.GetAddressById()
	mct, _, _ := models.GetMerchantDetail(ord_dtl.MerchantId)
	gdsdtl, _, _ := models.GetGoodsDetail(ord_dtl.GoodsId)
	var goods_last_price float64
	if gdsdtl.IsDiscount == 0 {
		goods_last_price = gdsdtl.GoodsPrice
	}  else {
		goods_last_price = gdsdtl.GoodsDisPrice
	}
	var ret_ordr *order.ReturnOrderProcess
	if ord_dtl.IsCancle != 0 {
		order_process, _, err := models.GetOrderProcessDetail(ord_dtl.Id)
		if err == nil && order_process != nil{
			ret_ordr = &order.ReturnOrderProcess{
				ProcessId: order_process.Id,
				ReturnUser: mct.ContactUser,
				ReturnPhone: mct.Phone,
				ReturnAddress: mct.Address,
				ReturnReson: order_process.RetGoodsRs,
				ReturnAmount: ord_dtl.PayCoinAmount,
				AcceptRejectRs: order_process.RetPayRs,
				QuestionDes: order_process.QsDescribe,
				AdjustContent: order_process.AdjustContent,
				QsImgOne: order_process.QsImgOne,
				QsImgTwo: order_process.QsImgTwo,
				QsImgThree: order_process.QsImgThree,
				AskTime: order_process.CreatedAt,
				Process: order_process.Process,
				IsRecvGoods: order_process.IsRecvGoods,
				LeftTime: order_process.LeftTime,
			}
		} else {
			ret_ordr = nil
		}
	} else {
		ret_ordr = nil
	}
	odl := order.OrderDetailRet{
		OrderId: ord_dtl.Id,
		GoodsId: ord_dtl.GoodsId,
		Logistics: ord_dtl.Logistics,
		ShipNumber: ord_dtl.ShipNumber,
		RetShipNumber: ord_dtl.RetShipNumber,
		RecUser: addrs.UserName,
		RecPhone: addrs.Phone,
		RecAddress:addrs.Address,
		MerchantId: mct.Id,
		MerchantName: mct.MerchantName,
		GoodsName: gdsdtl.GoodsName,
		GoodsLogo:  gdsdtl.Logo,
		GoodsPrice: goods_last_price,
		OrderStatus: ord_dtl.OrderStatus,
		BuyNums: ord_dtl.BuyNums,
		PayCnyPrice: ord_dtl.PayCnyPrice,
		PayCoinAmount: ord_dtl.PayCoinAmount,
		ShipFee: 0,
		PayWay: ord_dtl.PayWay,
		OrderNumber: ord_dtl.OrderNumber,
		PayTime: ord_dtl.PayAt,
		CreateTime: ord_dtl.CreatedAt,
		IsCancle: ord_dtl.IsCancle,
		IsComment: ord_dtl.IsComment,
		IsDiscount: gdsdtl.IsDiscount,
		GoodsTypes: ord_dtl.GoodsAttrs,
		RetrurnOrder: ret_ordr,
		IsAdmin: gdsdtl.IsAdmin,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, odl, "获取订单详情成功")
	this.ServeJSON()
	return
}

// ReturnGoodsOrder @Title ReturnGoodsOrder finished
// @Description 订单换退货 ReturnGoodsOrder
// @Success 200 status bool, data interface{}, msg string
// @router /return_goods_order [post]
func (this *OrderController) ReturnGoodsOrder() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	_, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	var order_ret order.ReturnGoodsOrderReq
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &order_ret); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := order_ret.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	odr, _, _ := models.GetGoodsOrderDetail(order_ret.OrderId)
	if odr.IsCancle != 0 {
		this.Data["json"] = RetResource(false, types.AlreadyCancleOrder, nil, "该订单已经发起退换货")
		this.ServeJSON()
		return
	}
	ord_dtl, code, err := models.ReturnGoodsOrder(order_ret)
	if err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	data := map[string]interface{}{
		"order_id": ord_dtl.Id,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "退换货成功")
	this.ServeJSON()
	return
}

// ReturnGoodsShipNumber @Title ReturnGoodsShipNumber finished
// @Description 输入退货的运单号 ReturnGoodsShipNumber
// @Success 200 status bool, data interface{}, msg string
// @router /return_ship_number [post]
func (this *OrderController) ReturnGoodsShipNumber() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	_, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	var order_ship order.ReturnShipNumber
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &order_ship); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := order_ship.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	err = models.UpdReturnShipNumber(order_ship.OrderId, order_ship.RetShipNumber)
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "添加退货快递单号失败")
		this.ServeJSON()
		return
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "添加快递单号成功")
	this.ServeJSON()
	return
}


// ReturnGoodsApproval @Title ReturnGoodsApproval finished
// @Description 退货申述 ReturnGoodsApproval
// @Success 200 status bool, data interface{}, msg string
// @router /return_goods_approval [post]
func (this *OrderController) ReturnGoodsApproval() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	_, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	var order_approval order.ReturnGoodsApprovalReq
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &order_approval); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := order_approval.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	err, msg := models.OrderApproval(order_approval.OrderId, order_approval.AdjustContent)
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, nil, msg)
		this.ServeJSON()
		return
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "申诉成功")
	this.ServeJSON()
	return
}


// CancleReturnGoodsOrder @Title CancleReturnGoodsOrder finished
// @Description 撤销换退货 CancleReturnGoodsOrder
// @Success 200 status bool, data interface{}, msg string
// @router /cancle_return_goods_order [post]
func (this *OrderController) CancleReturnGoodsOrder() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	user_, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	var cancle_order order.CancleReturnGoodsOrderReq
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &cancle_order); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := cancle_order.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	err = models.RemoveOrderProcess(cancle_order.OrderId, user_.Id)
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "没有这个退货记录")
		this.ServeJSON()
		return
	}
	order_dtl, code, err := models.GetGoodsOrderDetail(cancle_order.OrderId)
	if err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	order_dtl.IsCancle = 0
	err = order_dtl.Update()
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, err, err.Error())
		this.ServeJSON()
		return
	}
	data := map[string]interface{}{
		"order_id": order_dtl.Id,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "撤销退换货成功")
	this.ServeJSON()
	return
}

// ConfirmRecvGoods @Title ConfirmRecvGoods finished
// @Description 确认收货 ConfirmRecvGoods
// @Success 200 status bool, data interface{}, msg string
// @router /confirm_revc_goods [post]
func (this *OrderController) ConfirmRecvGoods() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	_, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	var cancle_order order.CancleReturnGoodsOrderReq
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &cancle_order); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := cancle_order.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	order_dtl, code, err := models.GetGoodsOrderDetail(cancle_order.OrderId)
	if err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	order_dtl.OrderStatus = models.OrderStatusRecvGoods
	err = order_dtl.Update()
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, err, err.Error())
		this.ServeJSON()
		return
	}
	data := map[string]interface{}{
		"order_id": order_dtl.Id,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "确认收货成功")
	this.ServeJSON()
	return
}
