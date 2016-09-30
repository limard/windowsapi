package main

import (
	"log"

	"bitbucket.org/Limard/win/osinfo"
	"bitbucket.org/Limard/win/systempath"
)

func main() {
	log.Println("osinfo.Is64bitOS()", osinfo.Is64bitOS())
	log.Println("systempath.GetCommmonAppDataDirectory()", systempath.GetCommmonAppDataDirectory())
	log.Println("systempath.GetCommonDesktopDir()", systempath.GetCommonDesktopDir())
}
