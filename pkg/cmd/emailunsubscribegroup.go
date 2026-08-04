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

var emailUnsubscribeGroupsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create an unsubscribe group",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "description",
			BodyPath: "description",
		},
	},
	Action:          handleEmailUnsubscribeGroupsCreate,
	HideHelpCommand: true,
}

var emailUnsubscribeGroupsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve an unsubscribe group",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleEmailUnsubscribeGroupsRetrieve,
	HideHelpCommand: true,
}

var emailUnsubscribeGroupsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Partial update (only `name` / `description`). `PUT` is not routed.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[*string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
	},
	Action:          handleEmailUnsubscribeGroupsUpdate,
	HideHelpCommand: true,
}

var emailUnsubscribeGroupsList = cli.Command{
	Name:    "list",
	Usage:   "Offset pagination only (`page[number]` default 1, `page[size]` default 25, max\n100). No `sort`/`filter`/cursor — ordering fixed `desc created_at, desc id`.\nUses the shared `QueryParser.parse_offset/1` — a malformed `page` (e.g. flat\n`?page=1` instead of `?page[number]=1`) returns `400` (code `10015`), consistent\nwith `GET /v2/email_blocks`. `meta` includes `total_pages`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "Offset page number (≥1, default 1).",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Page size (1–100, default 25).",
			Default:   25,
			QueryPath: "page[size]",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleEmailUnsubscribeGroupsList,
	HideHelpCommand: true,
}

var emailUnsubscribeGroupsDelete = cli.Command{
	Name:    "delete",
	Usage:   "If the group has 0 active suppressions, hard-deletes the row. With `force=true`,\nsoft-deletes all active suppressions first (status → `removed`, `group_id`\ncleared, `removed` audit event per block) in a single transaction, then\nhard-deletes the group. Without `force` and active suppressions present → `409`.\nAudit trail is preserved. `force` only accepts the string `\"true\"` or boolean\n`true`; all other values are false.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[any]{
			Name:      "force",
			Usage:     "Force-delete a group with active suppressions. Only `\"true\"` (string) or `true` (bool) are truthy; all other values are false.",
			QueryPath: "force",
		},
	},
	Action:          handleEmailUnsubscribeGroupsDelete,
	HideHelpCommand: true,
}

func handleEmailUnsubscribeGroupsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailUnsubscribeGroupNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailUnsubscribeGroups.New(ctx, params, options...)
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
		Title:          "email-unsubscribe-groups create",
		Transform:      transform,
	})
}

func handleEmailUnsubscribeGroupsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.EmailUnsubscribeGroups.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "email-unsubscribe-groups retrieve",
		Transform:      transform,
	})
}

func handleEmailUnsubscribeGroupsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailUnsubscribeGroupUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailUnsubscribeGroups.Update(
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
		Title:          "email-unsubscribe-groups update",
		Transform:      transform,
	})
}

func handleEmailUnsubscribeGroupsList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailUnsubscribeGroupListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.EmailUnsubscribeGroups.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "email-unsubscribe-groups list",
			Transform:      transform,
		})
	} else {
		iter := client.EmailUnsubscribeGroups.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "email-unsubscribe-groups list",
			Transform:      transform,
		})
	}
}

func handleEmailUnsubscribeGroupsDelete(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailUnsubscribeGroupDeleteParams{}

	return client.EmailUnsubscribeGroups.Delete(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}
