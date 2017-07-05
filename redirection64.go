// +build windows

package windowsapi

import (
	"log"
	"unsafe"
)

func Wow64DisableWow64FsRedirection() (oldvalue uintptr, err error) {
	if err = pWow64DisableWow64FsRedirection.Find(); err != nil {
		return
	}

	if Is64bitOS() == false {
		return oldvalue, nil
	}

	ret, _, err := pWow64DisableWow64FsRedirection.Call(uintptr(unsafe.Pointer(&oldvalue)))
	if ret == 0 {
		log.Println("ERROR (Wow64DisableWow64FsRedirection):", err.Error())
		return
	}

	return oldvalue, nil
}

func Wow64EnableWow64FsRedirection(enable uint) (err error) {
	if err = pWow64EnableWow64FsRedirection.Find(); err != nil {
		return
	}

	if Is64bitOS() == false {
		return nil
	}

	ret, _, err := pWow64EnableWow64FsRedirection.Call(uintptr(enable))
	if ret == 0 {
		log.Println("ERROR (Wow64EnableWow64FsRedirection):", err.Error())
		return
	}

	return nil
}

func Wow64RevertWow64FsRedirection(oldValue uintptr) (err error) {
	if err = pWow64RevertWow64FsRedirection.Find(); err != nil {
		return
	}

	// if osinfo.Is64bitOS() == false {
	// 	return nil
	// }

	ret, _, err := pWow64RevertWow64FsRedirection.Call(oldValue)
	if ret == 0 {
		log.Println("ERROR (Wow64RevertWow64FsRedirection):", err.Error())
		return
	}

	return nil
}
