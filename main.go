package main

import (
	_ "blockshop/common/template"
	"blockshop/cron"
	_ "blockshop/routers"
	_ "blockshop/session"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	//orm.Debug = true
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	beego.BConfig.WebConfig.Session.SessionOn = true             //开启Session模块
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 86400 //设置Session有效期,单位秒
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 86400
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	go cron.Run()
	beego.Run()
}


