package tool

import (
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
)

var (
	config              models.Config
	accountFile         utils.FileService
	configFile          utils.FileService
	timestampFormatFile utils.FileService
	configPath          = utils.ConfigPath
	urlDir              = utils.UrlPath
	accountPath         = utils.AccountPath
	timestampFormatPath = utils.TimestampFormatFilePath
	log                 = utils.Log
	listUtility         = utils.ListUtility
	fPrintln            = utils.Fprintln
	fPrintf             = utils.Fprintf
	fPrint              = utils.Fprint
	println             = utils.Println
	printf              = utils.Printf
	print               = utils.Print
	path                = utils.Path
	assert              = utils.EAssert
	is                  = utils.Is
)
