package device

import (
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
	ftil "github.com/optim-corp/cios-cli/utils/go_advance_type/file"
)

var (
	listUtility  = utils.ListUtility
	fPrintln     = utils.Fprintln
	is           = utils.Is
	println      = utils.Println
	print        = utils.Print
	printf       = utils.Printf
	assert       = utils.EAssert
	lifecycleDir = models.LifecycleDir
	path         = ftil.Path
)
