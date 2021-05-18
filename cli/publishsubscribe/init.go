package publishsubscribe

import (
	cnv "github.com/fcfcqloow/go-advance/convert"
	"github.com/fcfcqloow/go-advance/ftil"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
)

var (
	out          = utils.Out
	is           = utils.Is
	listUtility  = utils.ListUtility
	fPrintln     = utils.Console.Fprintln
	fPrintf      = utils.Console.Fprintf
	fPrint       = utils.Console.Fprint
	datastoreDir = models.DatastoreDir
	path         = ftil.Path
	assert       = utils.EAssert
	str          = cnv.MustStr
)
