package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	// "io"
	// "log"
	// "strings"
)

// type History struct {
// 	Year, Day, Title, Detail string
// }

// const jsonStream = `
// 		{"Year":"2015","Day":"11.11 13:30","Title":"创建项目","Detail":"采用beego的框架+ORM操作数据库"}
// 		{"Year":"2015","Day":"11.11 14:30","Title":"增加路由","Detail":"添加路由控制部分"}
//    `

type AboutController struct {
	beego.Controller
}

func (c *AboutController) Get() {
	fmt.Println("About page")
	c.Data["Page"] = "About"
	c.TplNames = "about.html"
}
