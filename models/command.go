package models

import (
	"fmt"
	//"github.com/astaxie/beego"
)

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
const GWCMD_FIRMWARE_UPDATE string = "FWUpdate"
const GWCMD_UPDATE_SBFILE string = "UpdateSBFile"
const GWCMD_UPDATE_IVTTABLE string = "UpdateIVTTable"
const GWCMD_SYNC_TIME string = "SyncTime"
const GWCMD_DEVICE_RESET string = "Reset"
const GWCMD_DEVIDE_REAUTH string = "ReAuth"
const GWCMD_REPORT_SELFINFO string = "ReportSelfInfo"

const GWCMD_RUN_PREDEFCMD string = "RunPreDefCmd"
const GWCMD_RUN_CUSTOMCMD string = "RunCustomCmd"

const DEFAULT_COMMAND_NUM int = 10

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

type GWCommand struct {
	Cmd  string
	Para string
}

//type GWCommandList []*GWCommand

var GWCommandMap map[string][]*GWCommand
var OneTimeCommands = []string{GWCMD_FIRMWARE_UPDATE, GWCMD_UPDATE_SBFILE, GWCMD_UPDATE_IVTTABLE,
	GWCMD_SYNC_TIME, GWCMD_DEVICE_RESET, GWCMD_DEVIDE_REAUTH, GWCMD_REPORT_SELFINFO}
var MultiTimeCommands = []string{GWCMD_RUN_PREDEFCMD, GWCMD_RUN_CUSTOMCMD}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
func init() {
	GWCommandMap = make(map[string][]*GWCommand)

	//cmdList = make([]*GWCommand, )
	//cmd := NewGWCommand(GWCMD_DEVICE_RESET, "")

	fmt.Println("init the command map!")

	//AppendNewCommand("nooo001", GWCMD_DEVICE_RESET, "")
	//DeleteAllCommands("nooo001")
	//AppendNewCommand("nooo001", GWCMD_FIRMWARE_UPDATE, "1.txt")
	//AppendNewCommand("nooo001", GWCMD_DEVIDE_REAUTH, "")
	//fmt.Printf("GWCommandMap=%v\n", GWCommandMap)
	//list := GWCommandMap["nooo001"]
	//fmt.Println("list=", list)
}

func NewGWCommand(cmd, para string) *GWCommand {
	return &GWCommand{cmd, para}
}

func isPredefineCommand(cmd string) bool {
	for _, oneTimeCommand := range OneTimeCommands {
		if cmd == oneTimeCommand {
			return true
		}
	}

	for _, multiTimeCommand := range MultiTimeCommands {
		if cmd == multiTimeCommand {
			return true
		}
	}

	return false
}

func isOneTimeCommand(cmd string) bool {
	for _, oneTimeCommand := range OneTimeCommands {
		if cmd == oneTimeCommand {
			return true
		}
	}

	return false
}

func isHaveTheCommand(cmd string, commands []*GWCommand) bool {
	for _, oldCmd := range commands {
		if oldCmd != nil && oldCmd.Cmd == cmd {
			return true
		}
	}
	return false
}

func getCommandIndex(cmd string, commands []*GWCommand) int {
	for i, oldCmd := range commands {
		//fmt.Printf("i=%d, oldCmd=%v\n", i, oldCmd)
		if oldCmd != nil && oldCmd.Cmd == cmd {
			return i
		}
	}
	return -1
}

func doAppendNewCommand(cmd, content string, commands []*GWCommand) []*GWCommand {
	//fmt.Println("doAppendNewCommand Entry!")
	//fmt.Printf("command=%s, para=%s\n", cmd, content)
	if ok := isOneTimeCommand(cmd); ok {
		// once command, need cheek if exist
		//fmt.Println("One Time Run Command!")
		if i := getCommandIndex(cmd, commands); i == -1 {
			//fmt.Println("NOT Exist!")
			item := NewGWCommand(cmd, content)
			commands = append(commands, item)
			//fmt.Println("append:commands=", commands)
			//commands[len(commands)] = item
		} else {
			//fmt.Println("Already Exist!")
			item := NewGWCommand(cmd, content)
			commands[i] = item
		}
	} else {
		// multi command, append another instance
		//fmt.Println("Multi Time Run Command!")
		item := NewGWCommand(cmd, content)
		commands = append(commands, item)
	}
	return commands
	//fmt.Println("commands=", commands)
}

func AppendNewCommand(gwsn, cmd, content string) {
	if gwsn != "" {
		if cmdList, ok := GWCommandMap[gwsn]; ok {
			// add the command
			//fmt.Printf("Append COMMAND to the sn(%s)!, cmdList=%p\n", gwsn, cmdList)

			if cmdOK := isPredefineCommand(cmd); cmdOK {
				cmdList = doAppendNewCommand(cmd, content, cmdList)
				GWCommandMap[gwsn] = cmdList
			}
		} else {
			// create the command map & add the command
			//fmt.Printf("Create the sn(%s) COMMAND list!\n", gwsn)
			cmdList = make([]*GWCommand, 0)
			//fmt.Printf("cmdList=%p\n", cmdList)

			//item := NewGWCommand(cmd, content)
			//cmdList = append(cmdList, item, item, item)
			//GWCommandMap[gwsn] = cmdList
			cmdList = doAppendNewCommand(cmd, content, cmdList)
			GWCommandMap[gwsn] = cmdList

			//fmt.Println("After AppendNewCommand: GWCommandMap=", GWCommandMap)
		}
	}
}

// how can this happen!!!!
func DeleteCommand(gwsn, cmd string) {
	if gwsn != "" {
		if _, ok := GWCommandMap[gwsn]; ok {
			//delete(cmdItem, cmd)
		}
	}
}

func DeleteAllCommands(gwsn string) {
	if gwsn != "" {
		if _, ok := GWCommandMap[gwsn]; ok {
			GWCommandMap[gwsn] = nil
		}
	}
}

// serial a command to the JSON style
func SerialOneCommand(cmd string, para string) string {
	s := ""
	s += fmt.Sprintf("{\"command\":\"%s\", \"parameters\":\"%s\"}", cmd, para)
	return s
}

// serial command list to the JSON style
func SerialCommands(gwsn string) string {
	if gwsn != "" {
		if cmds, ok := GWCommandMap[gwsn]; ok {
			s := "["
			for i, cmd := range cmds {
				if i != 0 {
					s += ","
				}
				s += SerialOneCommand(cmd.Cmd, cmd.Para)
			}
			s += "]"
			if s == "[]" {
				s = "\"\""
			}
			return s
		}
	}
	return "\"\""
}
