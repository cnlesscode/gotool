package gotool

import (
	"os"
	"path/filepath"
	"runtime"
)

var Root string
var SystemSeparator string
var OS string

func init() {
	SystemSeparator = string(filepath.Separator)
	exePath, err := os.Executable()
	if err != nil {
		return
	}
	Root, _ = filepath.Abs(filepath.Dir(exePath))
	Root += SystemSeparator
	OS = runtime.GOOS
}
