package device

import (
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"

	"github.com/optim-corp/cios-cli/utils/console"

	. "github.com/optim-corp/cios-cli/cli"
	"github.com/optim-corp/cios-cli/models"
	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"
	"github.com/urfave/cli/v2"
)

func GetDeviceLifecycleCommand() *cli.Command {
	return &cli.Command{
		Name:    "lifecycle",
		Aliases: []string{"lc", "life", "cycle"},
		Usage:   "cios device lifecycle | lc | cycle",
		Subcommands: []*cli.Command{
			deleteDeviceLifecycle(),
			listDeviceLifecycle(),
		},
	}
}

func deleteDeviceLifecycle() *cli.Command {
	return &cli.Command{
		Name:    models.DELETE,
		Aliases: models.ALIAS_DELETE,
		Usage:   "cios device lifecycle delete | cycle del",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "start_timestamp", Aliases: []string{"st"}},
			&cli.StringFlag{Name: "end_timestamp", Aliases: []string{"et"}},
			&cli.StringFlag{Name: "key", Aliases: []string{"k"}, Required: true},
		},
		Action: func(c *cli.Context) error {
			var (
				key            = c.String("key")
				startTimestamp = c.String("start_timestamp")
				endTimestamp   = c.String("end_timestamp")
				// wg             = sync.WaitGroup{}
			)
			lifecycles, _, err := Client.DeviceAssetManagement.GetLifecyclesAll(ciosctx.Background(), key,
				ciossdk.MakeGetLifecyclesOpts().
					StartEventAt(startTimestamp).
					EndEventAt(endTimestamp))
			assert(err).Log().NoneErr(func() {
				for _, lifecycle := range lifecycles {
					//time.Sleep(time.Millisecond * 50)
					_, err := Client.DeviceAssetManagement.DeleteLifecycle(ciosctx.Background(), key, lifecycle.Id)
					assert(err).Log().NoneErrPrintln("Completed ", lifecycle.Id)
				}
			})
			return nil
		},
	}
}

func listDeviceLifecycle() *cli.Command {
	return &cli.Command{
		Name:    models.LIST,
		Aliases: models.ALIAS_LIST,
		Usage:   "cios device lifecycle list | lc ls",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "key", Aliases: []string{"k"}, Required: true},
			&cli.StringFlag{Name: "order_by", Aliases: []string{"ob"}},
			&cli.StringFlag{Name: "component_id", Aliases: []string{"cid"}},
			&cli.StringFlag{Name: "start_timestamp", Aliases: []string{"st"}},
			&cli.StringFlag{Name: "end_timestamp", Aliases: []string{"et"}},
			&cli.BoolFlag{Name: "save", Aliases: []string{"s"}},
		},
		Action: func(c *cli.Context) error {
			var (
				key            = c.String("key")
				orderBy        = c.String("order_by")
				componentId    = c.String("component_id")
				startTimestamp = c.String("start_timestamp")
				endTimestamp   = c.String("end_timestamp")
				save           = c.Bool("save")
			)
			if c.Args().Len() == 0 {
				lifecycles, _, err := Client.DeviceAssetManagement.GetLifecyclesAll(key, ciossdk.MakeGetLifecyclesOpts().
					OrderBy(orderBy).
					ComponentId(componentId).
					StartEventAt(startTimestamp).
					EndEventAt(endTimestamp))
				stageDSDir := lifecycleDir + "/" + models.GetStage()
				if save {
					path(lifecycleDir).CreateDir()
					path(stageDSDir).CreateDir()
					_path := stageDSDir + "/" + key + ".json"
					assert(path(_path).WriteJson(lifecycles)).Log()
				}
				assert(err).Log().NoneErr(func() {
					listUtility(func() {
						for _, lifecycle := range lifecycles {
							console.FOutStructJson(lifecycle)
						}
					})

				})
			}
			return nil
		},
	}
}
