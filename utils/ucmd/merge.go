package ucmd

import (
	"fmt"
)

type Merge struct {
	h    uint64 // high
	l    uint64 // low
	n    uint8  // move bytes
	unit float64
}

func NewMerge(h, l uint64, n uint8, unit float64) *Merge {
	return &Merge{h, l, n, unit}
}

func (p *Merge) Run() float64 {
	v := (p.h << (p.n * 8)) | p.l

	fmt.Println("Merge command run!")
	return float64(float64(v) * p.unit)
}
