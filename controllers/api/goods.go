package api

import (
	"blockshop/models"
	"blockshop/types"
	"blockshop/types/goods"
	"blockshop/types/user"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
)

type GoodsController struct {
	beego.Controller
}

func orderByCdt() interface{} {
	var odb_list []*goods.OrderBy
	odb_time := &goods.OrderBy{
		Way: 1,
		WayName: "时间",
	}
	odb_list = append(odb_list, odb_time)
	odb_xl := &goods.OrderBy{
		Way: 2,
		WayName: "销量",
	}
	odb_list = append(odb_list, odb_xl)
	odb_jg := &goods.OrderBy{
		Way: 3,
		WayName: "价格",
	}
	odb_list = append(odb_list, odb_jg)
	odb_sj := &goods.OrderBy{
		Way: 4,
		WayName: "商家",
	}
	odb_list = append(odb_list, odb_sj)
	return odb_list
}


// GoodsQueryCondition @Title GoodsQueryCondition
// @Description 商品查询条件 GoodsQueryCondition
// @Success 200 status bool, data interface{}, msg string
// @router /goods_query_condition [post]
func (this *GoodsController) GoodsQueryCondition() {
	goods_type_list := models.GetGdsTypeList()
	var gds_types []*goods.GoodsType
	if goods_type_list != nil {
		for _, goods_type := range goods_type_list {
			gsc := &goods.GoodsType{
				Id: goods_type.Id,
				TypeName: goods_type.Name,
			}
			gds_types = append(gds_types, gsc)
		}
	} else {
		gds_types = nil
	}
	goods_cat_list := models.GetGdsCatList()
	var gds_cat_list []*goods.GoodsCat
	if goods_cat_list != nil {
		for _, value := range goods_cat_list {
			gcat_item := &goods.GoodsCat{
				Id: value.Id,
				CatName: value.Name,
			}
			gds_cat_list = append(gds_cat_list, gcat_item)
		}
	} else {
		gds_cat_list = nil
	}
	gds_origin_states := models.GetGdsCatList()
	var origin_state_list []*goods.OriginState
	if goods_cat_list != nil {
		for _, value := range gds_origin_states {
			o_i := &goods.OriginState{
				Id: value.Id,
				StateName: value.Name,
			}
			origin_state_list = append(origin_state_list, o_i)
		}
	} else {
		origin_state_list = nil
	}
	var btc_c_price user.CoinPrice
	btc_price := models.GetAssetByName("BTC")
	if btc_price != nil {
		btc_c_price = user.CoinPrice {
			Asset: "BTC",
			ChainName: "Bitcoin",
			UsdPrice: btc_price.UsdPrice,
			CnyPrice: btc_price.CnyPrice,
		}
	} else {
		btc_c_price = user.CoinPrice {
			Asset: "BTC",
			ChainName: "Bitcoin",
			UsdPrice: "0",
			CnyPrice: "0",
		}
	}
	var usdt_price user.CoinPrice
	usdt_price_md := models.GetAssetByName("USDT")
	if usdt_price_md != nil {
		usdt_price = user.CoinPrice {
			Asset: "USDT",
			ChainName: "Trc20",
			UsdPrice: usdt_price_md.UsdPrice,
			CnyPrice: usdt_price_md.CnyPrice,
		}
	} else {
		usdt_price = user.CoinPrice {
			Asset: "USDT",
			ChainName: "Trc20",
			UsdPrice: "0",
			CnyPrice: "0",
		}
	}
	data := map[string]interface{}{
		"goods_type": gds_types,
		"goods_cat": gds_cat_list,
		"origin_state": origin_state_list,
		"order_by": orderByCdt(),
		"pay_way": [2]string{"BTC", "USDT"},
		"btc_price": btc_c_price,
		"usdt_price": usdt_price,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取查询条件成功")
	this.ServeJSON()
	return
}

// GoodsList @Title GoodsList
// @Description 随机商品列表 GoodsList
// @Success 200 status bool, data interface{}, msg string
// @router /goods_list [post]
func (this *GoodsController) GoodsList() {
	goods_lst_req := goods.GoodsListReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &goods_lst_req); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := goods_lst_req.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}

	good_list, total, err := models.GetGoodsList(goods_lst_req)
	if err != nil {
		this.Data["json"] = RetResource(false, types.GetGoodsListFail, nil, err.Error())
		this.ServeJSON()
		return
	}
	//image_path := beego.AppConfig.String("img_root_path")
	var goods_ret_list []goods.GoodsListRep
	for _, value := range good_list {
		merchant, code, err := models.GetMerchantDetail(value.MerchantId)
		if err != nil {
			this.Data["json"] = RetResource(false, code, err.Error(), "获取商家信息失败")
			this.ServeJSON()
			return
		}
		var type_id int64
		var type_name string
		gds_type := models.GetGdsTypeById(value.GoodsTypeId)
		if gds_type != nil {
			type_id =  gds_type.Id
			type_name = gds_type.Name
		} else {
			type_id = 0
			type_name = "未知"
		}
		btc_price_b := 1.0
		usdt_price_u := 1.0
		btc_price := models.GetAssetByName("BTC")
		if btc_price != nil {
			btc_price_b_, _ := strconv.ParseFloat(btc_price.CnyPrice, 64)
			btc_price_b = btc_price_b_
		}
		usdt_price := models.GetAssetByName("USDT")
		if usdt_price != nil {
			usdt_price_u_, _ := strconv.ParseFloat(usdt_price.CnyPrice, 64)
			usdt_price_u = usdt_price_u_
		}
		gds_ret := goods.GoodsListRep{
			GoodsId:   value.Id,
			Title: value.Title,
			MerchantId: merchant.Id,
			MerchantName: merchant.MerchantName,
			TypeId:type_id,
			TypeName: type_name,
			Logo: value.Logo,
			GoodsPrice: value.GoodsPrice,
			GoodsDisPrice: value.GoodsDisPrice,
			Views:value.Views,
			LeftAmount: value.LeftAmount,
			SellNum: value.SellNums,
			BtcPrice: value.GoodsPrice / btc_price_b,
			UsdtPrice: value.GoodsPrice / usdt_price_u,
			IsDiscount: value.IsDiscount,
			IsAdmin: value.IsAdmin,
			IsSale: value.IsSale,
		}
		goods_ret_list = append(goods_ret_list, gds_ret)
	}
	data := map[string]interface{}{
		"total":     total,
		"gds_lst":   goods_ret_list,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取分类商品列表成功")
	this.ServeJSON()
	return
}

// GoodsDetail @Title GoodsDetail
// @Description 商品详情接口 GoodsDetail
// @Success 200 status bool, data interface{}, msg string
// @router /goods_detail [post]
func (this *GoodsController) GoodsDetail() {
	goods_detil := goods.GoodsDetailReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &goods_detil); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := goods_detil.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	goods_dtl, code, err := models.GetGoodsDetail(goods_detil.GoodsId)
	if err != nil {
		this.Data["json"] = RetResource(false, code, err.Error(), "获取商品列表失败")
		this.ServeJSON()
		return
	}
	//img_path := beego.AppConfig.String("img_root_path")
	merchant, code, err := models.GetMerchantDetail(goods_dtl.MerchantId)
	if err != nil {
		this.Data["json"] = RetResource(false, code, err.Error(), "获取商家信息失败")
		this.ServeJSON()
		return
	}
	merchant_info :=  map[string]interface{}{
		"merchant_id": merchant.Id,
		"merchant_logo": merchant.Logo,
		"merchant_name": merchant.MerchantName,
		"merchant_level": merchant.ShopLevel,
	}
	goods_img_lst, code, err := models.GetGoodsImgList(goods_dtl.Id)
	if err != nil {
		this.Data["json"] = RetResource(false, code, err.Error(), "获取商品图片失败")
		this.ServeJSON()
		return
	}
	gds_img_lst := []goods.GoodsImagesRet{}
	for _, v := range goods_img_lst {
		gds_img := goods.GoodsImagesRet{
			GoodsImgId:v.Id,
			ImageUrl: v.Image,
		}
		gds_img_lst = append(gds_img_lst, gds_img)
	}
	user_address := make(map[string]interface{})
	if goods_detil.UserId > 0 {
		user_addr, _, err := models.GetUserAddressDefault(goods_detil.UserId)
		if err != nil {
			user_address = nil
		} else {
			user_address["address_id"] = user_addr.Id
			user_address["address_name"] = user_addr.Address
		}
	} else {
		user_address = nil
	}
	attr_list_data, _, err := models.GetGoodsAttrList(goods_dtl.Id)
	var attr_list []goods.GoodsAttrRet
	if err != nil || attr_list_data == nil {
		attr_list = nil
	} else {
		for _, value_t := range attr_list_data {
			var value_list []string
			json.Unmarshal([]byte(value_t.TypeVale), &value_list)
			c_gds_attr := goods.GoodsAttrRet{
				GdsAttrKey: value_t.TypeKey,
				GdsAttrValue: value_list,
			}
			attr_list = append(attr_list, c_gds_attr)
		}
	}
	var member_level int8
	user_ll, _ := models.GetUserById(goods_dtl.UserId)
	if user_ll != nil {
		member_level = user_ll.MemberLevel
	} else {
		member_level = 0
	}
	var os_sta string
	ors_state := models.GetGdsOsById(goods_dtl.OriginStateId)
	if ors_state != nil {
		os_sta = ors_state.Name
	} else {
		os_sta = "未知"
	}
	btc_price_b := 1.0
	usdt_price_u := 1.0
    btc_price := models.GetAssetByName("BTC")
    if btc_price != nil {
		btc_price_b_, _ := strconv.ParseFloat(btc_price.CnyPrice, 64)
		btc_price_b = btc_price_b_
	}
    usdt_price := models.GetAssetByName("USDT")
    if usdt_price != nil {
		usdt_price_u_, _ := strconv.ParseFloat(usdt_price.CnyPrice, 64)
		usdt_price_u = usdt_price_u_
	}
	var goods_type_name string
	goods_tpe := models.GetGdsTypeById(goods_dtl.GoodsTypeId)
	if goods_tpe != nil {
		goods_type_name = goods_tpe.Name
	} else {
		goods_type_name = "未知类别商品"
	}
	goods_detail := map[string]interface{}{
		"id": goods_dtl.Id,
		"trust_level": member_level,
		"title": goods_dtl.Title,
		"mark": goods_dtl.GoodsMark,
		"logo": goods_dtl.Logo,
		"serveice": goods_dtl.Serveice,
		"origin_state": os_sta,
		"calc_way": goods_dtl.CalcWay,
		"sell_nums": goods_dtl.SellNums,
		"total_amount": goods_dtl.TotalAmount,
		"left_amount": goods_dtl.LeftAmount,
		"goods_price": goods_dtl.GoodsPrice,
		"btc_price": goods_dtl.GoodsPrice / btc_price_b,
		"usdt_price": goods_dtl.GoodsPrice / usdt_price_u,
		"goods_name": goods_dtl.GoodsName,
		"goods_params": goods_dtl.GoodsParams,
		"goods_detail": goods_dtl.GoodsDetail,
		"goods_img": gds_img_lst,
		"user_address": user_address,
		"merchant_info": merchant_info,
		"is_discount": goods_dtl.IsDiscount,
		"is_sale": goods_dtl.IsSale,
		"goods_attr": attr_list,
		"goods_type": goods_type_name,
		"is_admin": goods_dtl.IsAdmin,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, goods_detail, "获取商品详情成功")
	this.ServeJSON()
	return
}
