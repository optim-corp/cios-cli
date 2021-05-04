package tool

import (
	"encoding/json"

	"github.com/fcfcqloow/go-advance/ftil"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
	. "github.com/optim-corp/cios-cli/utils"
	"github.com/urfave/cli/v2"
	"gopkg.in/AlecAivazis/survey.v1"
)

func GetLogCommand() *cli.Command {
	return &cli.Command{
		Name: "log",
		Action: func(c *cli.Context) error {
			var (
				config   models.Config
				levelStr = ""
			)

			if c.Args().Len() == 0 {
				ans := struct{ Value string }{}
				Question([]*survey.Question{
					{
						Name: "value",
						Prompt: &survey.Select{
							Message: "Pick Level",
							Options: []string{
								"debug",
								"info",
								"warn",
								"error",
								"emergency",
							},
						},
					},
				}, &ans)
				levelStr = ans.Value
			} else {
				levelStr = c.Args().First()
			}
			configPath = models.ConfigPath
			file, err := path(configPath).ReadFile()
			assert(err).
				Log().
				NoneErrAssert(json.Unmarshal(file, &config)).
				NoneErr(func() { config.LogLevel = levelStr }).
				Log().
				NoneErrAssert(ftil.Path(configPath).WriteJson(config)).
				NoneErrPrintln("Config Level: ", config.LogLevel).
				Log()

			return nil
		},
	}
}
func GetSwitchCommand() *cli.Command {
	return &cli.Command{
		Name:    "switch",
		Aliases: []string{"sw", "swap"},
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			var (
				account = models.Account{}
				in      = struct {
					Name string
				}{}
				accounts    = models.Account{}
				accountFile = path(accountPath)
				configFile  = path(configPath)
			)

			assert(accountFile.LoadJsonStruct(&accounts)).Log().NoneErr(func() {
				assert(path(models.AccountPath).LoadJsonStruct(&account)).Log().NoneErr(func() {
					var keys []string
					for key, _ := range account {
						keys = append(keys, key)
					}
					utils.Question(
						[]*survey.Question{
							{
								Name:   "name",
								Prompt: &survey.Select{Message: "Select a Stage", Options: keys},
							},
						}, &in)
					stage := in.Name
					keys = []string{}
					for k := range accounts[stage] {
						keys = append(keys, k)
					}

					utils.Question(
						[]*survey.Question{
							{
								Name: "name",
								Prompt: &survey.Select{
									Message: "Choose a name:",
									Options: keys,
								},
							},
						}, &in,
					)
					name := in.Name
					assert(configFile.WriteJson(accounts[stage][name])).Log()
				})

			})

			return nil
		},
	}
}
func GetConfigCommand() *cli.Command {
	return &cli.Command{
		Name:    "config",
		Aliases: []string{"conf", "cf"},
		Usage:   "cios config | conf",
		Before: func(c *cli.Context) error {
			accountFile = path(accountPath)
			configFile = path(configPath)
			return nil
		},
		Subcommands: []*cli.Command{
			GetSwitchCommand(),
			GetLogCommand(),
			{
				Name:    models.LIST,
				Aliases: []string{"lis", "ls"},
				Flags:   []cli.Flag{},
				Action: func(c *cli.Context) error {
					if config, ok := models.GetConfig(); ok {
						ListUtility(func() {
							fPrintln("\tStage:          " + config.Stage)
							fPrintln("\tClient ID:      " + config.ClientID)
							fPrintln("\tClient Secret:  " + config.ClientSecret)
							fPrintln("\tLog Level:      " + config.LogLevel)
						})
					}
					return nil
				},
			},
			{
				Name:    "save",
				Aliases: []string{"store", "sv"},
				Action: func(c *cli.Context) error {
					name := ""
					if c.Args().Len() == 0 {
						in := struct {
							Name string
						}{}
						utils.Question(
							[]*survey.Question{
								{
									Name:   "name",
									Prompt: &survey.Input{Message: "Name: "},
								},
							}, &in)
						name = in.Name
					} else {
						name = c.Args().Get(0)
					}
					accounts := models.Account{}
					config := models.Config{}
					return assert(accountFile.LoadJsonStruct(&accounts)).Log().
						NoneErrAssert(configFile.LoadJsonStruct(&config)).Log().
						NoneErr(func() {
							configs, ok := accounts[models.GetStage()]
							if ok {
								configs[name] = config
								accounts[models.GetStage()] = configs
							} else {
								accounts[models.GetStage()] = map[string]models.Config{name: config}
							}
						}).
						NoneErrAssert(accountFile.WriteJson(accounts)).Log().
						NoneErrPrintln("Success to save config").Err
				},
			},
			{
				Name:    "clean",
				Aliases: []string{"rm", "del", "remove", "delete", "cl"},
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "store", Aliases: []string{"s", "history", "past", "p"}},
				},
				Action: func(c *cli.Context) error {
					if c.Bool("store") {
						assert(accountFile.WriteFileAsString("")).Log()
					} else {
						assert(configFile.WriteFileAsString("")).Log()
					}
					return nil
				},
			},
		},
	}
}
