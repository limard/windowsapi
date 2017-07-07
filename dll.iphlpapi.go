package windowsapi

// copy from
// https://github.com/azyshalom/netstat/blob/1ab548d3df44c1a0c59b80ce5fe3340c7d4728f2/netstat_windows.go

import (
	"encoding/binary"
	"fmt"
	"syscall"
	"unsafe"
)

var (
	dIphlpapi            = syscall.NewLazyDLL("iphlpapi.dll")
	pGetExtendedTcpTable = dIphlpapi.NewProc("GetExtendedTcpTable")
	pGetExtendedUdpTable = dIphlpapi.NewProc("GetExtendedUdpTable")
)

const (
	AF_INET = 2
)

const (
	TCPTableBasicListener = iota
	TCPTableBasicConnections
	TCPTableBasicAll
	TCPTableOwnerPIDListener
	TCPTableOwnerPIDConnections
	TCPTableOwnerPIDAll
	TCPTableOwnerModuleListener
	TCPTableOwnerModuleConnections
	TCPTableOwnerModuleAll
)

const (
	UDPTableBasicListener = iota
	UDPTableOwnerPIDAll
	UDPTableOwnerModuleAll
)

type State int

const (
	MIB_TCP_STATE_CLOSED State = 1 + iota
	MIB_TCP_STATE_LISTEN
	MIB_TCP_STATE_SYN_SENT
	MIB_TCP_STATE_SYN_RCVD
	MIB_TCP_STATE_ESTAB
	MIB_TCP_STATE_FIN_WAIT1
	MIB_TCP_STATE_FIN_WAIT2
	MIB_TCP_STATE_CLOSE_WAIT
	MIB_TCP_STATE_CLOSING
	MIB_TCP_STATE_LAST_ACK
	MIB_TCP_STATE_TIME_WAIT
	MIB_TCP_STATE_DELETE_TCB
)

type MIB_TCPROW_OWNER_PID struct {
	dwState      uint32
	dwLocalAddr  uint32
	dwLocalPort  uint32
	dwRemoteAddr uint32
	dwRemotePort uint32
	dwOwningPid  uint32
}

type MIB_TCPTABLE_OWNER_PID struct {
	dwNumEntries uint32
	table        [200]MIB_TCPROW_OWNER_PID
}

type MIB_UDPROW_OWNER_PID struct {
	dwLocalAddr uint32
	dwLocalPort uint32
	dwOwningPid uint32
}

type MIB_UDPTABLE_OWNER_PID struct {
	dwNumEntries uint32
	table        [200]MIB_UDPROW_OWNER_PID
}

type PortEnumRow struct {
	Proto      string
	State      State
	LocalAddr  string
	LocalPort  uint32
	RemoteAddr string
	RemotePort uint32
	ProcessId  uint32
}

func GetExtendedTcpTableEx() (<-chan PortEnumRow, error) {
	b := make([]byte, 100)
	size := uint32(len(b))
	ret, _, _ := pGetExtendedTcpTable.Call(
		uintptr(unsafe.Pointer(&b[0])),
		uintptr(unsafe.Pointer(&size)),
		0,
		uintptr(AF_INET),
		TCPTableOwnerPIDAll,
		0)

	if ret == uintptr(syscall.ERROR_INSUFFICIENT_BUFFER) {
		b = make([]byte, size)
		ret, _, _ = pGetExtendedTcpTable.Call(
			uintptr(unsafe.Pointer(&b[0])),
			uintptr(unsafe.Pointer(&size)),
			0,
			uintptr(AF_INET),
			TCPTableOwnerPIDAll,
			0)
	}
	if ret != 0 {
		return nil, syscall.GetLastError()
	}
	ch := make(chan PortEnumRow)
	go func() {
		table := (*MIB_TCPTABLE_OWNER_PID)(unsafe.Pointer(&b[0]))
		for i := 0; i < int(table.dwNumEntries) && i < 200; i++ {
			row := PortEnumRow{
				Proto:      "TCP",
				State:      State(table.table[i].dwState),
				LocalAddr:  getIpAddress(table.table[i].dwLocalAddr),
				LocalPort:  getPortNumber(table.table[i].dwLocalPort),
				RemoteAddr: getIpAddress(table.table[i].dwRemoteAddr),
				RemotePort: getPortNumber(table.table[i].dwRemotePort),
				ProcessId:  table.table[i].dwOwningPid,
			}
			ch <- row
		}
		close(ch)
	}()
	return ch, nil
}

func GetExtendedUdpTableEx() (<-chan PortEnumRow, error) {
	b := make([]byte, 100)
	size := uint32(len(b))
	ret, _, _ := pGetExtendedUdpTable.Call(
		uintptr(unsafe.Pointer(&b[0])),
		uintptr(unsafe.Pointer(&size)),
		0,
		uintptr(AF_INET),
		UDPTableOwnerPIDAll,
		0)

	if ret == uintptr(syscall.ERROR_INSUFFICIENT_BUFFER) {
		b = make([]byte, size)
		ret, _, _ = pGetExtendedUdpTable.Call(
			uintptr(unsafe.Pointer(&b[0])),
			uintptr(unsafe.Pointer(&size)),
			0,
			uintptr(AF_INET),
			UDPTableOwnerPIDAll,
			0)
	}
	if ret != 0 {
		return nil, syscall.GetLastError()
	}
	ch := make(chan PortEnumRow)
	go func() {
		table := (*MIB_UDPTABLE_OWNER_PID)(unsafe.Pointer(&b[0]))
		for i := 0; i < int(table.dwNumEntries) && i < 200; i++ {
			row := PortEnumRow{
				Proto:     "UDP",
				LocalAddr: getIpAddress(table.table[i].dwLocalAddr),
				LocalPort: getPortNumber(table.table[i].dwLocalPort),
				ProcessId: table.table[i].dwOwningPid,
			}
			ch <- row
		}
		close(ch)
	}()
	return ch, nil
}

func getPortNumber(port uint32) uint32 {
	return port/256 + (port%256)*256
}

func getIpAddress(ip uint32) string {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, ip)
	return fmt.Sprintf("%d.%d.%d.%d", b[0], b[1], b[2], b[3])
}

func getState(state State) string {
	m := map[State]string{
		MIB_TCP_STATE_CLOSED:     "CLOSED",
		MIB_TCP_STATE_LISTEN:     "LISTEN",
		MIB_TCP_STATE_SYN_SENT:   "SYN_SEND",
		MIB_TCP_STATE_SYN_RCVD:   "SYN_RECV",
		MIB_TCP_STATE_ESTAB:      "ESTABLISHED",
		MIB_TCP_STATE_FIN_WAIT1:  "FIN_WAIT_1",
		MIB_TCP_STATE_FIN_WAIT2:  "FIN_WAIT_2",
		MIB_TCP_STATE_CLOSE_WAIT: "CLOSE_WAIT",
		MIB_TCP_STATE_CLOSING:    "CLOSING",
		MIB_TCP_STATE_LAST_ACK:   "LAST_ACK",
		MIB_TCP_STATE_TIME_WAIT:  "TIME_WAIT",
		MIB_TCP_STATE_DELETE_TCB: "DELETE_TBC",
	}
	return m[state]
}
