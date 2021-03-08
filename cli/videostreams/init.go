package video

import (
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
	"github.com/optim-kazuhiro-seida/ftil"
)

var (
	out         = utils.Out
	is          = utils.Is
	listUtility = utils.ListUtility
	spaceRight  = utils.SpaceRight
	fPrintln    = utils.Fprintln
	fPrintf     = utils.Fprintf
	fPrint      = utils.Fprint
	println     = utils.Println
	printf      = utils.Printf
	print       = utils.Print
	dir         = models.Dir
	path        = ftil.Path
	assert      = utils.EAssert
)
