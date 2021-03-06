package ucmd

import (
//"fmt"
)

type Merge struct {
	h     uint64 // high
	l     uint64 // low
	n     uint64 // move bytes
	digit float64
}

func NewMerge(h, l uint64, n uint64, digit float64) *Merge {
	return &Merge{h, l, n, digit}
}

func (p *Merge) Run() interface{} {
	v := (p.h << (p.n * 8)) | p.l

	//fmt.Println("Merge command run!")
	return float64(float64(v) * p.digit)
}
