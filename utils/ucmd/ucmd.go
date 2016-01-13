package ucmd

import (
	"fmt"
)

type UCmd interface {
	Run() interface{}
}

const FNAME_SUM string = "Sum"
const FNAME_AVG string = "Average" // averge without 0
const FNAME_DIV string = "Division"
const FNAME_MUL string = "Multiply"
const FNAME_HL string = "H_L"
const FNAME_STDEV string = "STDEV"
const FNAME_AVG0 string = "Averge0" // averge include 0
const FNAME_GDEF string = "GetDefault"
const FNAME_ISEQUAL string = "Equal"
const FNAME_NVSTRCAT string = "NVStrCat"

func Run(source string, data interface{}, digit float64) interface{} {
	var cmd UCmd

	//fmt.Println("UCmd run! command =", source)
	switch source {
	case FNAME_HL:
		para, _ := data.([]uint64)
		cmd = &Merge{para[0], para[1], para[2], digit}
	case FNAME_SUM:
		para, _ := data.([]float64)
		cmd = &Sum{para, digit}
	case FNAME_AVG:
		para, _ := data.([]float64)
		cmd = &Average{para, digit}
	case FNAME_AVG0:
		para, _ := data.([]float64)
		cmd = &Average0{para, digit}
	case FNAME_MUL:
		para, _ := data.([]float64)
		cmd = &Multiply{para, digit}
	case FNAME_DIV:
		para, _ := data.([]float64)
		cmd = &Division{para[0], para[1], digit}
	case FNAME_STDEV:
		para, _ := data.([]float64)
		cmd = &Stdev{para, digit}
	case FNAME_GDEF:
		cmd = &GetDefault{data}
	case FNAME_ISEQUAL:
		var para2 interface{}
		var para3 interface{}
		var ok bool
		para, _ := data.([]interface{})
		para0, _ := para[0].(float64)
		para1, _ := para[1].(float64)

		para2, ok = para[2].(string)
		if !ok {
			para2, _ = para[2].(float64)
		}

		para3, ok = para[3].(string)
		if !ok {
			para3, _ = para[3].(float64)
		}

		//para2, _ := para[2].(string)
		//para3, _ := para[3].(string)

		//fmt.Printf("Run equal cmd: para0=%v, para1=%v, para2=%v, para3=%v\n", para0, para1, para2, para3)
		//fmt.Println("data value=", data)
		cmd = &Equal2{int64(para0), int64(para1), para2, para3}
	case FNAME_NVSTRCAT:
		para, _ := data.([]string)
		cmd = &NVStrCat{para}
	default:
		fmt.Println("Unsupported Command!")
		return 0.0
	}

	return cmd.Run()
}
