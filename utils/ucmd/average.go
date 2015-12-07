package ucmd

import (
	"fmt"
)

type Average struct {
	data []float64
	unit float64
}

func NewAverage(data []float64, unit float64) *Average {
	return &Average{data, unit}
}

func (p *Average) Run() float64 {
	var sum float64 = 0.0

	for _, v := range p.data {
		sum += v
	}

	fmt.Println("Average command run!")
	return float64(sum / float64(len(p.data)) * p.unit)
}
