package windowsapi

import (
	"os"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows/registry"
)

var (
	pathRegKey string = "Path"
)

func BroadcastSettingChange() {
	rawParam, _ := syscall.UTF16PtrFromString("ENVIRONMENT")
	param := uintptr(unsafe.Pointer(rawParam))
	pSendMessageTimeout.Call(uintptr(HWND_BROADCAST), uintptr(WM_SETTINGCHANGE), 0, param,
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
