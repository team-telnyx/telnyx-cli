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

var numberReservationsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a Phone Number Reservation for multiple numbers.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "customer-reference",
			Usage:    "A customer reference string for customer look ups.",
			BodyPath: "customer_reference",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "phone-number",
			BodyPath: "phone_numbers",
		},
	},
	Action:          handleNumberReservationsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"phone-number": {
		&requestflag.InnerFlag[string]{
			Name:       "phone-number.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[any]{
			Name:       "phone-number.created-at",
			Usage:      "An ISO 8901 datetime string denoting when the individual number reservation was created.",
			InnerField: "created_at",
		},
		&requestflag.InnerFlag[any]{
			Name:       "phone-number.expired-at",
			Usage:      "An ISO 8901 datetime string for when the individual number reservation is going to expire",
			InnerField: "expired_at",
		},
		&requestflag.InnerFlag[string]{
			Name:       "phone-number.phone-number",
			InnerField: "phone_number",
		},
		&requestflag.InnerFlag[string]{
			Name:       "phone-number.record-type",
			InnerField: "record_type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "phone-number.status",
			Usage:      "The status of the phone number's reservation.",
			InnerField: "status",
		},
		&requestflag.InnerFlag[any]{
			Name:       "phone-number.updated-at",
			Usage:      "An ISO 8901 datetime string for when the the individual number reservation was updated.",
			InnerField: "updated_at",
		},
	},
})

var numberReservationsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Gets a single phone number reservation.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "number-reservation-id",
			Required: true,
		},
	},
	Action:          handleNumberReservationsRetrieve,
	HideHelpCommand: true,
}

var numberReservationsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Gets a paginated list of phone number reservations.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[status], filter[created_at], filter[phone_numbers.phone_number], filter[customer_reference]",
			QueryPath: "filter",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "page",
			Usage:     "Consolidated page parameter (deepObject style). Originally: page[size], page[number]",
			QueryPath: "page",
		},
	},
	Action:          handleNumberReservationsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.created-at",
			Usage:      "Filter number reservations by date range.",
			InnerField: "created_at",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.customer-reference",
			Usage:      "Filter number reservations via the customer reference set.",
			InnerField: "customer_reference",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.phone-numbers-phone-number",
			Usage:      "Filter number reservations having these phone numbers.",
			InnerField: "phone_numbers.phone_number",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.status",
			Usage:      "Filter number reservations by status.",
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

func handleNumberReservationsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.NumberReservationNewParams{}

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
	_, err = client.NumberReservations.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "number-reservations create", obj, format, transform)
}

func handleNumberReservationsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("number-reservation-id") && len(unusedArgs) > 0 {
		cmd.Set("number-reservation-id", unusedArgs[0])
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
	_, err = client.NumberReservations.Get(ctx, cmd.Value("number-reservation-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "number-reservations retrieve", obj, format, transform)
}

func handleNumberReservationsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.NumberReservationListParams{}

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
		_, err = client.NumberReservations.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "number-reservations list", obj, format, transform)
	} else {
		iter := client.NumberReservations.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "number-reservations list", iter, format, transform)
	}
}
