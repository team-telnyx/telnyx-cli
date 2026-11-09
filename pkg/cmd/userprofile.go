// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/team-telnyx/telnyx-cli/internal/apiquery"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var userProfileList = cli.Command{
	Name:            "list",
	Usage:           "Retrieve the authenticated user's profile, including data locality and migration\nstatus.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleUserProfileList,
	HideHelpCommand: true,
}

var userProfilePatchAll = cli.Command{
	Name:    "patch-all",
	Usage:   "Update the authenticated user's profile using top-level JSON fields. Returns the\nupdated profile.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*string]{
			Name:     "address-1",
			Usage:    "First line of the profile address.",
			BodyPath: "address_1",
		},
		&requestflag.Flag[*string]{
			Name:     "address-2",
			Usage:    "Second line of the profile address.",
			BodyPath: "address_2",
		},
		&requestflag.Flag[*string]{
			Name:     "business-name",
			Usage:    "Business name associated with the profile.",
			BodyPath: "business_name",
		},
		&requestflag.Flag[*string]{
			Name:     "city",
			Usage:    "City in the profile address.",
			BodyPath: "city",
		},
		&requestflag.Flag[string]{
			Name:     "company-size",
			Usage:    "Company size reported by the user.",
			BodyPath: "company_size",
		},
		&requestflag.Flag[*string]{
			Name:     "country",
			Usage:    "Country code in the profile address.",
			BodyPath: "country",
		},
		&requestflag.Flag[string]{
			Name:     "data-locality",
			Usage:    "Data locality: USA (United States), DEU (Germany), AUS (Australia), UAE (United Arab Emirates), or CAN (Canada). A change is permitted only if data locality has not previously been changed. A successful change sets data_locality_updated_at and initializes data_locality_migration_status to not_started.",
			BodyPath: "data_locality",
		},
		&requestflag.Flag[bool]{
			Name:     "default-managed-account-allow-custom-pricing",
			Usage:    "Default custom-pricing setting for managed accounts.",
			BodyPath: "default_managed_account_allow_custom_pricing",
		},
		&requestflag.Flag[bool]{
			Name:     "default-managed-account-inherit-emergency-terms-of-service",
			Usage:    "Default emergency terms-of-service inheritance setting for managed accounts, where supported.",
			BodyPath: "default_managed_account_inherit_emergency_terms_of_service",
		},
		&requestflag.Flag[bool]{
			Name:     "default-managed-account-rollup-billing",
			Usage:    "Default rollup-billing setting for managed accounts.",
			BodyPath: "default_managed_account_rollup_billing",
		},
		&requestflag.Flag[bool]{
			Name:     "default-notify-manager-account",
			Usage:    "Default setting for notifying the manager account.",
			BodyPath: "default_notify_manager_account",
		},
		&requestflag.Flag[string]{
			Name:     "default-numbers-external-pin",
			Usage:    "Default alphanumeric external PIN for phone numbers.",
			BodyPath: "default_numbers_external_pin",
		},
		&requestflag.Flag[*string]{
			Name:     "first-name",
			Usage:    "The user's first name.",
			BodyPath: "first_name",
		},
		&requestflag.Flag[bool]{
			Name:     "got-free-number-at-signup",
			Usage:    "Whether the user received a free number at signup.",
			BodyPath: "got_free_number_at_signup",
		},
		&requestflag.Flag[*string]{
			Name:     "last-name",
			Usage:    "The user's last name.",
			BodyPath: "last_name",
		},
		&requestflag.Flag[*int64]{
			Name:     "low-balance-notification-frequency",
			Usage:    "Low-balance notification frequency. Organization sub-users cannot set a non-null frequency. The resulting profile must have both a threshold and frequency, or neither.",
			BodyPath: "low_balance_notification_frequency",
		},
		&requestflag.Flag[*float64]{
			Name:     "low-balance-threshold-amount",
			Usage:    "Low-balance notification threshold in the billing currency. The resulting profile must have both a threshold and notification frequency, or neither.",
			BodyPath: "low_balance_threshold_amount",
		},
		&requestflag.Flag[bool]{
			Name:     "notify-on-invoice-available",
			Usage:    "Whether to notify the user when an invoice is available.",
			BodyPath: "notify_on_invoice_available",
		},
		&requestflag.Flag[bool]{
			Name:     "notify-on-invoice-payment",
			Usage:    "Whether to notify the user when an invoice is paid.",
			BodyPath: "notify_on_invoice_payment",
		},
		&requestflag.Flag[bool]{
			Name:     "npac-user",
			Usage:    "Whether the user is an NPAC user.",
			BodyPath: "npac_user",
		},
		&requestflag.Flag[[]string]{
			Name:     "onboarding-build-method",
			Usage:    "Build methods reported during onboarding.",
			BodyPath: "onboarding_build_method",
		},
		&requestflag.Flag[[]string]{
			Name:     "onboarding-industry",
			Usage:    "Industries reported during onboarding.",
			BodyPath: "onboarding_industries",
		},
		&requestflag.Flag[[]string]{
			Name:     "onboarding-monthly-budget",
			Usage:    "Monthly budget ranges reported during onboarding.",
			BodyPath: "onboarding_monthly_budget",
		},
		&requestflag.Flag[[]string]{
			Name:     "onboarding-product-interest",
			Usage:    "Products of interest reported during onboarding.",
			BodyPath: "onboarding_product_interest",
		},
		&requestflag.Flag[[]string]{
			Name:     "onboarding-project-type",
			Usage:    "Onboarding project types.",
			BodyPath: "onboarding_project_types",
		},
		&requestflag.Flag[[]string]{
			Name:     "onboarding-reasons-for-consideration",
			Usage:    "Reasons for considering Telnyx.",
			BodyPath: "onboarding_reasons_for_consideration",
		},
		&requestflag.Flag[[]string]{
			Name:     "onboarding-referrer",
			Usage:    "Referral sources reported during onboarding.",
			BodyPath: "onboarding_referrer",
		},
		&requestflag.Flag[[]string]{
			Name:     "onboarding-use-cases-other",
			Usage:    "Other onboarding use cases.",
			BodyPath: "onboarding_use_cases_other",
		},
		&requestflag.Flag[*string]{
			Name:     "phone-number",
			Usage:    "Contact phone number associated with the profile.",
			BodyPath: "phone_number",
		},
		&requestflag.Flag[bool]{
			Name:     "port-out-pin-validation",
			Usage:    "Whether port-out PIN validation is enabled.",
			BodyPath: "port_out_pin_validation",
		},
		&requestflag.Flag[string]{
			Name:     "primary-use-case",
			Usage:    "Primary use case for the account.",
			BodyPath: "primary_use_case",
		},
		&requestflag.Flag[string]{
			Name:     "role",
			Usage:    "The user's role.",
			BodyPath: "role",
		},
		&requestflag.Flag[*string]{
			Name:     "state",
			Usage:    "State or region in the profile address.",
			BodyPath: "state",
		},
		&requestflag.Flag[int64]{
			Name:     "ui-theme",
			Usage:    "User interface theme: 0 for dark or 1 for light.",
			BodyPath: "ui_theme",
		},
		&requestflag.Flag[[]string]{
			Name:     "use-case",
			Usage:    "Use cases reported by the user.",
			BodyPath: "use_cases",
		},
		&requestflag.Flag[bool]{
			Name:     "user-subscription-opt-in",
			Usage:    "Whether the user opts in to subscriptions.",
			BodyPath: "user_subscription_opt_in",
		},
		&requestflag.Flag[*string]{
			Name:     "zip",
			Usage:    "Postal code in the profile address.",
			BodyPath: "zip",
		},
	},
	Action:          handleUserProfilePatchAll,
	HideHelpCommand: true,
}

func handleUserProfileList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.UserProfile.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "user-profile list",
		Transform:      transform,
	})
}

func handleUserProfilePatchAll(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
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

	params := telnyx.UserProfilePatchAllParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.UserProfile.PatchAll(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "user-profile patch-all",
		Transform:      transform,
	})
}
