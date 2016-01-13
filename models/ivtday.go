package models

import (
	"fmt"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"solarzoom/utils"
	"strconv"
	"time"
)

type PvInverterDayData struct {
	Id                  int64 `orm:"pk;auto"`
	IvtId               int64
	InputTime           int64
	Day                 int64
	DataValidate        int32   // solarzoom
	AcActivePowerTotal  float64 `orm:"digits(8);decimals(2)"`
	EnergyToday         float64 `orm:"digits(8);decimals(2)"`
	EnergyTotal         float64 `orm:"digits(10);decimals(2)"`
	PowerContent        string
	NominalHours        int32   // count number
	TodayHours          float32 `orm:"digits(5);decimals(3)"` // solarzoom
	AvgDirectPower      float32 `orm:"digits(8);decimals(1)"` // solarzoom
	AvgAlternatingPower float32 `orm:"digits(8);decimals(1)"` // solarzoom
	AvgEfficiency       float32 `orm:"digits(4);decimals(3)"` // solarzoom
}

const DAY_TABLE_PREFIX string = "pv_inverter_day_data_"

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func init() {
	fmt.Println("register PvInverterDayData table")
	//orm.RegisterModel(new(PvInverterDayData))

	// fmt.Printf("current seconds is:%v\n", time.Now().Unix())
	// fmt.Printf("today seconds is:%v\n", GetCurrentDay())
}

func NewPvInverterDayData() *PvInverterDayData {
	return &PvInverterDayData{}
}

func CalcDayTableDayItem(t int64) int64 {
	year, month, day := time.Unix(t, 0).Date()
	date := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return date.Unix()
}

func GetCurrentDay() int64 {
	year, month, day := time.Now().Date()
	date := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return date.Unix()
}

func (u *PvInverterDayData) TableName() string {
	fmt.Println("Get Inverter day data table name!")
	year := time.Now().Year()
	//fmt.Println("year=", year)
	return "pv_inverter_day_data_" + strconv.Itoa(year)
}

func (u *PvInverterDayData) TableEngine() string {
	return "MyISAM"
}

///////////////////////////////////////////////////////////////////////////////
func AddPvInverterDayData(ivtid int64) (int64, error) {
	o := orm.NewOrm()
	item := &PvInverterDayData{
		IvtId:     ivtid,
		InputTime: time.Now().Unix(),
		Day:       GetCurrentDay(),
	}
	id, err := o.Insert(item)

	return id, err
}

// get inverter today's record
func GetPvInverterTodayRecord(ivtid int64) (id int64, err error) {
	o := orm.NewOrm()
	record := PvInverterDayData{IvtId: ivtid, Day: GetCurrentDay()}
	err = o.Read(&record, "ivt_id", "day")

	return record.Id, err
}

// get inverter special day's record
func GetPvInverterDayRecord(ivtid, day int64) (id int64, err error) {
	o := orm.NewOrm()
	record := PvInverterDayData{IvtId: ivtid, Day: day}
	err = o.Read(&record, "ivt_id", "day")

	return record.Id, err
}

func GetPVInverterTodayHisPower(ivtid int64) (hisPower string, err error) {
	o := orm.NewOrm()
	record := PvInverterDayData{IvtId: ivtid, Day: GetCurrentDay()}

	err = o.Read(&record, "ivt_id", "power_content")

	return record.PowerContent, err
}

func UpdatePvInverterTodayRecord(id int64, power, totalEnergy, todayEnergy float64, count int32, content string) error {
	o := orm.NewOrm()
	record := PvInverterDayData{Id: id}
	if o.Read(&record) == nil {
		//fmt.Println("update DB day table item!")
		record.InputTime = time.Now().Unix()
		record.AcActivePowerTotal = power
		record.EnergyToday = todayEnergy
		record.EnergyTotal = totalEnergy
		record.PowerContent = content
		record.NominalHours = count

		//fmt.Printf("power=%v, todayEnergy=%v, total=%v, content=%v\n", record.AcActivePowerTotal, record.EnergyToday, totalEnergy, content)
		_, err := o.Update(&record)

		return err
	}
	return nil
}

///////////////////////////////////////////////////////////////////////////////
/////////////////////////////// for solarzoom /////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func UpdateRecord(id int64, valid int32, hours, dpwr, apwr, efficiency float32) error {
	o := orm.NewOrm()
	record := PvInverterDayData{Id: id}
	if o.Read(&record) == nil {
		record.InputTime = time.Now().Unix()
		record.DataValidate = valid
		record.TodayHours = hours
		record.AvgDirectPower = dpwr
		record.AvgAlternatingPower = apwr
		record.AvgEfficiency = efficiency
		_, err := o.Update(&record)
		//fmt.Printf("Update ivt:%d at day %d SUCCESS!\n", record.IvtId, record.Day)
		return err
	}

	return nil
}

///////////////////////////////////////////////////////////////////////////////
//                             USE SQL Sentence                              //
///////////////////////////////////////////////////////////////////////////////
// test for SQL sentence
func getDayTableName(t int64) string {
	year := time.Unix(t, 0).Year()

	s := fmt.Sprintf("%s%d", DAY_TABLE_PREFIX, year)
	return s
}

func CreateDayTableBySQL(t int64) {
	year := time.Unix(t, 0).Year()
	//month := time.Unix(t, 0).Month()
	tableName := getDayTableName(t)

	s := "CREATE TABLE IF NOT EXISTS"
	s = fmt.Sprintf("%s `%s`", s, tableName)
	s = fmt.Sprintf("%s ( `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,", s)
	s = fmt.Sprintf("%s `ivt_id` int(11) NOT NULL DEFAULT '0' COMMENT '逆变器ID',", s)
	s = fmt.Sprintf("%s `input_time` int(10) DEFAULT NULL DEFAULT '0' COMMENT '最后更新时间',", s)
	s = fmt.Sprintf("%s `day` int(10) NOT NULL DEFAULT '0' COMMENT '数据日期',", s)
	s = fmt.Sprintf("%s `data_validate` smallint(1) NOT NULL DEFAULT '0' COMMENT '数据完整性验算',", s)
	s = fmt.Sprintf("%s `ac_active_power_total` float(8,2) NOT NULL DEFAULT '0.00' COMMENT '实时总功率',", s)
	s = fmt.Sprintf("%s `energy_today` float(8,2) NOT NULL DEFAULT '0.00' COMMENT '当日实时总发电量',", s)
	s = fmt.Sprintf("%s `energy_total` double(10,2) NOT NULL DEFAULT '0.00' COMMENT '总发电量',", s)
	s = fmt.Sprintf("%s `power_content` text NOT NULL DEFAULT '' COMMENT '当日历史功率',", s)
	s = fmt.Sprintf("%s `nominal_hours` int(11) NOT NULL DEFAULT '0' COMMENT '名义发电小时数',", s)
	s = fmt.Sprintf("%s `today_hours` float(5,3) NOT NULL DEFAULT '0.000' COMMENT '当日有效发电小时数',", s)
	s = fmt.Sprintf("%s `avg_direct_power` float(8,1) NOT NULL DEFAULT '0.0' COMMENT '当日有效直流平均功率',", s)
	s = fmt.Sprintf("%s `avg_alternating_power` float(8,1) NOT NULL DEFAULT '0.0' COMMENT '当日截止到目前的有效交流平均功率',", s)
	s = fmt.Sprintf("%s `avg_efficiency` float(4,3) NOT NULL DEFAULT '0.000' COMMENT '当日截止到目前的逆变器平均效率')", s)
	s = fmt.Sprintf("%s ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='%d年所有逆变器每日数据统计表';", s, year)
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	_, err := o.Raw(s).Exec()
	if err == nil {
		fmt.Println("Create %s table SUCCESS!", tableName)
		utils.WriteDebugLog("Create %s table ...... DONE", tableName)
	} else {
		fmt.Printf("Create err=%v\n", err)
		fmt.Println("Create table ERROR!")
		utils.WriteErrorLog("Create %s table ...... ERROR", tableName)
	}

}

func DoInsertDayTableRecordBySQL(r *PvInverterDayData) {
	tableName := getDayTableName(r.Day)

	s := fmt.Sprintf("INSERT INTO `%s`", tableName)
	s = fmt.Sprintf("%s (`ivt_id`, `input_time`, `day`, `data_validate`, `ac_active_power_total`, `energy_today`, `energy_total`, `power_content`, `nominal_hours`, `today_hours`, `avg_direct_power`, `avg_alternating_power`, `avg_efficiency`)", s)
	s = fmt.Sprintf("%s VALUES ", s)
	s = fmt.Sprintf("%s ('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v');", s, r.IvtId, time.Now().Unix(), r.Day, r.DataValidate, r.AcActivePowerTotal, r.EnergyToday, r.EnergyTotal, r.PowerContent, r.NominalHours, r.TodayHours, r.AvgDirectPower, r.AvgAlternatingPower, r.AvgEfficiency)
	//fmt.Println("s=", s)

	// Run the SQL
	o := orm.NewOrm()
	res, err := o.Raw(s).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		utils.WriteDebugLog("Insert a record to %s table ...... DONE", tableName)
	} else {
		fmt.Printf("err=%v\n", err)
		fmt.Println("mysql insert data have an ERROR!")
		utils.WriteErrorLog("Insert a record to %s table ...... ERROR", tableName)
	}
}

func InsertDayTableItemBySQL(r *PvInverterDayData) {
	// first, try to create the table
	CreateDayTableBySQL(r.Day)

	// second: generate the INSERT SQL sentence
	DoInsertDayTableRecordBySQL(r)

}

func DoUpdateDayTableItemBySQL(r *PvInverterDayData) error {
	tableName := getDayTableName(r.Day)

	s := fmt.Sprintf("UPDATE `%s`", tableName)
	s = fmt.Sprintf("%s SET ", s)
	s = fmt.Sprintf("%s `input_time` = '%v',", s, time.Now().Unix())
	s = fmt.Sprintf("%s `data_validate` = '%v',", s, r.DataValidate)
	s = fmt.Sprintf("%s `ac_active_power_total` = '%v',", s, r.AcActivePowerTotal)
	s = fmt.Sprintf("%s `energy_today` = '%v',", s, r.EnergyToday)
	s = fmt.Sprintf("%s `energy_total` = '%v',", s, r.EnergyTotal)
	s = fmt.Sprintf("%s `power_content` = '%v',", s, r.PowerContent)
	s = fmt.Sprintf("%s `nominal_hours` = '%v',", s, r.NominalHours)
	s = fmt.Sprintf("%s `today_hours` = '%v',", s, r.TodayHours)
	s = fmt.Sprintf("%s `avg_direct_power` = '%v',", s, r.AvgDirectPower)
	s = fmt.Sprintf("%s `avg_alternating_power` = '%v',", s, r.AvgAlternatingPower)
	s = fmt.Sprintf("%s `avg_efficiency` = '%v'", s, r.AvgEfficiency)
	s = fmt.Sprintf("%s WHERE (`ivt_id` = %d AND `day` = %v);", s, r.IvtId, r.Day)
	//fmt.Println("s=", s)

	o := orm.NewOrm()
	_, err := o.Raw(s).Exec()
	if err == nil {
		utils.WriteDebugLog("Update record(ivt_id=%d & day=%v) in table %s  ...... DONE", r.IvtId, r.Day, tableName)
	} else {
		utils.WriteErrorLog("Update record(ivt_id=%d & day=%v) in table %s  ...... ERROR", r.IvtId, r.Day, tableName)
	}

	return err
}

// This API can used by other module
// Parameters: r --- *PvInverterDayData
//             this para record all the data, r.Day is the second in a special day's 0:00:00
func UpdateDayTableRecordBySQL(r *PvInverterDayData) {
	// create the day table
	CreateDayTableBySQL(r.Day)

	// select the record
	err := SelectDayTableRecordBySQL(r)
	if err == nil {
		// update
		DoUpdateDayTableItemBySQL(r)
	} else {
		// insert
		DoInsertDayTableRecordBySQL(r)
	}
}

func SelectDayTableRecordBySQL(r *PvInverterDayData) error {
	tableName := getDayTableName(r.Day)

	s := fmt.Sprintf("SELECT * FROM `%s`", tableName)
	s = fmt.Sprintf("%s WHERE (`ivt_id` = %d AND `day` = %v);", s, r.IvtId, r.Day)

	var selRecord PvInverterDayData
	o := orm.NewOrm()
	err := o.Raw(s).QueryRow(&selRecord)

	if err == nil {
		fmt.Printf("SelectItem=%v\n", selRecord)
	} else {
		fmt.Println("Select NONE! Error=%v", err)
	}

	return err
}

func GetPowerContentInDayTable(r *PvInverterDayData) (string, error) {
	tableName := getDayTableName(r.Day)

	s := fmt.Sprintf("SELECT * FROM `%s`", tableName)
	s = fmt.Sprintf("%s WHERE (`ivt_id` = %d AND `day` = %v);", s, r.IvtId, r.Day)

	var selRecord PvInverterDayData
	o := orm.NewOrm()
	err := o.Raw(s).QueryRow(&selRecord)

	if err == nil {
		return selRecord.PowerContent, nil
	}

	return "", err
}

// for solarzoom API
func UpdateDayTableRecordWith() {

}
