// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/telnyx-cli/internal/apiquery"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var mobileVoiceConnectionsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a Mobile Voice Connection",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[bool]{
			Name:     "active",
			Default:  true,
			BodyPath: "active",
		},
		&requestflag.Flag[string]{
			Name:     "connection-name",
			Default:  "Telnyx Mobile Voice IMS",
			BodyPath: "connection_name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "inbound",
			BodyPath: "inbound",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "outbound",
			BodyPath: "outbound",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-api-version",
			Default:  "2",
			BodyPath: "webhook_api_version",
		},
		&requestflag.Flag[any]{
			Name:     "webhook-event-failover-url",
			BodyPath: "webhook_event_failover_url",
		},
		&requestflag.Flag[any]{
			Name:     "webhook-event-url",
			BodyPath: "webhook_event_url",
		},
		&requestflag.Flag[any]{
			Name:     "webhook-timeout-secs",
			BodyPath: "webhook_timeout_secs",
		},
	},
	Action:          handleMobileVoiceConnectionsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"inbound": {
		&requestflag.InnerFlag[int64]{
			Name:       "inbound.channel-limit",
			InnerField: "channel_limit",
		},
	},
	"outbound": {
		&requestflag.InnerFlag[int64]{
			Name:       "outbound.channel-limit",
			InnerField: "channel_limit",
		},
		&requestflag.InnerFlag[string]{
			Name:       "outbound.outbound-voice-profile-id",
			InnerField: "outbound_voice_profile_id",
		},
	},
})

var mobileVoiceConnectionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a Mobile Voice Connection",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleMobileVoiceConnectionsRetrieve,
	HideHelpCommand: true,
}

var mobileVoiceConnectionsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update a Mobile Voice Connection",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:     "active",
			BodyPath: "active",
		},
		&requestflag.Flag[string]{
			Name:     "connection-name",
			BodyPath: "connection_name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "inbound",
			BodyPath: "inbound",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "outbound",
			BodyPath: "outbound",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-api-version",
			BodyPath: "webhook_api_version",
		},
		&requestflag.Flag[any]{
			Name:     "webhook-event-failover-url",
			BodyPath: "webhook_event_failover_url",
		},
		&requestflag.Flag[any]{
			Name:     "webhook-event-url",
			BodyPath: "webhook_event_url",
		},
		&requestflag.Flag[int64]{
			Name:     "webhook-timeout-secs",
			BodyPath: "webhook_timeout_secs",
		},
	},
	Action:          handleMobileVoiceConnectionsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"inbound": {
		&requestflag.InnerFlag[int64]{
			Name:       "inbound.channel-limit",
			InnerField: "channel_limit",
		},
	},
	"outbound": {
		&requestflag.InnerFlag[int64]{
			Name:       "outbound.channel-limit",
			InnerField: "channel_limit",
		},
		&requestflag.InnerFlag[string]{
			Name:       "outbound.outbound-voice-profile-id",
			InnerField: "outbound_voice_profile_id",
		},
	},
})

var mobileVoiceConnectionsList = cli.Command{
	Name:    "list",
	Usage:   "List Mobile Voice Connections",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "filter-connection-name-contains",
			Usage:     "Filter by connection name containing the given string",
			QueryPath: "filter[connection_name][contains]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "The page number to load",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "The size of the page",
			QueryPath: "page[size]",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Sort by field (e.g., created_at, connection_name, active). Prefix with - for descending order.",
			QueryPath: "sort",
		},
	},
	Action:          handleMobileVoiceConnectionsList,
	HideHelpCommand: true,
}

var mobileVoiceConnectionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a Mobile Voice Connection",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleMobileVoiceConnectionsDelete,
	HideHelpCommand: true,
}

func handleMobileVoiceConnectionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MobileVoiceConnectionNewParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.MobileVoiceConnections.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "mobile-voice-connections create", obj, format, transform)
}

func handleMobileVoiceConnectionsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.MobileVoiceConnections.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "mobile-voice-connections retrieve", obj, format, transform)
}

func handleMobileVoiceConnectionsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MobileVoiceConnectionUpdateParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.MobileVoiceConnections.Update(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "mobile-voice-connections update", obj, format, transform)
}

func handleMobileVoiceConnectionsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MobileVoiceConnectionListParams{}

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

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.MobileVoiceConnections.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "mobile-voice-connections list", obj, format, transform)
	} else {
		iter := client.MobileVoiceConnections.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "mobile-voice-connections list", iter, format, transform)
	}
}

func handleMobileVoiceConnectionsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.MobileVoiceConnections.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "mobile-voice-connections delete", obj, format, transform)
}
