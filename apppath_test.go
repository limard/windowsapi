package windowsapi

import "testing"

func TestParseCommand(t *testing.T) {
	compare := func(str string) {
		s0 := ParseCommand(str)
		for key, value := range s0 {
			t.Log(key, value)
		}
	}

	compare(`psexec \\machine -u MYDOMAIN\myuser -p mypassword copy "c:\path to my dir"`)
	compare(`psexec \\machine -u MYDOMAIN\myuser -p mypassword copy "c:\path to my dir" hallo`)
	compare(`psexec \\machine -u MYDOMAIN\myuser -p mypassword copy 'c:\path to my dir'`)
	compare(`psexec \\machine -u MYDOMAIN\myuser -p mypassword copy 'c:\path to my dir' hallo`)
	compare(`  psexec \\machine -u MYDOMAIN\myuser -p mypassword copy "c:\path to my dir"  `)
	compare(`psexec   \\machine   -u MYDOMAIN\myuser    -p mypassword   copy    "c:\path to my dir"`)
}