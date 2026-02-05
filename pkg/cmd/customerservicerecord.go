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

var customerServiceRecordsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a new customer service record for the provided phone number.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "phone-number",
			Usage:    "A valid US phone number in E164 format.",
			Required: true,
			BodyPath: "phone_number",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "additional-data",
			BodyPath: "additional_data",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "Callback URL to receive webhook notifications.",
			BodyPath: "webhook_url",
		},
	},
	Action:          handleCustomerServiceRecordsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"additional-data": {
		&requestflag.InnerFlag[string]{
			Name:       "additional-data.account-number",
			Usage:      "The account number of the customer service record.",
			InnerField: "account_number",
		},
		&requestflag.InnerFlag[string]{
			Name:       "additional-data.address-line-1",
			Usage:      "The first line of the address of the customer service record.",
			InnerField: "address_line_1",
		},
		&requestflag.InnerFlag[string]{
			Name:       "additional-data.authorized-person-name",
			Usage:      "The name of the authorized person.",
			InnerField: "authorized_person_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "additional-data.billing-phone-number",
			Usage:      "The billing phone number of the customer service record.",
			InnerField: "billing_phone_number",
		},
		&requestflag.InnerFlag[string]{
			Name:       "additional-data.city",
			Usage:      "The city of the customer service record.",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "additional-data.customer-code",
			Usage:      "The customer code of the customer service record.",
			InnerField: "customer_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "additional-data.name",
			Usage:      "The name of the administrator of CSR.",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "additional-data.pin",
			Usage:      "The PIN of the customer service record.",
			InnerField: "pin",
		},
		&requestflag.InnerFlag[string]{
			Name:       "additional-data.state",
			Usage:      "The state of the customer service record.",
			InnerField: "state",
		},
		&requestflag.InnerFlag[string]{
			Name:       "additional-data.zip-code",
			Usage:      "The zip code of the customer service record.",
			InnerField: "zip_code",
		},
	},
})

var customerServiceRecordsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a specific customer service record.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "customer-service-record-id",
			Required: true,
		},
	},
	Action:          handleCustomerServiceRecordsRetrieve,
	HideHelpCommand: true,
}

var customerServiceRecordsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List customer service records.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[phone_number][eq], filter[phone_number][in][], filter[status][eq], filter[status][in][], filter[created_at][lt], filter[created_at][gt]",
			QueryPath: "filter",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page[size]",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "sort",
			Usage:     "Consolidated sort parameter (deepObject style). Originally: sort[value]",
			QueryPath: "sort",
		},
	},
	Action:          handleCustomerServiceRecordsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.created-at",
			InnerField: "created_at",
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
	"sort": {
		&requestflag.InnerFlag[string]{
			Name:       "sort.value",
			Usage:      "Specifies the sort order for results. If not given, results are sorted by created_at in descending order.",
			InnerField: "value",
		},
	},
})

var customerServiceRecordsVerifyPhoneNumberCoverage = cli.Command{
	Name:    "verify-phone-number-coverage",
	Usage:   "Verify the coverage for a list of phone numbers.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "phone-number",
			Usage:    "The phone numbers list to be verified.",
			Required: true,
			BodyPath: "phone_numbers",
		},
	},
	Action:          handleCustomerServiceRecordsVerifyPhoneNumberCoverage,
	HideHelpCommand: true,
}

func handleCustomerServiceRecordsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CustomerServiceRecordNewParams{}

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
	_, err = client.CustomerServiceRecords.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "customer-service-records create", obj, format, transform)
}

func handleCustomerServiceRecordsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("customer-service-record-id") && len(unusedArgs) > 0 {
		cmd.Set("customer-service-record-id", unusedArgs[0])
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
	_, err = client.CustomerServiceRecords.Get(ctx, cmd.Value("customer-service-record-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "customer-service-records retrieve", obj, format, transform)
}

func handleCustomerServiceRecordsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CustomerServiceRecordListParams{}

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
		_, err = client.CustomerServiceRecords.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "customer-service-records list", obj, format, transform)
	} else {
		iter := client.CustomerServiceRecords.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "customer-service-records list", iter, format, transform)
	}
}

func handleCustomerServiceRecordsVerifyPhoneNumberCoverage(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.CustomerServiceRecordVerifyPhoneNumberCoverageParams{}

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
	_, err = client.CustomerServiceRecords.VerifyPhoneNumberCoverage(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "customer-service-records verify-phone-number-coverage", obj, format, transform)
}
