package ucmd

import (
	//"fmt"
	"math"
)

type Average struct {
	data  []float64
	digit float64
}

func NewAverage(data []float64, digit float64) *Average {
	return &Average{data, digit}
}

func (p *Average) Run() interface{} {
	var sum float64 = 0.0
	var length float64 = 0.0

	for _, v := range p.data {
		sum += v
		if math.Abs(v) >= 0.000001 {
			// v == 0
			length += 1.0
		}
	}

	//fmt.Println("Average without 0 command run!")
	return float64(sum / length * p.digit)
	//return float64(sum / float64(len(p.data)) * p.digit)
}
