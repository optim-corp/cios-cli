package utils

import (
	"io"
	"os/exec"
	"runtime"
	"strings"
)

func Sh(cmds ...string) {
	Log.Info(strings.Join(cmds, " "))
	cmd := Is(len(cmds) >= 2).
		T(exec.Command(cmds[0], cmds[1:]...)).
		F(exec.Command(cmds[0])).Value.(*exec.Cmd)
	stdin, _ := cmd.StdinPipe()
	io.WriteString(stdin, "")
	stdin.Close()
	out, _ := cmd.Output()
	Println(string(out))
}

var (
	IsWindows = (runtime.GOOS == "windows")
	IsLinux   = (runtime.GOOS == "linux")
	IsMac     = (runtime.GOOS == "darwin")
)
