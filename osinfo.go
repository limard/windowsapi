package win

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

// Windows version

type OSVERSIONINFO struct {
	dwOSVersionInfoSize int32
	dwMajorVersion      int32
	dwMinorVersion      int32
	dwBuildNumber       int32
	dwPlatformId        int32
	szCSDVersion        [128]uint16
}

type OSVERSIONINFOEX struct {
	dwOSVersionInfoSize uint32
	dwMajorVersion      uint32
	dwMinorVersion      uint32
	dwBuildNumber       uint32
	dwPlatformId        uint32
	szCSDVersion        [128]uint16
	wServicePackMajor   uint16
	wServicePackMinor   uint16
	wSuiteMask          uint16
	wProductType        uint8
	wReserved           uint8
}

func GetOSVersion() (string, uint32, uint32) {
	const VER_NT_WORKSTATION = 0x1

	var version string = "Unknown Version"
	kernel32 := syscall.NewLazyDLL("kernel32.dll")

	var os OSVERSIONINFOEX
	os.dwOSVersionInfoSize = uint32(unsafe.Sizeof(os))

	GetVersionExW := kernel32.NewProc("GetVersionExW")

	rt, _, _ := GetVersionExW.Call(uintptr(unsafe.Pointer(&os)))
	if int(rt) == 1 {
		switch {
			// 4
		case os.dwMajorVersion == 4 && os.dwMinorVersion == 0 && os.dwPlatformId == 1:
			version = "Windows 95"
		case os.dwMajorVersion == 4 && os.dwMinorVersion == 10:
			version = "Windows 98"
		case os.dwMajorVersion == 4 && os.dwMinorVersion == 90:
			version = "Windows Me"
		case os.dwMajorVersion == 4 && os.dwMinorVersion == 0 && os.dwPlatformId == 2:
			version = "Windows NT4"

			// 5
		case os.dwMajorVersion == 5 && os.dwMinorVersion == 0:
			version = "Windows Server 2000"
		case os.dwMajorVersion == 5 && os.dwMinorVersion == 1:
			version = "Windows XP"
		case os.dwMajorVersion == 5 && os.dwMinorVersion == 2:
			version = "Windows Server 2003"

			// 6
		case os.dwMajorVersion == 6 && os.dwMinorVersion == 0 && os.wProductType == VER_NT_WORKSTATION:
			version = "Windows Vista"
		case os.dwMajorVersion == 6 && os.dwMinorVersion == 0 && os.wProductType != VER_NT_WORKSTATION:
			version = "Windows Server 2008"
		case os.dwMajorVersion == 6 && os.dwMinorVersion == 1 && os.wProductType == VER_NT_WORKSTATION:
			version = "Windows 7"
		case os.dwMajorVersion == 6 && os.dwMinorVersion == 1 && os.wProductType != VER_NT_WORKSTATION:
			version = "Windows Server 2008 R2"
		case os.dwMajorVersion == 6 && os.dwMinorVersion == 2 && os.wProductType == VER_NT_WORKSTATION:
			version = "Windows 8"
		case os.dwMajorVersion == 6 && os.dwMinorVersion == 2 && os.wProductType != VER_NT_WORKSTATION:
			version = "Windows Server 2012"
		case os.dwMajorVersion == 6 && os.dwMinorVersion == 3 && os.wProductType == VER_NT_WORKSTATION:
			version = "Windows 8.1"
		case os.dwMajorVersion == 6 && os.dwMinorVersion == 3 && os.wProductType != VER_NT_WORKSTATION:
			version = "Windows Server 2012 R2"

			// 10
		case os.dwMajorVersion == 10 && os.dwMinorVersion == 0 && os.wProductType == VER_NT_WORKSTATION:
			version = "Windows 10"
		case os.dwMajorVersion == 10 && os.dwMinorVersion == 0 && os.wProductType != VER_NT_WORKSTATION:
			version = "Windows Server 2016"

		default:
			version = "Unknown Version"
		}
	}
	return version, os.dwMajorVersion, os.dwMinorVersion
}
