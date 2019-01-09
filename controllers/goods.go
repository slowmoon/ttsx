package controllers

import (
	"github.com/astaxie/beego/utils/pagination"
	"github.com/gomodule/redigo/redis"
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
	var goodsType []*models.GoodsType
	o := orm.NewOrm()
	o.QueryTable("GoodsType").All(&goodsType)

	var indexGoodsBanners []*models.IndexGoodsBanner
	o.QueryTable("IndexGoodsBanner").OrderBy("Index").All(&indexGoodsBanners)

	var goodsPromotions []*models.IndexPromotionBanner
	o.QueryTable("IndexPromotionBanner").OrderBy("Index").All(&goodsPromotions)

	var goods []map[string]interface{}

	for _, k := range goodsType{
		 cont := make(map[string]interface{})
		 cont["goodsType"] = k 
		 //cont["goodsType"] = k 
		 qs := o.QueryTable("IndexTypeGoodsBanner").RelatedSel("GoodsType", "GoodsSKU").Filter("GoodsType", k)
		 
		 var goodsTexts []*models.IndexTypeGoodsBanner
		 qs.Filter("DisplayType",0).OrderBy("Index").All(&goodsTexts)
		 
		 var goodsImages []*models.IndexTypeGoodsBanner
		 qs.Filter("DisplayType", 1).OrderBy("Index").All(&goodsImages)

		 cont["goodsTexts"] = goodsTexts
		 cont["goodsImages"] = goodsImages
		 goods = append(goods, cont)
	}

	beego.Info("userName", username)
	this.Data["goodsType"] = goodsType
	this.Data["lunbo"] = indexGoodsBanners
	this.Data["promotion"] = goodsPromotions
	this.Data["goods"] = goods

	this.TplName = "index.html"
}

func(this *GoodsController)ShowUserCenterInfo(){
	userName := this.GetSession("userName").(string)
	var receiver models.Receiver
	o := orm.NewOrm()
	if err := o.QueryTable("Receiver").RelatedSel("User").Filter("User__UserName", userName).Filter("IsDefault", true).One(&receiver);err!=nil{
		beego.Info("user default address is not specified ....")
	}	
	
	var goods []models.GoodsSKU
	if userName != ""{
		client ,_  := redis.Dial("tcp", beego.AppConfig.String("hostaddr")+":6379")
		defer client.Close()
		ints, err := redis.Ints(client.Do("lrange", "history_"+userName, 0 , 4))
		if err!= nil{
			beego.Error(err)
		}
		for _, k := range ints{
		var good models.GoodsSKU
			good.Id = k
			if err := o.Read(&good);err!=nil{
				beego.Error(err)
			}else{
				goods = append(goods, good)
			}
		}
	}
	this.Data["goods"] = goods
	this.Data["receiver"] = receiver
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


func(this *GoodsController)ShowDetail(){
	username := this.GetSession("userName")
	if username == nil{
		beego.Error("uername is empty")
		this.Data["userName"] = ""
	}else{
		this.Data["userName"] = username
	}
	
	id ,err := this.GetInt("goodsId")
	if err != nil{
		beego.Error("error info....", err)
		return
	}
	var good models.GoodsSKU
	good.Id = id
	o := orm.NewOrm()
	if err := o.QueryTable("GoodsSKU").RelatedSel("Goods").Filter("Id", good.Id).One(&good);err!= nil{    //商品详情
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	var goods []*models.GoodsSKU     
	_ ,err = o.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType", good.GoodsType).OrderBy("Time", "desc").Limit(2).All(&goods)  
	if  err!= nil{
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	var goodsType []*models.GoodsType
	o.QueryTable("GoodsType").All(&goodsType)
	this.Data["goodsType"] = goodsType

	userName := this.GetSession("userName")
	if userName!= nil{
		name := userName.(string)
		client , err := redis.Dial("tcp", beego.AppConfig.String("hostaddr")+":6379")
		if err!= nil{
			beego.Error(err)
			goto label
		}
		defer client.Close()
		reply, err := client.Do("lrem", "history_"+name, 0, this.GetString("goodsId"))
		beego.Error(reply, err)
		reply, err =client.Do("lpush","history_"+name , this.GetString("goodsId"))
		beego.Error(reply, err)
	}
	label:
	this.Data["good"] = good
	this.Data["goods"] = goods
	this.Layout = "listshow.html"
	this.TplName = "detail.html"

}


func(this *GoodsController)ShowList(){
	username := this.GetSession("userName")
	if username == nil{
		beego.Error("uername is empty")
		this.Data["userName"] = ""
	}else{
		this.Data["userName"] = username
	}
	typeId, err := this.GetInt("typeId")
	if err!= nil{
		beego.Error(err)
		return
	}
	pageSize:=5
	var goods []*models.GoodsSKU
	o := orm.NewOrm()
	qs := o.QueryTable("GoodsSKU").Filter("GoodsType__id", typeId)
	total, _ := qs.Count()
	pag := pagination.SetPaginator(this.Ctx, pageSize, total)
	
	beego.Info(pag.Offset(), pag.PerPageNums)
	qs = qs.Limit(pag.PerPageNums, pag.Offset())
	qs.All(&goods)

	this.Data["goods"] = goods
	var recommends []*models.GoodsSKU     
	_ ,err = o.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType__Id", typeId).OrderBy("Time", "desc").Limit(2).All(&recommends)  
	if  err!= nil{
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	var goodsType []*models.GoodsType
	o.QueryTable("GoodsType").All(&goodsType)
	this.Data["goodsType"] = goodsType

	this.Data["recommends"] = recommends
	this.Layout = "listshow.html"
	this.TplName = "list.html"

}