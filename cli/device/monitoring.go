package device

import (
	"context"

	"github.com/optim-corp/cios-cli/utils/console"

	. "github.com/optim-corp/cios-cli/cli"
	"github.com/optim-corp/cios-cli/models"
	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"
	"github.com/urfave/cli/v2"
)

func GetDeviceMonitoringCommand() *cli.Command {
	return &cli.Command{
		Name:    "monitoring",
		Aliases: []string{"monitor", "m"},
		Usage:   "cios monitoring | monitor | m | d m",
		Subcommands: []*cli.Command{
			listDeviceMonitoringCommand(),
		},
	}
}

func listDeviceMonitoringCommand() *cli.Command {
	return &cli.Command{
		Name:      models.LIST,
		Aliases:   models.ALIAS_LIST,
		UsageText: "cios device monitoring ls [command options] [device_id...]",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "all", Aliases: []string{"a"}},
		},
		Action: func(c *cli.Context) error {
			deviceIDs := []string{}
			if c.Args().Len() == 0 {
				devices, _, err := Client.DeviceManagement.GetDevicesAll(ciossdk.MakeGetDevicesOpts(), context.Background())
				assert(err).Log().NoneErr(func() {
					for _, device := range devices {
						deviceIDs = append(deviceIDs, device.Id)
					}
				})
			} else {
				console.CliArgsForEach(c, func(a string) { deviceIDs = append(deviceIDs, a) })
			}
			monitorings, _, err := Client.DeviceManagement.GetMonitoringLatestList(deviceIDs, context.Background())
			assert(err).
				Log().
				NoneErrPrintln(func() { console.OutStructJson(monitorings) })
			return nil
		},
	}
}
