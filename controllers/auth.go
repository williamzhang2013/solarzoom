package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"solarzoom/utils"
	// "solarzoom/utils/ucmd"
	"strconv"
)

// session auth item --- record the auth state
// none --- nothing happens
// sn --- get sn
// ok --- passed auth
//const SessAuth string = "auth"
const URLAuth string = "/gw/auth"

const snlength int = 16
const cipherlength int = 32

///////////////////////////////////////////////////////////////////////////////
var ChipSN [8]uint8
var ChipCipherText [16]uint8

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
type AuthController struct {
	beego.Controller
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func getChipSNAtString(src string) error {
	length := len(src)
	for i := 0; i < length; i += 2 {
		var s string = string(src[i]) + string(src[i+1])
		//fmt.Println("getChipSNAtString:", s)
		if data, err := strconv.ParseInt(s, 16, 16); err == nil {
			ChipSN[i/2] = uint8(data)
			//fmt.Printf("ChipSN[%d]=0x%x\n", i/2, ChipSN[i/2])
		} else {
			return err
		}
	}
	return nil
}

func getCipherTextAtString(src string) error {
	length := len(src)
	for i := 0; i < length; i += 2 {
		var s string = string(src[i]) + string(src[i+1])
		if data, err := strconv.ParseInt(s, 16, 16); err == nil {
			ChipCipherText[i/2] = uint8(data)
			//fmt.Printf("ChipCipherText[%d]=0x%x\n", i/2, ChipCipherText[i/2])
		} else {
			return err
		}
	}

	return nil
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func DoSetSN(sn string, sess session.SessionStore) {
	if len(sn) == snlength {
		// sn works
		if err := getChipSNAtString(sn); err == nil {
			utils.SetChipSNArrayItem(ChipSN)
			utils.PrintChipSN()
			sess.Set(utils.SessAuth, "sn")
			utils.UpdateSolarMapItem(utils.SessAuth, "sn")
		}
	}
}

func DoSetCipher(cipher string, sess session.SessionStore) {
	if len(cipher) == cipherlength {
		// cipher works
		if err := getCipherTextAtString(cipher); err == nil {
			utils.SetChipCipherArrayItem(ChipCipherText)
			//utils.PrintAlpuCipherText()
			DoAuth(sess)
		}
	}
}

func DoAuth(sess session.SessionStore) {
	if utils.IsPassedAuth() {
		sess.Set(utils.SessAuth, "ok")
		utils.UpdateSolarMapItem(utils.SessAuth, "ok")
		fmt.Println("AUTH OK!")
		utils.WriteDebugLog("AUTH OK!")
	} else {
		sess.Set(utils.SessAuth, "none")
		utils.UpdateSolarMapItem(utils.SessAuth, "none")
	}
}

func handleAuthOKState(ctrl *AuthController, sess session.SessionStore) {
	ctrl.Data["command1"] = "cmd"
	ctrl.Data["value1"] = "cipher"
	ctrl.Data["command2"] = "errcode"
	ctrl.Data["value2"] = 0

	ctrl.TplNames = "cmd2.tpl"
	fmt.Println("Auth OK!")
	utils.WriteDebugLog("AUTH already OK!")
}

func handleAuthSNState(ctrl *AuthController, sess session.SessionStore) {
	// get cipher
	cipher := ctrl.GetString("cipher")
	fmt.Println("cipher=", cipher)

	utils.WriteDebugLog("Auth: get cipher")
	if len(cipher) == 0 {
		ctrl.Data["command1"] = "cmd"
		ctrl.Data["value1"] = "cipher"
		ctrl.Data["command2"] = "errcode"
		ctrl.Data["value2"] = 1
		ctrl.TplNames = "cmd2.tpl"
	} else if len(cipher) != 32 {
		ctrl.Data["command1"] = "cmd"
		ctrl.Data["value1"] = "cipher"
		ctrl.Data["command2"] = "errcode"
		ctrl.Data["value2"] = 2
		ctrl.TplNames = "cmd2.tpl"
	} else {
		DoSetCipher(cipher, sess)
		ctrl.Data["command1"] = "cmd"
		ctrl.Data["value1"] = "cipher"
		ctrl.Data["command2"] = "errcode"
		ctrl.Data["value2"] = 0

		ctrl.TplNames = "cmd2.tpl"
	}
}

func handleAuthInitState(ctrl *AuthController, sess session.SessionStore) {
	// get sn
	sn := ctrl.GetString("sn")
	fmt.Println("sn=", sn)

	utils.WriteDebugLog("Auth: get sn")
	if len(sn) == 0 {
		ctrl.Data["command1"] = "cmd"
		ctrl.Data["value1"] = "sn"
		ctrl.Data["command2"] = "errcode"
		ctrl.Data["value2"] = 1
		ctrl.TplNames = "cmd2.tpl"
	} else if len(sn) != 16 {
		ctrl.Data["command1"] = "cmd"
		ctrl.Data["value1"] = "sn"
		ctrl.Data["command2"] = "errcode"
		ctrl.Data["value2"] = 2
		ctrl.TplNames = "cmd2.tpl"
	} else {
		DoSetSN(sn, sess)
		ctrl.Data["command1"] = "cmd"
		ctrl.Data["value1"] = "sn"
		ctrl.Data["command2"] = "errcode"
		ctrl.Data["value2"] = 0

		ctrl.Data["command3"] = "ciphertext"
		ctrl.Data["value3"] = "12345678123456781234567812345678"
		ctrl.TplNames = "cmd3.tpl"
	}
}

func handleAuthRequest(ctrl *AuthController) {
	sess := ctrl.StartSession()
	state := sess.Get(utils.SessAuth)
	//state := utils.GetSolarMapItem(utils.SessAuth)

	fmt.Println("current state=", state)
	switch state {
	case "ok":
		handleAuthOKState(ctrl, sess)
	case "sn":
		handleAuthSNState(ctrl, sess)
	default:
		handleAuthInitState(ctrl, sess)
	}
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func (c *AuthController) Get() {
	handleAuthRequest(c)
}

func (c *AuthController) Post() {
	handleAuthRequest(c)
}
