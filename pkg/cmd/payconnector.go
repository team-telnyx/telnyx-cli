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

var payConnectorsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a generic payment connector in the authenticated organization. The name\nmust be unique in that organization. Authentication supplies ownership; do not\nsend user_id or organization_id. Config is encrypted at rest; passwords and\ntokens are omitted from responses. This operation stores configuration and does\nnot initiate a payment. See the\n[processor contract](/docs/voice/programmable-voice/pay#request-sent-to-the-payment-processor).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "config",
			Usage:    "Configuration for a generic connector. A config update replaces the complete object, not individual keys. Unknown config keys are discarded. Authentication credentials must be nonblank strings when required by the selected auth_type.",
			Required: true,
			BodyPath: "config",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Unique within the authenticated organization. Used in resource paths and as connector_name in the Call Control Pay request. Cannot be renamed by update.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "mode",
			Usage:    "Defaults to test on creation. Test restricts collected card numbers but still contacts the configured processor; use a processor sandbox URL and sandbox credentials. Live removes the test-card restriction. Neither mode locks API updates.",
			Default:  "test",
			BodyPath: "mode",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Connector type covered by this reference.",
			Default:  "generic",
			BodyPath: "type",
		},
	},
	Action:          handlePayConnectorsCreate,
	HideHelpCommand: true,
}

var payConnectorsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a connector by name within the authenticated organization. Password and\ntoken are omitted. A missing name or a connector owned by another organization\nreturns 404.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "name",
			Required:  true,
			PathParam: "name",
		},
	},
	Action:          handlePayConnectorsRetrieve,
	HideHelpCommand: true,
}

var payConnectorsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a generic connector by name in the authenticated organization. Both test\nand live connectors can be updated through the API. Omitted fields remain\nunchanged. Supplying config replaces the entire config; it does not merge keys.\nInclude endpoint_url and the credentials required for the selected auth_type on\nevery config replacement. Omitting config preserves existing credentials,\nallowing a mode-only update. The path name cannot be renamed. Public updates use\nPATCH.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "name",
			Required:  true,
			PathParam: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "config",
			Usage:    "Configuration for a generic connector. A config update replaces the complete object, not individual keys. Unknown config keys are discarded. Authentication credentials must be nonblank strings when required by the selected auth_type.",
			BodyPath: "config",
		},
		&requestflag.Flag[string]{
			Name:     "mode",
			Usage:    "Omit to preserve the existing mode. Use an explicit test or live value.",
			BodyPath: "mode",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Generic connector type. Omit to preserve the existing value.",
			BodyPath: "type",
		},
	},
	Action:          handlePayConnectorsUpdate,
	HideHelpCommand: true,
}

var payConnectorsList = cli.Command{
	Name:    "list",
	Usage:   "List connectors belonging to the authenticated organization, including\nconnectors created by other users in that organization. Set filter[type]=generic\nto restrict results to the type covered here. Results default to newest\ncreated_at first. Use meta.end_cursor as page[after] for the next page, or\nmeta.start_cursor as page[before] for the preceding page. Omit both cursors for\nthe first page; if both are supplied, page[after] takes precedence. Responses\nomit passwords and tokens.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "filter-created-at-eq",
			Usage:     "Equal to this created_at timestamp.",
			QueryPath: "filter[created_at][eq]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-created-at-gt",
			Usage:     "After this created_at timestamp.",
			QueryPath: "filter[created_at][gt]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-created-at-gte",
			Usage:     "At or after this created_at timestamp.",
			QueryPath: "filter[created_at][gte]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-created-at-lt",
			Usage:     "Before this created_at timestamp.",
			QueryPath: "filter[created_at][lt]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-created-at-lte",
			Usage:     "At or before this created_at timestamp.",
			QueryPath: "filter[created_at][lte]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-mode",
			Usage:     "Exact match on mode.",
			QueryPath: "filter[mode]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-name",
			Usage:     "Exact match on name.",
			QueryPath: "filter[name]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-type",
			Usage:     "Exact match on type.",
			QueryPath: "filter[type]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-updated-at-eq",
			Usage:     "Equal to this updated_at timestamp.",
			QueryPath: "filter[updated_at][eq]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-updated-at-gt",
			Usage:     "After this updated_at timestamp.",
			QueryPath: "filter[updated_at][gt]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-updated-at-gte",
			Usage:     "At or after this updated_at timestamp.",
			QueryPath: "filter[updated_at][gte]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-updated-at-lt",
			Usage:     "Before this updated_at timestamp.",
			QueryPath: "filter[updated_at][lt]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-updated-at-lte",
			Usage:     "At or before this updated_at timestamp.",
			QueryPath: "filter[updated_at][lte]",
		},
		&requestflag.Flag[string]{
			Name:      "page-after",
			Usage:     "Opaque cursor from meta.end_cursor.",
			QueryPath: "page[after]",
		},
		&requestflag.Flag[string]{
			Name:      "page-before",
			Usage:     "Opaque cursor from meta.start_cursor.",
			QueryPath: "page[before]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Number of records per page.",
			Default:   25,
			QueryPath: "page[size]",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Comma-separated created_at or updated_at fields, optionally prefixed with - for descending order.",
			Default:   "-created_at",
			QueryPath: "sort",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handlePayConnectorsList,
	HideHelpCommand: true,
}

var payConnectorsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a connector by name in the authenticated organization. Both test and live\nconnectors can be deleted. Returns no response body on success; a missing name\nor a connector in another organization returns 404.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "name",
			Required:  true,
			PathParam: "name",
		},
	},
	Action:          handlePayConnectorsDelete,
	HideHelpCommand: true,
}

func handlePayConnectorsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.PayConnectorNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.PayConnectors.New(ctx, params, options...)
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
		Title:          "pay-connectors create",
		Transform:      transform,
	})
}

func handlePayConnectorsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("name") && len(unusedArgs) > 0 {
		cmd.Set("name", unusedArgs[0])
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
	_, err = client.PayConnectors.Get(ctx, cmd.Value("name").(string), options...)
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
		Title:          "pay-connectors retrieve",
		Transform:      transform,
	})
}

func handlePayConnectorsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("name") && len(unusedArgs) > 0 {
		cmd.Set("name", unusedArgs[0])
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

	params := telnyx.PayConnectorUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.PayConnectors.Update(
		ctx,
		cmd.Value("name").(string),
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
		Title:          "pay-connectors update",
		Transform:      transform,
	})
}

func handlePayConnectorsList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.PayConnectorListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.PayConnectors.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "pay-connectors list",
			Transform:      transform,
		})
	} else {
		iter := client.PayConnectors.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "pay-connectors list",
			Transform:      transform,
		})
	}
}

func handlePayConnectorsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("name") && len(unusedArgs) > 0 {
		cmd.Set("name", unusedArgs[0])
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

	return client.PayConnectors.Delete(ctx, cmd.Value("name").(string), options...)
}
