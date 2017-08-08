package windowsapi

import (
	"syscall"
	"unsafe"
)

func WTSQueryUserToken(sessionID uint32) (hd syscall.Token, err error) {
	proc, err := loadProc("Wtsapi32.dll", "WTSQueryUserToken")
	if err != nil {
		return
	}
	r1, _, err := proc.Call(
		uintptr(sessionID),
		uintptr(unsafe.Pointer(&hd)),
	)
	if r1 == 1 {
		err = nil
	}
	return
}