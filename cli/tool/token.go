package tool

import (
	"github.com/atotto/clipboard"
	. "github.com/optim-corp/cios-cli/cli"

	"github.com/urfave/cli/v2"
)

func GetTokenCommand() *cli.Command {
	return &cli.Command{
		Name:    "token",
		Aliases: []string{"t", "ac", "at"},
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "bearer", Aliases: []string{"b"}},
			&cli.BoolFlag{Name: "scope", Aliases: []string{"s"}},
			&cli.BoolFlag{Name: "clipboard", Aliases: []string{"c"}},
			&cli.BoolFlag{Name: "bearer_clipboard", Aliases: []string{"bc", "cb"}},
		},
		Usage: "cios token | ac",
		Action: func(c *cli.Context) error {
			token, scope, _, _, err := Client.Auth.GetAccessTokenByRefreshToken()
			var (
				isBearer    = c.Bool("bearer")
				isClipboard = c.Bool("clipboard")
				isScope     = c.Bool("scope")
				isBC        = c.Bool("bearer_clipboard")
			)
			if isBC {
				isBearer = true
				isClipboard = true
			}
			assert(err).Log().NoneErr(func() {
				text := is(isBearer).T("Bearer " + token).F(token).Value.(string)
				printf("\n%s\n\n", text)
				if isScope {
					println("\n" + scope + "\n")
				}
				if isClipboard {
					assert(clipboard.WriteAll(text)).Log().NoneErrPrintln("Clipped!!")
				}
			})
			return nil
		},
	}
}
