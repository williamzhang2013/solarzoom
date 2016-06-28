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
	o := orm.NewOrm()
	item := PvCollectorInverter{GwSn: sn, IvtAddress: addr}
	err = o.Read(&item, "gw_sn", "ivt_address")
	return item.IvtId, err
}

func GetIvtIdByIvtSN(sn string) (id int64, err error) {
	//name := GetBoxTableName(order)
	var ivtsn []byte = make([]byte, strings.IndexByte(sn, 0))
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