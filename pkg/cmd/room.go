// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/team-telnyx/telnyx-cli/internal/apiquery"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var roomsCreate = cli.Command{
	Name:    "create",
	Usage:   "Synchronously create a Room.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[bool]{
			Name:     "enable-recording",
			Usage:    "Enable or disable recording for that room.",
			Default:  false,
			BodyPath: "enable_recording",
		},
		&requestflag.Flag[int64]{
			Name:     "max-participants",
			Usage:    "The maximum amount of participants allowed in a room. If new participants try to join after that limit is reached, their request will be rejected.",
			Default:  10,
			BodyPath: "max_participants",
		},
		&requestflag.Flag[string]{
			Name:     "unique-name",
			Usage:    "The unique (within the Telnyx account scope) name of the room.",
			BodyPath: "unique_name",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-event-failover-url",
			Usage:    "The failover URL where webhooks related to this room will be sent if sending to the primary URL fails. Must include a scheme, such as 'https'.",
			Default:  "",
			BodyPath: "webhook_event_failover_url",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-event-url",
			Usage:    "The URL where webhooks related to this room will be sent. Must include a scheme, such as 'https'.",
			BodyPath: "webhook_event_url",
		},
		&requestflag.Flag[int64]{
			Name:     "webhook-timeout-secs",
			Usage:    "Specifies how many seconds to wait before timing out a webhook.",
			BodyPath: "webhook_timeout_secs",
		},
	},
	Action:          handleRoomsCreate,
	HideHelpCommand: true,
}

var roomsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "View a room.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "room-id",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:      "include-sessions",
			Usage:     "To decide if room sessions should be included in the response.",
			QueryPath: "include_sessions",
		},
	},
	Action:          handleRoomsRetrieve,
	HideHelpCommand: true,
}

var roomsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Synchronously update a Room.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "room-id",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:     "enable-recording",
			Usage:    "Enable or disable recording for that room.",
			Default:  false,
			BodyPath: "enable_recording",
		},
		&requestflag.Flag[int64]{
			Name:     "max-participants",
			Usage:    "The maximum amount of participants allowed in a room. If new participants try to join after that limit is reached, their request will be rejected.",
			Default:  10,
			BodyPath: "max_participants",
		},
		&requestflag.Flag[string]{
			Name:     "unique-name",
			Usage:    "The unique (within the Telnyx account scope) name of the room.",
			BodyPath: "unique_name",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-event-failover-url",
			Usage:    "The failover URL where webhooks related to this room will be sent if sending to the primary URL fails. Must include a scheme, such as 'https'.",
			Default:  "",
			BodyPath: "webhook_event_failover_url",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-event-url",
			Usage:    "The URL where webhooks related to this room will be sent. Must include a scheme, such as 'https'.",
			BodyPath: "webhook_event_url",
		},
		&requestflag.Flag[int64]{
			Name:     "webhook-timeout-secs",
			Usage:    "Specifies how many seconds to wait before timing out a webhook.",
			BodyPath: "webhook_timeout_secs",
		},
	},
	Action:          handleRoomsUpdate,
	HideHelpCommand: true,
}

var roomsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "View a list of rooms.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[date_created_at][eq], filter[date_created_at][gte], filter[date_created_at][lte], filter[date_updated_at][eq], filter[date_updated_at][gte], filter[date_updated_at][lte], filter[unique_name]",
			QueryPath: "filter",
		},
		&requestflag.Flag[bool]{
			Name:      "include-sessions",
			Usage:     "To decide if room sessions should be included in the response.",
			QueryPath: "include_sessions",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page[size]",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleRoomsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.date-created-at",
			InnerField: "date_created_at",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.date-updated-at",
			InnerField: "date_updated_at",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.unique-name",
			Usage:      "Unique_name for filtering rooms.",
			InnerField: "unique_name",
		},
	},
})

var roomsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Synchronously delete a Room. Participants from that room will be kicked out,\nthey won't be able to join that room anymore, and you won't be charged anymore\nfor that room.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "room-id",
			Required: true,
		},
	},
	Action:          handleRoomsDelete,
	HideHelpCommand: true,
}

func handleRoomsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.RoomNewParams{}

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
	_, err = client.Rooms.New(ctx, params, options...)
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
		Title:          "rooms create",
		Transform:      transform,
	})
}

func handleRoomsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("room-id") && len(unusedArgs) > 0 {
		cmd.Set("room-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.RoomGetParams{}

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
	_, err = client.Rooms.Get(
		ctx,
		cmd.Value("room-id").(string),
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
		Title:          "rooms retrieve",
		Transform:      transform,
	})
}

func handleRoomsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("room-id") && len(unusedArgs) > 0 {
		cmd.Set("room-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.RoomUpdateParams{}

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
	_, err = client.Rooms.Update(
		ctx,
		cmd.Value("room-id").(string),
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
		Title:          "rooms update",
		Transform:      transform,
	})
}

func handleRoomsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.RoomListParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Rooms.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "rooms list",
			Transform:      transform,
		})
	} else {
		iter := client.Rooms.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "rooms list",
			Transform:      transform,
		})
	}
}

func handleRoomsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("room-id") && len(unusedArgs) > 0 {
		cmd.Set("room-id", unusedArgs[0])
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

	return client.Rooms.Delete(ctx, cmd.Value("room-id").(string), options...)
}
