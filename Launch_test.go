package windowsapi

import "testing"

func TestLaunchInActiveSesstion(t *testing.T) {
	_, _, err := LaunchInActiveSesstion(`E:\111.txt`)
	if err != nil {
		t.Error(err)
	}
}
