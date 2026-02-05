// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/team-telnyx/telnyx-cli/internal/apiquery"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var roomsSessionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "View a room session.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "room-session-id",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:      "include-participants",
			Usage:     "To decide if room participants should be included in the response.",
			QueryPath: "include_participants",
		},
	},
	Action:          handleRoomsSessionsRetrieve,
	HideHelpCommand: true,
}

var roomsSessionsList0 = requestflag.WithInnerFlags(cli.Command{
	Name:    "list-0",
	Usage:   "View a list of room sessions.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[date_created_at][eq], filter[date_created_at][gte], filter[date_created_at][lte], filter[date_updated_at][eq], filter[date_updated_at][gte], filter[date_updated_at][lte], filter[date_ended_at][eq], filter[date_ended_at][gte], filter[date_ended_at][lte], filter[room_id], filter[active]",
			QueryPath: "filter",
		},
		&requestflag.Flag[bool]{
			Name:      "include-participants",
			Usage:     "To decide if room participants should be included in the response.",
			QueryPath: "include_participants",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page[size]",
		},
	},
	Action:          handleRoomsSessionsList0,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[bool]{
			Name:       "filter.active",
			Usage:      "Filter active or inactive room sessions.",
			InnerField: "active",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.date-created-at",
			InnerField: "date_created_at",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.date-ended-at",
			InnerField: "date_ended_at",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.date-updated-at",
			InnerField: "date_updated_at",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.room-id",
			Usage:      "Room_id for filtering room sessions.",
			InnerField: "room_id",
		},
	},
})

var roomsSessionsList1 = requestflag.WithInnerFlags(cli.Command{
	Name:    "list-1",
	Usage:   "View a list of room sessions.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "room-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[date_created_at][eq], filter[date_created_at][gte], filter[date_created_at][lte], filter[date_updated_at][eq], filter[date_updated_at][gte], filter[date_updated_at][lte], filter[date_ended_at][eq], filter[date_ended_at][gte], filter[date_ended_at][lte], filter[active]",
			QueryPath: "filter",
		},
		&requestflag.Flag[bool]{
			Name:      "include-participants",
			Usage:     "To decide if room participants should be included in the response.",
			QueryPath: "include_participants",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page[size]",
		},
	},
	Action:          handleRoomsSessionsList1,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[bool]{
			Name:       "filter.active",
			Usage:      "Filter active or inactive room sessions.",
			InnerField: "active",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.date-created-at",
			InnerField: "date_created_at",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.date-ended-at",
			InnerField: "date_ended_at",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.date-updated-at",
			InnerField: "date_updated_at",
		},
	},
})

var roomsSessionsRetrieveParticipants = requestflag.WithInnerFlags(cli.Command{
	Name:    "retrieve-participants",
	Usage:   "View a list of room participants.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "room-session-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[date_joined_at][eq], filter[date_joined_at][gte], filter[date_joined_at][lte], filter[date_updated_at][eq], filter[date_updated_at][gte], filter[date_updated_at][lte], filter[date_left_at][eq], filter[date_left_at][gte], filter[date_left_at][lte], filter[context]",
			QueryPath: "filter",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page[size]",
		},
	},
	Action:          handleRoomsSessionsRetrieveParticipants,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.context",
			Usage:      "Filter room participants based on the context.",
			InnerField: "context",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.date-joined-at",
			InnerField: "date_joined_at",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.date-left-at",
			InnerField: "date_left_at",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.date-updated-at",
			InnerField: "date_updated_at",
		},
	},
})

func handleRoomsSessionsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("room-session-id") && len(unusedArgs) > 0 {
		cmd.Set("room-session-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.RoomSessionGetParams{}

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
	_, err = client.Rooms.Sessions.Get(
		ctx,
		cmd.Value("room-session-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rooms:sessions retrieve", obj, format, transform)
}

func handleRoomsSessionsList0(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.RoomSessionList0Params{}

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
		_, err = client.Rooms.Sessions.List0(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "rooms:sessions list-0", obj, format, transform)
	} else {
		iter := client.Rooms.Sessions.List0AutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "rooms:sessions list-0", iter, format, transform)
	}
}

func handleRoomsSessionsList1(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("room-id") && len(unusedArgs) > 0 {
		cmd.Set("room-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.RoomSessionList1Params{}

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
		_, err = client.Rooms.Sessions.List1(
			ctx,
			cmd.Value("room-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "rooms:sessions list-1", obj, format, transform)
	} else {
		iter := client.Rooms.Sessions.List1AutoPaging(
			ctx,
			cmd.Value("room-id").(string),
			params,
			options...,
		)
		return ShowJSONIterator(os.Stdout, "rooms:sessions list-1", iter, format, transform)
	}
}

func handleRoomsSessionsRetrieveParticipants(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("room-session-id") && len(unusedArgs) > 0 {
		cmd.Set("room-session-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.RoomSessionGetParticipantsParams{}

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
		_, err = client.Rooms.Sessions.GetParticipants(
			ctx,
			cmd.Value("room-session-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "rooms:sessions retrieve-participants", obj, format, transform)
	} else {
		iter := client.Rooms.Sessions.GetParticipantsAutoPaging(
			ctx,
			cmd.Value("room-session-id").(string),
			params,
			options...,
		)
		return ShowJSONIterator(os.Stdout, "rooms:sessions retrieve-participants", iter, format, transform)
	}
}
