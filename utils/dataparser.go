package utils

import (
	"errors"
	"fmt"
	"solarzoom/utils/simplejson"
	"solarzoom/utils/ucmd"
	// "os"
	// "strconv"
	//"math"
)

type CalcUnit struct {
	Cmd       string        // calc unit name ---
	SubCmd    []string      // sub command, in items array
	SubResult []interface{} // every sub command should be a return value
	Digital   float64       // digital
	Oper      string        // function name --- some function should be special treatment
	Default   interface{}   // default value
	//Value   interface{}
}

///////////////////////////////////////////////////////////////////////////////
const FItem_Function string = "Function"
const FItem_Items string = "Items"
const FItem_Digit string = "Digit"
const FItem_Dividend string = "Dividend"
const FItem_Divisor string = "Divisor"
const FItem_H string = "H"
const FItem_L string = "L"
const FItem_N string = "N"
const FItem_Default string = "Default"

const Cmd_WorkStatus string = "WorkStatus"
const Cmd_RunTimeTotal string = "RunTimeTotal"
const Cmd_EnergyTotal string = "EnergyTotal"
const Cmd_EnergyDay string = "EnergyDay"
const Cmd_ITemp string = "InternalTemperature"
const Cmd_VdcPV1 string = "VdcPV1"
const Cmd_IdcPV1 string = "IdcPV1"
const Cmd_DCPowerPV1 string = "DCPowerPV1"
const Cmd_VdcPV2 string = "VdcPV2"
const Cmd_IdcPV2 string = "IdcPV2"
const Cmd_DCPowerPV2 string = "DCPowerPV2"
const Cmd_VdcPV3 string = "VdcPV3"
const Cmd_IdcPV3 string = "IdcPV3"
const Cmd_DCPowerPV3 string = "DCPowerPV3"
const Cmd_VdcPV4 string = "VdcPV4"
const Cmd_IdcPV4 string = "IdcPV4"
const Cmd_DCPowerPV4 string = "DCPowerPV4"
const Cmd_GFCIResistorPV1 string = "GFCIResistorPV1"
const Cmd_GFCIResistorPV2 string = "GFCIResistorPV2"
const Cmd_GFCIResistorPV3 string = "GFCIResistorPV3"
const Cmd_GFCIResistorPV4 string = "GFCIResistorPV4"
const Cmd_AverVdcPV string = "AverVdcPV"
const Cmd_IdcTotal string = "IdcTotal"
const Cmd_DCPowerTotal string = "DCPowerTotal"
const Cmd_VacR string = "VacR"
const Cmd_IacR string = "IacR"
const Cmd_ACPwerR string = "ACPwerR"
const Cmd_FacR string = "FacR"
const Cmd_VacS string = "VacS"
const Cmd_IacS string = "IacS"
const Cmd_ACPwerS string = "ACPwerS"
const Cmd_FacS string = "FacS"
const Cmd_VacT string = "VacT"
const Cmd_IacT string = "IacT"
const Cmd_ACPwerT string = "ACPwerT"
const Cmd_FacT string = "FacT"
const Cmd_AverVac string = "AverVac"
const Cmd_ACActivePowerTotal string = "ACActivePowerTotal"
const Cmd_IacTotal string = "IacTotal"
const Cmd_VacBalance string = "VacBalance"
const Cmd_IacBalance string = "IacBalance"
const Cmd_Fgrid string = "Fgrid"
const Cmd_Efficiency string = "Efficiency"
const Cmd_SPLPEnergy string = "SPLPEnergy"
const Cmd_ErrorMessage string = "ErrorMessage"

///////////////////////////////////////////////////////////////////////////////
var dataCalcDB map[string]*CalcUnit

///////////////////////////////////////////////////////////////////////////////
func init() {
	// generate the CalcUnit map
	dataCalcDB = make(map[string]*CalcUnit, 50)

	dataCalcDB[Cmd_WorkStatus] = NewCalcUnit(Cmd_WorkStatus)
	dataCalcDB[Cmd_RunTimeTotal] = NewCalcUnit(Cmd_RunTimeTotal)
	dataCalcDB[Cmd_EnergyTotal] = NewCalcUnit(Cmd_EnergyTotal)
	dataCalcDB[Cmd_EnergyDay] = NewCalcUnit(Cmd_EnergyDay)
	dataCalcDB[Cmd_ITemp] = NewCalcUnit(Cmd_ITemp)
	dataCalcDB[Cmd_VdcPV1] = NewCalcUnit(Cmd_VdcPV1)
	dataCalcDB[Cmd_IdcPV1] = NewCalcUnit(Cmd_IdcPV1)
	dataCalcDB[Cmd_DCPowerPV1] = NewCalcUnit(Cmd_DCPowerPV1)
	dataCalcDB[Cmd_VdcPV2] = NewCalcUnit(Cmd_VdcPV2)
	dataCalcDB[Cmd_IdcPV2] = NewCalcUnit(Cmd_IdcPV2)
	dataCalcDB[Cmd_DCPowerPV2] = NewCalcUnit(Cmd_DCPowerPV2)
	dataCalcDB[Cmd_VdcPV3] = NewCalcUnit(Cmd_VdcPV3)
	dataCalcDB[Cmd_IdcPV3] = NewCalcUnit(Cmd_IdcPV3)
	dataCalcDB[Cmd_DCPowerPV3] = NewCalcUnit(Cmd_DCPowerPV3)
	dataCalcDB[Cmd_VdcPV4] = NewCalcUnit(Cmd_VdcPV4)
	dataCalcDB[Cmd_IdcPV4] = NewCalcUnit(Cmd_IdcPV4)
	dataCalcDB[Cmd_DCPowerPV4] = NewCalcUnit(Cmd_DCPowerPV4)
	dataCalcDB[Cmd_GFCIResistorPV1] = NewCalcUnit(Cmd_GFCIResistorPV1)
	dataCalcDB[Cmd_GFCIResistorPV2] = NewCalcUnit(Cmd_GFCIResistorPV2)
	dataCalcDB[Cmd_GFCIResistorPV3] = NewCalcUnit(Cmd_GFCIResistorPV3)
	dataCalcDB[Cmd_GFCIResistorPV4] = NewCalcUnit(Cmd_GFCIResistorPV4)
	dataCalcDB[Cmd_AverVdcPV] = NewCalcUnit(Cmd_AverVdcPV)
	dataCalcDB[Cmd_IdcTotal] = NewCalcUnit(Cmd_IdcTotal)
	dataCalcDB[Cmd_DCPowerTotal] = NewCalcUnit(Cmd_DCPowerTotal)
	dataCalcDB[Cmd_VacR] = NewCalcUnit(Cmd_VacR)
	dataCalcDB[Cmd_IacR] = NewCalcUnit(Cmd_IacR)
	dataCalcDB[Cmd_ACPwerR] = NewCalcUnit(Cmd_ACPwerR)
	dataCalcDB[Cmd_FacR] = NewCalcUnit(Cmd_FacR)
	dataCalcDB[Cmd_VacS] = NewCalcUnit(Cmd_VacS)
	dataCalcDB[Cmd_IacS] = NewCalcUnit(Cmd_IacS)
	dataCalcDB[Cmd_ACPwerS] = NewCalcUnit(Cmd_ACPwerS)
	dataCalcDB[Cmd_FacS] = NewCalcUnit(Cmd_FacS)
	dataCalcDB[Cmd_VacT] = NewCalcUnit(Cmd_VacT)
	dataCalcDB[Cmd_IacT] = NewCalcUnit(Cmd_IacT)
	dataCalcDB[Cmd_ACPwerT] = NewCalcUnit(Cmd_ACPwerT)
	dataCalcDB[Cmd_FacT] = NewCalcUnit(Cmd_FacT)
	dataCalcDB[Cmd_AverVac] = NewCalcUnit(Cmd_AverVac)
	dataCalcDB[Cmd_ACActivePowerTotal] = NewCalcUnit(Cmd_ACActivePowerTotal)
	dataCalcDB[Cmd_IacTotal] = NewCalcUnit(Cmd_IacTotal)
	dataCalcDB[Cmd_VacBalance] = NewCalcUnit(Cmd_VacBalance)
	dataCalcDB[Cmd_IacBalance] = NewCalcUnit(Cmd_IacBalance)
	dataCalcDB[Cmd_Fgrid] = NewCalcUnit(Cmd_Fgrid)
	dataCalcDB[Cmd_Efficiency] = NewCalcUnit(Cmd_Efficiency)
	dataCalcDB[Cmd_SPLPEnergy] = NewCalcUnit(Cmd_SPLPEnergy)
	dataCalcDB[Cmd_ErrorMessage] = NewCalcUnit(Cmd_ErrorMessage)
}

func NewCalcUnit(cmd string) *CalcUnit {
	return &CalcUnit{Cmd: cmd}
}

func RefreshDataCalcDB() {
	//
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func genSubCmdUnit(v *simplejson.Json, data *CalcUnit) {
	// generate the items & digital
	switch data.Oper {
	case ucmd.FNAME_DIV:
		item, _ := v.Get(FItem_Dividend).String()
		data.SubCmd = append(data.SubCmd, item)

		item, _ = v.Get(FItem_Divisor).String()
		data.SubCmd = append(data.SubCmd, item)
	case ucmd.FNAME_HL:
		item, _ := v.Get(FItem_H).String()
		data.SubCmd = append(data.SubCmd, item)

		item, _ = v.Get(FItem_L).String()
		data.SubCmd = append(data.SubCmd, item)

		item, _ = v.Get(FItem_N).String()
		data.SubCmd = append(data.SubCmd, item)
	case ucmd.FNAME_GDEF:
		var err error
		if data.Default, err = v.Get(FItem_Default).String(); err != nil {
			if data.Default, err = v.Get(FItem_Default).Float64(); err != nil {
				data.Default, _ = v.Get(FItem_Default).Int64()
			}
		}
		//fmt.Println("default value:", data.Default)
	default:
		// normal sub command
		data.SubCmd, _ = v.Get(FItem_Items).StringArray()
	}
	data.Digital = v.Get(FItem_Digit).MustFloat64()
}

func genCalcPara(oper string, orig []interface{}) interface{} {
	switch oper {
	case ucmd.FNAME_HL:
		length := len(orig)
		para := make([]uint64, length)
		for i, v := range orig {
			para[i], _ = v.(uint64)
		}
		return para
	case ucmd.FNAME_SUM:
		fallthrough
	case ucmd.FNAME_AVG:
		fallthrough
	case ucmd.FNAME_DIV:
		fallthrough
	case ucmd.FNAME_MUL:
		fallthrough
	case ucmd.FNAME_STDEV:
		length := len(orig)
		para := make([]float64, length)
		for i, v := range orig {
			para[i], _ = v.(float64)
		}
		return para
	case ucmd.FNAME_ISEQUAL:
		length := len(orig)
		para := make([]interface{}, length)
		for i, v := range orig {
			para[i] = v
		}
		return para
	case ucmd.FNAME_NVSTRCAT:
		length := len(orig)
		para := make([]string, length)
		for i, v := range orig {
			para[i], _ = v.(string)
		}
		return para
	}

	return nil
}
func reInitCalcUnit(unit *CalcUnit) {
	unit.SubCmd = nil
	unit.SubResult = nil
	unit.Digital = 0.0
	unit.Default = nil
}

func DoRunCalcFunc(fname string, unit *CalcUnit, value *simplejson.Json, dataMap map[string]interface{}) (interface{}, error) {
	unit.Oper, _ = value.Get(FItem_Function).String()
	reInitCalcUnit(unit)
	genSubCmdUnit(value, unit)
	//fmt.Printf("function unit:%v\n", unit)

	//traverse the sub command
	for _, subCmd := range unit.SubCmd {
		subUnit := NewCalcUnit(subCmd)

		subVal, _ := DoRunCalcUnit(fname, subUnit, dataMap)
		unit.SubResult = append(unit.SubResult, subVal)
	}

	// calc the result
	var cmdPara interface{}
	switch unit.Oper {
	case ucmd.FNAME_GDEF:
		cmdPara = unit.Default

		//cmdPara0, _ := cmdPara.(int64)
		//cmdPara1, _ := cmdPara.(float64)
		//fmt.Printf("cmdPara=%v, cmdParaf=%v\n", cmdPara0, cmdPara1)
	default:
		cmdPara = genCalcPara(unit.Oper, unit.SubResult)
		//fmt.Println("cmdPara=", cmdPara)
	}
	//fmt.Printf("Oper=%v, SubResult=%v, Digital=%v, cmdPara=%v\n", unit.Oper, unit.SubResult, unit.Digital, cmdPara)
	result := ucmd.Run(unit.Oper, cmdPara, unit.Digital)
	//fmt.Println("result=", result)
	return result, nil
}

// The function used calculate the input unit's value
func DoRunCalcUnit(fname string, unit *CalcUnit, dataMap map[string]interface{}) (interface{}, error) {
	data, ok := dataMap[unit.Cmd]
	if ok {
		// find the value
		//fmt.Println("1 ------ in value data: data=", data)
		return data, nil
	} else {
		// can't find the value
		if value, err := HandleJSONCmd(fname, unit.Cmd); err != nil {
			// no this command!
			//fmt.Println("NO this command")
			//fmt.Println("2 ------ No this command")
			return nil, err
		} else {
			// parse sub command, traverse all items
			//fmt.Println("3 ------ parse this command")
			return DoRunCalcFunc(fname, unit, value, dataMap)
		}
	}

	return nil, nil
}

func RunCalcUnit(fname, cmd string, dataMap map[string]interface{}) (interface{}, error) {
	unit, ok := dataCalcDB[cmd]
	if ok {
		return DoRunCalcUnit(fname, unit, dataMap)
	} else {
		return nil, errors.New("No this calculator unit!")
	}
}

func GenerateRealTimeObject() {
	//fmt.Printf("dataCalcDB=%v\n", dataCalcDB)
	for k, v := range dataCalcDB {
		fmt.Printf("k=%s, cmd=%s\n", k, v.Cmd)
	}
}
