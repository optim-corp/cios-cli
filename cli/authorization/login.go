package authorization

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/fcfcqloow/go-advance/ftil"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
	"github.com/skratchdot/open-golang/open"
	"github.com/urfave/cli/v2"
	"gopkg.in/AlecAivazis/survey.v1"
)

type input struct {
	Stage        string
	ClientID     string
	ClientSecret string
	RedirectUri  string
}

var (
	configPath = models.ConfigPath
	assert     = utils.EAssert
)

func GetLoginCommand() *cli.Command {
	return &cli.Command{
		Name:  "login",
		Usage: "cios login",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "client_type", Aliases: []string{"c"}},
		},
		Action: func(c *cli.Context) error {
			if c.Bool("client_type") {
				clientLogin()
				return nil
			}
			login()
			return nil
		},
	}
}

func setPath(stage string) string {
	if urls, ok := models.GetUrls(); ok {
		if _url, ok := urls[stage]; ok {
			return "https://" + _url.Auth
		}
	}
	return ""
}

func login() {
	answers := input{}
	stage := struct {
		Stage string
	}{}
	utils.Question([]*survey.Question{
		{
			Name: "stage",
			Prompt: &survey.Select{
				Message: "Choose a stage:",
				Options: models.Stages,
			},
		},
	}, &stage)
	utils.Question([]*survey.Question{
		{
			Name:   "redirectUri",
			Prompt: &survey.Input{Message: "Redirect URI: "},
		},
		{
			Name:   "clientID",
			Prompt: &survey.Input{Message: "Client ID: "},
		},
		{
			Name:   "clientSecret",
			Prompt: &survey.Password{Message: "Client Secret"},
		},
	}, &answers)
	answers.Stage = stage.Stage
	basePath := setPath(answers.Stage)
	port := answers.RedirectUri[16:21]
	path := answers.RedirectUri[21:]
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		form := url.Values{}
		form.Add("grant_type", "authorization_code")
		form.Add("code", r.FormValue("code"))
		form.Add("redirect_uri", answers.RedirectUri)
		form.Add("client_id", answers.ClientID)
		form.Add("client_secret", answers.ClientSecret)
		read := strings.NewReader(form.Encode())
		resp, err := http.Post(basePath+"/connect/token", "application/x-www-form-urlencoded", read)
		assert(err).Log().ExitWith(1).NoneErr(func() {
			var token struct {
				RefreshToken string `json:"refresh_token,omitempty"`
			}
			assert(json.NewDecoder(resp.Body).Decode(&token)).ExitWith(1).Log().NoneErr(func() {
				config := &models.Config{}
				config.Refresh = token.RefreshToken
				config.ClientID = answers.ClientID
				config.ClientSecret = answers.ClientSecret
				config.Stage = answers.Stage
				config.LogLevel = "info"
				config.AuthType = "refresh_token"
				assert(ftil.Path(configPath).WriteJson(config)).Log().NoneErrPrintln("\n\nfinish")
				os.Exit(0)
			})

		})
	})
	scope := models.FullScope
	_url := basePath +
		"/connect/authorize?response_type=code&redirect_uri=" +
		answers.RedirectUri +
		"&scope=" + scope +
		"&client_id=" + answers.ClientID
	println("req url: " + _url)
	assert(open.Start(_url)).Log().NoneErr(func() {
		assert(http.ListenAndServe(port, nil)).Log()
	})

}
func clientLogin() {
	stage := struct {
		Stage string
	}{}
	in := input{}
	utils.Question([]*survey.Question{
		{
			Name: "stage",
			Prompt: &survey.Select{
				Message: "Choose a stage:",
				Options: models.Stages,
			},
		},
	}, &stage)
	utils.Question([]*survey.Question{
		{
			Name:   "clientID",
			Prompt: &survey.Input{Message: "Client ID: "},
		},
		{
			Name:   "clientSecret",
			Prompt: &survey.Password{Message: "Client Secret"},
		},
	}, &in)
	config := &models.Config{}
	config.ClientID = in.ClientID
	config.ClientSecret = in.ClientSecret
	config.AuthType = "client"
	config.Stage = stage.Stage
	config.LogLevel = "info"
	config.Refresh = ""
	assert(ftil.Path(configPath).WriteJson(config)).
		Log().
		NoneErrPrintln("finish")
}
