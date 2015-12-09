package utils

import (
	"fmt"
)

// Record the constant command --- with constant response length
// Record the var need to save to Database
//////////////////////////////////////////////////////
//////////////////////////////////////////////////////
//////////////////////////////////////////////////////
const cmdSliceMax int = 50
const cmdLogicInit int = 10
const cmdRealInit int = 40

const lenStyleVersion int = 8
const lenStyleCode1 int = 2
const lenStyleCode2 int = 5
const lenStyleCode3 int = 5
const lenStyleCode int = 12
const lenGWSN int = 10 //18
const lenDAdrr int = 2 // 1
const lenCmdRsltTblName int = 5
const lenBatchOrder int = 2 // 3
const lenSampleTime int = 4
const lenCmdFail int = 1
const lenCRCL int = 1
const lenCRCH int = 1

//////////////////////////////////////////////////////
//////////////////////////////////////////////////////
//////////////////////////////////////////////////////
type SysDataHeader struct {
	StyleCode    []string // style code, [3]string
	StyleVersion string   // style version
	CreateDate   string   // SD file create date
}

type LogicCmdRspType struct {
	command     string // command
	description string // command description
	length      int    // response data length
	mode        int    // 0 --- string, 1 --- number
}

type RealCmdRspType struct {
	command     string // command
	description string // command description
	mode        int    // 0 --- string, 1 --- number
}

//////////////////////////////////////////////////////
//////////////////////////////////////////////////////
//////////////////////////////////////////////////////
func GetLogicCmdArray() []LogicCmdRspType {
	var LogicCmd []LogicCmdRspType = make([]LogicCmdRspType, cmdLogicInit, cmdSliceMax)

	LogicCmd[0] = LogicCmdRspType{"StyleVersion", "设备版本号", lenStyleVersion, 0}
	LogicCmd[1] = LogicCmdRspType{"StyleCode", "设备代码", lenStyleCode, 0}
	LogicCmd[2] = LogicCmdRspType{"CmdRsltTblName", "返回数据表名字", lenCmdRsltTblName, 0}
	LogicCmd[3] = LogicCmdRspType{"GWSN", "数据采集器SN号", lenGWSN, 0}
	LogicCmd[4] = LogicCmdRspType{"DAddr", "设备地址", lenDAdrr, 1}
	LogicCmd[5] = LogicCmdRspType{"SmplTime", "上次采样时间", lenSampleTime, 1}
	LogicCmd[6] = LogicCmdRspType{"BatchOrder", "批处理命令运行次数", lenBatchOrder, 1}
	LogicCmd[7] = LogicCmdRspType{"CRCL", "CRCL", lenCRCL, 1}
	LogicCmd[8] = LogicCmdRspType{"CRCH", "CRCH", lenCRCH, 1}

	return LogicCmd
}

func GetRealCmdArray() []RealCmdRspType {
	var realCmd []RealCmdRspType = make([]RealCmdRspType, cmdRealInit, cmdSliceMax)

	realCmd[0] = RealCmdRspType{"DVCSN", "逆变器SN", 0}
	realCmd[1] = RealCmdRspType{"VdcPV1", "PV1直流电压", 1}
	realCmd[2] = RealCmdRspType{"IdcPV1", "PV1直流电流", 1}
	realCmd[3] = RealCmdRspType{"DCPowerPV1", "PV1直流功率", 1}
	realCmd[4] = RealCmdRspType{"VdcPV2", "PV2直流电压", 1}
	realCmd[5] = RealCmdRspType{"IdcPV2", "PV2直流电流", 1}
	realCmd[6] = RealCmdRspType{"DCPowerPV2", "PV2直流功率", 1}
	realCmd[7] = RealCmdRspType{"VdcPV3", "PV3直流电压", 1}
	realCmd[8] = RealCmdRspType{"IdcPV3", "PV3直流电流", 1}
	realCmd[9] = RealCmdRspType{"DCPowerPV3", "PV3直流功率", 1}
	realCmd[10] = RealCmdRspType{"VdcPV4", "PV4直流电压", 1}
	realCmd[11] = RealCmdRspType{"IdcPV4", "PV4直流电流", 1}
	realCmd[12] = RealCmdRspType{"DCPowerPV4", "PV4直流功率", 1}
	realCmd[13] = RealCmdRspType{"GFCIResistorPV1", "PV1正对地绝缘阻抗", 1}
	realCmd[14] = RealCmdRspType{"GFCIResistorPV2", "PV2正对地绝缘阻抗", 1}
	realCmd[15] = RealCmdRspType{"GFCIResistorPV3", "PV3正对地绝缘阻抗", 1}
	realCmd[16] = RealCmdRspType{"GFCIResistorPV4", "PV4正对地绝缘阻抗", 1}
	realCmd[17] = RealCmdRspType{"VacR", "R相电压", 1}
	realCmd[18] = RealCmdRspType{"IacR", "R相电流", 1}
	realCmd[19] = RealCmdRspType{"FacR", "R相频率", 1}
	realCmd[20] = RealCmdRspType{"ACPowerR", "R相输出功率", 1}
	realCmd[21] = RealCmdRspType{"VacS", "S相电压", 1}
	realCmd[22] = RealCmdRspType{"IacS", "S相电流", 1}
	realCmd[23] = RealCmdRspType{"FacS", "S相频率", 1}
	realCmd[24] = RealCmdRspType{"ACPowerS", "S相输出功率", 1}
	realCmd[25] = RealCmdRspType{"VacT", "T相电压", 1}
	realCmd[26] = RealCmdRspType{"IacT", "T相电流", 1}
	realCmd[27] = RealCmdRspType{"FacT", "T相频率", 1}
	realCmd[28] = RealCmdRspType{"ACPowerT", "T相输出功率", 1}
	realCmd[29] = RealCmdRspType{"Fgrid", "电网频率", 1}
	realCmd[30] = RealCmdRspType{"DCPowerTotal", "总直流功率", 1}
	realCmd[31] = RealCmdRspType{"ACActivePowerTotal", "总有功功率", 1}
	realCmd[32] = RealCmdRspType{"RunTimeTotal", "累计运行时间", 1}
	realCmd[33] = RealCmdRspType{"EnergyTotal", "累计总发电量", 1}
	realCmd[34] = RealCmdRspType{"EnergyDay", "当日总发电量", 1}

	return realCmd
}

//////////////////////////////////////////////////////
func GetLogicCmdItem(cmd string, list []LogicCmdRspType) LogicCmdRspType {
	var noneCmd LogicCmdRspType
	for _, v := range list {
		if v.command == cmd {
			return v
		}
	}
	return noneCmd
}

func GetRealCmdItem(cmd string, list []RealCmdRspType) RealCmdRspType {
	var noneCmd RealCmdRspType
	for _, v := range list {
		if v.command == cmd {
			return v
		}
	}
	return noneCmd
}

//////////////////////////////////////////////////////
func printLogicCmdArray(list []LogicCmdRspType) {
	for _, v := range list {
		if v.command != "" {
			fmt.Println("command=", v.command,
				", description=", v.description,
				", length=", v.length)
		}
	}
}

func printReadCmdArray(list []RealCmdRspType) {
	for _, v := range list {
		if v.command != "" {
			fmt.Println("command=", v.command, ", description=", v.description)
		}
	}
}
