package ucmd

import (
//"fmt"
)

type NVStrCat struct {
	data []string
}

func NewStrCat(data []string) *NVStrCat {
	return &NVStrCat{data}
}

func (p *NVStrCat) Run() interface{} {
	var s string = ""
	for _, v := range p.data {
		s += v
	}

	//fmt.Println("Sum command run!")
	return s
}
