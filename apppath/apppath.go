package apppath

import (
	"os"
	"os/exec"
	"path/filepath"
)

func GetAppPath() string {
	lp, _ := exec.LookPath(os.Args[0])
	p, err := filepath.Abs(lp)
	if err != nil {
		p = os.Args[0]
	}
	dir := filepath.Dir(p)
	return dir
}

// ParseCommand is spilt command to args
func ParseCommand(command string) (commandArgs []string) {
	var inQuote = false
	var tempStr []byte
	for i := 0; i < len(command); i++ {
		if command[i] == '"' {
			inQuote = inQuote == false
			continue
		}

		if command[i] == ' ' && inQuote == false {
			commandArgs = append(commandArgs, string(tempStr))
			tempStr = make([]byte, 0)
			continue
		}

		tempStr = append(tempStr, command[i])
	}
	return
}
