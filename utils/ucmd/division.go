package ucmd

import (
	"fmt"
)

type Division struct {
	dividend float64
	divisor  float64
	unit     float64
}

func NewDivision(dividend, divisor, unit float64) *Division {
	return &Division{dividend, divisor, unit}
}

func (p *Division) Run() float64 {

	fmt.Println("Multiply command run!")
	return p.dividend / p.divisor * p.unit
}
