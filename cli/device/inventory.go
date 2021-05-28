package device

import (
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"

	"github.com/optim-corp/cios-cli/utils/console"

	. "github.com/optim-corp/cios-cli/cli"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-golang-sdk/cios"
	"github.com/urfave/cli/v2"
)

func GetDeviceInventoryCommand() *cli.Command {
	return &cli.Command{
		Name:    "inventory",
		Aliases: []string{"i", "invt"},
		Usage:   "cios inventory | invt | i |  device inventory | d i",
		Subcommands: []*cli.Command{
			listDeviceInventoryCommand(),
		},
	}
}

func listDeviceInventoryCommand() *cli.Command {
	return &cli.Command{
		Name:      models.LIST,
		Aliases:   models.ALIAS_LIST,
		UsageText: "cios device inventory ls [command options] [device_id...]",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "all", Aliases: []string{"a"}},
		},
		Action: func(c *cli.Context) error {
			list := func(deviceID string) {
				inventory, _, err := Client.DeviceManagement.GetDeviceInventory(ciosctx.Background(), deviceID)
				assert(err).
					Log().NoneErr(func() {
					fPrintln("\n\n\nDevice ID: " + deviceID + "\n\n\n")
					console.FOutStructJsonSlim(inventory)
					fPrintln("------------------------------------------------------------------------------------------------------------------------")
				})
			}
			listUtility(func() {
				if c.Args().Len() == 0 {
					devices, _, err := Client.DeviceManagement.GetDevicesAll(ciosctx.Background(), cios.ApiGetDevicesRequest{})
					assert(err).Log().NoneErr(func() {
						for _, device := range devices {
							list(device.Id)
						}
					})
				} else {
					console.CliArgsForEach(c, func(id string) { list(id) })
				}
			})
			return nil
		},
	}
}
