package utils

import (
	"bufio"
	"os"
	"runtime"
)

var (
	Out = bufio.NewWriter(os.Stdout)

	IsWindows = (runtime.GOOS == "windows")
	IsLinux   = (runtime.GOOS == "linux")
	IsMac     = (runtime.GOOS == "darwin")
	Console   = out{}
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
