package models

import (
	"fmt"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//"strconv"
	"solarzoom/utils"
	"time"
)

const RUNDATA_TABLE_PREFIX string = "pv_inverter_run_data_"

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

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func init() {
	fmt.Println("package models: init function")
	//orm.RegisterModel(new(PvInverterRunData))
}

///////////////////////////////////////////////////////////////////////////////
//                             USE ORM Engine                                //
///////////////////////////////////////////////////////////////////////////////
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

func AddInverterRunData(record *PvInverterRunData) (int64, error) {
	o := orm.NewOrm()
	record.InputTime = time.Now().Unix()

	id, err := o.Insert(record)
	return id, err
}

///////////////////////////////////////////////////////////////////////////////
//                             USE SQL Sentence                              //
///////////////////////////////////////////////////////////////////////////////
func getRunDataTableName(t int64) string {
	year := time.Unix(t, 0).Year()
	month := time.Unix(t, 0).Month()

	s := fmt.Sprintf("%s%d%02d", RUNDATA_TABLE_PREFIX, year, month)
	return s
}

func CreateRunDataTableBySQL(t int64) {
	year := time.Unix(t, 0).Year()
	month := time.Unix(t, 0).Month()

	tableName := getRunDataTableName(t)
	s := "CREATE TABLE IF NOT EXISTS"
	s = fmt.Sprintf("%s `%s` (", s, tableName)
	s = fmt.Sprintf("%s `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,", s)
	s = fmt.Sprintf("%s `ivt_id` int(11) NOT NULL DEFAULT '0' COMMENT '逆变器ID',", s)
	s = fmt.Sprintf("%s `batch_order` int(4) NOT NULL DEFAULT '0' COMMENT '数据采集序号',", s)
	s = fmt.Sprintf("%s `smpl_time` int(10) NOT NULL DEFAULT '0' COMMENT '数据采集时间',", s)
	s = fmt.Sprintf("%s `input_time` int(10) NOT NULL DEFAULT '0' COMMENT '数据写入时间',", s)
	s = fmt.Sprintf("%s `work_status` varchar(32) NOT NULL DEFAULT '' COMMENT '当前最新运行状态',", s)
	s = fmt.Sprintf("%s `run_time_total` double(8,2) NOT NULL DEFAULT '0' COMMENT '总运行时间',", s)
	s = fmt.Sprintf("%s `energy_total` double(10,2) NOT NULL DEFAULT '0' COMMENT '总发电量',", s)
	s = fmt.Sprintf("%s `energy_day` float(8,2) NOT NULL DEFAULT '0' COMMENT '今日总发电量',", s)
	s = fmt.Sprintf("%s `internal_temperature` float(5, 2) NOT NULL DEFAULT '0' COMMENT '逆变器内部温度',", s)
	s = fmt.Sprintf("%s `vdc_pv1` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'PV1直流电压',", s)
	s = fmt.Sprintf("%s `idc_pv1` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'PV1直流电流',", s)
	s = fmt.Sprintf("%s `dcpower_pv1` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'PV1直流功率',", s)
	s = fmt.Sprintf("%s `vdc_pv2` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'PV2直流电压',", s)
	s = fmt.Sprintf("%s `idc_pv2` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'PV2直流电流',", s)
	s = fmt.Sprintf("%s `dcpower_pv2` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'PV2直流功率',", s)
	s = fmt.Sprintf("%s `vdc_pv3` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'PV3直流电压',", s)
	s = fmt.Sprintf("%s `idc_pv3` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'PV3直流电流',", s)
	s = fmt.Sprintf("%s `dcpower_pv3` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'PV3直流功率',", s)
	s = fmt.Sprintf("%s `vdc_pv4` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'PV4直流电压',", s)
	s = fmt.Sprintf("%s `idc_pv4` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'PV4直流电流',", s)
	s = fmt.Sprintf("%s `dcpower_pv4` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'PV4直流功率',", s)
	s = fmt.Sprintf("%s `pv1_resistor` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'PV1正对地绝缘阻抗',", s)
	s = fmt.Sprintf("%s `pv2_resistor` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'PV2正对地绝缘阻抗',", s)
	s = fmt.Sprintf("%s `pv3_resistor` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'PV3正对地绝缘阻抗',", s)
	s = fmt.Sprintf("%s `pv4_resistor` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'PV4正对地绝缘阻抗',", s)
	s = fmt.Sprintf("%s `aver_vdc_pv` float(4, 1) NOT NULL DEFAULT '0' COMMENT '平均直流电压',", s)
	s = fmt.Sprintf("%s `idc_total` float(4, 1) NOT NULL DEFAULT '0' COMMENT '总直流电流',", s)
	s = fmt.Sprintf("%s `dcpower_total` float(9, 1) NOT NULL DEFAULT '0' COMMENT '总直流功率',", s)
	s = fmt.Sprintf("%s `vac_r` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'R相电压',", s)
	s = fmt.Sprintf("%s `iac_r` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'R相电流',", s)
	s = fmt.Sprintf("%s `acpower_r` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'R相输出功率',", s)
	s = fmt.Sprintf("%s `fac_r` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'R相频率',", s)
	s = fmt.Sprintf("%s `vac_s` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'S相电压',", s)
	s = fmt.Sprintf("%s `iac_s` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'S相电流',", s)
	s = fmt.Sprintf("%s `acpower_s` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'S相输出功率',", s)
	s = fmt.Sprintf("%s `fac_s` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'S相频率',", s)
	s = fmt.Sprintf("%s `vac_t` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'T相电压',", s)
	s = fmt.Sprintf("%s `iac_t` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'T相电流',", s)
	s = fmt.Sprintf("%s `acpower_t` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'T相输出功率',", s)
	s = fmt.Sprintf("%s `fac_t` float(5, 2) NOT NULL DEFAULT '0' COMMENT 'T相频率',", s)
	s = fmt.Sprintf("%s `aver_vac` float(4, 1) NOT NULL DEFAULT '0' COMMENT '平均交流电压',", s)
	s = fmt.Sprintf("%s `ac_active_power_total` float(8, 2) NOT NULL DEFAULT '0' COMMENT '总有功功率',", s)
	s = fmt.Sprintf("%s `iac_total` float(4, 1) NOT NULL DEFAULT '0' COMMENT '总交流电流',", s)
	s = fmt.Sprintf("%s `vac_balance` float(5, 2) NOT NULL DEFAULT '0' COMMENT '交流电压三相平衡度',", s)
	s = fmt.Sprintf("%s `iac_balance` float(5, 2) NOT NULL DEFAULT '0' COMMENT '交流直流三相平衡度',", s)
	s = fmt.Sprintf("%s `fgrid` float(5, 2) NOT NULL DEFAULT '0' COMMENT '电网频率',", s)
	s = fmt.Sprintf("%s `efficiency` float(7, 3) NOT NULL DEFAULT '0' COMMENT '逆变器效率',", s)
	s = fmt.Sprintf("%s `simu_kwh5_min` float(7, 2) NOT NULL DEFAULT '0' COMMENT '五分钟模拟发电量'", s)
	s = fmt.Sprintf("%s ) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='%d年%02d月份所有逆变器实时运行数据表';", s, year, month)

	//fmt.Println("s=", s)
	o := orm.NewOrm()
	_, err := o.Raw(s).Exec()
	if err == nil {
		utils.WriteDebugLog("Create %s table ...... DONE", tableName)
		fmt.Println("Create %s table SUCCESS!", tableName)
	} else {
		fmt.Printf("Create err=%v\n", err)
		fmt.Println("Create table ERROR!")
		utils.WriteErrorLog("Create %s table ...... ERROR", tableName)
	}
}

func InsertRunDataTableItemBySQL(r *PvInverterRunData) {
	// use the r.SmplTime to generate the table
	t := time.Unix(r.SmplTime, 0)
	fmt.Println("t=%v", t)
	//fmt.Println("t.Unix=%v", t.Unix())

	// first, try to create the table
	CreateRunDataTableBySQL(t.Unix())

	// insert data to the table
	tableName := getRunDataTableName(t.Unix())

	s := fmt.Sprintf("INSERT INTO `%s`", tableName)
	s = fmt.Sprintf("%s (`ivt_id`, `batch_order`, `smpl_time`, `input_time`, `work_status`, `run_time_total`, `energy_total`, `energy_day`, `internal_temperature`, `vdc_pv1`, `idc_pv1`, `dcpower_pv1`, `vdc_pv2`, `idc_pv2`, `dcpower_pv2`, `vdc_pv3`, `idc_pv3`, `dcpower_pv3`, `vdc_pv4`, `idc_pv4`, `dcpower_pv4`, `pv1_resistor`, `pv2_resistor`, `pv3_resistor`, `pv4_resistor`, `aver_vdc_pv`, `idc_total`, `dcpower_total`, `vac_r`, `iac_r`, `acpower_r`, `fac_r`, `vac_s`, `iac_s`, `acpower_s`, `fac_s`, `vac_t`, `iac_t`, `acpower_t`, `fac_t`, `aver_vac`, `ac_active_power_total`, `iac_total`, `vac_balance`, `iac_balance`, `fgrid`, `efficiency`, `simu_kwh5_min`)", s)
	s = fmt.Sprintf("%s VALUES ", s)
	s = fmt.Sprintf("%s ('%v', '%v', '%v', '%v', '%s', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v');", s, r.IvtId, r.BatchOrder, r.SmplTime, time.Now().Unix(), r.WorkStatus, r.RunTimeTotal, r.EnergyTotal, r.EnergyDay, r.InternalTemperature, r.VdcPv1, r.IdcPv1, r.DcpowerPv1, r.VdcPv2, r.IdcPv2, r.DcpowerPv2, r.VdcPv3, r.IdcPv3, r.DcpowerPv3, r.VdcPv4, r.IdcPv4, r.DcpowerPv4, r.Pv1Resistor, r.Pv2Resistor, r.Pv3Resistor, r.Pv4Resistor, r.AverVdcPv, r.IdcTotal, r.DcpowerTotal, r.VacR, r.IacR, r.AcpowerR, r.FacR, r.VacS, r.IacS, r.AcpowerS, r.FacS, r.VacT, r.IacT, r.AcpowerT, r.FacT, r.AverVac, r.AcActivePowerTotal, r.IacTotal, r.VacBalance, r.IacBalance, r.Fgrid, r.Efficiency, r.SimuKwh5Min)
	//fmt.Println("s=", s)

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
