package utils

import (
	"bufio"
	"os"
	"runtime"
)

var (
	IsWindows = (runtime.GOOS == "windows")
	IsLinux   = (runtime.GOOS == "linux")
	IsMac     = (runtime.GOOS == "darwin")
	Console   = &out{writer: bufio.NewWriter(os.Stdout)}
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
