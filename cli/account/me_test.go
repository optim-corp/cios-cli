package account_test

import (
	"encoding/json"
	"flag"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/optim-corp/cios-golang-sdk/cios"

	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"

	app "github.com/optim-corp/cios-cli/cli"
	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"

	"github.com/urfave/cli/v2"

	"github.com/optim-corp/cios-cli/models"

	"github.com/optim-corp/cios-cli/cli/account"
	"github.com/stretchr/testify/assert"
)

func TestGetMeCommand(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != "GET" {
			t.Fatal(r.Method)
		}
		if r.URL.Path == "/v2/me" {
			response := cios.Me{
				Id: "test",
			}
			json.NewEncoder(w).Encode(response)
		}
	}))
	defer ts.Close()
	app.Client = ciossdk.NewCiosClient(ciossdk.CiosClientConfig{Urls: sdkmodel.CIOSUrl{AccountsUrl: ts.URL}})
	cmd := account.GetMeCommand()
	assert.Equal(t, "me", cmd.Name)
	assert.Equal(t, "list", cmd.Subcommands[0].Name)
	assert.Equal(t, models.ALIAS_LIST, cmd.Subcommands[0].Aliases)
	ctx := cli.NewContext(cli.NewApp(), &flag.FlagSet{}, &cli.Context{})
	cmd.Subcommands[0].Action(ctx)

}
