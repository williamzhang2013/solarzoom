package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	fmt.Println("Get  Method")
	c.Data["Website"] = "beego.me"
	c.Data["Page"] = "Login Page"
	c.TplNames = "cgi_login.html"

	//fmt.Printf("user=%s, password=%s\n", c.GetString("name"), c.GetString("password"))

	// myStruct := "{\"BUSY\":0, \"ERRCODE\":0}"
	// c.Data["json"] = &myStruct
	// c.ServeJson()
	// fmt.Println("myStruct=", myStruct)
}

func (c *LoginController) Post() {
	fmt.Println("Post Method")
	c.Data["Website"] = "beego.me"
	c.Data["Page"] = "Login Page"
	c.TplNames = "cgi_login.html"
}

func (c *LoginController) Judge() {
	fmt.Printf("user=%s, password=%s\n", c.GetString("name"), c.GetString("password"))
	myStruct := "{\"BUSY\":0, \"ERRCODE\":1}"
	//myStruct := {"BUSY":0, "ERRCODE":1}
	c.Data["json"] = &myStruct
	c.ServeJson()
	fmt.Println("myStruct=", myStruct)
}
