package utils

import (
	"bufio"
	"fmt"
	"os"
	"runtime"

	"github.com/mitchellh/go-homedir"
)

var (
	Out                     = bufio.NewWriter(os.Stdout)
	Dir, _                  = homedir.Dir()
	TopDir                  = Dir + "/.cios-cli"
	DatastoreDir            = TopDir + "/datastore"
	UrlPath                 = Is(os.Getenv("CIOS_CLI_URL_PATH") == "").T(TopDir + "/URL.json").F(os.Getenv("CIOS_CLI_URL_PATH")).Value.(string)
	ConfigPath              = Is(os.Getenv("CIOS_CLI_CONFIG_PATH") == "").T(TopDir + "/config.json").F(os.Getenv("CIOS_CLI_CONFIG_PATH")).Value.(string)
	AccountPath             = TopDir + "/accounts.json"
	TimestampFormatFilePath = TopDir + "/.timestamp_format"
	LifecycleDir            = TopDir + "/lifecycle"
	IsWindows               = (runtime.GOOS == "windows")
	IsLinux                 = (runtime.GOOS == "linux")
	IsMac                   = (runtime.GOOS == "darwin")
)

type (
	Logging struct {
		LogLevel int
	}

	Judge struct {
		Value interface{}
		flag  bool
	}

	FileService struct {
		Path string
	}

	Assert struct {
		Err error
	}
	DirByt struct {
		Value   []byte
		AbsPath string
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
