package filestorage

import (
	"github.com/fcfcqloow/go-advance/ftil"
	"github.com/optim-corp/cios-cli/utils"
	"github.com/optim-corp/cios-cli/utils/console"
)

var (
	is          = utils.Is
	listUtility = console.ListUtility
	spaceRight  = console.SpaceRight
	fPrintln    = console.Fprintln
	fPrintf     = console.Fprintf
	path        = ftil.Path
	assert      = utils.EAssert
)
