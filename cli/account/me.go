package account

import (
	"context"

	"github.com/optim-corp/cios-cli/utils/console"

	"github.com/fcfcqloow/go-advance/check"

	cnv "github.com/fcfcqloow/go-advance/convert"

	. "github.com/optim-corp/cios-cli/cli"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
	"github.com/urfave/cli/v2"
)

var (
	listUtility = console.ListUtility
	fPrintln    = console.Fprintln
	str         = cnv.MustStr
	assert      = utils.EAssert
)

func GetMeCommand() *cli.Command {
	return &cli.Command{
		Name:  "me",
		Usage: "cios me",
		Subcommands: []*cli.Command{
			listMe(),
		},
	}
}
func listMe() *cli.Command {
	return &cli.Command{
		Name:    models.LIST,
		Aliases: models.ALIAS_LIST,
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			listUtility(func() {
				value, _, err := Client.Account.GetMe(context.Background())
				assert(err).Log().NoneErr(func() {
					fPrintln("|Name|                : " + str(value.Name))
					fPrintln("|Email|               : " + value.Email)
					if !check.IsNil(value.Corporation) && !check.IsNil(value.Corporation.Name) {
						fPrintln("|Corporation|         : " + str(value.Corporation.Name))
					}
					fPrintln("\t     |group id|\t\t\t\t|resource_owner_id|\t\t|name / type|")
					if value.Groups != nil {
						// 1000件超えたら積
						resourceOwnerMap, err := Client.Account.GetResourceOwnersMapByGroupID(nil)
						assert(err).Log()
						for _, group := range *value.Groups {
							resourceOwner, ok := resourceOwnerMap[group.Id]
							resourceOwnerId := ""
							if ok {
								resourceOwnerId = resourceOwner.Id
							}
							assert(err).Log().NoneErr(func() {
								fPrintln(group.Id + "\t" + resourceOwnerId + "\t" + group.Name + " / " + group.Type)
							})
						}
					}
				})
			})
			return nil
		},
	}
}
