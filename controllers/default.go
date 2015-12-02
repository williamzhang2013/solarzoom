package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	fmt.Println("Input:", c.Input())

	cmd := c.GetString("synctime")
	if cmd == "" {
		fmt.Println("NOT Receive SyncTime command!")
		c.Data["time"] = uint32(123456)
	} else {
		fmt.Println("Receive SyncTime command!")
		curtime := time.Now().Unix()
		c.Data["time"] = uint32(curtime)
	}
	c.TplNames = "cmd.tpl"
}

func (c *MainController) Post() {
	fmt.Println("Input:", c.Input())

	if v, err := c.GetInt("Data"); err == nil {
		fmt.Println("DataInt=", v)
	}

	fmt.Println("DataString=", c.GetString("Data"))

	// v := c.Ctx.Input.GetData("Data")
	// fmt.Println("Input Data=", v)
	// fmt.Println("RequestBody=", c.Ctx.Input.RequestBody)
	cmd := c.GetString("synctime")
	if cmd == "" {
		fmt.Println("NOT Receive SyncTime command!")
		c.Data["time"] = uint32(123456)
	} else {
		fmt.Println("Receive SyncTime command!")
		curtime := time.Now().Unix()
		c.Data["time"] = uint32(curtime)
		//s := "Content-type:application/json\n\n {\"time\":\"OK\"}"
	}
	c.TplNames = "cmd.tpl"
}
