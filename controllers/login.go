package controllers

import (
	"github.com/astaxie/beego/orm"
	"ttsx/models"
	"github.com/astaxie/beego"
)

type LoginController struct{
	beego.Controller
}

func(this *LoginController)ShowLogin(){
	this.Data["userName"] = this.Ctx.GetCookie("userName")
	this.Data["checked"] = this.Ctx.GetCookie("checked")

	this.TplName = "login.html"
}

func(this *LoginController)HandleLogin(){
	 userName := this.GetString("username")
	 passwd :=   this.GetString("pwd")
	 if userName == "" || passwd == ""{
		 beego.Error("参数缺失")
		 this.Data["error"] = "参数失败"
		 this.TplName = "login.html"
		 return
	 }

	 var user models.User
	 user.UserName = userName
	 
	 o := orm.NewOrm()
	 err :=o.Read(&user, "UserName")
	 if err != nil{
		 beego.Error("用户不存在", userName)
		 this.Data["error"] = "用户不存在"
		 this.TplName = "login.html"
		 return
	 }
	 if user.Password != passwd{
		 beego.Error("用户密码错误", userName)
		 this.Data["error"] = "用户密码错误"
		 this.TplName = "login.html"
		 return
	 }
	if !user.Active {
		 beego.Error("用户未激活", userName)
		 this.Data["error"] = "用户未激活"
		 this.TplName = "login.html"
		 return
	}
	remember := this.GetString("remember")
	beego.Info("remember me status ", remember)
	if remember!= "" && remember == "on"{
		//记住用户名
		this.Ctx.SetCookie("userName", userName)
		this.Ctx.SetCookie("checked", "checked")
	}
	this.SetSession("userName", userName)
	beego.Info("用户登录成功", userName)
	this.Redirect("/", 302)

}


func(this *LoginController)Logout(){
	 this.DelSession("userName")
	 this.Redirect("/", 302)
}