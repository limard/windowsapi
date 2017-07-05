package windowsapi

import "syscall"

var (
	dUser32         = syscall.NewLazyDLL("user32.dll")

	pSendMessageTimeout = dUser32.NewProc("SendMessageTimeoutW")
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