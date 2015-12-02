package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"solarzoom/models"
)

type DataController struct {
	beego.Controller
}

func (c *DataController) Get() {
	sess := c.StartSession()
	state := sess.Get(SessAuth)

	if state != "ok" {
		c.Redirect(URLAuth, 302)
	} else {
		data := c.GetString("data")
		fmt.Println("get data:", data)
		c.Data["command"] = "data"
		c.Data["value"] = data
		c.TplNames = "cmd.tpl"

		infoid, _ := c.GetInt32("infoid")
		ivtsn := c.GetString("ivtsn")
		gwsn := c.GetString("gwsn")
		ivtaddr := c.GetString("ivtaddr")

		fmt.Printf("infoid=%d, ivtsn=%v, gwsn=%v, ivtaddr=%v\n", infoid, ivtsn, gwsn, ivtaddr)
		_, err := models.AddGwIVTItem(infoid, ivtsn, gwsn, ivtaddr)
		if err != nil {
			beego.Error("write database PvCollectorInverter error!")
		}
	}

	// test the ORM CRUD operate
	// _, err0 := models.AddInverterInfo("SN0004", "SH", "DESCRIPTION_YG", "0.0.1", 1456.54)

	// if err0 != nil {
	// 	beego.Error("write database InverterInfo error!")
	// }

	// _, err := models.AddUser(1, 2, "noovo", "john.yin", "shanghai", "william.zhang@noovo.co", "13817503955", "vip")
	// if err != nil {
	// 	beego.Error("write database User error!")
	// }

	//models.ReadInverterInfoById(1)
	//models.UpdateInvertInfoById(4)
	//models.DeleteInvertInfoById(4)
}

func (c *DataController) Post() {
	sess := c.StartSession()
	state := sess.Get(SessAuth)
	if state != "ok" {
		c.Redirect(URLAuth, 302)
	} else {
		data := c.GetString("data")
		fmt.Println("get data:", data)

		c.Data["command"] = "data"
		c.Data["value"] = data
		c.TplNames = "cmd.tpl"
	}
}

func (c *DataController) Command() {
	sess := c.StartSession()
	state := sess.Get(SessAuth)
	if state != "ok" {
		c.Redirect(URLAuth, 302)
	} else {
		fmt.Println("Command")
		c.Data["command"] = "Command"
		c.Data["value"] = "Send back command"
		c.TplNames = "cmd.tpl"
	}
}
