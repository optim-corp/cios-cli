package account_test

import (
	"encoding/json"
	"flag"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	wrp "github.com/fcfcqloow/go-advance/wrapper"

	xos "github.com/fcfcqloow/go-advance/os"

	cnv "github.com/fcfcqloow/go-advance/convert"

	"github.com/optim-corp/cios-golang-sdk/cios"

	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"

	app "github.com/optim-corp/cios-cli/cli"
	"github.com/optim-corp/cios-cli/utils/console"
	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"

	"github.com/urfave/cli/v2"

	"github.com/optim-corp/cios-cli/models"

	"github.com/optim-corp/cios-cli/cli/account"
	"github.com/stretchr/testify/assert"
)

func mockCiosClient(fun http.HandlerFunc) (server *httptest.Server, fin func()) {
	server = httptest.NewServer(fun)
	app.Client = ciossdk.NewCiosClient(ciossdk.CiosClientConfig{Urls: sdkmodel.CIOSUrl{AccountsUrl: server.URL}})
	fin = server.Close
	return
}

var ctx = cli.NewContext(cli.NewApp(), &flag.FlagSet{}, &cli.Context{})

func TestGetMeCommand(t *testing.T) {
	_, fin := mockCiosClient(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "GET" && r.URL.Path == "/v2/me" {
			json.NewEncoder(w).Encode(cios.Me{
				Id:    "test",
				Name:  cnv.StrPtr("test-name"),
				Email: "example@sample.com",
			})
		}
	})

	buff := xos.CaptureStdout(func() {
		console.SetWriter(os.Stdout)
		account.GetMeCommand().Subcommands[0].Action(ctx)
	})

	assert.Equal(t, "me", account.GetMeCommand().Name)
	assert.Equal(t, "list", account.GetMeCommand().Subcommands[0].Name)
	assert.Equal(t, models.ALIAS_LIST, account.GetMeCommand().Subcommands[0].Aliases)

	result := wrp.MakeString(buff.String())
	if !result.ContainsAll("test-name", "example@sample.com") {
		t.Fatal("Fail Plot Name")
	}
	fin()
	_, fin = mockCiosClient(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "GET" && r.URL.Path == "/v2/me" {
			json.NewEncoder(w).Encode(cios.Me{
				Groups: &[]cios.MeGroups{},
			})
		}
	})

}
