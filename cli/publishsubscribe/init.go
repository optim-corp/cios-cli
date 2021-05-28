package publishsubscribe

import (
	cnv "github.com/fcfcqloow/go-advance/convert"
	"github.com/fcfcqloow/go-advance/ftil"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
	"github.com/optim-corp/cios-cli/utils/console"
)

var (
	is           = utils.Is
	listUtility  = console.ListUtility
	fPrintln     = console.Fprintln
	fPrintf      = console.Fprintf
	fPrint       = console.Fprint
	datastoreDir = models.DatastoreDir
	path         = ftil.Path
	assert       = utils.EAssert
	str          = cnv.MustStr
)
