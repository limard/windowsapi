// +build windows

package windowsapi

import (
	"syscall"
	"unsafe"
)

var (
	dKernel32 = syscall.NewLazyDLL("Kernel32.dll")

	pWow64DisableWow64FsRedirection = dKernel32.NewProc("Wow64DisableWow64FsRedirection")
	pWow64EnableWow64FsRedirection  = dKernel32.NewProc("Wow64EnableWow64FsRedirection")
	pWow64RevertWow64FsRedirection  = dKernel32.NewProc("Wow64RevertWow64FsRedirection")

	pGetSystemDirectoryW = dKernel32.NewProc("GetSystemDirectoryW")
	pGetTempPathW        = dKernel32.NewProc("GetTempPathW")

	pGetNativeSystemInfo = dKernel32.NewProc("GetNativeSystemInfo")
	pGetVersionExW       = dKernel32.NewProc("GetVersionExW")
	pVerSetConditionMask = dKernel32.NewProc("VerSetConditionMask")
	pVerifyVersionInfo   = dKernel32.NewProc("VerifyVersionInfoW")

	pWTSGetActiveConsoleSessionId = dKernel32.NewProc("WTSGetActiveConsoleSessionId")
)

const (
	PROCESSOR_ARCHITECTURE_AMD64   = 9
	PROCESSOR_ARCHITECTURE_ARM     = 5
	PROCESSOR_ARCHITECTURE_IA64    = 6
	PROCESSOR_ARCHITECTURE_INTEL   = 0
	PROCESSOR_ARCHITECTURE_UNKNOWN = 0xffff
)

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

func GetTempPath() (string, error) {
	pt := make([]uint16, syscall.MAX_PATH)
	ret, _, err := pGetTempPathW.Call(syscall.MAX_PATH, uintptr(unsafe.Pointer(&pt[0])))
	if ret != 0 {
		err = nil
	}

	return syscall.UTF16ToString(pt), err
}

func WTSGetActiveConsoleSessionId() (sessionId uint32, err error) {
	r1, _, err := pWTSGetActiveConsoleSessionId.Call()
	if r1 == 0xFFFFFFFF {
		return
	}

	sessionId = uint32(r1)
	err = nil
	return
}
