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

var reportsMdrUsageReportsCreate = cli.Command{
	Name:    "create",
	Usage:   "Submit request for new new messaging usage report. This endpoint will pull and\naggregate messaging data in specified time period.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "aggregation-type",
			Required: true,
			BodyPath: "aggregation_type",
		},
		&requestflag.Flag[any]{
			Name:     "end-date",
			Required: true,
			BodyPath: "end_date",
		},
		&requestflag.Flag[any]{
			Name:     "start-date",
			Required: true,
			BodyPath: "start_date",
		},
		&requestflag.Flag[string]{
			Name:     "profiles",
			BodyPath: "profiles",
		},
	},
	Action:          handleReportsMdrUsageReportsCreate,
	HideHelpCommand: true,
}

var reportsMdrUsageReportsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Fetch a single messaging usage report by id",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleReportsMdrUsageReportsRetrieve,
	HideHelpCommand: true,
}

var reportsMdrUsageReportsList = cli.Command{
	Name:    "list",
	Usage:   "Fetch all messaging usage reports. Usage reports are aggregated messaging data\nfor specified time period and breakdown",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "page-number",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page[size]",
		},
	},
	Action:          handleReportsMdrUsageReportsList,
	HideHelpCommand: true,
}

var reportsMdrUsageReportsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete messaging usage report by id",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleReportsMdrUsageReportsDelete,
	HideHelpCommand: true,
}

var reportsMdrUsageReportsFetchSync = cli.Command{
	Name:    "fetch-sync",
	Usage:   "Generate and fetch messaging usage report synchronously. This endpoint will both\ngenerate and fetch the messaging report over a specified time period. No polling\nis necessary but the response may take up to a couple of minutes.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "aggregation-type",
			Required:  true,
			QueryPath: "aggregation_type",
		},
		&requestflag.Flag[any]{
			Name:      "end-date",
			QueryPath: "end_date",
		},
		&requestflag.Flag[[]string]{
			Name:      "profile",
			QueryPath: "profiles",
		},
		&requestflag.Flag[any]{
			Name:      "start-date",
			QueryPath: "start_date",
		},
	},
	Action:          handleReportsMdrUsageReportsFetchSync,
	HideHelpCommand: true,
}

func handleReportsMdrUsageReportsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ReportMdrUsageReportNewParams{}

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
	_, err = client.Reports.MdrUsageReports.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "reports:mdr-usage-reports create", obj, format, transform)
}

func handleReportsMdrUsageReportsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Reports.MdrUsageReports.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "reports:mdr-usage-reports retrieve", obj, format, transform)
}

func handleReportsMdrUsageReportsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ReportMdrUsageReportListParams{}

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
		_, err = client.Reports.MdrUsageReports.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "reports:mdr-usage-reports list", obj, format, transform)
	} else {
		iter := client.Reports.MdrUsageReports.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "reports:mdr-usage-reports list", iter, format, transform)
	}
}

func handleReportsMdrUsageReportsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Reports.MdrUsageReports.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "reports:mdr-usage-reports delete", obj, format, transform)
}

func handleReportsMdrUsageReportsFetchSync(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ReportMdrUsageReportFetchSyncParams{}

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
	_, err = client.Reports.MdrUsageReports.FetchSync(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "reports:mdr-usage-reports fetch-sync", obj, format, transform)
}
