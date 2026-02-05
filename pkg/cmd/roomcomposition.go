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

var roomCompositionsCreate = cli.Command{
	Name:    "create",
	Usage:   "Asynchronously create a room composition.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "format",
			Usage:    "The desired format of the room composition.",
			Default:  "mp4",
			BodyPath: "format",
		},
		&requestflag.Flag[string]{
			Name:     "resolution",
			Usage:    "The desired resolution (width/height in pixels) of the resulting video of the room composition. Both width and height are required to be between 16 and 1280; and width * height should not exceed 1280 * 720",
			Default:  "1280x720",
			BodyPath: "resolution",
		},
		&requestflag.Flag[string]{
			Name:     "session-id",
			Usage:    "id of the room session associated with the room composition.",
			BodyPath: "session_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "video-layout",
			Usage:    "Describes the video layout of the room composition in terms of regions.",
			BodyPath: "video_layout",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-event-failover-url",
			Usage:    "The failover URL where webhooks related to this room composition will be sent if sending to the primary URL fails. Must include a scheme, such as 'https'.",
			Default:  "",
			BodyPath: "webhook_event_failover_url",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-event-url",
			Usage:    "The URL where webhooks related to this room composition will be sent. Must include a scheme, such as 'https'.",
			BodyPath: "webhook_event_url",
		},
		&requestflag.Flag[int64]{
			Name:     "webhook-timeout-secs",
			Usage:    "Specifies how many seconds to wait before timing out a webhook.",
			BodyPath: "webhook_timeout_secs",
		},
	},
	Action:          handleRoomCompositionsCreate,
	HideHelpCommand: true,
}

var roomCompositionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "View a room composition.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "room-composition-id",
			Required: true,
		},
	},
	Action:          handleRoomCompositionsRetrieve,
	HideHelpCommand: true,
}

var roomCompositionsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "View a list of room compositions.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[date_created_at][eq], filter[date_created_at][gte], filter[date_created_at][lte], filter[session_id], filter[status]",
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
	Action:          handleRoomCompositionsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.date-created-at",
			InnerField: "date_created_at",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.session-id",
			Usage:      "The session_id for filtering room compositions.",
			InnerField: "session_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.status",
			Usage:      "The status for filtering room compositions.",
			InnerField: "status",
		},
	},
})

var roomCompositionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Synchronously delete a room composition.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "room-composition-id",
			Required: true,
		},
	},
	Action:          handleRoomCompositionsDelete,
	HideHelpCommand: true,
}

func handleRoomCompositionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.RoomCompositionNewParams{}

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
	_, err = client.RoomCompositions.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "room-compositions create", obj, format, transform)
}

func handleRoomCompositionsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("room-composition-id") && len(unusedArgs) > 0 {
		cmd.Set("room-composition-id", unusedArgs[0])
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
	_, err = client.RoomCompositions.Get(ctx, cmd.Value("room-composition-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "room-compositions retrieve", obj, format, transform)
}

func handleRoomCompositionsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.RoomCompositionListParams{}

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
		_, err = client.RoomCompositions.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "room-compositions list", obj, format, transform)
	} else {
		iter := client.RoomCompositions.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "room-compositions list", iter, format, transform)
	}
}

func handleRoomCompositionsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("room-composition-id") && len(unusedArgs) > 0 {
		cmd.Set("room-composition-id", unusedArgs[0])
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

	return client.RoomCompositions.Delete(ctx, cmd.Value("room-composition-id").(string), options...)
}
