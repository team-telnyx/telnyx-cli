// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/team-telnyx/telnyx-cli/internal/apiquery"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var emailIPPoolsIPsCreate = cli.Command{
	Name:    "create",
	Usage:   "Add a dedicated IP to the pool with an optional routing weight. An IP can belong\nto one pool; to move an IP that is already in a pool, use\n`PATCH /email/dedicated_ips/{id}` with `pool_id`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "dedicated-ip-id",
			Usage:    "The dedicated IP to add. It must belong to the account, must not be released, and must not already be in a pool.",
			Required: true,
			BodyPath: "dedicated_ip_id",
		},
		&requestflag.Flag[int64]{
			Name:     "weight",
			Usage:    "Relative routing weight for this IP in the pool. Positive integer; defaults to 100.",
			BodyPath: "weight",
		},
	},
	Action:          handleEmailIPPoolsIPsCreate,
	HideHelpCommand: true,
}

var emailIPPoolsIPsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Remove the dedicated IP from the pool. The IP is kept but carries no mail until\nadded to a pool again. Removing the last IP requires `force=true` and\n`drain=true`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "ip-id",
			Required:  true,
			PathParam: "ip_id",
		},
		&requestflag.Flag[bool]{
			Name:      "drain",
			Usage:     "Set together with `force=true` to remove the last IP in the pool. Acknowledges that mail routed to the emptied pool is held until the pool has a sendable IP again.",
			Default:   false,
			QueryPath: "drain",
		},
		&requestflag.Flag[bool]{
			Name:      "force",
			Usage:     "Set together with `drain=true` to remove the last IP in the pool. Either flag alone has no effect.",
			Default:   false,
			QueryPath: "force",
		},
	},
	Action:          handleEmailIPPoolsIPsDelete,
	HideHelpCommand: true,
}

func handleEmailIPPoolsIPsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := telnyx.EmailIPPoolIPNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Email.IPPools.IPs.New(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "email:ip-pools:ips create",
		Transform:      transform,
	})
}

func handleEmailIPPoolsIPsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("ip-id") && len(unusedArgs) > 0 {
		cmd.Set("ip-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := telnyx.EmailIPPoolIPDeleteParams{
		ID: cmd.Value("id").(string),
	}

	return client.Email.IPPools.IPs.Delete(
		ctx,
		cmd.Value("ip-id").(string),
		params,
		options...,
	)
}
