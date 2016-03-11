package models

import (
	"fmt"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"solarzoom/utils"
	//"strconv"
	//"time"
)

type PvInverterFaultData struct {
	Id           int64 `orm:"pk;auto"`
	IvtId        int64
	FaultMessage string
	StartTime    int64
	EndTime      int64 // solarzoom
	Status       int8
	Method       string
	Person       string
}

const FAULT_TABLE_PREFIX string = "pv_inverter_fault"

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func init() {
	fmt.Println("register PvInverterFaultData table")
	//orm.RegisterModel(new(PvInverterFaultData))

	// fmt.Printf("current seconds is:%v\n", time.Now().Unix())
	// fmt.Printf("today seconds is:%v\n", GetCurrentDay())
}

func NewPvInverterFaultData() *PvInverterFaultData {
	return &PvInverterFaultData{}
}

///////////////////////////////////////////////////////////////////////////////
//                             USE SQL Sentence                              //
///////////////////////////////////////////////////////////////////////////////
// test for SQL sentence
func CreateFaultTableBySQL() {
	s := "CREATE TABLE IF NOT EXISTS"
	s = fmt.Sprintf("%s `%s`", s, FAULT_TABLE_PREFIX)
	s = fmt.Sprintf("%s ( `id` int(11) AUTO_INCREMENT NOT NULL PRIMARY KEY,", s)
	s = fmt.Sprintf("%s `ivt_id` int(11) NOT NULL DEFAULT '0' COMMENT '逆变器ID',", s)
	s = fmt.Sprintf("%s `faultMessage` varchar(255) DEFAULT NULL COMMENT '故障信息内容',", s)
	s = fmt.Sprintf("%s `startTime` int(10) NOT NULL DEFAULT '0' COMMENT '故障发生时间',", s)
	s = fmt.Sprintf("%s `endTime` int(10) NOT NULL DEFAULT '0' COMMENT '故障结束时间',", s)
	s = fmt.Sprintf("%s `status` smallint(1) DEFAULT '0' COMMENT '处理状态,0未处理1处理中2处理完',", s)
	s = fmt.Sprintf("%s `method` text COMMENT '处理方法说明',", s)
	s = fmt.Sprintf("%s `person` varchar(32) DEFAULT NULL COMMENT '处理人'", s)
	s = fmt.Sprintf("%s) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='逆变器故障记录表';", s)
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	_, err := o.Raw(s).Exec()
	if err == nil {
		fmt.Println("Create %s table SUCCESS!", FAULT_TABLE_PREFIX)
		utils.WriteDebugLog("Create %s table ...... DONE", FAULT_TABLE_PREFIX)
	} else {
		fmt.Printf("Create err=%v\n", err)
		fmt.Println("Create table ERROR!")
		utils.WriteErrorLog("Create %s table ...... ERROR", FAULT_TABLE_PREFIX)
	}

}

func DoInsertFaultTableRecordBySQL(r *PvInverterFaultData) {
	s := fmt.Sprintf("INSERT INTO `%s`", FAULT_TABLE_PREFIX)
	s = fmt.Sprintf("%s (`ivt_id`, `faultMessage`, `startTime`, `endTime`, `status`, `method`, `person`)", s)
	s = fmt.Sprintf("%s VALUES ", s)
	s = fmt.Sprintf("%s ('%v', '%v', '%v', '%v', '%v', '%v', '%v');", s, r.IvtId, r.FaultMessage, r.StartTime, r.EndTime, r.Status, r.Method, r.Person)
	//fmt.Println("s=", s)

	// Run the SQL
	o := orm.NewOrm()
	res, err := o.Raw(s).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		utils.WriteDebugLog("Insert a record to %s table ...... DONE", FAULT_TABLE_PREFIX)
	} else {
		fmt.Printf("err=%v\n", err)
		fmt.Println("mysql insert data have an ERROR!")
		utils.WriteErrorLog("Insert a record to %s table ...... ERROR", FAULT_TABLE_PREFIX)
	}
}

func InsertFaultTableItemBySQL(r *PvInverterFaultData) {
	// first, try to create the table
	CreateFaultTableBySQL()

	// second: generate the INSERT SQL sentence
	DoInsertFaultTableRecordBySQL(r)

}
