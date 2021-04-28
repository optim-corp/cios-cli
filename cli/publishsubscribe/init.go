package publishsubscribe

import (
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
	"github.com/optim-corp/cios-cli/utils/go_advance_type/convert"
	ftil "github.com/optim-corp/cios-cli/utils/go_advance_type/file"
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
