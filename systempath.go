package windowsapi

import (
	"syscall"
	"unsafe"
)

// GetPrintProcessorDirectory is used for get print processor directory
func GetPrintProcessorDirectory(platform string) (path string, err error) {
	pt := make([]uint16, syscall.MAX_PATH)
	num := 0
	ret, _, err := pGetPrintProcessorDirectoryW.Call(
		uintptr(0),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(platform))),
		uintptr(1),
		uintptr(unsafe.Pointer(&pt[0])),
		uintptr(256),
		uintptr(unsafe.Pointer(&num)))
	if ret != 0 {
		err = nil
	}

	return syscall.UTF16ToString(pt), err
}

// C:\Windows\System32\spool\prtprocs\x64
func GetPrintProcessorDirectory64() (path string, err error) {
	return GetPrintProcessorDirectory("Windows x64")
}

// C:\Windows\System32\spool\prtprocs\x86
func GetPrintProcessorDirectory86() (path string, err error) {
	return GetPrintProcessorDirectory("Windows NT x86")
}

////////////////////////////////////////////////////////

// GetPrinterDriverDirectory ...
func GetPrinterDriverDirectory(platform string) (path string, err error) {
	pt := make([]uint16, syscall.MAX_PATH)
	num := 0
	ret, _, err := pGetPrinterDriverDirectoryW.Call(
		uintptr(0),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(platform))),
		uintptr(1),
		uintptr(unsafe.Pointer(&pt[0])),
		uintptr(256),
		uintptr(unsafe.Pointer(&num)))
	if ret != 0 {
		err = nil
	}

	return syscall.UTF16ToString(pt), err
}

func GetPrinterDriverDirectory64() (path string, err error) {
	return GetPrinterDriverDirectory("Windows x64")
}

func GetPrinterDriverDirectory86() (path string, err error) {
	return GetPrinterDriverDirectory("Windows NT x86")
}

////////////////////////////////////////////////////////

// GetSystemDirectory get C:\Windows\System32
func GetSystemDirectory() (path string, err error) {
	pt := make([]uint16, syscall.MAX_PATH)
	num := 0
	ret, _, err := pGetSystemDirectoryW.Call(uintptr(unsafe.Pointer(&pt[0])), uintptr(unsafe.Pointer(&num)))
	if ret != 0 {
		err = nil
	}

	return syscall.UTF16ToString(pt), err
}

////////////////////////////////////////////////////////

func shGetFolderPath(nFolder int) (string, error) {
	if err := pSHGetFolderPathW.Find(); err != nil {
		return "", err
	}

	pt := make([]uint16, syscall.MAX_PATH)
	ret, _, err := pSHGetFolderPathW.Call(
		uintptr(0),
		uintptr(nFolder),
		uintptr(0), // token
		uintptr(0), // dwFlags
		uintptr(unsafe.Pointer(&pt[0])))
	if ret == 0 {
		// err = nil
	}

	return syscall.UTF16ToString(pt), err
}

func shGetSpecialFolderPath(nFolder int) (string, error) {
	if err := pSHGetSpecialFolderPath.Find(); err != nil {
		return "", err
	}
	pt := make([]uint16, syscall.MAX_PATH)
	ret, _, err := pSHGetSpecialFolderPath.Call(
		uintptr(0),
		uintptr(unsafe.Pointer(&pt[0])),
		uintptr(nFolder),
		uintptr(1))
	if ret != 0 {
		err = nil
	}

	return syscall.UTF16ToString(pt), err
}

// CSIDL_APPDATA: C:\Documents and Settings\username\Application Data
func GetAppDataDirectory() (string, error) {
	return shGetSpecialFolderPath(CSIDL_APPDATA)
}

// CSIDL_LOCAL_APPDATA: C:\Documents and Settings\username\Local Settings\Application Data.
func GetLocalAppDataDirectory() (string, error) {
	return shGetSpecialFolderPath(CSIDL_LOCAL_APPDATA)
}

// CSIDL_COMMON_APPDATA: "C:\Documents and Settings\All Users\Application Data" or "C:\ProgramData"
func GetCommmonAppDataDirectory() (string, error) {
	return shGetSpecialFolderPath(CSIDL_COMMON_APPDATA)
}

// // C:\Documents and Settings\...\Desktop
func GetDesktopDir() (string, error) {
	return shGetFolderPath(CSIDL_DESKTOP)
}

// C:\Documents and Settings\All Users\Desktop
func GetCommonDesktopDir() (string, error) {
	return shGetFolderPath(CSIDL_COMMON_DESKTOPDIRECTORY)
}

// C:\Windows
func GetWindowsDir() (string, error) {
	return shGetFolderPath(CSIDL_WINDOWS)
}

// C:\Windows\System32
func GetSystemDir() (string, error) {
	return shGetFolderPath(CSIDL_SYSTEM)
}

// C:\Windows\SysWOW64
func GetSystem86Dir() (string, error) {
	return shGetFolderPath(CSIDL_SYSTEMX86)
}

// C:\Program Files
func GetProgramFilesDir() (string, error) {
	return shGetFolderPath(CSIDL_PROGRAM_FILES)
}

// C:\Program Files (x86)
func GetProgramFiles86Dir() (string, error) {
	return shGetFolderPath(CSIDL_PROGRAM_FILESX86)
}

//  C:\Documents and Settings\username\Templates
func GetUserTempSystemDir() (string, error) {
	return shGetFolderPath(CSIDL_TEMPLATES)
}

// C:\Program Files\Common
func GetProgramFilesCommonDir() (string, error) {
	return shGetFolderPath(CSIDL_PROGRAM_FILES_COMMON)
}

// C:\Program Files (x86)\Common
func GetProgramFilesCommon86Dir() (string, error) {
	return shGetFolderPath(CSIDL_PROGRAM_FILES_COMMONX86)
}

func GetTempPath() (string, error) {
	pt := make([]uint16, syscall.MAX_PATH)
	ret, _, err := pGetTempPathW.Call(syscall.MAX_PATH, uintptr(unsafe.Pointer(&pt[0])))
	if ret != 0 {
		err = nil
	}

	return syscall.UTF16ToString(pt), err
}
