package systempath

import (
	"log"
	"strings"
	"syscall"
	"unsafe"

	"github.com/lxn/win"
)

// GetPrintProcessorDirectory is used for get print processor directory
func GetPrintProcessorDirectory(platform string) (path string) {
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
	if ret == 0 {
		log.Printf("err: %#+v\n", err.Error())
		return path
	}

	path = strings.Trim(string(pt[0:num/2]), string(0))

	return path
}

// GetPrintProcessorDirectory64 ...
func GetPrintProcessorDirectory64() (path string) {
	return GetPrintProcessorDirectory("Windows x64")
}

// GetPrintProcessorDirectory86 ...
func GetPrintProcessorDirectory86() (path string) {
	return GetPrintProcessorDirectory("Windows x86")
}

////////////////////

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
	if ret == 0 {
		log.Printf("err: %#+v\n", err.Error())
		return path
	}

	path = strings.Trim(string(pt[0:num/2]), string(0))

	return path
}

func GetPrinterDriverDirectory64() (path string) {
	return GetPrinterDriverDirectory("Windows x64")
}

func GetPrinterDriverDirectory86() (path string) {
	return GetPrinterDriverDirectory("Windows x86")
}

// GetSystemDirectory ...
func GetSystemDirectory() (path string) {
	d := syscall.NewLazyDLL("Kernel32.dll")
	p := d.NewProc("GetSystemDirectoryA")

	pt := make([]byte, 256)
	num := 0
	ret, _, err := p.Call(uintptr(unsafe.Pointer(&pt[0])), uintptr(unsafe.Pointer(&num)))
	if ret == 0 {
		log.Printf("err: %#+v\n", err.Error())
		return path
	}

	path = strings.Trim(string(pt), string(0))

	return
}

////////////////////////////////////////////////////////

func shGetFolderPath(nFolder int) string {
	d := syscall.NewLazyDLL("Shell32.dll")
	p := d.NewProc("SHGetFolderPathA")

	if err := p.Find(); err != nil {
		log.Println("ERROR", err.Error())
		return ""
	}

	pt := make([]byte, 256)
	ret, _, err := p.Call(uintptr(0), uintptr(nFolder), uintptr(0), uintptr(0), uintptr(unsafe.Pointer(&pt[0])))
	if ret == 0 {
		log.Println("ERROR:", err.Error())
	}

	path := strings.Trim(string(pt), string(0))
	return path
}

// GetCommmonAppDataDirectory ...
func GetCommmonAppDataDirectory() string {
	return shGetFolderPath(win.CSIDL_COMMON_APPDATA)
}

func GetDesktopDir() string {
	return shGetFolderPath(win.CSIDL_DESKTOP)
}

func GetCommonDesktopDir() string {
	return shGetFolderPath(win.CSIDL_COMMON_DESKTOPDIRECTORY)
}

func GetWindowsDir() string {
	return shGetFolderPath(win.CSIDL_WINDOWS)
}

func GetProgramFilesDir() string {
	return shGetFolderPath(win.CSIDL_SYSTEM)
}

func GetProgramFiles86Dir() string {
	return shGetFolderPath(win.CSIDL_SYSTEMX86)
}

func GetUserTempSystemDir() string {
	return shGetFolderPath(win.CSIDL_TEMPLATES)
}

func GetProgramFilesCommonDir() string {
	return shGetFolderPath(win.CSIDL_PROGRAM_FILES_COMMON)
}

func GetProgramFilesCommon86Dir() string {
	return shGetFolderPath(win.CSIDL_PROGRAM_FILES_COMMONX86)
}
