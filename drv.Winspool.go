package windowsapi

import (
	"syscall"
	"unsafe"
)

var (
	dWinspool = syscall.NewLazyDLL("Winspool.drv")

	pGetPrintProcessorDirectoryW = dWinspool.NewProc("GetPrintProcessorDirectoryW")
	pGetPrinterDriverDirectoryW  = dWinspool.NewProc("GetPrinterDriverDirectoryW")
)

// GetPrintProcessorDirectory is used for get print processor directory
func GetPrintProcessorDirectory(platform string) (path string, err error) {
	if err = pGetPrintProcessorDirectoryW.Find(); err != nil {
		return "", err
	}

	pt := make([]uint16, syscall.MAX_PATH)
	num := 0
	ptrPlatform, _ := syscall.UTF16PtrFromString(platform)
	ret, _, err := pGetPrintProcessorDirectoryW.Call(
		uintptr(0),
		uintptr(unsafe.Pointer(ptrPlatform)),
		uintptr(1),
		uintptr(unsafe.Pointer(&pt[0])),
		uintptr(256),
		uintptr(unsafe.Pointer(&num)))
	if ret != 0 {
		err = nil
	}

	return syscall.UTF16ToString(pt), err
}

// GetPrinterDriverDirectory ...
func GetPrinterDriverDirectory(platform string) (path string, err error) {
	if err = pGetPrinterDriverDirectoryW.Find(); err != nil {
		return "", err
	}

	pt := make([]uint16, syscall.MAX_PATH)
	num := 0
	ptrPlatform, _ := syscall.UTF16PtrFromString(platform)
	ret, _, err := pGetPrinterDriverDirectoryW.Call(
		uintptr(0),
		uintptr(unsafe.Pointer(ptrPlatform)),
		uintptr(1),
		uintptr(unsafe.Pointer(&pt[0])),
		uintptr(256),
		uintptr(unsafe.Pointer(&num)))
	if ret != 0 {
		err = nil
	}

	return syscall.UTF16ToString(pt), err
}
