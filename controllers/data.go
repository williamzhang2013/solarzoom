package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"solarzoom/models"
	"solarzoom/utils"
	//"solarzoom/utils/simplejson"
)

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
type DataController struct {
	beego.Controller
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
const jsonPath = "./static/json/"

var sData string = "5630312e30322e534254524e4759535030303154523030314e4f563030303030303150434c313330305231353831323030360001000012270001aa550001010011823200db00000ca4000000010000000000000000000000000000000000000000000200000000ffff00000000000000000000020005542471"

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func setEnergyDay(m *models.PvInverterRunData, dataMap map[string]interface{}) {
	if energy, ok := dataMap["EnergyDay"].(float64); ok {
		m.EnergyDay = energy
	}
}

func setSampleTime(m *models.PvInverterRunData, dataMap map[string]interface{}) {
	if time, ok := dataMap["SmplTime"].(uint64); ok {
		m.SmplTime = int64(time)
	}
}

func setInterTemperature(m *models.PvInverterRunData, dataMap map[string]interface{}) {
	if t, ok := dataMap["InternalTemperature"].(float64); ok {
		m.InternalTemperature = t
	}
}

func setVdcPV1(m *models.PvInverterRunData, dataMap map[string]interface{}) {
	if v, ok := dataMap["VdcPV1"].(float64); ok {
		m.VdcPv1 = v
	}
}

func setVdcPV2(m *models.PvInverterRunData, dataMap map[string]interface{}) {
	if v, ok := dataMap["VdcPV2"].(float64); ok {
		m.VdcPv2 = v
	}
}

func setIdcPV1(m *models.PvInverterRunData, dataMap map[string]interface{}) {
	if v, ok := dataMap["IdcPV1"].(float64); ok {
		m.IdcPv1 = v
	}
}

func setIdcPV2(m *models.PvInverterRunData, dataMap map[string]interface{}) {
	if v, ok := dataMap["IdcPV2"].(float64); ok {
		m.IdcPv2 = v
	}
}

func setIacR(m *models.PvInverterRunData, dataMap map[string]interface{}) {
	if v, ok := dataMap["IacR"].(float64); ok {
		m.IacR = v
	}
}

func setVacR(m *models.PvInverterRunData, dataMap map[string]interface{}) {
	if v, ok := dataMap["VacR"].(float64); ok {
		m.VacR = v
	}
}

func setFacR(m *models.PvInverterRunData, dataMap map[string]interface{}) {
	if v, ok := dataMap["FacR"].(float64); ok {
		m.FacR = v
	}
}

func setACPwerR(m *models.PvInverterRunData, dataMap map[string]interface{}) {
	if v, ok := dataMap["ACPwerR"].(float64); ok {
		m.AcpowerR = v
	}
}

///////////////////////////////////////////////////////////////////////////////
func getInverterSN(dataMap map[string]interface{}) string {
	if sn, ok := dataMap["DVCSN_Len"].(string); ok {
		return sn
	}
	return ""
}

func getGWSN(dataMap map[string]interface{}) string {
	if sn, ok := dataMap["GWSN"].(string); ok {
		return sn
	}
	return ""
}

///////////////////////////////////////////////////////////////////////////////
func handleDataRequest(ctrl *DataController) {
	data := ctrl.GetString("data")
	fmt.Println("data=", data)

	//var s []byte = []byte(sData)
	var s []byte = []byte(data)

	stylecode := utils.PeekStyleCode(s)
	fmt.Printf("stylecode=%v\n", stylecode)
	fname := jsonPath + "SD" + stylecode[1] + stylecode[2] + ".json"
	_, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println("ReadJSONFile:", err.Error())
	} else {
		fmt.Println("ReadJSONFile SUCCESS!")
	}

	// fmt.Println("styleVersion=", utils.PeekStyleVersion(s))
	// fmt.Println("stylecode=", stylecode)
	// fmt.Println("resultTableName=", utils.PeekRstTblName(s))
	item := models.NewPvInverterRunData()
	// parse the map
	dataMap := utils.HandleSDData(fname, s)
	setSampleTime(item, dataMap)
	setInterTemperature(item, dataMap)
	setVdcPV1(item, dataMap)
	setVdcPV2(item, dataMap)
	setIdcPV1(item, dataMap)
	setIdcPV2(item, dataMap)
	setVacR(item, dataMap)
	setIacR(item, dataMap)
	setFacR(item, dataMap)
	setEnergyDay(item, dataMap)
	setACPwerR(item, dataMap)

	// fmt.Printf("myMap=%v\n", myMap)
	// time := myMap["SmplTime"]
	// tmpTime, ok := time.(uint64)
	// if ok {
	// 	item.SmplTime = int64(tmpTime)
	// }
	// get the inverter sn
	sn := getInverterSN(dataMap)
	if len(sn) != 0 {
		item.IvtId, _ = models.GetIvtIdByIvtSN(sn)
	}

	_, err = models.AddInverterRunData(item)
	if err != nil {
		beego.Error("write database User error!")
	}

	// infoid, _ := ctrl.GetInt32("infoid")
	// ivtsn := ctrl.GetString("ivtsn")
	// gwsn := ctrl.GetString("gwsn")
	// ivtaddr := ctrl.GetString("ivtaddr")

	// fmt.Printf("infoid=%d, ivtsn=%v, gwsn=%v, ivtaddr=%v\n", infoid, ivtsn, gwsn, ivtaddr)
	// _, err := models.AddGwIVTItem(infoid, ivtsn, gwsn, ivtaddr)
	// if err != nil {
	// 	beego.Error("write database PvCollectorInverter error!")
	// }

	// test the ORM CRUD operate
	// _, err0 := models.AddInverterInfo("SN0004", "SH", "DESCRIPTION_YG", "0.0.1", 1456.54)

	// if err0 != nil {
	// 	beego.Error("write database InverterInfo error!")
	// }

	// _, err := models.AddUser(1, 2, "noovo", "john.yin", "shanghai", "william.zhang@noovo.co", "13817503955", "vip")
	// if err != nil {
	// 	beego.Error("write database User error!")
	// }

	//models.ReadInverterInfoById(1)
	//models.UpdateInvertInfoById(4)
	//models.DeleteInvertInfoById(4)
	ctrl.Data["command1"] = "cmd"
	ctrl.Data["value1"] = "data"
	ctrl.Data["command2"] = "errcode"
	ctrl.Data["value2"] = 0
	ctrl.TplNames = "cmd2.tpl"
}

func (ctrl *DataController) Get() {
	//sess := ctrl.StartSession()
	//state := sess.Get(SessAuth)

	state := utils.GetSolarMapItem(utils.SessAuth)
	if state != "ok" {
		ctrl.Redirect(URLAuth, 302)
	} else {
		handleDataRequest(ctrl)
	}
	//handleDataRequest(ctrl)
}

func (ctrl *DataController) Post() {
	//sess := ctrl.StartSession()
	//state := sess.Get(SessAuth)

	state := utils.GetSolarMapItem(utils.SessAuth)
	if state != "ok" {
		ctrl.Redirect(URLAuth, 302)
	} else {
		handleDataRequest(ctrl)
	}
}

func (ctrl *DataController) Command() {
	// sess := ctrl.StartSession()
	// state := sess.Get(SessAuth)
	state := utils.GetSolarMapItem(utils.SessAuth)
	if state != "ok" {
		ctrl.Redirect(URLAuth, 302)
	} else {
		fmt.Println("Command")
		ctrl.Data["command"] = "Command"
		ctrl.Data["value"] = "Send back command"
		ctrl.TplNames = "cmd.tpl"
	}
}
