package ucmd

import (
//"fmt"
)

type Division struct {
	dividend float64
	divisor  float64
	digit    float64
}

func NewDivision(dividend, divisor, digit float64) *Division {
	return &Division{dividend, divisor, digit}
}

func (p *Division) Run() float64 {
	//fmt.Printf("Division command run! dividend=%f, divisor=%d\n", p.dividend, p.divisor)
	return p.dividend / p.divisor * p.digit
}
