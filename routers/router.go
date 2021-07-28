package routers

import (
	"blockshop_service/controllers/api"
	"github.com/astaxie/beego"
)

func init() {
	api_path := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&api.UserController{},
			),
		),
	)
	beego.AddNamespace(api_path)
}
