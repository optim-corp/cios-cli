package tool

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/optim-corp/cios-cli/utils/console"

	"github.com/fcfcqloow/go-advance/log"

	"github.com/optim-corp/cios-cli/models"
	"github.com/urfave/cli/v2"
)

func GetTimestampCommand() *cli.Command {
	return &cli.Command{
		Name:    "timestamp",
		Aliases: []string{"time", "ts"},
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "add", Aliases: []string{"a"}},
		},
		Before: func(c *cli.Context) error {
			timestampFormatFile = path(timestampFormatPath)
			return nil
		},
		Subcommands: []*cli.Command{
			{
				Name:    models.LIST,
				Aliases: models.ALIAS_LIST,
				Action: func(c *cli.Context) error {
					file, err := timestampFormatFile.ReadFile()
					if err != nil {
						log.Error(err.Error())
						return err
					}
					formats := strings.Split(string(file), "\n")
					println(formats)
					return nil
				},
			},
		},
		UsageText: "cios timestamp | time | ts",
		Action: func(c *cli.Context) error {
			if c.Bool("add") {
				timestampFormatFile.CreateFile()
				file, err := timestampFormatFile.ReadFile()
				assert(err).Log().NoneErr(func() {
					str := string(file)
					console.CliArgsForEach(c, func(format string) { str += format + "\n" })
					assert(timestampFormatFile.WriteFileAsString(str)).Log()
				})
				return nil
			}
			if _, err := os.Stat(timestampFormatPath); err != nil {
				str := ""
				str += time.ANSIC + "\n"
				str += time.UnixDate + "\n"
				str += time.RubyDate + "\n"
				str += time.RFC822 + "\n"
				str += time.RFC822Z + "\n"
				str += time.RFC850 + "\n"
				str += time.RFC1123 + "\n"
				str += time.RFC1123Z + "\n"
				str += time.RFC3339 + "\n"
				str += time.RFC3339Nano + "\n"
				str += time.Kitchen + "\n"
				str += time.Stamp + "\n"
				str += time.StampMilli + "\n"
				str += time.StampMicro + "\n"
				str += time.StampNano + "\n"
				timestampFormatFile.CreateFile()
				assert(timestampFormatFile.WriteFileAsString(str)).Log()
			}

			byt, _ := timestampFormatFile.ReadFile()
			formats := strings.Split(string(byt), "\n")
			console.CliArgsForEach(c, func(arg string) {
				arg = strings.ReplaceAll(arg, "\"", "")
				ok := false
				for _, format := range formats {
					log.Debug(format, arg)
					if result, err := time.Parse(format, arg); err == nil {
						println("Time: ", result)
						println("Unix: ", result.Unix())
						println("Unix Nano: ", result.UnixNano())
						ok = true
					}
				}
				if !ok {
					if integer, err := strconv.ParseInt(arg, 10, 64); err != nil {
						log.Error(err.Error())
					} else {
						nano := integer
						if integer/1e9 > 1 {
							integer = 0
						}
						result := time.Unix(integer, nano)
						println("Time: ", result)
						println("Unix: ", result.Unix())
						println("Unix Nano: ", result.UnixNano())
					}
				}
			})
			return nil
		},
	}
}
