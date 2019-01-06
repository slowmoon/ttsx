package models

import (
	"github.com/astaxie/beego/orm"
	
)




type User struct{
	Id  int64  
	UserName string `orm:"unique;size(100)"`  //用户名
	Password string `orm:"size(20)"`    //密码
	Email string          //邮箱
	Power byte   `orm:"default(0)"` //0:普通用户， 1: 管理员 （用户权限）
	Active bool  `orm:"default(0)"` //0 :未激活 1：激活 
	Receivers []*Receiver `orm:"reverse(many)"`
}

type Receiver struct{
	Id int64
	Name string   //收件人
	ZipCode string   //邮编
	Address string  // 地址
	Phone string    //电话
	IsDefault bool `orm:"default(false)"`  //默认
	User *User    `orm:"rel(fk)"`  
}


func init(){
	  orm.RegisterDataBase("default", "mysql", "root:244121@tcp(host.docker.internal:3306)/ttsx?charset=utf8")
	  orm.RegisterModel(new(User), new(Receiver))
	  orm.RunSyncdb("default", false, true)
}