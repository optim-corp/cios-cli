package publishsubscribe

import (
	"context"
	"strings"

	. "github.com/optim-corp/cios-cli/cli"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
	"github.com/optim-corp/cios-golang-sdk/cios"
	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"
	"github.com/urfave/cli/v2"
	"gopkg.in/AlecAivazis/survey.v1"
)

func GetChannelCommand() *cli.Command {
	return &cli.Command{
		Name:    "channel",
		Aliases: []string{"c", "ch", "channels"},
		Usage:   "cios channel | ch",
		Subcommands: []*cli.Command{
			createChannel(),
			deleteChannel(),
			listChannel(),
			updateChannel(),
		},
	}
}

func createChannel() *cli.Command {
	return &cli.Command{
		Name:    models.CREATE,
		Aliases: models.ALIAS_CREATE,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "name", Aliases: []string{"nm", "n"}},
			&cli.StringFlag{Name: "label", Aliases: []string{"l", "lab"}, Usage: "key1=value1,key=value2"},
			&cli.StringFlag{Name: "description", Aliases: []string{"d"}},
			&cli.StringFlag{Name: "language", Aliases: []string{"lan"}, DefaultText: "ja", Value: "ja"},
			&cli.BoolFlag{Name: "messaging_enabled", Aliases: []string{"me"}, DefaultText: "true", Value: true},
			&cli.BoolFlag{Name: "messaging_persisted", Aliases: []string{"mp"}, DefaultText: "true", Value: true},
			&cli.BoolFlag{Name: "datastore_enabled", Aliases: []string{"de"}, DefaultText: "true", Value: true},
			&cli.StringFlag{Name: "datastore_max_count", Aliases: []string{"dmc"}, DefaultText: "0", Value: "0"},
			&cli.StringFlag{Name: "datastore_max_size", Aliases: []string{"dms"}, DefaultText: "0", Value: "0"},
			&cli.StringFlag{Name: "resource_owner_id", Aliases: []string{"r", "ro"}},
		},
		Action: func(c *cli.Context) error {
			name := c.String("name")
			resourceOwnerID := is(c.String("resource_owner_id") != "").
				T(c.String("resource_owner_id")).
				F("").Value.(string)
			body := cios.ChannelProposal{}
			if name != "" {
				var (
					labelArg    = c.String("label")
					labels      = strings.Split(labelArg, ",")
					labelReq    []cios.Label
					description = c.String("description")
					language    = c.String("language")
					enabled     = c.Bool("messaging_enabled")
					persisted   = c.Bool("messaging_persisted")
					dEnabled    = c.Bool("datastore_enabled")
					maxCount    = c.String("datastore_max_count")
					maxSize     = c.String("datastore_max_size")
				)

				if labelArg != "" {
					for _, l := range labels {
						kv := strings.Split(l, "=")
						labelReq = append(labelReq, cios.Label{
							Key:   kv[0],
							Value: kv[1],
						})
					}
				}

				body = cios.ChannelProposal{
					ResourceOwnerId: resourceOwnerID,
					DisplayInfo: []cios.DisplayInfo{
						{
							Name:        name,
							Description: &description,
							Language:    language,
							IsDefault:   true,
						},
					},
					Labels:          &labelReq,
					MessagingConfig: &cios.MessagingConfig{Enabled: &enabled, Persisted: &persisted},
					DatastoreConfig: &cios.DataStoreConfig{Enabled: &dEnabled, MaxCount: &maxCount, MaxSize: &maxSize},
				}
			} else {
				answers := struct {
					Name               string
					Description        string
					Language           string
					IsDefault          bool
					ResourceOwnerID    string
					Label              string
					MessagingEnabled   bool
					MessagingPersisted bool
					DataStoreEnabled   bool
				}{}
				dataStoreConfig := struct {
					MaxSize  string
					MaxCount string
				}{
					MaxCount: "0",
					MaxSize:  "0",
				}
				utils.Question(
					[]*survey.Question{
						{
							Name:   "name",
							Prompt: &survey.Input{Message: "Name: "},
						},
						{
							Name:   "description",
							Prompt: &survey.Input{Message: "Description: "},
						},
						{
							Name:   "language",
							Prompt: &survey.Input{Message: "Language: ", Default: "ja"},
						},
						// {
						// 	Name:   "isDefault",
						// 	Prompt: &survey.Confirm{Message: "is default", Default: true},
						// },
						{
							Name: "resourceOwnerID",
							Prompt: &survey.Input{
								Message: "Resource Owner ID: ",
								Default: resourceOwnerID,
							},
						},
						{
							Name:   "label",
							Prompt: &survey.Multiline{Message: "Labels(key=value): "},
						},
						{
							Name:   "messagingEnabled",
							Prompt: &survey.Confirm{Message: "Messaging Enabled", Default: true},
						},
						{
							Name:   "messagingPersisted",
							Prompt: &survey.Confirm{Message: "Messaging Persisted", Default: true},
						},
						{
							Name:   "dataStoreEnabled",
							Prompt: &survey.Confirm{Message: "Data Store Enabled", Default: true},
						},
					}, &answers,
				)
				if answers.DataStoreEnabled {
					utils.Question([]*survey.Question{
						{Name: "maxCount", Prompt: &survey.Input{Message: "Max Count: ", Default: "0"}},
						{Name: "maxSize", Prompt: &survey.Input{Message: "Max Size: ", Default: "0"}},
					}, &dataStoreConfig)
				}
				labelExp := func(exp bool) []cios.Label {
					if exp {
						labels := strings.Split(answers.Label, "\n")
						var result []cios.Label
						for _, l := range labels {
							kv := strings.Split(l, "=")
							result = append(result, cios.Label{Key: kv[0], Value: kv[1]})
						}
						return result
					}
					return []cios.Label{}
				}
				labels := labelExp(answers.Label != "")
				body = cios.ChannelProposal{
					ResourceOwnerId: answers.ResourceOwnerID,
					DisplayInfo: []cios.DisplayInfo{
						{
							Name:        answers.Name,
							Description: &answers.Description,
							Language:    answers.Language,
							IsDefault:   true,
						},
					},
					Labels: &labels,
					MessagingConfig: &cios.MessagingConfig{
						Enabled:   &answers.MessagingEnabled,
						Persisted: &answers.MessagingPersisted,
					},
					DatastoreConfig: &cios.DataStoreConfig{
						Enabled:  &answers.DataStoreEnabled,
						MaxCount: &dataStoreConfig.MaxCount,
						MaxSize:  &dataStoreConfig.MaxSize,
					},
				}

			}
			channel, _, err := Client.PubSub.CreateChannel(body, context.Background())
			assert(err).
				Log().
				NoneErrPrintln("Completed ", channel.Id)
			return nil
		},
	}
}
func deleteChannel() *cli.Command {
	return &cli.Command{
		Name:    models.DELETE,
		Aliases: models.ALIAS_DELETE,
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			utils.CliArgsForEach(c, func(id string) {
				_, err := Client.PubSub.DeleteChannel(id, nil)
				assert(err).
					Log().
					NoneErrPrintln("Completed ", id)
			})
			return nil
		},
	}
}
func listChannel() *cli.Command {
	return &cli.Command{
		Name:    models.LIST,
		Aliases: models.ALIAS_LIST,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
			&cli.BoolFlag{Name: "all", Aliases: []string{"a"}},
			&cli.BoolFlag{Name: "detail", Aliases: []string{"d"}},
			&cli.StringFlag{Name: "label", Aliases: []string{"lb"}},
			&cli.StringFlag{Name: "resource_owner_id", Aliases: []string{"r", "ro"}},
			&cli.Int64Flag{Name: "limit", Aliases: []string{"l"}, Value: 30},
			&cli.Int64Flag{Name: "offset", Aliases: []string{"o"}, Value: 0},
		},
		Action: func(c *cli.Context) error {
			var (
				name             = c.String("name")
				label            = c.String("label")
				isAll            = c.Bool("all")
				isDetail         = c.Bool("detail")
				resourceOwnerID  = c.String("resource_owner_id")
				limit            = c.Int64("limit")
				offset           = c.Int64("offset")
				resourceOwnerMap map[string]cios.ResourceOwner
			)
			channels, _, err := Client.PubSub.GetChannelsAll(ciossdk.MakeGetChannelsOpts().
				ResourceOwnerId(resourceOwnerID).
				Label(label).
				Name(name).
				Limit(limit).
				Offset(offset), nil)
			assert(err).Log()
			if isDetail {
				listUtility(func() {
					for _, channel := range channels {
						utils.FOutStructJson(channel)
						fPrintln()
					}
				})
				return nil
			}
			if isAll {
				resourceOwnerMap, _, err = Client.Account.GetResourceOwnersMapByID(context.Background())
			}
			listUtility(func() {
				fPrint("\t|ID|\t\t\t|Resource Owner ID|\t\t  |Name|\t\t|Labels|")
				if isAll {
					fPrint(" : |Resource Owner Name|")
				}
				fPrintln()
				for _, value := range channels {
					fPrint(value.Id, "\t", value.ResourceOwnerId, "\t", value.Name, "\t", value.Labels)
					if isAll {
						resourceOwner, _ := resourceOwnerMap[value.ResourceOwnerId]
						name := ""
						if resourceOwner.Profile.DisplayName != nil {
							name = *resourceOwner.Profile.DisplayName
						}
						fPrint(" : ", name)
					}
					fPrintln()
				}
			})
			return nil
		},
	}
}
func updateChannel() *cli.Command {
	return &cli.Command{
		Name:    models.PATCH,
		Aliases: models.ALIAS_PATCH,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "label", Aliases: []string{"l", "labels"}},
			&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
		},
		Action: func(c *cli.Context) error {
			labels := is(c.String("label") == "").T([]string{}).F(strings.Split(c.String("label"), "=")).Value.([]string)
			name := c.String("name")
			if len(labels) >= 2 && name != "" {
				utils.CliArgsForEach(c, func(channelID string) {
					_, _, err := Client.PubSub.UpdateChannel(channelID, cios.ChannelUpdateProposal{
						DisplayInfo: []cios.DisplayInfo{
							{
								Name:      c.String("name"),
								Language:  "ja",
								IsDefault: true,
							},
						},
						Labels: &[]cios.Label{
							{
								Key:   labels[0],
								Value: labels[1],
							},
						},
					}, context.Background())
					assert(err).Log().NoneErrPrintln("Completed " + channelID)
				})
			} else if name != "" {
				utils.CliArgsForEach(c, func(channelID string) {
					_, _, err := Client.PubSub.UpdateChannel(channelID, cios.ChannelUpdateProposal{
						DisplayInfo: []cios.DisplayInfo{
							{
								Name:      c.String("name"),
								Language:  "ja",
								IsDefault: true,
							},
						},
					}, context.Background())
					assert(err).Log().NoneErrPrintln("Completed " + channelID)
				})
			} else if len(labels) >= 2 {
				utils.CliArgsForEach(c, func(channelID string) {
					_, _, err := Client.PubSub.UpdateChannel(channelID, cios.ChannelUpdateProposal{
						DisplayInfo: nil,
						Labels: &[]cios.Label{
							{
								Key:   labels[0],
								Value: labels[1],
							},
						}}, context.Background())
					assert(err).Log().NoneErrPrintln("Completed ", channelID)
				})
			} else if c.String("label") == "" {
				answers := struct {
					Name        string
					Description string
					Language    string
					isDefault   bool
					Label       string
				}{}
				utils.CliArgsForEach(c, func(channelID string) {
					utils.Question([]*survey.Question{
						{
							Name:   "name",
							Prompt: &survey.Input{Message: "name: "},
						},
						{
							Name:   "description",
							Prompt: &survey.Input{Message: "description: "},
						},
						{
							Name:   "language",
							Prompt: &survey.Input{Message: "language: ", Default: "ja"},
						},
						{
							Name:   "isDefault",
							Prompt: &survey.Confirm{Message: "is default", Default: true},
						},
						{
							Name:   "label",
							Prompt: &survey.Input{Message: "label(key=value): "},
						},
					}, &answers)
					labelExp := func(exp bool) []cios.Label {
						if exp {
							return []cios.Label{
								{
									Key:   strings.Split(answers.Label, "=")[0],
									Value: strings.Split(answers.Label, "=")[1],
								},
							}
						}
						return []cios.Label{}
					}
					labels := labelExp(answers.Label != "")
					_, _, err := Client.PubSub.UpdateChannel(
						channelID,
						cios.ChannelUpdateProposal{
							DisplayInfo: []cios.DisplayInfo{
								{
									Name:        answers.Name,
									Description: &answers.Description,
									Language:    answers.Language,
									IsDefault:   answers.isDefault,
								},
							},
							Labels: &labels,
						}, context.Background())
					assert(err).Log().NoneErrPrintln()
				})
			}

			return nil
		},
	}
}
