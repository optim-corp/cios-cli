package device

import (
	"context"

	"github.com/optim-corp/cios-golang-sdk/cios"

	. "github.com/optim-corp/cios-cli/cli"
	"github.com/optim-corp/cios-cli/models"
	"github.com/optim-corp/cios-cli/utils"
	"github.com/urfave/cli/v2"
)

func GetDevicePolicyCommand() *cli.Command {
	return &cli.Command{
		Name:    "policy",
		Aliases: []string{"p"},
		Usage:   "cios policy | p | device policy | d p",
		Subcommands: []*cli.Command{
			createDevicePolicy(),
			deleteDevicePolicy(),
			listDevicePolicy(),
		},
	}
}

func createDevicePolicy() *cli.Command {
	return &cli.Command{
		Name:    models.CREATE,
		Aliases: []string{"add", "cre"},
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "resourceOwnerId", Required: true},
		},
		Action: func(c *cli.Context) error {
			policy, _, err := Client.DeviceManagement.CreatePolicy(c.String("resourceOwnerId"), context.Background())
			assert(err).Log().NoneErr(func() { utils.OutStructJson(policy) })
			return nil
		},
	}
}
func listDevicePolicy() *cli.Command {
	return &cli.Command{
		Name:      models.LIST,
		UsageText: "cios device policy ls",
		Aliases:   models.ALIAS_LIST,
		Action: func(c *cli.Context) error {
			policies, _, err := Client.DeviceManagement.GetPolicies(cios.ApiGetDevicePoliciesRequest{}, context.Background())
			assert(err).Log().NoneErr(func() {
				listUtility(func() {
					fPrintln("\t\t|id|\t\t\t\t|resource owner|\t      |create at|\t|updated at|")
					for _, policy := range policies.Policies {
						fPrintln(policy.PolicyId + " " +
							policy.ResourceOwnerId + " " +
							policy.CreatedAt + " " +
							policy.UpdatedAt)
					}
				})
			})
			return nil
		},
	}
}

func deleteDevicePolicy() *cli.Command {
	return &cli.Command{
		Name:      models.DELETE,
		Aliases:   models.ALIAS_DELETE,
		UsageText: "cios device monitoring del [command options] [policy_id...]",
		Action: func(c *cli.Context) error {
			utils.CliArgsForEach(c, func(t string) {
				_, err := Client.DeviceManagement.DeletePolicy(t, context.Background())
				assert(err).Log()
			})
			return nil
		},
	}
}
