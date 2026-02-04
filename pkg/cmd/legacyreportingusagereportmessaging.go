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

var legacyReportingUsageReportsMessagingCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a new legacy usage V2 MDR report request with the specified filters",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "aggregation-type",
			Usage:    "Aggregation type: No aggregation = 0, By Messaging Profile = 1, By Tags = 2",
			Required: true,
			BodyPath: "aggregation_type",
		},
		&requestflag.Flag[any]{
			Name:     "end-time",
			BodyPath: "end_time",
		},
		&requestflag.Flag[[]string]{
			Name:     "managed-account",
			Usage:    "List of managed accounts to include",
			BodyPath: "managed_accounts",
		},
		&requestflag.Flag[[]string]{
			Name:     "profile",
			Usage:    "List of messaging profile IDs to filter by",
			BodyPath: "profiles",
		},
		&requestflag.Flag[bool]{
			Name:     "select-all-managed-accounts",
			BodyPath: "select_all_managed_accounts",
		},
		&requestflag.Flag[any]{
			Name:     "start-time",
			BodyPath: "start_time",
		},
	},
	Action:          handleLegacyReportingUsageReportsMessagingCreate,
	HideHelpCommand: true,
}

var legacyReportingUsageReportsMessagingRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Fetch single MDR usage report by id.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleLegacyReportingUsageReportsMessagingRetrieve,
	HideHelpCommand: true,
}

var legacyReportingUsageReportsMessagingList = cli.Command{
	Name:    "list",
	Usage:   "Fetch all previous requests for MDR usage reports.",
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
	Action:          handleLegacyReportingUsageReportsMessagingList,
	HideHelpCommand: true,
}

var legacyReportingUsageReportsMessagingDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a specific V2 legacy usage MDR report request by ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleLegacyReportingUsageReportsMessagingDelete,
	HideHelpCommand: true,
}

func handleLegacyReportingUsageReportsMessagingCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.LegacyReportingUsageReportMessagingNewParams{}

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
	_, err = client.Legacy.Reporting.UsageReports.Messaging.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legacy:reporting:usage-reports:messaging create", obj, format, transform)
}

func handleLegacyReportingUsageReportsMessagingRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Legacy.Reporting.UsageReports.Messaging.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legacy:reporting:usage-reports:messaging retrieve", obj, format, transform)
}

func handleLegacyReportingUsageReportsMessagingList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.LegacyReportingUsageReportMessagingListParams{}

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
		_, err = client.Legacy.Reporting.UsageReports.Messaging.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "legacy:reporting:usage-reports:messaging list", obj, format, transform)
	} else {
		iter := client.Legacy.Reporting.UsageReports.Messaging.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "legacy:reporting:usage-reports:messaging list", iter, format, transform)
	}
}

func handleLegacyReportingUsageReportsMessagingDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Legacy.Reporting.UsageReports.Messaging.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legacy:reporting:usage-reports:messaging delete", obj, format, transform)
}
