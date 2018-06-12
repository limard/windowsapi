package windowsapi

import (
	"syscall"
	"unsafe"
)

var (
	dshell32 = syscall.NewLazyDLL("Shell32.dll")

	pSHGetFolderPath        = dshell32.NewProc("SHGetFolderPathW")
	pSHGetSpecialFolderPath = dshell32.NewProc("SHGetSpecialFolderPathW")
)

const (
	CSIDL_DESKTOP                 = 0x00
	CSIDL_INTERNET                = 0x01
	CSIDL_PROGRAMS                = 0x02
	CSIDL_CONTROLS                = 0x03
	CSIDL_PRINTERS                = 0x04
	CSIDL_PERSONAL                = 0x05
	CSIDL_FAVORITES               = 0x06
	CSIDL_STARTUP                 = 0x07
	CSIDL_RECENT                  = 0x08
	CSIDL_SENDTO                  = 0x09
	CSIDL_BITBUCKET               = 0x0A
	CSIDL_STARTMENU               = 0x0B
	CSIDL_MYDOCUMENTS             = 0x0C
	CSIDL_MYMUSIC                 = 0x0D
	CSIDL_MYVIDEO                 = 0x0E
	CSIDL_DESKTOPDIRECTORY        = 0x10
	CSIDL_DRIVES                  = 0x11
	CSIDL_NETWORK                 = 0x12
	CSIDL_NETHOOD                 = 0x13
	CSIDL_FONTS                   = 0x14
	CSIDL_TEMPLATES               = 0x15
	CSIDL_COMMON_STARTMENU        = 0x16
	CSIDL_COMMON_PROGRAMS         = 0x17
	CSIDL_COMMON_STARTUP          = 0x18
	CSIDL_COMMON_DESKTOPDIRECTORY = 0x19
	CSIDL_APPDATA                 = 0x1A
	CSIDL_PRINTHOOD               = 0x1B
	CSIDL_LOCAL_APPDATA           = 0x1C
	CSIDL_ALTSTARTUP              = 0x1D
	CSIDL_COMMON_ALTSTARTUP       = 0x1E
	CSIDL_COMMON_FAVORITES        = 0x1F
	CSIDL_INTERNET_CACHE          = 0x20
	CSIDL_COOKIES                 = 0x21
	CSIDL_HISTORY                 = 0x22
	CSIDL_COMMON_APPDATA          = 0x23
	CSIDL_WINDOWS                 = 0x24
	CSIDL_SYSTEM                  = 0x25
	CSIDL_PROGRAM_FILES           = 0x26
	CSIDL_MYPICTURES              = 0x27
	CSIDL_PROFILE                 = 0x28
	CSIDL_SYSTEMX86               = 0x29
	CSIDL_PROGRAM_FILESX86        = 0x2A
	CSIDL_PROGRAM_FILES_COMMON    = 0x2B
	CSIDL_PROGRAM_FILES_COMMONX86 = 0x2C
	CSIDL_COMMON_TEMPLATES        = 0x2D
	CSIDL_COMMON_DOCUMENTS        = 0x2E
	CSIDL_COMMON_ADMINTOOLS       = 0x2F
	CSIDL_ADMINTOOLS              = 0x30
	CSIDL_CONNECTIONS             = 0x31
	CSIDL_COMMON_MUSIC            = 0x35
	CSIDL_COMMON_PICTURES         = 0x36
	CSIDL_COMMON_VIDEO            = 0x37
	CSIDL_RESOURCES               = 0x38
	CSIDL_RESOURCES_LOCALIZED     = 0x39
	CSIDL_COMMON_OEM_LINKS        = 0x3A
	CSIDL_CDBURN_AREA             = 0x3B
	CSIDL_COMPUTERSNEARME         = 0x3D
	CSIDL_FLAG_CREATE             = 0x8000
	CSIDL_FLAG_DONT_VERIFY        = 0x4000
	CSIDL_FLAG_NO_ALIAS           = 0x1000
	CSIDL_FLAG_PER_USER_INIT      = 0x8000
	CSIDL_FLAG_MASK               = 0xFF00
)

func SHGetFolderPath(nFolder int) (string, error) {
	if err := pSHGetFolderPath.Find(); err != nil {
		return "", err
	}

	pt := make([]uint16, syscall.MAX_PATH)
	ret, _, err := pSHGetFolderPath.Call(
		uintptr(0),
		uintptr(nFolder),
		uintptr(0), // token
		uintptr(0), // dwFlags
		uintptr(unsafe.Pointer(&pt[0])))
	if ret == 0 {
		err = nil
	}

	return syscall.UTF16ToString(pt), err
}

func SHGetSpecialFolderPath(nFolder int) (string, error) {
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
