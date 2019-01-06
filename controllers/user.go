package controllers

import (
	"regexp"
	"strconv"
	"ttsx/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) ShowRegister() {
	this.TplName = "register.html"
}

func (this *UserController) HandleRegister() {
	userName := this.GetString("user_name")
	psw := this.GetString("pwd")
	cpwd := this.GetString("cpwd")
	email := this.GetString("email")

	if userName == "" || psw == "" || cpwd == "" || email == "" {
		beego.Error("注册数据为空")
		this.Data["error"] = "数据为空"
		this.TplName = "register.html"
	}
	if psw != cpwd {
		beego.Error("两次密码不同")
		this.Data["error"] = "密码输入错误"
		this.TplName = "register.html"
	}

	reg := regexp.MustCompile(`\w+@.*\.(.*)?`)
	result := reg.FindStringSubmatch(email)
	if len(result) == 0 {
		beego.Error("邮箱创建失败")
		this.Data["error"] = "邮箱数据错误"
		this.TplName = "register.html"
	}

	var user models.User
	user.UserName = userName
	user.Password = psw
	user.Email = email
	o := orm.NewOrm()
	if _, err := o.Insert(&user); err != nil {
		beego.Error("用户创建失败")
		this.Data["error"] = "用户创建失败"
		this.TplName = "register.html"
	}
	beego.Info("user register succ:", userName)
	//this.Redirect("/index", 302)
	//发送激活邮箱
	emailSender := utils.NewEMail(`{"username":"935233292@qq.com","password":"zaydhykiohxvbcij","host":"smtp.qq.com","port":587}`)
	emailSender.From = "935233292@qq.com"
	emailSender.To = []string{email}
	emailSender.Subject = "天天生鲜用户注册"
	
	emailSender.HTML = "<a href=\"http://127.0.0.1:8888/active?id=" + strconv.Itoa(int(user.Id)) + " \">点击激活</a>"
	emailSender.Send()

	this.Ctx.WriteString("注册成功,请前往页面激活")
}

func (this *UserController) ActiveUser() {
	id, err := this.GetInt64("id")
	if err != nil {
		beego.Error("参数缺失:")
		this.Ctx.WriteString("激活失败，参数缺失")
		return
	}
	var user models.User
	user.Id = id
	o := orm.NewOrm()
	if err = o.Read(&user); err != nil {
		beego.Error("用户不存在")
		this.Ctx.WriteString("激活失败，参数缺失")
		return
	}
	user.Active = true
	o.Update(&user)
	this.Redirect("/login", 302)
}



func(this *UserController)AddAddress(){
	name := this.GetString("name")
	address:= this.GetString("address")
	zipcode := this.GetString("zipcode")
	phone := this.GetString("phone")
	if name=="" || address == "" ||zipcode=="" || phone == ""{
		beego.Error("the specify param is not present!")
		this.Redirect("/", 302)
		return
	}
	var receiver models.Receiver
	receiver.Address = address
	receiver.Name = name
	receiver.Phone = phone
	receiver.ZipCode = zipcode

	userName := this.GetSession("userName")
	if userName==nil{
		beego.Error("user is not in login status")
		this.Redirect("/", 302)
		return
	}
	o := orm.NewOrm()
	o.Begin()
	var user models.User
	user.UserName = userName.(string)
    if err := o.Read(&user, "UserName");err!=nil{
		beego.Error("user not exists!", userName, err)
		this.Redirect("/", 302)
		return
	}
	if c, err := o.QueryTable("Receiver").Filter("User__Id", user.Id).Count();err!=nil || c==0{
		receiver.IsDefault = true
	}
	receiver.User = &user
	if _, err :=o.Insert(&receiver);err!=nil{
		beego.Error("error insert the receiver")
		o.Rollback()
		this.Redirect("/", 302)
		return
	}
	o.Commit()
	
	this.Redirect("/goods/usercentersite", 302)
}
