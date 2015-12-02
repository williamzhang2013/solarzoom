package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type DownloadController struct {
	beego.Controller
}

func (c *DownloadController) Get() {
	c.Data["Website"] = "beego.me"

	c.TplNames = "login.html"

	file := c.GetString("file")
	c.Data["Page"] = fmt.Sprintf("Download file is: %s", file)
	//fmt.Println("Need download %s file!", file)
}
