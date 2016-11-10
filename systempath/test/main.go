package main

import (
	"log"

	"bitbucket.org/Limard/win/redirection"
	"bitbucket.org/Limard/win/systempath"
)

func main() {
	o, _ := redirection.Wow64DisableWow64FsRedirection()

	log.Println("GetPrintProcessorDirectory64")
	log.Println(systempath.GetPrintProcessorDirectory64())

	log.Println("GetPrintProcessorDirectory86")
	log.Println(systempath.GetPrintProcessorDirectory86())

	log.Println("GetPrinterDriverDirectory64")
	log.Println(systempath.GetPrinterDriverDirectory64())

	log.Println("GetPrinterDriverDirectory86")
	log.Println(systempath.GetPrinterDriverDirectory86())

	log.Println("GetSystemDirectory")
	log.Println(systempath.GetSystemDirectory())

	log.Println("GetCommmonAppDataDirectory")
	log.Println(systempath.GetCommmonAppDataDirectory())

	log.Println("GetDesktopDir")
	log.Println(systempath.GetDesktopDir())

	log.Println("GetCommonDesktopDir")
	log.Println(systempath.GetCommonDesktopDir())

	log.Println("GetWindowsDir")
	log.Println(systempath.GetWindowsDir())

	log.Println("GetSystemDir")
	log.Println(systempath.GetSystemDir())

	log.Println("GetSystem86Dir")
	log.Println(systempath.GetSystem86Dir())

	log.Println("GetProgramFilesDir")
	log.Println(systempath.GetProgramFilesDir())

	log.Println("GetProgramFiles86Dir")
	log.Println(systempath.GetProgramFiles86Dir())

	log.Println("GetProgramFilesCommonDir")
	log.Println(systempath.GetProgramFilesCommonDir())

	log.Println("GetProgramFilesCommon86Dir")
	log.Println(systempath.GetProgramFilesCommon86Dir())

	log.Println("GetTempDir")
	log.Println(systempath.GetTempDir())

	redirection.Wow64RevertWow64FsRedirection(o)
}
