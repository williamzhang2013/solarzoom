package models

import (
	"fmt"
	//"github.com/astaxie/beego"
)

type GWInfo struct {
	FWVersion string
	IPAddr    string
}

var GWInfoMap map[string]*GWInfo

func init() {
	GWInfoMap = make(map[string]*GWInfo)

	fmt.Println("init the gateway info map!")
}

func NewGWInfo(ver, ip string) *GWInfo {
	return &GWInfo{ver, ip}
}
