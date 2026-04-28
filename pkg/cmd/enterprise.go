// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/team-telnyx/telnyx-cli/internal/apiquery"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var enterprisesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a new enterprise for Branded Calling / Number Reputation services.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "billing-address",
			Required: true,
			BodyPath: "billing_address",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "billing-contact",
			Required: true,
			BodyPath: "billing_contact",
		},
		&requestflag.Flag[string]{
			Name:     "country-code",
			Usage:    "Country code. Currently only 'US' is accepted.",
			Required: true,
			BodyPath: "country_code",
		},
		&requestflag.Flag[string]{
			Name:     "doing-business-as",
			Usage:    "Primary business name / DBA name",
			Required: true,
			BodyPath: "doing_business_as",
		},
		&requestflag.Flag[string]{
			Name:     "fein",
			Usage:    "Federal Employer Identification Number. Format: XX-XXXXXXX or 9-digit number (minimum 9 digits).",
			Required: true,
			BodyPath: "fein",
		},
		&requestflag.Flag[string]{
			Name:     "industry",
			Usage:    "Industry classification. Case-insensitive. Accepted values: accounting, finance, billing, collections, business, charity, nonprofit, communications, telecom, customer service, support, delivery, shipping, logistics, education, financial, banking, government, public, healthcare, health, pharmacy, medical, insurance, legal, law, notifications, scheduling, real estate, property, retail, ecommerce, sales, marketing, software, technology, tech, media, surveys, market research, travel, hospitality, hotel",
			Required: true,
			BodyPath: "industry",
		},
		&requestflag.Flag[string]{
			Name:     "legal-name",
			Usage:    "Legal name of the enterprise",
			Required: true,
			BodyPath: "legal_name",
		},
		&requestflag.Flag[string]{
			Name:     "number-of-employees",
			Usage:    "Employee count range",
			Required: true,
			BodyPath: "number_of_employees",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "organization-contact",
			Usage:    "Organization contact information. Note: the response returns this object with the phone field as 'phone' (not 'phone_number').",
			Required: true,
			BodyPath: "organization_contact",
		},
		&requestflag.Flag[string]{
			Name:     "organization-legal-type",
			Usage:    "Legal structure type",
			Required: true,
			BodyPath: "organization_legal_type",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "organization-physical-address",
			Required: true,
			BodyPath: "organization_physical_address",
		},
		&requestflag.Flag[string]{
			Name:     "organization-type",
			Usage:    "Type of organization",
			Required: true,
			BodyPath: "organization_type",
		},
		&requestflag.Flag[string]{
			Name:     "website",
			Usage:    "Enterprise website URL. Accepts any string — no URL format validation enforced.",
			Required: true,
			BodyPath: "website",
		},
		&requestflag.Flag[string]{
			Name:     "corporate-registration-number",
			Usage:    "Corporate registration number (optional)",
			BodyPath: "corporate_registration_number",
		},
		&requestflag.Flag[string]{
			Name:     "customer-reference",
			Usage:    "Optional customer reference identifier for your own tracking",
			BodyPath: "customer_reference",
		},
		&requestflag.Flag[string]{
			Name:     "dun-bradstreet-number",
			Usage:    "D-U-N-S Number (optional)",
			BodyPath: "dun_bradstreet_number",
		},
		&requestflag.Flag[string]{
			Name:     "primary-business-domain-sic-code",
			Usage:    "SIC Code (optional)",
			BodyPath: "primary_business_domain_sic_code",
		},
		&requestflag.Flag[string]{
			Name:     "professional-license-number",
			Usage:    "Professional license number (optional)",
			BodyPath: "professional_license_number",
		},
		&requestflag.Flag[string]{
			Name:     "role-type",
			Usage:    "Role type in Branded Calling / Number Reputation services",
			Default:  "enterprise",
			BodyPath: "role_type",
		},
	},
	Action:          handleEnterprisesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"billing-address": {
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.administrative-area",
			Usage:      "State or province",
			InnerField: "administrative_area",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.city",
			Usage:      "City name",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.country",
			Usage:      "Country name (e.g., United States)",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.postal-code",
			Usage:      "ZIP or postal code",
			InnerField: "postal_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.street-address",
			Usage:      "Street address",
			InnerField: "street_address",
		},
		&requestflag.InnerFlag[any]{
			Name:       "billing-address.extended-address",
			Usage:      "Additional address line (suite, apt, etc.)",
			InnerField: "extended_address",
		},
	},
	"billing-contact": {
		&requestflag.InnerFlag[string]{
			Name:       "billing-contact.email",
			Usage:      "Contact's email address",
			InnerField: "email",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-contact.first-name",
			Usage:      "Contact's first name",
			InnerField: "first_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-contact.last-name",
			Usage:      "Contact's last name",
			InnerField: "last_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-contact.phone-number",
			Usage:      "Contact's phone number (10-15 digits)",
			InnerField: "phone_number",
		},
	},
	"organization-contact": {
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.email",
			Usage:      "Contact's email address",
			InnerField: "email",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.first-name",
			Usage:      "Contact's first name",
			InnerField: "first_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.job-title",
			Usage:      "Contact's job title (required)",
			InnerField: "job_title",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.last-name",
			Usage:      "Contact's last name",
			InnerField: "last_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.phone",
			Usage:      "Contact's phone number in E.164 format",
			InnerField: "phone",
		},
	},
	"organization-physical-address": {
		&requestflag.InnerFlag[string]{
			Name:       "organization-physical-address.administrative-area",
			Usage:      "State or province",
			InnerField: "administrative_area",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-physical-address.city",
			Usage:      "City name",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-physical-address.country",
			Usage:      "Country name (e.g., United States)",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-physical-address.postal-code",
			Usage:      "ZIP or postal code",
			InnerField: "postal_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-physical-address.street-address",
			Usage:      "Street address",
			InnerField: "street_address",
		},
		&requestflag.InnerFlag[any]{
			Name:       "organization-physical-address.extended-address",
			Usage:      "Additional address line (suite, apt, etc.)",
			InnerField: "extended_address",
		},
	},
})

var enterprisesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve details of a specific enterprise by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "enterprise-id",
			Required: true,
		},
	},
	Action:          handleEnterprisesRetrieve,
	HideHelpCommand: true,
}

var enterprisesUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update enterprise information. All fields are optional — only the provided\nfields will be updated.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "enterprise-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "billing-address",
			BodyPath: "billing_address",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "billing-contact",
			BodyPath: "billing_contact",
		},
		&requestflag.Flag[string]{
			Name:     "corporate-registration-number",
			Usage:    "Corporate registration number",
			BodyPath: "corporate_registration_number",
		},
		&requestflag.Flag[string]{
			Name:     "customer-reference",
			Usage:    "Customer reference identifier",
			BodyPath: "customer_reference",
		},
		&requestflag.Flag[string]{
			Name:     "doing-business-as",
			Usage:    "DBA name",
			BodyPath: "doing_business_as",
		},
		&requestflag.Flag[string]{
			Name:     "dun-bradstreet-number",
			Usage:    "D-U-N-S Number",
			BodyPath: "dun_bradstreet_number",
		},
		&requestflag.Flag[string]{
			Name:     "fein",
			Usage:    "Federal Employer Identification Number. Format: XX-XXXXXXX or XXXXXXXXX",
			BodyPath: "fein",
		},
		&requestflag.Flag[string]{
			Name:     "industry",
			Usage:    "Industry classification",
			BodyPath: "industry",
		},
		&requestflag.Flag[string]{
			Name:     "legal-name",
			Usage:    "Legal name of the enterprise",
			BodyPath: "legal_name",
		},
		&requestflag.Flag[string]{
			Name:     "number-of-employees",
			Usage:    "Employee count range",
			BodyPath: "number_of_employees",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "organization-contact",
			Usage:    "Organization contact information. Note: the response returns this object with the phone field as 'phone' (not 'phone_number').",
			BodyPath: "organization_contact",
		},
		&requestflag.Flag[string]{
			Name:     "organization-legal-type",
			Usage:    "Legal structure type",
			BodyPath: "organization_legal_type",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "organization-physical-address",
			BodyPath: "organization_physical_address",
		},
		&requestflag.Flag[string]{
			Name:     "primary-business-domain-sic-code",
			Usage:    "SIC Code",
			BodyPath: "primary_business_domain_sic_code",
		},
		&requestflag.Flag[string]{
			Name:     "professional-license-number",
			Usage:    "Professional license number",
			BodyPath: "professional_license_number",
		},
		&requestflag.Flag[string]{
			Name:     "website",
			Usage:    "Company website URL",
			BodyPath: "website",
		},
	},
	Action:          handleEnterprisesUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"billing-address": {
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.administrative-area",
			Usage:      "State or province",
			InnerField: "administrative_area",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.city",
			Usage:      "City name",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.country",
			Usage:      "Country name (e.g., United States)",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.postal-code",
			Usage:      "ZIP or postal code",
			InnerField: "postal_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.street-address",
			Usage:      "Street address",
			InnerField: "street_address",
		},
		&requestflag.InnerFlag[any]{
			Name:       "billing-address.extended-address",
			Usage:      "Additional address line (suite, apt, etc.)",
			InnerField: "extended_address",
		},
	},
	"billing-contact": {
		&requestflag.InnerFlag[string]{
			Name:       "billing-contact.email",
			Usage:      "Contact's email address",
			InnerField: "email",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-contact.first-name",
			Usage:      "Contact's first name",
			InnerField: "first_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-contact.last-name",
			Usage:      "Contact's last name",
			InnerField: "last_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-contact.phone-number",
			Usage:      "Contact's phone number (10-15 digits)",
			InnerField: "phone_number",
		},
	},
	"organization-contact": {
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.email",
			Usage:      "Contact's email address",
			InnerField: "email",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.first-name",
			Usage:      "Contact's first name",
			InnerField: "first_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.job-title",
			Usage:      "Contact's job title (required)",
			InnerField: "job_title",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.last-name",
			Usage:      "Contact's last name",
			InnerField: "last_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.phone",
			Usage:      "Contact's phone number in E.164 format",
			InnerField: "phone",
		},
	},
	"organization-physical-address": {
		&requestflag.InnerFlag[string]{
			Name:       "organization-physical-address.administrative-area",
			Usage:      "State or province",
			InnerField: "administrative_area",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-physical-address.city",
			Usage:      "City name",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-physical-address.country",
			Usage:      "Country name (e.g., United States)",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-physical-address.postal-code",
			Usage:      "ZIP or postal code",
			InnerField: "postal_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-physical-address.street-address",
			Usage:      "Street address",
			InnerField: "street_address",
		},
		&requestflag.InnerFlag[any]{
			Name:       "organization-physical-address.extended-address",
			Usage:      "Additional address line (suite, apt, etc.)",
			InnerField: "extended_address",
		},
	},
})

var enterprisesList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve a paginated list of enterprises associated with your account.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "legal-name",
			Usage:     "Filter by legal name (partial match)",
			QueryPath: "legal_name",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "Page number (1-indexed)",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Number of items per page",
			Default:   10,
			QueryPath: "page[size]",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleEnterprisesList,
	HideHelpCommand: true,
}

var enterprisesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete an enterprise and all associated resources. This action is irreversible.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "enterprise-id",
			Required: true,
		},
	},
	Action:          handleEnterprisesDelete,
	HideHelpCommand: true,
}

func handleEnterprisesCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.EnterpriseNewParams{}

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
	_, err = client.Enterprises.New(ctx, params, options...)
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
		Title:          "enterprises create",
		Transform:      transform,
	})
}

func handleEnterprisesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("enterprise-id") && len(unusedArgs) > 0 {
		cmd.Set("enterprise-id", unusedArgs[0])
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
	_, err = client.Enterprises.Get(ctx, cmd.Value("enterprise-id").(string), options...)
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
		Title:          "enterprises retrieve",
		Transform:      transform,
	})
}

func handleEnterprisesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("enterprise-id") && len(unusedArgs) > 0 {
		cmd.Set("enterprise-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.EnterpriseUpdateParams{}

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
	_, err = client.Enterprises.Update(
		ctx,
		cmd.Value("enterprise-id").(string),
		params,
		options...,
	)
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
		Title:          "enterprises update",
		Transform:      transform,
	})
}

func handleEnterprisesList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.EnterpriseListParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Enterprises.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "enterprises list",
			Transform:      transform,
		})
	} else {
		iter := client.Enterprises.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "enterprises list",
			Transform:      transform,
		})
	}
}

func handleEnterprisesDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("enterprise-id") && len(unusedArgs) > 0 {
		cmd.Set("enterprise-id", unusedArgs[0])
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

	return client.Enterprises.Delete(ctx, cmd.Value("enterprise-id").(string), options...)
}
