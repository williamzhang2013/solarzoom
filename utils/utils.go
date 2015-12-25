package utils

import (
	"fmt"
	"os"
	"strconv"
	//"math"
)

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
const SessAuth string = "auth"

///////////////////////////////////////////////////////////////////////////////
var solarMap map[string]string

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func init() {
	fmt.Println("package utils: init function!")
	initSolarMap()
}

func initSolarMap() {
	solarMap = make(map[string]string)
	solarMap[SessAuth] = "none"
}

func PrintSolarMap() {
	fmt.Println("solarMap=", solarMap)
}

func GetSolarMapItem(k string) string {
	v, ok := solarMap[k]
	if ok {
		return v
	}

	return ""
}

func UpdateSolarMapItem(k, v string) {
	_, ok := solarMap[k]
	//fmt.Printf("k=%s, v=%s\n", k, v)
	if ok {
		// update item
		solarMap[k] = v
	} else {
		// add new item
		solarMap[k] = v
	}
}

func DeleteSolarMapItem(k string) {
	_, ok := solarMap[k]
	if ok {
		delete(solarMap, k)
	}
}

///////////////////////////////////////////////////////////////////////////////
func Byte2str(s []byte) string {
	var ret string
	var ch [2]byte

	for i, v := range s {
		ch[i%2] = v
		if i%2 == 1 {
			tmp := string(ch[0]) + string(ch[1])
			if ver, err := strconv.ParseUint(tmp, 16, 32); err == nil {
				ret += string(ver)
			}

		}
	}
	//fmt.Println("ret=", ret)
	return ret
}

func Byte2Uint(s []byte) uint64 {
	var ret uint64
	var ch [2]byte

	for i, v := range s {
		ch[i%2] = v
		if i%2 == 1 {
			tmp := string(ch[0]) + string(ch[1])
			if ver, err := strconv.ParseUint(tmp, 16, 32); err == nil {
				ret = ret<<8 | uint64(ver)
			}
		}
	}
	//fmt.Println("ret=", ret)
	return ret
}

// style version is 8 letters
func PeekStyleVersion(s []byte) string {
	// only the header 16 bytes
	tmp := s[0 : lenStyleVersion<<1]
	return Byte2str(tmp)
}

//style should be made of by 3 different strings
func PeekStyleCodePart1(s []byte) string {
	begin := lenStyleVersion << 1
	end := begin + lenStyleCode1<<1
	tmp := s[begin:end]
	return Byte2str(tmp)
}

func PeekStyleCodePart2(s []byte) string {
	begin := lenStyleVersion<<1 + lenStyleCode1<<1
	end := begin + lenStyleCode2<<1
	tmp := s[begin:end]
	return Byte2str(tmp)
}

func PeekStyleCodePart3(s []byte) string {
	begin := lenStyleVersion<<1 + lenStyleCode1<<1 + lenStyleCode2<<1
	end := begin + lenStyleCode3<<1
	tmp := s[begin:end]
	return Byte2str(tmp)
}

func PeekStyleCode(s []byte) []string {
	stylecode := make([]string, 3, 3)

	stylecode[0] = PeekStyleCodePart1(s)
	stylecode[1] = PeekStyleCodePart2(s)
	stylecode[2] = PeekStyleCodePart3(s)

	return stylecode
}

func PeekRstTblName(s []byte) string {
	begin := lenStyleVersion<<1 + lenStyleCode<<1
	end := begin + lenCmdRsltTblName<<1
	tmp := s[begin:end]
	return Byte2str(tmp)
}

///////////////////////////////////////////////////////////////////////////////
func SolarzoomDelay(t int64) float64 {
	var s float64 = 1.0
	for i := 0; i < 3000; i++ {
		s *= float64(i)
	}
	return s
}

///////////////////////////////////////////////////////////////////////////////
func IsFileExist(name string) bool {
	fmt.Println("name=", name)
	_, err := os.Stat(name)
	return err == nil || os.IsExist(err)
}
