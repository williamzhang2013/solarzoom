package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"solarzoom/models"
	"solarzoom/utils"
	"strconv"
	"strings"
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
//var sData string = "5630312e30322e34534254524e4759535030303154523030314e4f563030303030303150434c313330305231353831323030320025000043370001aa55000101001182320083000005b90000000000000000089b138e000000020000000000000000000000000000ffff000000000000000000020200064fdeb2"
//var sData string = "5630312e30322e34534254524e4759535030303154523030314e4f5630303030303031303030303030303050434c31333030523135383132303032000a0000192901aa550001010011823200e2002204c200000005000000020883138e002b00000000000500000007000100000000ffff000000000000000000000000006f99151"
//var sData string = "5602300231022e02300233022e02310253024202540252024e024702590253025002300230023102540252023002300231024e024f0256023002300230023002300230023102500243024c023102330230023002520231023502380231023202300230023602002102560271023902f02102aa025502002102102002110282023202002d802002002d02402002002002102002002002002002002002002002002002002002002002002002002002002002202002002002002ff02ff02002002002002002002002002002002202002402b2026b02e102"
//var sData string = "5633102e30332e31534254524e4759535030303154523030314e4f5630303030303031303030303030303050434c3133303052313538313230303100035680d0c601aa5500010100118232011a012905570000002f0000001b088b138c026500000000001f00000005000100000000ffff000000000000000000000000066d23b5"
var sData string = "5630312e30332e33534254524e4759535030303154523030314e4f5630303030303031303030303030303050434c3133303052313538313230303100025697138c01aa5500010100118232011700a6058d0000003d0000002108a3138902f10000000001d400000092000100000000ffff0000000000000000000000000914b0e9"

// SUN
//var sData string = "5630312e30332e32534253554e47524b544c30314b544c30314e4f56303030303030313030303030303030000100000a3b010104e60001009500700000006c0000015c000000001cf50006ffffffff0000000001d7000008f208f408db0007000700070000000000000000000000000159000000000000000001f402dc0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000013300000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000219000200000000580e8f250000"
//var sData string = "5630312e30332e32534253554e47524b544c30314b544c30310104E60001009500700000006C0000015E000000001CCD0006FFFFFFFF0000000001F6000008EE08FA08D50008000800080000000000000000000000000178000000000000000001F402ED00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000134000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002160002000000003602"
//var sData string = "5630312e30332e32534253554e47524b544c30314b544c30314e4f56303030303030313030303030303030000100000a3b010104e60001009500700000006c0000015c000000001cf50006ffffffff0000000001d7000008f208f408db0007000700070000000000000000000000000159000000000000000001f402dc0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000013300000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000219000200000000580e8f25"

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func setBatchOrder(m *models.PvInverterRunData, dataMap map[string]interface{}) {
	if order, ok := dataMap["BatchOrder"].(uint64); ok {
		m.BatchOrder = int32(order)
	}
}

func setSampleTime(m *models.PvInverterRunData, dataMap map[string]interface{}) {
	if time, ok := dataMap["SmplTime"].(uint64); ok {
		m.SmplTime = int64(time)
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
	setBatchOrder(item, dataMap)
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
		}

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
