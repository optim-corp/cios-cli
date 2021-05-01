package filestorage

import (
	"strconv"

	"github.com/fcfcqloow/go-advance/log"

	xmath "github.com/fcfcqloow/go-advance/math"

	cnv "github.com/fcfcqloow/go-advance/convert"

	. "github.com/optim-corp/cios-cli/cli"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
	"github.com/optim-corp/cios-golang-sdk/cios"
	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"
	"github.com/urfave/cli/v2"
)

func GetBucketCommand() *cli.Command {
	return &cli.Command{
		Name:    "bucket",
		Aliases: []string{"b", "buckets"},
		Usage:   "cios bucket | b",
		Subcommands: []*cli.Command{
			createBucket(),
			deleteBucket(),
			listBucket(),
			updateBucket(),
		},
	}
}

func createBucket() *cli.Command {
	return &cli.Command{
		Name:      models.CREATE,
		Aliases:   models.ALIAS_CREATE,
		UsageText: "cios bucket create -resource_owner_id <Resource Owner ID> -name <name1> -name <name2> ",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "resource_owner_id", Aliases: []string{"r"}, Required: true},
			&cli.StringSliceFlag{Name: "name", Aliases: []string{"n"}, Required: true},
		},
		Action: func(c *cli.Context) error {
			var (
				resourceOwnerId = c.String("resource_owner_id")
				names           = c.StringSlice("name")
			)
			for _, name := range names {
				_, _, err := Client.FileStorage.CreateBucket(resourceOwnerId, name, nil)
				assert(err).Log().NoneErrPrintln("Completed ", name)
			}
			return nil
		},
	}
}
func deleteBucket() *cli.Command {
	return &cli.Command{
		Name:      models.DELETE,
		Aliases:   models.ALIAS_DELETE,
		UsageText: "cios bucket delete  [id...]",
		Flags:     []cli.Flag{},
		Action: func(c *cli.Context) error {
			utils.CliArgsForEach(c, func(id string) {
				_, err := Client.FileStorage.DeleteBucket(id, nil)
				assert(err).Log().NoneErrPrintln("Completed " + id)
			})
			return nil
		},
	}
}
func listBucket() *cli.Command {
	return &cli.Command{
		Name:    models.LIST,
		Aliases: models.ALIAS_LIST,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
			&cli.StringFlag{Name: "resource_owner_id", Aliases: []string{"r"}},
			&cli.BoolFlag{Name: "all", Aliases: []string{"a"}},
			&cli.StringFlag{Name: "order"},
			&cli.StringFlag{Name: "order_by"},
			&cli.Int64Flag{Name: "limit", Aliases: []string{"l"}, Value: 30},
			&cli.Int64Flag{Name: "offset", Aliases: []string{"o"}, Value: 0},
		},
		Action: func(c *cli.Context) error {
			var (
				name            = c.String("name")
				isAll           = c.Bool("all")
				resourceOwnerId = c.String("resource_owner_id")
				order           = c.String("order")
				orderBy         = c.String("order_by")
				limit           = xmath.MinInt64(c.Int64("limit"), 3000)
				offset          = c.Int64("offset")
			)
			buckets, _, err := Client.FileStorage.GetBucketsAll(ciossdk.MakeGetBucketsOpts().
				Name(name).
				ResourceOwnerId(resourceOwnerId).
				OrderBy(orderBy).
				Order(order).
				Limit(limit).
				Offset(offset), nil)
			assert(err).Log().NoneErr(func() {
				log.Info("Getting Size: ", strconv.Itoa(len(buckets)))
				listUtility(func() {
					if isAll {
						var resourceOwnerMap map[string]cios.ResourceOwner
						resourceOwnerMap, _, err = Client.Account.GetResourceOwnersMapByID(nil)
						assert(err).Log().NoneErr(func() {
							fPrintln("\t|ID|\t\t\t|Resource Owner ID|\t\t|Name| : |Resource Owner Name|")
							for _, value := range buckets {
								resourceOwner, _ := resourceOwnerMap[value.ResourceOwnerId]
								resourceOwnerName := cnv.MustStr(resourceOwner.Profile.DisplayName)
								fPrintln(value.Id+"\t"+value.ResourceOwnerId+"\t", value.Name, " : ", resourceOwnerName)
							}
						})

					} else {
						fPrintln("\t|ID|\t\t\t|Resource Owner ID|\t\t |Name|")
						for _, value := range buckets {
							fPrintln(value.Id + "\t" + value.ResourceOwnerId + "\t" + value.Name)
						}
					}

				})
			})
			return nil
		},
	}
}
func updateBucket() *cli.Command {
	return &cli.Command{
		Name:    models.PATCH,
		Aliases: models.ALIAS_PATCH,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "bucket_id", Aliases: []string{"b"}, Required: true},
			&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Required: true},
		},
		Action: func(c *cli.Context) error {
			var (
				bucketID = c.String("bucket_id")
				name     = c.String("")
			)
			_, err := Client.FileStorage.UpdateBucket(bucketID, name, nil)
			assert(err).Log().NoneErrPrintln("Completed " + bucketID)
			return nil
		},
	}
}
