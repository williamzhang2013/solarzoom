package ucmd

import (
	"fmt"
)

type Multiply struct {
	data []float64
	unit float64
}

func NewMultiply(data []float64, unit float64) *Multiply {
	return &Multiply{data, unit}
}

func (p *Multiply) Run() float64 {
	var mul float64 = 1.0
	for _, v := range p.data {
		mul *= v
	}

	fmt.Println("Multiply command run!")
	return mul * p.unit
}
