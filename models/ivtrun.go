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
	Id                  int64 `orm:"pk;auto"` // main key, auto increase
	IvtId               int64 // ???
	BatchOrder          int32 // data serial 1 - 288
	SmplTime            int64 //
	InputTime           int64
	WorkStatus          string // ??? checkcmd???
	RunTimeTotal        float64
	EnergyTotal         float64
	EnergyDay           float64
	InternalTemperature float64
	VdcPv1              float64
	IdcPv1              float64
	DcpowerPv1          float64
	VdcPv2              float64
	IdcPv2              float64
	DcpowerPv2          float64
	VdcPv3              float64
	IdcPv3              float64
	DcpowerPv3          float64
	VdcPv4              float64
	IdcPv4              float64
	DcpowerPv4          float64
	Pv1Resistor         float64
	Pv2Resistor         float64
	Pv3Resistor         float64
	Pv4Resistor         float64
	AverVdcPv           float64
	IdcTotal            float64
	DcpowerTotal        float64
	VacR                float64
	IacR                float64
	AcpowerR            float64
	FacR                float64
	VacS                float64
	IacS                float64
	AcpowerS            float64
	FacS                float64
	VacT                float64
	IacT                float64
	AcpowerT            float64
	FacT                float64
	AverVac             float64
	AcActivePowerTotal  float64
	IacTotal            float64
	VacBalance          float64
	IacBalance          float64
	Fgrid               float64
	Efficiency          float64
	SimuKwh5Min         float64
}

func init() {
	fmt.Println("package models: init function")
	orm.RegisterModel(new(PvInverterRunData))
}

func NewPvInverterRunData() *PvInverterRunData {
	return &PvInverterRunData{}
}

func AddInverterRunData(record *PvInverterRunData) (int64, error) {
	o := orm.NewOrm()
	record.InputTime = time.Now().Unix()

	id, err := o.Insert(record)
	return id, err
}

// func GetInverterRunDataById() {

// }
