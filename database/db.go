package database

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"solarzoom/models"
	"time"
)

func init() {
	username := beego.AppConfig.String("username")
	password := beego.AppConfig.String("password")
	dbname := beego.AppConfig.String("dbname")

	fmt.Println("package database: init function")
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql",
		fmt.Sprintf("%s:%s@/%s?charset=utf8", username, password, dbname), 30)
}

func CreateTable() {
	name := "default"                          //数据库别名
	force := false                             //不强制建数据库
	verbose := true                            //打印建表过程
	err := orm.RunSyncdb(name, force, verbose) //建表
	if err != nil {
		beego.Error(err)
	}

	// create table by SQL
	//models.CreateDayTableBySQL()
	models.CreateRunDataTableBySQL(time.Now().Unix())
}
