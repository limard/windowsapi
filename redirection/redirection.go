package redirection

import (
	"log"
	"strings"
	"syscall"
	"unsafe"
)

// Wow64DisableWow64FsRedirection ...
func Wow64DisableWow64FsRedirection() (oldvalue int, err error) {
	d := syscall.NewLazyDLL("Kernel32.dll")
	p := d.NewProc("Wow64DisableWow64FsRedirection")

	ret, _, err := p.Call(
		uintptr(unsafe.Pointer(&oldvalue)))

	if strings.Contains(err.Error(), "successfully") == false {
		log.Printf("ret: %#+v\n", ret)
		log.Printf("err: %#+v\n", err.Error())
		return
	}

	return
}

// Wow64EnableWow64FsRedirection ...
func Wow64EnableWow64FsRedirection(enable int) (err error) {
	d := syscall.NewLazyDLL("Kernel32.dll")
	p := d.NewProc("Wow64EnableWow64FsRedirection")

	ret, _, err := p.Call(
		uintptr(enable))

	if strings.Contains(err.Error(), "successfully") == false {
		log.Printf("ret: %#+v\n", ret)
		log.Printf("err: %#+v\n", err.Error())
		return
	}

	return
}

// Wow64RevertWow64FsRedirection ...
func Wow64RevertWow64FsRedirection(oldValue int) (err error) {
	d := syscall.NewLazyDLL("Kernel32.dll")
	p := d.NewProc("Wow64RevertWow64FsRedirection")

	ret, _, err := p.Call(
		uintptr(oldValue))

	if strings.Contains(err.Error(), "successfully") == false {
		log.Printf("ret: %#+v\n", ret)
		log.Printf("err: %#+v\n", err.Error())
		return
	}

	return
}
