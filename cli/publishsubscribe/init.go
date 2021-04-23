package publishsubscribe

import (
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
	"github.com/optim-kazuhiro-seida/ftil"
	"github.com/optim-kazuhiro-seida/go-advance-type/convert"
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
	str          = convert.MustStr
)
