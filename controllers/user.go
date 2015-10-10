package controllers

import (
	. "Go_Mvc/models"
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Index() {
	var user = new(TbUsers).Init()
	fmt.Println(user)
	c.TplNames = "user/index.html"
}
