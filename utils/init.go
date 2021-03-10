package utils

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

var (
	Out = bufio.NewWriter(os.Stdout)

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

func Fprintln(v ...interface{}) {
	fmt.Fprintln(Out, v...)
}
func Fprint(v ...interface{}) {
	fmt.Fprint(Out, v...)
}
func Fprintf(format string, v ...interface{}) {
	fmt.Fprintf(Out, format, v...)
}
func Println(v ...interface{}) {
	fmt.Println(v...)
}
func Print(v ...interface{}) {
	fmt.Print(v...)
}
func Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}
