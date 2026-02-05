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

var roomsSessionsActionsEnd = cli.Command{
	Name:    "end",
	Usage:   "Note: this will also kick all participants currently present in the room",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "room-session-id",
			Required: true,
		},
	},
	Action:          handleRoomsSessionsActionsEnd,
	HideHelpCommand: true,
}

var roomsSessionsActionsKick = cli.Command{
	Name:    "kick",
	Usage:   "Kick participants from a room session.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "room-session-id",
			Required: true,
		},
		&requestflag.Flag[[]string]{
			Name:     "exclude",
			Usage:    "List of participant id to exclude from the action.",
			BodyPath: "exclude",
		},
		&requestflag.Flag[any]{
			Name:     "participants",
			Usage:    `Either a list of participant id to perform the action on, or the keyword "all" to perform the action on all participant.`,
			BodyPath: "participants",
		},
	},
	Action:          handleRoomsSessionsActionsKick,
	HideHelpCommand: true,
}

var roomsSessionsActionsMute = cli.Command{
	Name:    "mute",
	Usage:   "Mute participants in room session.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "room-session-id",
			Required: true,
		},
		&requestflag.Flag[[]string]{
			Name:     "exclude",
			Usage:    "List of participant id to exclude from the action.",
			BodyPath: "exclude",
		},
		&requestflag.Flag[any]{
			Name:     "participants",
			Usage:    `Either a list of participant id to perform the action on, or the keyword "all" to perform the action on all participant.`,
			BodyPath: "participants",
		},
	},
	Action:          handleRoomsSessionsActionsMute,
	HideHelpCommand: true,
}

var roomsSessionsActionsUnmute = cli.Command{
	Name:    "unmute",
	Usage:   "Unmute participants in room session.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "room-session-id",
			Required: true,
		},
		&requestflag.Flag[[]string]{
			Name:     "exclude",
			Usage:    "List of participant id to exclude from the action.",
			BodyPath: "exclude",
		},
		&requestflag.Flag[any]{
			Name:     "participants",
			Usage:    `Either a list of participant id to perform the action on, or the keyword "all" to perform the action on all participant.`,
			BodyPath: "participants",
		},
	},
	Action:          handleRoomsSessionsActionsUnmute,
	HideHelpCommand: true,
}

func handleRoomsSessionsActionsEnd(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("room-session-id") && len(unusedArgs) > 0 {
		cmd.Set("room-session-id", unusedArgs[0])
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
	_, err = client.Rooms.Sessions.Actions.End(ctx, cmd.Value("room-session-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rooms:sessions:actions end", obj, format, transform)
}

func handleRoomsSessionsActionsKick(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("room-session-id") && len(unusedArgs) > 0 {
		cmd.Set("room-session-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.RoomSessionActionKickParams{}

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
	_, err = client.Rooms.Sessions.Actions.Kick(
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
	return ShowJSON(os.Stdout, "rooms:sessions:actions kick", obj, format, transform)
}

func handleRoomsSessionsActionsMute(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("room-session-id") && len(unusedArgs) > 0 {
		cmd.Set("room-session-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.RoomSessionActionMuteParams{}

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
	_, err = client.Rooms.Sessions.Actions.Mute(
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
	return ShowJSON(os.Stdout, "rooms:sessions:actions mute", obj, format, transform)
}

func handleRoomsSessionsActionsUnmute(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("room-session-id") && len(unusedArgs) > 0 {
		cmd.Set("room-session-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.RoomSessionActionUnmuteParams{}

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
	_, err = client.Rooms.Sessions.Actions.Unmute(
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
	return ShowJSON(os.Stdout, "rooms:sessions:actions unmute", obj, format, transform)
}
