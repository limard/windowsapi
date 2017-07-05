package windowsapi

import "syscall"

var (
	dWinspool                    = syscall.NewLazyDLL("Winspool.drv")
	pGetPrintProcessorDirectoryW = dWinspool.NewProc("GetPrintProcessorDirectoryW")
	pGetPrinterDriverDirectoryW  = dWinspool.NewProc("GetPrinterDriverDirectoryW")
)
