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

var sessionAnalysisRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a full session analysis tree for a given event, including costs, child\nevents, and product linkages.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "record-type",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "event-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:      "date-time",
			Usage:     "ISO 8601 timestamp or date to narrow index selection for faster lookups. Accepts full datetime (e.g., 2026-03-17T10:00:00Z) or date-only format (e.g., 2026-03-17).",
			QueryPath: "date_time",
		},
		&requestflag.Flag[string]{
			Name:      "expand",
			Usage:     "Controls what data to expand on each event node.",
			Default:   "record",
			QueryPath: "expand",
		},
		&requestflag.Flag[bool]{
			Name:      "include-children",
			Usage:     "Whether to include child events in the response.",
			Default:   true,
			QueryPath: "include_children",
		},
		&requestflag.Flag[int64]{
			Name:      "max-depth",
			Usage:     "Maximum traversal depth for the event tree.",
			Default:   2,
			QueryPath: "max_depth",
		},
	},
	Action:          handleSessionAnalysisRetrieve,
	HideHelpCommand: true,
}

func handleSessionAnalysisRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("event-id") && len(unusedArgs) > 0 {
		cmd.Set("event-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SessionAnalysisGetParams{
		RecordType: cmd.Value("record-type").(string),
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
	_, err = client.SessionAnalysis.Get(
		ctx,
		cmd.Value("event-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "session-analysis retrieve", obj, format, transform)
}
