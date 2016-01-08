package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"solarzoom/controllers"
)

func init() {
	fmt.Println("package routers: init function")
	beego.Router("/gw/auth", &controllers.AuthController{})
	beego.Router("/gw/config/", &controllers.ConfigController{})
	beego.Router("/gw/config/synctime", &controllers.ConfigController{}, "*:SyncTime")
	beego.Router("/gw/config/getivttable", &controllers.ConfigController{}, "*:GetIVTTable")
	beego.Router("/gw", &controllers.DataController{})
	beego.Router("/gw/cmd", &controllers.DataController{}, "*:Command")

	beego.Router("/", &controllers.MainController{})
	// beego.Router("/synctime", &controllers.MainController{}, "*:SyncTime")
	// beego.Router("/auth", &controllers.AuthController{})
	// beego.Router("/login", &controllers.LoginController{})
	// beego.Router("/login/judge", &controllers.LoginController{}, "*:Judge")
	beego.Router("/upload", &controllers.UploadController{})
	// beego.Router("/download", &controllers.DownloadController{})
	// beego.Router("/data", &controllers.DataController{})
	// beego.Router("/config", &controllers.ConfigController{})
	beego.Router("/about", &controllers.AboutController{})
}
