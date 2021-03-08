package models

type (
	Config struct {
		Refresh      string `json:"refresh"`
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		LogLevel     string `json:"log_level"`
		Stage        string `json:"stage"`
		AuthType     string `json:"auth_type"`
	}
	URL struct {
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

	Account = map[string]map[string]Config
	URLs    = map[string]URL
)
