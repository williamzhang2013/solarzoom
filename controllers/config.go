package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
	"solarzoom/utils"
	"time"
)

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
const FILE_PREFIX string = "./static/json/"
const FILE_JSON string = ".json"
const IVT_CONFIG_JSON_FILE string = "./static/json/STNOOVODNL01.json"

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
type ConfigController struct {
	beego.Controller
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func handleConfigRequest(ctrl *ConfigController) {
	filename := ctrl.GetString("getstylefile")
	fmt.Println("style file name=", filename)

	if len(filename) == 0 {
		ctrl.Data["value"] = 1
		ctrl.Data["command"] = "errcode"
		ctrl.TplNames = "cmd.tpl"
	} else if utils.IsFileExist(FILE_PREFIX + filename + FILE_JSON) {
		// find the file
		http.ServeFile(ctrl.Ctx.ResponseWriter, ctrl.Ctx.Request, FILE_PREFIX+filename+FILE_JSON)
		utils.WriteDebugLog("Send back the device style file: %s", FILE_PREFIX+filename+FILE_JSON)
		//ctrl.Data["value"] = 0
	} else {
		ctrl.Data["value"] = 2
		ctrl.Data["command"] = "errcode"
		ctrl.TplNames = "cmd.tpl"
	}
}

func (ctrl *ConfigController) Get() {
	fmt.Println("config controller get method")

	sess := ctrl.StartSession()
	state := sess.Get(utils.SessAuth)
	//state := utils.GetSolarMapItem(utils.SessAuth)
	//state = "ok"

	if state != "ok" {
		// redirect auth
		utils.WriteDebugLog("Config: Redirect AUTH")
		ctrl.Redirect(URLAuth, 302)
	} else {
		handleConfigRequest(ctrl)
	}
}

func (ctrl *ConfigController) Post() {
	fmt.Println("config controller post method")

	sess := ctrl.StartSession()
	state := sess.Get(utils.SessAuth)
	//state := utils.GetSolarMapItem(utils.SessAuth)
	//fmt.Println("auth state=", state)

	if state != "ok" {
		utils.WriteDebugLog("Config: Redirect AUTH")
		ctrl.Redirect(URLAuth, 302)
	} else {
		handleConfigRequest(ctrl)
	}
}

///////////////////////////////////////////////////////////////////////////////
func (ctrl *ConfigController) SyncTime() {
	sess := ctrl.StartSession()
	state := sess.Get(utils.SessAuth)
	//state := utils.GetSolarMapItem(utils.SessAuth)
	fmt.Println("auth state=", state)

	if state != "ok" {
		utils.WriteDebugLog("Config: Redirect AUTH")
		ctrl.Redirect(URLAuth, 302)
	} else {
		fmt.Println("SyncTime")
		utils.WriteDebugLog("SyncTime")

		curtime := time.Now().Unix()
		ctrl.Data["command1"] = "cmd"
		ctrl.Data["value1"] = "synctime"
		ctrl.Data["command2"] = "time"
		ctrl.Data["value2"] = uint32(curtime)
		ctrl.TplNames = "cmd2.tpl"
	}
}

func (ctrl *ConfigController) GetIVTTable() {
	sess := ctrl.StartSession()
	state := sess.Get(utils.SessAuth)

	//state := utils.GetSolarMapItem(utils.SessAuth)
	//state := "ok"
	//fmt.Println("auth state=", state)

	if state != "ok" {
		utils.WriteDebugLog("Config: Redirect AUTH")
		ctrl.Redirect(URLAuth, 302)
	} else {
		fmt.Println("Get Inverter Config File!")

		utils.WriteDebugLog("Get Inverter Table: %s", IVT_CONFIG_JSON_FILE)
		http.ServeFile(ctrl.Ctx.ResponseWriter, ctrl.Ctx.Request, IVT_CONFIG_JSON_FILE)

		ctrl.Data["value"] = 0
		ctrl.Data["command"] = "errcode"
		ctrl.TplNames = "cmd.tpl"
	}
}
