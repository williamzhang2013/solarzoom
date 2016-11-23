package models

import (
	"fmt"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//_ "github.com/go-sql-driver/mysql"
	"strings"
	"strconv"
)

type PvCollectorInverter struct {
	IvtId      int64  `orm:"pk;auto"` // primary key, inverter id
	IvtSn      string // inverter sn
	GwSn       string // gateway sn
	IvtAddress string // inverter address
	Status     string // 0 --- can use; 1 --- can't use
}

func init() {
	fmt.Println("register PvCollectorInverter table")
	//orm.RegisterModel(new(PvCollectorInverter))
}

// func CreateCollectorInverterTable(order string) {
// 	s := "CREATE TABLE IF NOT EXISTS"
// 	s = fmt.Sprintf("%s `%s`", s, "pv_collector_inverter")
// 	s = fmt.Sprintf("%s ( `ivt_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '逆变器ID',", s)
// 	s = fmt.Sprintf("%s `ivt_sn` varchar(32) DEFAULT NULL COMMENT '逆变器SN',", s)
// 	s = fmt.Sprintf("%s `gw_sn` varchar(32) DEFAULT NULL COMMENT '采集器SN',", s)
// 	s = fmt.Sprintf("%s `ivt_address` varchar(255) DEFAULT NULL COMMENT '逆变器地址',", s)
// 	s = fmt.Sprintf("%s `status` varchar(255) DEFAULT '1' COMMENT '0不可用，1可用',", s)
// 	s = fmt.Sprintf("%s PRIMARY KEY (`ivt_id`)", s)
// 	s = fmt.Sprintf("%s) ENGINE=InnoDB AUTO_INCREMENT=69 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='采集器逆变器关联表';", s)
// 	fmt.Println("s=", s)

// 	o := orm.NewOrm()
// 	_, err := o.Raw(s).Exec()
// 	if err == nil {
// 		fmt.Println("Create %s table SUCCESS!", "pv_collector_inverter")
// 		dglogs.WriteDebugLog("Create %s table ...... DONE", "pv_collector_inverter")
// 	} else {
// 		fmt.Printf("Create err=%v\n", err)
// 		fmt.Println("Create table ERROR!")
// 		dglogs.WriteErrorLog("Create %s table ...... ERROR", "pv_collector_inverter")
// 	}		
// }

func AddGwIVTItem(ivtsn, gwsn, ivtaddr, status string) (int64, error) {
	o := orm.NewOrm()
	item := &PvCollectorInverter{
		IvtSn:      ivtsn,
		GwSn:       gwsn,
		IvtAddress: ivtaddr,
		Status:		status,
	}
	id, err := o.Insert(item)

	return id, err
}

// func GetIvtIdByIvtSN(sn string) (id int64, err error) {
// 	o := orm.NewOrm()
// 	item := PvCollectorInverter{IvtSn: sn}
// 	err = o.Read(&item, "ivt_sn")
// 	return item.IvtId, err
// }

func GetIvtIdByGWInfo(sn, addr string) (id int64, err error) {
	// o := orm.NewOrm()
	// item := PvCollectorInverter{GwSn: sn, IvtAddress: addr}
	// err = o.Read(&item, "gw_sn", "ivt_address")
	// return item.IvtId, err

	s := fmt.Sprintf("SELECT * FROM `%s`", "pv_collector_inverter")
	s = fmt.Sprintf("%s WHERE (`gw_sn` = '%s' AND `ivt_address` = '%s');", s, sn, addr)
	fmt.Println("s=", s)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(s).Values(&maps)
	fmt.Printf("num=%d, maps=%v\n", num, maps)

	if err == nil && num > 0 {
		ivt_id := (maps[0]["ivt_id"]).(string)
		id , _ = strconv.ParseInt(ivt_id, 10, 64)
		return id, nil
	} else {
		return 0, err
	}
}

func GetIvtIdByIvtSN(sn string) (id int64, err error) {
	//name := GetBoxTableName(order)
	fmt.Println("GetIvtIdByIvtSN:sn=",sn)
	//fmt.Printf("GetIvtIdByIvtSN: sn length=%d, indexByte 0=%d\n", len(sn), strings.IndexByte(sn, 0))
	var bytesLen = strings.IndexByte(sn, 0)
	if bytesLen == -1 {
		bytesLen = len(sn)
	}
	var ivtsn []byte = make([]byte, bytesLen)
	for i, ch := range sn {
		if ch != 0 {
			ivtsn[i] = sn[i]
		}
	}

	s := fmt.Sprintf("SELECT * FROM `%s`", "pv_collector_inverter")
	s = fmt.Sprintf("%s WHERE (`ivt_sn` = '%s');", s, string(ivtsn))
	fmt.Println("s=", s)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(s).Values(&maps)
	fmt.Printf("num=%d, maps=%v\n", num, maps)

	if err == nil && num > 0 {
		ivt_id := (maps[0]["ivt_id"]).(string)
		id , _ = strconv.ParseInt(ivt_id, 10, 64)
		return id, nil
	} else {
		return 0, err
	}
}