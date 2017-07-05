package windowsapi

import "syscall"

var (
	dNtdll = syscall.NewLazyDLL("ntdll.dll")
	pRtlGetVersion = dNtdll.NewProc("RtlGetVersion")
)
