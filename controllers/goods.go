package controllers

import (
	"github.com/astaxie/beego"
)


type GoodsController struct{
	beego.Controller
}

func(this *GoodsController)ShowIndex(){
	username := this.GetSession("userName")
	if username == nil{
		beego.Error("uername is empty")
		this.Data["userName"] = ""
	}else{
		this.Data["userName"] = username
	}
	beego.Info("userName", username)
	this.TplName = "index.html"
}

func(this *GoodsController)ShowUserCenterInfo(){
	userName := this.GetSession("userName")
	if userName == nil{
		beego.Error("user not login")
	}else{
		this.Data["userName"] = userName
	}
	this.Layout = "layout.html"
	this.TplName = "user_center_info.html"
}

func(this *GoodsController)ShowUserCenterOrder(){
	this.Layout = "layout.html"
	this.TplName = "user_center_order.html"
}
func(this *GoodsController)ShowUserCenterSite(){
	this.Layout = "layout.html"
	this.TplName = "user_center_site.html"
}