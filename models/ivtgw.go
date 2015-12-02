package models

import (
	"fmt"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//_ "github.com/go-sql-driver/mysql"
)

type PvCollectorInverter struct {
	IvtId      int64  `orm:"pk;auto"` // primary key, inverter id
	IvtSn      string // inverter sn
	GwSn       string // gateway sn
	IvtInfoId  int32  // inverter basic info id
	IvtAddress string // inverter address
}

func init() {
	fmt.Println("register PvCollectorInverter table")
	orm.RegisterModel(new(PvCollectorInverter))
}

func AddGwIVTItem(infoid int32, ivtsn, gwsn, ivtaddr string) (int64, error) {
	o := orm.NewOrm()
	item := &PvCollectorInverter{
		IvtSn:      ivtsn,
		GwSn:       gwsn,
		IvtInfoId:  infoid,
		IvtAddress: ivtaddr,
	}
	id, err := o.Insert(item)

	return id, err
}

func GetIvtIdByIvtSN(sn string) (id int64, err error) {
	o := orm.NewOrm()
	item := PvCollectorInverter{IvtSn: sn}
	err = o.Read(&item, "ivt_sn")
	return item.IvtId, err
}

func GetIvtIdByGWInfo(sn, addr string) (id int64, err error) {
	o := orm.NewOrm()
	item := PvCollectorInverter{GwSn: sn, IvtAddress: addr}
	err = o.Read(&item, "gw_sn", "ivt_address")
	return item.IvtId, err
}
