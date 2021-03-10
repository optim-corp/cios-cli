package tool

import (
	app "github.com/optim-corp/cios-cli/cli"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
	"github.com/skratchdot/open-golang/open"
	"github.com/urfave/cli/v2"
	"gopkg.in/AlecAivazis/survey.v1"
)

func GetURLCommand() *cli.Command {
	return &cli.Command{
		Name:    "url",
		Aliases: []string{"uri", "domain"},
		Usage:   "cios url | uri",
		Action: func(c *cli.Context) error {
			listUtility(func() {
				fPrintln("|DeviceAssetManagement_URL| : " + app.Client.DeviceAssetManagement.Url)
				fPrintln("|DeviceManagement_URL|      : " + app.Client.DeviceManagement.Url)
				//fPrintln("|VideoStreams_URL|          : " + app.Client.Vide.Url)
				//fPrintln("|Monitoring_URL|            : " + app.Client.DeviceManagement.Url)
				fPrintln("|Messaging_URL|             : " + app.Client.PubSub.Url)
				fPrintln("|Location_URL|              : " + app.Client.Geography.Url)
				fPrintln("|Accounts_URL|              : " + app.Client.Account.Url)
				fPrintln("|Storage_URL|               : " + app.Client.FileStorage.Url)
				fPrintln("|Auth_URL|                  : " + app.Client.Auth.Url)
				//fPrintln("|Iam_URL|                   : " + app.Client.Iam.Url)
			})
			return nil
		},
		Subcommands: []*cli.Command{
			{
				Name:    models.PATCH,
				Aliases: []string{"up", "sync"},
				Action: func(c *cli.Context) error {
					urlFile := path(models.UrlPath)
					urlFile.CreateFile()
					utils.EAssert(urlFile.WriteFileAsString(models.URL_JSON)).Log()
					return nil
				},
			},
			addUrls(),
		},
	}
}

func addUrls() *cli.Command {
	return &cli.Command{
		Name:    "edit",
		Aliases: []string{"control"},
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "interactive", Aliases: []string{"i"}},
			&cli.BoolFlag{Name: "interactive_all", Aliases: []string{"ia", "ai"}},
			&cli.StringFlag{Name: "domain", Aliases: []string{"d"}},
			&cli.StringFlag{Name: "stage", Aliases: []string{"s"}},
			&cli.StringFlag{Name: "auth_url", Aliases: []string{"auth"}},
		},
		Action: func(c *cli.Context) error {
			var (
				in struct {
					models.URL
					Stage                 string
					Domain                string
					DeviceManagement      string
					DeviceAssetManagement string
					Monitoring            string
					Messaging             string
					Location              string
					Accounts              string
					Storage               string
					Iam                   string
					Auth                  string
					VideoStreams          string
				}
				domain           = c.String("domain")
				stage            = c.String("stage")
				authUrl          = c.String("auth_url")
				isInteractive    = c.Bool("interactive")
				isInteractiveAll = c.Bool("interactive_all")
				urls, ok         = models.GetUrls()
			)
			if !ok {
				panic("Not Found URL.json")
			}
			if domain != "" && stage != "" {
				urls[stage] = models.URL{
					DeviceManagement:      "device-management." + domain,
					DeviceAssetManagement: "device-asset-lifecycle." + domain,
					Monitoring:            "monitoring." + domain,
					Messaging:             "messaging." + domain,
					Location:              "location." + domain,
					Accounts:              "accounts." + domain,
					Storage:               "storage." + domain,
					Iam:                   "iam." + domain,
					Auth:                  is(authUrl == "").T("auth." + domain).F(authUrl).Value.(string),
					VideoStreams:          "video-streaming." + domain,
				}
			} else if isInteractive {
				utils.Question(
					[]*survey.Question{
						{Name: "stage", Prompt: &survey.Input{Message: "Stage: "}},
						{Name: "domain", Prompt: &survey.Input{Message: "Domain: "}},
					}, &in)
				utils.Question(
					[]*survey.Question{
						{Name: "auth", Prompt: &survey.Input{Message: "Auth URL: ", Default: "auth." + in.Domain}},
					}, &in)
				urls[in.Stage] = models.URL{
					DeviceManagement:      "device-management." + in.Domain,
					DeviceAssetManagement: "device-asset-lifecycle." + in.Domain,
					Monitoring:            "monitoring." + in.Domain,
					Messaging:             "messaging." + in.Domain,
					Location:              "location." + in.Domain,
					Accounts:              "accounts." + in.Domain,
					Storage:               "storage." + in.Domain,
					Iam:                   "iam." + in.Domain,
					Auth:                  in.Auth,
					VideoStreams:          "video-streaming." + in.Domain,
				}
			} else if isInteractiveAll {
				utils.Question(
					[]*survey.Question{
						{Name: "stage", Prompt: &survey.Input{Message: "Stage: "}},
						{Name: "device", Prompt: &survey.Input{Message: "Device URL: "}},
						{Name: "device_asset", Prompt: &survey.Input{Message: "Device Asset URL: "}},
						{Name: "monitoring", Prompt: &survey.Input{Message: "Monitoring URL: "}},
						{Name: "messaging", Prompt: &survey.Input{Message: "Messaging URL: "}},
						{Name: "location", Prompt: &survey.Input{Message: "Geography URL: "}},
						{Name: "account", Prompt: &survey.Input{Message: "Account URL: "}},
						{Name: "storage", Prompt: &survey.Input{Message: "File Storage URL: "}},
						{Name: "iam", Prompt: &survey.Input{Message: "Iam URL: "}},
						{Name: "auth", Prompt: &survey.Input{Message: "Auth URL: "}},
						{Name: "video_streaming", Prompt: &survey.Input{Message: "VideoStreaming URL: "}},
					}, &in)
				urls[in.Stage] = models.URL{
					DeviceManagement:      in.DeviceManagement,
					DeviceAssetManagement: in.DeviceAssetManagement,
					Monitoring:            in.Monitoring,
					Messaging:             in.Messaging,
					Location:              in.Location,
					Accounts:              in.Accounts,
					Storage:               in.Storage,
					Iam:                   in.Iam,
					Auth:                  in.Auth,
					VideoStreams:          in.VideoStreams,
				}
			} else {
				open.Start(models.UrlPath)
			}
			models.WriteUrls(urls)
			return nil
		},
	}
}
