package routers

import (
	"ttsx/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.UserController{}, "get:ShowRegister;post:HandleRegister")

	beego.Router("/index", &controllers.IndexController{}, "get:ShowIndex")
	beego.Router("/login", &controllers.LoginController{}, "get:ShowLogin")
	beego.Router("/active", &controllers.UserController{}, "get:ActiveUser")
}
