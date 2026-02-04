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

var messaging10dlcBrandCreate = cli.Command{
	Name:    "create",
	Usage:   "This endpoint is used to create a new brand. A brand is an entity created by The\nCampaign Registry (TCR) that represents an organization or a company. It is this\nentity that TCR created campaigns will be associated with. Each brand creation\nwill entail an upfront, non-refundable $4 expense.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "country",
			Usage:    "ISO2 2 characters country code. Example: US - United States",
			Required: true,
			BodyPath: "country",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			Usage:    "Display name, marketing name, or DBA name of the brand.",
			Required: true,
			BodyPath: "displayName",
		},
		&requestflag.Flag[string]{
			Name:     "email",
			Usage:    "Valid email address of brand support contact.",
			Required: true,
			BodyPath: "email",
		},
		&requestflag.Flag[string]{
			Name:     "entity-type",
			Usage:    "Entity type behind the brand. This is the form of business establishment.",
			Required: true,
			BodyPath: "entityType",
		},
		&requestflag.Flag[string]{
			Name:     "vertical",
			Usage:    "Vertical or industry segment of the brand or campaign.",
			Required: true,
			BodyPath: "vertical",
		},
		&requestflag.Flag[string]{
			Name:     "business-contact-email",
			Usage:    "Business contact email.\n\nRequired if `entityType` is `PUBLIC_PROFIT`. Otherwise, it is recommended to either omit this field or set it to `null`.",
			BodyPath: "businessContactEmail",
		},
		&requestflag.Flag[string]{
			Name:     "city",
			Usage:    "City name",
			BodyPath: "city",
		},
		&requestflag.Flag[string]{
			Name:     "company-name",
			Usage:    "(Required for Non-profit/private/public) Legal company name.",
			BodyPath: "companyName",
		},
		&requestflag.Flag[string]{
			Name:     "ein",
			Usage:    "(Required for Non-profit) Government assigned corporate tax ID. EIN is 9-digits in U.S.",
			BodyPath: "ein",
		},
		&requestflag.Flag[string]{
			Name:     "first-name",
			Usage:    "First name of business contact.",
			BodyPath: "firstName",
		},
		&requestflag.Flag[string]{
			Name:     "ip-address",
			Usage:    "IP address of the browser requesting to create brand identity.",
			BodyPath: "ipAddress",
		},
		&requestflag.Flag[bool]{
			Name:     "is-reseller",
			Default:  false,
			BodyPath: "isReseller",
		},
		&requestflag.Flag[string]{
			Name:     "last-name",
			Usage:    "Last name of business contact.",
			BodyPath: "lastName",
		},
		&requestflag.Flag[string]{
			Name:     "mobile-phone",
			Usage:    "Valid mobile phone number in e.164 international format.",
			BodyPath: "mobilePhone",
		},
		&requestflag.Flag[bool]{
			Name:     "mock",
			Usage:    "Mock brand for testing purposes. Defaults to false.",
			Default:  false,
			BodyPath: "mock",
		},
		&requestflag.Flag[string]{
			Name:     "phone",
			Usage:    "Valid phone number in e.164 international format.",
			BodyPath: "phone",
		},
		&requestflag.Flag[string]{
			Name:     "postal-code",
			Usage:    "Postal codes. Use 5 digit zipcode for United States",
			BodyPath: "postalCode",
		},
		&requestflag.Flag[string]{
			Name:     "state",
			Usage:    "State. Must be 2 letters code for United States.",
			BodyPath: "state",
		},
		&requestflag.Flag[string]{
			Name:     "stock-exchange",
			Usage:    "(Required for public company) stock exchange.",
			BodyPath: "stockExchange",
		},
		&requestflag.Flag[string]{
			Name:     "stock-symbol",
			Usage:    "(Required for public company) stock symbol.",
			BodyPath: "stockSymbol",
		},
		&requestflag.Flag[string]{
			Name:     "street",
			Usage:    "Street number and name.",
			BodyPath: "street",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-failover-url",
			Usage:    "Webhook failover URL for brand status updates.",
			BodyPath: "webhookFailoverURL",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "Webhook URL for brand status updates.",
			BodyPath: "webhookURL",
		},
		&requestflag.Flag[string]{
			Name:     "website",
			Usage:    "Brand website URL.",
			BodyPath: "website",
		},
	},
	Action:          handleMessaging10dlcBrandCreate,
	HideHelpCommand: true,
}

var messaging10dlcBrandRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a brand by `brandId`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "brand-id",
			Required: true,
		},
	},
	Action:          handleMessaging10dlcBrandRetrieve,
	HideHelpCommand: true,
}

var messaging10dlcBrandUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a brand's attributes by `brandId`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "brand-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "country",
			Usage:    "ISO2 2 characters country code. Example: US - United States",
			Required: true,
			BodyPath: "country",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			Usage:    "Display or marketing name of the brand.",
			Required: true,
			BodyPath: "displayName",
		},
		&requestflag.Flag[string]{
			Name:     "email",
			Usage:    "Valid email address of brand support contact.",
			Required: true,
			BodyPath: "email",
		},
		&requestflag.Flag[string]{
			Name:     "entity-type",
			Usage:    "Entity type behind the brand. This is the form of business establishment.",
			Required: true,
			BodyPath: "entityType",
		},
		&requestflag.Flag[string]{
			Name:     "vertical",
			Usage:    "Vertical or industry segment of the brand or campaign.",
			Required: true,
			BodyPath: "vertical",
		},
		&requestflag.Flag[string]{
			Name:     "alt-business-id",
			Usage:    "Alternate business identifier such as DUNS, LEI, or GIIN",
			BodyPath: "altBusiness_id",
		},
		&requestflag.Flag[string]{
			Name:     "alt-business-id-type",
			Usage:    "An enumeration.",
			BodyPath: "altBusinessIdType",
		},
		&requestflag.Flag[string]{
			Name:     "business-contact-email",
			Usage:    "Business contact email.\n\nRequired if `entityType` will be changed to `PUBLIC_PROFIT`. Otherwise, it is recommended to either omit this field or set it to `null`.",
			BodyPath: "businessContactEmail",
		},
		&requestflag.Flag[string]{
			Name:     "city",
			Usage:    "City name",
			BodyPath: "city",
		},
		&requestflag.Flag[string]{
			Name:     "company-name",
			Usage:    "(Required for Non-profit/private/public) Legal company name.",
			BodyPath: "companyName",
		},
		&requestflag.Flag[string]{
			Name:     "ein",
			Usage:    "(Required for Non-profit) Government assigned corporate tax ID. EIN is 9-digits in U.S.",
			BodyPath: "ein",
		},
		&requestflag.Flag[string]{
			Name:     "first-name",
			Usage:    "First name of business contact.",
			BodyPath: "firstName",
		},
		&requestflag.Flag[string]{
			Name:     "identity-status",
			Usage:    "The verification status of an active brand",
			BodyPath: "identityStatus",
		},
		&requestflag.Flag[string]{
			Name:     "ip-address",
			Usage:    "IP address of the browser requesting to create brand identity.",
			BodyPath: "ipAddress",
		},
		&requestflag.Flag[bool]{
			Name:     "is-reseller",
			BodyPath: "isReseller",
		},
		&requestflag.Flag[string]{
			Name:     "last-name",
			Usage:    "Last name of business contact.",
			BodyPath: "lastName",
		},
		&requestflag.Flag[string]{
			Name:     "phone",
			Usage:    "Valid phone number in e.164 international format.",
			BodyPath: "phone",
		},
		&requestflag.Flag[string]{
			Name:     "postal-code",
			Usage:    "Postal codes. Use 5 digit zipcode for United States",
			BodyPath: "postalCode",
		},
		&requestflag.Flag[string]{
			Name:     "state",
			Usage:    "State. Must be 2 letters code for United States.",
			BodyPath: "state",
		},
		&requestflag.Flag[string]{
			Name:     "stock-exchange",
			Usage:    "(Required for public company) stock exchange.",
			BodyPath: "stockExchange",
		},
		&requestflag.Flag[string]{
			Name:     "stock-symbol",
			Usage:    "(Required for public company) stock symbol.",
			BodyPath: "stockSymbol",
		},
		&requestflag.Flag[string]{
			Name:     "street",
			Usage:    "Street number and name.",
			BodyPath: "street",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-failover-url",
			Usage:    "Webhook failover URL for brand status updates.",
			BodyPath: "webhookFailoverURL",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "Webhook URL for brand status updates.",
			BodyPath: "webhookURL",
		},
		&requestflag.Flag[string]{
			Name:     "website",
			Usage:    "Brand website URL.",
			BodyPath: "website",
		},
	},
	Action:          handleMessaging10dlcBrandUpdate,
	HideHelpCommand: true,
}

var messaging10dlcBrandList = cli.Command{
	Name:    "list",
	Usage:   "This endpoint is used to list all brands associated with your organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "brand-id",
			Usage:     "Filter results by the Telnyx Brand id",
			QueryPath: "brandId",
		},
		&requestflag.Flag[string]{
			Name:      "country",
			QueryPath: "country",
		},
		&requestflag.Flag[string]{
			Name:      "display-name",
			QueryPath: "displayName",
		},
		&requestflag.Flag[string]{
			Name:      "entity-type",
			QueryPath: "entityType",
		},
		&requestflag.Flag[int64]{
			Name:      "page",
			Default:   1,
			QueryPath: "page",
		},
		&requestflag.Flag[int64]{
			Name:      "records-per-page",
			Usage:     "number of records per page. maximum of 500",
			Default:   10,
			QueryPath: "recordsPerPage",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Specifies the sort order for results. If not given, results are sorted by createdAt in descending order.",
			Default:   "-createdAt",
			QueryPath: "sort",
		},
		&requestflag.Flag[string]{
			Name:      "state",
			QueryPath: "state",
		},
		&requestflag.Flag[string]{
			Name:      "tcr-brand-id",
			Usage:     "Filter results by the TCR Brand id",
			QueryPath: "tcrBrandId",
		},
	},
	Action:          handleMessaging10dlcBrandList,
	HideHelpCommand: true,
}

var messaging10dlcBrandDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete Brand. This endpoint is used to delete a brand. Note the brand cannot be\ndeleted if it contains one or more active campaigns, the campaigns need to be\ninactive and at least 3 months old due to billing purposes.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "brand-id",
			Required: true,
		},
	},
	Action:          handleMessaging10dlcBrandDelete,
	HideHelpCommand: true,
}

var messaging10dlcBrandGetFeedback = cli.Command{
	Name:    "get-feedback",
	Usage:   "Get feedback about a brand by ID. This endpoint can be used after creating or\nrevetting a brand.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "brand-id",
			Required: true,
		},
	},
	Action:          handleMessaging10dlcBrandGetFeedback,
	HideHelpCommand: true,
}

var messaging10dlcBrandResend2faEmail = cli.Command{
	Name:    "resend-2fa-email",
	Usage:   "Resend brand 2FA email",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "brand-id",
			Required: true,
		},
	},
	Action:          handleMessaging10dlcBrandResend2faEmail,
	HideHelpCommand: true,
}

var messaging10dlcBrandRetrieveSMSOtpStatus = cli.Command{
	Name:    "retrieve-sms-otp-status",
	Usage:   "Query the status of an SMS OTP (One-Time Password) for Sole Proprietor brand\nverification using the Brand ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "brand-id",
			Required: true,
		},
	},
	Action:          handleMessaging10dlcBrandRetrieveSMSOtpStatus,
	HideHelpCommand: true,
}

var messaging10dlcBrandRevet = cli.Command{
	Name:    "revet",
	Usage:   "This operation allows you to revet the brand. However, revetting is allowed once\nafter the successful brand registration and thereafter limited to once every 3\nmonths.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "brand-id",
			Required: true,
		},
	},
	Action:          handleMessaging10dlcBrandRevet,
	HideHelpCommand: true,
}

var messaging10dlcBrandTriggerSMSOtp = cli.Command{
	Name:    "trigger-sms-otp",
	Usage:   "Trigger or re-trigger an SMS OTP (One-Time Password) for Sole Proprietor brand\nverification.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "brand-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "pin-sms",
			Usage:    "SMS message template to send the OTP. Must include `@OTP_PIN@` placeholder which will be replaced with the actual PIN",
			Required: true,
			BodyPath: "pinSms",
		},
		&requestflag.Flag[string]{
			Name:     "success-sms",
			Usage:    "SMS message to send upon successful OTP verification",
			Required: true,
			BodyPath: "successSms",
		},
	},
	Action:          handleMessaging10dlcBrandTriggerSMSOtp,
	HideHelpCommand: true,
}

var messaging10dlcBrandVerifySMSOtp = cli.Command{
	Name:    "verify-sms-otp",
	Usage:   "Verify the SMS OTP (One-Time Password) for Sole Proprietor brand verification.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "brand-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "otp-pin",
			Usage:    "The OTP PIN received via SMS",
			Required: true,
			BodyPath: "otpPin",
		},
	},
	Action:          handleMessaging10dlcBrandVerifySMSOtp,
	HideHelpCommand: true,
}

func handleMessaging10dlcBrandCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.Messaging10dlcBrandNewParams{}

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
	_, err = client.Messaging10dlc.Brand.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:brand create", obj, format, transform)
}

func handleMessaging10dlcBrandRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("brand-id") && len(unusedArgs) > 0 {
		cmd.Set("brand-id", unusedArgs[0])
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
	_, err = client.Messaging10dlc.Brand.Get(ctx, cmd.Value("brand-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:brand retrieve", obj, format, transform)
}

func handleMessaging10dlcBrandUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("brand-id") && len(unusedArgs) > 0 {
		cmd.Set("brand-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.Messaging10dlcBrandUpdateParams{}

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
	_, err = client.Messaging10dlc.Brand.Update(
		ctx,
		cmd.Value("brand-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:brand update", obj, format, transform)
}

func handleMessaging10dlcBrandList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.Messaging10dlcBrandListParams{}

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
		_, err = client.Messaging10dlc.Brand.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "messaging-10dlc:brand list", obj, format, transform)
	} else {
		iter := client.Messaging10dlc.Brand.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "messaging-10dlc:brand list", iter, format, transform)
	}
}

func handleMessaging10dlcBrandDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("brand-id") && len(unusedArgs) > 0 {
		cmd.Set("brand-id", unusedArgs[0])
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

	return client.Messaging10dlc.Brand.Delete(ctx, cmd.Value("brand-id").(string), options...)
}

func handleMessaging10dlcBrandGetFeedback(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("brand-id") && len(unusedArgs) > 0 {
		cmd.Set("brand-id", unusedArgs[0])
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
	_, err = client.Messaging10dlc.Brand.GetFeedback(ctx, cmd.Value("brand-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:brand get-feedback", obj, format, transform)
}

func handleMessaging10dlcBrandResend2faEmail(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("brand-id") && len(unusedArgs) > 0 {
		cmd.Set("brand-id", unusedArgs[0])
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

	return client.Messaging10dlc.Brand.Resend2faEmail(ctx, cmd.Value("brand-id").(string), options...)
}

func handleMessaging10dlcBrandRetrieveSMSOtpStatus(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("brand-id") && len(unusedArgs) > 0 {
		cmd.Set("brand-id", unusedArgs[0])
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
	_, err = client.Messaging10dlc.Brand.GetSMSOtpStatus(ctx, cmd.Value("brand-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:brand retrieve-sms-otp-status", obj, format, transform)
}

func handleMessaging10dlcBrandRevet(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("brand-id") && len(unusedArgs) > 0 {
		cmd.Set("brand-id", unusedArgs[0])
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
	_, err = client.Messaging10dlc.Brand.Revet(ctx, cmd.Value("brand-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:brand revet", obj, format, transform)
}

func handleMessaging10dlcBrandTriggerSMSOtp(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("brand-id") && len(unusedArgs) > 0 {
		cmd.Set("brand-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.Messaging10dlcBrandTriggerSMSOtpParams{}

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
	_, err = client.Messaging10dlc.Brand.TriggerSMSOtp(
		ctx,
		cmd.Value("brand-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-10dlc:brand trigger-sms-otp", obj, format, transform)
}

func handleMessaging10dlcBrandVerifySMSOtp(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("brand-id") && len(unusedArgs) > 0 {
		cmd.Set("brand-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.Messaging10dlcBrandVerifySMSOtpParams{}

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

	return client.Messaging10dlc.Brand.VerifySMSOtp(
		ctx,
		cmd.Value("brand-id").(string),
		params,
		options...,
	)
}
