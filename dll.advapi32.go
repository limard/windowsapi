// +build windows

package windowsapi

import (
	"syscall"
	"unsafe"
)

var (
	dAdvapi32 = syscall.NewLazyDLL("advapi32.dll")

	pDuplicateTokenEx    = dAdvapi32.NewProc("DuplicateTokenEx")
	pSetTokenInformation = dAdvapi32.NewProc("SetTokenInformation")
)

const (
	SecurityAnonymous      = 0
	SecurityIdentification = 1
	SecurityImpersonation  = 2
	SecurityDelegation     = 3

	TokenPrimary = 1
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

func DuplicateTokenEx(hExistingToken syscall.Token, dwDesiredAccess uint32) (syscall.Token, error) {
	var phNewToken syscall.Token

	// SecurityImpersonation
	// TokenPrimary
	r, _, e := pDuplicateTokenEx.Call(
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
	r, _, e := pSetTokenInformation.Call(
		uintptr(TokenHandle),
		uintptr(TokenSessionId),
		uintptr(unsafe.Pointer(TokenInformation)),
		uintptr(TokenInformationLength))
	if r == 0 {
		return e
	}
	return nil
}

func CreateProcessAsUser(
	token syscall.Token,
// applicationName string,
	cmd string,
	procSecurity *syscall.SecurityAttributes,
	threadSecurity *syscall.SecurityAttributes,
	inheritHandles bool,
	creationFlags uint32,
	environment *uint16,
	currentDir *uint16,
	startupInfo *syscall.StartupInfo,
	outProcInfo *syscall.ProcessInformation,
) error {
	proc, err := loadProc("advapi32.dll", "CreateProcessAsUserW")
	if err != nil {
		return err
	}

	iInheritHandles := 0
	if inheritHandles {
		iInheritHandles = 1
	}

	r1, _, err := proc.Call(
		uintptr(token),
		// uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(applicationName))),
		uintptr(0),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(cmd))),
		uintptr(unsafe.Pointer(procSecurity)),
		uintptr(unsafe.Pointer(threadSecurity)),
		uintptr(iInheritHandles),
		uintptr(creationFlags),
		uintptr(unsafe.Pointer(environment)),
		uintptr(unsafe.Pointer(currentDir)),
		uintptr(unsafe.Pointer(startupInfo)),
		uintptr(unsafe.Pointer(outProcInfo)),
	)

	if r1 == 1 {
		return nil
	}
	return err
}

type luid struct {
	lowPart  uint32
	highPart uint32
}

func LookupPrivilegeValue(systemName string, name string) (*luid, error) {
	proc, err := loadProc("advapi32.dll", "LookupPrivilegeValueW")
	if err != nil {
		return nil, err
	}

	l := luid{}

	wsSystemName := uintptr(0)
	if len(systemName) > 0 {
		wsSystemName = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(systemName)))
	}

	r1, _, err := proc.Call(
		wsSystemName,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name))),
		uintptr(unsafe.Pointer(&l)),
	)
	if r1 == 1 {
		return &l, nil
	}
	return nil, err
}

type luidAndAttributes struct {
	luid       luid
	attributes uint32
}

type tokenPrivileges struct {
	privilegeCount uint32
	privileges     *luidAndAttributes
}

func AdjustTokenPrivileges(token syscall.Token, uid luid) error {
	proc, err := loadProc("advapi32.dll", "AdjustTokenPrivileges")
	if err != nil {
		return err
	}

	var sePrivilegeEnabled = uint32(0x00000002)

	newState := tokenPrivileges{
		privilegeCount: 1,
		privileges: &luidAndAttributes{
			luid:       uid,
			attributes: sePrivilegeEnabled,
		},
	}

	r1, _, err := proc.Call(
		uintptr(token),
		uintptr(0),
		uintptr(unsafe.Pointer(&newState)),
		uintptr(unsafe.Sizeof(newState)),
		uintptr(0),
		uintptr(0),
	)
	if r1 == 1 {
		return nil
	}
	return err
}