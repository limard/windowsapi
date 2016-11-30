package main

import (
	"log"

	"bitbucket.org/Limard/win/redirection"
	"bitbucket.org/Limard/win/systempath"
)

func main() {
	o, _ := redirection.Wow64DisableWow64FsRedirection()

	log.Println("GetPrintProcessorDirectory64\t", systempath.GetPrintProcessorDirectory64())
	log.Println("GetPrintProcessorDirectory86\t", systempath.GetPrintProcessorDirectory86())
	log.Println("GetPrinterDriverDirectory64\t", systempath.GetPrinterDriverDirectory64())
	log.Println("GetPrinterDriverDirectory86\t", systempath.GetPrinterDriverDirectory86())
	log.Println("GetSystemDirectory\t", systempath.GetSystemDirectory())
	log.Println("GetCommmonAppDataDirectory\t", systempath.GetCommmonAppDataDirectory())
	log.Println("GetDesktopDir\t", systempath.GetDesktopDir())
	log.Println("GetCommonDesktopDir\t", systempath.GetCommonDesktopDir())
	log.Println("GetWindowsDir\t", systempath.GetWindowsDir())
	log.Println("GetSystemDir\t", systempath.GetSystemDir())
	log.Println("GetSystem86Dir\t", systempath.GetSystem86Dir())
	log.Println("GetProgramFilesDir\t", systempath.GetProgramFilesDir())
	log.Println("GetProgramFiles86Dir\t", systempath.GetProgramFiles86Dir())
	log.Println("GetProgramFilesCommonDir\t", systempath.GetProgramFilesCommonDir())
	log.Println("GetProgramFilesCommon86Dir\t", systempath.GetProgramFilesCommon86Dir())
	log.Println("GetTempDir\t", systempath.GetTempDir())

	redirection.Wow64RevertWow64FsRedirection(o)
}
