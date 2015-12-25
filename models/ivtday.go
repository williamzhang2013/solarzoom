package models

import (
	"fmt"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type PvInverterDayData struct {
	Id                  int64 `orm:"pk;auto"`
	IvtId               int64
	InputTime           int64
	Day                 int64
	DataValidate        int32 // solarzoom
	AcActivePowerTotal  float64
	EnergyToday         float64
	EnergyTotal         float64
	PowerContent        string
	NominalHours        int32   // count number
	TodayHours          float32 // solarzoom
	AvgDirectPower      float32 // solarzoom
	AvgAlternatingPower float32 // solarzoom
	AvgEfficiency       float32 // solarzoom
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func init() {
	fmt.Println("register PvInverterDayData table")
	orm.RegisterModel(new(PvInverterDayData))

	// fmt.Printf("current seconds is:%v\n", time.Now().Unix())
	// fmt.Printf("today seconds is:%v\n", GetCurrentDay())
}

func GetCurrentDay() int64 {
	year, month, day := time.Now().Date()
	date := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return date.Unix()
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
