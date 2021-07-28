package api


import "github.com/astaxie/beego"

const HttpAuthKey = "Authorization"


type baseController struct {
	beego.Controller
}

type RetJson struct {
	Status bool        `json:"status"`
	Code   int         `json:"code"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}

func RetResource(status bool, code int, data interface{}, msg string) (apijson *RetJson) {
	apijson = &RetJson{Status: status, Code: code, Data: data, Msg: msg}
	return
}

