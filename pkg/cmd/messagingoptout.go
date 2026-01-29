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

var messagingOptoutsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Retrieve a list of opt-out blocks.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "created-at",
			Usage:     "Consolidated created_at parameter (deepObject style). Originally: created_at[gte], created_at[lte]",
			QueryPath: "created_at",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[messaging_profile_id], filter[from]",
			QueryPath: "filter",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "page",
			Usage:     "Consolidated page parameter (deepObject style). Originally: page[number], page[size]",
			QueryPath: "page",
		},
		&requestflag.Flag[string]{
			Name:      "redaction-enabled",
			Usage:     "If receiving address (+E.164 formatted phone number) should be redacted",
			QueryPath: "redaction_enabled",
		},
	},
	Action:          handleMessagingOptoutsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"created-at": {
		&requestflag.InnerFlag[any]{
			Name:       "created-at.gte",
			Usage:      "Filter opt-outs created after this date (ISO-8601 format)",
			InnerField: "gte",
		},
		&requestflag.InnerFlag[any]{
			Name:       "created-at.lte",
			Usage:      "Filter opt-outs created before this date (ISO-8601 format)",
			InnerField: "lte",
		},
	},
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.from",
			Usage:      "The sending address (+E.164 formatted phone number, alphanumeric sender ID, or short code) to retrieve opt-outs for",
			InnerField: "from",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.messaging-profile-id",
			Usage:      "The ID of the messaging profile to retrieve opt-outs for",
			InnerField: "messaging_profile_id",
		},
	},
	"page": {
		&requestflag.InnerFlag[int64]{
			Name:       "page.number",
			Usage:      "The page number to load",
			InnerField: "number",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "page.size",
			Usage:      "The size of the page",
			InnerField: "size",
		},
	},
})

func handleMessagingOptoutsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MessagingOptoutListParams{}

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
		_, err = client.MessagingOptouts.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "messaging-optouts list", obj, format, transform)
	} else {
		iter := client.MessagingOptouts.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "messaging-optouts list", iter, format, transform)
	}
}
