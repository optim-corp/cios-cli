package models

type Config struct {
	Refresh      string `json:"refresh"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	LogLevel     string `json:"log_level"`
	Stage        string `json:"stage"`
	AuthType     string `json:"auth_type"`
}

type Account = map[string]map[string]Config
