package ucmd

import (
	//"fmt"
	"math"
)

type Stdev struct {
	data  []float64
	digit float64
}

func NewStdev(data []float64, digit float64) *Stdev {
	return &Stdev{data, digit}
}

func (p *Stdev) Run() interface{} {
	var sum float64 = 0.0
	var avg float64 = 0.0

	for _, v := range p.data {
		sum += v
	}
	avg = sum / float64(len(p.data))

	sum = 0.0
	for _, v := range p.data {
		sum += math.Pow((v - avg), 2)
	}
	avg = sum / float64(len(p.data))

	//fmt.Println("Stdev command run!")
	return math.Sqrt(avg) * float64(p.digit)
}
