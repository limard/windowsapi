// +build windows

package win

import (
"os"
"strings"
"syscall"
"unsafe"

"golang.org/x/sys/windows/registry"
)

const (
	HWND_BROADCAST   uintptr = 0xffff
	WM_SETTINGCHANGE uint    = 0x1a
)

// SendMessageTimeout Flags
// https://msdn.microsoft.com/en-us/library/windows/desktop/ms644952(v=vs.85).aspx
const (
	SMTO_ABORTIFHUNG        = 0x0002
	SMTO_BLOCK              = 0x0001
	SMTO_NORMAL             = 0x0000
	SMTO_NOTIMEOUTIFNOTHUNG = 0x0008
	SMTO_ERRORONEXIT        = 0x0020
)

var (
	libUser32         = syscall.NewLazyDLL("user32.dll")
	pathRegKey string = "Path"
)

func BroadcastSettingChange() {
	rawParam, _ := syscall.UTF16PtrFromString("ENVIRONMENT")
	param := uintptr(unsafe.Pointer(rawParam))
	sendMessageProcedure := libUser32.NewProc("SendMessageTimeoutW")
	sendMessageProcedure.Call(uintptr(HWND_BROADCAST), uintptr(WM_SETTINGCHANGE), 0, param,
		uintptr(SMTO_ABORTIFHUNG), uintptr(5000), uintptr(0))
}

func GetRegEnvValue(key string) (string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `System\CurrentControlSet\Control\Session Manager\Environment`, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}

	defer k.Close()
	s, _, err := k.GetStringValue(key)
	return s, err
}

func SetRegEnvValue(key string, value string) error {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `System\CurrentControlSet\Control\Session Manager\Environment`, registry.SET_VALUE)
	if err != nil {
		return err
	}

	defer k.Close()
	return k.SetStringValue(key, value)
}

func AddPathToEnv(newPath string) error {
	v, err := GetRegEnvValue(pathRegKey)
	if err != nil {
		return err
	}

	v = v + string(os.PathListSeparator) + newPath
	// TODO: do logging here
	return SetRegEnvValue(pathRegKey, v)
}

func RemovePathFormEnv(removePath string) error {
	sep := string(os.PathListSeparator)
	v, err := GetRegEnvValue(pathRegKey)
	if err != nil {
		return err
	}

	paths := strings.Split(v, sep)
	for i, p := range paths {
		if p == removePath {
			continue
		}

		if i == 0 {
			v = p
		} else {
			v += sep + p
		}
	}

	// TODO: do logging here
	return SetRegEnvValue(pathRegKey, v)
}
