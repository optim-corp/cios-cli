package tool

import (
	app "github.com/optim-corp/cios-cli/cli"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
	"github.com/urfave/cli/v2"
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
					urlFile := path(urlDir)
					urlFile.CreateFile()
					utils.EAssert(urlFile.WriteFileAsString(models.URL_JSON)).Log()
					return nil
				},
			},
		},
	}
}
