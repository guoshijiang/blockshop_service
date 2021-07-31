package api

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}


// FeaturedStore @Title FeaturedStore
// @Description 特色商家 FeaturedStore
// @Success 200 status bool, data interface{}, msg string
// @router /featured_store [post]
func (this *UserController) FeaturedStore() {

}


// IndexOtherInfo @Title IndexOtherInfo
// @Description 首页信息数据 IndexOtherInfo
// @Success 200 status bool, data interface{}, msg string
// @router /index_info [post]
func (this *UserController) IndexOtherInfo() {

}
