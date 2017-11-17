package windowsapi

import (
	"syscall"
	"unsafe"
)

var (
	dUser32 = syscall.NewLazyDLL("user32.dll")

	pSendMessageTimeout = dUser32.NewProc("SendMessageTimeoutW")
	pmessageBox = dUser32.NewProc("MessageBoxW")
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

// MessageBox
const (
	MB_OK                = 0x00000000
	MB_OKCANCEL          = 0x00000001
	MB_ABORTRETRYIGNORE  = 0x00000002
	MB_YESNOCANCEL       = 0x00000003
	MB_YESNO             = 0x00000004
	MB_RETRYCANCEL       = 0x00000005
	MB_CANCELTRYCONTINUE = 0x00000006
	MB_ICONHAND          = 0x00000010
	MB_ICONQUESTION      = 0x00000020
	MB_ICONEXCLAMATION   = 0x00000030
	MB_ICONASTERISK      = 0x00000040
	MB_USERICON          = 0x00000080
	MB_ICONWARNING       = MB_ICONEXCLAMATION
	MB_ICONERROR         = MB_ICONHAND
	MB_ICONINFORMATION   = MB_ICONASTERISK
	MB_ICONSTOP          = MB_ICONHAND

	MB_DEFBUTTON1 = 0x00000000
	MB_DEFBUTTON2 = 0x00000100
	MB_DEFBUTTON3 = 0x00000200
	MB_DEFBUTTON4 = 0x00000300
)

const (
	IdAbort    = 3
	IdCancel   = 2
	IdContinue = 11
	IdIgnore   = 5
	IdNo       = 7
	IdOk       = 1
	IdRetry    = 4
	IdTryAgain = 10
	IdYes      = 6
)

func MessageBox(text string, caption string, typ int) (int, error) {
	textUtf16, _ := syscall.UTF16PtrFromString(text)
	captionUtf16, _ := syscall.UTF16PtrFromString(caption)

	ret, _, err := pmessageBox.Call(
		0,
		uintptr(unsafe.Pointer(textUtf16)),
		uintptr(unsafe.Pointer(captionUtf16)),
		uintptr(typ))
	if ret == 0 {
		return 0, err
	}
	return int(ret), nil
}