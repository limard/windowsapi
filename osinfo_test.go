package win

import "testing"

func TestGetOSVersion(t *testing.T) {
	t.Log(GetOSVersion())
	t.Log(Is64bitOS())
}
