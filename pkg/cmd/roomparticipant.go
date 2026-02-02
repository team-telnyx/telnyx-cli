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

var roomParticipantsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "View a room participant.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "room-participant-id",
			Required: true,
		},
	},
	Action:          handleRoomParticipantsRetrieve,
	HideHelpCommand: true,
}

var roomParticipantsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "View a list of room participants.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[date_joined_at][eq], filter[date_joined_at][gte], filter[date_joined_at][lte], filter[date_updated_at][eq], filter[date_updated_at][gte], filter[date_updated_at][lte], filter[date_left_at][eq], filter[date_left_at][gte], filter[date_left_at][lte], filter[context], filter[session_id]",
			QueryPath: "filter",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "page",
			Usage:     "Consolidated page parameter (deepObject style). Originally: page[size], page[number]",
			QueryPath: "page",
		},
	},
	Action:          handleRoomParticipantsList,
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
		&requestflag.InnerFlag[string]{
			Name:       "filter.session-id",
			Usage:      "Session_id for filtering room participants.",
			InnerField: "session_id",
		},
	},
	"page": {
		&requestflag.InnerFlag[int64]{
			Name:       "page.number",
			Usage:      "The page number to load.",
			InnerField: "number",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "page.size",
			Usage:      "The size of the page.",
			InnerField: "size",
		},
	},
})

func handleRoomParticipantsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("room-participant-id") && len(unusedArgs) > 0 {
		cmd.Set("room-participant-id", unusedArgs[0])
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
	_, err = client.RoomParticipants.Get(ctx, cmd.Value("room-participant-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "room-participants retrieve", obj, format, transform)
}

func handleRoomParticipantsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.RoomParticipantListParams{}

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
		_, err = client.RoomParticipants.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "room-participants list", obj, format, transform)
	} else {
		iter := client.RoomParticipants.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "room-participants list", iter, format, transform)
	}
}
