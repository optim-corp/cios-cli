package account_test

import (
	"encoding/json"
	"flag"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/optim-corp/cios-cli/utils/xstring"

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
	cases := []struct {
		title    string
		me       cios.Me
		expected string
	}{
		{
			title: "Plot Name",
			me:    cios.Me{Name: cnv.StrPtr("test-name"), Email: "example@sample.com"},
			expected: "****************************************************************************************************************" +
				"|Name|: test-name|Email| : example@sample.com |group id||resource_owner_id||name / type|" +
				"****************************************************************************************************************",
		},
		{
			title: "Plot corporation",
			me:    cios.Me{Corporation: &cios.GroupChildren{Name: cnv.StrPtr("test-corpo")}},
			expected: "****************************************************************************************************************" +
				"|Name|: |Email| : |Corporation| : test-corpo |group id||resource_owner_id||name / type|" +
				"****************************************************************************************************************",
		},
		{
			title: "Plot group id",
			me: cios.Me{Corporation: &cios.GroupChildren{Name: cnv.StrPtr("test-corpo")}, Groups: &[]cios.MeGroups{
				{
					Id:            "test-group-id",
					Name:          "",
					Type:          "",
					CorporationId: cios.NullableString{},
				},
			}},
			expected: "****************************************************************************************************************" +
				"|Name|: |Email| : |Corporation| : test-corpo |group id||resource_owner_id||name / type|test-group-id / " +
				"****************************************************************************************************************",
		},
		{
			title: "Plot groups",
			me: cios.Me{Corporation: &cios.GroupChildren{Name: cnv.StrPtr("test-corpo")}, Groups: &[]cios.MeGroups{
				{
					Id:            "ead53ded-003d-4525-816b-e471f11814d9",
					Name:          "test1",
					Type:          "Group",
					CorporationId: cios.NullableString{},
				},
				{
					Id:            "ead53ded-003d-4525-816b-e471f11814d9",
					Name:          "test2",
					Type:          "Group",
					CorporationId: cios.NullableString{},
				},
				{
					Id:            "ead53ded-003d-4525-816b-e471f11814d9",
					Name:          "test2",
					Type:          "Group",
					CorporationId: cios.NullableString{},
				},
				{
					Id:            "ead53ded-003d-4525-816b-e471f11814d9",
					Name:          "test1",
					Type:          "Corporation",
					CorporationId: cios.NullableString{},
				},
			}},
			expected: "****************************************************************************************************************" +
				"|Name|: |Email| : |Corporation| : test-corpo |group id||resource_owner_id||name / type|ead53ded-003d-4525-816b-e471f11814d9test1 / Groupead53ded-003d-4525-816b-e471f11814d9test2 / Groupead53ded-003d-4525-816b-e471f11814d9test2 / Groupead53ded-003d-4525-816b-e471f11814d9test1 / Corporation" +
				"****************************************************************************************************************",
		},
	}
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			_, fin := mockCiosClientMe(c.me)
			t.Log(getConsoleResult().Str())
			assert.Equal(t, c.expected, xstring.ToOneLine(getConsoleResult().Str()))
			fin()
		})
	}
}
