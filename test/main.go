package main

import (
	"log"

	"bitbucket.org/Limard/win"
	"bitbucket.org/Limard/win/apppath"
)

func main() {
	// log.Println("osinfo.Is64bitOS()", osinfo.Is64bitOS())
	// log.Println("systempath.GetCommmonAppDataDirectory()", systempath.GetCommmonAppDataDirectory())
	// log.Println("systempath.GetCommonDesktopDir()", systempath.GetCommonDesktopDir())
	// win.LaunchInActiveSesstion(`C:\Windows\System32\calc.exe`)
	// args := apppath.ParseCommand("cmd /c REG IMPORT [TempDir]\\UpdateDataDir\\mPortM.reg")
}

func TestParseCommand() {
	args := apppath.ParseCommand("cmd /c REG IMPORT C:\\Windows\\Temp\\UpdateDataDir\\mPortM.reg")
	for i := 0; i < len(args); i++ {
		log.Println(args[i])
	}
}

func TestLaunchInActiveSesstion() {
	_, _, err := win.LaunchInActiveSesstion(`notepad.exe`)
	if err != nil {
		log.Println(err)
	}
}
