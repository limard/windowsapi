package windowsapi

import "testing"

func TestParseCommand(t *testing.T) {
	parse := func(str string) {
		s0 := ParseCommand(str)
		for key, value := range s0 {
			t.Log(key, value)
		}
	}

	//parse(`psexec \\machine -u MYDOMAIN\myuser -p mypassword copy "c:\path to my dir"`)
	//parse(`psexec \\machine -u MYDOMAIN\myuser -p mypassword copy "c:\path to my dir" hallo`)
	//parse(`psexec \\machine -u MYDOMAIN\myuser -p mypassword copy 'c:\path to my dir'`)
	//parse(`psexec \\machine -u MYDOMAIN\myuser -p mypassword copy 'c:\path to my dir' hallo`)
	//parse(`  psexec \\machine -u MYDOMAIN\myuser -p mypassword copy "c:\path to my dir"  `)
	//parse(`psexec   \\machine   -u MYDOMAIN\myuser    -p mypassword   copy    "c:\path to my dir"`)
	//parse(`psexec \\machine 测试  字符串 "c:\program files"`)
	parse(`set Title 打印输出端`)
	//parse(`set Title 打印输出端 字符串`)
	//parse(`set Title 打印输出端 "字符串"`)
}