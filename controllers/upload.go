package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
)

type UploadController struct {
	beego.Controller
}

func (c *UploadController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Page"] = "Upload Page"
	c.TplNames = "device_update.html"
	// curtime := time.Now().Unix()
	// h := md5.New()
	// io.WriteString(h, strconv.FormatInt(curtime, 10))
	// token := fmt.Sprintf("%x", h.Sum(nil))

	// t, _ := template.ParseFiles("upload.gtpl")
	// t.Execute(w, token)
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
