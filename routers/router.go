package routers

import (
	"github.com/astaxie/beego/context"
	"ttsx/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.InsertFilter("/goods/*", beego.BeforeExec, filterFunc)
	beego.Router("/", &controllers.GoodsController{}, "get:ShowIndex")
	beego.Router("/register", &controllers.UserController{}, "get:ShowRegister;post:HandleRegister")

	beego.Router("/login", &controllers.LoginController{}, "get:ShowLogin;post:HandleLogin")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("/active", &controllers.UserController{}, "get:ActiveUser")
	beego.Router("/receiver", &controllers.UserController{}, "post:AddAddress")
	//beego.Router("/receiver", &controllers.UserController{}, "post:AddAddress")
	beego.Router("/goods/usercenterinfo", &controllers.GoodsController{}, "get:ShowUserCenterInfo")
	beego.Router("/goods/usercenterorder", &controllers.GoodsController{}, "get:ShowUserCenterOrder")
	beego.Router("/goods/usercentersite", &controllers.GoodsController{}, "get:ShowUserCenterSite")
	beego.Router("/goods/detail", &controllers.GoodsController{}, "get:ShowDetail")
	beego.Router("/goods/types", &controllers.GoodsController{}, "get:ShowList")
}


var filterFunc = func(ctx *context.Context){
	 username := ctx.Input.Session("userName")
	 if username == nil{
		 beego.Error("user not login ....")
		 ctx.Redirect(302, "/login")
		 return
	 }
	 ctx.Input.SetData("userName" ,username)
}