package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "ttsx/routers"
	_ "ttsx/models"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()

}

