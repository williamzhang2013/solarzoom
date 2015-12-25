package ucmd

import (
//"fmt"
)

type Average0 struct {
	data  []float64
	digit float64
}

func NewAverage0(data []float64, digit float64) *Average0 {
	return &Average0{data, digit}
}

func (p *Average0) Run() float64 {
	var sum float64 = 0.0

	for _, v := range p.data {
		sum += v
	}

	//fmt.Println("Average with 0 command run!")
	//return float64(sum / length * p.digit)
	return float64(sum / float64(len(p.data)) * p.digit)
}
