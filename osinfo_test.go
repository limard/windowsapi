package win

import "testing"

func TestGetOSVersion(t *testing.T) {
	t.Log(GetOSVersion())
	t.Log(Is64bitOS())
}

func TestGetOSVersion2(t *testing.T) {
	//t.Log(equalOSVersion(5,1))

	t.Log(equalOSVersion(6,0))
	t.Log(equalOSVersion(6,1))
	t.Log(equalOSVersion(6,2))

	//t.Log(GetOSVersion2(7,0, 0,0))
	//t.Log(GetOSVersion2(7,1, 0,0))
	//t.Log(GetOSVersion2(7,2, 0,0))

	//t.Log(GetOSVersion2(8,0, 0,0))
	//t.Log(GetOSVersion2(10,0, 0,0))
}

func TestIsOSWorkstation(t *testing.T) {
	t.Log("isOSWorkstation")
	t.Log(isOSWorkstation())
}