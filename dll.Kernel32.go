package windowsapi

import "syscall"

var (
	dKernel32                       = syscall.NewLazyDLL("Kernel32.dll")

	pWow64DisableWow64FsRedirection = dKernel32.NewProc("Wow64DisableWow64FsRedirection")
	pWow64EnableWow64FsRedirection  = dKernel32.NewProc("Wow64EnableWow64FsRedirection")
	pWow64RevertWow64FsRedirection  = dKernel32.NewProc("Wow64RevertWow64FsRedirection")

	pGetSystemDirectoryW            = dKernel32.NewProc("GetSystemDirectoryW")
	pGetTempPathW                   = dKernel32.NewProc("GetTempPathW")

	pGetNativeSystemInfo            = dKernel32.NewProc("GetNativeSystemInfo")
	pGetVersionExW                  = dKernel32.NewProc("GetVersionExW")
	pVerSetConditionMask            = dKernel32.NewProc("VerSetConditionMask")
	pVerifyVersionInfo              = dKernel32.NewProc("VerifyVersionInfoW")

	pWTSGetActiveConsoleSessionId 	= dKernel32.NewProc("WTSGetActiveConsoleSessionId")
)

const (
	PROCESSOR_ARCHITECTURE_AMD64   = 9
	PROCESSOR_ARCHITECTURE_ARM     = 5
	PROCESSOR_ARCHITECTURE_IA64    = 6
	PROCESSOR_ARCHITECTURE_INTEL   = 0
	PROCESSOR_ARCHITECTURE_UNKNOWN = 0xffff
)
