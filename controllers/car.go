package controllers

import (
	"github.com/astaxie/beego"
)
type CarController struct{
	beego.Controller
}
func (this *CarController) List(){
	this.TplNames="main/layout.html"
}
