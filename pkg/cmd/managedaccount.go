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

var managedAccountsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new managed account owned by the authenticated user. You need to be\nexplictly approved by Telnyx in order to become a manager account.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "business-name",
			Usage:    "The name of the business for which the new managed account is being created, that will be used as the managed accounts's organization's name.",
			Required: true,
			BodyPath: "business_name",
		},
		&requestflag.Flag[string]{
			Name:     "email",
			Usage:    "The email address for the managed account. If not provided, the email address will be generated based on the email address of the manager account.",
			BodyPath: "email",
		},
		&requestflag.Flag[bool]{
			Name:     "managed-account-allow-custom-pricing",
			Usage:    "Boolean value that indicates if the managed account is able to have custom pricing set for it or not. If false, uses the pricing of the manager account. Defaults to false. This value may be changed after creation, but there may be time lag between when the value is changed and pricing changes take effect.",
			BodyPath: "managed_account_allow_custom_pricing",
		},
		&requestflag.Flag[string]{
			Name:     "password",
			Usage:    "Password for the managed account. If a password is not supplied, the account will not be able to be signed into directly. (A password reset may still be performed later to enable sign-in via password.)",
			BodyPath: "password",
		},
		&requestflag.Flag[bool]{
			Name:     "rollup-billing",
			Usage:    `Boolean value that indicates if the billing information and charges to the managed account "roll up" to the manager account. If true, the managed account will not have its own balance and will use the shared balance with the manager account. This value cannot be changed after account creation without going through Telnyx support as changes require manual updates to the account ledger. Defaults to false.`,
			BodyPath: "rollup_billing",
		},
	},
	Action:          handleManagedAccountsCreate,
	HideHelpCommand: true,
}

var managedAccountsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves the details of a single managed account.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleManagedAccountsRetrieve,
	HideHelpCommand: true,
}

var managedAccountsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a single managed account.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:     "managed-account-allow-custom-pricing",
			Usage:    "Boolean value that indicates if the managed account is able to have custom pricing set for it or not. If false, uses the pricing of the manager account. Defaults to false. This value may be changed, but there may be time lag between when the value is changed and pricing changes take effect.",
			BodyPath: "managed_account_allow_custom_pricing",
		},
	},
	Action:          handleManagedAccountsUpdate,
	HideHelpCommand: true,
}

var managedAccountsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Lists the accounts managed by the current user. Users need to be explictly\napproved by Telnyx in order to become manager accounts.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[email][contains], filter[email][eq], filter[organization_name][contains], filter[organization_name][eq]",
			QueryPath: "filter",
		},
		&requestflag.Flag[bool]{
			Name:      "include-cancelled-accounts",
			Usage:     "Specifies if cancelled accounts should be included in the results.",
			Default:   false,
			QueryPath: "include_cancelled_accounts",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page[size]",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the <code> -</code> prefix.<br/><br/>\nThat is: <ul>\n  <li>\n    <code>email</code>: sorts the result by the\n    <code>email</code> field in ascending order.\n  </li>\n\n  <li>\n    <code>-email</code>: sorts the result by the\n    <code>email</code> field in descending order.\n  </li>\n</ul> <br/> If not given, results are sorted by <code>created_at</code> in descending order.",
			Default:   "created_at",
			QueryPath: "sort",
		},
	},
	Action:          handleManagedAccountsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.email",
			InnerField: "email",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.organization-name",
			InnerField: "organization_name",
		},
	},
})

var managedAccountsGetAllocatableGlobalOutboundChannels = cli.Command{
	Name:            "get-allocatable-global-outbound-channels",
	Usage:           "Display information about allocatable global outbound channels for the current\nuser. Only usable by account managers.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleManagedAccountsGetAllocatableGlobalOutboundChannels,
	HideHelpCommand: true,
}

var managedAccountsUpdateGlobalChannelLimit = cli.Command{
	Name:    "update-global-channel-limit",
	Usage:   "Update the amount of allocatable global outbound channels allocated to a\nspecific managed account.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:     "channel-limit",
			Usage:    "Integer value that indicates the number of allocatable global outbound channels that should be allocated to the managed account. Must be 0 or more. If the value is 0 then the account will have no usable channels and will not be able to perform outbound calling.",
			BodyPath: "channel_limit",
		},
	},
	Action:          handleManagedAccountsUpdateGlobalChannelLimit,
	HideHelpCommand: true,
}

func handleManagedAccountsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ManagedAccountNewParams{}

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
	_, err = client.ManagedAccounts.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "managed-accounts create", obj, format, transform)
}

func handleManagedAccountsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.ManagedAccounts.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "managed-accounts retrieve", obj, format, transform)
}

func handleManagedAccountsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ManagedAccountUpdateParams{}

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
	_, err = client.ManagedAccounts.Update(
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
	return ShowJSON(os.Stdout, "managed-accounts update", obj, format, transform)
}

func handleManagedAccountsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ManagedAccountListParams{}

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
		_, err = client.ManagedAccounts.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "managed-accounts list", obj, format, transform)
	} else {
		iter := client.ManagedAccounts.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "managed-accounts list", iter, format, transform)
	}
}

func handleManagedAccountsGetAllocatableGlobalOutboundChannels(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.ManagedAccounts.GetAllocatableGlobalOutboundChannels(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "managed-accounts get-allocatable-global-outbound-channels", obj, format, transform)
}

func handleManagedAccountsUpdateGlobalChannelLimit(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ManagedAccountUpdateGlobalChannelLimitParams{}

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
	_, err = client.ManagedAccounts.UpdateGlobalChannelLimit(
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
	return ShowJSON(os.Stdout, "managed-accounts update-global-channel-limit", obj, format, transform)
}
