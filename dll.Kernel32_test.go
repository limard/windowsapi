package windowsapi

import (
	"testing"
	"os"
)

func TestProcessIdToSessionId(t *testing.T) {
	t.Log(ProcessIdToSessionId(uint32(os.Getpid())))
}
