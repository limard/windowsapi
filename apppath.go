// +build windows

package windowsapi

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

//ParseCommand is spilt command to args
//https://github.com/golang-devops/parsecommand
func ParseCommand(line string) ([]string) {
	args := [][]rune{}

	var quoteChar rune
	var isInQuote bool = false

	trimmedLine := strings.TrimSpace(line)

	currentArg := []rune{}
	trimmedLineLen := len(trimmedLine)
	for i, c := range trimmedLine {
		isLastChar := i + len(string(c)) == trimmedLineLen

		if !isInQuote && (c == '"' || c == '\'') {
			isInQuote = true
			quoteChar = c
			if isLastChar {
				args = append(args, currentArg)
			}
			continue
		}

		if isInQuote && c == quoteChar {
			//Ensure it is not escaped with a slash beforehand
			if i == 0 || trimmedLine[i-1] != '\\' {
				isInQuote = false
				if isLastChar {
					args = append(args, currentArg)
				}
				continue
			}
		}

		if !isInQuote && c == ' ' {
			//Ignore multiple spaces
			if i > 0 && trimmedLine[i-1] != ' ' {
				args = append(args, currentArg)
				currentArg = []rune{}
				continue
			}
		}

		currentArg = append(currentArg, c)

		if isLastChar {
			args = append(args, currentArg)
		}
	}

	strArgs := []string{}
	for _, a := range args {
		strArgs = append(strArgs, strings.TrimSpace(string(a)))
	}
	return strArgs
}