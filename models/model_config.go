package models

import (
	"os"

	"github.com/optim-corp/cios-cli/utils"
	"github.com/optim-kazuhiro-seida/ftil"
	"github.com/optim-kazuhiro-seida/go-advance-type/convert"
)

type Config struct {
	Refresh      string `json:"refresh"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	LogLevel     string `json:"log_level"`
	Stage        string `json:"stage"`
	AuthType     string `json:"auth_type"`
}
type Account = map[string]map[string]Config
type URLs = map[string]URL
type URL struct {
	DeviceManagement      string `json:"device"`
	DeviceAssetManagement string `json:"device_asset"`
	Monitoring            string `json:"monitoring"`
	Messaging             string `json:"messaging"`
	Location              string `json:"location"`
	Accounts              string `json:"account"`
	Storage               string `json:"storage"`
	Iam                   string `json:"iam"`
	Auth                  string `json:"auth"`
	VideoStreams          string `json:"video_streaming"`
}

const (
	StageStr = "STAGE"
)

func GetStage() string            { return os.Getenv(StageStr) }
func SetStage(stage string) error { return os.Setenv(StageStr, stage) }

func createStages() []string {
	var urls URLs
	if err := ftil.Path(utils.UrlPath).LoadJsonStruct(&urls); err != nil {
		return make([]string, 0)
	}
	return convert.GetObjectKeys(urls)

}

func GetConfig() (config Config, ok bool) {
	ok = ftil.Path(utils.ConfigPath).LoadJsonStruct(&config) == nil
	return
}

func GetUrls() (urls URLs, ok bool) {
	ok = ftil.Path(utils.UrlPath).LoadJsonStruct(&urls) == nil
	return
}

func GetAccounts() (account Account, ok bool) {
	ok = ftil.Path(utils.AccountPath).LoadJsonStruct(&account) == nil
	return
}

func WriteConfig(config Config) bool {
	return ftil.Path(utils.ConfigPath).WriteJson(config) == nil
}

func WriteUrls(urls URLs) bool {
	return ftil.Path(utils.UrlPath).WriteJson(urls) == nil
}

func WriteAccounts(account Account) bool {
	return ftil.Path(utils.AccountPath).WriteJson(account) == nil
}
