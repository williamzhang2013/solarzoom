package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
	"solarzoom/models"
	"solarzoom/utils"
	"time"
)

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
const FILE_STYLE_PATH string = "./static/json/"
const FILE_FW_PATH string = "./static/fw/"
const FILE_DEF_PATH string = "./static/"

const FILE_JSON string = ".json"
const IVT_CONFIG_JSON_FILE string = "./static/json/STNOOVODNL01.json"

// ATTENTION!!!!:
// the command CMD_GET_FW_FILE && CMD_GET_STYLE_FILE can't be replaced using by
// CMD_GET_FILE, because the default path is different
const CMD_GET_FW_FILE string = "getfwfile"
const CMD_GET_STYLE_FILE string = "getstylefile"
const CMD_GET_FILE string = "getfile"

const CMD_REPORT_FWVER string = "fmver"
const CMD_REPORT_GWSN string = "gwsn"
const CMD_REPORT_IPADDR string = "ipaddr"

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
type ConfigController struct {
	beego.Controller
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func handleConfigGWInfo(gwsn, fwver, ipaddr string) {
	if infoItem, ok := models.GWInfoMap[gwsn]; ok {
		infoItem.FWVersion = fwver
		infoItem.IPAddr = ipaddr
	} else {
		infoItem = models.NewGWInfo(fwver, ipaddr)
		models.GWInfoMap[gwsn] = infoItem
	}
	//fmt.Println("models.GWInfoMap[gwsn]=", models.GWInfoMap[gwsn].FWVersion)
}

func handleConfigRequest(ctrl *ConfigController) {
	stylename := ctrl.GetString(CMD_GET_STYLE_FILE)
	fwname := ctrl.GetString(CMD_GET_FW_FILE)
	filename := ctrl.GetString(CMD_GET_FILE)

	gwsn := ctrl.GetString(CMD_REPORT_GWSN)
	fwver := ctrl.GetString(CMD_REPORT_FWVER)
	ipaddr := ctrl.GetString(CMD_REPORT_IPADDR)

	ctrl.Data["value"] = 0
	ctrl.Data["command"] = "errcode"

	// first, handle the gwsn & fwver parameters
	// these two parameters MUST come same time
	if gwsn != "" && (fwver != "" || ipaddr != "") {
		handleConfigGWInfo(gwsn, fwver, ipaddr)
	} else if gwsn == "" && fwver == "" && ipaddr == "" {
		// do nothing
	} else {
		ctrl.Data["value"] = 2
	}

	// second, handle style name
	if len(stylename) != 0 {
		if utils.IsFileExist(FILE_STYLE_PATH + stylename + FILE_JSON) {
			http.ServeFile(ctrl.Ctx.ResponseWriter, ctrl.Ctx.Request, FILE_STYLE_PATH+stylename+FILE_JSON)
			utils.WriteDebugLog("Send back the device style file: %s", FILE_STYLE_PATH+stylename+FILE_JSON)
		} else {
			ctrl.Data["value"] = 4
		}
	}

	// third , handle fireware name
	if len(fwname) != 0 {
		if utils.IsFileExist(FILE_FW_PATH + fwname) {
			http.ServeFile(ctrl.Ctx.ResponseWriter, ctrl.Ctx.Request, FILE_FW_PATH+fwname)
			utils.WriteDebugLog("Send back the firmware image: %s", FILE_FW_PATH+fwname)
		} else {
			ctrl.Data["value"] = 4
		}
	}

	// last, handle file name
	if len(filename) != 0 {
		if utils.IsFileExist(FILE_DEF_PATH + filename) {
			http.ServeFile(ctrl.Ctx.ResponseWriter, ctrl.Ctx.Request, FILE_DEF_PATH+filename)
			utils.WriteDebugLog("Send back the file: %s", FILE_DEF_PATH+filename)
		} else {
			ctrl.Data["value"] = 4
		}
	}

	ctrl.TplNames = "cmd.tpl"

	// fmt.Printf("stylefile=%v, firmware=%v, file=%v\n", stylename, fwname, filename)

	// if len(stylename) == 0 {
	// 	ctrl.Data["value"] = 1
	// 	ctrl.Data["command"] = "errcode"
	// 	ctrl.TplNames = "cmd.tpl"
	// } else if utils.IsFileExist(FILE_STYLE_PATH + stylename + FILE_JSON) {
	// 	// find the file
	// 	http.ServeFile(ctrl.Ctx.ResponseWriter, ctrl.Ctx.Request, FILE_STYLE_PATH+stylename+FILE_JSON)
	// 	utils.WriteDebugLog("Send back the device style file: %s", FILE_STYLE_PATH+stylename+FILE_JSON)
	// 	//ctrl.Data["value"] = 0
	// } else {
	// 	ctrl.Data["value"] = 2
	// 	ctrl.Data["command"] = "errcode"
	// 	ctrl.TplNames = "cmd.tpl"
	// }
}

func (ctrl *ConfigController) Get() {
	fmt.Println("config controller get method")

	sess := ctrl.StartSession()
	state := sess.Get(utils.SessAuth)
	//state := utils.GetSolarMapItem(utils.SessAuth)
	//state = "ok"

	if state != "ok" {
		// redirect auth
		utils.WriteDebugLog("Config: AUTH ERROR!")
		//ctrl.Redirect(URLAuth, 302)
		ctrl.Data["command1"] = "cmd"
		ctrl.Data["value1"] = "data"
		ctrl.Data["command2"] = "errcode"
		ctrl.Data["value2"] = 3
		ctrl.TplNames = "cmd2.tpl"
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
		utils.WriteDebugLog("Config: AUTH ERROR!")
		//ctrl.Redirect(URLAuth, 302)
		ctrl.Data["command1"] = "cmd"
		ctrl.Data["value1"] = "data"
		ctrl.Data["command2"] = "errcode"
		ctrl.Data["value2"] = 3
		ctrl.TplNames = "cmd2.tpl"
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
		utils.WriteDebugLog("Config: AUTH ERROR!")
		//ctrl.Redirect(URLAuth, 302)
		ctrl.Data["command1"] = "cmd"
		ctrl.Data["value1"] = "data"
		ctrl.Data["command2"] = "errcode"
		ctrl.Data["value2"] = 3
		ctrl.TplNames = "cmd2.tpl"
	} else {
		fmt.Println("SyncTime")
		utils.WriteDebugLog("SyncTime: mac=%s", ctrl.GetString("mac"))

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
		utils.WriteDebugLog("Config: AUTH ERROR!")
		//ctrl.Redirect(URLAuth, 302)
		ctrl.Data["command1"] = "cmd"
		ctrl.Data["value1"] = "data"
		ctrl.Data["command2"] = "errcode"
		ctrl.Data["value2"] = 3
		ctrl.TplNames = "cmd2.tpl"
	} else {
		fmt.Println("Get Inverter Config File!")

		utils.WriteDebugLog("Get Inverter Table: %s", IVT_CONFIG_JSON_FILE)
		http.ServeFile(ctrl.Ctx.ResponseWriter, ctrl.Ctx.Request, IVT_CONFIG_JSON_FILE)

		ctrl.Data["value"] = 0
		ctrl.Data["command"] = "errcode"
		ctrl.TplNames = "cmd.tpl"
	}
}
