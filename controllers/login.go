package controllers

import (
	. "carmaint/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.Data["website"] = "beego.me"

	fmt.Println("login.user", User2.Query("select * from tbusers"))
	fmt.Println("login.Service.User", User2.Query("select * from tbusers"))
	fmt.Println("login.Service.User.Get(2)", User2.Get(2))
	this.TplNames = "login/login.html"
}

func (this *LoginController) post() {
	var username = this.GetString("uname")
	var password = this.GetString("pwd")

	var user = new(TbUsers).Init()
	var b = user.Login(username, password)
	if false == b {
		this.Ctx.WriteString("用户名或密码错误！")
	}
	this.SetSession("uid", user.ID)
	this.SetSession("uname", username)
	this.Ctx.WriteString("1")
}
