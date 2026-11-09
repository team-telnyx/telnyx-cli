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
	"github.com/urfave/cli/v3"
)

var meetingBotCalendarRetrieveOAuthCallback = cli.Command{
	Name:    "retrieve-oauth-callback",
	Usage:   "Browser redirect after calendar consent. Use the returned oauth_url rather than\nconstructing this request. Verified state identifies the connecting user.\nReturns HTML for success, cancellation, or failure; no JSON error envelope or\nbearer token is used.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "code",
			Usage:     "Authorization code supplied by the calendar provider.",
			QueryPath: "code",
		},
		&requestflag.Flag[string]{
			Name:      "error",
			Usage:     "Authorization error supplied by the calendar provider, such as access_denied.",
			QueryPath: "error",
		},
		&requestflag.Flag[string]{
			Name:      "state",
			Usage:     "Unmodified signed state from the authorization URL.",
			QueryPath: "state",
		},
	},
	Action:          handleMeetingBotCalendarRetrieveOAuthCallback,
	HideHelpCommand: true,
}

func handleMeetingBotCalendarRetrieveOAuthCallback(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.MeetingBotCalendarGetOAuthCallbackParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.MeetingBot.Calendar.GetOAuthCallback(ctx, params, options...)
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(res)
	return err
}
