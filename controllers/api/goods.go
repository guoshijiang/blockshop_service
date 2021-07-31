package api

import "github.com/astaxie/beego"

type GoodsController struct {
	beego.Controller
}

// GoodsList @Title GoodsList
// @Description 随机商品列表 GoodsList
// @Success 200 status bool, data interface{}, msg string
// @router /goods_list [post]
func (this *UserController) GoodsList() {

}