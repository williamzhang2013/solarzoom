package ucmd

import (
//"fmt"
)

type Multiply struct {
	data  []float64
	digit float64
}

func NewMultiply(data []float64, digit float64) *Multiply {
	return &Multiply{data, digit}
}

func (p *Multiply) Run() float64 {
	var mul float64 = 1.0
	//fmt.Printf("Multiply=%v\n", p)
	for _, v := range p.data {
		mul *= v
	}

	//fmt.Println("Multiply command run!")
	return mul * p.digit
}
