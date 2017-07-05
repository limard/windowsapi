// +build windows

package windowsapi

import (
	"os"
	"os/exec"
	"path/filepath"
)

func GetAppDir() string {
	lp, _ := exec.LookPath(os.Args[0])
	p, err := filepath.Abs(lp)
	if err != nil {
		p = os.Args[0]
	}
	dir := filepath.Dir(p)
	return dir
}

func GetAppExePath() string {
	lp, err := exec.LookPath(os.Args[0])
	if err != nil {
		lp = os.Args[0]
	}
	p, err := filepath.Abs(lp)
	if err != nil {
		p = lp
	}
	return p
}

// ParseCommand is spilt command to args
func ParseCommand(command string) (commandArgs []string) {
	var inQuote = false
	var tempStr []byte
	count := len(command)
	for i := 0; i < count; i++ {
		if command[i] == '"' {
			// log.Println("\"")
			inQuote = inQuote == false
			continue
		}

		if command[i] == ' ' && inQuote == false {
			// log.Println(string(tempStr))
			commandArgs = append(commandArgs, string(tempStr))
			tempStr = make([]byte, 0)
			continue
		}

		tempStr = append(tempStr, command[i])
	}

	commandArgs = append(commandArgs, string(tempStr))

	return
}
