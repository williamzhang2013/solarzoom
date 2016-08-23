package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"solarzoom/models"
	"solarzoom/utils"
	"strconv"
	"strings"
	"time"
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

const STR_FAULT string = "Fault"
const STR_OFF string = "Off"

// TRNGY Data
//var sData string = "5630312e30332e33534254524e4759535030303154523030314e4f5630303030303031303030303030303050434c3133303052313538313230303100025697138c01aa5500010100118232011700a6058d0000003d0000002108a3138902f10000000001d400000092000100000000ffff0000000000000000000000000914b0e9"
//var sData string = "5630312e30332e34534254524e4759535030303154523030314e4f5630303030303031303030303030303050434c3133303052313538313230303500a456d91f4b03aa55000301001182320141013404b60000002a0000001308ca138e01ba00000000023900000072000100000000ffff00000000000000000000000008106ba10000"
var sData string = "5630312e30332e34534254524e47595350303031545230303131323a63633a31383a65663a66663a65643a50434c31333030523135383132303036000157ba953c01aa5500010100118232012600000cab000000000000000000000000000000000000000000000000000200000000ffff00000000000000000000020004a6a028"

// SUN
//var sData string = "5630312e30332e36534253554e47524b544c30314b544c303130303a34643a31303a30363a31383a30343a4131363034323733353035000000000000000000007f57bbb7d7020204e6000103b438870000033f0000020fffffffff180200af180c00af180600afa8510000099009700950025002510253000000000000000000000000a4f200000061000003e801f303d90000000000000000000000000000000000000000000000fa00000000000000000000000000000000000000000000000000000000000000000000000000000000000001d9ffff0000000000000000d6d800000000000000010002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000013c000e17f800b0e57f88b5"

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func setBatchOrder(m *models.PvInverterRunData, dataMap map[string]interface{}) {
	if order, ok := dataMap["BatchOrder"].(uint64); ok {
		m.BatchOrder = int32(order)
	}
}

func setSampleTime(m *models.PvInverterRunData, dataMap map[string]interface{}) {
	if sTime, ok := dataMap["SmplTime"].(uint64); ok {
		m.SmplTime = int64(sTime)

		hour, min, _ := time.Unix(m.SmplTime, 0).Clock()
		m.BatchOrder = int32(hour*12 + min/5 + 1)
		//fmt.Printf("SampleTime ---> BatchOrder:hour=%d, min=%d, batchOrder=%d\n", hour, min, m.BatchOrder)
	}
}

///////////////////////////////////////////////////////////////////////////////
func setWorkStatus(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	var workStatus = "Normal"

	if v, err := utils.RunCalcUnit(fname, utils.Cmd_WorkStatus, dataMap); err == nil {
		if s, ok := v.(string); ok {
			m.WorkStatus = s
			if b := strings.Contains(s, STR_FAULT); !b {
				if b = strings.Contains(s, STR_OFF); !b {
					m.WorkStatus = workStatus
				}
			}
		} else {
			m.WorkStatus = workStatus
		}
	}
}

func setRunTimeTotal(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_RunTimeTotal, dataMap); err == nil {
		if time, ok := v.(float64); ok {
			m.RunTimeTotal = time
		}
	}
}

func setEnergyTotal(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_EnergyTotal, dataMap); err == nil {
		if energy, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", energy)
			m.EnergyTotal, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setEnergyDay(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_EnergyDay, dataMap); err == nil {
		if energy, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", energy)
			m.EnergyDay, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setInterTemperature(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_ITemp, dataMap); err == nil {
		if t, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", t)
			m.InternalTemperature, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setVdcPV1(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_VdcPV1, dataMap); err == nil {
		if vdc, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", vdc)
			m.VdcPv1, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setIdcPV1(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_IdcPV1, dataMap); err == nil {
		if idc, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", idc)
			m.IdcPv1, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setDCPowerPV1(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_DCPowerPV1, dataMap); err == nil {
		if power, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", power)
			m.DcpowerPv1, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setVdcPV2(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_VdcPV2, dataMap); err == nil {
		if vdc, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", vdc)
			m.VdcPv2, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setIdcPV2(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_IdcPV2, dataMap); err == nil {
		if idc, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", idc)
			m.IdcPv2, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setDCPowerPV2(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_DCPowerPV2, dataMap); err == nil {
		if power, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", power)
			m.DcpowerPv2, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setVdcPV3(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_VdcPV3, dataMap); err == nil {
		if vdc, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", vdc)
			m.VdcPv3, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setIdcPV3(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_IdcPV3, dataMap); err == nil {
		if idc, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", idc)
			m.IdcPv3, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setDCPowerPV3(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_DCPowerPV3, dataMap); err == nil {
		if power, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", power)
			m.DcpowerPv3, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setVdcPV4(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_VdcPV4, dataMap); err == nil {
		if vdc, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", vdc)
			m.VdcPv4, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setIdcPV4(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_IdcPV4, dataMap); err == nil {
		if idc, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", idc)
			m.IdcPv4, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setDCPowerPV4(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_DCPowerPV4, dataMap); err == nil {
		if power, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", power)
			m.DcpowerPv4, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setGFCIResistorPV1(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_GFCIResistorPV1, dataMap); err == nil {
		if r, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", r)
			m.Pv1Resistor, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setGFCIResistorPV2(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_GFCIResistorPV2, dataMap); err == nil {
		if r, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", r)
			m.Pv2Resistor, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setGFCIResistorPV3(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_GFCIResistorPV3, dataMap); err == nil {
		if r, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", r)
			m.Pv3Resistor, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setGFCIResistorPV4(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_GFCIResistorPV4, dataMap); err == nil {
		if r, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", r)
			m.Pv4Resistor, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setAverVdcPV(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_AverVdcPV, dataMap); err == nil {
		if avgVdc, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", avgVdc)
			m.AverVdcPv, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setIdcTotal(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_IdcTotal, dataMap); err == nil {
		if idcTotal, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", idcTotal)
			m.IdcTotal, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setDCPowerTotal(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_DCPowerTotal, dataMap); err == nil {
		if dcTotal, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", dcTotal)
			m.DcpowerTotal, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setVacR(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_VacR, dataMap); err == nil {
		if vac, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", vac)
			m.VacR, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setIacR(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_IacR, dataMap); err == nil {
		if iac, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", iac)
			m.IacR, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setACPwerR(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_ACPwerR, dataMap); err == nil {
		if power, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", power)
			m.AcpowerR, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setFacR(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_FacR, dataMap); err == nil {
		if f, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", f)
			m.FacR, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setVacS(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_VacS, dataMap); err == nil {
		if vac, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", vac)
			m.VacS, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setIacS(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_IacS, dataMap); err == nil {
		if iac, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", iac)
			m.IacS, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setACPwerS(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_ACPwerS, dataMap); err == nil {
		if power, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", power)
			m.AcpowerS, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setFacS(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_FacS, dataMap); err == nil {
		if f, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", f)
			m.FacS, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setVacT(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_VacT, dataMap); err == nil {
		if vac, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", vac)
			m.VacT, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setIacT(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_IacT, dataMap); err == nil {
		if iac, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", iac)
			m.IacT, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setACPwerT(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_ACPwerT, dataMap); err == nil {
		if power, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", power)
			m.AcpowerT, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setFacT(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_FacT, dataMap); err == nil {
		if f, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", f)
			m.FacT, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setAverVac(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_AverVac, dataMap); err == nil {
		if avgVac, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", avgVac)
			m.AverVac, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setACActivePowerTotal(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_ACActivePowerTotal, dataMap); err == nil {
		if power, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", power)
			m.AcActivePowerTotal, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setIacTotal(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_IacTotal, dataMap); err == nil {
		if iac, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", iac)
			//fmt.Printf("iac=%v, siac=%v\n", iac, s)
			m.IacTotal, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setVacBalance(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_VacBalance, dataMap); err == nil {
		if vBlc, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", vBlc)
			m.VacBalance, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setIacBalance(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_IacBalance, dataMap); err == nil {
		if iBlc, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", iBlc)
			m.IacBalance, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setFgrid(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_Fgrid, dataMap); err == nil {
		if grid, ok := v.(float64); ok {
			m.Fgrid = grid
		}
	}
}

func setEfficiency(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_Efficiency, dataMap); err == nil {
		if ef, ok := v.(float64); ok {
			s := fmt.Sprintf("%.1f", ef*100)
			m.Efficiency, _ = strconv.ParseFloat(s, 64)
		}
	}
}

func setSPLPEnergy(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_SPLPEnergy, dataMap); err == nil {
		if simu, ok := v.(float64); ok {
			s := fmt.Sprintf("%.02f", simu)
			//fmt.Printf("simu=%v, s=%s\n", simu, s)
			m.SimuKwh5Min, _ = strconv.ParseFloat(s, 64)
		}
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

func getErrorMessage(m *models.PvInverterRunData, fname string, dataMap map[string]interface{}) string {
	if v, err := utils.RunCalcUnit(fname, utils.Cmd_ErrorMessage, dataMap); err == nil {
		if msg, ok := v.(string); ok {
			//s := fmt.Sprintf("%.02f", simu)
			//fmt.Printf("simu=%v, s=%s\n", simu, s)
			return msg
		}
	}

	return ""
}

///////////////////////////////////////////////////////////////////////////////
func genIvtRunDataDBItem(item *models.PvInverterRunData, fname string, dataMap map[string]interface{}) {
	//setBatchOrder(item, dataMap)
	setSampleTime(item, dataMap)

	setWorkStatus(item, fname, dataMap)
	//fmt.Println("ErrorMessage:", getErrorMessage(item, fname, dataMap))
	setRunTimeTotal(item, fname, dataMap)
	setEnergyTotal(item, fname, dataMap)
	setEnergyDay(item, fname, dataMap)
	setInterTemperature(item, fname, dataMap)
	setVdcPV1(item, fname, dataMap)
	setIdcPV1(item, fname, dataMap)
	setDCPowerPV1(item, fname, dataMap)
	setVdcPV2(item, fname, dataMap)
	setIdcPV2(item, fname, dataMap)
	setDCPowerPV2(item, fname, dataMap)
	setVdcPV3(item, fname, dataMap)
	setIdcPV3(item, fname, dataMap)
	setDCPowerPV3(item, fname, dataMap)
	setVdcPV4(item, fname, dataMap)
	setIdcPV4(item, fname, dataMap)
	setDCPowerPV4(item, fname, dataMap)
	setGFCIResistorPV1(item, fname, dataMap)
	setGFCIResistorPV2(item, fname, dataMap)
	setGFCIResistorPV3(item, fname, dataMap)
	setGFCIResistorPV4(item, fname, dataMap)
	setAverVdcPV(item, fname, dataMap)
	setIdcTotal(item, fname, dataMap)
	setDCPowerTotal(item, fname, dataMap)
	setVacR(item, fname, dataMap)
	setIacR(item, fname, dataMap)
	setACPwerR(item, fname, dataMap)
	setFacR(item, fname, dataMap)
	setVacS(item, fname, dataMap)
	setIacS(item, fname, dataMap)
	setACPwerS(item, fname, dataMap)
	setFacS(item, fname, dataMap)
	setVacT(item, fname, dataMap)
	setIacT(item, fname, dataMap)
	setACPwerT(item, fname, dataMap)
	setFacT(item, fname, dataMap)
	setAverVac(item, fname, dataMap)
	setACActivePowerTotal(item, fname, dataMap)
	setIacTotal(item, fname, dataMap)
	setVacBalance(item, fname, dataMap)
	setIacBalance(item, fname, dataMap)
	setFgrid(item, fname, dataMap)
	setEfficiency(item, fname, dataMap)
	setSPLPEnergy(item, fname, dataMap)
}

func handleDataRequest(ctrl *DataController) {
	data := ctrl.GetString("data")

	// send back the json file
	ctrl.Data["command1"] = "cmd"
	ctrl.Data["value1"] = "data"
	ctrl.Data["command2"] = "errcode"
	ctrl.Data["value2"] = 2
	ctrl.TplNames = "cmd2.tpl"

	if data != "" {
		//var s []byte = []byte(sData)
		var s []byte = []byte(data)

		fmt.Println("data=", data)
		utils.WriteDebugLog("Handle Data: data=%v", data)
		stylecode := utils.PeekStyleCode(s)
		fmt.Printf("stylecode=%v\n", stylecode)
		fname := FILE_STYLE_PATH + "SD" + stylecode[1] + stylecode[2] + ".json"
		_, err := ioutil.ReadFile(fname)
		if err != nil {
			fmt.Println("ReadJSONFile:", err.Error())
		} else {
			fmt.Println("ReadJSONFile SUCCESS!")
		}

		item := models.NewPvInverterRunData()
		dataMap := utils.HandleSDData(fname, s)
		genIvtRunDataDBItem(item, fname, dataMap)

		// get the inverter sn
		sn := getInverterSN(dataMap)
		if len(sn) != 0 {
			item.IvtId, _ = models.GetIvtIdByIvtSN(sn)
			//fmt.Println("item.IvtId=", item.IvtId)
		}	
		utils.WriteDebugLog("Parse Data: ivt_id=%v, batch_order=%v, smpl_time=%v, ", item.IvtId, item.BatchOrder, item.SmplTime)
		utils.WriteDebugLog("            input_time=%v, work_status=%v, run_time_total=%v, ", item.InputTime, item.WorkStatus, item.RunTimeTotal)
		utils.WriteDebugLog("            energy_total=%v, energy_day=%v, internal_temperature=%v, ", item.EnergyTotal, item.EnergyDay, item.InternalTemperature)
		utils.WriteDebugLog("            vdc_pv1=%v, idc_pv1=%v, dcpower_pv1=%v, ", item.VdcPv1, item.IdcPv1, item.DcpowerPv1)
		utils.WriteDebugLog("            vdc_pv2=%v, idc_pv2=%v, dcpower_pv2=%v, ", item.VdcPv2, item.IdcPv2, item.DcpowerPv2)
		utils.WriteDebugLog("            vdc_pv3=%v, idc_pv3=%v, dcpower_pv3=%v, ", item.VdcPv3, item.IdcPv3, item.DcpowerPv3)
		utils.WriteDebugLog("            vdc_pv4=%v, idc_pv4=%v, dcpower_pv4=%v, ", item.VdcPv4, item.IdcPv4, item.DcpowerPv4)
		utils.WriteDebugLog("            pv1_resistor=%v, pv2_resistor=%v, pv3_resistor=%v, pv4_resistor=%v, ", item.Pv1Resistor, item.Pv2Resistor, item.Pv3Resistor, item.Pv4Resistor)
		utils.WriteDebugLog("            aver_vdc_pv=%v, idc_total=%v, dcpower_total=%v, ", item.AverVdcPv, item.IdcTotal, item.DcpowerTotal)
		utils.WriteDebugLog("            vac_r=%v, iac_r=%v, acpower_r=%v, fac_r=%v ", item.VacR, item.IacR, item.AcpowerR, item.FacR)
		utils.WriteDebugLog("            vac_s=%v, iac_s=%v, acpower_s=%v, fac_s=%v ", item.VacS, item.IacS, item.AcpowerS, item.FacS)
		utils.WriteDebugLog("            vac_t=%v, iac_t=%v, acpower_t=%v, fac_t=%v ", item.VacT, item.IacT, item.AcpowerT, item.FacT)
		utils.WriteDebugLog("            aver_vac=%v, ac_active_power_total=%v, iac_total=%v, ", item.AverVac, item.AcActivePowerTotal, item.IacTotal)
		utils.WriteDebugLog("            vac_balance=%v, iac_balance=%v, ", item.VacBalance, item.IacBalance)
		utils.WriteDebugLog("            fgrid=%v, efficiency=%v, simu_kwh5_min=%v", item.Fgrid, item.Efficiency, item.SimuKwh5Min)

		// calculate the day data & update it transfer to Solarzoom
		// so comment the code
		// dayRecord := models.NewPvInverterDayData()
		// dayRecord.IvtId = item.IvtId
		// dayRecord.Day = models.CalcDayTableDayItem(item.SmplTime)
		// dayRecord.AcActivePowerTotal = item.AcActivePowerTotal
		// dayRecord.EnergyTotal = item.EnergyTotal
		// dayRecord.EnergyToday = item.EnergyDay
		// dayRecord.PowerContent, _ = models.GetPowerContentInDayTable(dayRecord)
		// //fmt.Println("1 --- dayRecord.PowerContent=", dayRecord.PowerContent)

		// dayRecord.PowerContent = fmt.Sprintf("%s#%v:%v:%v", dayRecord.PowerContent, item.BatchOrder, item.SmplTime, item.DcpowerTotal)

		//fmt.Println("2 --- dayRecord.PowerContent=", dayRecord.PowerContent)
		// careate the new table
		// item.TableName()
		//models.CreateDayTableBySQL()
		//models.InsertDayTableItemBySQL()
		models.InsertRunDataTableItemBySQL(item)
		// models.UpdateDayTableRecordBySQL(dayRecord)

		//item.WorkStatus = "Error"
		// check the workstatus
		//fmt.Println("workstatus=", item.WorkStatus)
		if item.WorkStatus == STR_FAULT {
			// write the fault table
			fault := models.NewPvInverterFaultData()
			fault.IvtId = item.IvtId
			fault.FaultMessage = getErrorMessage(item, fname, dataMap)

			models.InsertFaultTableItemBySQL(fault)
		}

		ctrl.Data["value2"] = 0
	}
}

func (ctrl *DataController) Get() {
	sess := ctrl.StartSession()
	state := sess.Get(utils.SessAuth)

	//state := utils.GetSolarMapItem(utils.SessAuth)
	//state = "ok"
	utils.WriteDebugLog("/gw/data GET request")
	if state != "ok" {
		utils.WriteDebugLog("Data: AUTH ERROR!")
		//ctrl.Redirect(URLAuth, 302)
		ctrl.Data["command1"] = "cmd"
		ctrl.Data["value1"] = "data"
		ctrl.Data["command2"] = "errcode"
		ctrl.Data["value2"] = 3
		ctrl.TplNames = "cmd2.tpl"
	} else {
		handleDataRequest(ctrl)
	}
	//handleDataRequest(ctrl)
}

func (ctrl *DataController) Post() {
	sess := ctrl.StartSession()
	state := sess.Get(utils.SessAuth)

	//state := utils.GetSolarMapItem(utils.SessAuth)
	utils.WriteDebugLog("/gw/data POST request")
	if state != "ok" {
		utils.WriteDebugLog("Data: AUTH ERROR!")
		//ctrl.Redirect(URLAuth, 302)
		ctrl.Data["command1"] = "cmd"
		ctrl.Data["value1"] = "data"
		ctrl.Data["command2"] = "errcode"
		ctrl.Data["value2"] = 3
		ctrl.TplNames = "cmd2.tpl"
	} else {
		handleDataRequest(ctrl)
	}
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

func (ctrl *DataController) Command() {
	sess := ctrl.StartSession()
	state := sess.Get(utils.SessAuth)

	//state := utils.GetSolarMapItem(utils.SessAuth)
	gwsn := ctrl.GetString("cmd")
	//fmt.Println("gwsn=", gwsn)
	if state != "ok" {
		utils.WriteDebugLog("Data: AUTH ERROR!")
		//ctrl.Redirect(URLAuth, 302)
		ctrl.Data["command1"] = "cmd"
		ctrl.Data["value1"] = "data"
		ctrl.Data["command2"] = "errcode"
		ctrl.Data["value2"] = 3
		ctrl.TplNames = "cmd2.tpl"
	} else {
		fmt.Println("Command")
		//fmt.Println("content:=", models.SerialCommands(gwsn))

		utils.WriteDebugLog("/gw/cmd request")
		ctrl.Data["command"] = "commands"
		ctrl.Data["value"] = models.SerialCommands(gwsn)
		ctrl.TplNames = "cmd.tpl"
	}
}
