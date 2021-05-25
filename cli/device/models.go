package device

import (
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"

	"github.com/optim-corp/cios-cli/utils/console"

	"github.com/AlecAivazis/survey/v2"
	cnv "github.com/fcfcqloow/go-advance/convert"
	"github.com/fcfcqloow/go-advance/log"
	. "github.com/optim-corp/cios-cli/cli"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-golang-sdk/cios"
	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"
	"github.com/urfave/cli/v2"
)

func GetDeviceModelsCommand() *cli.Command {
	return &cli.Command{
		Name:    "models",
		Aliases: []string{"m", "model"},
		Usage:   "cios device models | model | m",
		Subcommands: []*cli.Command{
			listDeviceModels(),
			createDeviceModel(),
			deleteDeviceModel(),
			entityDeviceModel(),
		},
	}
}

func listDeviceModels() *cli.Command {
	return &cli.Command{
		Name:    models.LIST,
		Aliases: models.ALIAS_LIST,
		Usage:   "cios device model list | model ls",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "all", Aliases: []string{"a"}},
		},
		Action: func(c *cli.Context) error {
			printModel := func(model cios.DeviceModel) {
				fPrintln()
				fPrintln("|Device Model ID|  : ", model.Id)
				fPrintln("|Resource Owner ID|: ", model.ResourceOwnerId)
				fPrintln("|Component ID|     : ", model.Components.Id)
				fPrintln("|Name|             : ", model.Name)
				fPrintln("|Version|          : ", model.Version)
				fPrintln("|Watch|            : ", model.Watch)
				fPrintln("|Created at|       : ", model.CreatedAt)
				fPrintln("|Updated at|       : ", model.UpdatedAt, "\n")
				console.FOutStructJsonSlim(model.Components)
				fPrintln("\n------------------------------------------------------------------------------------------------------------------------------------------------------------\n")
			}
			if c.Args().Len() > 0 {
				console.CliArgsForEach(c, func(arg string) {
					modelMap, err := Client.DeviceAssetManagement.GetModelsMapByID(ciosctx.Background(), ciossdk.MakeGetModelsOpts())
					assert(err).Log().NoneErr(func() {
						model, ok := modelMap[arg]
						if !ok {
							model, _, err = Client.DeviceAssetManagement.GetModel(ciosctx.Background(), arg)
							if assert(err).Log().ErrNotNil() {
								return
							}
						}
						printModel(model)
					})
				})
				return nil
			}

			ms, _, err := Client.DeviceAssetManagement.GetModelsAll(ciosctx.Background(), ciossdk.MakeGetModelsOpts())
			isAll := c.Bool("all")
			assert(err).Log().NoneErr(func() {
				listUtility(func() {
					if isAll {
						for _, model := range ms {
							printModel(model)
						}
					} else {
						fPrintln("\t|ID|\t \t\t|Resource Owner ID|\t\t       |Component ID / Name  / Maker ID|")
						for _, model := range ms {
							fPrintln(model.Id, "\t", model.ResourceOwnerId, "\t  ", model.Components.Id, " / ", model.Name, " / ", model.MakerId)

						}
					}

				})
			})

			return nil
		},
	}
}

func deleteDeviceModel() *cli.Command {
	return &cli.Command{
		Name:    models.DELETE,
		Aliases: models.ALIAS_DELETE,
		Usage:   "cios device model list | model ls",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "all", Aliases: []string{"a"}},
		},
		Action: func(c *cli.Context) error {
			console.CliArgsForEach(c, func(arg string) {
				_, err := Client.DeviceAssetManagement.DeleteModel(ciosctx.Background(), arg)
				assert(err).Log().NoneErr(func() {
					modelMap, err := Client.DeviceAssetManagement.GetModelsMapByID(ciosctx.Background(), ciossdk.MakeGetModelsOpts())
					assert(err).Log().NoneErr(func() {
						if model, ok := modelMap[arg]; ok {
							_, err = Client.DeviceAssetManagement.DeleteModel(ciosctx.Background(), model.Name)
							assert(err).Log().NoneErrPrintln("Completed ", arg)
						}
					})
				})
			})
			return nil
		},
	}
}

func createDeviceModel() *cli.Command {
	return &cli.Command{
		Name:    models.CREATE,
		Aliases: models.ALIAS_CREATE,
		Usage:   "cios device model create | model add",
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			req := cios.DeviceModelRequest{}
			input := console.GetConsoleMultipleLine(">>")
			assert(cnv.UnMarshalJson(input, &req)).Log().NoneErr(func() {
				model, _, err := Client.DeviceAssetManagement.CreateModel(ciosctx.Background(), req)
				assert(err).Log().NoneErr(func() { console.OutStructJsonSlim(model) })
			})
			return nil
		},
	}
}

func entityDeviceModel() *cli.Command {
	return &cli.Command{
		Name:    "entity",
		Aliases: []string{"e", "entt", "ett"},
		Usage:   "cios device model create | model add",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
			&cli.StringFlag{Name: "id", Aliases: []string{"i"}},
		},
		Action: func(c *cli.Context) error {
			name := c.String("name")
			id := c.String("id")

			if name == "" {
				if id == "" {
					log.Error("No Name and ID")
					return nil
				}
				modelMap, err := Client.DeviceAssetManagement.GetModelsMapByID(ciosctx.Background(), ciossdk.MakeGetModelsOpts())
				if assert(err).Log().ErrNotNil() {
					if model, ok := modelMap[id]; ok {
						name = model.Name
					} else {
						log.Error("No Model ID")
						return nil
					}
				}

			}
			ans := struct {
				SerialNumber    string
				ResourceOwnerID string
				StartAt         string
				Value           string
			}{}

			console.Question([]*survey.Question{
				{
					Name:   "serialNumber",
					Prompt: &survey.Input{Message: "Serial Number: "},
				},
				{
					Name:   "resourceOwnerId",
					Prompt: &survey.Input{Message: "Resource Owner Id: "},
				},
				{
					Name:   "startAt",
					Prompt: &survey.Input{Message: "Start At: "},
				},
				{
					Name:   "value",
					Prompt: &survey.Multiline{Message: "Custom Inventory: "},
				},
			}, &ans)
			body := cios.Inventory{
				ResourceOwnerId: &ans.ResourceOwnerID,
				SerialNumber:    &ans.SerialNumber,
				StartAt:         &ans.StartAt,
			}
			if ans.Value != "" {
				assert(cnv.UnMarshalJson(ans.Value, &body.CustomInventory)).
					Log().NoneErr(func() {
					_, _, err := Client.DeviceAssetManagement.CreateEntity(ciosctx.Background(), name, body)
					assert(err).Log()
				})
			}
			return nil
		},
	}
}
