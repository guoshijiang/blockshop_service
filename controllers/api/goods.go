package api

import (
	"blockshop/models"
	"blockshop/types"
	"blockshop/types/goods"
	"encoding/json"
	"github.com/astaxie/beego"
)

type GoodsController struct {
	beego.Controller
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
	image_path := beego.AppConfig.String("img_root_path")
	var goods_ret_list []goods.GoodsListRep
	for _, value := range good_list {
		gds_ret := goods.GoodsListRep{
			GoodsId:   value.Id,
			Title: value.Title,
			Logo: image_path + value.Logo,
			GoodsPrice: value.GoodsPrice,
			GoodsDisPrice: value.GoodsDisPrice,
			IsDiscount: value.IsDiscount,
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
	img_path := beego.AppConfig.String("img_root_path")
	merchant, code, err := models.GetMerchantDetail(goods_dtl.MerchantId)
	if err != nil {
		this.Data["json"] = RetResource(false, code, err.Error(), "获取商家信息失败")
		this.ServeJSON()
		return
	}
	merchant_info :=  map[string]interface{}{
		"merchant_id": merchant.Id,
		"merchant_logo": img_path + merchant.Logo,
		"merchant_name": merchant.MerchantName,
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
			ImageUrl: img_path + v.Image,
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
	type_list_data, _, err := models.GetGoodsAttrList(goods_dtl.Id)
	var type_list []goods.GoodsAttrRet
	if err != nil || type_list_data == nil {
		type_list = nil
	} else {
		for _, value_t := range type_list_data {
			var value_list []string
			json.Unmarshal([]byte(value_t.TypeVale), &value_list)
			c_gds_type := goods.GoodsAttrRet{
				GdsAttrKey: value_t.TypeKey,
				GdsAttrValue: value_list,
			}
			type_list = append(type_list, c_gds_type)
		}
	}
	goods_detail := map[string]interface{}{
		"id": goods_dtl.Id,
		"title": goods_dtl.Title,
		"mark": goods_dtl.GoodsMark,
		"logo": img_path + goods_dtl.Logo,
		"serveice": goods_dtl.Serveice,
		"calc_way": goods_dtl.CalcWay,
		"sell_nums": goods_dtl.SellNums,
		"total_amount": goods_dtl.TotalAmount,
		"left_amount": goods_dtl.LeftAmount,
		"goods_price": goods_dtl.GoodsPrice,
		"goods_dis_price": goods_dtl.GoodsDisPrice,
		"goods_name": goods_dtl.GoodsName,
		"goods_params": goods_dtl.GoodsParams,
		"goods_detail": goods_dtl.GoodsDetail,
		"goods_img": gds_img_lst,
		"user_address": user_address,
		"merchant_info": merchant_info,
		"is_hot": goods_dtl.IsHot,
		"is_discount": goods_dtl.IsDiscount,
		"goods_types": type_list,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, goods_detail, "获取商品详情成功")
	this.ServeJSON()
	return
}
