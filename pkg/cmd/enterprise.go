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
	Usage:   "Create the legal entity (enterprise) that represents your business on the Telnyx\nplatform.",
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
			Usage:    "The trade name your business operates under if it is different from your legal name, also called a Doing Business As (DBA) name. Leave blank if you only use your legal name.",
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
			Usage:    "The industry your business operates in. Choose the closest match from the list; if your value is not accepted, pick the nearest category.",
			Required: true,
			BodyPath: "industry",
		},
		&requestflag.Flag[string]{
			Name:     "jurisdiction-of-incorporation",
			Usage:    "The state, province, or country where your business was legally incorporated, for example Delaware.",
			Required: true,
			BodyPath: "jurisdiction_of_incorporation",
		},
		&requestflag.Flag[string]{
			Name:     "legal-name",
			Usage:    "Your business's full registered legal name, exactly as it appears on your incorporation or tax documents, 3 to 64 characters.",
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
			Usage:    "Legal-entity form. Pick the form that matches your incorporation documents:\n- `corporation` - C-corp or S-corp.\n- `llc` - limited liability company.\n- `partnership` - general/limited partnership.\n- `nonprofit` - non-profit corporation, charitable trust, or 501(c)(3)/equivalent.\n- `other` - anything else (sole proprietorships, government bodies, DBAs, etc.). You may be asked for additional documents during vetting.",
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
			Usage:    "Organization category for vetting purposes:\n- `commercial` - for-profit business entities (LLC, corp, partnership, sole proprietorship). Most callers fall here.\n- `government` - federal/state/local government bodies.\n- `non_profit` - registered 501(c)(3)/equivalent (incl. educational institutions, charities, religious organisations).",
			Required: true,
			BodyPath: "organization_type",
		},
		&requestflag.Flag[string]{
			Name:     "website",
			Usage:    "Your business's public website address, including https://. Leave blank if your business has no website.",
			Required: true,
			BodyPath: "website",
		},
		&requestflag.Flag[*string]{
			Name:     "corporate-registration-number",
			Usage:    "The official number your company received when it was legally registered or incorporated (for example from your state or national business registry). It is on your certificate of incorporation.",
			BodyPath: "corporate_registration_number",
		},
		&requestflag.Flag[string]{
			Name:     "customer-reference",
			Usage:    "Your own label for this account. Enter any reference that helps you find it in your records. Telnyx does not use it during vetting.",
			BodyPath: "customer_reference",
		},
		&requestflag.Flag[*string]{
			Name:     "dun-bradstreet-number",
			Usage:    "Your optional 9-digit D-U-N-S Number issued by Dun & Bradstreet, a unique identifier for your business. Leave blank if you do not have one.",
			BodyPath: "dun_bradstreet_number",
		},
		&requestflag.Flag[*string]{
			Name:     "primary-business-domain-sic-code",
			Usage:    "The 4-digit Standard Industrial Classification code for your main line of business, which tells us what industry you operate in. Look it up in the SIC code directory if you are unsure.",
			BodyPath: "primary_business_domain_sic_code",
		},
		&requestflag.Flag[*string]{
			Name:     "professional-license-number",
			Usage:    "If your business operates under a professional license (for example legal, medical, or financial services), enter the license number issued by the licensing authority. Leave blank if it does not apply.",
			BodyPath: "professional_license_number",
		},
		&requestflag.Flag[string]{
			Name:     "role-type",
			Usage:    "`enterprise` for an organization registering its own DIRs (the default, and the right choice when the calls display your own brand). `bpo` for a Business Process Outsourcer: a call center that places calls on behalf of other enterprises and displays their brand. A `bpo` enterprise describes the call center itself and cannot own a DIR. Each client the call center calls for gets its own `enterprise` in the same account, with the client's DIR under it; that DIR is then linked to the `bpo` enterprise through `bpo_authorizations`. Fixed at creation.",
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
			Usage:      "The city of your registered business address.",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.country",
			Usage:      "ISO 3166-1 alpha-2 code (currently `US` or `CA`).",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.postal-code",
			Usage:      "The postal or ZIP code of your registered business address.",
			InnerField: "postal_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.street-address",
			Usage:      "The street address of your registered business, including the building number and street name.",
			InnerField: "street_address",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "billing-address.extended-address",
			Usage:      "An optional second address line, such as a suite, unit, or floor. Leave blank if it does not apply.",
			InnerField: "extended_address",
		},
	},
	"billing-contact": {
		&requestflag.InnerFlag[string]{
			Name:       "billing-contact.email",
			Usage:      "The email address of the person Telnyx should contact about billing for this account.",
			InnerField: "email",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-contact.first-name",
			Usage:      "The first name of the person Telnyx should contact about billing for this account.",
			InnerField: "first_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-contact.last-name",
			Usage:      "The last name of the person Telnyx should contact about billing for this account.",
			InnerField: "last_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-contact.phone-number",
			Usage:      "The phone number of the billing contact, in E.164 format, for example +12125551234.",
			InnerField: "phone_number",
		},
	},
	"organization-contact": {
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.email",
			Usage:      "The email address of the main person Telnyx should contact about this account. For a call center (BPO) account this is the email you will verify later, so use a mailbox you can access.",
			InnerField: "email",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.first-name",
			Usage:      "The first name of the main person Telnyx should contact about this account.",
			InnerField: "first_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.job-title",
			Usage:      "The job title of the main person Telnyx should contact about this account.",
			InnerField: "job_title",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.last-name",
			Usage:      "The last name of the main person Telnyx should contact about this account.",
			InnerField: "last_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.phone-number",
			Usage:      "The phone number of the main contact, in E.164 format, for example +12125551234.",
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
			Usage:      "The city of your registered business address.",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-physical-address.country",
			Usage:      "ISO 3166-1 alpha-2 code (currently `US` or `CA`).",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-physical-address.postal-code",
			Usage:      "The postal or ZIP code of your registered business address.",
			InnerField: "postal_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-physical-address.street-address",
			Usage:      "The street address of your registered business, including the building number and street name.",
			InnerField: "street_address",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "organization-physical-address.extended-address",
			Usage:      "An optional second address line, such as a suite, unit, or floor. Leave blank if it does not apply.",
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
			Usage:    "The official number your company received when it was legally registered or incorporated (for example from your state or national business registry). It is on your certificate of incorporation.",
			BodyPath: "corporate_registration_number",
		},
		&requestflag.Flag[string]{
			Name:     "customer-reference",
			Usage:    "Your own label for this account. Enter any reference that helps you find it in your records. Telnyx does not use it during vetting.",
			BodyPath: "customer_reference",
		},
		&requestflag.Flag[string]{
			Name:     "doing-business-as",
			Usage:    "The trade name your business operates under if it is different from your legal name, also called a Doing Business As (DBA) name. Leave blank if you only use your legal name.",
			BodyPath: "doing_business_as",
		},
		&requestflag.Flag[*string]{
			Name:     "dun-bradstreet-number",
			Usage:    "Your optional 9-digit D-U-N-S Number issued by Dun & Bradstreet, a unique identifier for your business. Leave blank if you do not have one.",
			BodyPath: "dun_bradstreet_number",
		},
		&requestflag.Flag[string]{
			Name:     "fein",
			Usage:    "US Federal Employer Identification Number (`NN-NNNNNNN`) or Canadian equivalent.",
			BodyPath: "fein",
		},
		&requestflag.Flag[string]{
			Name:     "industry",
			Usage:    "The industry your business operates in. Choose the closest match from the list; if your value is not accepted, pick the nearest category.",
			BodyPath: "industry",
		},
		&requestflag.Flag[string]{
			Name:     "jurisdiction-of-incorporation",
			Usage:    "The state, province, or country where your business was legally incorporated, for example Delaware.",
			BodyPath: "jurisdiction_of_incorporation",
		},
		&requestflag.Flag[string]{
			Name:     "legal-name",
			Usage:    "Your business's full registered legal name, exactly as it appears on your incorporation or tax documents, 3 to 64 characters.",
			BodyPath: "legal_name",
		},
		&requestflag.Flag[string]{
			Name:     "number-of-employees",
			Usage:    "Approximate headcount range. Used for vetting heuristics; pick the bucket that contains your current employee count.",
			BodyPath: "number_of_employees",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "organization-contact",
			BodyPath: "organization_contact",
		},
		&requestflag.Flag[string]{
			Name:     "organization-legal-type",
			Usage:    "Legal-entity form. Pick the form that matches your incorporation documents:\n- `corporation` - C-corp or S-corp.\n- `llc` - limited liability company.\n- `partnership` - general/limited partnership.\n- `nonprofit` - non-profit corporation, charitable trust, or 501(c)(3)/equivalent.\n- `other` - anything else (sole proprietorships, government bodies, DBAs, etc.). You may be asked for additional documents during vetting.",
			BodyPath: "organization_legal_type",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "organization-physical-address",
			BodyPath: "organization_physical_address",
		},
		&requestflag.Flag[*string]{
			Name:     "primary-business-domain-sic-code",
			Usage:    "The 4-digit Standard Industrial Classification code for your main line of business, which tells us what industry you operate in. Look it up in the SIC code directory if you are unsure.",
			BodyPath: "primary_business_domain_sic_code",
		},
		&requestflag.Flag[*string]{
			Name:     "professional-license-number",
			Usage:    "If your business operates under a professional license (for example legal, medical, or financial services), enter the license number issued by the licensing authority. Leave blank if it does not apply.",
			BodyPath: "professional_license_number",
		},
		&requestflag.Flag[string]{
			Name:     "website",
			Usage:    "Your business's public website address, including https://. Leave blank if your business has no website.",
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
			Usage:      "The city of your registered business address.",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.country",
			Usage:      "ISO 3166-1 alpha-2 code (currently `US` or `CA`).",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.postal-code",
			Usage:      "The postal or ZIP code of your registered business address.",
			InnerField: "postal_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.street-address",
			Usage:      "The street address of your registered business, including the building number and street name.",
			InnerField: "street_address",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "billing-address.extended-address",
			Usage:      "An optional second address line, such as a suite, unit, or floor. Leave blank if it does not apply.",
			InnerField: "extended_address",
		},
	},
	"billing-contact": {
		&requestflag.InnerFlag[string]{
			Name:       "billing-contact.email",
			Usage:      "The email address of the person Telnyx should contact about billing for this account.",
			InnerField: "email",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-contact.first-name",
			Usage:      "The first name of the person Telnyx should contact about billing for this account.",
			InnerField: "first_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-contact.last-name",
			Usage:      "The last name of the person Telnyx should contact about billing for this account.",
			InnerField: "last_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-contact.phone-number",
			Usage:      "The phone number of the billing contact, in E.164 format, for example +12125551234.",
			InnerField: "phone_number",
		},
	},
	"organization-contact": {
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.email",
			Usage:      "The email address of the main person Telnyx should contact about this account. For a call center (BPO) account this is the email you will verify later, so use a mailbox you can access.",
			InnerField: "email",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.first-name",
			Usage:      "The first name of the main person Telnyx should contact about this account.",
			InnerField: "first_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.job-title",
			Usage:      "The job title of the main person Telnyx should contact about this account.",
			InnerField: "job_title",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.last-name",
			Usage:      "The last name of the main person Telnyx should contact about this account.",
			InnerField: "last_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-contact.phone-number",
			Usage:      "The phone number of the main contact, in E.164 format, for example +12125551234.",
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
			Usage:      "The city of your registered business address.",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-physical-address.country",
			Usage:      "ISO 3166-1 alpha-2 code (currently `US` or `CA`).",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-physical-address.postal-code",
			Usage:      "The postal or ZIP code of your registered business address.",
			InnerField: "postal_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "organization-physical-address.street-address",
			Usage:      "The street address of your registered business, including the building number and street name.",
			InnerField: "street_address",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "organization-physical-address.extended-address",
			Usage:      "An optional second address line, such as a suite, unit, or floor. Leave blank if it does not apply.",
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
			Name:      "filter-legal-name-contains",
			Usage:     "Case-insensitive partial match on legal name.",
			QueryPath: "filter[legal_name][contains]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-role-type",
			Usage:     "Only return enterprises of this type: `bpo` for call-center (BPO) enterprises, `enterprise` for normal enterprises. Omit to return both.",
			QueryPath: "filter[role_type]",
		},
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
	Usage:   "Soft-delete an enterprise.",
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

var enterprisesBrandedCalling = cli.Command{
	Name:    "branded-calling",
	Usage:   "Branded Calling must be activated on each enterprise. Activation is idempotent:",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "enterprise-id",
			Required:  true,
			PathParam: "enterprise_id",
		},
	},
	Action:          handleEnterprisesBrandedCalling,
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

func handleEnterprisesBrandedCalling(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Enterprises.BrandedCalling(ctx, cmd.Value("enterprise-id").(string), options...)
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
		Title:          "enterprises branded-calling",
		Transform:      transform,
	})
}
