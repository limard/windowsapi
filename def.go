package win

const (
	DEBUG_PROCESS                    = 0x00000001
	DEBUG_ONLY_THIS_PROCESS          = 0x00000002
	CREATE_SUSPENDED                 = 0x00000004
	DETACHED_PROCESS                 = 0x00000008
	CREATE_NEW_CONSOLE               = 0x00000010
	NORMAL_PRIORITY_CLASS            = 0x00000020
	IDLE_PRIORITY_CLASS              = 0x00000040
	HIGH_PRIORITY_CLASS              = 0x00000080
	REALTIME_PRIORITY_CLASS          = 0x00000100
	CREATE_NEW_PROCESS_GROUP         = 0x00000200
	CREATE_UNICODE_ENVIRONMENT       = 0x00000400
	CREATE_SEPARATE_WOW_VDM          = 0x00000800
	CREATE_SHARED_WOW_VDM            = 0x00001000
	BELOW_NORMAL_PRIORITY_CLASS      = 0x00004000
	ABOVE_NORMAL_PRIORITY_CLASS      = 0x00008000
	INHERIT_PARENT_AFFINITY          = 0x00010000
	CREATE_PROTECTED_PROCESS         = 0x00040000
	EXTENDED_STARTUPINFO_PRESENT     = 0x00080000
	PROCESS_MODE_BACKGROUND_BEGIN    = 0x00100000
	PROCESS_MODE_BACKGROUND_END      = 0x00200000
	CREATE_BREAKAWAY_FROM_JOB        = 0x01000000
	CREATE_PRESERVE_CODE_AUTHZ_LEVEL = 0x02000000
	CREATE_DEFAULT_ERROR_MODE        = 0x04000000
	CREATE_NO_WINDOW                 = 0x08000000
)

type startupinfo struct {
	/* DWORD */ cb uint32
	/* LPSTR */ lpReserved uintptr
	/* LPSTR */ lpDesktop uintptr
	/* LPSTR */ lpTitle uintptr
	/* DWORD */ dwX uint32
	/* DWORD */ dwY uint32
	/* DWORD */ dwXSize uint32
	/* DWORD */ dwYSize uint32
	/* DWORD */ dwXCountChars uint32
	/* DWORD */ dwYCountChars uint32
	/* DWORD */ dwFillAttribute uint32
	/* DWORD */ dwFlags uint32
	/* WORD */ wShowWindow uint16
	/* WORD */ cbReserved2 uint16
	/* LPBYTE */ lpReserved2 uintptr
	/* HANDLE */ hStdInput uintptr
	/* HANDLE */ hStdOutput uintptr
	/* HANDLE */ hStdError uintptr
}

type processinfo struct {
	/* HANDLE */ hProcess uintptr
	/* HANDLE */ hThread uintptr
	/* DWORD */ dwProcessId uint32
	/* DWORD */ dwThreadId uint32
}
