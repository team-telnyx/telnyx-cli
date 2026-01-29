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

var externalConnectionsUploadsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a new Upload request to Microsoft teams with the included phone numbers.\nOnly one of civic_address_id or location_id must be provided, not both. The\nmaximum allowed phone numbers for the numbers_ids array is 1000.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[[]string]{
			Name:     "number-id",
			Required: true,
			BodyPath: "number_ids",
		},
		&requestflag.Flag[[]string]{
			Name:     "additional-usage",
			BodyPath: "additional_usages",
		},
		&requestflag.Flag[string]{
			Name:     "civic-address-id",
			Usage:    "Identifies the civic address to assign all phone numbers to.",
			BodyPath: "civic_address_id",
		},
		&requestflag.Flag[string]{
			Name:     "location-id",
			Usage:    "Identifies the location to assign all phone numbers to.",
			BodyPath: "location_id",
		},
		&requestflag.Flag[string]{
			Name:     "usage",
			Usage:    "The use case of the upload request. NOTE: `calling_user_assignment` is not supported for toll free numbers.",
			BodyPath: "usage",
		},
	},
	Action:          handleExternalConnectionsUploadsCreate,
	HideHelpCommand: true,
}

var externalConnectionsUploadsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Return the details of an Upload request and its phone numbers.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "ticket-id",
			Required: true,
		},
	},
	Action:          handleExternalConnectionsUploadsRetrieve,
	HideHelpCommand: true,
}

var externalConnectionsUploadsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Returns a list of your Upload requests for the given external connection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Filter parameter for uploads (deepObject style). Supports filtering by status, civic_address_id, location_id, and phone_number with eq/contains operations.",
			QueryPath: "filter",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "page",
			Usage:     "Consolidated page parameter (deepObject style). Originally: page[size], page[number]",
			QueryPath: "page",
		},
	},
	Action:          handleExternalConnectionsUploadsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.civic-address-id",
			InnerField: "civic_address_id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.location-id",
			InnerField: "location_id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.phone-number",
			InnerField: "phone_number",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.status",
			InnerField: "status",
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

var externalConnectionsUploadsPendingCount = cli.Command{
	Name:    "pending-count",
	Usage:   "Returns the count of all pending upload requests for the given external\nconnection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleExternalConnectionsUploadsPendingCount,
	HideHelpCommand: true,
}

var externalConnectionsUploadsRefreshStatus = cli.Command{
	Name:    "refresh-status",
	Usage:   "Forces a recheck of the status of all pending Upload requests for the given\nexternal connection in the background.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleExternalConnectionsUploadsRefreshStatus,
	HideHelpCommand: true,
}

var externalConnectionsUploadsRetry = cli.Command{
	Name:    "retry",
	Usage:   "If there were any errors during the upload process, this endpoint will retry the\nupload request. In some cases this will reattempt the existing upload request,\nin other cases it may create a new upload request. Please check the ticket_id in\nthe response to determine if a new upload request was created.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "ticket-id",
			Required: true,
		},
	},
	Action:          handleExternalConnectionsUploadsRetry,
	HideHelpCommand: true,
}

func handleExternalConnectionsUploadsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ExternalConnectionUploadNewParams{}

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
	_, err = client.ExternalConnections.Uploads.New(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "external-connections:uploads create", obj, format, transform)
}

func handleExternalConnectionsUploadsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("ticket-id") && len(unusedArgs) > 0 {
		cmd.Set("ticket-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ExternalConnectionUploadGetParams{
		ID: cmd.Value("id").(string),
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
	_, err = client.ExternalConnections.Uploads.Get(
		ctx,
		cmd.Value("ticket-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "external-connections:uploads retrieve", obj, format, transform)
}

func handleExternalConnectionsUploadsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ExternalConnectionUploadListParams{}

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
		_, err = client.ExternalConnections.Uploads.List(
			ctx,
			cmd.Value("id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "external-connections:uploads list", obj, format, transform)
	} else {
		iter := client.ExternalConnections.Uploads.ListAutoPaging(
			ctx,
			cmd.Value("id").(string),
			params,
			options...,
		)
		return ShowJSONIterator(os.Stdout, "external-connections:uploads list", iter, format, transform)
	}
}

func handleExternalConnectionsUploadsPendingCount(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.ExternalConnections.Uploads.PendingCount(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "external-connections:uploads pending-count", obj, format, transform)
}

func handleExternalConnectionsUploadsRefreshStatus(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.ExternalConnections.Uploads.RefreshStatus(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "external-connections:uploads refresh-status", obj, format, transform)
}

func handleExternalConnectionsUploadsRetry(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("ticket-id") && len(unusedArgs) > 0 {
		cmd.Set("ticket-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ExternalConnectionUploadRetryParams{
		ID: cmd.Value("id").(string),
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
	_, err = client.ExternalConnections.Uploads.Retry(
		ctx,
		cmd.Value("ticket-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "external-connections:uploads retry", obj, format, transform)
}
