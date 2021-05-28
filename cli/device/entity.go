package device

import (
	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"

	"github.com/optim-corp/cios-cli/utils/console"

	xmath "github.com/fcfcqloow/go-advance/math"

	. "github.com/optim-corp/cios-cli/cli"
	"github.com/optim-corp/cios-cli/models"
	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"
	"github.com/urfave/cli/v2"
)

func GetDeviceEntityCommand() *cli.Command {
	return &cli.Command{
		Name:    "entity",
		Aliases: []string{"ent", "ett", "entt"},
		Usage:   "cios device entity | ett | entt",
		Subcommands: []*cli.Command{
			listDeviceEntities(),
			deleteDeviceEntity(),
		},
	}
}

func listDeviceEntities() *cli.Command {
	return &cli.Command{
		Name:    models.LIST,
		Aliases: models.ALIAS_LIST,
		Usage:   "cios device entity list | entity ls",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "all", Aliases: []string{"a"}},
			&cli.StringFlag{Name: "resource_owner_id", Aliases: []string{"r"}},
			&cli.StringFlag{Name: "order"},
			&cli.StringFlag{Name: "order_by"},
			&cli.Int64Flag{Name: "limit", Aliases: []string{"l"}, Value: 30},
			&cli.Int64Flag{Name: "offset", Aliases: []string{"o"}, Value: 0},
		},
		Action: func(c *cli.Context) error {
			var (
				resourceOwnerId = c.String("resource_owner_id")
				order           = c.String("order")
				orderBy         = c.String("order_by")
				limit           = xmath.MinInt64(c.Int64("limit"), 3000)
				offset          = c.Int64("offset")
				isAll           = c.Bool("all")
			)
			m, _, err := Client.DeviceAssetManagement.GetEntitiesAll(ciosctx.Background(), ciossdk.MakeGetEntitiesOpts().
				Limit(limit).
				Offset(offset).
				Order(order).
				OrderBy(orderBy).
				ResourceOwnerId(resourceOwnerId))
			assert(err).Log().NoneErr(func() {
				listUtility(func() {
					if isAll {
						for _, model := range m {
							fPrintln("|Device Model ID|  : ", model.Id)
							fPrintln("|Resource Owner ID|: ", model.ResourceOwnerId)
							fPrintln("|Component ID|     : ", model.Components.Get().Id)
							fPrintln("|Key|             : ", model.Key, "\n")
							console.FOutStructJsonSlim(model)
							fPrintln("\n------------------------------------------------------------------------------------------------------------------------------------------------------------\n")
						}
					} else {
						fPrintln("\t|id|\t\t         |resource owner id|\t\t   |component id|\t\t|key|")
						for _, model := range m {
							fPrintln(model.Id, "\t", model.ResourceOwnerId, "\t  ", model.Components.Get().Id, "\t", model.Key)
						}
					}

				})
			})
			return nil
		},
	}
}

func deleteDeviceEntity() *cli.Command {
	return &cli.Command{
		Name:    models.DELETE,
		Aliases: models.ALIAS_DELETE,
		Usage:   "cios device model delete  ett del",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "key", Aliases: []string{"k"}},
		},
		Action: func(c *cli.Context) error {
			var (
				key = c.String("key")
			)
			entityMap, err := Client.DeviceAssetManagement.GetEntitiesMapByID(ciosctx.Background(), ciossdk.MakeGetEntitiesOpts())
			assert(err).Log().NoneErr(func() {
				if entity, ok := entityMap[key]; ok {
					_, err = Client.DeviceAssetManagement.DeleteEntity(ciosctx.Background(), entity.Key)
					assert(err).NoneErrPrintln("Completed ", entity.Key)
				} else {
					_, err = Client.DeviceAssetManagement.DeleteEntity(ciosctx.Background(), key)
					assert(err).NoneErrPrintln("Completed ", key)
				}
			})
			return nil
		},
	}
}
