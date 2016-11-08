package systempath

import (
	"log"
	"syscall"
	"unsafe"

	"github.com/lxn/win"
)

var (
	dshell32          = syscall.NewLazyDLL("Shell32.dll")
	pSHGetFolderPathW = dshell32.NewProc("SHGetFolderPathW")

	dkernel32            = syscall.NewLazyDLL("Kernel32.dll")
	pGetSystemDirectoryW = dkernel32.NewProc("GetSystemDirectoryW")
	pGetTempPathW        = dkernel32.NewProc("GetTempPathW")

	dWinspool                    = syscall.NewLazyDLL("Winspool.drv")
	pGetPrintProcessorDirectoryW = dWinspool.NewProc("GetPrintProcessorDirectoryW")
	pGetPrinterDriverDirectoryW  = dWinspool.NewProc("GetPrinterDriverDirectoryW")
)

////////////////////////////////////////////////////////

// GetPrintProcessorDirectory is used for get print processor directory
func GetPrintProcessorDirectory(platform string) (path string) {
	pt := make([]uint16, syscall.MAX_PATH)
	num := 0
	ret, _, err := pGetPrintProcessorDirectoryW.Call(
		uintptr(0),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(platform))),
		uintptr(1),
		uintptr(unsafe.Pointer(&pt[0])),
		uintptr(256),
		uintptr(unsafe.Pointer(&num)))
	if ret == 0 {
		if err.Error() != "An attempt was made to reference a token that does not exist." {
			log.Println("ERROR:", err.Error())
		}
		return path
	}

	return syscall.UTF16ToString(pt)
}

// C:\Windows\System32\spool\prtprocs\x64
func GetPrintProcessorDirectory64() (path string) {
	return GetPrintProcessorDirectory("Windows x64")
}

// C:\Windows\System32\spool\prtprocs\x86
func GetPrintProcessorDirectory86() (path string) {
	return GetPrintProcessorDirectory("Windows NT x86")
}

////////////////////////////////////////////////////////

// GetPrinterDriverDirectory ...
func GetPrinterDriverDirectory(platform string) (path string) {
	pt := make([]uint16, syscall.MAX_PATH)
	num := 0
	ret, _, err := pGetPrinterDriverDirectoryW.Call(
		uintptr(0),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(platform))),
		uintptr(1),
		uintptr(unsafe.Pointer(&pt[0])),
		uintptr(256),
		uintptr(unsafe.Pointer(&num)))
	if ret == 0 {
		if err.Error() != "An attempt was made to reference a token that does not exist." {
			log.Println("ERROR:", err.Error())
		}
		return path
	}

	return syscall.UTF16ToString(pt)
}

func GetPrinterDriverDirectory64() (path string) {
	return GetPrinterDriverDirectory("Windows x64")
}

func GetPrinterDriverDirectory86() (path string) {
	return GetPrinterDriverDirectory("Windows NT x86")
}

////////////////////////////////////////////////////////

// GetSystemDirectory ...
func GetSystemDirectory() (path string) {
	pt := make([]uint16, syscall.MAX_PATH)
	num := 0
	ret, _, err := pGetSystemDirectoryW.Call(uintptr(unsafe.Pointer(&pt[0])), uintptr(unsafe.Pointer(&num)))
	if ret == 0 {
		log.Println("ERROR:", err)
		return path
	}

	return syscall.UTF16ToString(pt)
}

////////////////////////////////////////////////////////

func shGetFolderPath(nFolder int) string {
	if err := pSHGetFolderPathW.Find(); err != nil {
		log.Println("ERROR", err.Error())
		return ""
	}

	pt := make([]uint16, syscall.MAX_PATH)
	ret, _, err := pSHGetFolderPathW.Call(uintptr(0), uintptr(nFolder), uintptr(0), uintptr(0), uintptr(unsafe.Pointer(&pt[0])))
	if ret == 0 {
		if err.Error() != "An attempt was made to reference a token that does not exist." {
			log.Println("ERROR:", err.Error())
		}
	}

	return syscall.UTF16ToString(pt)
}

// C:\Documents and Settings\All Users\Application Data
func GetCommmonAppDataDirectory() string {
	return shGetFolderPath(win.CSIDL_COMMON_APPDATA)
}

// // C:\Documents and Settings\...\Desktop
func GetDesktopDir() string {
	return shGetFolderPath(win.CSIDL_DESKTOP)
}

// C:\Documents and Settings\All Users\Desktop
func GetCommonDesktopDir() string {
	return shGetFolderPath(win.CSIDL_COMMON_DESKTOPDIRECTORY)
}

// C:\Windows
func GetWindowsDir() string {
	return shGetFolderPath(win.CSIDL_WINDOWS)
}

// C:\Windows\System32
func GetSystemDir() string {
	return shGetFolderPath(win.CSIDL_SYSTEM)
}

// C:\Windows\SysWOW64
func GetSystem86Dir() string {
	return shGetFolderPath(win.CSIDL_SYSTEMX86)
}

// C:\Program Files
func GetProgramFilesDir() string {
	return shGetFolderPath(win.CSIDL_PROGRAM_FILES)
}

// C:\Program Files (x86)
func GetProgramFiles86Dir() string {
	return shGetFolderPath(win.CSIDL_PROGRAM_FILESX86)
}

//  C:\Documents and Settings\username\Templates
func GetUserTempSystemDir() string {
	return shGetFolderPath(win.CSIDL_TEMPLATES)
}

// C:\Program Files\Common
func GetProgramFilesCommonDir() string {
	return shGetFolderPath(win.CSIDL_PROGRAM_FILES_COMMON)
}

// C:\Program Files (x86)\Common
func GetProgramFilesCommon86Dir() string {
	return shGetFolderPath(win.CSIDL_PROGRAM_FILES_COMMONX86)
}

func GetTempDir() string {
	pt := make([]uint16, syscall.MAX_PATH)
	ret, _, err := pGetTempPathW.Call(syscall.MAX_PATH, uintptr(unsafe.Pointer(&pt[0])))
	if ret == 0 {
		log.Println("ERROR", err)
	}

	return syscall.UTF16ToString(pt)
}
