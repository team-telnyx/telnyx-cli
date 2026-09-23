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

var botSessionsList = cli.Command{
	Name:    "list",
	Usage:   "Consumes the one-time portal redirect (magic link) token emailed during bot\nsignup and returns an API session. The token is a UUIDv7 that encodes its\ncreation time; it expires after a configurable validity window (15 minutes by\ndefault) and is cleared on first use. Although the action creates a session, the\nroute uses the GET verb because it is opened from an email link. On first use\nthe account is also initialized. For bot signup (freemium) accounts the response\nis a minimal envelope containing only the `api_v2_token`; accounts that are\npermitted to use magic links but are not freemium accounts may instead receive\nan extended session payload when additional steps (such as two-factor\nauthentication or identity verification) are required. This endpoint is public;\nthe magic link token in the query string is the credential.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "email",
			Usage:     "Email address associated with the magic link token.",
			Required:  true,
			QueryPath: "email",
		},
		&requestflag.Flag[string]{
			Name:      "portal-redirect-token",
			Usage:     "Single-use portal redirect (magic link) token, a UUIDv7 sent to the account owner's email.",
			Required:  true,
			QueryPath: "portal_redirect_token",
		},
	},
	Action:          handleBotSessionsList,
	HideHelpCommand: true,
}

func handleBotSessionsList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.BotSessionListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.BotSessions.List(ctx, params, options...)
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
		Title:          "bot-sessions list",
		Transform:      transform,
	})
}
