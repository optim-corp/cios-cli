package device

import (
	"context"

	. "github.com/optim-corp/cios-cli/cli"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
	"github.com/optim-corp/cios-golang-sdk/cios"
	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"
	"github.com/urfave/cli/v2"
	"gopkg.in/AlecAivazis/survey.v1"
)

func GetDeviceCommand() *cli.Command {
	return &cli.Command{
		Name:    "device",
		Aliases: []string{"d", "dv"},
		Usage:   "cios device | dv | d",
		Subcommands: []*cli.Command{
			createDevice(),
			deleteDevice(),
			listDevice(),
			GetDevicePolicyCommand(),
			GetDeviceMonitoringCommand(),
			GetDeviceInventoryCommand(),
			GetDeviceModelsCommand(),
			GetDeviceEntityCommand(),
			GetDeviceLifecycleCommand(),
		},
	}
}
func createDevice() *cli.Command {
	return &cli.Command{
		Name:    models.CREATE,
		Aliases: models.ALIAS_CREATE,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "custom", Aliases: []string{"c"}},
			&cli.StringFlag{Name: "resource_owner_id", Aliases: []string{"r"}, Required: true},
		},
		Action: func(c *cli.Context) error {
			in := struct {
				ResourceOwnerID string
				IDNumber        string
				Description     string
				Name            string
				ClientList      string
				IsManaged       string
				RsaPublickey    string
			}{}

			utils.Question([]*survey.Question{
				{
					Name:   "idNumber",
					Prompt: &survey.Input{Message: "ID Number: "},
				},
				{
					Name:   "description",
					Prompt: &survey.Input{Message: "Description: "},
				},
				{
					Name:   "name",
					Prompt: &survey.Input{Message: "Name: "},
				},
				{
					Name: "clientList",
					Prompt: &survey.Multiline{
						Message: "Client ID: ",
					},
				},
				{
					Name:   "isManaged",
					Prompt: &survey.Input{Message: "Is managed: "},
				},
				{
					Name:   "rsaPublickey",
					Prompt: &survey.Input{Message: "RSA Public Key: "},
				},
			}, &in)

			assert(
				func() error {
					_, _, err := Client.DeviceManagement.CreateDevice(cios.DeviceInfo{
						ResourceOwnerId: c.String("resource_owner_id"),
						Name:            &in.Name,
						IdNumber:        &in.IDNumber,
						IsManaged:       &in.IsManaged,
						RsaPublickey:    &in.RsaPublickey,
						Description:     &in.Description,
					}, context.Background())
					return err
				}()).Log().NoneErrPrintln("Completed " + in.Name)
			return nil
		},
	}
}
func deleteDevice() *cli.Command {
	return &cli.Command{
		Name:      models.DELETE,
		Aliases:   models.ALIAS_DELETE,
		UsageText: "cios device delete [command options] [device_id...]",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "all", Aliases: []string{"a"}},
		},
		Action: func(c *cli.Context) error {
			utils.CliArgsForEach(c, func(id string) {
				_, err := Client.DeviceManagement.DeleteDevice(id, context.Background())
				assert(err).Log().NoneErrPrintln("Completed ", id)
			})

			return nil
		},
	}
}
func listDevice() *cli.Command {
	return &cli.Command{
		Name:      models.LIST,
		Aliases:   models.ALIAS_LIST,
		UsageText: "cios device ls [command options] [device_id...]",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
			&cli.BoolFlag{Name: "all", Aliases: []string{"a"}},
		},
		Action: func(c *cli.Context) error {
			if c.Args().Len() > 0 {
				utils.CliArgsForEach(c, func(id string) {
					device, _, err := Client.DeviceManagement.GetDevice(id, nil, nil, context.Background())
					assert(err).Log().NoneErr(
						func() {
							value, err := utils.StructToJsonStr(device)
							assert(err).Log().ExitWith(1)
							value, err = utils.IndentJson(value)
							assert(err).Log().ExitWith(1)
							println(value)
						})

				})
			} else {
				isAll := c.Bool("all")
				devices, _, err := Client.DeviceManagement.GetDevicesAll(ciossdk.MakeGetDevicesOpts().Name(c.String("name")), context.Background())
				assert(err).
					Log().
					NoneErr(func() {
						listUtility(func() {
							if !isAll {
								fPrintln("\t|id|\t\t\t|resource_owner_id|\t\t |name <id number>|")
							}
							for _, device := range devices {
								fPrintln(
									is(isAll).
										T(func() string {
											str, err := utils.StructToJsonStr(device)
											assert(err).Log().ExitWith(1)
											str, err = utils.IndentJson(str)
											assert(err).Log().ExitWith(1)
											return str
										}()).
										F(func() string {
											name := ""
											idN := ""
											if device.Name != nil {
												name = *device.Name
											}
											if device.IdNumber != nil {
												idN = *device.IdNumber
											}
											return device.Id + "\t" + device.ResourceOwnerId + "\t" + name + " <" + idN + ">"
										}()).
										Value.(string),
								)
							}
						})
					})
			}
			return nil
		},
	}
}
