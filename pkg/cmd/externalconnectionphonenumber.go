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

var externalConnectionsPhoneNumbersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Return the details of a phone number associated with the given external\nconnection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "phone-number-id",
			Required: true,
		},
	},
	Action:          handleExternalConnectionsPhoneNumbersRetrieve,
	HideHelpCommand: true,
}

var externalConnectionsPhoneNumbersUpdate = cli.Command{
	Name:    "update",
	Usage:   "Asynchronously update settings of the phone number associated with the given\nexternal connection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "phone-number-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "location-id",
			Usage:    "Identifies the location to assign the phone number to.",
			BodyPath: "location_id",
		},
	},
	Action:          handleExternalConnectionsPhoneNumbersUpdate,
	HideHelpCommand: true,
}

var externalConnectionsPhoneNumbersList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Returns a list of all active phone numbers associated with the given external\nconnection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Filter parameter for phone numbers (deepObject style). Supports filtering by phone_number, civic_address_id, and location_id with eq/contains operations.",
			QueryPath: "filter",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "page",
			Usage:     "Consolidated page parameter (deepObject style). Originally: page[size], page[number]",
			QueryPath: "page",
		},
	},
	Action:          handleExternalConnectionsPhoneNumbersList,
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

func handleExternalConnectionsPhoneNumbersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("phone-number-id") && len(unusedArgs) > 0 {
		cmd.Set("phone-number-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ExternalConnectionPhoneNumberGetParams{
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
	_, err = client.ExternalConnections.PhoneNumbers.Get(
		ctx,
		cmd.Value("phone-number-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "external-connections:phone-numbers retrieve", obj, format, transform)
}

func handleExternalConnectionsPhoneNumbersUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("phone-number-id") && len(unusedArgs) > 0 {
		cmd.Set("phone-number-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ExternalConnectionPhoneNumberUpdateParams{
		ID: cmd.Value("id").(string),
	}

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
	_, err = client.ExternalConnections.PhoneNumbers.Update(
		ctx,
		cmd.Value("phone-number-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "external-connections:phone-numbers update", obj, format, transform)
}

func handleExternalConnectionsPhoneNumbersList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ExternalConnectionPhoneNumberListParams{}

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
		_, err = client.ExternalConnections.PhoneNumbers.List(
			ctx,
			cmd.Value("id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "external-connections:phone-numbers list", obj, format, transform)
	} else {
		iter := client.ExternalConnections.PhoneNumbers.ListAutoPaging(
			ctx,
			cmd.Value("id").(string),
			params,
			options...,
		)
		return ShowJSONIterator(os.Stdout, "external-connections:phone-numbers list", iter, format, transform)
	}
}
