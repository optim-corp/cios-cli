package geography

// import (
// 	"context"
// 	"fmt"
// 	"os"
// 	"strings"

//
// 	. "github.com/optim-corp/cios-cli/client"
// 	. "github.com/optim-corp/cios-cli/utils"
// 	auth "github.com/optim-corp/cios-cli/authorization"
// 	"github.com/optim-corp/cios-cli/resourceowner"

// 	"gopkg.in/AlecAivazis/survey.v1"
// 	"github.com/urfave/cli/v2"
// )

// func CreateRouteCommand() *cli.Command {
// 	return &cli.Command{
// 		Name:    "route",
// 		Aliases: []string{"rt"},
// 		Usage:   "cios route || rt",
// 		Before: func(c *cli.Context) error {
// 			RefreshAccessToken()
// 			SetPath(os.Getenv("Location_URL"))
// 			return nil
// 		},
// 		Subcommands: []*cli.Command{
// 			listRoute(),
// 			deleteRoute(),
// 			createRoute(),
// 		},
// 	}
// }

// func createRoute() *cli.Command {
// 	return &cli.Command{
// 		Name:    models.CREATE,
// 		Aliases: models.ALIAS_CREATE,
// 		Flags:   []cli.Flag{},
// 		Action: func(c *cli.Context) error {
// 			qs := []*survey.Question{
// 				{
// 					Name:   "name",
// 					Prompt: &survey.Input{Message: "name: "},
// 				},
// 				{
// 					Name:   "latitude",
// 					Prompt: &survey.Input{Message: "latitude: "},
// 				},
// 				{
// 					Name:   "longitude",
// 					Prompt: &survey.Input{Message: "longitude: "},
// 				},
// 				{
// 					Name:   "altitude",
// 					Prompt: &survey.Input{Message: "altitude: "},
// 				},
// 				{
// 					Name:   "language",
// 					Prompt: &survey.Input{Message: "language: ", Default: "ja"},
// 				},
// 				{
// 					Name:   "isDefault",
// 					Prompt: &survey.Confirm{Message: "is default", Default: true},
// 				},
// 				{
// 					Name:   "resourceOwnerID",
// 					Prompt: &survey.Input{Message: "resource owner id: ", Default: GetMyResourceOwnerId()},
// 				},
// 				{
// 					Name:   "label",
// 					Prompt: &survey.Input{Message: "label(key=value): "},
// 				},
// 			}
// 			SetPath(os.Getenv("Location_URL"))
// 			answers := struct {
// 				Name            string
// 				Description     string
// 				Latitude        float32
// 				Longitude       float32
// 				Altitude        float32
// 				Language        string
// 				IsDefault       bool
// 				ResourceOwnerId string
// 				Label           string
// 			}{}
// 			survey.Ask(qs, &answers)
// 			labelExp := func(exp bool) []cios.Label {
// 				if exp {
// 					return []cios.Label{
// 						cios.Label{
// 							Key:   strings.Split(answers.Label, "=")[0],
// 							Value: strings.Split(answers.Label, "=")[1],
// 						},
// 					}
// 				}
// 				return []cios.Label{}

// 			}

// 			_, _, err := ApiClient.GeographyApi.CreateRoute(
// 				context.Background(),
// 				cios.RouteRequest{
// 					Location:        cios.Location{Latitude: answers.Latitude, Longitude: answers.Longitude},
// 					Altitude:        answers.Altitude,
// 					ResourceOwnerId: answers.ResourceOwnerId,
// 					DisplayInfo: []cios.DisplayInfo{
// 						cios.DisplayInfo{
// 							Name:        answers.Name,
// 							Language:    answers.Language,
// 							Description: answers.Description,
// 							IsDefault:   answers.IsDefault,
// 						},
// 					},
// 					Labels: labelExp(answers.Label != ""),
// 				},
// 			)
// 			if err != nil {
// 				Log.Error(err.Error())
// 			}
// 			return nil
// 		},
// 	}
// }
// func deleteRoute() *cli.Command {
// 	return &cli.Command{
// 		Name:      models.DELETE,
// 		Aliases: models.ALIAS_DELETE,
// 		Action: func(c *cli.Context) error {
// 			for i := 0; i < c.Args().Len(); i++ {
// 				_, _, err := ApiClient.GeographyApi.DeleteRoute(
// 					context.Background(),
// 					c.Args().Get(i),
// 				)
// 				if err != nil {
// 					Log.Info("Cannot delete a " + c.Args().Get(i))
// 				}
// 			}
// 			return nil
// 		},
// 	}
// }
// func listRoute() *cli.Command {
// 	return &cli.Command{
// 		Name: models.LIST,
// 		Aliases: models.ALIAS_LIST,
// 		Action: func(c *cli.Context) error {
// 			ListUtility(func() {
// 				fmt.Fprintln(Out, "\t|id|    \t\t|resource owner id|\t\t   |name -- latitude -- longitude -- altitude|\t\t|label|")
// 				value, _, _ := ApiClient.GeographyApi.GetRoutes(context.Background(), &cios.GetRoutesOpts{})
// 				for _, val := range value.Routes {
// 					fmt.Fprintf(
// 						Out,
// 						"%s\t%s       %s -- %f -- %f --%f\t%s\n",
// 						val.Id,
// 						val.ResourceOwnerId,
// 						val.Name,
// 						val.Location.Latitude,
// 						val.Location.Longitude,
// 						val.Altitude,
// 						val.Labels,
// 					)
// 				}
// 			})
// 			return nil
// 		},
// 	}
// }
