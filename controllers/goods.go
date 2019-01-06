package controllers

import (
	"github.com/astaxie/beego/orm"
	"ttsx/models"
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
	this.Layout = "layout.html"
	this.TplName = "user_center_info.html"
}

func(this *GoodsController)ShowUserCenterOrder(){
	this.Layout = "layout.html"
	this.TplName = "user_center_order.html"
}
func(this *GoodsController)ShowUserCenterSite(){

	//显示defaultaddress
	userName := this.GetSession("userName").(string)
	var receiver models.Receiver
	o := orm.NewOrm()
	err := o.QueryTable("Receiver").Filter("User__UserName", userName).Filter("IsDefault", true).One(&receiver)
	if err ==nil{
		this.Data["addr"] = receiver
	}
	this.Layout = "layout.html"
	this.TplName = "user_center_site.html"
}