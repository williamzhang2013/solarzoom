package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type ConfigController struct {
	beego.Controller
}

func (c *ConfigController) Get() {
	fmt.Println("config controller get method")

	sess := c.StartSession()
	state := sess.Get(SessAuth)
	fmt.Println("auth state=", state)
	if state != "ok" {
		// redirect auth
		c.Redirect(URLAuth, 302)
	} else {
		filename := c.GetString("getstylefile")
		fmt.Println("style file name=", filename)

		c.Data["Website"] = c.GetSession("user")
		c.Data["Email"] = c.Ctx.GetCookie("cookie_user")
		c.TplNames = "index.tpl"
	}
}

func (c *ConfigController) Post() {
	fmt.Println("config controller post method")

	sess := c.StartSession()
	state := sess.Get(SessAuth)
	fmt.Println("auth state=", state)

	if state != "ok" {
		c.Redirect(URLAuth, 302)
	} else {
		filename := c.GetString("getstylefile")
		fmt.Println("style file name=", filename)

		c.Data["Website"] = c.GetSession("user")
		c.Data["Email"] = c.Ctx.GetCookie("cookie_user")
		c.TplNames = "index.tpl"
	}
}

func (c *ConfigController) SyncTime() {
	sess := c.StartSession()
	state := sess.Get(SessAuth)
	fmt.Println("auth state=", state)

	if state != "ok" {
		c.Redirect(URLAuth, 302)
	} else {
		fmt.Println("SyncTime")

		curtime := time.Now().Unix()
		c.Data["command"] = "time"
		c.Data["value"] = uint32(curtime)
		c.TplNames = "cmd.tpl"
	}
}

func (c *ConfigController) GetSingleIVTTable() {
	sess := c.StartSession()
	state := sess.Get(SessAuth)
	fmt.Println("auth state=", state)

	if state != "ok" {
		c.Redirect(URLAuth, 302)
	} else {
		fmt.Println("GetSingleIVTTable")

		c.TplNames = "cmd.tpl"
	}
}

func (c *ConfigController) GetHybridIVTTable() {
	sess := c.StartSession()
	state := sess.Get(SessAuth)
	fmt.Println("auth state=", state)

	if state != "ok" {
		c.Redirect(URLAuth, 302)
	} else {
		fmt.Println("GetHybridIVTTable")

		c.TplNames = "cmd.tpl"
	}
}
