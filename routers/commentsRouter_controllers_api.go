package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["blockshop/controllers/api:ImageController"] = append(beego.GlobalControllerRouter["blockshop/controllers/api:ImageController"],
        beego.ControllerComments{
            Method: "UploadFiles",
            Router: "/upload_file",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blockshop/controllers/api:UserController"] = append(beego.GlobalControllerRouter["blockshop/controllers/api:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blockshop/controllers/api:UserController"] = append(beego.GlobalControllerRouter["blockshop/controllers/api:UserController"],
        beego.ControllerComments{
            Method: "Register",
            Router: "/register",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
