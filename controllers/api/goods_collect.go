package api

import (
	"blockshop/models"
	"blockshop/types"
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
}
