package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"solarzoom/utils"
	"time"
)

type ConfigController struct {
	beego.Controller
}

func handleConfigRequest(ctrl *ConfigController) {
	filename := ctrl.GetString("getstylefile")
	fmt.Println("style file name=", filename)

	if len(filename) == 0 {
		ctrl.Data["value"] = 1
	} else {
		ctrl.Data["value"] = 0
	}
	ctrl.Data["command"] = "errcode"
	ctrl.TplNames = "cmd.tpl"
}

func (ctrl *ConfigController) Get() {
	fmt.Println("config controller get method")

	// sess := ctrl.StartSession()
	// state := sess.Get(SessAuth)
	state := utils.GetSolarMapItem(utils.SessAuth)
	//state = "ok"
	//fmt.Println("request=", ctrl.Ctx.Request.URL)
	//fmt.Println("auth state=", state)
	if state != "ok" {
		// redirect auth
		ctrl.Redirect(URLAuth, 302)
	} else {
		handleConfigRequest(ctrl)
		// filename := c.GetString("getstylefile")
		// fmt.Println("style file name=", filename)

		// c.Data["Website"] = c.GetSession("user")
		// c.Data["Email"] = c.Ctx.GetCookie("cookie_user")
		// c.TplNames = "index.tpl"
	}
}

func (ctrl *ConfigController) Post() {
	fmt.Println("config controller post method")

	// sess := ctrl.StartSession()
	// state := sess.Get(SessAuth)
	state := utils.GetSolarMapItem(utils.SessAuth)
	fmt.Println("auth state=", state)

	if state != "ok" {
		ctrl.Redirect(URLAuth, 302)
	} else {
		handleConfigRequest(ctrl)
		// filename := c.GetString("getstylefile")
		// fmt.Println("style file name=", filename)

		// c.Data["Website"] = c.GetSession("user")
		// c.Data["Email"] = c.Ctx.GetCookie("cookie_user")
		// c.TplNames = "index.tpl"
	}
}

func (c *ConfigController) SyncTime() {
	//sess := c.StartSession()
	//state := sess.Get(SessAuth)
	state := utils.GetSolarMapItem(utils.SessAuth)
	fmt.Println("auth state=", state)

	if state != "ok" {
		c.Redirect(URLAuth, 302)
	} else {
		fmt.Println("SyncTime")

		curtime := time.Now().Unix()
		c.Data["command1"] = "cmd"
		c.Data["value1"] = "synctime"
		c.Data["command2"] = "time"
		c.Data["value2"] = uint32(curtime)
		c.TplNames = "cmd2.tpl"
	}
}

func (c *ConfigController) GetSingleIVTTable() {
	// sess := c.StartSession()
	// state := sess.Get(SessAuth)
	state := utils.GetSolarMapItem(utils.SessAuth)
	fmt.Println("auth state=", state)

	if state != "ok" {
		c.Redirect(URLAuth, 302)
	} else {
		fmt.Println("GetSingleIVTTable")

		c.TplNames = "cmd.tpl"
	}
}

func (c *ConfigController) GetHybridIVTTable() {
	// sess := c.StartSession()
	// state := sess.Get(SessAuth)
	state := utils.GetSolarMapItem(utils.SessAuth)
	fmt.Println("auth state=", state)

	if state != "ok" {
		c.Redirect(URLAuth, 302)
	} else {
		fmt.Println("GetHybridIVTTable")

		c.TplNames = "cmd.tpl"
	}
}
