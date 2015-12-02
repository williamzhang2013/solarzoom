package models

import (
	"fmt"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// inverter runtime data
type PvInverterRunData struct {
	Id                int64 `orm:"pk;auto"` // main key, auto increase
	IvtId             int64 // ???
	DataIndex         int32 // data serial 1 - 288
	DataSampleTime    int64
	DataInputTime     int64
	CurrWorkStatus    int32
	TotalRunTime      float64
	TotalKwh          float64
	TodayKwh          float32
	InsideTemperature float32
	PV1Volt           float32
	PV1Curr           float32
	PV1Power          float32
	PV2Volt           float32
	PV2Curr           float32
	PV2Power          float32
	PV3Volt           float32
	PV3Curr           float32
	PV3Power          float32
	PV4Volt           float32
	PV4Curr           float32
	PV4Power          float32
	PV1Resistor       float32
	PV2Resistor       float32
	PV3Resistor       float32
	PV4Resistor       float32
	AverDirectVolt    float32
	TotalDirectCurr   float32
	TotalDirectPower  float32
	RVolt             float32
	RCurr             float32
	RPower            float32
	RFreq             float32
	SVolt             float32
	SCurr             float32
	SPower            float32
	SFreq             float32
	TVolt             float32
	TCurr             float32
	TPower            float32
	TFreq             float32
	AverAlterVolt     float32
	TotalEffecPower   float32
	TotalAlterCurr    float32
	StdevVolt         float32
	StdevCurr         float32
	Frequency         float32
	IvtProductivity   float32
	SimuKwh5Min       float32
}

func init() {
	fmt.Println("package models: init function")
	orm.RegisterModel(new(PvInverterRunData))
}

func AddInverterRunData(record PvInverterRunData) (int64, error) {
	o := orm.NewOrm()
	record.DataInputTime = time.Now().Unix()

	id, err := o.Insert(&record)
	return id, err
}

// func GetInverterRunDataById() {

// }
