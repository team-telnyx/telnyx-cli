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

var roomRecordingsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "View a room recording.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "room-recording-id",
			Required: true,
		},
	},
	Action:          handleRoomRecordingsRetrieve,
	HideHelpCommand: true,
}

var roomRecordingsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "View a list of room recordings.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[date_ended_at][eq], filter[date_ended_at][gte], filter[date_ended_at][lte], filter[date_started_at][eq], filter[date_started_at][gte], filter[date_started_at][lte], filter[room_id], filter[participant_id], filter[session_id], filter[status], filter[type], filter[duration_secs]",
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
	Action:          handleRoomRecordingsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.date-ended-at",
			InnerField: "date_ended_at",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.date-started-at",
			InnerField: "date_started_at",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "filter.duration-secs",
			Usage:      "duration_secs greater or equal for filtering room recordings.",
			InnerField: "duration_secs",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.participant-id",
			Usage:      "participant_id for filtering room recordings.",
			InnerField: "participant_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.room-id",
			Usage:      "room_id for filtering room recordings.",
			InnerField: "room_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.session-id",
			Usage:      "session_id for filtering room recordings.",
			InnerField: "session_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.status",
			Usage:      "status for filtering room recordings.",
			InnerField: "status",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.type",
			Usage:      "type for filtering room recordings.",
			InnerField: "type",
		},
	},
})

var roomRecordingsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Synchronously delete a Room Recording.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "room-recording-id",
			Required: true,
		},
	},
	Action:          handleRoomRecordingsDelete,
	HideHelpCommand: true,
}

var roomRecordingsDeleteBulk = requestflag.WithInnerFlags(cli.Command{
	Name:    "delete-bulk",
	Usage:   "Delete several room recordings in a bulk.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[date_ended_at][eq], filter[date_ended_at][gte], filter[date_ended_at][lte], filter[date_started_at][eq], filter[date_started_at][gte], filter[date_started_at][lte], filter[room_id], filter[participant_id], filter[session_id], filter[status], filter[type], filter[duration_secs]",
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
	Action:          handleRoomRecordingsDeleteBulk,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.date-ended-at",
			InnerField: "date_ended_at",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.date-started-at",
			InnerField: "date_started_at",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "filter.duration-secs",
			Usage:      "duration_secs greater or equal for filtering room recordings.",
			InnerField: "duration_secs",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.participant-id",
			Usage:      "participant_id for filtering room recordings.",
			InnerField: "participant_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.room-id",
			Usage:      "room_id for filtering room recordings.",
			InnerField: "room_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.session-id",
			Usage:      "session_id for filtering room recordings.",
			InnerField: "session_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.status",
			Usage:      "status for filtering room recordings.",
			InnerField: "status",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.type",
			Usage:      "type for filtering room recordings.",
			InnerField: "type",
		},
	},
})

func handleRoomRecordingsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("room-recording-id") && len(unusedArgs) > 0 {
		cmd.Set("room-recording-id", unusedArgs[0])
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
	_, err = client.RoomRecordings.Get(ctx, cmd.Value("room-recording-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "room-recordings retrieve", obj, format, transform)
}

func handleRoomRecordingsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.RoomRecordingListParams{}

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
		_, err = client.RoomRecordings.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "room-recordings list", obj, format, transform)
	} else {
		iter := client.RoomRecordings.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "room-recordings list", iter, format, transform)
	}
}

func handleRoomRecordingsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("room-recording-id") && len(unusedArgs) > 0 {
		cmd.Set("room-recording-id", unusedArgs[0])
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

	return client.RoomRecordings.Delete(ctx, cmd.Value("room-recording-id").(string), options...)
}

func handleRoomRecordingsDeleteBulk(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.RoomRecordingDeleteBulkParams{}

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
	_, err = client.RoomRecordings.DeleteBulk(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "room-recordings delete-bulk", obj, format, transform)
}
