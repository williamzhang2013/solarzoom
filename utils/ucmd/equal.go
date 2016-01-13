package ucmd

import (
//"fmt"
)

type Equal2 struct {
	data0 int64
	data1 int64
	ok    interface{}
	fail  interface{}
}

func NewEqual2(data0, data1 int64, ok, fail interface{}) *Equal2 {
	return &Equal2{data0, data1, ok, fail}
}

func (p *Equal2) Run() interface{} {
	//var equal bool = false
	//fmt.Printf("NewEqual2:data0=%v, data1=%v, ok=%v, fail=%v\n", p.data0, p.data1, p.ok, p.fail)
	if p.data0 == p.data1 {
		return p.ok
	} else {
		return p.fail
	}
}
