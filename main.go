package main

import (
  _ "blockshop/common/template"
  _ "blockshop/routers"
  _ "blockshop/session"
  "github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

