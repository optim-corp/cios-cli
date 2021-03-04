package publishsubscribe

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/optim-kazuhiro-seida/go-advance-type/convert"

	"github.com/optim-corp/cios-golang-sdk/cios"

	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"

	. "github.com/optim-corp/cios-cli/cli"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
	"github.com/urfave/cli/v2"
)

func GetDataStoreCommand() *cli.Command {
	return &cli.Command{
		Name:    "datastore",
		Aliases: []string{"ds", "DS"},
		Usage:   "cios datastore | ds",
		Subcommands: []*cli.Command{
			createDataStore(),
			deleteDataStore(),
			listDataStore(),
			embezzleDateStore(),
		},
	}
}

func createDataStore() *cli.Command {
	return &cli.Command{
		Name:    models.CREATE,
		Aliases: models.ALIAS_CREATE,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "resource_owner_id", Aliases: []string{"r"}},
		},
		Action: func(c *cli.Context) error {
			utils.CliArgsForEach(c, func(channelID string) {})
			fmt.Println("未実装")
			return nil
		},
	}
}
func deleteDataStore() *cli.Command {
	return &cli.Command{
		Name:    models.DELETE,
		Aliases: models.ALIAS_DELETE,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "channel_id", Aliases: []string{"cID", "c"}, Required: true},
			&cli.StringFlag{Name: "resource_owner_id", Aliases: []string{"r", "ro"}},
			&cli.StringFlag{Name: "timestamp_range", Aliases: []string{"tr"}},
		},
		Action: func(c *cli.Context) error {
			var (
				channelID      = c.String("channel_id")
				timestampRange = c.String("timestamp_range")
			)
			if channelID != "" {
				if timestampRange != "" {
					objects, _, err := Client.PubSub.GetObjectsAll(channelID, ciossdk.MakeGetObjectsOpts().TimestampRange(timestampRange), context.Background())
					assert(err).
						Log().
						NoneErr(func() {
							for _, object := range objects {
								_, err := Client.PubSub.DeleteObject(channelID, object.Id, context.Background())
								assert(err).
									Log().
									NoneErrPrintln("Completed ", object.Id)
							}
						})

				} else if c.Args().Len() == 0 {
					_, err := Client.PubSub.DeleteDataByChannel(channelID, context.Background())
					assert(err).Log()
				} else {
					utils.CliArgsForEach(c, func(objectID string) {
						_, err := Client.PubSub.DeleteObject(channelID, objectID, nil)
						assert(err).Log().NoneErrPrintln("Completed ", channelID, objectID)
					})
				}
			}

			return nil
		},
	}
}
func listDataStore() *cli.Command {
	return &cli.Command{
		Name:    models.LIST,
		Aliases: models.ALIAS_LIST,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "channel_id", Aliases: []string{"c"}},
			&cli.StringFlag{Name: "packer_format", Aliases: []string{"pf"}},
			&cli.StringFlag{Name: "timestamp_range", Aliases: []string{"tr"}},
			&cli.StringFlag{Name: "save_dir", Aliases: []string{"out"}},
			&cli.StringFlag{Name: "label", Aliases: []string{"lb"}},
			&cli.BoolFlag{Name: "data", Aliases: []string{"d"}},
			&cli.BoolFlag{Name: "save", Aliases: []string{"s"}},
			&cli.Int64Flag{Name: "limit", Aliases: []string{"l"}, DefaultText: "30", Value: 30},
			&cli.Int64Flag{Name: "offset", Aliases: []string{"o"}, DefaultText: "0", Value: 0},
			&cli.StringFlag{Name: "resource_owner_id", Aliases: []string{"r", "ro"}},
			&cli.StringFlag{Name: "change_channel_id", Aliases: []string{"chg-c"}},
			&cli.StringFlag{Name: "data_replace", Aliases: []string{"replace", "rep"}, Usage: ""},
		},
		Action: func(c *cli.Context) error {
			var (
				channelID         = c.String("channel_id")
				packerFormat      = c.String("packer_format")
				outputDir         = c.String("save_dir")
				limit             = c.Int64("limit")
				offset            = c.Int64("offset")
				timestampRange    = is(c.String("timestamp_range") != "").T(c.String("timestamp_range")).F(":" + convert.MustStr(time.Now().UnixNano())).Value.(string)
				resourceOwnerID   = c.String("resource_owner_id")
				dataFlag          = c.Bool("data")
				saveFlag          = c.Bool("save")
				label             = c.String("label")
				replace           = c.String("data_replace")
				channelsMap, _, _ = Client.PubSub.GetChannelsMapByID(ciossdk.MakeGetChannelsOpts(), context.Background())
			)

			printObj := func(channel cios.Channel, limit int64) {
				if dataFlag {
					stageDSDir := datastoreDir + "/" + utils.GetStage()
					channelDir := datastoreDir + "/" + utils.GetStage() + "/" + channelsMap[channel.Id].Name + "___" + channel.Id
					data, err := Client.PubSub.GetStreamAll(channel.Id, ciossdk.MakeGetStreamOpts().Limit(limit).PackerFormat(packerFormat).TimestampRange(timestampRange).Label(label).Offset(offset), context.Background())
					assert(err).Log()
					fPrintf("\n|Channel ID|  : %s \n|Channel Name|: %s\n\n", channel.Id, channelsMap[channel.Id].Name)
					if saveFlag {
						path(datastoreDir).CreateDir()
						path(stageDSDir).CreateDir()
						path(channelDir).CreateDir()
					}

					for count, val := range data {
						if replace != "" {
							splitInComma := strings.Split(replace, ",")
							for _, _val := range splitInComma {
								splitVal := strings.Split(_val, ":")
								if len(splitVal) >= 2 {
									val = strings.Replace(val, splitVal[0], splitVal[1], -1)
								} else {
									log.Warn("Replace Missing format string")
								}
							}
						}
						if saveFlag {
							if outputDir != "" {
								channelDir = outputDir
							}
							filePrefix := strings.Repeat("0", len(strconv.Itoa(len(data)))-len(strconv.Itoa(count)))
							_path := channelDir + "/" + filePrefix + strconv.Itoa(count) + ".txt"
							assert(path(_path).WriteFileAsString(val)).Log()
							fPrintln(val)
						}
						fPrintln(val)
					}
				} else {
					objects, _, err := Client.PubSub.GetObjectsAll(channel.Id, ciossdk.MakeGetObjectsOpts().Limit(limit).TimestampRange(timestampRange).Label(label).Offset(offset), context.Background())
					if len(objects) == 0 || resourceOwnerID != "" && channelsMap[channel.Id].ResourceOwnerId != resourceOwnerID {
						return
					}
					fPrintf("\n|Channel ID|  : %s \n|Channel Name|: %s\n\n", channel.Id, channelsMap[channel.Id].Name)
					fPrintln("\t|ID|\t\t|Timestamp|\t  |Mime Type|")
					assert(err).
						Log().
						NoneErr(func() {
							for _, obj := range objects {
								fPrintf("%s %s %s\n", obj.Id, obj.Timestamp, obj.MimeType)
							}
						})
				}
				assert(out.Flush()).Log()
			}

			listUtility(func() {
				if channelID != "" {
					channel, _, err := Client.PubSub.GetChannel(channelID, nil, nil, context.Background())
					assert(err).
						Log().
						NoneErr(func() { printObj(channel, limit) })
				} else {
					channels, _, err := Client.PubSub.GetChannelsAll(ciossdk.MakeGetChannelsOpts().ResourceOwnerId(resourceOwnerID), context.Background())
					assert(err).
						Log().
						NoneErr(func() {
							for _, channel := range channels {
								printObj(channel, limit)
							}
						})
				}
			})
			return nil
		},
	}
}
func embezzleDateStore() *cli.Command {
	return &cli.Command{
		Name:    "embezzle",
		Aliases: []string{"emb", "emz", "agent", "broker"},
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "source_channel_id", Aliases: []string{"sc"}, Required: true},
			&cli.StringFlag{Name: "target_channel_id", Aliases: []string{"tc"}, Required: true},
			&cli.StringFlag{Name: "timestamp_range", Aliases: []string{"tr"}},
			&cli.StringFlag{Name: "label", Aliases: []string{"lb"}},
			&cli.BoolFlag{Name: "ascending", Aliases: []string{"oder_by", "ob", "asc"}},
		},
		Action: func(c *cli.Context) error {
			var (
				sourceChannel  = c.String("source_channel_id")
				targetChannel  = c.String("target_channel_id")
				timestampRange = c.String("timestamp_range")
				label          = c.String("label")
				ascending      = c.Bool("ascending")
				packerFormat   = "json"
				receiver       = make(chan *string)
				done           = make(chan bool)
			)
			if err := Client.PubSub.ConnectWebSocket(targetChannel, done, ciossdk.ConnectWebSocketOptions{PackerFormat: &packerFormat, PublishStr: &receiver}); err != nil {
				return err
			}
			values, err := Client.PubSub.GetStreamAll(sourceChannel, ciossdk.MakeGetStreamOpts().
				PackerFormat(packerFormat).TimestampRange(timestampRange).Label(label).Ascending(ascending), context.Background())
			if err != nil {
				log.Error(err.Error())
			} else {
				for index, value := range values {
					value := strings.Replace(value, sourceChannel, targetChannel, 1)
					time.Sleep(time.Millisecond * 5)
					receiver <- &value
					log.Info("Count ", strconv.Itoa(index))
					println(value)
				}
			}
			return nil
		},
	}
}
