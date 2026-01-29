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

var legacyReportingUsageReportsNumberLookupCreate = cli.Command{
	Name:    "create",
	Usage:   "Submit a new telco data usage report",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "aggregation-type",
			Usage:    "Type of aggregation for the report",
			BodyPath: "aggregationType",
		},
		&requestflag.Flag[any]{
			Name:     "end-date",
			Usage:    "End date for the usage report",
			BodyPath: "endDate",
		},
		&requestflag.Flag[[]string]{
			Name:     "managed-account",
			Usage:    "List of managed accounts to include in the report",
			BodyPath: "managedAccounts",
		},
		&requestflag.Flag[any]{
			Name:     "start-date",
			Usage:    "Start date for the usage report",
			BodyPath: "startDate",
		},
	},
	Action:          handleLegacyReportingUsageReportsNumberLookupCreate,
	HideHelpCommand: true,
}

var legacyReportingUsageReportsNumberLookupRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a specific telco data usage report by its ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleLegacyReportingUsageReportsNumberLookupRetrieve,
	HideHelpCommand: true,
}

var legacyReportingUsageReportsNumberLookupList = cli.Command{
	Name:            "list",
	Usage:           "Retrieve a paginated list of telco data usage reports",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleLegacyReportingUsageReportsNumberLookupList,
	HideHelpCommand: true,
}

var legacyReportingUsageReportsNumberLookupDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a specific telco data usage report by its ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleLegacyReportingUsageReportsNumberLookupDelete,
	HideHelpCommand: true,
}

func handleLegacyReportingUsageReportsNumberLookupCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.LegacyReportingUsageReportNumberLookupNewParams{}

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
	_, err = client.Legacy.Reporting.UsageReports.NumberLookup.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legacy:reporting:usage-reports:number-lookup create", obj, format, transform)
}

func handleLegacyReportingUsageReportsNumberLookupRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Legacy.Reporting.UsageReports.NumberLookup.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legacy:reporting:usage-reports:number-lookup retrieve", obj, format, transform)
}

func handleLegacyReportingUsageReportsNumberLookupList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Legacy.Reporting.UsageReports.NumberLookup.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "legacy:reporting:usage-reports:number-lookup list", obj, format, transform)
}

func handleLegacyReportingUsageReportsNumberLookupDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.Legacy.Reporting.UsageReports.NumberLookup.Delete(ctx, cmd.Value("id").(string), options...)
}
