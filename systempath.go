package windowsapi

// C:\Windows\System32\spool\prtprocs\x64
func GetPrintProcessorDirectory64() (path string, err error) {
	return GetPrintProcessorDirectory("Windows x64")
}

// C:\Windows\System32\spool\prtprocs\x86
func GetPrintProcessorDirectory86() (path string, err error) {
	return GetPrintProcessorDirectory("Windows NT x86")
}

func GetPrinterDriverDirectory64() (path string, err error) {
	return GetPrinterDriverDirectory("Windows x64")
}

func GetPrinterDriverDirectory86() (path string, err error) {
	return GetPrinterDriverDirectory("Windows NT x86")
}

// CSIDL_APPDATA: C:\Documents and Settings\username\Application Data
func GetAppDataDirectory() (string, error) {
	return SHGetSpecialFolderPath(CSIDL_APPDATA)
}

// CSIDL_LOCAL_APPDATA: C:\Documents and Settings\username\Local Settings\Application Data.
func GetLocalAppDataDirectory() (string, error) {
	return SHGetSpecialFolderPath(CSIDL_LOCAL_APPDATA)
}

// CSIDL_COMMON_APPDATA: "C:\Documents and Settings\All Users\Application Data" or "C:\ProgramData"
func GetCommmonAppDataDirectory() (string, error) {
	return SHGetSpecialFolderPath(CSIDL_COMMON_APPDATA)
}

// // C:\Documents and Settings\...\Desktop
func GetDesktopDir() (string, error) {
	return SHGetFolderPath(CSIDL_DESKTOP)
}

// C:\Documents and Settings\All Users\Desktop
func GetCommonDesktopDir() (string, error) {
	return SHGetFolderPath(CSIDL_COMMON_DESKTOPDIRECTORY)
}

// C:\Windows
func GetWindowsDir() (string, error) {
	return SHGetFolderPath(CSIDL_WINDOWS)
}

// C:\Windows\System32
func GetSystemDir() (string, error) {
	return SHGetFolderPath(CSIDL_SYSTEM)
}

// C:\Windows\SysWOW64
func GetSystem86Dir() (string, error) {
	return SHGetFolderPath(CSIDL_SYSTEMX86)
}

// C:\Program Files
func GetProgramFilesDir() (string, error) {
	return SHGetFolderPath(CSIDL_PROGRAM_FILES)
}

// C:\Program Files (x86)
func GetProgramFiles86Dir() (string, error) {
	return SHGetFolderPath(CSIDL_PROGRAM_FILESX86)
}

//  C:\Documents and Settings\username\Templates
func GetUserTempSystemDir() (string, error) {
	return SHGetFolderPath(CSIDL_TEMPLATES)
}

// C:\Program Files\Common
func GetProgramFilesCommonDir() (string, error) {
	return SHGetFolderPath(CSIDL_PROGRAM_FILES_COMMON)
}

// C:\Program Files (x86)\Common
func GetProgramFilesCommon86Dir() (string, error) {
	return SHGetFolderPath(CSIDL_PROGRAM_FILES_COMMONX86)
}
