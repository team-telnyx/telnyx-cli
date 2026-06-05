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

var enterprisesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create the legal entity that owns your Number Reputation registrations.",
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
			Usage:    "ISO 3166-1 alpha-2 country code. Currently `US` and `CA` are supported.",
			Required: true,
			BodyPath: "country_code",
		},
		&requestflag.Flag[string]{
			Name:     "doing-business-as",
			Required: true,
			BodyPath: "doing_business_as",
		},
		&requestflag.Flag[string]{
			Name:     "fein",
			Usage:    "US Federal Employer Identification Number (`NN-NNNNNNN`) or Canadian equivalent.",
			Required: true,
			BodyPath: "fein",
		},
		&requestflag.Flag[string]{
			Name:     "industry",
			Usage:    "Industry classification.",
			Required: true,
			BodyPath: "industry",
		},
		&requestflag.Flag[string]{
			Name:     "jurisdiction-of-incorporation",
			Required: true,
			BodyPath: "jurisdiction_of_incorporation",
		},
		&requestflag.Flag[string]{
			Name:     "legal-name",
			Usage:    "Legal name of the enterprise.",
			Required: true,
			BodyPath: "legal_name",
		},
		&requestflag.Flag[string]{
			Name:     "number-of-employees",
			Usage:    "Approximate headcount range. Used for vetting heuristics; pick the bucket that contains your current employee count.",
			Required: true,
			BodyPath: "number_of_employees",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "organization-contact",
			Required: true,
			BodyPath: "organization_contact",
		},
		&requestflag.Flag[string]{
			Name:     "organization-legal-type",
			Usage:    "Legal-entity form. Pick the form that matches your incorporation documents:\n- `corporation` — C-corp or S-corp.\n- `llc` — limited liability company.\n- `partnership` — general/limited partnership.\n- `nonprofit` — non-profit corporation, charitable trust, or 501(c)(3)/equivalent.\n- `other` — anything else (sole proprietorships, government bodies, DBAs, etc.). You may be asked for additional documents during vetting.",
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
			Usage:    "Organization category for vetting purposes:\n- `commercial` — for-profit business entities (LLC, corp, partnership, sole proprietorship). Most callers fall here.\n- `government` — federal/state/local government bodies.\n- `non_profit` — registered 501(c)(3)/equivalent (incl. educational institutions, charities, religious organisations).",
			Required: true,
			BodyPath: "organization_type",
		},
		&requestflag.Flag[string]{
			Name:     "website",
			Required: true,
			BodyPath: "website",
		},
		&requestflag.Flag[*string]{
			Name:     "corporate-registration-number",
			Usage:    "Optional corporate-registration / company-number identifier.",
			BodyPath: "corporate_registration_number",
		},
		&requestflag.Flag[string]{
			Name:     "customer-reference",
			Usage:    "Optional free-form string the caller can attach for their own bookkeeping. Telnyx does not interpret it.",
			BodyPath: "customer_reference",
		},
		&requestflag.Flag[*string]{
			Name:     "dun-bradstreet-number",
			Usage:    "Optional D-U-N-S Number.",
			BodyPath: "dun_bradstreet_number",
		},
		&requestflag.Flag[*string]{
			Name:     "primary-business-domain-sic-code",
			Usage:    "Optional SIC code for the primary line of business.",
			BodyPath: "primary_business_domain_sic_code",
		},
		&requestflag.Flag[*string]{
			Name:     "professional-license-number",
			Usage:    "Optional professional-license number for regulated industries.",
			BodyPath: "professional_license_number",
		},
		&requestflag.Flag[string]{
			Name:     "role-type",
			Usage:    "`enterprise` for an organization registering its own DIRs; `bpo` for a Business Process Outsourcer placing calls on behalf of one or more enterprises.",
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
			Usage:      "State or province code (e.g. `IL`, `ON`).",
			InnerField: "administrative_area",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.city",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.country",
			Usage:      "ISO 3166-1 alpha-2 code (currently `US` or `CA`).",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.postal-code",
			InnerField: "postal_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.street-address",
			InnerField: "street_address",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "billing-address.extended-address",
			InnerField: "extended_address",
		},
	},
	"billing-contact": {
		&requestflag.InnerFlag[string]{
			Name:       "billing-contact.email",
			InnerField: "email",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-contact.first-name",
			InnerField: "first_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-contact.last-name",
			InnerField: "last_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-contact.phone-number",
			Usage:      "E.164 format with leading `+`.",
			InnerField: "phone_number",
		},
	},
	"organization-contact": {
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.email",
			InnerField: "email",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.first-name",
			InnerField: "first_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.job-title",
			InnerField: "job_title",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.last-name",
			InnerField: "last_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.phone-number",
			Usage:      "E.164 format with leading `+`.",
			InnerField: "phone_number",
		},
	},
	"organization-physical-address": {
		&requestflag.InnerFlag[string]{
			Name:       "organization-physical-address.administrative-area",
			Usage:      "State or province code (e.g. `IL`, `ON`).",
			InnerField: "administrative_area",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-physical-address.city",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-physical-address.country",
			Usage:      "ISO 3166-1 alpha-2 code (currently `US` or `CA`).",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-physical-address.postal-code",
			InnerField: "postal_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-physical-address.street-address",
			InnerField: "street_address",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "organization-physical-address.extended-address",
			InnerField: "extended_address",
		},
	},
})

var enterprisesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a single enterprise by id. Returns `404` if the id does not exist or\ndoes not belong to your account.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "enterprise-id",
			Required:  true,
			PathParam: "enterprise_id",
		},
	},
	Action:          handleEnterprisesRetrieve,
	HideHelpCommand: true,
}

var enterprisesUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Replace the enterprise's mutable fields. Only mutable fields may be sent.\nServer-assigned and immutable fields (`id`, `record_type`, `created_at`,\n`updated_at`, status fields, `organization_type`, `country_code`, `role_type`)\ncannot be changed: including any of them in the body is rejected with\n`400 Bad Request` (`Field 'X' is not allowed in this request`).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "enterprise-id",
			Required:  true,
			PathParam: "enterprise_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "billing-address",
			BodyPath: "billing_address",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "billing-contact",
			BodyPath: "billing_contact",
		},
		&requestflag.Flag[*string]{
			Name:     "corporate-registration-number",
			BodyPath: "corporate_registration_number",
		},
		&requestflag.Flag[string]{
			Name:     "customer-reference",
			BodyPath: "customer_reference",
		},
		&requestflag.Flag[string]{
			Name:     "doing-business-as",
			BodyPath: "doing_business_as",
		},
		&requestflag.Flag[*string]{
			Name:     "dun-bradstreet-number",
			BodyPath: "dun_bradstreet_number",
		},
		&requestflag.Flag[string]{
			Name:     "fein",
			BodyPath: "fein",
		},
		&requestflag.Flag[string]{
			Name:     "industry",
			Usage:    `Allowed values: "accounting", "finance", "billing", "collections", "business", "charity", "nonprofit", "communications", "telecom", "customer service", "support", "delivery", "shipping", "logistics", "education", "financial", "banking", "government", "public", "healthcare", "health", "pharmacy", "medical", "insurance", "legal", "law", "notifications", "scheduling", "real estate", "property", "retail", "ecommerce", "sales", "marketing", "software", "technology", "tech", "media", "surveys", "market research", "travel", "hospitality", "hotel".`,
			BodyPath: "industry",
		},
		&requestflag.Flag[string]{
			Name:     "jurisdiction-of-incorporation",
			Usage:    "Updated state/province/country of incorporation. Optional on update.",
			BodyPath: "jurisdiction_of_incorporation",
		},
		&requestflag.Flag[string]{
			Name:     "legal-name",
			Usage:    "Legal name of the enterprise.",
			BodyPath: "legal_name",
		},
		&requestflag.Flag[string]{
			Name:     "number-of-employees",
			BodyPath: "number_of_employees",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "organization-contact",
			BodyPath: "organization_contact",
		},
		&requestflag.Flag[string]{
			Name:     "organization-legal-type",
			BodyPath: "organization_legal_type",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "organization-physical-address",
			BodyPath: "organization_physical_address",
		},
		&requestflag.Flag[*string]{
			Name:     "primary-business-domain-sic-code",
			BodyPath: "primary_business_domain_sic_code",
		},
		&requestflag.Flag[*string]{
			Name:     "professional-license-number",
			BodyPath: "professional_license_number",
		},
		&requestflag.Flag[string]{
			Name:     "website",
			BodyPath: "website",
		},
	},
	Action:          handleEnterprisesUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"billing-address": {
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.administrative-area",
			Usage:      "State or province code (e.g. `IL`, `ON`).",
			InnerField: "administrative_area",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.city",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.country",
			Usage:      "ISO 3166-1 alpha-2 code (currently `US` or `CA`).",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.postal-code",
			InnerField: "postal_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.street-address",
			InnerField: "street_address",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "billing-address.extended-address",
			InnerField: "extended_address",
		},
	},
	"billing-contact": {
		&requestflag.InnerFlag[string]{
			Name:       "billing-contact.email",
			InnerField: "email",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-contact.first-name",
			InnerField: "first_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-contact.last-name",
			InnerField: "last_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-contact.phone-number",
			Usage:      "E.164 format with leading `+`.",
			InnerField: "phone_number",
		},
	},
	"organization-contact": {
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.email",
			InnerField: "email",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.first-name",
			InnerField: "first_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.job-title",
			InnerField: "job_title",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.last-name",
			InnerField: "last_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.phone-number",
			Usage:      "E.164 format with leading `+`.",
			InnerField: "phone_number",
		},
	},
	"organization-physical-address": {
		&requestflag.InnerFlag[string]{
			Name:       "organization-physical-address.administrative-area",
			Usage:      "State or province code (e.g. `IL`, `ON`).",
			InnerField: "administrative_area",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-physical-address.city",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-physical-address.country",
			Usage:      "ISO 3166-1 alpha-2 code (currently `US` or `CA`).",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-physical-address.postal-code",
			InnerField: "postal_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-physical-address.street-address",
			InnerField: "street_address",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "organization-physical-address.extended-address",
			InnerField: "extended_address",
		},
	},
})

var enterprisesList = cli.Command{
	Name:    "list",
	Usage:   "Return the enterprises you own, paginated. The default page size is 20; the\nmaximum is 250.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "legal-name",
			Usage:     "Filter by legal name (partial match).",
			QueryPath: "legal_name",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "1-based page number. Out-of-range values return an empty page with correct meta.",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Items per page. Default 10. Maximum 250; values above are clamped to 250.",
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
	Usage:   "Delete an enterprise. Fails with `400` if the enterprise still has dependent\nresources (e.g. active reputation settings or registered numbers); remove those\nfirst. Returns `404` if the enterprise does not exist or does not belong to your\naccount.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "enterprise-id",
			Required:  true,
			PathParam: "enterprise_id",
		},
	},
	Action:          handleEnterprisesDelete,
	HideHelpCommand: true,
}

var enterprisesActivateBrandedCalling = cli.Command{
	Name:    "activate-branded-calling",
	Usage:   "Branded Calling is a paid product that must be activated on each enterprise.\nActivation is idempotent:",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "enterprise-id",
			Required:  true,
			PathParam: "enterprise_id",
		},
	},
	Action:          handleEnterprisesActivateBrandedCalling,
	HideHelpCommand: true,
}

func handleEnterprisesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EnterpriseNewParams{}

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

	params := telnyx.EnterpriseUpdateParams{}

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

	params := telnyx.EnterpriseListParams{}

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

func handleEnterprisesActivateBrandedCalling(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Enterprises.ActivateBrandedCalling(ctx, cmd.Value("enterprise-id").(string), options...)
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
		Title:          "enterprises activate-branded-calling",
		Transform:      transform,
	})
}
