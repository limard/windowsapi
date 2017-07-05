package windowsapi

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	dNtdll = syscall.NewLazyDLL("ntdll.dll")

	pRtlGetVersion = dNtdll.NewProc("RtlGetVersion")
)

type OSVERSIONINFO struct {
	dwOSVersionInfoSize uint32
	dwMajorVersion      uint32
	dwMinorVersion      uint32
	dwBuildNumber       uint32
	dwPlatformId        uint32
	szCSDVersion        [128]uint16
}

func RtlGetVersion() (OSVERSIONINFO, error) {
	info := OSVERSIONINFO{}
	info.dwOSVersionInfoSize = uint32(unsafe.Sizeof(info))

	ret, _, err := pRtlGetVersion.Call(uintptr(unsafe.Pointer(&info)))
	if ret > 0xC0000000 {
		fmt.Printf("RtlGetVersion failed: " + err.Error())
	}

	return info, nil
}
