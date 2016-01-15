package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
	"solarzoom/utils"
)

type UploadController struct {
	beego.Controller
}

func (c *UploadController) Get() {
	ip := utils.GetLocalIP()

	if supportHttps, _ := beego.AppConfig.Bool("EnableHttpTLS"); supportHttps {
		port, _ := beego.AppConfig.Int("HttpsPort")
		c.Data["protocol"] = "https"
		c.Data["port"] = port
	} else {
		port, _ := beego.AppConfig.Int("httpport")
		c.Data["protocol"] = "http"
		c.Data["port"] = port
	}

	c.Data["localIP"] = ip
	c.TplNames = "device_update.html"
}

func (u *UploadController) Post() {
	fmt.Println("upload post method")
	r := u.Ctx.Request
	w := u.Ctx.ResponseWriter
	r.ParseMultipartForm(32 << 30)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}
