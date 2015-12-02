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
	DataInputTime       int64
	Day                 int
	DataValidate        int32 // solarzoom
	TotalEffecPower     float32
	TodayKwh            float32
	TotalKwh            float64
	PowerContent        string
	WorkHours           float32
	EffectHours         float32 // solarzoom
	AverDirectPower     float32 // solarzoom
	AverAlterPower      float32 // solarzoom
	AverIvtProductivity float32 // solarzoom
}

func init() {
	fmt.Println("register PvStationDayData table")
	orm.RegisterModel(new(PvStationDayData))
}

func AddPvStationDayData(ivtid int64) (int64, error) {
	o := orm.NewOrm()
	item := &PvStationDayData{
		IvtId:         ivtid,
		DataInputTime: time.Now().Unix(),
		Day:           time.Now().Day(),
	}
	id, err := o.Insert(item)

	return id, err
}

func GetPVStationCurrDayRecord(ivtid int64, day int) (id int64, err error) {
	o := orm.NewOrm()
	item := PvStationDayData{IvtId: ivtid, Day: day}
	err = o.Read(&item, "ivt_id", "day")

	return item.Id, err
}

// for solarzoom
func UpdateRecord(id int64, valid int32, hours, dpwr, apwr, productivity float32) error {
	o := orm.NewOrm()
	record := PvStationDayData{Id: id}
	if o.Read(&record) == nil {
		record.DataValidate = valid
		record.EffectHours = hours
		record.AverDirectPower = dpwr
		record.AverAlterPower = apwr
		record.AverIvtProductivity = productivity
		_, err := o.Update(&record)
		//fmt.Printf("Update ivt:%d at day %d SUCCESS!\n", record.IvtId, record.Day)
		return err
	}

	return nil
}
