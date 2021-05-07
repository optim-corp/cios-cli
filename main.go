package main

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/dimiro1/banner"
	"github.com/fcfcqloow/go-advance/ftil"
	"github.com/fcfcqloow/go-advance/log"
	"github.com/mattn/go-colorable"
	. "github.com/optim-corp/cios-cli/cli"
	"github.com/optim-corp/cios-cli/cli/account"
	"github.com/optim-corp/cios-cli/cli/authorization"
	"github.com/optim-corp/cios-cli/cli/device"
	"github.com/optim-corp/cios-cli/cli/filestorage"
	"github.com/optim-corp/cios-cli/cli/geography"
	"github.com/optim-corp/cios-cli/cli/group"
	"github.com/optim-corp/cios-cli/cli/publishsubscribe"
	"github.com/optim-corp/cios-cli/cli/resourceowner"
	"github.com/optim-corp/cios-cli/cli/tool"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v2"
)

const (
	NAME             = "Could IoT OS CLI"
	VERSION          = "0.3.1"
	COPYRIGHT        = "OPTiM Corporation"
	APPLICATION_LOGO = `
{{.AnsiColor.Black}}
█████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████
████████{{.AnsiColor.Cyan}}██████████████{{.AnsiColor.Black}}█████{{.AnsiColor.Cyan}}████████{{.AnsiColor.Black}}████████{{.AnsiColor.Cyan}}███████████████{{.AnsiColor.Black}}███████████{{.AnsiColor.Cyan}}████████████████{{.AnsiColor.Black}}████████████████████████████████████████████████
█████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}████████████████████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}███████{{.AnsiColor.Cyan}}█████████████████████{{.AnsiColor.Black}}█████{{.AnsiColor.Cyan}}███████████████████{{.AnsiColor.Black}}████████████████████████████████████████████████
█████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}████████████████████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}███████{{.AnsiColor.Cyan}}█████{{.AnsiColor.Black}}███████████{{.AnsiColor.Cyan}}█████{{.AnsiColor.Black}}████{{.AnsiColor.Cyan}}██████{{.AnsiColor.Black}}██████████████████████████████████████████████████████████████
█████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}████████████████████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}███████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}█████████████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}█████{{.AnsiColor.Cyan}}█████{{.AnsiColor.Black}}██████████████████████████████████████████████████████████████
█████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}████████████████████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}███████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}█████████████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}███████{{.AnsiColor.Cyan}}█████{{.AnsiColor.Black}}████████████████████████████████████████████████████████████
█████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}████████████████████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}███████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}█████████████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}███████████{{.AnsiColor.Cyan}}██████████████{{.AnsiColor.Black}}███████████████████████████████████████████████
█████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}████████████████████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}███████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}█████████████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}██████████████████████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}██████████████████████████████████████████████
█████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}████████████████████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}███████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}█████████████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}███████████████████████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}█████████████████████████████████████████████
█████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}████████████████████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}███████{{.AnsiColor.Cyan}}█████{{.AnsiColor.Black}}███████████{{.AnsiColor.Cyan}}█████{{.AnsiColor.Black}}██████████████████████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}██████████████████████████████████████████████
█████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}████████████████████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}███████{{.AnsiColor.Cyan}}█████████████████████{{.AnsiColor.Black}}████████████████████{{.AnsiColor.Cyan}}████{{.AnsiColor.Black}}████████████████████████████████████████████████
████████{{.AnsiColor.Cyan}}██████████████{{.AnsiColor.Black}}█████{{.AnsiColor.Cyan}}████████{{.AnsiColor.Black}}████████{{.AnsiColor.Cyan}}███████████████{{.AnsiColor.Black}}███████████{{.AnsiColor.Cyan}}██████████████{{.AnsiColor.Black}}██████████████████████████████████████████████████
█████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████
█████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████
██████████████████████████████████████████████████████████████████{{.AnsiColor.BrightMagenta}}█████████████{{.AnsiColor.Black}}██████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}████████████████████{{.AnsiColor.BrightMagenta}}████████{{.AnsiColor.Black}}████████████████
████████████████████████████████████████████████████████████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}█████████████████████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}██████████████████████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}██████████████████
████████████████████████████████████████████████████████████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}█████████████████████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}██████████████████████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}██████████████████
████████████████████████████████████████████████████████████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}█████████████████████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}██████████████████████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}██████████████████
████████████████████████████████████████████████████████████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}█████████████████████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}██████████████████████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}██████████████████
████████████████████████████████████████████████████████████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}█████████████████████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}██████████████████████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}██████████████████
████████████████████████████████████████████████████████████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}█████████████████████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}██████████████████████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}██████████████████
████████████████████████████████████████████████████████████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}█████████████████████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}██████████████████████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}██████████████████
████████████████████████████████████████████████████████████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}█████████████████████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}██████████████████████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}██████████████████
████████████████████████████████████████████████████████████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}█████████████████████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}██████████████████████{{.AnsiColor.BrightMagenta}}████{{.AnsiColor.Black}}██████████████████
█████████████████████████████████████████████████████████████████{{.AnsiColor.BrightMagenta}}██████████████{{.AnsiColor.Black}}██████{{.AnsiColor.BrightMagenta}}█████████████████████{{.AnsiColor.Black}}███{{.AnsiColor.BrightMagenta}}████████{{.AnsiColor.Black}}████████████████
█████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████
`
)

var (
	config     models.Config
	topDir     = models.TopDir
	configPath = models.ConfigPath
	urlDir     = models.UrlPath
	is         = utils.Is
	path       = ftil.Path
	assert     = utils.EAssert
)

func init() {
	if os.Getenv("AUTH_TYPE") == "client" {
		setClientType()
	} else {
		assert(path(topDir).CreateDir()).Log()
		configFile := path(configPath)
		urlFile := path(urlDir)
		file, configErr := configFile.ReadFile()
		urls, urlErr := urlFile.ReadFile()
		assert(urlErr).
			OnErr(func() { assert(urlFile.CreateFile()).Log() }).
			OnErr(func() { assert(urlFile.WriteFileAsString(models.URL_JSON)) })
		assert(configErr).
			OnErr(func() { assert(configFile.CreateFile()).Log() }).
			ExitWith(0)
		assert(json.Unmarshal(file, &config)).
			NoneErr(func() { setConfig(config, urls, config.Stage) })
	}
}
func main() {
	app := &cli.App{
		Name:      NAME,
		Version:   VERSION,
		Copyright: COPYRIGHT,
		Action: func(context *cli.Context) error {
			banner.Init(colorable.NewColorableStdout(), true, true, bytes.NewBufferString(APPLICATION_LOGO))
			println("\n\n\nPlease $cios help !!!!\n\n")
			return nil
		},
		Commands: []*cli.Command{
			publishsubscribe.GetDataStoreCommand(),
			publishsubscribe.GetMessagingCommand(),
			publishsubscribe.GetChannelCommand(),
			device.GetDeviceMonitoringCommand(),
			device.GetDeviceInventoryCommand(),
			device.GetDeviceLifecycleCommand(),
			device.GetDevicePolicyCommand(),
			device.GetDeviceEntityCommand(),
			device.GetDeviceModelsCommand(),
			device.GetDeviceCommand(),
			resourceowner.GetResourceOwnerCommand(),
			authorization.GetLoginCommand(),
			filestorage.GetBucketCommand(),
			filestorage.GetFileCommand(),
			filestorage.GetNodeCommand(),
			geography.GetPointCommand(),
			group.GetGroupCommand(),
			//video.GetVideoCommand(),
			tool.GetTimestampCommand(),
			tool.GetConfigCommand(),
			tool.GetSwitchCommand(),
			tool.GetTokenCommand(),
			tool.GetURLCommand(),
			tool.GetLogCommand(),
			account.GetMeCommand(),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Error(err.Error())
	}
}
func setClientType() {
	logLevel := os.Getenv("LOG_LEVEL")
	log.SetLevelOrDefault(is(logLevel == "").T("info").F(logLevel).Value.(string), log.LOG_LEVEL_WARN)
	Client = ciossdk.NewCiosClient(ciossdk.CiosClientConfig{
		Debug:       log.GetLevel() == log.LOG_LEVEL_DEBUG,
		AuthConfig:  ciossdk.ClientAuthConf(os.Getenv("CIOS_CLIENT_ID"), os.Getenv("CIOS_CLIENT_SECRET"), os.Getenv("CIOS_SCOPE")),
		AutoRefresh: true,
		Urls: sdkmodel.CIOSUrl{
			MessagingUrl:             os.Getenv("CIOS_MESSAGING_URL"),
			LocationUrl:              os.Getenv("CIOS_LOCATION_URL"),
			AccountsUrl:              os.Getenv("CIOS_ACCOUNT_URL"),
			StorageUrl:               os.Getenv("CIOS_STORAGE_URL"),
			IamUrl:                   os.Getenv("CIOS_IAM_URL"),
			AuthUrl:                  os.Getenv("CIOS_AUTH_URL"),
			VideoStreamingUrl:        os.Getenv("CIOS_STREAMING_URL"),
			DeviceManagementUrl:      os.Getenv("CIOS_DEVICE_URL"),
			DeviceMonitoringUrl:      os.Getenv("CIOS_MONITORING_URL"),
			DeviceAssetManagementUrl: os.Getenv("CIOS_ASSET_URL"),
		},
	})
}
func setConfig(config models.Config, urls []byte, stage string) {
	_ = models.SetStage(stage)
	log.SetLevelOrDefault(config.LogLevel, log.LOG_LEVEL_WARN)
	https := "https://"
	domain := func(name string) string {
		return gjson.GetBytes(urls, stage+"."+name).String()
	}
	Client = ciossdk.NewCiosClient(ciossdk.CiosClientConfig{
		Debug: log.GetLevel() == log.LOG_LEVEL_DEBUG,
		AuthConfig: ciossdk.RefreshTokenAuth(
			config.ClientID,
			config.ClientSecret,
			config.Refresh,
			models.FullScope,
		),
		AutoRefresh: true,
		Urls: sdkmodel.CIOSUrl{
			MessagingUrl:             https + domain("messaging"),
			LocationUrl:              https + domain("location"),
			AccountsUrl:              https + domain("account"),
			StorageUrl:               https + domain("storage"),
			IamUrl:                   https + domain("iam"),
			AuthUrl:                  https + domain("auth"),
			VideoStreamingUrl:        https + domain("video_streaming"),
			DeviceMonitoringUrl:      https + domain("monitoring"),
			DeviceManagementUrl:      https + domain("device"),
			DeviceAssetManagementUrl: https + domain("device_asset"),
		},
	})
}
