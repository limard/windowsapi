package win

import (
	"syscall"
	"time"
	"unsafe"

	"bitbucket.org/Limard/logx"
)

const (
	PROCESS_TERMINATE                 = 0x0001
	PROCESS_CREATE_THREAD             = 0x0002
	PROCESS_SET_SESSIONID             = 0x0004
	PROCESS_VM_OPERATION              = 0x0008
	PROCESS_VM_READ                   = 0x0010
	PROCESS_VM_WRITE                  = 0x0020
	PROCESS_DUP_HANDLE                = 0x0040
	PROCESS_CREATE_PROCESS            = 0x0080
	PROCESS_SET_QUOTA                 = 0x0100
	PROCESS_SET_INFORMATION           = 0x0200
	PROCESS_QUERY_INFORMATION         = 0x0400
	PROCESS_SUSPEND_RESUME            = 0x0800
	PROCESS_QUERY_LIMITED_INFORMATION = 0x1000
	PROCESS_SET_LIMITED_INFORMATION   = 0x2000

	STANDARD_RIGHTS_REQUIRED = 0x000F0000
	SYNCHRONIZE              = 0x00100000

	PROCESS_ALL_ACCESS = uint32(STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0xFFFF)
)

const (
	// do not reorder
	TokenUser = 1 + iota
	TokenGroups
	TokenPrivileges
	TokenOwner
	TokenPrimaryGroup
	TokenDefaultDacl
	TokenSource
	TokenType
	TokenImpersonationLevel
	TokenStatistics
	TokenRestrictedSids
	TokenSessionId
	TokenGroupsAndPrivileges
	TokenSessionReference
	TokenSandBoxInert
	TokenAuditPolicy
	TokenOrigin
	TokenElevationType
	TokenLinkedToken
	TokenElevation
	TokenHasRestrictions
	TokenAccessInformation
	TokenVirtualizationAllowed
	TokenVirtualizationEnabled
	TokenIntegrityLevel
	TokenUIAccess
	TokenMandatoryPolicy
	TokenLogonSid
	MaxTokenInfoClass
)

var (
	dllKernel32                      = syscall.NewLazyDLL("kernel32.dll")
	procWTSGetActiveConsoleSessionId = dllKernel32.NewProc("WTSGetActiveConsoleSessionId")

	// dllWtsapi32           = syscall.NewLazyDLL("Wtsapi32.dll")
	// procWTSQueryUserToken = dllWtsapi32.NewProc("WTSQueryUserToken")

	dllAdvapi32             = syscall.NewLazyDLL("advapi32.dll")
	procDuplicateTokenEx    = dllAdvapi32.NewProc("DuplicateTokenEx")
	procSetTokenInformation = dllAdvapi32.NewProc("SetTokenInformation")
)

const (
	TOKEN_ADJUST_SESSIONID = 0x0100
)

func LaunchInActiveSesstion(cmd string) (pid int, handle uintptr, err error) {
	logx.Println(cmd)

	// DWORD sesstionId = WTSGetActiveConsoleSessionId();
	var sesstionId uint32
	for {
		sesstionId, err = WTSGetActiveConsoleSessionId()
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
	logx.Println("SesstionId", sesstionId)

	var hToken syscall.Token
	var hNewToken syscall.Token
	var lpEnvironment *uint16

	// si
	var si = new(syscall.StartupInfo)
	si.Cb = uint32(unsafe.Sizeof(*si))
	si.Flags = syscall.STARTF_USESTDHANDLES
	si.Desktop = syscall.StringToUTF16Ptr(`winsta0\default`)

	// pi
	pi := new(syscall.ProcessInformation)

	// // hProc = OpenProcess(PROCESS_ALL_ACCESS, FALSE, GetCurrentProcessId());
	// curPid := syscall.Getpid()
	// logx.Println("pid:", curPid)
	// hProc, err := syscall.OpenProcess(PROCESS_ALL_ACCESS, false, uint32(curPid))
	// if err != nil {
	// 	logx.Println("OpenProcess:", err)
	// 	goto Exit
	// }

	// // if (OpenProcessToken(hProc, TOKEN_ALL_ACCESS, &hToken) == FALSE)
	// if err := syscall.OpenProcessToken(hProc, syscall.TOKEN_ALL_ACCESS, &hToken); err != nil {
	// 	logx.Println("OpenProcessToken", err)
	// 	goto Exit
	// }
	// defer hToken.Close()

	hToken, err = wtsQueryUserToken(sesstionId)
	if err != nil {
		logx.Println("WTSQueryUserToken:", err)
		goto Exit
	}
	defer hToken.Close()

	err = enableAllPrivileges(hToken)
	if err != nil {
		logx.Println("EnableAllPrivileges:", err)
		// goto Exit
	}

	// if (DuplicateTokenEx(hToken, TOKEN_ALL_ACCESS, NULL, SecurityImpersonation, TokenPrimary, &hNewToken) == FALSE)
	hNewToken, err = DuplicateTokenEx(hToken, syscall.TOKEN_ALL_ACCESS)
	if err != nil {
		logx.Println("DuplicateTokenEx:", err)
		goto Exit
	}

	// if (SetTokenInformation(hNewToken, TokenSessionId, &sesstionId, sizeof(sesstionId)) == FALSE)
	err = SetTokenInformation(hNewToken, TokenSessionId, &sesstionId, 4)
	if err != nil {
		logx.Println("SetTokenInformation:", err)
		goto Exit
	}

	// if (CreateEnvironmentBlock(&lpEnvironment, hToken, FALSE) == FALSE)
	lpEnvironment, err = CreateEnvironmentBlock(hToken, false)
	if err != nil {
		logx.Println("CreateEnvironmentBlock:", err)
		goto Exit
	}

	// if (CreateProcessAsUser(
	// 	hNewToken,
	// 	NULL,
	// 	(LPWSTR)cmd.c_str(),
	// 	NULL,
	// 	NULL,
	// 	FALSE,
	// 	CREATE_NO_WINDOW | CREATE_UNICODE_ENVIRONMENT,
	// 	lpEnvironment,
	// 	NULL,
	// 	&si,
	// 	&pi) == FALSE)
	err = CreateProcessAsUser(hNewToken, cmd, nil, nil, false,
		syscall.CREATE_UNICODE_ENVIRONMENT, lpEnvironment, nil, si, pi)
	if err != nil {
		logx.Println("CreateProcessAsUser:", err)
		goto Exit
	}

	defer syscall.CloseHandle(syscall.Handle(pi.Thread))

	// ::WaitForSingleObject(pthis->m_hProcess, INFINITE);

Exit:
	// CloseHandle(hProc)

	// CloseHandle(hToken);

	// CloseHandle(hNewToken);

	if err != nil {
		logx.Println(err)
		return 0, 0, err
	}

	return 0, 0, nil
}

const (
	SecurityAnonymous      = 0
	SecurityIdentification = 1
	SecurityImpersonation  = 2
	SecurityDelegation     = 3

	TokenPrimary = 1
)

func DuplicateTokenEx(hExistingToken syscall.Token, dwDesiredAccess uint32) (syscall.Token, error) {
	var phNewToken syscall.Token

	// SecurityImpersonation
	// TokenPrimary
	r, _, e := procDuplicateTokenEx.Call(
		uintptr(hExistingToken),
		uintptr(dwDesiredAccess),
		uintptr(0),
		SecurityImpersonation,
		TokenPrimary,
		uintptr(unsafe.Pointer(&phNewToken)))
	if r == 0 {
		return phNewToken, e
	}
	return phNewToken, nil
}

func SetTokenInformation(TokenHandle syscall.Token, TokenSessionId int, TokenInformation *uint32, TokenInformationLength uint32) error {
	r, _, e := procSetTokenInformation.Call(
		uintptr(TokenHandle),
		uintptr(TokenSessionId),
		uintptr(unsafe.Pointer(TokenInformation)),
		uintptr(TokenInformationLength))
	if r == 0 {
		return e
	}
	return nil
}

// func CreateEnvironmentBlock(token syscall.Token, inherit bool) (*uint16, error) {
// 	proc, err := win.LoadProc("Userenv.dll", "CreateEnvironmentBlock")
// 	if err != nil {
// 		return nil, err
// 	}

// 	iInherit := 0
// 	if inherit {
// 		iInherit = 1
// 	}

// 	var env *uint16

// 	r1, _, err := proc.Call(
// 		uintptr(unsafe.Pointer(&env)),
// 		uintptr(token),
// 		uintptr(iInherit),
// 	)

// 	if r1 == 1 {
// 		return env, nil
// 	}
// 	return nil, err
// }

func enableAllPrivileges(token syscall.Token) error {
	privileges := []string{
		"SeCreateTokenPrivilege",
		"SeAssignPrimaryTokenPrivilege",
		"SeLockMemoryPrivilege",
		"SeIncreaseQuotaPrivilege",
		"SeMachineAccountPrivilege",
		"SeTcbPrivilege",
		"SeSecurityPrivilege",
		"SeTakeOwnershipPrivilege",
		"SeLoadDriverPrivilege",
		"SeSystemProfilePrivilege",
		"SeSystemtimePrivilege",
		"SeProfileSingleProcessPrivilege",
		"SeIncreaseBasePriorityPrivilege",
		"SeCreatePagefilePrivilege",
		"SeCreatePermanentPrivilege",
		"SeBackupPrivilege",
		"SeRestorePrivilege",
		"SeShutdownPrivilege",
		"SeDebugPrivilege",
		"SeAuditPrivilege",
		"SeSystemEnvironmentPrivilege",
		"SeChangeNotifyPrivilege",
		"SeRemoteShutdownPrivilege",
		"SeUndockPrivilege",
		"SeSyncAgentPrivilege",
		"SeEnableDelegationPrivilege",
		"SeManageVolumePrivilege",
		"SeImpersonatePrivilege",
		"SeCreateGlobalPrivilege",
		"SeTrustedCredManAccessPrivilege",
		"SeRelabelPrivilege",
		"SeIncreaseWorkingSetPrivilege",
		"SeTimeZonePrivilege",
		"SeCreateSymbolicLinkPrivilege",
	}

	for _, privilege := range privileges {
		err := EnablePrivilege(token, privilege)
		if err != nil {
			return err
		}
	}
	return nil
}
