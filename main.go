package main

import (
	"encoding/json"
	"os"

	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
	log "github.com/optim-kazuhiro-seida/loglog"

	"github.com/optim-corp/cios-cli/cli/device"

	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"

	. "github.com/optim-corp/cios-cli/cli"
	"github.com/optim-corp/cios-cli/cli/account"
	"github.com/optim-corp/cios-cli/cli/authorization"
	"github.com/optim-corp/cios-cli/cli/filestorage"
	"github.com/optim-corp/cios-cli/cli/geography"
	"github.com/optim-corp/cios-cli/cli/group"
	"github.com/optim-corp/cios-cli/cli/publishsubscribe"
	"github.com/optim-corp/cios-cli/cli/resourceowner"
	"github.com/optim-corp/cios-cli/cli/tool"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"

	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v2"
)

const (
	NAME    = "Could CLI"
	VERSION = "0.0.1"
)

var (
	config     models.Config
	topDir     = utils.TopDir
	configPath = utils.ConfigPath
	urlDir     = utils.UrlPath
	is         = utils.Is
	path       = utils.Path
	assert     = utils.EAssert
)

func init() {
	if os.Getenv("AUTH_TYPE") == "client" {
		setClientType()
	} else {
		path(topDir).CreateDir()
		configFile := path(configPath)
		urlFile := path(urlDir)
		file, configErr := configFile.ReadFile()
		urls, urlErr := urlFile.ReadFile()
		assert(urlErr).
			OnErr(urlFile.CreateFile).
			OnErr(func() { assert(urlFile.WriteFileAsString(models.URL_JSON)) })
		assert(configErr).
			OnErr(configFile.CreateFile).
			ExitWith(0)
		assert(json.Unmarshal(file, &config)).
			NoneErr(func() { setConfig(config, urls, config.Stage) })
	}
}
func main() {
	app := &cli.App{
		Name:    NAME,
		Version: VERSION,
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
	logLevel = is(logLevel == "").T("info").F(logLevel).Value.(string)
	switch logLevel {
	case "emergency":
		utils.Log.LogLevel = utils.EMERGENCY_LOG_LEVEL
	case "error":
		utils.Log.LogLevel = utils.ERROR_LOG_LEVEL
	case "warn":
		utils.Log.LogLevel = utils.WARN_LOG_LEVEL
	case "info":
		utils.Log.LogLevel = utils.INFO_LOG_LEVEL
	case "debug":
		utils.Log.LogLevel = utils.DEBUG_LOG_LEVEL
	}
	Client = ciossdk.NewCiosClient(ciossdk.CiosClientConfig{
		Debug:       logLevel == "debug",
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
	_ = utils.SetStage(stage)
	https := "https://"
	domain := func(name string) string {
		return gjson.GetBytes(urls, stage+"."+name).String()
	}
	switch config.LogLevel {
	case "emergency":
		utils.Log.LogLevel = utils.EMERGENCY_LOG_LEVEL
	case "error":
		utils.Log.LogLevel = utils.ERROR_LOG_LEVEL
	case "warn":
		utils.Log.LogLevel = utils.WARN_LOG_LEVEL
	case "info":
		utils.Log.LogLevel = utils.INFO_LOG_LEVEL
	case "debug":
		utils.Log.LogLevel = utils.DEBUG_LOG_LEVEL
	}
	Client = ciossdk.NewCiosClient(ciossdk.CiosClientConfig{
		Debug: config.LogLevel == "debug",
		AuthConfig: ciossdk.RefreshTokenAuth(
			config.ClientID,
			config.ClientSecret,
			config.Refresh,
			models.FullScope,
		),
		AutoRefresh: true,
		Urls: sdkmodel.CIOSUrl{
			MessagingUrl:             https + domain("Messaging"),
			LocationUrl:              https + domain("Location"),
			AccountsUrl:              https + domain("Accounts"),
			StorageUrl:               https + domain("Storage"),
			IamUrl:                   https + domain("Iam"),
			AuthUrl:                  https + domain("Auth"),
			VideoStreamingUrl:        https + domain("VideoStreams"),
			DeviceMonitoringUrl:      https + domain("Monitoring"),
			DeviceManagementUrl:      https + domain("DeviceManagement"),
			DeviceAssetManagementUrl: https + domain("DeviceAssetManagement"),
		},
	})
}
