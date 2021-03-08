package tool

import (
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
	ftil "github.com/optim-kazuhiro-seida/go-advance-type/file"
)

var (
	config              models.Config
	accountFile         ftil.FileService
	configFile          ftil.FileService
	timestampFormatFile ftil.FileService
	configPath          = utils.ConfigPath
	accountPath         = utils.AccountPath
	timestampFormatPath = utils.TimestampFormatFilePath
	listUtility         = utils.ListUtility
	fPrintln            = utils.Fprintln
	fPrint              = utils.Fprint
	println             = utils.Println
	printf              = utils.Printf
	print               = utils.Print
	path                = ftil.Path
	assert              = utils.EAssert
	is                  = utils.Is
)
