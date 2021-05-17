package tool

import (
	"fmt"

	"github.com/fcfcqloow/go-advance/ftil"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
)

var (
	accountFile         ftil.FileService
	configFile          ftil.FileService
	timestampFormatFile ftil.FileService
	configPath          = models.ConfigPath
	accountPath         = models.AccountPath
	timestampFormatPath = models.TimestampFormatFilePath
	listUtility         = utils.ListUtility
	fPrintln            = utils.Console.Fprintln
	fPrint              = utils.Console.Fprint
	printf              = utils.Console.Printf
	path                = ftil.Path
	println             = fmt.Println
	assert              = utils.EAssert
	is                  = utils.Is
)
