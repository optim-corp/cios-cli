package device

import (
	"context"

	"github.com/optim-corp/cios-golang-sdk/cios"

	. "github.com/optim-corp/cios-cli/cli"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
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
				inventory, _, err := Client.DeviceManagement.GetDeviceInventory(deviceID, context.Background())
				assert(err).
					Log().NoneErr(func() {
					fPrintln("\n\n\nDevice ID: " + deviceID + "\n\n\n")
					utils.FOutStructJsonSlim(inventory)
					fPrintln("------------------------------------------------------------------------------------------------------------------------")
				})
			}
			listUtility(func() {
				if c.Args().Len() == 0 {
					devices, _, err := Client.DeviceManagement.GetDevicesAll(cios.ApiGetDevicesRequest{}, context.Background())
					assert(err).Log().NoneErr(func() {
						for _, device := range devices {
							list(device.Id)
						}
					})
				} else {
					utils.CliArgsForEach(c, func(id string) { list(id) })
				}
			})
			return nil
		},
	}
}
