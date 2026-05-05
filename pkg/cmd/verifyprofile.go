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

var verifyProfilesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new Verify profile to associate verifications with.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "call",
			BodyPath: "call",
		},
		&requestflag.Flag[float64]{
			Name:     "daily-spend-limit",
			Usage:    "The maximum daily spend allowed on this verify profile, in USD.",
			BodyPath: "daily_spend_limit",
		},
		&requestflag.Flag[bool]{
			Name:     "daily-spend-limit-enabled",
			Usage:    "Whether the daily spend limit is enforced for this verify profile.",
			Default:  false,
			BodyPath: "daily_spend_limit_enabled",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "flashcall",
			BodyPath: "flashcall",
		},
		&requestflag.Flag[string]{
			Name:     "language",
			BodyPath: "language",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "sms",
			BodyPath: "sms",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-failover-url",
			BodyPath: "webhook_failover_url",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			BodyPath: "webhook_url",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "whatsapp",
			BodyPath: "whatsapp",
		},
	},
	Action:          handleVerifyProfilesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"call": {
		&requestflag.InnerFlag[string]{
			Name:       "call.app-name",
			Usage:      "The name that identifies the application requesting 2fa in the verification message.",
			InnerField: "app_name",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "call.code-length",
			Usage:      "The length of the verify code to generate.",
			InnerField: "code_length",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "call.default-verification-timeout-secs",
			Usage:      "For every request that is initiated via this Verify profile, this sets the number of seconds before a verification request code expires. Once the verification request expires, the user cannot use the code to verify their identity.",
			InnerField: "default_verification_timeout_secs",
		},
		&requestflag.InnerFlag[string]{
			Name:       "call.messaging-template-id",
			Usage:      "The message template identifier selected from /verify_profiles/templates",
			InnerField: "messaging_template_id",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "call.whitelisted-destinations",
			Usage:      "Enabled country destinations to send verification codes. The elements in the list must be valid ISO 3166-1 alpha-2 country codes. If set to `[\"*\"]`, all destinations will be allowed. **Conditionally required:** this field must be provided when your organization is configured to require explicit whitelisted destinations; otherwise it is optional.",
			InnerField: "whitelisted_destinations",
		},
	},
	"flashcall": {
		&requestflag.InnerFlag[string]{
			Name:       "flashcall.app-name",
			Usage:      "The name that identifies the application requesting 2fa in the verification message.",
			InnerField: "app_name",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "flashcall.default-verification-timeout-secs",
			Usage:      "For every request that is initiated via this Verify profile, this sets the number of seconds before a verification request code expires. Once the verification request expires, the user cannot use the code to verify their identity.",
			InnerField: "default_verification_timeout_secs",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "flashcall.whitelisted-destinations",
			Usage:      "Enabled country destinations to send verification codes. The elements in the list must be valid ISO 3166-1 alpha-2 country codes. If set to `[\"*\"]`, all destinations will be allowed. **Conditionally required:** this field must be provided when your organization is configured to require explicit whitelisted destinations; otherwise it is optional.",
			InnerField: "whitelisted_destinations",
		},
	},
	"sms": {
		&requestflag.InnerFlag[*string]{
			Name:       "sms.alpha-sender",
			Usage:      "The alphanumeric sender ID to use when sending to destinations that require an alphanumeric sender ID.",
			InnerField: "alpha_sender",
		},
		&requestflag.InnerFlag[string]{
			Name:       "sms.app-name",
			Usage:      "The name that identifies the application requesting 2fa in the verification message.",
			InnerField: "app_name",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "sms.code-length",
			Usage:      "The length of the verify code to generate.",
			InnerField: "code_length",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "sms.default-verification-timeout-secs",
			Usage:      "For every request that is initiated via this Verify profile, this sets the number of seconds before a verification request code expires. Once the verification request expires, the user cannot use the code to verify their identity.",
			InnerField: "default_verification_timeout_secs",
		},
		&requestflag.InnerFlag[string]{
			Name:       "sms.messaging-template-id",
			Usage:      "The message template identifier selected from /verify_profiles/templates",
			InnerField: "messaging_template_id",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "sms.whitelisted-destinations",
			Usage:      "Enabled country destinations to send verification codes. The elements in the list must be valid ISO 3166-1 alpha-2 country codes. If set to `[\"*\"]`, all destinations will be allowed. **Conditionally required:** this field must be provided when your organization is configured to require explicit whitelisted destinations; otherwise it is optional.",
			InnerField: "whitelisted_destinations",
		},
	},
	"whatsapp": {
		&requestflag.InnerFlag[int64]{
			Name:       "whatsapp.default-verification-timeout-secs",
			Usage:      "For every request that is initiated via this Verify profile, this sets the number of seconds before a verification request code expires. Once the verification request expires, the user cannot use the code to verify their identity.",
			InnerField: "default_verification_timeout_secs",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "whatsapp.sender-phone-number",
			Usage:      "Phone number registered on the customer WABA to send OTPs from",
			InnerField: "sender_phone_number",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "whatsapp.template-id",
			Usage:      "Customer pre-approved authentication template name registered on Meta",
			InnerField: "template_id",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "whatsapp.waba-id",
			Usage:      "Customer Meta WABA ID for Bring-Your-Own-WABA sending",
			InnerField: "waba_id",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "whatsapp.whitelisted-destinations",
			Usage:      "Enabled country destinations to send verification codes. The elements in the list must be valid ISO 3166-1 alpha-2 country codes. If set to `[\"*\"]`, all destinations will be allowed. **Conditionally required:** this field must be provided when your organization is configured to require explicit whitelisted destinations; otherwise it is optional.",
			InnerField: "whitelisted_destinations",
		},
	},
})

var verifyProfilesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Gets a single Verify profile.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "verify-profile-id",
			Required:  true,
			PathParam: "verify_profile_id",
		},
	},
	Action:          handleVerifyProfilesRetrieve,
	HideHelpCommand: true,
}

var verifyProfilesUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update Verify profile",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "verify-profile-id",
			Required:  true,
			PathParam: "verify_profile_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "call",
			BodyPath: "call",
		},
		&requestflag.Flag[float64]{
			Name:     "daily-spend-limit",
			Usage:    "The maximum daily spend allowed on this verify profile, in USD.",
			BodyPath: "daily_spend_limit",
		},
		&requestflag.Flag[bool]{
			Name:     "daily-spend-limit-enabled",
			Usage:    "Whether the daily spend limit is enforced for this verify profile.",
			BodyPath: "daily_spend_limit_enabled",
		},
		&requestflag.Flag[string]{
			Name:     "language",
			BodyPath: "language",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "sms",
			BodyPath: "sms",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-failover-url",
			BodyPath: "webhook_failover_url",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			BodyPath: "webhook_url",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "whatsapp",
			BodyPath: "whatsapp",
		},
	},
	Action:          handleVerifyProfilesUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"call": {
		&requestflag.InnerFlag[string]{
			Name:       "call.app-name",
			Usage:      "The name that identifies the application requesting 2fa in the verification message.",
			InnerField: "app_name",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "call.code-length",
			Usage:      "The length of the verify code to generate.",
			InnerField: "code_length",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "call.default-verification-timeout-secs",
			Usage:      "For every request that is initiated via this Verify profile, this sets the number of seconds before a verification request code expires. Once the verification request expires, the user cannot use the code to verify their identity.",
			InnerField: "default_verification_timeout_secs",
		},
		&requestflag.InnerFlag[string]{
			Name:       "call.messaging-template-id",
			Usage:      "The message template identifier selected from /verify_profiles/templates",
			InnerField: "messaging_template_id",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "call.whitelisted-destinations",
			Usage:      "Enabled country destinations to send verification codes. The elements in the list must be valid ISO 3166-1 alpha-2 country codes. If set to `[\"*\"]`, all destinations will be allowed. **Conditionally required:** this field must be provided when your organization is configured to require explicit whitelisted destinations; otherwise it is optional.",
			InnerField: "whitelisted_destinations",
		},
	},
	"sms": {
		&requestflag.InnerFlag[*string]{
			Name:       "sms.alpha-sender",
			Usage:      "The alphanumeric sender ID to use when sending to destinations that require an alphanumeric sender ID.",
			InnerField: "alpha_sender",
		},
		&requestflag.InnerFlag[string]{
			Name:       "sms.app-name",
			Usage:      "The name that identifies the application requesting 2fa in the verification message.",
			InnerField: "app_name",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "sms.code-length",
			Usage:      "The length of the verify code to generate.",
			InnerField: "code_length",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "sms.default-verification-timeout-secs",
			Usage:      "For every request that is initiated via this Verify profile, this sets the number of seconds before a verification request code expires. Once the verification request expires, the user cannot use the code to verify their identity.",
			InnerField: "default_verification_timeout_secs",
		},
		&requestflag.InnerFlag[string]{
			Name:       "sms.messaging-template-id",
			Usage:      "The message template identifier selected from /verify_profiles/templates",
			InnerField: "messaging_template_id",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "sms.whitelisted-destinations",
			Usage:      "Enabled country destinations to send verification codes. The elements in the list must be valid ISO 3166-1 alpha-2 country codes. If set to `[\"*\"]`, all destinations will be allowed. **Conditionally required:** this field must be provided when your organization is configured to require explicit whitelisted destinations; otherwise it is optional.",
			InnerField: "whitelisted_destinations",
		},
	},
	"whatsapp": {
		&requestflag.InnerFlag[int64]{
			Name:       "whatsapp.default-verification-timeout-secs",
			Usage:      "For every request that is initiated via this Verify profile, this sets the number of seconds before a verification request code expires. Once the verification request expires, the user cannot use the code to verify their identity.",
			InnerField: "default_verification_timeout_secs",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "whatsapp.sender-phone-number",
			Usage:      "Phone number registered on the customer WABA to send OTPs from",
			InnerField: "sender_phone_number",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "whatsapp.template-id",
			Usage:      "Customer pre-approved authentication template name registered on Meta",
			InnerField: "template_id",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "whatsapp.waba-id",
			Usage:      "Customer Meta WABA ID for Bring-Your-Own-WABA sending",
			InnerField: "waba_id",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "whatsapp.whitelisted-destinations",
			Usage:      "Enabled country destinations to send verification codes. The elements in the list must be valid ISO 3166-1 alpha-2 country codes. If set to `[\"*\"]`, all destinations will be allowed. **Conditionally required:** this field must be provided when your organization is configured to require explicit whitelisted destinations; otherwise it is optional.",
			InnerField: "whitelisted_destinations",
		},
	},
})

var verifyProfilesList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Gets a paginated list of Verify profiles.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[name]",
			QueryPath: "filter",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page[size]",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleVerifyProfilesList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.name",
			Usage:      "Optional filter for profile names.",
			InnerField: "name",
		},
	},
})

var verifyProfilesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete Verify profile",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "verify-profile-id",
			Required:  true,
			PathParam: "verify_profile_id",
		},
	},
	Action:          handleVerifyProfilesDelete,
	HideHelpCommand: true,
}

var verifyProfilesCreateTemplate = cli.Command{
	Name:    "create-template",
	Usage:   "Create a new Verify profile message template.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "The text content of the message template.",
			Required: true,
			BodyPath: "text",
		},
	},
	Action:          handleVerifyProfilesCreateTemplate,
	HideHelpCommand: true,
}

var verifyProfilesRetrieveTemplates = cli.Command{
	Name:            "retrieve-templates",
	Usage:           "List all Verify profile message templates.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleVerifyProfilesRetrieveTemplates,
	HideHelpCommand: true,
}

var verifyProfilesUpdateTemplate = cli.Command{
	Name:    "update-template",
	Usage:   "Update an existing Verify profile message template.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "template-id",
			Required:  true,
			PathParam: "template_id",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "The text content of the message template.",
			Required: true,
			BodyPath: "text",
		},
	},
	Action:          handleVerifyProfilesUpdateTemplate,
	HideHelpCommand: true,
}

func handleVerifyProfilesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.VerifyProfileNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.VerifyProfiles.New(ctx, params, options...)
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
		Title:          "verify-profiles create",
		Transform:      transform,
	})
}

func handleVerifyProfilesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("verify-profile-id") && len(unusedArgs) > 0 {
		cmd.Set("verify-profile-id", unusedArgs[0])
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
	_, err = client.VerifyProfiles.Get(ctx, cmd.Value("verify-profile-id").(string), options...)
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
		Title:          "verify-profiles retrieve",
		Transform:      transform,
	})
}

func handleVerifyProfilesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("verify-profile-id") && len(unusedArgs) > 0 {
		cmd.Set("verify-profile-id", unusedArgs[0])
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

	params := telnyx.VerifyProfileUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.VerifyProfiles.Update(
		ctx,
		cmd.Value("verify-profile-id").(string),
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
		Title:          "verify-profiles update",
		Transform:      transform,
	})
}

func handleVerifyProfilesList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.VerifyProfileListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.VerifyProfiles.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "verify-profiles list",
			Transform:      transform,
		})
	} else {
		iter := client.VerifyProfiles.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "verify-profiles list",
			Transform:      transform,
		})
	}
}

func handleVerifyProfilesDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("verify-profile-id") && len(unusedArgs) > 0 {
		cmd.Set("verify-profile-id", unusedArgs[0])
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
	_, err = client.VerifyProfiles.Delete(ctx, cmd.Value("verify-profile-id").(string), options...)
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
		Title:          "verify-profiles delete",
		Transform:      transform,
	})
}

func handleVerifyProfilesCreateTemplate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.VerifyProfileNewTemplateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.VerifyProfiles.NewTemplate(ctx, params, options...)
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
		Title:          "verify-profiles create-template",
		Transform:      transform,
	})
}

func handleVerifyProfilesRetrieveTemplates(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.VerifyProfiles.GetTemplates(ctx, options...)
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
		Title:          "verify-profiles retrieve-templates",
		Transform:      transform,
	})
}

func handleVerifyProfilesUpdateTemplate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("template-id") && len(unusedArgs) > 0 {
		cmd.Set("template-id", unusedArgs[0])
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

	params := telnyx.VerifyProfileUpdateTemplateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.VerifyProfiles.UpdateTemplate(
		ctx,
		cmd.Value("template-id").(string),
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
		Title:          "verify-profiles update-template",
		Transform:      transform,
	})
}
