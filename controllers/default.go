package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Layout = "main/layout.html"
	c.TplNames = "main/default.html"
}

// func (this *MainController) Welcome() {
// 	this.TplNames = "main/welcome.html"
// }

// func (this *MainController) Test() {
// 	var showtype = this.GetString("st")
// 	if showtype != "jbox" {
// 		this.Layout = "main/layout.html"
// 		this.TplNames = "main/test.html"

// 	} else {
// 		this.TplNames = "main/test3.html"
// 	}
// }
