package account

import (
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"

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
			value, _, err := Client.Account.GetMe(ciosctx.Background())
			if err != nil {
				return err
			}
			resourceOwnerMap, err := Client.Account.GetResourceOwnersMapByGroupID(ciosctx.Background())
			if err != nil {
				return nil
			}
			listUtility(func() {
				fPrintln("|Name|                : " + str(value.Name))
				fPrintln("|Email|               : " + value.Email)
				if !check.IsNil(value.Corporation) && !check.IsNil(value.Corporation.Name) {
					fPrintln("|Corporation|         : " + str(value.Corporation.Name))
				}
				fPrintln("\t     |group id|\t\t\t\t|resource_owner_id|\t\t|name / type|")
				if !check.IsNil(value.Groups) {
					// 1000件超えたら積
					for _, group := range *value.Groups {
						resourceOwner, _ := resourceOwnerMap[group.Id]
						fPrintln(group.Id + "\t" + resourceOwner.GetId() + "\t" + group.Name + " / " + group.Type)
					}
				}
			})
			return nil
		},
	}
}
