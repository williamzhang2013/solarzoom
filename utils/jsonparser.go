package utils

import (
	"fmt"
	"io/ioutil"
	"solarzoom/utils/simplejson"
	//"strconv"
)

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func parseJSONFile(name string) (*simplejson.Json, error) {
	bytes, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("ReadJSONFile:", err.Error())
		return nil, err
	}
	//fmt.Printf("sd file contents=%v\n", bytes)

	js, err := simplejson.NewJson(bytes)
	return js, err
}

func getCmdLength(cmd string) int {
	v := GetLogicCmdItem(cmd, GetLogicCmdArray())
	return v.length
}

func getCmdInfo(cmd string) (length int, mode int) {
	v := GetLogicCmdItem(cmd, GetLogicCmdArray())
	length = v.length
	mode = v.mode
	return
}

// TODO
// return map[string]interface{}
func HandleSDData(fname string, content []byte) map[string]interface{} {
	fmt.Printf("sd fname=%v\n", fname)
	retMap := make(map[string]interface{})
	goBook, err := parseJSONFile(fname)

	if err == nil {
		// travse, find Data first
		resultTbl := PeekRstTblName(content)
		valueData := goBook.Get("Data").Get(resultTbl)
		//valueData := 001
		//fmt.Println("valueData=", valueData)
		// find DataOrder item, this item must be a string array
		dataOrder, _ := valueData.Get("DataOrder").StringArray()
		fmt.Println("dataOrder=", dataOrder)
		var offset, cmdLength, cmdMode = 0, 0, 0
		var cmdUnit float64 = 1.0
		for _, v := range dataOrder {
			//fmt.Println(i, ": value=", v)
			// first, travse the commandArray array to get the length
			//fmt.Printf("%s=", v)
			//cmdLength = getCmdLength(v)
			cmdLength, cmdMode = getCmdInfo(v)
			//fmt.Printf("length=%v, mode=%v, ", cmdLength, cmdMode)
			if cmdLength != 0 {
				// logic command
				if cmdMode == 0 {
					// string
					s := Byte2str(content[offset : offset+cmdLength*2])
					//fmt.Printf("String: orig=%v, v=%s s=%s\n", content[offset : offset+cmdLength*2], v, s)
					retMap[v] = s
				} else {
					// uint64
					s := Byte2Uint(content[offset : offset+cmdLength*2])
					//fmt.Printf("Number: orig=%v, v=%s, s=%v\n", content[offset : offset+cmdLength*2], v, s)
					retMap[v] = s
				}
				offset += cmdLength * 2
			} else {
				detailData := valueData.Get(v) //{}interface

				dataLen := detailData.Get("Len")
				cmdLength = dataLen.MustInt()
				dataMode, isNum := detailData.CheckGet("Digit")

				if isNum == true {
					cmdUnit = dataMode.MustFloat64()
					cmdMode = 1
				} else {
					cmdMode = 0
				}
				//fmt.Println("key:=", v, ", data=", detailData, ", mode=", cmdMode, ", unit=", cmdUnit)

				if cmdMode == 0 {
					// string
					s := Byte2str(content[offset : offset+cmdLength*2])
					fmt.Println(v, "=", s)
					retMap[v] = s
				} else {
					// uint64
					s := Byte2Uint(content[offset : offset+cmdLength*2])
					fmt.Println(v, "=", s)
					retMap[v] = float64(s) * cmdUnit
				}
				offset += cmdLength * 2
			}
		}
	} else {
		fmt.Printf("Parse JSON file %s Error!\n", fname)
	}

	return retMap
}

func HandleJSONCmd(fname string, cmd string) (*simplejson.Json, error) {
	goBook, err := parseJSONFile(fname)
	if err == nil {
		value := goBook.Get(cmd)
		//fmt.Printf("HandleJSONCmd: value=%v\n", value)
		if _, err := value.Map(); err == nil {
			//fmt.Println("get value ", value)
			return value, nil
		} else {
			//fmt.Println("error:=", err)
			return nil, err
		}
	} else {
		return nil, err
	}
}
