package ucmd

import (
	"fmt"
)

type UCmd interface {
	Run() float64
}

func Run(source string, data interface{}) float64 {
	var cmd UCmd

	fmt.Println("UCmd run! command =", source)
	switch source {
	case "H_L":
		para, _ := data.(*Merge)
		cmd = &Merge{para.h, para.l, para.n, para.unit}
	case "Sum":
		para, _ := data.(*Sum)
		cmd = &Sum{para.data, para.unit}
	case "Average":
		para, _ := data.(*Average)
		cmd = &Average{para.data, para.unit}
	case "Multiply":
		para, _ := data.(*Multiply)
		cmd = &Multiply{para.data, para.unit}
	case "Division":
		para, _ := data.(*Division)
		cmd = &Division{para.dividend, para.divisor, para.unit}
	case "STDEV":
		para, _ := data.(*Stdev)
		cmd = &Stdev{para.data, para.unit}
	default:
		fmt.Println("Unsupported Command!")
		return 0.0
	}

	return cmd.Run()
}
