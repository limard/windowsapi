package windowsapi

import (
	"syscall"
	"unsafe"
)

func CreateEnvironmentBlock(token syscall.Token, inherit bool) (*uint16, error) {
	proc, err := loadProc("Userenv.dll", "CreateEnvironmentBlock")
	if err != nil {
		return nil, err
	}

	iInherit := 0
	if inherit {
		iInherit = 1
	}

	var env *uint16

	r1, _, err := proc.Call(
		uintptr(unsafe.Pointer(&env)),
		uintptr(token),
		uintptr(iInherit),
	)

	if r1 == 1 {
		return env, nil
	}
	return nil, err
}