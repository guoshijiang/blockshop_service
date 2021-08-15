package api

import (
  "blockshop/models"
  "blockshop/types"
  "blockshop/types/merchant"
  "encoding/json"
  "github.com/astaxie/beego"
  "strings"
)

type MerchantController struct {
	beego.Controller
}

// MerchantOpenFee @Title MerchantOpenFee
// @Description 商家开通支付费用 MerchantOpenFee
// @Success 200 status bool, data interface{}, msg string
// @router /marchant_open_fee [post]
func (this *MerchantController) MerchantOpenFee() {
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
	merchant_config := models.GetMerchantConfig()
	if merchant_config != nil {
		data := map[string]interface{}{
			"btc_amount": merchant_config.BtcAmount,
			"usdt_amount": merchant_config.UsdtAmount,
		}
		this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取商家开通费率成功")
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = RetResource(false, types.OpenMerchantFail, nil, "商家模块权限没有对您开放")
		this.ServeJSON()
		return
	}
}


// OpenMerchant @Title OpenMerchant
// @Description 商家开通支付费用 OpenMerchant
// @Success 200 status bool, data interface{}, msg string
// @router /open_marchant [post]
func (this *MerchantController) OpenMerchant() {
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
	open_merchant := merchant.OpenMerchantReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &open_merchant); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	msg, err, id := models.OpenMerchant(open_merchant)
	if id == 0 {
		this.Data["json"] = RetResource(false, types.OpenMerchantFail, msg, "开通商家失败")
		this.ServeJSON()
		return
	}
	data := map[string]interface{}{
		"id":id,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "开通商家成功")
	this.ServeJSON()
	return
}

// MerchantList @Title MerchantList
// @Description 商家列表接口 MerchantList
// @Success 200 status bool, data interface{}, msg string
// @router /marchant_list [post]
func (this *MerchantController) MerchantList() {
	gds_merchant := merchant.MerchantListReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &gds_merchant); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := gds_merchant.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	merchant_list, total, err := models.GetMerchantList(gds_merchant.Page, gds_merchant.PageSize, gds_merchant.IsShow)
	if err != nil {
		this.Data["json"] = RetResource(false, types.GetMerchantListFail, nil, "获取商家列表失败")
		this.ServeJSON()
		return
	}
	image_path := beego.AppConfig.String("img_root_path")
	var mct_list_ret []merchant.MerchantListRep
	for _, mct := range merchant_list {
		mct_ret := merchant.MerchantListRep {
			MctId: mct.Id,
			MctName: mct.MerchantName,
			MctIntroduce: mct.MerchantIntro,
			MerchantDetail: mct.MerchantDetail,
			MctLogo: image_path + mct.Logo,
			MctWay: mct.MerchantWay,
			ShopLevel: mct.ShopLevel,
			ShopServer: mct.ShopServer,
		}
		mct_list_ret = append(mct_list_ret, mct_ret)
	}
	data := map[string]interface{}{
		"total":     total,
		"mct_lst": mct_list_ret,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取商家列表成功")
	this.ServeJSON()
	return
}

// MerchantDetail @Title MerchantDetail
// @Description 商家详情接口 MerchantDetail
// @Success 200 status bool, data interface{}, msg string
// @router /marchant_detail [post]
func (this *MerchantController) MerchantDetail() {
	merchant_dtil := merchant.MerchantDetailReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &merchant_dtil); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := merchant_dtil.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	image_path := beego.AppConfig.String("img_root_path")
	mcrt_detail, code, err := models.GetMerchantDetail(merchant_dtil.MerchantId)
	if err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	m_goods_nums := models.GetMerchantGoodsNums(merchant_dtil.MerchantId)
  WaidPayOrderNum,WaitSendOrderNum,SendOrderNum := (new(models.GoodsOrder)).Aggregation(merchant_dtil.MerchantId)
  WaitReturnOrderNum := new(models.OrderProcess).WaitReturnOrderTotal()
	order_stat := &merchant.OrderDataStat{
		WaidPayOrderNum: WaidPayOrderNum,
		WaitSendOrderNum:  WaitSendOrderNum,
		WaitReturnOrderNum: WaitReturnOrderNum,
		SendOrderNum:SendOrderNum,
	}
	goods_stat := &merchant.GoodsDataStat{
		OnSaleNum: models.GetMerchantGoodsIsSale(merchant_dtil.MerchantId,0),
		SoldOutNum:models.GetMerchantGoodsIsSale(merchant_dtil.MerchantId,1),
		OffShelfNum: models.GetMerchantGoodsEmpty(merchant_dtil.MerchantId),
	}

  merchant_state,_ := (&models.MerchantStat{MerchantId:merchant_dtil.MerchantId}).QueryByMerchant()
	comment_stat :=  &merchant.CommentDataStat{
		SericeBest: merchant_state.ServiceBest,
		ServiceGood: merchant_state.ServiceGood,
		ServiceBad: merchant_state.ServiceBad,
		ServiceAvg: merchant_state.ServiceAvg,
		TradeBest: merchant_state.TradeBest,
		TradeGood: merchant_state.TradeGood,
		TradeBad: merchant_state.TradeBad,
		TradeAvg:merchant_state.ServiceAvg,
		QualityBest: merchant_state.TradeBest,
		QualityGood: merchant_state.QualityGood,
		QualityBad: merchant_state.QualityBad,
		QualityAvg:merchant_state.QualityAvg,
	}
	mct_ret_dtl := merchant.MerchantDetailRep{
		MctId: mcrt_detail.Id,
		MctLogo: image_path + mcrt_detail.Logo,
		MctName: mcrt_detail.MerchantName,
		MctIntroduce: mcrt_detail.MerchantIntro,
		MerchantDetail: mcrt_detail.MerchantDetail,
		Address: mcrt_detail.Address,
		GoodsNum: m_goods_nums,
		MctWay: mcrt_detail.MerchantWay,
		ShopLevel: mcrt_detail.ShopLevel,
		ShopServer: mcrt_detail.ShopServer,
		OrderStat: order_stat,
		GoodsStat: goods_stat,
		CommentStat: comment_stat,
		ShopScore: mcrt_detail.ShopScore,
		MonthSellNum: mcrt_detail.MonthSellNum,
		MonthSellAmount: mcrt_detail.MonthSellAmount,
		TotalSellNum: mcrt_detail.TotalSellNum,
		TotalSellAmount: mcrt_detail.TotalSellAmount,
		AdjustVictor: mcrt_detail.AdjustVictor,
		AdjustFail: mcrt_detail.AdjustFail,
		JoinTime:mcrt_detail.CreatedAt.Format("2006-01-02 15:04:05"),
		LstLoginTime: mcrt_detail.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, mct_ret_dtl, "获取商家详情成功")
	this.ServeJSON()
	return
}

// MerchantAddGoods @Title MerchantAddGoods
// @Description 商家新增商品 MerchantAddGoods
// @Success 200 status bool, data interface{}, msg string
// @router /mct_add_goods [post]
func (this *MerchantController) MerchantAddGoods() {
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
	merchant_goods_add := merchant.MerchantAddUpdGoodsReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &merchant_goods_add); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := merchant_goods_add.GoodsAddParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	code, err := models.CreateMerchantGoods(merchant_goods_add)
	if code == types.ReturnSuccess {
		this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "添加商品成功")
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = RetResource(false, types.CreateGoodsFail, nil, "添加商品失败")
		this.ServeJSON()
		return
	}
}


// MerchantUpdGoods @Title MerchantUpdGoods
// @Description 商家修改商品 MerchantUpdGoods
// @Success 200 status bool, data interface{}, msg string
// @router /mct_upd_goods [post]
func (this *MerchantController) MerchantUpdGoods() {
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
	merchant_goods_upd := merchant.UpdateGoodsReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &merchant_goods_upd); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := merchant_goods_upd.GoodsUpdParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	code, err := models.UpdateMerchantGoods(merchant_goods_upd)
	if code == types.ReturnSuccess {
		this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "修改商品信息成功")
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = RetResource(false, types.CreateGoodsFail, nil, "修改商品信息失败")
		this.ServeJSON()
		return
	}
}


// MerchantDelGoods @Title MerchantDelGoods
// @Description 商家商品删除 MerchantDelGoods
// @Success 200 status bool, data interface{}, msg string
// @router /mct_del_goods [post]
func (this *MerchantController) MerchantDelGoods() {
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
	merchant_goods_del := merchant.DeleteGoodsReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &merchant_goods_del); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := merchant_goods_del.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	code, err := models.DeleteGoodsById(merchant_goods_del.GoodsId)
	if code == types.ReturnSuccess {
		this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "删除商品成功")
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = RetResource(false, types.CreateGoodsFail, nil, "删除商品失败")
		this.ServeJSON()
		return
	}
}


// MerchantStaticDetail @Title MerchantStaticDetail
// @Description 商家统计 MerchantStaticDetail
// @Success 200 status bool, data interface{}, msg string
// @router /mct_static_detail [post]
func (this *MerchantController) MerchantStaticDetail() {
}
