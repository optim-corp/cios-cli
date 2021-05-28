package device

import (
	"github.com/fcfcqloow/go-advance/ftil"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
	"github.com/optim-corp/cios-cli/utils/console"
)

var (
	listUtility  = console.ListUtility
	fPrintln     = console.Fprintln
	is           = utils.Is
	printf       = console.Printf
	assert       = utils.EAssert
	lifecycleDir = models.LifecycleDir
	path         = ftil.Path
)
