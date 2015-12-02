package models

// import (
// 	"fmt"
// 	//"github.com/astaxie/beego"
// 	"github.com/astaxie/beego/orm"
// 	_ "github.com/go-sql-driver/mysql"
// 	"time"
// )

// type PvCollectorInverter struct {
// 	IvtId      int64  `orm:"pk;auto"` // primary key, inverter id
// 	IvtSn      string // inverter sn
// 	GwSn       string // gateway sn
// 	IvtInfoId  int32  // inverter basic info id
// 	IvtAddress string // inverter address
// }

// // inverter runtime data
// type PvInverterRunData struct {
// 	Id                int64 `orm:"pk;auto"` // main key, auto increase
// 	IvtId             int64 // ???
// 	DataIndex         int32 // data serial 1 - 288
// 	DataSampleTime    time.Time
// 	DataInputTime     time.Time
// 	CurrWorkStatus    int32
// 	TotalRunTime      float64
// 	TotalKwh          float64
// 	TodayKwh          float32
// 	InsideTemperature float32
// 	PV1Volt           float32
// 	PV1Curr           float32
// 	PV1Power          float32
// 	PV2Volt           float32
// 	PV2Curr           float32
// 	PV2Power          float32
// 	PV3Volt           float32
// 	PV3Curr           float32
// 	PV3Power          float32
// 	PV4Volt           float32
// 	PV4Curr           float32
// 	PV4Power          float32
// 	PV1Resistor       float32
// 	PV2Resistor       float32
// 	PV3Resistor       float32
// 	PV4Resistor       float32
// 	AverDirectVolt    float32
// 	TotalDirectCurr   float32
// 	TotalDirectPower  float32
// 	RVolt             float32
// 	RCurr             float32
// 	RPower            float32
// 	RFreq             float32
// 	SVolt             float32
// 	SCurr             float32
// 	SPower            float32
// 	SFreq             float32
// 	TVolt             float32
// 	TCurr             float32
// 	TPower            float32
// 	TFreq             float32
// 	AverAlterVolt     float32
// 	TotalEffecPower   float32
// 	TotalAlterCurr    float32
// 	StdevVolt         float32
// 	StdevCurr         float32
// 	Frequency         float32
// 	IvtProductivity   float32
// 	SimuKwh5Min       float32
// }

// type PvStationDayData struct {
// 	Id                  int64 `orm:"pk;auto"`
// 	DataInputTime       time.Time
// 	Day                 string
// 	DataValidate        int32 // solarzoom
// 	TotalEffecPower     float32
// 	TodayKwh            float32
// 	TotalKwh            float64
// 	PowerContent        string
// 	WorkHours           float32
// 	EffectHours         float32 // solarzoom
// 	AverDirectPower     float32 // solarzoom
// 	AverAlterPower      float32 // solarzoom
// 	AverIvtProductivity float32 // solarzoom
// }

// type InverterInfo struct {
// 	Id          int64 `orm:"pk;auto"` // main key, auto increase
// 	SN          string
// 	Style       string
// 	Description string
// 	Version     string
// 	Price       float32
// }

// type User struct {
// 	Id            int64 `orm:"pk;auto"`
// 	UserId        int64
// 	UserLoginName string
// 	Password      string
// 	Mail          string
// 	Mobile        string
// 	UserType      string
// 	UserName      string
// 	State         int64
// }

// func init() {
// 	fmt.Println("package models: init function")
// 	orm.RegisterModel(new(PvCollectorInverter))
// 	//orm.RegisterModel(new(InverterInfo), new(User))
// 	//orm.RegisterModel()
// }

// func AddGwIVTItem(infoid int32, ivtsn, gwsn, ivtaddr string) (int64, error) {
// 	o := orm.NewOrm()
// 	item := &PvCollectorInverter{
// 		IvtSn:      ivtsn,
// 		GwSn:       gwsn,
// 		IvtInfoId:  infoid,
// 		IvtAddress: ivtaddr,
// 	}
// 	id, err := o.Insert(item)

// 	return id, err
// }

// func AddInverterInfo(sn, style, description, version string, price float32) (int64, error) {
// 	o := orm.NewOrm()
// 	info := &InverterInfo{
// 		SN:          sn,
// 		Style:       style,
// 		Description: description,
// 		Version:     version,
// 		Price:       price,
// 	}
// 	id, err := o.Insert(info)
// 	return id, err
// }

// func ReadInverterInfoById(id int64) {
// 	o := orm.NewOrm()
// 	info := InverterInfo{Id: id}

// 	err := o.Read(&info)

// 	if err == orm.ErrNoRows {
// 		fmt.Println("查询不到")
// 	} else if err == orm.ErrMissPK {
// 		fmt.Println("找不到主键")
// 	} else {
// 		fmt.Printf("info: id=%d, SN=%v, Style=%v, Description=%v, Version=%s, Price=%v\n",
// 			info.Id, info.SN, info.Style, info.Description, info.Version, info.Price)
// 	}
// }

// func ReadInvertInfoBySN(sn string) {
// 	o := orm.NewOrm()
// 	info := InverterInfo{SN: "SNXXXx"}

// 	err := o.Read(&info)

// 	if err == orm.ErrNoRows {
// 		fmt.Println("查询不到")
// 	} else if err == orm.ErrMissPK {
// 		fmt.Println("找不到主键")
// 	} else {
// 		fmt.Printf("info: id=%d, SN=%v, Style=%v, Description=%v, Version=%s, Price=%v\n",
// 			info.Id, info.SN, info.Style, info.Description, info.Version, info.Price)
// 	}
// }

// func UpdateInvertInfoById(id int64) {
// 	o := orm.NewOrm()
// 	info := InverterInfo{Id: id}
// 	if o.Read(&info) == nil {
// 		info.Style = "Noovo"
// 		if num, err := o.Update(&info); err == nil {
// 			fmt.Println(num)
// 		}
// 	}
// }

// func UpdateInvertInfoBySN(sn string) {

// }

// func DeleteInvertInfoById(id int64) {
// 	o := orm.NewOrm()
// 	if num, err := o.Delete(&InverterInfo{Id: id}); err == nil {
// 		fmt.Println(num)
// 	} else if err == orm.ErrNoRows {
// 		fmt.Println("查询不到")
// 	} else {
// 		fmt.Println(err.Error())
// 	}
// }

// func DeleteInvertInfoBySN(sn string) {
// 	//
// }

// func AddUser(userid, state int64, loginName, name, password, mail, mobile, usertype string) (int64, error) {
// 	o := orm.NewOrm()
// 	user := &User{
// 		UserId:        userid,
// 		UserLoginName: loginName,
// 		Password:      password,
// 		Mail:          mail,
// 		Mobile:        mobile,
// 		UserType:      usertype,
// 		UserName:      name,
// 		State:         state,
// 	}
// 	id, err := o.Insert(user)
// 	return id, err
// }
