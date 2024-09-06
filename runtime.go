package xgo

import (
	"path/filepath"
	"runtime"
)

const (
	DarwinOS  = runtime.GOOS == "darwin"
	WindowsOS = runtime.GOOS == "windows"
	LinuxOS   = runtime.GOOS == "linux"
)

func RuntimeDir() string {
	_, file, _, _ := runtime.Caller(1)
	return filepath.Dir(file)
}
