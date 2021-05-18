package geography

import (
	"context"
	"strings"

	"github.com/optim-corp/cios-cli/utils/console"

	"github.com/AlecAivazis/survey/v2"
	. "github.com/optim-corp/cios-cli/cli"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-golang-sdk/cios"
	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"
	"github.com/urfave/cli/v2"
)

func GetPointCommand() *cli.Command {
	return &cli.Command{
		Name:    "point",
		Aliases: []string{"pt"},
		Subcommands: []*cli.Command{
			listPoint(),
			deletePoint(),
			createPoint(),
		},
	}
}

func createPoint() *cli.Command {
	return &cli.Command{
		Name:    models.CREATE,
		Aliases: models.ALIAS_CREATE,
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			answers := struct {
				Name            string
				Description     string
				Latitude        float32
				Longitude       float32
				Altitude        float32
				Language        string
				IsDefault       bool
				ResourceOwnerID string
				Label           string
			}{}
			console.Question([]*survey.Question{
				{
					Name:   "name",
					Prompt: &survey.Input{Message: "name: "},
				},
				{
					Name:   "latitude",
					Prompt: &survey.Input{Message: "latitude: "},
				},
				{
					Name:   "longitude",
					Prompt: &survey.Input{Message: "longitude: "},
				},
				{
					Name:   "altitude",
					Prompt: &survey.Input{Message: "altitude: "},
				},
				{
					Name:   "language",
					Prompt: &survey.Input{Message: "language: ", Default: "ja"},
				},
				{
					Name:   "isDefault",
					Prompt: &survey.Confirm{Message: "is default", Default: true},
				},
				{
					Name:   "resourceOwnerID",
					Prompt: &survey.Input{Message: "resource owner id: "},
				},
				{
					Name:   "label",
					Prompt: &survey.Input{Message: "label(key=value): "},
				},
			}, &answers)
			labelExp := func(exp bool) []cios.Label {
				if exp {
					return []cios.Label{
						{
							Key:   strings.Split(answers.Label, "=")[0],
							Value: strings.Split(answers.Label, "=")[1],
						},
					}
				}
				return []cios.Label{}

			}
			labels := labelExp(answers.Label != "")
			request := cios.PointRequest{
				Location:        &cios.Location{Latitude: answers.Latitude, Longitude: answers.Longitude},
				Altitude:        &answers.Altitude,
				ResourceOwnerId: answers.ResourceOwnerID,
				DisplayInfo: &[]cios.DisplayInfo{
					{
						Name:        answers.Name,
						Language:    answers.Language,
						Description: &answers.Description,
						IsDefault:   answers.IsDefault,
					},
				},
				Labels: &labels,
			}
			point, _, err := Client.Geography.CreatePoint(request, context.Background())
			assert(err).Log().NoneErr(func() { console.OutStructJson(point) })
			return nil
		},
	}
}
func deletePoint() *cli.Command {
	return &cli.Command{
		Name:    models.DELETE,
		Aliases: models.ALIAS_DELETE,
		Action: func(c *cli.Context) error {
			console.CliArgsForEach(c, func(id string) {
				_, _, err := Client.Geography.DeletePoint(id, context.Background())
				assert(err).Log().NoneErrPrintln("Completed ", id)
			})
			return nil
		},
	}
}
func listPoint() *cli.Command {
	return &cli.Command{
		Name:    models.LIST,
		Aliases: models.ALIAS_LIST,
		Flags:   []cli.Flag{
			// &cli.BoolFlag{Name: "gui", Aliases: []string{"g"}},
		},
		Action: func(c *cli.Context) error {
			listUtility(func() {
				fPrintln("\t|id|    \t\t|resource owner id|\t\t   |name -- latitude -- longitude -- altitude|\t\t|label|")
				response, _, err := Client.Geography.GetPoints(ciossdk.MakeGetPointsOpts(), context.Background())
				assert(err).Log().NoneErr(func() {
					for _, val := range response.Points {
						fPrintf(
							"%s\t%s       %s -- %f -- %f --%f\t%s\n",
							val.Id,
							val.ResourceOwnerId,
							val.Name,
							val.Location.Latitude,
							val.Location.Longitude,
							val.Altitude,
							val.Labels,
						)
					}
				})
			})
			return nil
		},
	}
}
