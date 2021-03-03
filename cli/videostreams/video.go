package video

//func GetVideoCommand() *cli.Command {
//	return &cli.Command{
//		Name:    "video",
//		Aliases: []string{"v"},
//		Usage:   "cios video | v",
//		Subcommands: []*cli.Command{
//			listVideo(),
//			deleteVideo(),
//			createVideo(),
//			playVideo(),
//			stopVideo(),
//		},
//	}
//}
//
//func createVideo() *cli.Command {
//	return &cli.Command{
//		Name:    models.CREATE,
//		Aliases: models.ALIAS_CREATE,
//		Flags:   []cli.Flag{},
//		Action: func(c *cli.Context) error {
//			return nil
//		},
//	}
//}
//func deleteVideo() *cli.Command {
//	return &cli.Command{
//		Name:    models.DELETE,
//		Aliases: models.ALIAS_DELETE,
//		Action: func(c *cli.Context) error {
//			return nil
//		},
//	}
//}
//func listVideo() *cli.Command {
//	return &cli.Command{
//		Name:    models.LIST,
//		Aliases: models.ALIAS_LIST,
//		Action: func(c *cli.Context) error {
//			printVideos := func(videos []cios.Video) {
//				if len(videos) == 0 {
//					return
//				}
//				for _, v := range videos {
//					fPrintf("%s\t%s\t%s\t%s\t%s\n", v.Id, v.ResourceOwnerId, v.DeviceId, v.Enabled, v.Name)
//				}
//			}
//			if c.Args().Len() == 0 {
//				ros, _ := client.GetResourceOwners()
//				utils.ListUtility(func() {
//					fPrintln("\t|id|    \t\t|resource owner id|\t\t   |device id|\t\t   |enabled|\t|name|\t")
//					for _, ro := range ros {
//						videos, err := client.VideoList(ro.Id)
//						assert(err).NoneErr(func() { printVideos(videos) })
//					}
//				})
//			} else {
//				utils.ListUtility(func() {
//					fPrintln("\t|id|    \t\t|resource owner id|\t\t   |device id|\t\t   |enabled|\t|name|\t")
//					utils.CliArgsForEach(c, func(id string) {
//						videos, err := client.VideoList(id)
//						assert(err).NoneErr(func() { printVideos(videos) })
//					})
//				})
//
//			}
//			return nil
//		},
//	}
//}
//func playVideo() *cli.Command {
//	return &cli.Command{
//		Name:    "play",
//		Aliases: []string{"p", "up", "start"},
//		Action: func(c *cli.Context) error {
//			utils.CliArgsForEach(c, func(videoID string) {
//				_, err := client.VideoPlay(videoID)
//				assert(err).NoneErr(func() { println("Completed ", videoID) })
//			})
//			return nil
//		},
//	}
//}
//func stopVideo() *cli.Command {
//	return &cli.Command{
//		Name:    "stop",
//		Aliases: []string{"st", "shutdown"},
//		Flags: []cli.Flag{
//			&cli.StringFlag{Name: "resource_owner_id", Aliases: []string{"r"}},
//		},
//		Action: func(c *cli.Context) error {
//			if c.String("resource_owner_id") != "" {
//				videos, err := client.VideoList(c.String("resource_owner_id"))
//				assert(err).
//					NoneErr(func() {
//						for _, v := range videos {
//							assert(client.VideoStop(v.Id)).
//								NoneErr(func() { println("Completed ", v.Id) })
//						}
//					})
//			} else if c.Args().Len() > 0 {
//				utils.CliArgsForEach(c, func(videoID string) {
//					assert(client.VideoStop(videoID)).
//						NoneErr(func() { println("Completed ", videoID) })
//				})
//			} else {
//				ros, _ := client.GetResourceOwners()
//				for _, ro := range ros {
//					videos, err := client.VideoList(ro.Id)
//					assert(err).NoneErr(func() {
//						for _, v := range videos {
//							assert(client.VideoStop(v.Id)).
//								NoneErr(func() { println("Completed ", v.Id) })
//						}
//					})
//				}
//			}
//			return nil
//		},
//	}
//}
