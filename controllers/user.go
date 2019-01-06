package controllers

import (
	"github.com/astaxie/beego/orm"
	"ttsx/models"
	"github.com/astaxie/beego"
	"regexp"
)



type UserController struct{
	beego.Controller
}

func(this *UserController)ShowRegister(){
	 this.TplName = "register.html"
}

func(this *UserController)HandleRegister(){
	userName := this.GetString("user_name")
	psw := this.GetString("pwd")
	cpwd := this.GetString("cpwd")
	email := this.GetString("email")

	if userName == "" || psw== "" || cpwd=="" || email==""{
		beego.Error("注册数据为空")
		this.Data["error"] = "数据为空"
		this.TplName = "register.html"
	}
	if psw !=cpwd{
		beego.Error("两次密码不同")
		this.Data["error"] = "密码输入错误"
		this.TplName = "register.html"
	}

	reg :=regexp.MustCompile(`\w+@.*\.(.*)?`)
	result := reg.FindStringSubmatch(email)
	if len(result)==0{
		beego.Error("邮箱创建失败")
		this.Data["error"] = "邮箱数据错误"
		this.TplName = "register.html"
	}

	var user models.User
	user.UserName = userName
	user.Password = psw
	user.Email = email
	o := orm.NewOrm()
	if _, err := o.Insert(&user);err!= nil{
		beego.Error("用户创建失败")
		this.Data["error"] = "用户创建失败"
		this.TplName = "register.html"
	}
	beego.Info("user register succ:", userName)
	this.Redirect("/index", 302)
}