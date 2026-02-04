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

var roomsActionsGenerateJoinClientToken = cli.Command{
	Name:    "generate-join-client-token",
	Usage:   "Synchronously create an Client Token to join a Room. Client Token is necessary\nto join a Telnyx Room. Client Token will expire after `token_ttl_secs`, a\nRefresh Token is also provided to refresh a Client Token, the Refresh Token\nexpires after `refresh_token_ttl_secs`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "room-id",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:     "refresh-token-ttl-secs",
			Usage:    "The time to live in seconds of the Refresh Token, after that time the Refresh Token is invalid and can't be used to refresh Client Token.",
			Default:  3600,
			BodyPath: "refresh_token_ttl_secs",
		},
		&requestflag.Flag[int64]{
			Name:     "token-ttl-secs",
			Usage:    "The time to live in seconds of the Client Token, after that time the Client Token is invalid and can't be used to join a Room.",
			Default:  600,
			BodyPath: "token_ttl_secs",
		},
	},
	Action:          handleRoomsActionsGenerateJoinClientToken,
	HideHelpCommand: true,
}

var roomsActionsRefreshClientToken = cli.Command{
	Name:    "refresh-client-token",
	Usage:   "Synchronously refresh an Client Token to join a Room. Client Token is necessary\nto join a Telnyx Room. Client Token will expire after `token_ttl_secs`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "room-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "refresh-token",
			Required: true,
			BodyPath: "refresh_token",
		},
		&requestflag.Flag[int64]{
			Name:     "token-ttl-secs",
			Usage:    "The time to live in seconds of the Client Token, after that time the Client Token is invalid and can't be used to join a Room.",
			Default:  600,
			BodyPath: "token_ttl_secs",
		},
	},
	Action:          handleRoomsActionsRefreshClientToken,
	HideHelpCommand: true,
}

func handleRoomsActionsGenerateJoinClientToken(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("room-id") && len(unusedArgs) > 0 {
		cmd.Set("room-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.RoomActionGenerateJoinClientTokenParams{}

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
	_, err = client.Rooms.Actions.GenerateJoinClientToken(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rooms:actions generate-join-client-token", obj, format, transform)
}

func handleRoomsActionsRefreshClientToken(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("room-id") && len(unusedArgs) > 0 {
		cmd.Set("room-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.RoomActionRefreshClientTokenParams{}

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
	_, err = client.Rooms.Actions.RefreshClientToken(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rooms:actions refresh-client-token", obj, format, transform)
}
