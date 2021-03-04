package resourceowner

import (
	"context"
	"unicode/utf8"

	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"

	. "github.com/optim-corp/cios-cli/cli"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
	"github.com/urfave/cli/v2"
)

var (
	out         = utils.Out
	listUtility = utils.ListUtility
	spaceRight  = utils.SpaceRight
	fPrintln    = utils.Fprintln
	fPrintf     = utils.Fprintf
	assert      = utils.EAssert
)

func GetResourceOwnerCommand() *cli.Command {
	return &cli.Command{
		Name:    "resourceowner",
		Aliases: []string{"ro", "RO"},
		Usage:   "cios resourceowner | ro",
		Subcommands: []*cli.Command{
			listResourceOwner(),
		},
	}
}

func listResourceOwner() *cli.Command {
	return &cli.Command{
		Name:    models.LIST,
		Aliases: models.ALIAS_LIST,
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			listUtility(func() {
				fPrintln("\t\t|id|\t\t\t\t|group_id|\t\t\t\t|user_id|                        |author_id|                  |profile|")
				ros, _, err := Client.Account.GetResourceOwnersAll(ciossdk.MakeGetResourceOwnersOpts(), context.Background())
				assert(err).Log().NoneErr(func() {
					length := utf8.RuneCountInString("000000000000000000000000000000000000")
					for _, val := range ros {
						groupId := ""
						userId := ""
						authorId := ""
						if val.GroupId != nil {
							groupId = *val.GroupId
						}
						if val.UserId != nil {
							userId = *val.UserId
						}
						if val.AuthorId != nil {
							authorId = *val.AuthorId
						}
						fPrintf("%s %s %s %s  %s\n",
							spaceRight(val.Id, length),
							spaceRight(groupId, length),
							spaceRight(userId, length),
							spaceRight(authorId, length),
							*val.Profile.DisplayName)
					}
				})
			})
			return nil
		},
	}
}
