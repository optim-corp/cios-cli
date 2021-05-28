package group

import (
	"strings"
	"unicode/utf8"

	ciosctx "github.com/optim-corp/cios-golang-sdk/ctx"

	"github.com/optim-corp/cios-cli/utils/console"

	cnv "github.com/fcfcqloow/go-advance/convert"

	"github.com/optim-corp/cios-golang-sdk/cios"

	ciossdk "github.com/optim-corp/cios-golang-sdk/sdk"

	"github.com/urfave/cli/v2"

	"github.com/AlecAivazis/survey/v2"
	. "github.com/optim-corp/cios-cli/cli"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
)

var (
	is          = utils.Is
	listUtility = console.ListUtility
	spaceRight  = console.SpaceRight
	fPrintln    = console.Fprintln
	fPrintf     = console.Fprintf
	assert      = utils.EAssert
)

func GetGroupCommand() *cli.Command {
	return &cli.Command{
		Name:    "group",
		Aliases: []string{"gp"},
		Usage:   "cios group | gp",
		Subcommands: []*cli.Command{
			createGroup(),
			deleteGroup(),
			listGroup(),
			inviteGroup(),
			updateGroup(),
		},
	}
}
func createGroup() *cli.Command {
	return &cli.Command{
		Name:    models.CREATE,
		Aliases: models.ALIAS_CREATE,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
			&cli.StringFlag{Name: "parent_group_id", Aliases: []string{"p"}},
			&cli.StringFlag{Name: "tag", Aliases: []string{"t"}},
		},
		Action: func(c *cli.Context) error {
			answers := struct {
				Name          string
				ParentGroupID string
				Tags          string
			}{}
			if c.String("name") == "" {
				console.Question(
					[]*survey.Question{
						{
							Name:   "name",
							Prompt: &survey.Input{Message: "name: "},
						},
						{
							Name:   "parentGroupID",
							Prompt: &survey.Input{Message: "parent group id: "},
						},
						{
							Name:   "tags",
							Prompt: &survey.Input{Message: "tags(tag1,tag2,tag3): "},
						},
					}, &answers,
				)
			} else {
				answers = struct {
					Name          string
					ParentGroupID string
					Tags          string
				}{
					Name:          c.String("name"),
					ParentGroupID: c.String("parent_group_id"),
					Tags:          c.String("tag"),
				}
			}

			tagExp := func(exp bool) []string {
				if exp {
					return strings.Split(answers.Tags, ",")
				}
				return []string{}
			}
			var opts cios.GroupCreateRequest
			if answers.ParentGroupID == "" {
				tags := tagExp(answers.Tags != "")
				opts = cios.GroupCreateRequest{
					Name: answers.Name,
					Tags: &tags,
					Type: "Group",
				}
			} else {
				tags := tagExp(answers.Tags != "")
				opts = cios.GroupCreateRequest{
					Name:          answers.Name,
					ParentGroupId: &answers.ParentGroupID,
					Tags:          &tags,
					Type:          "Group",
				}
			}
			_, _, err := Client.Account.CreateGroup(ciosctx.Background(), opts)
			assert(err).Log().NoneErrPrintln("Completed " + answers.Name)
			return nil
		},
	}
}
func updateGroup() *cli.Command {
	return &cli.Command{
		Name:    models.PATCH,
		Aliases: models.ALIAS_PATCH,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
			&cli.StringFlag{Name: "tag", Aliases: []string{"t"}},
		},
		Action: func(c *cli.Context) error {
			if c.Args().Len() == 0 {
				println("No Group ID")
				return nil
			}
			answers := struct {
				Name string
				Tags string
			}{
				Name: c.String("name"),
				Tags: c.String("tag"),
			}
			if answers.Name == "" {
				tags := is(answers.Tags != "").T(strings.Split(answers.Tags, ",")).F([]string{}).Value.([]string)
				_, _, err := Client.Account.UpdateGroup(ciosctx.Background(), c.Args().Get(0), cios.GroupUpdateRequest{Tags: &tags})
				assert(err).Log().NoneErrPrintln("Completed " + c.Args().Get(0))
			} else {
				tags := is(answers.Tags != "").T(strings.Split(answers.Tags, ",")).F([]string{}).Value.([]string)
				_, _, err := Client.Account.UpdateGroup(ciosctx.Background(), c.Args().Get(0), cios.GroupUpdateRequest{Name: &answers.Name, Tags: &tags})
				assert(err).Log().NoneErrPrintln("Completed " + c.Args().Get(0))
			}

			return nil
		},
	}
}
func deleteGroup() *cli.Command {
	return &cli.Command{
		Name:    models.DELETE,
		Aliases: models.ALIAS_DELETE,
		Action: func(c *cli.Context) error {
			console.CliArgsForEach(c, func(id string) {
				_, err := Client.Account.DeleteGroup(ciosctx.Background(), id)
				assert(err).Log().NoneErrPrintln("Completed " + id)
			})
			return nil
		},
	}
}
func listGroup() *cli.Command {
	return &cli.Command{
		Name:    models.LIST,
		Aliases: models.ALIAS_LIST,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
			&cli.StringFlag{Name: "label", Aliases: []string{"lb"}},
			&cli.StringFlag{Name: "tag", Aliases: []string{"t"}},
			&cli.Int64Flag{Name: "limit", Aliases: []string{"l"}, Value: 30},
			&cli.Int64Flag{Name: "offset", Aliases: []string{"o"}},
		},
		Action: func(c *cli.Context) error {
			var (
				name   = c.String("name")
				label  = c.String("label")
				limit  = c.Int64("limit")
				offset = c.Int64("offset")
				tag    = c.String("tag")
			)
			groups, _, _ := Client.Account.GetGroupsAll(ciosctx.Background(), ciossdk.MakeGetGroupsOpts().
				Limit(limit).
				Name(name).
				Label(label).
				Tags(tag).
				Offset(offset))
			listUtility(func() {
				length := utf8.RuneCountInString("0000000000000-0000-0000-000000000000")
				fPrintln("\t\t|id|\t\t\t\t|parent_group_id|\t\t\t|resource_owner_id|\t\t|type|         |name / tags|")
				for _, value := range groups {
					resource, _, err := Client.Account.GetResourceOwnerByGroupId(ciosctx.Background(), value.Id)
					assert(err).Log().NoneErr(func() {
						fPrintf(
							"%s\t%s\t%s\t%s　　%s / %s\n",
							value.Id,
							spaceRight(cnv.MustStr(value.ParentGroupId), length),
							resource.Id,
							spaceRight(value.Type, utf8.RuneCountInString("Corporation")),
							value.Name,
							value.Tags,
						)
					})
				}
			})
			return nil
		},
	}
}
func inviteGroup() *cli.Command {
	return &cli.Command{
		Name:    "invite",
		Aliases: []string{"call", "come"},
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "list", Aliases: []string{"l"}},
		},
		Action: func(c *cli.Context) error {
			if c.Args().Len() == 0 {
				println("no groups")
				return nil
			}
			emails := []string{}
			if c.Bool("list") {

			} else {
				for {
					answers := struct {
						Email string
					}{}
					console.Question([]*survey.Question{
						{
							Name:   "email",
							Prompt: &survey.Input{Message: "email or exit: "},
						},
					}, &answers)
					if answers.Email == "exit" {
						break
					} else {
						emails = append(emails, answers.Email)
					}
				}
				console.CliArgsForEach(c, func(id string) {
					for _, email := range emails {
						_, _, err := Client.Account.InviteGroup(ciosctx.Background(), id, email)
						assert(err).Log().NoneErrPrintln("Completed ", id, "\n", email)
					}
				})
			}
			return nil
		},
	}
}
