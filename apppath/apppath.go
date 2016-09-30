package apppath

import (
	"os"
	"path/filepath"
)

func GetAppPath() string {
	p, err := filepath.Abs(os.Args[0])
	if err != nil {
		p = os.Args[0]
	}
	dir := filepath.Dir(p)
	return dir
}
