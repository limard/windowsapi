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
