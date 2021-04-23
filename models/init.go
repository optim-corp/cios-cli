package models

import (
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/optim-corp/cios-cli/utils"
	"github.com/optim-kazuhiro-seida/go-advance-type/convert"
	ftil "github.com/optim-kazuhiro-seida/go-advance-type/file"
)

var (
	Dir, _                  = homedir.Dir()
	TopDir                  = Dir + "/.cios-cli"
	DatastoreDir            = TopDir + "/datastore"
	UrlPath                 = utils.Is(os.Getenv("CIOS_CLI_URL_PATH") == "").T(TopDir + "/URL.json").F(os.Getenv("CIOS_CLI_URL_PATH")).Value.(string)
	ConfigPath              = utils.Is(os.Getenv("CIOS_CLI_CONFIG_PATH") == "").T(TopDir + "/config.json").F(os.Getenv("CIOS_CLI_CONFIG_PATH")).Value.(string)
	AccountPath             = TopDir + "/accounts.json"
	TimestampFormatFilePath = TopDir + "/.timestamp_format"
	LifecycleDir            = TopDir + "/lifecycle"
)

const (
	StageStr = "STAGE"
)

func GetStage() string            { return os.Getenv(StageStr) }
func SetStage(stage string) error { return os.Setenv(StageStr, stage) }

func createStages() []string {
	var urls URLs
	if err := ftil.Path(UrlPath).LoadJsonStruct(&urls); err != nil {
		return make([]string, 0)
	}
	return convert.GetObjectKeys(urls)

}

func GetConfig() (config Config, ok bool) {
	ok = ftil.Path(ConfigPath).LoadJsonStruct(&config) == nil
	return
}

func GetUrls() (urls URLs, ok bool) {
	ok = ftil.Path(UrlPath).LoadJsonStruct(&urls) == nil
	return
}

func GetAccounts() (account Account, ok bool) {
	ok = ftil.Path(AccountPath).LoadJsonStruct(&account) == nil
	return
}

func WriteConfig(config Config) bool {
	return ftil.Path(ConfigPath).WriteJson(config) == nil
}

func WriteUrls(urls URLs) bool {
	return ftil.Path(UrlPath).WriteJson(urls) == nil
}

func WriteAccounts(account Account) bool {
	return ftil.Path(AccountPath).WriteJson(account) == nil
}
