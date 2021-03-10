package device

import (
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
	"github.com/optim-kazuhiro-seida/ftil"
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
