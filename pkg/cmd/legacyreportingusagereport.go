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

var legacyReportingUsageReportsRetrieveSpeechToText = cli.Command{
	Name:    "retrieve-speech-to-text",
	Usage:   "Generate and fetch speech to text usage report synchronously. This endpoint will\nboth generate and fetch the speech to text report over a specified time period.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "end-date",
			QueryPath: "end_date",
		},
		&requestflag.Flag[any]{
			Name:      "start-date",
			QueryPath: "start_date",
		},
	},
	Action:          handleLegacyReportingUsageReportsRetrieveSpeechToText,
	HideHelpCommand: true,
}

func handleLegacyReportingUsageReportsRetrieveSpeechToText(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.LegacyReportingUsageReportGetSpeechToTextParams{}

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
	_, err = client.Legacy.Reporting.UsageReports.GetSpeechToText(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legacy:reporting:usage-reports retrieve-speech-to-text", obj, format, transform)
}
