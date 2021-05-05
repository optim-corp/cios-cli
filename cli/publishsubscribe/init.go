package publishsubscribe

import (
	cnv "github.com/fcfcqloow/go-advance/convert"
	"github.com/fcfcqloow/go-advance/ftil"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
)

var (
	print        = utils.Print
	printf       = utils.Printf
	out          = utils.Out
	is           = utils.Is
	listUtility  = utils.ListUtility
	fPrintln     = utils.Fprintln
	fPrintf      = utils.Fprintf
	fPrint       = utils.Fprint
	datastoreDir = models.DatastoreDir
	path         = ftil.Path
	assert       = utils.EAssert
	str          = cnv.MustStr
)
