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

var calendarConnectionsMeetingsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Set an individual meeting override: true opts in, false opts out, null follows\nthe connection policy. A supported meeting link is required. Changes re-evaluate\nthe meeting immediately; poll the linked session for its final lifecycle\noutcome.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "meeting-id",
			Required:  true,
			PathParam: "meeting_id",
		},
		&requestflag.Flag[*bool]{
			Name:     "bot-will-join",
			Usage:    "True opts in, false opts out, null follows connection policy. A supported meeting URL is still required.",
			Required: true,
			BodyPath: "bot_will_join",
		},
	},
	Action:          handleCalendarConnectionsMeetingsUpdate,
	HideHelpCommand: true,
}

var calendarConnectionsMeetingsList = cli.Command{
	Name:    "list",
	Usage:   "List tracked meetings in a window. Defaults from now through seven days from\nnow. from must not exceed to and the window must not exceed 31 days.\nbot_will_join is the effective policy decision; check session_id and the linked\nsession status for actual attendance.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[any]{
			Name:      "from",
			Usage:     "Window start; defaults to now.",
			QueryPath: "from",
		},
		&requestflag.Flag[any]{
			Name:      "to",
			Usage:     "Window end; defaults to seven days from now.",
			QueryPath: "to",
		},
	},
	Action:          handleCalendarConnectionsMeetingsList,
	HideHelpCommand: true,
}

func handleCalendarConnectionsMeetingsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("meeting-id") && len(unusedArgs) > 0 {
		cmd.Set("meeting-id", unusedArgs[0])
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

	params := telnyx.CalendarConnectionMeetingUpdateParams{
		ID: cmd.Value("id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.CalendarConnections.Meetings.Update(
		ctx,
		cmd.Value("meeting-id").(string),
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
		Title:          "calendar-connections:meetings update",
		Transform:      transform,
	})
}

func handleCalendarConnectionsMeetingsList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.CalendarConnectionMeetingListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.CalendarConnections.Meetings.List(
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
		Title:          "calendar-connections:meetings list",
		Transform:      transform,
	})
}
