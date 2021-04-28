package filestorage

import (
	"context"
	"strconv"
	"unicode/utf8"

	. "github.com/optim-corp/cios-cli/cli"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
	"github.com/optim-corp/cios-cli/utils/go_advance_type/convert"
	log "github.com/optim-corp/cios-cli/utils/loglog"
	"github.com/optim-corp/cios-golang-sdk/cios"
	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"
	"github.com/urfave/cli/v2"
)

func GetNodeCommand() *cli.Command {
	return &cli.Command{
		Name:    "node",
		Aliases: []string{"n", "nodes"},
		Usage:   "cios node | n",
		Subcommands: []*cli.Command{
			createNode(),
			deleteNode(),
			listNode(),
			copyNode(),
			moveNode(),
			renameNode(),
		},
	}
}

func createNode() *cli.Command {
	return &cli.Command{
		Name:      models.CREATE,
		Aliases:   models.ALIAS_CREATE,
		UsageText: "cios node create [command options] [name...]",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "bucket_id", Aliases: []string{"b"}, Required: true},
			&cli.StringFlag{Name: "parent_node_id", Aliases: []string{"p"}},
		},
		Action: func(c *cli.Context) error {
			var parentNodeId = c.String("parent_node_id")
			utils.CliArgsForEach(c, func(name string) {
				_, _, err := Client.FileStorage.CreateNodeOnNodeID(c.String("bucket_id"), cios.NodeRequest{
					Name:         name,
					ParentNodeId: &parentNodeId,
				}, context.Background())
				assert(err).Log().NoneErrPrintln("Completed " + name)
			})
			return nil
		},
	}
}
func deleteNode() *cli.Command {
	return &cli.Command{
		Name:      models.DELETE,
		Aliases:   models.ALIAS_DELETE,
		UsageText: "cios node delete [command options] [id...]",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "bucket_id", Aliases: []string{"b"}, Required: true},
		},
		Action: func(c *cli.Context) error {
			bucketID := c.String("bucket_id")
			utils.CliArgsForEach(c, func(nodeID string) {
				_, err := Client.FileStorage.DeleteNode(bucketID, nodeID, context.Background())
				assert(err).Log().NoneErrPrintln("Completed " + nodeID)
			})
			return nil
		},
	}
}
func listNode() *cli.Command {
	return &cli.Command{
		Name:    models.LIST,
		Aliases: models.ALIAS_LIST,
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "all", Aliases: []string{"a"}},
			&cli.StringFlag{Name: "parent_node_id", Aliases: []string{"p"}},
			&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
			&cli.Int64Flag{Name: "limit", Aliases: []string{"l"}, Value: 30},
		},
		Action: func(c *cli.Context) error {
			var (
				name         = c.String("name")
				parentNodeID = c.String("parent_node_id")
				all          = c.Bool("all")
				limit        = c.Int64("limit")
				opts         = ciossdk.MakeGetNodesOpts().Name(name).Recursive(all).ParentNodeId(parentNodeID).Limit(limit)
			)
			lsNode := func(value cios.Bucket) {
				nodes, _, err := Client.FileStorage.GetNodesAll(value.Id, opts, context.Background())
				if len(nodes) == 0 {
					return
				}
				if err != nil {
					log.Error(err.Error())
					return
				}
				fPrintf("\n\n|Bucket ID        |: %s\n|Bucket Name      |: %s\n|Resource Owner ID|: %s\n", value.Id, value.Name, value.ResourceOwnerId)
				fPrintln("\n\t\t|ID|\t\t  |Parent Node ID|  |Is directory|\t|Name : Key|")
				for _, node := range nodes {
					parentNodeID := is(node.ParentNodeId.Get() == nil).
						T(spaceRight("null", utf8.RuneCountInString("0000000000000000000"))).
						F(convert.MustStr(node.ParentNodeId.Get())).Value.(string)
					fPrintf(
						"\t%s\t%s\t%s\t\t%s : %s\n",
						node.Id,
						parentNodeID,
						strconv.FormatBool(node.IsDirectory),
						node.Name,
						node.Key,
					)
				}
				fPrintln("\n\n--------------------------------------------------------------------------------------------------------------")
			}
			listUtility(func() {
				if c.Args().Len() > 0 {
					utils.CliArgsForEach(c, func(bucketID string) {
						res, _, err := Client.FileStorage.GetBucket(bucketID, nil)
						assert(err).Log().NoneErr(func() { lsNode(res) })
					})
				} else {
					buckets, _, err := Client.FileStorage.GetBucketsAll(ciossdk.MakeGetBucketsOpts().Limit(30), context.Background())
					assert(err).Log().NoneErr(func() {
						for _, value := range buckets {
							lsNode(value)
						}
					})
				}
			})
			return nil
		},
	}
}
func copyNode() *cli.Command {
	return &cli.Command{
		Name:      "copy",
		Aliases:   []string{"copy", "cp"},
		UsageText: "cios node copy [command options] [id...]",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "bucket_id", Aliases: []string{"b"}, Required: true},
			&cli.StringFlag{Name: "dest_bucket_id", Aliases: []string{"d"}},
			&cli.StringFlag{Name: "parent_node_id", Aliases: []string{"p"}},
		},
		Action: func(c *cli.Context) error {
			var (
				bucketId     = c.String("bucket_id")
				destBucketId = c.String("dest_bucket_id")
				parentNodeId = c.String("parent_node_id")
			)
			utils.CliArgsForEach(c, func(nodeID string) {
				_, _, err := Client.FileStorage.CopyNode(
					bucketId,
					nodeID,
					&destBucketId,
					&parentNodeId,
					context.Background(),
				)
				assert(err).Log().NoneErrPrintln("Completed ", nodeID)
			})
			return nil
		},
	}
}
func moveNode() *cli.Command {
	return &cli.Command{
		Name:      "move",
		Aliases:   []string{"move", "mv", "mo", "mn"},
		UsageText: "cios node move [command options] [id...]",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "bucket_id", Aliases: []string{"b"}, Required: true},
			&cli.StringFlag{Name: "dest_bucket_id", Aliases: []string{"d"}},
			&cli.StringFlag{Name: "parent_node_id", Aliases: []string{"p"}},
		},
		Action: func(c *cli.Context) error {
			var (
				bucketId     = c.String("bucket_id")
				destBucketId = c.String("dest_bucket_id")
				parentNodeId = c.String("parent_node_id")
			)
			utils.CliArgsForEach(c, func(nodeID string) {
				_, _, err := Client.FileStorage.MoveNode(
					bucketId,
					nodeID,
					&destBucketId,
					&parentNodeId,
					context.Background(),
				)
				assert(err).Log().NoneErrPrintln("Completed " + nodeID)
			})
			return nil
		},
	}
}
func renameNode() *cli.Command {
	return &cli.Command{
		Name:      "rename",
		Aliases:   []string{"rename", "re", "rn"},
		UsageText: "cios node rename [command options] [id...]",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "bucket_id", Aliases: []string{"b"}, Required: true},
			&cli.StringFlag{Name: "node_id", Aliases: []string{"n"}, Required: true},
			&cli.StringFlag{Name: "name", Aliases: []string{"nm"}},
		},
		Action: func(c *cli.Context) error {
			name := c.String("name")
			_, _, err := Client.FileStorage.RenameNode(c.String("bucket_id"), c.String("node_id"), name, context.Background())
			assert(err).Log().NoneErrPrintln("Completed ", name)
			return nil
		},
	}
}
