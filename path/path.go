package path

import (
	"log"
	"strings"
	"syscall"
	"unsafe"
)

// GetPrintProcessorDirectory is used for get print processor directory
func GetPrintProcessorDirectory(platform string) (path string) {
	// log.Printf("platform: %#+v\n", platform)
	d := syscall.NewLazyDLL("Winspool.drv")
	p := d.NewProc("GetPrintProcessorDirectoryA")

	pt := make([]byte, 256)
	num := 0
	ret, _, err := p.Call(
		uintptr(0),
		uintptr(unsafe.Pointer(syscall.StringBytePtr(platform))),
		uintptr(1),
		uintptr(unsafe.Pointer(&pt[0])),
		uintptr(256),
		uintptr(unsafe.Pointer(&num)))

	if strings.Contains(err.Error(), "successfully") == false {
		log.Printf("ret: %#+v\n", ret)
		log.Printf("err: %#+v\n", err.Error())
		return
	}

	path = strings.Trim(string(pt[0:num/2]), string(0))

	return path
}

// GetPrinterDriverDirectory ...
func GetPrinterDriverDirectory(platform string) (path string) {
	// log.Printf("platform: %#+v\n", platform)
	d := syscall.NewLazyDLL("Winspool.drv")
	p := d.NewProc("GetPrinterDriverDirectoryA")

	pt := make([]byte, 256)
	num := 0
	ret, _, err := p.Call(
		uintptr(0),
		uintptr(unsafe.Pointer(syscall.StringBytePtr(platform))),
		uintptr(1),
		uintptr(unsafe.Pointer(&pt[0])),
		uintptr(256),
		uintptr(unsafe.Pointer(&num)))

	if strings.Contains(err.Error(), "successfully") == false {
		log.Printf("ret: %#+v\n", ret)
		log.Printf("err: %#+v\n", err.Error())
		return
	}

	path = strings.Trim(string(pt[0:num/2]), string(0))

	return path
}

// GetSystemDirectory ...
func GetSystemDirectory() (path string) {
	d := syscall.NewLazyDLL("Kernel32.dll")
	p := d.NewProc("GetSystemDirectoryA")

	pt := make([]byte, 256)
	num := 0
	ret, _, err := p.Call(uintptr(unsafe.Pointer(&pt[0])), uintptr(unsafe.Pointer(&num)))

	if strings.Contains(err.Error(), "successfully") == false {
		log.Printf("ret: %#+v\n", ret)
		log.Printf("err: %#+v\n", err.Error())
		return
	}

	path = strings.Trim(string(pt), string(0))

	return
}
