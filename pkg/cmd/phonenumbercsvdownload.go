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

var phoneNumbersCsvDownloadsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a CSV download",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "csv-format",
			Usage:     "Which format to use when generating the CSV file. The default for backwards compatibility is 'V1'",
			Default:   "V1",
			QueryPath: "csv_format",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[has_bundle], filter[tag], filter[connection_id], filter[phone_number], filter[status], filter[voice.connection_name], filter[voice.usage_payment_method], filter[billing_group_id], filter[emergency_address_id], filter[customer_reference]",
			QueryPath: "filter",
		},
	},
	Action:          handlePhoneNumbersCsvDownloadsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.billing-group-id",
			Usage:      "Filter by the billing_group_id associated with phone numbers. To filter to only phone numbers that have no billing group associated them, set the value of this filter to the string 'null'.",
			InnerField: "billing_group_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.connection-id",
			Usage:      "Filter by connection_id.",
			InnerField: "connection_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.customer-reference",
			Usage:      "Filter numbers via the customer_reference set.",
			InnerField: "customer_reference",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.emergency-address-id",
			Usage:      "Filter by the emergency_address_id associated with phone numbers. To filter only phone numbers that have no emergency address associated with them, set the value of this filter to the string 'null'.",
			InnerField: "emergency_address_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.has-bundle",
			Usage:      "Filter by phone number that have bundles.",
			InnerField: "has_bundle",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.phone-number",
			Usage:      "Filter by phone number. Requires at least three digits.\n             Non-numerical characters will result in no values being returned.",
			InnerField: "phone_number",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.status",
			Usage:      "Filter by phone number status.",
			InnerField: "status",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.tag",
			Usage:      "Filter by phone number tags.",
			InnerField: "tag",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.voice-connection-name",
			Usage:      "Filter by voice connection name pattern matching.",
			InnerField: "voice.connection_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.voice-usage-payment-method",
			Usage:      "Filter by usage_payment_method.",
			InnerField: "voice.usage_payment_method",
		},
	},
})

var phoneNumbersCsvDownloadsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a CSV download",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handlePhoneNumbersCsvDownloadsRetrieve,
	HideHelpCommand: true,
}

var phoneNumbersCsvDownloadsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List CSV downloads",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "page",
			Usage:     "Consolidated page parameter (deepObject style). Originally: page[size], page[number]",
			QueryPath: "page",
		},
	},
	Action:          handlePhoneNumbersCsvDownloadsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
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

func handlePhoneNumbersCsvDownloadsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PhoneNumberCsvDownloadNewParams{}

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
	_, err = client.PhoneNumbers.CsvDownloads.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "phone-numbers:csv-downloads create", obj, format, transform)
}

func handlePhoneNumbersCsvDownloadsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.PhoneNumbers.CsvDownloads.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "phone-numbers:csv-downloads retrieve", obj, format, transform)
}

func handlePhoneNumbersCsvDownloadsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PhoneNumberCsvDownloadListParams{}

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
		_, err = client.PhoneNumbers.CsvDownloads.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "phone-numbers:csv-downloads list", obj, format, transform)
	} else {
		iter := client.PhoneNumbers.CsvDownloads.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "phone-numbers:csv-downloads list", iter, format, transform)
	}
}
