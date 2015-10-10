package routers

import (
	"Go_Mvc/controllers"
	_ "fmt"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/context"
)

// var FilterUser = func(ctx *context.Context) {

// }

func init() {
	beego.Router("/", &controllers.MainController{})
	//beego.Router("/hello", &controllers.HelloController{})
	beego.Router("/user/index", &controllers.UserController{}, "get:Index")
	//beego.AutoRouter(&controllers.MainController)
	//beego.AutoRouter(&controllers.HelloController)
	//beego.AutoRouter(&controllers.UserController)
}
