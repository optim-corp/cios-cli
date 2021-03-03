package resourceowner

import (
	"context"
	"fmt"
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
			//deleteByResourceOwner(),
		},
	}
}

func listResourceOwner() *cli.Command {
	return &cli.Command{
		Name:    models.LIST,
		Aliases: models.ALIAS_LIST,
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "all_resource", Aliases: []string{"a", "r"}},
		},
		Action: func(c *cli.Context) error {
			isAll := c.Bool("all_resource")
			line := "---------------------------------------------------------------------------------------------------------------------------------\n"
			allFunc := func(ownerID string) {
				buckets, _, _ := Client.FileStorage.GetBucketsAll(ciossdk.MakeGetBucketsOpts().ResourceOwnerId(ownerID), context.Background())
				if len(buckets) > 0 {
					fPrintln("\n\n\t\t\t\t\t<Bucket>\n\n")
					fPrintln("\n\n\t|id|\t\t\t|resource_owner_id|\t\t |name|")
					for _, value := range buckets {
						fPrintln(value.Id, "\t", value.ResourceOwnerId, "\t", value.Name)
					}
					fPrintln(line)
				}
				channels, _, _ := Client.PubSub.GetChannelsAll(ciossdk.MakeGetChannelsOpts().ResourceOwnerId(ownerID), context.Background())
				if len(channels) > 0 {
					fPrintln("\n\n\t\t\t\t\t<PubSub Channel>\n\n")
					fPrintln("\t|id|\t\t\t|resource_owner_id|\t\t  |name|\t\t|labels|")
					for _, value := range channels {
						fPrintln(value.Id, "\t", value.ResourceOwnerId, "\t", value.Name, "\t", value.Labels)
					}
					fPrintln(line)
				}

				response, _, _ := Client.Geography.GetPoints(ciossdk.MakeGetPointsOpts().ResourceOwnerId(ownerID), context.Background())
				if len(response.Points) > 0 {
					fPrintln("\n\n\t\t\t\t\t<Geo Point>\n\n")
					fPrintln("\t|id|    \t\t|resource owner id|\t\t   |name -- latitude -- longitude -- altitude|\t\t|label|")
					for _, val := range response.Points {
						fmt.Fprintf(
							out,
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
					fPrintln(line)

				}
				fPrintln("*******************************************************************************************************************************************************************")
			}
			listUtility(func() {
				if c.Args().Len() > 0 {
					utils.CliArgsForEach(c, func(id string) { allFunc(id) })
				} else if isAll {
					owners, _, err := Client.Account.GetResourceOwnersAll(ciossdk.MakeGetResourceOwnersOpts(), context.Background())
					assert(err).Log().NoneErr(func() {
						for _, owner := range owners {
							fPrintf("|Resource Owner ID|: %s\n|Group ID|: %s\n|Name|: %s\n|Type|: %s\n", owner.Id, owner.GroupId, owner.Profile.DisplayName, owner.Type)
							allFunc(owner.Id)
						}
					})
				} else {
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
				}
			})

			return nil
		},
	}
}

//
//func deleteByResourceOwner() *cli.Command {
//	return &cli.Command{
//		Name:      models.DELETE,
//		Aliases:   models.ALIAS_DELETE,
//		UsageText: "cios resourceowner delete [command options] [resource owner id...]",
//		Action: func(c *cli.Context) error {
//			utils.CliArgsForEach(c, func(id string) {
//				println("\n\nBucket Deleting...")
//				buckets, _, err := Client.FileStorage.GetBucketsAll(ciossdk.MakeGetBucketsOpts().ResourceOwnerId(id), context.Background())
//				assert(err).
//					Log().
//					NoneErrPrintln("Total: ", len(buckets)).
//					NoneErr(func() {
//						for _, bucket := range buckets {
//							_, err := Client.FileStorage.DeleteBucket(bucket.Id, context.Background())
//							assert(err).Log().NoneErrPrintln("Delete: " + bucket.Id)
//						}
//					})
//				println("\n\nChannel Deleting...")
//				channels, _, err := Client.PubSub.GetChannelsAll(ciossdk.MakeGetChannelsOpts().ResourceOwnerId(id), context.Background())
//				assert(err).
//					Log().
//					NoneErrPrintln("Total: ", len(channels)).
//					NoneErr(func() {
//						for _, channel := range channels {
//							_, err := Client.PubSub.DeleteDataByChannel(channel.Id, context.Background())
//							assert(err).Log().
//								NoneErrPrintln("Data Store Completed: " + channel.Id).
//								NoneErrPrintln("Deleting: " + channel.Id).
//								NoneErr(func() {
//									_, err = Client.PubSub.DeleteChannel(channel.Id, context.Background())
//									assert(err).Log().NoneErrPrintln("Completed: " + channel.Id)
//
//								})
//						}
//					})
//			})
//			return nil
//		},
//	}
//}
