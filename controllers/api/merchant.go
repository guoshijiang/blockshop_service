package api

import (
	"blockshop/models"
	"blockshop/types"
	"blockshop/types/merchant"
	"encoding/json"
	"github.com/astaxie/beego"
)

type MerchantController struct {
	beego.Controller
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
	merchant_list, total, err := models.GetMerchantList(gds_merchant.Page, gds_merchant.PageSize, gds_merchant.MerchantName, gds_merchant.MerchantAddress)
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
		CreatedAt:mcrt_detail.CreatedAt,
		UpdatedAt: mcrt_detail.UpdatedAt,
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

}


// MerchantUpdGoods @Title MerchantUpdGoods
// @Description 商家修改商品 MerchantUpdGoods
// @Success 200 status bool, data interface{}, msg string
// @router /mct_upd_goods [post]
func (this *MerchantController) MerchantUpdGoods() {

}


// MerchantGoodsList @Title MerchantGoodsList
// @Description 商家商品列表 MerchantGoodsList
// @Success 200 status bool, data interface{}, msg string
// @router /mct_goods_list [post]
func (this *MerchantController) MerchantGoodsList() {

}


// MerchantDelGoods @Title MerchantDelGoods
// @Description 商家商品删除 MerchantDelGoods
// @Success 200 status bool, data interface{}, msg string
// @router /mct_del_goods [post]
func (this *MerchantController) MerchantDelGoods() {

}


// MerchantGoodsCmtList @Title MerchantGoodsCmtList
// @Description 商品商品评价 MerchantGoodsCmtList
// @Success 200 status bool, data interface{}, msg string
// @router /mct_goods_cmt_list [post]
func (this *MerchantController) MerchantGoodsCmtList() {

}


// MerchantGoodsOrderList @Title MerchantGoodsOrderList
// @Description 商品商品订单 MerchantGoodsOrderList
// @Success 200 status bool, data interface{}, msg string
// @router /mct_goods_order_list [post]
func (this *MerchantController) MerchantGoodsOrderList() {

}


// MerchantGoodsOrderDetail @Title MerchantGoodsOrderDetail
// @Description 商品商品订单 MerchantGoodsOrderDetail
// @Success 200 status bool, data interface{}, msg string
// @router /mct_goods_order_detail [post]
func (this *MerchantController) MerchantGoodsOrderDetail() {

}


// MerchantGoodsUpdOrder @Title MerchantGoodsUpdOrder
// @Description 商品商品订单状态维护 MerchantGoodsUpdOrder
// @Success 200 status bool, data interface{}, msg string
// @router /mct_goods_upd_order [post]
func (this *MerchantController) MerchantGoodsUpdOrder() {

}