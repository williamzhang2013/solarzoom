package models

import (
	"fmt"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type PvStationDayData struct {
	Id                  int64 `orm:"pk;auto"`
	IvtId               int64
	InputTime           int64
	Day                 int64
	DataValidate        int32 // solarzoom
	AcActivePowerTotal  float32
	EnergyDay           float32
	EnergyTotal         float64
	PowerContent        string
	NominalHours        int32   // count number
	EffectHours         float32 // solarzoom
	AvgDirectPower      float32 // solarzoom
	AvgAlternatingPower float32 // solarzoom
	AvgEfficiency       float32 // solarzoom
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func init() {
	fmt.Println("register PvStationDayData table")
	orm.RegisterModel(new(PvStationDayData))

	// fmt.Printf("current seconds is:%v\n", time.Now().Unix())
	// fmt.Printf("today seconds is:%v\n", GetCurrentDay())
}

func GetCurrentDay() int64 {
	year, month, day := time.Now().Date()
	date := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return date.Unix()
}

///////////////////////////////////////////////////////////////////////////////
func AddPvStationDayData(ivtid int64) (int64, error) {
	o := orm.NewOrm()
	item := &PvStationDayData{
		IvtId:     ivtid,
		InputTime: time.Now().Unix(),
		Day:       GetCurrentDay(),
	}
	id, err := o.Insert(item)

	return id, err
}

// get inverter today's record
func GetPVStationTodayRecord(ivtid int64) (id int64, err error) {
	o := orm.NewOrm()
	record := PvStationDayData{IvtId: ivtid, Day: GetCurrentDay()}
	err = o.Read(&record, "ivt_id", "day")

	return record.Id, err
}

// get inverter special day's record
func GetPVStationDayRecord(ivtid, day int64) (id int64, err error) {
	o := orm.NewOrm()
	record := PvStationDayData{IvtId: ivtid, Day: day}
	err = o.Read(&record, "ivt_id", "day")

	return record.Id, err
}

func UpdatePVStationTodayRecord(id int64, power float32, totalEnergy float64, todayEnergy float32, count int32, content string) error {
	o := orm.NewOrm()
	record := PvStationDayData{Id: id}
	if o.Read(&record) == nil {
		record.InputTime = time.Now().Unix()
		record.AcActivePowerTotal = power
		record.EnergyDay = todayEnergy
		record.EnergyTotal = totalEnergy
		record.PowerContent = content
		record.NominalHours = count

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
	record := PvStationDayData{Id: id}
	if o.Read(&record) == nil {
		record.InputTime = time.Now().Unix()
		record.DataValidate = valid
		record.EffectHours = hours
		record.AvgDirectPower = dpwr
		record.AvgAlternatingPower = apwr
		record.AvgEfficiency = efficiency
		_, err := o.Update(&record)
		//fmt.Printf("Update ivt:%d at day %d SUCCESS!\n", record.IvtId, record.Day)
		return err
	}

	return nil
}
