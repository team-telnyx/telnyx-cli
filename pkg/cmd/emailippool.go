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

var emailIPPoolsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create an account-scoped pool of dedicated IPs. Set `is_default` or `domain_id`\nfor the pool to carry mail; a pool with neither carries none.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Pool name: 1-64 characters of lowercase letters, digits, `.`, `_`, and `-`, starting and ending with a letter or digit. Must be unique within the account and must not start with `dip_`.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "domain-id",
			Usage:    "Scope the pool to one of the account's email domains; the pool then carries that domain's mail. Scope at most one pool to each domain.",
			BodyPath: "domain_id",
		},
		&requestflag.Flag[bool]{
			Name:     "is-default",
			Usage:    "Make this the account's default pool, which carries mail for sending domains without a domain-scoped pool. An account can have one default pool. Defaults to `false`.",
			BodyPath: "is_default",
		},
	},
	Action:          handleEmailIPPoolsCreate,
	HideHelpCommand: true,
}

var emailIPPoolsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a single IP pool by its identifier, including its dedicated-IP members\nwith billing state.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleEmailIPPoolsRetrieve,
	HideHelpCommand: true,
}

var emailIPPoolsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Rename the pool, set or clear the default flag, or change its domain scope.\nRouting follows the new settings for subsequent mail.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[*string]{
			Name:     "domain-id",
			Usage:    "Scope the pool to one of the account's email domains, or `null` to remove the scope.",
			BodyPath: "domain_id",
		},
		&requestflag.Flag[bool]{
			Name:     "is-default",
			Usage:    "Set or clear the account default flag. An account can have one default pool.",
			BodyPath: "is_default",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Pool name: 1-64 characters of lowercase letters, digits, `.`, `_`, and `-`, starting and ending with a letter or digit. Must be unique within the account and must not start with `dip_`.",
			BodyPath: "name",
		},
	},
	Action:          handleEmailIPPoolsUpdate,
	HideHelpCommand: true,
}

var emailIPPoolsList = cli.Command{
	Name:            "list",
	Usage:           "List the account's dedicated-IP pools. Member IPs are not included (`ips` is\nempty); retrieve a pool to see its members.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleEmailIPPoolsList,
	HideHelpCommand: true,
}

var emailIPPoolsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete the pool. Member IPs are kept but leave the pool, so they carry no mail\nuntil added to another pool. Deleting a pool changes routing for subsequently\nresolved messages. It does not reroute messages already assigned to the deleted\npool. Deleting a pool that still has IPs requires `force=true` and `drain=true`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[bool]{
			Name:      "drain",
			Usage:     "Set together with `force=true` to delete a pool that still has IPs. Acknowledges deletion while the pool still has members; this flag does not reroute already queued messages.",
			Default:   false,
			QueryPath: "drain",
		},
		&requestflag.Flag[bool]{
			Name:      "force",
			Usage:     "Set together with `drain=true` to delete a pool that still has IPs. Either flag alone has no effect.",
			Default:   false,
			QueryPath: "force",
		},
	},
	Action:          handleEmailIPPoolsDelete,
	HideHelpCommand: true,
}

func handleEmailIPPoolsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := telnyx.EmailIPPoolNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Email.IPPools.New(ctx, params, options...)
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
		Title:          "email:ip-pools create",
		Transform:      transform,
	})
}

func handleEmailIPPoolsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Email.IPPools.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "email:ip-pools retrieve",
		Transform:      transform,
	})
}

func handleEmailIPPoolsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailIPPoolUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Email.IPPools.Update(
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
		Title:          "email:ip-pools update",
		Transform:      transform,
	})
}

func handleEmailIPPoolsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Email.IPPools.List(ctx, options...)
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
		Title:          "email:ip-pools list",
		Transform:      transform,
	})
}

func handleEmailIPPoolsDelete(ctx context.Context, cmd *cli.Command) error {
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
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := telnyx.EmailIPPoolDeleteParams{}

	return client.Email.IPPools.Delete(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}
