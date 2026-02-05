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

var legacyReportingBatchDetailRecordsSpeechToTextCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a new Speech to Text batch report request with the specified filters",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "end-date",
			Usage:    "End date in ISO format with timezone (date range must be up to one month)",
			Required: true,
			BodyPath: "end_date",
		},
		&requestflag.Flag[any]{
			Name:     "start-date",
			Usage:    "Start date in ISO format with timezone",
			Required: true,
			BodyPath: "start_date",
		},
	},
	Action:          handleLegacyReportingBatchDetailRecordsSpeechToTextCreate,
	HideHelpCommand: true,
}

var legacyReportingBatchDetailRecordsSpeechToTextRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a specific Speech to Text batch report request by ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleLegacyReportingBatchDetailRecordsSpeechToTextRetrieve,
	HideHelpCommand: true,
}

var legacyReportingBatchDetailRecordsSpeechToTextList = cli.Command{
	Name:            "list",
	Usage:           "Retrieves all Speech to Text batch report requests for the authenticated user",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleLegacyReportingBatchDetailRecordsSpeechToTextList,
	HideHelpCommand: true,
}

var legacyReportingBatchDetailRecordsSpeechToTextDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a specific Speech to Text batch report request by ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleLegacyReportingBatchDetailRecordsSpeechToTextDelete,
	HideHelpCommand: true,
}

func handleLegacyReportingBatchDetailRecordsSpeechToTextCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.LegacyReportingBatchDetailRecordSpeechToTextNewParams{}

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
	_, err = client.Legacy.Reporting.BatchDetailRecords.SpeechToText.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legacy:reporting:batch-detail-records:speech-to-text create", obj, format, transform)
}

func handleLegacyReportingBatchDetailRecordsSpeechToTextRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Legacy.Reporting.BatchDetailRecords.SpeechToText.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legacy:reporting:batch-detail-records:speech-to-text retrieve", obj, format, transform)
}

func handleLegacyReportingBatchDetailRecordsSpeechToTextList(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Legacy.Reporting.BatchDetailRecords.SpeechToText.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legacy:reporting:batch-detail-records:speech-to-text list", obj, format, transform)
}

func handleLegacyReportingBatchDetailRecordsSpeechToTextDelete(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Legacy.Reporting.BatchDetailRecords.SpeechToText.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legacy:reporting:batch-detail-records:speech-to-text delete", obj, format, transform)
}
