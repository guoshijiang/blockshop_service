package api

import (
	"blockshop/models"
	"blockshop/types"
	"blockshop/types/collect"
	"encoding/json"
	"github.com/astaxie/beego"
	"strings"
)

type GoodsCollectController struct {
	beego.Controller
}

// AddGoodsCollect @Title AddGoodsCollect finished
// @Description 收藏商品 AddGoodsCollect
// @Success 200 status bool, data interface{}, msg string
// @router /add_goods_collect [post]
func (this *GoodsCollectController) AddGoodsCollect() {
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
	goods_c := collect.GoodsCollectReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &goods_c); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	msg, code := models.AddGoodsCollect(goods_c)
	if code != types.ReturnSuccess {
		this.Data["json"] = RetResource(false, types.SystemDbErr, err, msg)
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = RetResource(true, types.ReturnSuccess, err, "收藏商品成功")
		this.ServeJSON()
		return
	}
}


// GetAddGoodsCollectList @Title GetAddGoodsCollectList finished
// @Description 收藏商品列表 GetAddGoodsCollectList
// @Success 200 status bool, data interface{}, msg string
// @router /get_goods_collect_list [post]
func (this *GoodsCollectController) GetAddGoodsCollectList() {
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
	page_s := types.PageSizeData{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &page_s); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	bl_lst, total := models.GoodsCollectList(page_s.Page, page_s.PageSize)
	var gdsc_list []*collect.GoodsListRep
	for _, value := range bl_lst {
		goods_dtl, _, _ := models.GetGoodsDetail(value.CtGdsId)
		if goods_dtl != nil {
			gds_c := &collect.GoodsListRep{
				GdsCollectId: value.CtGdsId,
				GoodsTitle: goods_dtl.Title,
				GoodsName: goods_dtl.GoodsName,
				Views: 1,
				SellNum: goods_dtl.SellNums,
				LeftNum: goods_dtl.LeftAmount,
				GoodsPrice: goods_dtl.GoodsPrice,
				GoodsBtcAmount: 10.0,
				GoodsUsdtAmount: 0.0,
				IsAdmin: goods_dtl.IsAdmin,
			}
			gdsc_list =append(gdsc_list, gds_c)
		} else {
			continue
		}
	}
	data := map[string]interface{}{
		"total": total,
		"data":  gdsc_list,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取收藏的商品列表成功")
	this.ServeJSON()
	return
}


// RemoveGoodsCollect @Title RemoveGoodsCollect finished
// @Description 移除收藏商品 RemoveGoodsCollect
// @Success 200 status bool, data interface{}, msg string
// @router /remove_goods_collect [post]
func (this *GoodsCollectController) RemoveGoodsCollect() {
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
	gds_c_del := collect.GoodsCollectDelReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &gds_c_del); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	msg, code := models.RemoveGoodsCollect(gds_c_del.GoodsCollecId)
	if code != types.ReturnSuccess {
		this.Data["json"] = RetResource(false, types.SystemDbErr, err, msg)
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = RetResource(true, types.ReturnSuccess, err, "移除收藏商品成功")
		this.ServeJSON()
		return
	}
}
