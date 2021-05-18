package utils

import (
	"runtime"
)

var (
	IsWindows = (runtime.GOOS == "windows")
	IsLinux   = (runtime.GOOS == "linux")
	IsMac     = (runtime.GOOS == "darwin")
)

type (
	Judge struct {
		Value interface{}
		flag  bool
	}
	Assert struct {
		Err error
	}
)
