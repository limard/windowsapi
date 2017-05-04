package win

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"
	//"log"
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
	dwOSVersionInfoSize uint32
	dwMajorVersion      uint32
	dwMinorVersion      uint32
	dwBuildNumber       uint32
	dwPlatformId        uint32
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

const (
	VER_NT_WORKSTATION       = 0x1
	VER_NT_DOMAIN_CONTROLLER = 0x2
	VER_NT_SERVER            = 0x3
)

const (
	VER_BUILDNUMBER      = 0x0000004
	VER_MAJORVERSION     = 0x0000002
	VER_MINORVERSION     = 0x0000001
	VER_PLATFORMID       = 0x0000008
	VER_PRODUCT_TYPE     = 0x0000080
	VER_SERVICEPACKMAJOR = 0x0000020
	VER_SERVICEPACKMINOR = 0x0000010
	VER_SUITENAME        = 0x0000040

	VER_EQUAL         = 1
	VER_GREATER       = 2
	VER_GREATER_EQUAL = 3
	VER_LESS          = 4
	VER_LESS_EQUAL    = 5

	ERROR_OLD_WIN_VERSION syscall.Errno = 1150
)

func getOSVersion_back() (string, uint32, uint32) {
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
			if b, _ := equalOSVersion(6, 2); b {
				return "Windows 8", 6, 2
			}
			if b, _ := equalOSVersion(6, 3); b {
				return "Windows 8.1", 6, 2
			}
			if b, _ := equalOSVersion(10, 0); b {
				return "Windows 10", 6, 2
			}
			version = "Windows 8"

		case os.dwMajorVersion == 6 && os.dwMinorVersion == 2 && os.wProductType != VER_NT_WORKSTATION:
			if b, _ := equalOSVersion(6, 2); b {
				return "Windows Server 2012", 6, 2
			}
			if b, _ := equalOSVersion(6, 3); b {
				return "Windows Server 2012 R2", 6, 2
			}
			if b, _ := equalOSVersion(10, 0); b {
				return "Windows Server 2016", 6, 2
			}
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
			return "Unknown Version", os.dwMajorVersion, os.dwMinorVersion
		}
	}
	return version, os.dwMajorVersion, os.dwMinorVersion
}

func GetOSVersion() (string, uint32, uint32) {
	var version string = "Unknown Version"

	kernel32 := syscall.NewLazyDLL("kernel32.dll")

	var os OSVERSIONINFOEX
	os.dwOSVersionInfoSize = uint32(unsafe.Sizeof(os))

	GetVersionExW := kernel32.NewProc("GetVersionExW")
	rt, _, _ := GetVersionExW.Call(uintptr(unsafe.Pointer(&os)))
	if rt == 0 {
		return "Unknown Version", 0, 0
	}

	os.dwMajorVersion, os.dwMinorVersion, _ = rtlGetVersion()

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
		return "Unknown Version", os.dwMajorVersion, os.dwMinorVersion
	}

	return version, os.dwMajorVersion, os.dwMinorVersion
}

func equalOSVersion(dwMajorVersion, dwMinorVersion uint32) (bool, error) {
	d := syscall.NewLazyDLL("kernel32.dll")

	var m1, m2 uintptr
	VerSetConditionMask := d.NewProc("VerSetConditionMask")
	m1, m2, _ = VerSetConditionMask.Call(m1, m2, VER_MAJORVERSION, VER_EQUAL)
	m1, m2, _ = VerSetConditionMask.Call(m1, m2, VER_MINORVERSION, VER_EQUAL)

	//log.Printf("%#v%#v", m1, m2)

	vi := OSVERSIONINFOEX{
		dwMajorVersion: dwMajorVersion,
		dwMinorVersion: dwMinorVersion,
	}
	vi.dwOSVersionInfoSize = uint32(unsafe.Sizeof(vi))
	r, _, e1 := d.NewProc("VerifyVersionInfoW").Call(
		uintptr(unsafe.Pointer(&vi)),
		VER_MAJORVERSION|VER_MINORVERSION,
		m1, m2)
	if r == 1 {
		return true, nil
	}

	if r == 0 && e1 == ERROR_OLD_WIN_VERSION {
		return false, nil
	}

	return false, fmt.Errorf("VerifyVersionInfo failed: %s", e1)
}

func isOSWorkstation() (bool, error) {
	d := syscall.NewLazyDLL("kernel32.dll")

	var m1, m2 uintptr
	VerSetConditionMask := d.NewProc("VerSetConditionMask")
	m1, m2, _ = VerSetConditionMask.Call(m1, m2, VER_MAJORVERSION, VER_EQUAL)
	m1, m2, _ = VerSetConditionMask.Call(m1, m2, VER_MINORVERSION, VER_EQUAL)
	m1, m2, _ = VerSetConditionMask.Call(m1, m2, VER_PRODUCT_TYPE, VER_EQUAL)

	//log.Printf("%#v", unsafe.Sizeof(m1))
	//log.Printf("%#v%#v", m1, m2)

	vi := OSVERSIONINFOEX{
		dwMajorVersion: 6,
		dwMinorVersion: 1,
		wProductType:   VER_NT_WORKSTATION,
	}
	vi.dwOSVersionInfoSize = uint32(unsafe.Sizeof(vi))

	//log.Printf("%#v", vi)

	r, _, e1 := d.NewProc("VerifyVersionInfoW").Call(
		uintptr(unsafe.Pointer(&vi)),
		//VER_PRODUCT_TYPE|
		VER_MAJORVERSION|
			VER_MINORVERSION,
		m1, m2)
	if r == 1 {
		return true, nil
	}

	if r == 0 && e1 == ERROR_OLD_WIN_VERSION {
		return false, nil
	}

	return false, fmt.Errorf("VerifyVersionInfo failed: %s", e1)
}

func rtlGetVersion() (uint32, uint32, error) {
	dll := syscall.NewLazyDLL("ntdll.dll")
	proc := dll.NewProc("RtlGetVersion")
	info := &OSVERSIONINFO{}
	info.dwOSVersionInfoSize = uint32(unsafe.Sizeof(info))

	ret, _, err := proc.Call(uintptr(unsafe.Pointer(info)))
	if ret > 0xC0000000 {
		fmt.Println(errors.New("RtlGetVersion failed: " + err.Error()))
	}
	return info.dwMajorVersion, info.dwMinorVersion, nil
}
