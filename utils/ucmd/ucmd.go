package ucmd

import (
	"fmt"
)

type UCmd interface {
	Run() float64
}

const FNAME_SUM string = "Sum"
const FNAME_AVG string = "Average" // averge without 0
const FNAME_DIV string = "Division"
const FNAME_MUL string = "Multiply"
const FNAME_HL string = "H_L"
const FNAME_STDEV string = "STDEV"
const FNAME_AVG0 string = "Averge0" // averge include 0

func Run(source string, data interface{}, digit float64) float64 {
	var cmd UCmd

	//fmt.Println("UCmd run! command =", source)
	switch source {
	case FNAME_HL:
		para, _ := data.([]uint64)
		cmd = &Merge{para[0], para[1], 8, digit}
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
	default:
		fmt.Println("Unsupported Command!")
		return 0.0
	}

	return cmd.Run()
}
