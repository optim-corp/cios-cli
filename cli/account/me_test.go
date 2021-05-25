package account_test

import (
	"encoding/json"
	"flag"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	cnv "github.com/fcfcqloow/go-advance/convert"

	xos "github.com/fcfcqloow/go-advance/os"
	wrp "github.com/fcfcqloow/go-advance/wrapper"
	app "github.com/optim-corp/cios-cli/cli"
	"github.com/optim-corp/cios-cli/cli/account"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils/console"
	"github.com/optim-corp/cios-golang-sdk/cios"
	sdkmodel "github.com/optim-corp/cios-golang-sdk/model"
	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func mockCiosClient(fun http.HandlerFunc) (server *httptest.Server, fin func()) {
	server = httptest.NewServer(fun)
	app.Client = ciossdk.NewCiosClient(ciossdk.CiosClientConfig{Urls: sdkmodel.CIOSUrl{AccountsUrl: server.URL}})
	fin = server.Close
	return
}
func mockCiosClientMe(me cios.Me) (server *httptest.Server, fin func()) {
	return mockCiosClient(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "GET" && r.URL.Path == "/v2/me" {
			json.NewEncoder(w).Encode(me)
		}
	})
}

var (
	ctx              = cli.NewContext(cli.NewApp(), &flag.FlagSet{}, &cli.Context{})
	getConsoleResult = func() *wrp.String {
		b := xos.CaptureStdout(func() {
			console.SetWriter(os.Stdout)
			account.GetMeCommand().Subcommands[0].Action(ctx)
		})
		return wrp.MakeString(b.String())
	}
)

func TestGetMeCommand(t *testing.T) {
	assert.Equal(t, "me", account.GetMeCommand().Name)
	assert.Equal(t, "list", account.GetMeCommand().Subcommands[0].Name)
	assert.Equal(t, models.ALIAS_LIST, account.GetMeCommand().Subcommands[0].Aliases)

	t.Run("Plot Name", func(t *testing.T) {
		var (
			name  = "test-name"
			email = "example@sample.com"
		)

		_, fin := mockCiosClientMe(cios.Me{Name: &name, Email: email})
		defer fin()

		assert.True(t, getConsoleResult().ContainsAll(name, email))
	})
	t.Run("Plot corporation", func(t *testing.T) {
		var (
			name = "test-corpo"
		)
		_, fin := mockCiosClientMe(cios.Me{Corporation: &cios.GroupChildren{Name: cnv.StrPtr("test-corpo")}})
		defer fin()

		assert.True(t, getConsoleResult().ContainsAll(name))
	})

}
