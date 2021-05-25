package filestorage

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"

	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"

	"github.com/optim-corp/cios-cli/utils/console"

	"github.com/fcfcqloow/go-advance/log"
	. "github.com/optim-corp/cios-cli/cli"
	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"
	"github.com/urfave/cli/v2"
)

func GetFileCommand() *cli.Command {
	return &cli.Command{
		Name:    "file",
		Aliases: []string{"f", "file"},
		Usage:   "cios file | f",
		Subcommands: []*cli.Command{
			uploadFile(),
			downloadFile(),
		},
	}
}
func downloadFile() *cli.Command {
	return &cli.Command{
		Name:      "download",
		Aliases:   []string{"get", "save"},
		UsageText: "cios  download [command options] [node id....]",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "bucket_id", Aliases: []string{"bucket", "b"}, Required: true},
			&cli.StringFlag{Name: "local_save_directory", Aliases: []string{"ld"}, Value: "./"},
		},
		Action: func(c *cli.Context) error {
			println("Download start....")
			var (
				bucketID = c.String("bucket_id")
				dir      = c.String("local_save_directory")
			)

			console.CliArgsForEach(c, func(nodeID string) {
				var (
					byt  []byte
					file *os.File
					err  error
				)
				defer func() { assert(file.Close()).Log() }()
				node, _, err := Client.FileStorage.GetNode(bucketID, nodeID, nil)
				assert(err).Log().
					NoneErrAssertFn(func() error {
						byt, _, err = Client.FileStorage.DownloadFile(ciosctx.Background(), bucketID, nodeID)
						return err
					}).Log().
					NoneErrAssertFn(func() error {
						file, err = os.Create(dir + node.Name)
						return err
					}).Log().
					NoneErrAssertFn(func() error {
						_, err = file.Write(byt)
						return err
					}).Log().
					NoneErrPrintln("Completed ", nodeID)
			})
			return nil
		},
	}
}

func uploadFile() *cli.Command {
	return &cli.Command{
		Name:      "upload",
		Aliases:   []string{"up", "add", "cre", "create", "make"},
		UsageText: "cios  download [command options] [local path....]",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "bucket_id", Aliases: []string{"b"}},
			&cli.StringFlag{Name: "bucket_name", Aliases: []string{"bn"}},
			&cli.StringFlag{Name: "bucket_resource_owner_id", Aliases: []string{"br", "bro"}},
			&cli.StringFlag{Name: "node_id", Aliases: []string{"n"}},
			&cli.StringFlag{Name: "directory", Aliases: []string{"d"}},
		},
		Action: func(c *cli.Context) error {
			var (
				directory           = c.String("directory")
				bucketID            = c.String("bucket_id")
				bucketName          = c.String("bucket_name")
				bucketResourceOwner = c.String("bucket_resource_owner_id")
				nodeID              = c.String("node_id")
			)

			if bucketID == "" {
				bucket, _, err := Client.FileStorage.GetBucketByResourceOwnerIDAndName(ciosctx.Background(), bucketResourceOwner, bucketName)
				if err != nil {
					return err
				}
				bucketID = bucket.Id
			}
			if directory != "" {
				if nodeID != "" {
					node, _, err := Client.FileStorage.GetNode(bucketID, nodeID, ciosctx.Background())
					if err != nil {
						log.Error(err.Error())
						return err
					}
					directoryUpload(bucketID, directory, node.Key)
				} else {
					directoryUpload(bucketID, directory, "/")
				}
				return nil
			}
			console.CliArgsForEach(c, func(localPath string) {
				byts, err := path(localPath).ReadFile()
				assert(err).Log().NoneErr(func() {
					_, err := Client.FileStorage.UploadFile(ciosctx.Background(), bucketID, byts, ciossdk.MakeUploadFileOpts().Name(filepath.Base(localPath)).NodeId(nodeID))
					assert(err).Log().NoneErrPrintln("Completed " + localPath)
				})
			})
			return nil
		},
	}
}

func directoryUpload(bucketID string, _dir string, key string) {
	var (
		dirs []fs.FileInfo
	)
	_dir, err := filepath.Abs(_dir)
	assert(err).Log().
		NoneErrAssertFn(func() error {
			dirs, err = ioutil.ReadDir(_dir)
			return err
		}).Log().
		NoneErr(func() {
			for _, dir := range dirs {
				f := string(key[len(key)-1])
				localPath := filepath.Join(_dir, dir.Name())
				updatedKey := key + is(f == "/" || f == "\\").T("").F("/").Value.(string) + dir.Name()
				if dir.IsDir() {
					directoryUpload(bucketID, localPath, updatedKey)
				} else {
					byts, err := path(localPath).ReadFile()
					assert(err).Log().
						NoneErrAssertFn(func() error {
							_, err := Client.FileStorage.UploadFile(ciosctx.Background(), bucketID, byts, ciossdk.MakeUploadFileOpts().Key(updatedKey))
							return err
						}).
						NoneErrPrintln("Completed " + updatedKey)
				}
			}
		})
}
