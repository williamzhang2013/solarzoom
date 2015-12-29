package ucmd

import (
//"fmt"
)

type GetDefault struct {
	digit interface{}
}

func NewGetDefault(digit interface{}) *GetDefault {
	return &GetDefault{digit}
}

func (p *GetDefault) Run() interface{} {
	//fmt.Println("Sum command run!")
	return p.digit
}
