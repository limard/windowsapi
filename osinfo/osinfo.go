package osinfo

import (
	"syscall"
	"unsafe"
)

type SSystemInfo struct {
	wProcessorArchitecture      uint16
	wReserved                   uint16
	dwPageSize                  uint32
	lpMinimumApplicationAddress uintptr
	lpMaximumApplicationAddress uintptr
	dwActiveProcessorMask       uintptr
	dwNumberOfProcessors        uint32
	dwProcessorType             uint32
	dwAllocationGranularity     uint32
	wProcessorLevel             uint16
	wProcessorRevision          uint16
}

const (
	PROCESSOR_ARCHITECTURE_AMD64   = 9
	PROCESSOR_ARCHITECTURE_ARM     = 5
	PROCESSOR_ARCHITECTURE_IA64    = 6
	PROCESSOR_ARCHITECTURE_INTEL   = 0
	PROCESSOR_ARCHITECTURE_UNKNOWN = 0xffff
)

func Is64bitOS() bool {
	d := syscall.NewLazyDLL("kernel32.dll")
	p := d.NewProc("GetNativeSystemInfo")

	if err := p.Find(); err != nil {
		return false
	}

	var info = SSystemInfo{}

	ret, _, _ := p.Call(uintptr(unsafe.Pointer(&info)))
	if ret == 0 {
		return false
	}

	if info.wProcessorArchitecture == PROCESSOR_ARCHITECTURE_AMD64 ||
		info.wProcessorArchitecture == PROCESSOR_ARCHITECTURE_IA64 {
		// log.Println("wProcessorArchitecture", info.wProcessorArchitecture)
		return true
	}

	return false
}
