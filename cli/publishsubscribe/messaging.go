package publishsubscribe

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	cnv "github.com/fcfcqloow/go-advance/convert"
	"github.com/fcfcqloow/go-advance/log"

	"gopkg.in/yaml.v2"

	. "github.com/optim-corp/cios-cli/cli"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
	"github.com/optim-corp/cios-golang-sdk/cios"
	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
)

func GetMessagingCommand() *cli.Command {
	return &cli.Command{
		Name:    "messaging",
		Aliases: []string{"ms", "MS"},
		Usage:   "cios messaging | ms",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "subscribe", Aliases: []string{"s"}},
			&cli.BoolFlag{Name: "publish", Aliases: []string{"p"}},
			&cli.StringFlag{Name: "resource_owner_id", Aliases: []string{"r", "ro"}},
			&cli.StringFlag{Name: "packer_format", Aliases: []string{"pf"}, Value: "payload_only"},
		},
		Subcommands: []*cli.Command{
			listChannels(),
			publishMessage(),
			autoMessage(),
			registerJob(),
		},
		Action: func(c *cli.Context) error {
			var (
				eg              = errgroup.Group{}
				pub             = c.Bool("publish")
				sub             = c.Bool("subscribe")
				packerFormat    = c.String("packer_format")
				resourceOwnerID = c.String("resource_owner_id")
				channelIDs      []string
				exit            = false
				str             = make(chan *string)
				done            = make(chan bool)
				subscribeLogic  = func(body []byte) (bool, error) {
					println(string(body))
					return exit, nil
				}
				scanLogic = func(message string) {
					for {
						in := utils.GetConsoleMultipleLine(message)
						if in == "/exit" {
							done <- true
							exit = true
							break
						}
						go func() { str <- &in }()
					}
					os.Exit(0)
				}
			)
			if resourceOwnerID != "" {
				if res, _, err := Client.PubSub.GetChannelsAll(ciossdk.MakeGetChannelsOpts().ResourceOwnerId(resourceOwnerID), context.Background()); assert(err).
					NoneErr(func() {
						for _, channel := range res {
							channelIDs = append(channelIDs, channel.Id)
						}
					}).
					Log().
					ErrNotNil() {
					return err
				}
			} else {
				channelIDs = utils.CliArgs(c)
			}
			go func() {
				err := eg.Wait()
				if err != nil {
					log.Error(err.Error())
				}
			}()

			if sub {
				for _, channelID := range channelIDs {
					eg.Go(func() error {
						return Client.PubSub.ConnectWebSocket(channelID, done, ciossdk.ConnectWebSocketOptions{
							PackerFormat:  &packerFormat,
							SubscribeFunc: &subscribeLogic,
							Context:       context.Background(),
						})
					})
				}
				scanLogic("Press /exit")
			} else if pub {
				for _, channelID := range channelIDs {
					eg.Go(func() error {
						return Client.PubSub.ConnectWebSocket(channelID, done, ciossdk.ConnectWebSocketOptions{
							PackerFormat: &packerFormat,
							PublishStr:   &str,
							Context:      context.Background(),
						})
					})

				}
				scanLogic("Press publish message or /exit")
			} else {
				for _, channelID := range channelIDs {
					eg.Go(func() error {
						return Client.PubSub.ConnectWebSocket(channelID, done, ciossdk.ConnectWebSocketOptions{
							PackerFormat:  &packerFormat,
							PublishStr:    &str,
							SubscribeFunc: &subscribeLogic,
							Context:       context.Background(),
						})
					})

				}
				scanLogic("Press publish message or /exit")
			}
			return nil
		},
	}
}

func listChannels() *cli.Command {
	return &cli.Command{
		Name:    models.LIST,
		Aliases: models.ALIAS_LIST,
		Usage:   "cios messaging  ls | ms ls",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "resource_owner_id", Aliases: []string{"r"}},
			&cli.BoolFlag{Name: "datastore", Aliases: []string{"d"}},
		},
		Action: func(c *cli.Context) error {
			resourceOwnerID := c.String("resource_owner_id")
			channels, _, err := Client.PubSub.GetChannelsAll(ciossdk.MakeGetChannelsOpts().ResourceOwnerId(resourceOwnerID), context.Background())
			assert(err).Log().NoneErr(func() {
				if c.Bool("datastore") {
					println(datastoreDir)
					utils.ListDirs(datastoreDir, "|-")
					return
				}

				listUtility(func() {
					fPrintln("   |channel id|\t\t\t|resource owner id|\t\t    |data store enabled / messaging persisted|\t\t  |name|")
					for _, value := range channels {
						fPrintln(value.Id, "\t", value.ResourceOwnerId, "\t\t  ", value.DatastoreConfig.Enabled, "\t/\t", value.MessagingConfig.Persisted, "\t\t\t", value.Name)
					}
				})
			})

			return nil
		},
	}
}

func publishMessage() *cli.Command {
	return &cli.Command{
		Name:    "publish",
		Aliases: []string{"add", "pub", "push", "create"},
		Usage:   "cios messaging publish | pub",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "directory", Aliases: []string{"d"}},
			&cli.StringFlag{Name: "file", Aliases: []string{"f"}},
			&cli.StringFlag{Name: "packer_format", Aliases: []string{"pf"}, Value: "payload_only"},
			&cli.BoolFlag{Name: "file_reverse", Aliases: []string{"fr"}},
			&cli.IntFlag{Name: "millisecond_interval", Aliases: []string{"i", "interval", "millisecond"}, Value: 1000},
		},
		Action: func(c *cli.Context) error {
			if c.Args().Len() == 0 {
				log.Emergency("No Channel ID params")
				return errors.New("No Channel ID params")
			}
			var (
				eg           = errgroup.Group{}
				directory    = c.String("directory")
				file         = c.String("file")
				packerFormat = c.String("packer_format")
				reverse      = c.Bool("file_reverse")
				interval     = c.Int("millisecond_interval")
				str          = make(chan *string)
				done         = make(chan bool)
			)
			utils.CliArgsForEach(c, func(channelID string) {
				eg.Go(func() error {
					return Client.PubSub.ConnectWebSocket(channelID, done, ciossdk.ConnectWebSocketOptions{
						PackerFormat: &packerFormat,
						PublishStr:   &str,
					})
				})
			})
			go func() { assert(eg.Wait()).Log().ExitWith(1) }()
			if file != "" {
				val, err := path(file).ReadString()
				assert(err).Log().
					NoneErr(func() {
						str <- &val
						println("Completed ", val)
					})
			} else if directory != "" {
				files, err := path(directory).ReadDir()
				assert(err).Log().NoneErr(func() {
					if reverse {
						for i := len(files) - 1; i >= 0; i-- {
							file := files[i]
							time.Sleep(time.Millisecond * time.Duration(interval))
							mes := string(file.Value)
							str <- &mes
							println("Completed ", mes)
						}
					} else {
						for _, file := range files {
							time.Sleep(time.Millisecond * time.Duration(interval))
							mes := string(file.Value)
							str <- &mes
							println("Completed ", mes)
						}
					}

				})
			} else {
				input := utils.GetConsoleMultipleLine("")
				str <- &input
			}
			str <- nil
			return nil
		},
	}
}

func autoMessage() *cli.Command {
	return &cli.Command{
		Name:    "bot",
		Aliases: []string{"auto"},
		Usage:   "cios messaging bot | auto",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "channel_id", Aliases: []string{"chid", "cid", "ch", "id", "c"}, Required: true},
			&cli.StringFlag{Name: "file", Aliases: []string{"f"}},
			&cli.StringFlag{Name: "directory", Aliases: []string{"d"}},
			&cli.StringFlag{Name: "packer_format", Aliases: []string{"pf"}, Value: "payload_only"},
			&cli.Int64Flag{Name: "interval", Aliases: []string{"i"}, Usage: "Milli Second", Value: 1000},
		},
		Action: func(c *cli.Context) error {
			var (
				file         = c.String("file")
				channelId    = c.String("channel_id")
				directory    = c.String("directory")
				packerFormat = c.String("packer_format")
				interval     = c.Int64("interval")
				pubStr       = make(chan *string)
				isTypeFile   = false
				done         = make(chan bool)
				eg           = errgroup.Group{}
			)
			if file != "" {
				isTypeFile = true
			} else if directory != "" {
				isTypeFile = false
			} else {
				log.Error("File and Directory are empty.")
				os.Exit(1)
			}
			go func() { assert(eg.Wait()).Log().ExitWith(1) }()
			eg.Go(func() error {
				return Client.PubSub.ConnectWebSocket(channelId, done, ciossdk.ConnectWebSocketOptions{
					PackerFormat: &packerFormat,
					PublishStr:   &pubStr,
				})
			})
			var arr []string
			if isTypeFile {
				if f, err := path(file).ReadString(); err != nil {
					log.Error(err.Error())
				} else {
					arr = []string{f}
				}
			} else {
				if dir, err := path(directory).ReadDir(); err != nil {
					log.Error(err.Error())
				} else {
					for index, f := range dir {
						log.Info(strconv.Itoa(index), ":", f.AbsPath)
						arr = append(arr, string(f.Value))
					}
				}
			}
			if len(arr) == 0 {
				panic("No File")
			}
			for {
				for _, v := range arr {
					pubStr <- &v
					println("Publish \n" + v)
					time.Sleep(time.Millisecond * time.Duration(interval))
				}
			}
			return nil
		},
	}
}

func registerJob() *cli.Command {
	return &cli.Command{
		Name:    "job",
		Aliases: []string{"register", "reg", "j"},
		Usage:   "cios messaging job | register",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "path", Aliases: []string{"file_path", "f", "p", "file"}},
			&cli.Int64Flag{Name: "beginning_timestamp", Aliases: []string{"start", "b", "start_timestamp", "beginning"}, Value: time.Now().UnixNano(), DefaultText: "Now Timestamp"},
		},
		Action: func(c *cli.Context) error {
			var (
				filePath           = c.String("file_path")
				beginningTimestamp = c.Int64("beginning_timestamp")
				jobYaml            models.Job
				byts, err          = path(filePath).ReadFile()
				wg                 = sync.WaitGroup{}
			)
			fmt.Println(filePath, beginningTimestamp)
			check := func(jobYaml models.Job) error {
				for _, jobs := range jobYaml {
					for _, job := range jobs {
						for _, v := range job.Value {
							var formatJson cios.PackerFormatJson
							if err := cnv.UnMarshalJson([]byte(v.Data), &formatJson); err != nil {
								return err
							}

						}
					}
				}
				return nil
			}
			send := func(v models.MessagingJobValue, formatJson cios.PackerFormatJson, loop int, SendJson func(interface{}) error) error {
				for i, cnt := int64(0), 0; cnt < loop; i++ {
					switch {
					case i%v.Timestamp == 0:
						formatJson.Header.Timestamp = str(beginningTimestamp + i)
						fallthrough
					case v.Timestamp == -1 || i%v.Timestamp == 0:
						fmt.Println("Publish", formatJson.Header.Timestamp)
						if err := SendJson(formatJson); err != nil {
							return err
						}
						cnt++
					}
					time.Sleep(time.Nanosecond)
				}
				return nil
			}
			execJob := func(jobs models.MessagingJobs, name string) {
				defer wg.Done()
				for _, job := range jobs {
					ms := Client.PubSub.NewMessaging(job.Channel, "publish", "json")
					if err := ms.Start(context.Background()); err != nil {
						log.Error(err)
						return
					}
					fmt.Println(fmt.Sprintf("Start Job: %s", name))

					for _, v := range job.Value {
						var formatJson cios.PackerFormatJson
						_ = cnv.UnMarshalJson([]byte(v.Data), &formatJson)
						if err := send(v, formatJson, job.Loop, ms.SendJson); err != nil {
							break
						}
					}

				}
			}
			return assert(err).Log().
				NoneErrAssert(yaml.Unmarshal(byts, &jobYaml)).Log().
				NoneErrAssert(check(jobYaml)).Log().
				NoneErr(func() {
					for name, jobs := range jobYaml {
						wg.Add(1)
						go execJob(jobs, name)
					}
					wg.Wait()
				}).
				Err
		},
	}
}
