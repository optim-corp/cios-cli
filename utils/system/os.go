package system

import "runtime"

var (
	IsWindows = runtime.GOOS == "windows"
	IsLinux   = runtime.GOOS == "linux"
	IsMac     = runtime.GOOS == "darwin"
)
