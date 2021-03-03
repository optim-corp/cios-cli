package utils

import (
	"bufio"
	"os"

	"github.com/mitchellh/go-homedir"
)

var (
	Out                     = bufio.NewWriter(os.Stdout)
	scanner                 = bufio.NewScanner(os.Stdin)
	Dir, _                  = homedir.Dir()
	TopDir                  = Dir + "/.cios-cli"
	FlowmenDir              = TopDir + "/flowmen"
	FlowmenNodePath         = FlowmenDir + "/node.yml"
	DatastoreDir            = TopDir + "/datastore"
	UrlPath                 = Is(os.Getenv("CIOS_CLI_URL_PATH") == "").T(TopDir + "/URL.json").F(os.Getenv("CIOS_CLI_URL_PATH")).Value.(string)
	ConfigPath              = Is(os.Getenv("CIOS_CLI_CONFIG_PATH") == "").T(TopDir + "/config.json").F(os.Getenv("CIOS_CLI_CONFIG_PATH")).Value.(string)
	AccountPath             = TopDir + "/accounts.json"
	TimestampFormatFilePath = TopDir + "/.timestamp_format"
	LifecycleDir = TopDir + "/lifecycle"
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
