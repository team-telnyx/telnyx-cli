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

var messagingTollfreeVerificationRequestsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Submit a new tollfree verification request",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "additional-information",
			Usage:    "Any additional information",
			Required: true,
			BodyPath: "additionalInformation",
		},
		&requestflag.Flag[string]{
			Name:     "business-addr1",
			Usage:    "Line 1 of the business address",
			Required: true,
			BodyPath: "businessAddr1",
		},
		&requestflag.Flag[string]{
			Name:     "business-city",
			Usage:    "The city of the business address; the first letter should be capitalized",
			Required: true,
			BodyPath: "businessCity",
		},
		&requestflag.Flag[string]{
			Name:     "business-contact-email",
			Usage:    "The email address of the business contact",
			Required: true,
			BodyPath: "businessContactEmail",
		},
		&requestflag.Flag[string]{
			Name:     "business-contact-first-name",
			Usage:    "First name of the business contact; there are no specific requirements on formatting",
			Required: true,
			BodyPath: "businessContactFirstName",
		},
		&requestflag.Flag[string]{
			Name:     "business-contact-last-name",
			Usage:    "Last name of the business contact; there are no specific requirements on formatting",
			Required: true,
			BodyPath: "businessContactLastName",
		},
		&requestflag.Flag[string]{
			Name:     "business-contact-phone",
			Usage:    "The phone number of the business contact in E.164 format",
			Required: true,
			BodyPath: "businessContactPhone",
		},
		&requestflag.Flag[string]{
			Name:     "business-name",
			Usage:    "Name of the business; there are no specific formatting requirements",
			Required: true,
			BodyPath: "businessName",
		},
		&requestflag.Flag[string]{
			Name:     "business-state",
			Usage:    "The full name of the state (not the 2 letter code) of the business address; the first letter should be capitalized",
			Required: true,
			BodyPath: "businessState",
		},
		&requestflag.Flag[string]{
			Name:     "business-zip",
			Usage:    "The ZIP code of the business address",
			Required: true,
			BodyPath: "businessZip",
		},
		&requestflag.Flag[string]{
			Name:     "corporate-website",
			Usage:    "A URL, including the scheme, pointing to the corporate website",
			Required: true,
			BodyPath: "corporateWebsite",
		},
		&requestflag.Flag[string]{
			Name:     "isv-reseller",
			Usage:    "ISV name",
			Required: true,
			BodyPath: "isvReseller",
		},
		&requestflag.Flag[string]{
			Name:     "message-volume",
			Usage:    "Message Volume Enums",
			Required: true,
			BodyPath: "messageVolume",
		},
		&requestflag.Flag[string]{
			Name:     "opt-in-workflow",
			Usage:    "Human-readable description of how end users will opt into receiving messages from the given phone numbers",
			Required: true,
			BodyPath: "optInWorkflow",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "opt-in-workflow-image-url",
			Usage:    "Images showing the opt-in workflow",
			Required: true,
			BodyPath: "optInWorkflowImageURLs",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "phone-number",
			Usage:    "The phone numbers to request the verification of",
			Required: true,
			BodyPath: "phoneNumbers",
		},
		&requestflag.Flag[string]{
			Name:     "production-message-content",
			Usage:    "An example of a message that will be sent from the given phone numbers",
			Required: true,
			BodyPath: "productionMessageContent",
		},
		&requestflag.Flag[string]{
			Name:     "use-case",
			Usage:    "Tollfree usecase categories",
			Required: true,
			BodyPath: "useCase",
		},
		&requestflag.Flag[string]{
			Name:     "use-case-summary",
			Usage:    "Human-readable summary of the desired use-case",
			Required: true,
			BodyPath: "useCaseSummary",
		},
		&requestflag.Flag[bool]{
			Name:     "age-gated-content",
			Usage:    "Indicates if messaging content requires age gating (e.g., 18+). Defaults to false if not provided.",
			Default:  false,
			BodyPath: "ageGatedContent",
		},
		&requestflag.Flag[string]{
			Name:     "business-addr2",
			Usage:    "Line 2 of the business address",
			BodyPath: "businessAddr2",
		},
		&requestflag.Flag[any]{
			Name:     "business-registration-country",
			Usage:    "ISO 3166-1 alpha-2 country code of the issuing business authority. Must be exactly 2 letters. Automatically converted to uppercase. Required from January 2026.",
			BodyPath: "businessRegistrationCountry",
		},
		&requestflag.Flag[any]{
			Name:     "business-registration-number",
			Usage:    "Official business registration number (e.g., Employer Identification Number (EIN) in the U.S.). Required from January 2026.",
			BodyPath: "businessRegistrationNumber",
		},
		&requestflag.Flag[any]{
			Name:     "business-registration-type",
			Usage:    "Type of business registration being provided. Required from January 2026.",
			BodyPath: "businessRegistrationType",
		},
		&requestflag.Flag[any]{
			Name:     "campaign-verify-authorization-token",
			Usage:    "Campaign Verify Authorization Token required for Political use case submissions starting February 17, 2026. This token is validated by Zipwhip and must be provided for all Political use case verifications after the deadline.",
			BodyPath: "campaignVerifyAuthorizationToken",
		},
		&requestflag.Flag[any]{
			Name:     "doing-business-as",
			Usage:    "Doing Business As (DBA) name if different from legal name",
			BodyPath: "doingBusinessAs",
		},
		&requestflag.Flag[string]{
			Name:     "entity-type",
			Usage:    "Business entity classification",
			BodyPath: "entityType",
		},
		&requestflag.Flag[any]{
			Name:     "help-message-response",
			Usage:    "The message returned when users text 'HELP'",
			BodyPath: "helpMessageResponse",
		},
		&requestflag.Flag[any]{
			Name:     "opt-in-confirmation-response",
			Usage:    "Message sent to users confirming their opt-in to receive messages",
			BodyPath: "optInConfirmationResponse",
		},
		&requestflag.Flag[any]{
			Name:     "opt-in-keywords",
			Usage:    "Keywords used to collect and process consumer opt-ins",
			BodyPath: "optInKeywords",
		},
		&requestflag.Flag[any]{
			Name:     "privacy-policy-url",
			Usage:    "URL pointing to the business's privacy policy. Plain string, no URL format validation.",
			BodyPath: "privacyPolicyURL",
		},
		&requestflag.Flag[any]{
			Name:     "terms-and-condition-url",
			Usage:    "URL pointing to the business's terms and conditions. Plain string, no URL format validation.",
			BodyPath: "termsAndConditionURL",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "URL that should receive webhooks relating to this verification request",
			BodyPath: "webhookUrl",
		},
	},
	Action:          handleMessagingTollfreeVerificationRequestsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"opt-in-workflow-image-url": {
		&requestflag.InnerFlag[string]{
			Name:       "opt-in-workflow-image-url.url",
			InnerField: "url",
		},
	},
	"phone-number": {
		&requestflag.InnerFlag[string]{
			Name:       "phone-number.phone-number",
			InnerField: "phoneNumber",
		},
	},
})

var messagingTollfreeVerificationRequestsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a single verification request by its ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleMessagingTollfreeVerificationRequestsRetrieve,
	HideHelpCommand: true,
}

var messagingTollfreeVerificationRequestsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update an existing tollfree verification request. This is particularly useful\nwhen there are pending customer actions to be taken.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "additional-information",
			Usage:    "Any additional information",
			Required: true,
			BodyPath: "additionalInformation",
		},
		&requestflag.Flag[string]{
			Name:     "business-addr1",
			Usage:    "Line 1 of the business address",
			Required: true,
			BodyPath: "businessAddr1",
		},
		&requestflag.Flag[string]{
			Name:     "business-city",
			Usage:    "The city of the business address; the first letter should be capitalized",
			Required: true,
			BodyPath: "businessCity",
		},
		&requestflag.Flag[string]{
			Name:     "business-contact-email",
			Usage:    "The email address of the business contact",
			Required: true,
			BodyPath: "businessContactEmail",
		},
		&requestflag.Flag[string]{
			Name:     "business-contact-first-name",
			Usage:    "First name of the business contact; there are no specific requirements on formatting",
			Required: true,
			BodyPath: "businessContactFirstName",
		},
		&requestflag.Flag[string]{
			Name:     "business-contact-last-name",
			Usage:    "Last name of the business contact; there are no specific requirements on formatting",
			Required: true,
			BodyPath: "businessContactLastName",
		},
		&requestflag.Flag[string]{
			Name:     "business-contact-phone",
			Usage:    "The phone number of the business contact in E.164 format",
			Required: true,
			BodyPath: "businessContactPhone",
		},
		&requestflag.Flag[string]{
			Name:     "business-name",
			Usage:    "Name of the business; there are no specific formatting requirements",
			Required: true,
			BodyPath: "businessName",
		},
		&requestflag.Flag[string]{
			Name:     "business-state",
			Usage:    "The full name of the state (not the 2 letter code) of the business address; the first letter should be capitalized",
			Required: true,
			BodyPath: "businessState",
		},
		&requestflag.Flag[string]{
			Name:     "business-zip",
			Usage:    "The ZIP code of the business address",
			Required: true,
			BodyPath: "businessZip",
		},
		&requestflag.Flag[string]{
			Name:     "corporate-website",
			Usage:    "A URL, including the scheme, pointing to the corporate website",
			Required: true,
			BodyPath: "corporateWebsite",
		},
		&requestflag.Flag[string]{
			Name:     "isv-reseller",
			Usage:    "ISV name",
			Required: true,
			BodyPath: "isvReseller",
		},
		&requestflag.Flag[string]{
			Name:     "message-volume",
			Usage:    "Message Volume Enums",
			Required: true,
			BodyPath: "messageVolume",
		},
		&requestflag.Flag[string]{
			Name:     "opt-in-workflow",
			Usage:    "Human-readable description of how end users will opt into receiving messages from the given phone numbers",
			Required: true,
			BodyPath: "optInWorkflow",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "opt-in-workflow-image-url",
			Usage:    "Images showing the opt-in workflow",
			Required: true,
			BodyPath: "optInWorkflowImageURLs",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "phone-number",
			Usage:    "The phone numbers to request the verification of",
			Required: true,
			BodyPath: "phoneNumbers",
		},
		&requestflag.Flag[string]{
			Name:     "production-message-content",
			Usage:    "An example of a message that will be sent from the given phone numbers",
			Required: true,
			BodyPath: "productionMessageContent",
		},
		&requestflag.Flag[string]{
			Name:     "use-case",
			Usage:    "Tollfree usecase categories",
			Required: true,
			BodyPath: "useCase",
		},
		&requestflag.Flag[string]{
			Name:     "use-case-summary",
			Usage:    "Human-readable summary of the desired use-case",
			Required: true,
			BodyPath: "useCaseSummary",
		},
		&requestflag.Flag[bool]{
			Name:     "age-gated-content",
			Usage:    "Indicates if messaging content requires age gating (e.g., 18+). Defaults to false if not provided.",
			Default:  false,
			BodyPath: "ageGatedContent",
		},
		&requestflag.Flag[string]{
			Name:     "business-addr2",
			Usage:    "Line 2 of the business address",
			BodyPath: "businessAddr2",
		},
		&requestflag.Flag[any]{
			Name:     "business-registration-country",
			Usage:    "ISO 3166-1 alpha-2 country code of the issuing business authority. Must be exactly 2 letters. Automatically converted to uppercase. Required from January 2026.",
			BodyPath: "businessRegistrationCountry",
		},
		&requestflag.Flag[any]{
			Name:     "business-registration-number",
			Usage:    "Official business registration number (e.g., Employer Identification Number (EIN) in the U.S.). Required from January 2026.",
			BodyPath: "businessRegistrationNumber",
		},
		&requestflag.Flag[any]{
			Name:     "business-registration-type",
			Usage:    "Type of business registration being provided. Required from January 2026.",
			BodyPath: "businessRegistrationType",
		},
		&requestflag.Flag[any]{
			Name:     "campaign-verify-authorization-token",
			Usage:    "Campaign Verify Authorization Token required for Political use case submissions starting February 17, 2026. This token is validated by Zipwhip and must be provided for all Political use case verifications after the deadline.",
			BodyPath: "campaignVerifyAuthorizationToken",
		},
		&requestflag.Flag[any]{
			Name:     "doing-business-as",
			Usage:    "Doing Business As (DBA) name if different from legal name",
			BodyPath: "doingBusinessAs",
		},
		&requestflag.Flag[string]{
			Name:     "entity-type",
			Usage:    "Business entity classification",
			BodyPath: "entityType",
		},
		&requestflag.Flag[any]{
			Name:     "help-message-response",
			Usage:    "The message returned when users text 'HELP'",
			BodyPath: "helpMessageResponse",
		},
		&requestflag.Flag[any]{
			Name:     "opt-in-confirmation-response",
			Usage:    "Message sent to users confirming their opt-in to receive messages",
			BodyPath: "optInConfirmationResponse",
		},
		&requestflag.Flag[any]{
			Name:     "opt-in-keywords",
			Usage:    "Keywords used to collect and process consumer opt-ins",
			BodyPath: "optInKeywords",
		},
		&requestflag.Flag[any]{
			Name:     "privacy-policy-url",
			Usage:    "URL pointing to the business's privacy policy. Plain string, no URL format validation.",
			BodyPath: "privacyPolicyURL",
		},
		&requestflag.Flag[any]{
			Name:     "terms-and-condition-url",
			Usage:    "URL pointing to the business's terms and conditions. Plain string, no URL format validation.",
			BodyPath: "termsAndConditionURL",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "URL that should receive webhooks relating to this verification request",
			BodyPath: "webhookUrl",
		},
	},
	Action:          handleMessagingTollfreeVerificationRequestsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"opt-in-workflow-image-url": {
		&requestflag.InnerFlag[string]{
			Name:       "opt-in-workflow-image-url.url",
			InnerField: "url",
		},
	},
	"phone-number": {
		&requestflag.InnerFlag[string]{
			Name:       "phone-number.phone-number",
			InnerField: "phoneNumber",
		},
	},
})

var messagingTollfreeVerificationRequestsList = cli.Command{
	Name:    "list",
	Usage:   "Get a list of previously-submitted tollfree verification requests",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "page",
			Required:  true,
			QueryPath: "page",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "\n        Request this many records per page\n\n        This value is automatically clamped if the provided value is too large.\n        ",
			Required:  true,
			QueryPath: "page_size",
		},
		&requestflag.Flag[any]{
			Name:      "date-end",
			QueryPath: "date_end",
		},
		&requestflag.Flag[any]{
			Name:      "date-start",
			QueryPath: "date_start",
		},
		&requestflag.Flag[string]{
			Name:      "phone-number",
			QueryPath: "phone_number",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Tollfree verification status",
			QueryPath: "status",
		},
	},
	Action:          handleMessagingTollfreeVerificationRequestsList,
	HideHelpCommand: true,
}

var messagingTollfreeVerificationRequestsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a verification request",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleMessagingTollfreeVerificationRequestsDelete,
	HideHelpCommand: true,
}

func handleMessagingTollfreeVerificationRequestsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MessagingTollfreeVerificationRequestNewParams{}

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
	_, err = client.MessagingTollfree.Verification.Requests.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-tollfree:verification:requests create", obj, format, transform)
}

func handleMessagingTollfreeVerificationRequestsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.MessagingTollfree.Verification.Requests.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-tollfree:verification:requests retrieve", obj, format, transform)
}

func handleMessagingTollfreeVerificationRequestsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MessagingTollfreeVerificationRequestUpdateParams{}

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
	_, err = client.MessagingTollfree.Verification.Requests.Update(
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
	return ShowJSON(os.Stdout, "messaging-tollfree:verification:requests update", obj, format, transform)
}

func handleMessagingTollfreeVerificationRequestsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MessagingTollfreeVerificationRequestListParams{}

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
		_, err = client.MessagingTollfree.Verification.Requests.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "messaging-tollfree:verification:requests list", obj, format, transform)
	} else {
		iter := client.MessagingTollfree.Verification.Requests.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "messaging-tollfree:verification:requests list", iter, format, transform)
	}
}

func handleMessagingTollfreeVerificationRequestsDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.MessagingTollfree.Verification.Requests.Delete(ctx, cmd.Value("id").(string), options...)
}
