package ucmd

import (
//"fmt"
)

type Sum struct {
	data  []float64
	digit float64
}

func NewSum(data []float64, digit float64) *Sum {
	return &Sum{data, digit}
}

func (p *Sum) Run() float64 {
	var sum float64 = 0.0
	for _, v := range p.data {
		sum += v
	}

	//fmt.Println("Sum command run!")
	return sum * p.digit
}
