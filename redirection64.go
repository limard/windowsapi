// +build windows

package windowsapi

import (
	"runtime"
	"unsafe"
)

func Wow64DisableWow64FsRedirection() (oldvalue uintptr, err error) {
	// This function is useful for 32-bit applications that want to gain access to the native system32 directory.
	if runtime.GOARCH == "amd64" {
		return
	}

	if err = pWow64DisableWow64FsRedirection.Find(); err != nil {
		return
	}
	//log.Printf("%v", unsafe.Sizeof(oldvalue))
	ret, _, err := pWow64DisableWow64FsRedirection.Call(uintptr(unsafe.Pointer(&oldvalue)))
	if ret == 0 {
		return
	}

	// If the function succeeds, the return value is a nonzero value.
	err = nil
	return
}

func Wow64EnableWow64FsRedirection(enable uint) (err error) {
	if runtime.GOARCH == "amd64" {
		return
	}

	if err = pWow64EnableWow64FsRedirection.Find(); err != nil {
		return
	}

	ret, _, err := pWow64EnableWow64FsRedirection.Call(uintptr(enable))
	if ret == 0 {
		return
	}

	return nil
}

func Wow64RevertWow64FsRedirection(oldValue uintptr) (err error) {
	if runtime.GOARCH == "amd64" {
		return
	}

	if err = pWow64RevertWow64FsRedirection.Find(); err != nil {
		return
	}

	ret, _, err := pWow64RevertWow64FsRedirection.Call(oldValue)
	if ret == 0 {
		return
	}

	return nil
}
