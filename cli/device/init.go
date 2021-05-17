package device

import (
	"github.com/fcfcqloow/go-advance/ftil"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
)

var (
	listUtility  = utils.ListUtility
	fPrintln     = utils.Console.Fprintln
	is           = utils.Is
	printf       = utils.Console.Printf
	assert       = utils.EAssert
	lifecycleDir = models.LifecycleDir
	path         = ftil.Path
)
