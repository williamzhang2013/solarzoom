package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"solarzoom/utils"
	"strconv"
)

// session auth item --- record the auth state
// none --- nothing happens
// sn --- get sn
// ok --- passed auth
const SessAuth string = "auth"
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
		fmt.Println("getChipSNAtString:", s)
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
			sess.Set(SessAuth, "sn")
		}
	}
}

func DoSetCipher(cipher string, sess session.SessionStore) {
	if len(cipher) == cipherlength {
		// cipher works
		if err := getCipherTextAtString(cipher); err == nil {
			utils.SetChipCipherArrayItem(ChipCipherText)
			utils.PrintAlpuCipherText()
			DoAuth(sess)
		}
	}
}

func DoAuth(sess session.SessionStore) {
	if utils.IsPassedAuth() {
		sess.Set(SessAuth, "ok")
		fmt.Println("AUTH OK!")
	} else {
		sess.Set(SessAuth, "none")
	}
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func (c *AuthController) Get() {
	sess := c.StartSession()
	state := sess.Get(SessAuth)
	//fmt.Println("auth get method:auth=", state)
	switch state {
	case "ok":
		// do nothing
		fmt.Println("auth already in OK state")
	case "sn":
		// get cipher
		cipher := c.GetString("cipher")
		fmt.Println("cipher=", cipher)
		DoSetCipher(cipher, sess)

	default:
		// get sn
		sn := c.GetString("sn")
		fmt.Println("sn=", sn)
		DoSetSN(sn, sess)
	}

	c.Data["command"] = "auth"
	c.Data["value"] = sess.Get(SessAuth)
	c.TplNames = "cmd.tpl"
}

func (c *AuthController) Post() {
	sess := c.StartSession()
	state := sess.Get(SessAuth)

	switch state {
	case "ok":
		// do nothing
		fmt.Println("auth already in OK state")
	case "sn":
		// get cipher
		cipher := c.GetString("cipher")
		fmt.Println("cipher=", cipher)
		DoSetCipher(cipher, sess)

	default:
		// get sn
		sn := c.GetString("sn")
		fmt.Println("sn=", sn)
		DoSetSN(sn, sess)
	}

	c.Data["command"] = "auth"
	c.Data["value"] = sess.Get(SessAuth)
	c.TplNames = "cmd.tpl"
}
