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

var legacyReportingBatchDetailRecordsVoiceCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new CDR report request with the specified filters",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "end-time",
			Usage:    "End time in ISO format",
			Required: true,
			BodyPath: "end_time",
		},
		&requestflag.Flag[any]{
			Name:     "start-time",
			Usage:    "Start time in ISO format",
			Required: true,
			BodyPath: "start_time",
		},
		&requestflag.Flag[[]int64]{
			Name:     "call-type",
			Usage:    "List of call types to filter by (Inbound = 1, Outbound = 2)",
			BodyPath: "call_types",
		},
		&requestflag.Flag[[]int64]{
			Name:     "connection",
			Usage:    "List of connections to filter by",
			BodyPath: "connections",
		},
		&requestflag.Flag[[]string]{
			Name:     "field",
			Usage:    "Set of fields to include in the report",
			BodyPath: "fields",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "filter",
			Usage:    "List of filters to apply",
			BodyPath: "filters",
		},
		&requestflag.Flag[bool]{
			Name:     "include-all-metadata",
			Usage:    "Whether to include all metadata",
			BodyPath: "include_all_metadata",
		},
		&requestflag.Flag[[]string]{
			Name:     "managed-account",
			Usage:    "List of managed accounts to include",
			BodyPath: "managed_accounts",
		},
		&requestflag.Flag[[]int64]{
			Name:     "record-type",
			Usage:    "List of record types to filter by (Complete = 1, Incomplete = 2, Errors = 3)",
			BodyPath: "record_types",
		},
		&requestflag.Flag[string]{
			Name:     "report-name",
			Usage:    "Name of the report",
			BodyPath: "report_name",
		},
		&requestflag.Flag[bool]{
			Name:     "select-all-managed-accounts",
			Usage:    "Whether to select all managed accounts",
			BodyPath: "select_all_managed_accounts",
		},
		&requestflag.Flag[string]{
			Name:     "source",
			Usage:    "Source of the report. Valid values: calls (default), call-control, fax-api, webrtc",
			BodyPath: "source",
		},
		&requestflag.Flag[string]{
			Name:     "timezone",
			Usage:    "Timezone for the report",
			BodyPath: "timezone",
		},
	},
	Action:          handleLegacyReportingBatchDetailRecordsVoiceCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.billing-group",
			Usage:      "Billing group UUID to filter by",
			InnerField: "billing_group",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.cld",
			Usage:      "Called line identification (destination number)",
			InnerField: "cld",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.cld-filter",
			Usage:      "Filter type for CLD matching",
			InnerField: "cld_filter",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.cli",
			Usage:      "Calling line identification (caller ID)",
			InnerField: "cli",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.cli-filter",
			Usage:      "Filter type for CLI matching",
			InnerField: "cli_filter",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.filter-type",
			Usage:      "Logical operator for combining filters",
			InnerField: "filter_type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.tags-list",
			Usage:      "Tag name to filter by",
			InnerField: "tags_list",
		},
	},
})

var legacyReportingBatchDetailRecordsVoiceRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a specific CDR report request by ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleLegacyReportingBatchDetailRecordsVoiceRetrieve,
	HideHelpCommand: true,
}

var legacyReportingBatchDetailRecordsVoiceList = cli.Command{
	Name:            "list",
	Usage:           "Retrieves all CDR report requests for the authenticated user",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleLegacyReportingBatchDetailRecordsVoiceList,
	HideHelpCommand: true,
}

var legacyReportingBatchDetailRecordsVoiceDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a specific CDR report request by ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleLegacyReportingBatchDetailRecordsVoiceDelete,
	HideHelpCommand: true,
}

var legacyReportingBatchDetailRecordsVoiceRetrieveFields = cli.Command{
	Name:            "retrieve-fields",
	Usage:           "Retrieves all available fields that can be used in CDR reports",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleLegacyReportingBatchDetailRecordsVoiceRetrieveFields,
	HideHelpCommand: true,
}

func handleLegacyReportingBatchDetailRecordsVoiceCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.LegacyReportingBatchDetailRecordVoiceNewParams{}

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
	_, err = client.Legacy.Reporting.BatchDetailRecords.Voice.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legacy:reporting:batch-detail-records:voice create", obj, format, transform)
}

func handleLegacyReportingBatchDetailRecordsVoiceRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Legacy.Reporting.BatchDetailRecords.Voice.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legacy:reporting:batch-detail-records:voice retrieve", obj, format, transform)
}

func handleLegacyReportingBatchDetailRecordsVoiceList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Legacy.Reporting.BatchDetailRecords.Voice.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legacy:reporting:batch-detail-records:voice list", obj, format, transform)
}

func handleLegacyReportingBatchDetailRecordsVoiceDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Legacy.Reporting.BatchDetailRecords.Voice.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legacy:reporting:batch-detail-records:voice delete", obj, format, transform)
}

func handleLegacyReportingBatchDetailRecordsVoiceRetrieveFields(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Legacy.Reporting.BatchDetailRecords.Voice.GetFields(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legacy:reporting:batch-detail-records:voice retrieve-fields", obj, format, transform)
}
