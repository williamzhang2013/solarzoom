package ucmd

import (
	"fmt"
)

type Sum struct {
	data []float64
	unit float64
}

func NewSum(data []float64, unit float64) *Sum {
	return &Sum{data, unit}
}

func (p *Sum) Run() float64 {
	var sum float64 = 0.0
	for _, v := range p.data {
		sum += v
	}

	fmt.Println("Sum command run!")
	return sum * p.unit
}
