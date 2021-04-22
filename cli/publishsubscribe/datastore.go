package publishsubscribe

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/optim-kazuhiro-seida/go-advance-type/wrapper"

	. "github.com/optim-corp/cios-cli/cli"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
	"github.com/optim-corp/cios-golang-sdk/cios"
	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"
	"github.com/optim-kazuhiro-seida/go-advance-type/convert"
	log "github.com/optim-kazuhiro-seida/loglog"
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
			saveDataStore(),
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
			&cli.StringFlag{Name: "packer_format", Aliases: []string{"pf"}, Value: "payload_only"},
			&cli.StringFlag{Name: "timestamp_range", Aliases: []string{"tr"}, DefaultText: "Now Time", Value: ":" + convert.MustStr(time.Now().UnixNano())},
			&cli.StringFlag{Name: "label", Aliases: []string{"lb"}},
			&cli.BoolFlag{Name: "data", Aliases: []string{"d"}},
			&cli.Int64Flag{Name: "limit", Aliases: []string{"l"}, DefaultText: "30", Value: 30},
			&cli.Int64Flag{Name: "offset", Aliases: []string{"o"}, DefaultText: "0", Value: 0},
			&cli.StringFlag{Name: "resource_owner_id", Aliases: []string{"r", "ro"}},
		},
		Action: func(c *cli.Context) error {
			var (
				channelID         = c.String("channel_id")
				packerFormat      = c.String("packer_format")
				limit             = c.Int64("limit")
				offset            = c.Int64("offset")
				timestampRange    = c.String("timestamp_range")
				resourceOwnerID   = c.String("resource_owner_id")
				label             = c.String("label")
				dataFlag          = c.Bool("data")
				channelsMap, _, _ = Client.PubSub.GetChannelsMapByID(ciossdk.MakeGetChannelsOpts(), context.Background())
			)

			printObject := func(channelId string) {
				objects, _, err := Client.PubSub.GetObjectsAll(channelId, ciossdk.MakeGetObjectsOpts().
					Limit(limit).
					TimestampRange(timestampRange).
					Label(label).
					Offset(offset),
					context.Background())
				if len(objects) == 0 || resourceOwnerID != "" && channelsMap[channelId].ResourceOwnerId != resourceOwnerID {
					return
				}
				fPrintf("\n|Channel ID|  : %s \n|Channel Name|: %s\n\n", channelId, channelsMap[channelId].Name)
				fPrintln("\t|ID|\t\t|Timestamp|\t  |Mime Type|")
				assert(err).
					Log().
					NoneErr(func() {
						for _, obj := range objects {
							fPrintf("%s %s %s\n", obj.Id, obj.Timestamp, obj.MimeType)
						}
					})

			}
			printData := func(channelId string) []string {
				data, err := Client.PubSub.GetStreamAll(channelId, ciossdk.MakeGetStreamOpts().
					Limit(limit).Offset(offset).
					PackerFormat(packerFormat).
					TimestampRange(timestampRange).
					Label(label),
					context.Background())
				assert(err).Log()
				fPrintf("\n|Channel ID|  : %s \n|Channel Name|: %s\n\n", channelId, channelsMap[channelId].Name)
				for _, val := range data {
					fPrintln(val)
				}
				return data
			}
			printJob := func(channel cios.Channel, limit int64) {
				switch {
				case dataFlag:
					printData(channel.Id)
				default:
					printObject(channel.Id)
				}
				assert(out.Flush()).Log()
			}

			listUtility(func() {
				if channelID != "" {
					channel, _, err := Client.PubSub.GetChannel(channelID, nil, nil, context.Background())
					assert(err).
						Log().
						NoneErr(func() { printJob(channel, limit) })
				} else {
					channels, _, err := Client.PubSub.GetChannelsAll(ciossdk.MakeGetChannelsOpts().ResourceOwnerId(resourceOwnerID), context.Background())
					assert(err).
						Log().
						NoneErr(func() {
							for _, channel := range channels {
								printJob(channel, limit)
							}
						})
				}
			})
			return nil
		},
	}
}
func saveDataStore() *cli.Command {
	return &cli.Command{
		Name:    "save",
		Aliases: []string{"download", "get", "sv", "dl"},
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "channel_id", Aliases: []string{"c"}},
			&cli.StringFlag{Name: "packer_format", Aliases: []string{"pf"}, Value: "payload_only"},
			&cli.StringFlag{Name: "timestamp_range", Aliases: []string{"tr"}, DefaultText: "Now Time", Value: ":" + convert.MustStr(time.Now().UnixNano())},
			&cli.StringFlag{Name: "save_dir", Aliases: []string{"out"}},
			&cli.StringFlag{Name: "label", Aliases: []string{"lb"}},
			&cli.BoolFlag{Name: "indent", Aliases: []string{"idt", "idnt", "i"}},
			&cli.BoolFlag{Name: "collective", Aliases: []string{"compact", "coll", "collect"}},
			&cli.Int64Flag{Name: "limit", Aliases: []string{"l"}, DefaultText: "30", Value: 30},
			&cli.Int64Flag{Name: "offset", Aliases: []string{"o"}, DefaultText: "0", Value: 0},
			&cli.StringFlag{Name: "resource_owner_id", Aliases: []string{"r", "ro"}},
			&cli.StringFlag{Name: "replace_save_data_channel", Aliases: []string{"replace", "rep"}, Usage: "-replace <channel id>"},
		},
		Action: func(c *cli.Context) error {
			var (
				channelID         = c.String("channel_id")
				packerFormat      = c.String("packer_format")
				limit             = c.Int64("limit")
				offset            = c.Int64("offset")
				timestampRange    = c.String("timestamp_range")
				resourceOwnerID   = c.String("resource_owner_id")
				label             = c.String("label")
				outputDir         = wrapper.AsString(c.String("save_dir"))
				replaced          = wrapper.AsString(c.String("replace_save_data_channel"))
				indent            = c.Bool("indent")
				collective        = c.Bool("collective")
				channelsMap, _, _ = Client.PubSub.GetChannelsMapByID(ciossdk.MakeGetChannelsOpts(), context.Background())
				stageDSDir        = fmt.Sprintf("%s/%s", datastoreDir, models.GetStage())
			)
			replaceChannelId := func(data string) string {
				var jsonFormat cios.PackerFormatJson
				assert(convert.UnMarshalJson([]byte(data), &jsonFormat)).Log()
				jsonFormat.Header.ChannelId = replaced.Str()
				return convert.MustCompactJson(jsonFormat)
			}
			indentJson := func(data string) string {
				var jsonFormat cios.PackerFormatJson
				assert(convert.UnMarshalJson([]byte(data), &jsonFormat)).Log()
				return convert.MustIndentJson(jsonFormat)
			}
			job := func(channel cios.Channel, limit int64) {
				switch {
				case replaced != "":
					packerFormat = "json"
					fallthrough
				case outputDir == "":
					outputDir = wrapper.String(fmt.Sprintf("%s/%s/%s___%s", datastoreDir, models.GetStage(), channelsMap[channel.Id].Name, channel.Id))
					fallthrough
				default:
					data, err := Client.PubSub.GetStreamAll(channel.Id, ciossdk.MakeGetStreamOpts().
						Limit(limit).Offset(offset).
						PackerFormat(packerFormat).
						TimestampRange(timestampRange).
						Label(label),
						context.Background())
					assert(err).Log()
					for idx, val := range data {
						switch {
						case replaced.IsPreset():
							data[idx] = replaceChannelId(val)
							fallthrough
						case indent:
							data[idx] = indentJson(val)
						}
					}
					path(datastoreDir).CreateDir()
					path(stageDSDir).CreateDir()
					path(outputDir.Str()).CreateDir()
					switch {
					case collective:
						fileName := fmt.Sprintf("%s/%s_%s.txt",
							outputDir,
							packerFormat,
							strings.Replace(timestampRange, ":", "-", -1))
						assert(path(fileName).WriteFileAsString("[\n"+strings.Join(data, ",")+"]\n")).
							Log().
							NoneErrPrintln("Completed ", fileName)
					default:
						for idx, val := range data {
							allLength := len(str(len(data)))
							currentLength := len(str(idx))
							filePrefixZero := strings.Repeat("0", allLength-currentLength)
							fileName := fmt.Sprintf("%s/%s%d.txt", outputDir, filePrefixZero, idx)
							assert(path(fileName).WriteFileAsString(val)).
								Log().
								NoneErrPrintln("Completed ", fileName)
						}
					}
				}
			}

			if channelID != "" {
				channel, _, err := Client.PubSub.GetChannel(channelID, nil, nil, context.Background())
				assert(err).
					Log().
					NoneErr(func() { job(channel, limit) })
			} else {
				channels, _, err := Client.PubSub.GetChannelsAll(ciossdk.MakeGetChannelsOpts().ResourceOwnerId(resourceOwnerID), context.Background())
				assert(err).
					Log().
					NoneErr(func() {
						for _, channel := range channels {
							job(channel, limit)
						}
					})
			}
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
