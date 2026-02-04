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

var legacyReportingUsageReportsVoiceCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a new legacy usage V2 CDR report request with the specified filters",
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
		&requestflag.Flag[int64]{
			Name:     "aggregation-type",
			Usage:    "Aggregation type: All = 0, By Connections = 1, By Tags = 2, By Billing Group = 3",
			BodyPath: "aggregation_type",
		},
		&requestflag.Flag[[]int64]{
			Name:     "connection",
			Usage:    "List of connections to filter by",
			BodyPath: "connections",
		},
		&requestflag.Flag[[]string]{
			Name:     "managed-account",
			Usage:    "List of managed accounts to include",
			BodyPath: "managed_accounts",
		},
		&requestflag.Flag[int64]{
			Name:     "product-breakdown",
			Usage:    "Product breakdown type: No breakdown = 0, DID vs Toll-free = 1, Country = 2, DID vs Toll-free per Country = 3",
			BodyPath: "product_breakdown",
		},
		&requestflag.Flag[bool]{
			Name:     "select-all-managed-accounts",
			Usage:    "Whether to select all managed accounts",
			BodyPath: "select_all_managed_accounts",
		},
	},
	Action:          handleLegacyReportingUsageReportsVoiceCreate,
	HideHelpCommand: true,
}

var legacyReportingUsageReportsVoiceRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Fetch single cdr usage report by id.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleLegacyReportingUsageReportsVoiceRetrieve,
	HideHelpCommand: true,
}

var legacyReportingUsageReportsVoiceList = cli.Command{
	Name:    "list",
	Usage:   "Fetch all previous requests for cdr usage reports.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "page",
			Usage:     "Page number",
			Default:   1,
			QueryPath: "page",
		},
		&requestflag.Flag[int64]{
			Name:      "per-page",
			Usage:     "Size of the page",
			Default:   20,
			QueryPath: "per_page",
		},
	},
	Action:          handleLegacyReportingUsageReportsVoiceList,
	HideHelpCommand: true,
}

var legacyReportingUsageReportsVoiceDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a specific V2 legacy usage CDR report request by ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleLegacyReportingUsageReportsVoiceDelete,
	HideHelpCommand: true,
}

func handleLegacyReportingUsageReportsVoiceCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.LegacyReportingUsageReportVoiceNewParams{}

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
	_, err = client.Legacy.Reporting.UsageReports.Voice.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legacy:reporting:usage-reports:voice create", obj, format, transform)
}

func handleLegacyReportingUsageReportsVoiceRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Legacy.Reporting.UsageReports.Voice.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legacy:reporting:usage-reports:voice retrieve", obj, format, transform)
}

func handleLegacyReportingUsageReportsVoiceList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.LegacyReportingUsageReportVoiceListParams{}

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
		_, err = client.Legacy.Reporting.UsageReports.Voice.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "legacy:reporting:usage-reports:voice list", obj, format, transform)
	} else {
		iter := client.Legacy.Reporting.UsageReports.Voice.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "legacy:reporting:usage-reports:voice list", iter, format, transform)
	}
}

func handleLegacyReportingUsageReportsVoiceDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Legacy.Reporting.UsageReports.Voice.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legacy:reporting:usage-reports:voice delete", obj, format, transform)
}
