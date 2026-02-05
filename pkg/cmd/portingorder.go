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

var portingOrdersCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a new porting order object.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "phone-number",
			Usage:    "The list of +E.164 formatted phone numbers",
			Required: true,
			BodyPath: "phone_numbers",
		},
		&requestflag.Flag[string]{
			Name:     "customer-group-reference",
			Usage:    "A customer-specified group reference for customer bookkeeping purposes",
			BodyPath: "customer_group_reference",
		},
		&requestflag.Flag[any]{
			Name:     "customer-reference",
			Usage:    "A customer-specified reference number for customer bookkeeping purposes",
			BodyPath: "customer_reference",
		},
	},
	Action:          handlePortingOrdersCreate,
	HideHelpCommand: true,
}

var portingOrdersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves the details of an existing porting order.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:      "include-phone-numbers",
			Usage:     "Include the first 50 phone number objects in the results",
			Default:   true,
			QueryPath: "include_phone_numbers",
		},
	},
	Action:          handlePortingOrdersRetrieve,
	HideHelpCommand: true,
}

var portingOrdersUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Edits the details of an existing porting order.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "activation-settings",
			BodyPath: "activation_settings",
		},
		&requestflag.Flag[string]{
			Name:     "customer-group-reference",
			BodyPath: "customer_group_reference",
		},
		&requestflag.Flag[string]{
			Name:     "customer-reference",
			BodyPath: "customer_reference",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "documents",
			Usage:    "Can be specified directly or via the `requirement_group_id` parameter.",
			BodyPath: "documents",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "end-user",
			BodyPath: "end_user",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "messaging",
			BodyPath: "messaging",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "misc",
			BodyPath: "misc",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "phone-number-configuration",
			BodyPath: "phone_number_configuration",
		},
		&requestflag.Flag[string]{
			Name:     "requirement-group-id",
			Usage:    "If present, we will read the current values from the specified Requirement Group into the Documents and Requirements for this Porting Order. Note that any future changes in the Requirement Group would have no impact on this Porting Order. We will return an error if a specified Requirement Group conflicts with documents or requirements in the same request.",
			BodyPath: "requirement_group_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "requirement",
			Usage:    "List of requirements for porting numbers. ",
			BodyPath: "requirements",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "user-feedback",
			BodyPath: "user_feedback",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			BodyPath: "webhook_url",
		},
	},
	Action:          handlePortingOrdersUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"activation-settings": {
		&requestflag.InnerFlag[any]{
			Name:       "activation-settings.foc-datetime-requested",
			Usage:      "ISO 8601 formatted Date/Time requested for the FOC date",
			InnerField: "foc_datetime_requested",
		},
	},
	"documents": {
		&requestflag.InnerFlag[any]{
			Name:       "documents.invoice",
			Usage:      "Returned ID of the submitted Invoice via the Documents endpoint",
			InnerField: "invoice",
		},
		&requestflag.InnerFlag[any]{
			Name:       "documents.loa",
			Usage:      "Returned ID of the submitted LOA via the Documents endpoint",
			InnerField: "loa",
		},
	},
	"end-user": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "end-user.admin",
			InnerField: "admin",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "end-user.location",
			InnerField: "location",
		},
	},
	"messaging": {
		&requestflag.InnerFlag[bool]{
			Name:       "messaging.enable-messaging",
			Usage:      "Indicates whether Telnyx will port messaging capabilities from the losing carrier. If false, any messaging capabilities will stay with their current provider.",
			InnerField: "enable_messaging",
		},
	},
	"misc": {
		&requestflag.InnerFlag[any]{
			Name:       "misc.new-billing-phone-number",
			Usage:      "New billing phone number for the remaining numbers. Used in case the current billing phone number is being ported to Telnyx. This will be set on your account with your current service provider and should be one of the numbers remaining on that account.",
			InnerField: "new_billing_phone_number",
		},
		&requestflag.InnerFlag[any]{
			Name:       "misc.remaining-numbers-action",
			Usage:      "Remaining numbers can be either kept with their current service provider or disconnected. 'new_billing_telephone_number' is required when 'remaining_numbers_action' is 'keep'.",
			InnerField: "remaining_numbers_action",
		},
		&requestflag.InnerFlag[string]{
			Name:       "misc.type",
			Usage:      "A port can be either 'full' or 'partial'. When type is 'full' the other attributes should be omitted.",
			InnerField: "type",
		},
	},
	"phone-number-configuration": {
		&requestflag.InnerFlag[any]{
			Name:       "phone-number-configuration.billing-group-id",
			Usage:      "identifies the billing group to set on the numbers when ported",
			InnerField: "billing_group_id",
		},
		&requestflag.InnerFlag[any]{
			Name:       "phone-number-configuration.connection-id",
			Usage:      "identifies the connection to set on the numbers when ported",
			InnerField: "connection_id",
		},
		&requestflag.InnerFlag[any]{
			Name:       "phone-number-configuration.emergency-address-id",
			Usage:      "identifies the emergency address to set on the numbers when ported",
			InnerField: "emergency_address_id",
		},
		&requestflag.InnerFlag[any]{
			Name:       "phone-number-configuration.messaging-profile-id",
			Usage:      "identifies the messaging profile to set on the numbers when ported",
			InnerField: "messaging_profile_id",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "phone-number-configuration.tags",
			InnerField: "tags",
		},
	},
	"requirement": {
		&requestflag.InnerFlag[string]{
			Name:       "requirement.field-value",
			Usage:      "identifies the document or provides the text value that satisfies this requirement",
			InnerField: "field_value",
		},
		&requestflag.InnerFlag[string]{
			Name:       "requirement.requirement-type-id",
			Usage:      "Identifies the requirement type that the `field_value` fulfills",
			InnerField: "requirement_type_id",
		},
	},
	"user-feedback": {
		&requestflag.InnerFlag[any]{
			Name:       "user-feedback.user-comment",
			Usage:      "A comment related to the customer rating.",
			InnerField: "user_comment",
		},
		&requestflag.InnerFlag[any]{
			Name:       "user-feedback.user-rating",
			Usage:      "Once an order is ported, cancellation is requested or the request is cancelled, the user may rate their experience",
			InnerField: "user_rating",
		},
	},
})

var portingOrdersList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Returns a list of your porting order.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[customer_reference], filter[customer_group_reference], filter[parent_support_key], filter[phone_numbers.country_code], filter[phone_numbers.carrier_name], filter[misc.type], filter[end_user.admin.entity_name], filter[end_user.admin.auth_person_name], filter[activation_settings.fast_port_eligible], filter[activation_settings.foc_datetime_requested][gt], filter[activation_settings.foc_datetime_requested][lt], filter[phone_numbers.phone_number][contains]",
			QueryPath: "filter",
		},
		&requestflag.Flag[bool]{
			Name:      "include-phone-numbers",
			Usage:     "Include the first 50 phone number objects in the results",
			Default:   true,
			QueryPath: "include_phone_numbers",
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
	Action:          handlePortingOrdersList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.activation-settings",
			InnerField: "activation_settings",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.customer-group-reference",
			Usage:      "Filter results by customer_group_reference",
			InnerField: "customer_group_reference",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.customer-reference",
			Usage:      "Filter results by customer_reference",
			InnerField: "customer_reference",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.end-user",
			InnerField: "end_user",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.misc",
			InnerField: "misc",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.parent-support-key",
			Usage:      "Filter results by parent_support_key",
			InnerField: "parent_support_key",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.phone-numbers",
			InnerField: "phone_numbers",
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

var portingOrdersDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes an existing porting order. This operation is restrict to porting orders\nin draft state.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handlePortingOrdersDelete,
	HideHelpCommand: true,
}

var portingOrdersRetrieveAllowedFocWindows = cli.Command{
	Name:    "retrieve-allowed-foc-windows",
	Usage:   "Returns a list of allowed FOC dates for a porting order.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handlePortingOrdersRetrieveAllowedFocWindows,
	HideHelpCommand: true,
}

var portingOrdersRetrieveExceptionTypes = cli.Command{
	Name:            "retrieve-exception-types",
	Usage:           "Returns a list of all possible exception types for a porting order.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handlePortingOrdersRetrieveExceptionTypes,
	HideHelpCommand: true,
}

var portingOrdersRetrieveRequirements = cli.Command{
	Name:    "retrieve-requirements",
	Usage:   "Returns a list of all requirements based on country/number type for this porting\norder.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page[size]",
		},
	},
	Action:          handlePortingOrdersRetrieveRequirements,
	HideHelpCommand: true,
}

var portingOrdersRetrieveSubRequest = cli.Command{
	Name:    "retrieve-sub-request",
	Usage:   "Retrieve the associated V1 sub_request_id and port_request_id",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handlePortingOrdersRetrieveSubRequest,
	HideHelpCommand: true,
}

func handlePortingOrdersCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortingOrderNewParams{}

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
	_, err = client.PortingOrders.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "porting-orders create", obj, format, transform)
}

func handlePortingOrdersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortingOrderGetParams{}

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
	_, err = client.PortingOrders.Get(
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
	return ShowJSON(os.Stdout, "porting-orders retrieve", obj, format, transform)
}

func handlePortingOrdersUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortingOrderUpdateParams{}

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
	_, err = client.PortingOrders.Update(
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
	return ShowJSON(os.Stdout, "porting-orders update", obj, format, transform)
}

func handlePortingOrdersList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortingOrderListParams{}

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
		_, err = client.PortingOrders.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "porting-orders list", obj, format, transform)
	} else {
		iter := client.PortingOrders.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "porting-orders list", iter, format, transform)
	}
}

func handlePortingOrdersDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.PortingOrders.Delete(ctx, cmd.Value("id").(string), options...)
}

func handlePortingOrdersRetrieveAllowedFocWindows(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.PortingOrders.GetAllowedFocWindows(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "porting-orders retrieve-allowed-foc-windows", obj, format, transform)
}

func handlePortingOrdersRetrieveExceptionTypes(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.PortingOrders.GetExceptionTypes(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "porting-orders retrieve-exception-types", obj, format, transform)
}

func handlePortingOrdersRetrieveRequirements(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortingOrderGetRequirementsParams{}

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
		_, err = client.PortingOrders.GetRequirements(
			ctx,
			cmd.Value("id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "porting-orders retrieve-requirements", obj, format, transform)
	} else {
		iter := client.PortingOrders.GetRequirementsAutoPaging(
			ctx,
			cmd.Value("id").(string),
			params,
			options...,
		)
		return ShowJSONIterator(os.Stdout, "porting-orders retrieve-requirements", iter, format, transform)
	}
}

func handlePortingOrdersRetrieveSubRequest(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.PortingOrders.GetSubRequest(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "porting-orders retrieve-sub-request", obj, format, transform)
}
