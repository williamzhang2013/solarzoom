package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"runtime"
	"solarzoom/database"
	_ "solarzoom/routers"
)

func main() {
	fmt.Println("main function")
	runtime.GOMAXPROCS(runtime.NumCPU())

	o := orm.NewOrm()
	o.Using("default")
	database.CreateTable()

	beego.SessionOn = true
	beego.SessionProvider = "memory"
	beego.SessionGCMaxLifetime = 3600 //60 seconds
	beego.SessionName = "session_solarzoom"
	beego.SessionCookieLifeTime = 3600 //60 seconds
	beego.SessionAutoSetCookie = true
	beego.SessionSavePath = "/"

	beego.Run()
}
