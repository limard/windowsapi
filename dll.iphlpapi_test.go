package windowsapi

import (
	"log"
	"testing"
)

func TestGetExtendedTcpTableEx(t *testing.T) {
	list, err := GetExtendedTcpTableEx()
	if err != nil {
		log.Print(err)
		return
	}

	for l := range list {
		log.Printf("pid: %d - %s:%d\n",
			l.ProcessId,
			l.LocalAddr,
			l.LocalPort)
	}
}
