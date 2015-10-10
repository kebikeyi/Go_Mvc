package main

import (
	_ "Go_Mvc/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SessionOn = true
	beego.SetStaticPath("/img", "static/img")
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/js", "static/js")
	beego.Run()
}
