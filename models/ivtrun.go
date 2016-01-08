package models

import (
	"fmt"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//"strconv"
	"time"
)

// inverter runtime data
type PvInverterRunData struct {
	Id                  int64 `orm:"pk;auto"`   // main key, auto increase
	IvtId               int64 `orm:"size(11)"`  // inverter id
	BatchOrder          int32 `orm:"digits(4)"` // data serial 1 - 288
	SmplTime            int64 //
	InputTime           int64
	WorkStatus          string  `orm:"size(32)"` // ??? checkcmd???
	RunTimeTotal        float64 `orm:"digits(8);decimals(2)"`
	EnergyTotal         float64 `orm:"digits(10);decimals(2)"`
	EnergyDay           float64 `orm:"digits(8);decimals(2)"`
	InternalTemperature float64 `orm:"digits(5);decimals(2)"`
	VdcPv1              float64 `orm:"digits(5);decimals(2)"`
	IdcPv1              float64 `orm:"digits(5);decimals(2)"`
	DcpowerPv1          float64 `orm:"digits(5);decimals(2)"`
	VdcPv2              float64 `orm:"digits(5);decimals(2)"`
	IdcPv2              float64 `orm:"digits(5);decimals(2)"`
	DcpowerPv2          float64 `orm:"digits(5);decimals(2)"`
	VdcPv3              float64 `orm:"digits(5);decimals(2)"`
	IdcPv3              float64 `orm:"digits(5);decimals(2)"`
	DcpowerPv3          float64 `orm:"digits(5);decimals(2)"`
	VdcPv4              float64 `orm:"digits(5);decimals(2)"`
	IdcPv4              float64 `orm:"digits(5);decimals(2)"`
	DcpowerPv4          float64 `orm:"digits(5);decimals(2)"`
	Pv1Resistor         float64 `orm:"digits(5);decimals(2)"`
	Pv2Resistor         float64 `orm:"digits(5);decimals(2)"`
	Pv3Resistor         float64 `orm:"digits(5);decimals(2)"`
	Pv4Resistor         float64 `orm:"digits(5);decimals(2)"`
	AverVdcPv           float64 `orm:"digits(4);decimals(1)"`
	IdcTotal            float64 `orm:"digits(4);decimals(1)"`
	DcpowerTotal        float64 `orm:"digits(9);decimals(1)"`
	VacR                float64 `orm:"digits(5);decimals(2)"`
	IacR                float64 `orm:"digits(5);decimals(2)"`
	AcpowerR            float64 `orm:"digits(5);decimals(2)"`
	FacR                float64 `orm:"digits(5);decimals(2)"`
	VacS                float64 `orm:"digits(5);decimals(2)"`
	IacS                float64 `orm:"digits(5);decimals(2)"`
	AcpowerS            float64 `orm:"digits(5);decimals(2)"`
	FacS                float64 `orm:"digits(5);decimals(2)"`
	VacT                float64 `orm:"digits(5);decimals(2)"`
	IacT                float64 `orm:"digits(5);decimals(2)"`
	AcpowerT            float64 `orm:"digits(5);decimals(2)"`
	FacT                float64 `orm:"digits(5);decimals(2)"`
	AverVac             float64 `orm:"digits(4);decimals(1)"`
	AcActivePowerTotal  float64 `orm:"digits(8);decimals(2)"`
	IacTotal            float64 `orm:"digits(4);decimals(1)"`
	VacBalance          float64 `orm:"digits(5);decimals(2)"`
	IacBalance          float64 `orm:"digits(5);decimals(2)"`
	Fgrid               float64 `orm:"digits(5);decimals(2)"`
	Efficiency          float64 `orm:"digits(7);decimals(3)"`
	SimuKwh5Min         float64 `orm:"digits(7);decimals(2)"`
}

func init() {
	fmt.Println("package models: init function")
	orm.RegisterModel(new(PvInverterRunData))
}

func NewPvInverterRunData() *PvInverterRunData {
	return &PvInverterRunData{}
}

func (u *PvInverterRunData) TableName() string {
	fmt.Println("Get Inverter run data table name!")
	year := time.Now().Year()
	month := time.Now().Month()
	//fmt.Println("min=", min)
	s := fmt.Sprintf("pv_inverter_run_data_%d%02d", year, month)
	return s
}

func (u *PvInverterRunData) TableEngine() string {
	return "MyISAM"
}

func GenPvRunDataTable() {
	fmt.Println("Generate pv run data table")

	// err := orm.RunSyncdb("default", false, true)
	// if err != nil {
	// 	fmt.Println("Create the table ERROR!")
	// 	beego.Error(err)
	// }
}

func AddInverterRunData(record *PvInverterRunData) (int64, error) {
	o := orm.NewOrm()
	record.InputTime = time.Now().Unix()

	id, err := o.Insert(record)
	return id, err
}

// func GetInverterRunDataById() {

// }
