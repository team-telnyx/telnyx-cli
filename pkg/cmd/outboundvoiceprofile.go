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

var outboundVoiceProfilesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create an outbound voice profile.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "A user-supplied name to help with organization.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "billing-group-id",
			Usage:    "The ID of the billing group associated with the outbound proflile. Defaults to null (for no group assigned).",
			Default:  nil,
			BodyPath: "billing_group_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "call-recording",
			BodyPath: "call_recording",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "calling-window",
			Usage:    "(BETA) Specifies the time window and call limits for calls made using this outbound voice profile. Note that all times are UTC in 24-hour clock time.",
			BodyPath: "calling_window",
		},
		&requestflag.Flag[any]{
			Name:     "concurrent-call-limit",
			Usage:    "Must be no more than your global concurrent call limit. Null means no limit.",
			BodyPath: "concurrent_call_limit",
		},
		&requestflag.Flag[string]{
			Name:     "daily-spend-limit",
			Usage:    "The maximum amount of usage charges, in USD, you want Telnyx to allow on this outbound voice profile in a day before disallowing new calls.",
			BodyPath: "daily_spend_limit",
		},
		&requestflag.Flag[bool]{
			Name:     "daily-spend-limit-enabled",
			Usage:    "Specifies whether to enforce the daily_spend_limit on this outbound voice profile.",
			Default:  false,
			BodyPath: "daily_spend_limit_enabled",
		},
		&requestflag.Flag[bool]{
			Name:     "enabled",
			Usage:    "Specifies whether the outbound voice profile can be used. Disabled profiles will result in outbound calls being blocked for the associated Connections.",
			Default:  true,
			BodyPath: "enabled",
		},
		&requestflag.Flag[float64]{
			Name:     "max-destination-rate",
			Usage:    "Maximum rate (price per minute) for a Destination to be allowed when making outbound calls.",
			BodyPath: "max_destination_rate",
		},
		&requestflag.Flag[string]{
			Name:     "service-plan",
			Usage:    "Indicates the coverage of the termination regions.",
			Default:  "global",
			BodyPath: "service_plan",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "traffic-type",
			Usage:    "Specifies the type of traffic allowed in this profile.",
			Default:  "conversational",
			BodyPath: "traffic_type",
		},
		&requestflag.Flag[string]{
			Name:     "usage-payment-method",
			Usage:    "Setting for how costs for outbound profile are calculated.",
			Default:  "rate-deck",
			BodyPath: "usage_payment_method",
		},
		&requestflag.Flag[[]string]{
			Name:     "whitelisted-destination",
			Usage:    "The list of destinations you want to be able to call using this outbound voice profile formatted in alpha2.",
			Default:  []string{"US", "CA"},
			BodyPath: "whitelisted_destinations",
		},
	},
	Action:          handleOutboundVoiceProfilesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"call-recording": {
		&requestflag.InnerFlag[[]string]{
			Name:       "call-recording.call-recording-caller-phone-numbers",
			Usage:      "When call_recording_type is 'by_caller_phone_number', only outbound calls using one of these numbers will be recorded. Numbers must be specified in E164 format.",
			InnerField: "call_recording_caller_phone_numbers",
		},
		&requestflag.InnerFlag[string]{
			Name:       "call-recording.call-recording-channels",
			Usage:      "When using 'dual' channels, the final audio file will be a stereo recording with the first leg on channel A, and the rest on channel B.",
			InnerField: "call_recording_channels",
		},
		&requestflag.InnerFlag[string]{
			Name:       "call-recording.call-recording-format",
			Usage:      "The audio file format for calls being recorded.",
			InnerField: "call_recording_format",
		},
		&requestflag.InnerFlag[string]{
			Name:       "call-recording.call-recording-type",
			Usage:      "Specifies which calls are recorded.",
			InnerField: "call_recording_type",
		},
	},
	"calling-window": {
		&requestflag.InnerFlag[int64]{
			Name:       "calling-window.calls-per-cld",
			Usage:      "(BETA) The maximum number of calls that can be initiated to a single called party (CLD) within the calling window. A null value means no limit.",
			InnerField: "calls_per_cld",
		},
		&requestflag.InnerFlag[any]{
			Name:       "calling-window.end-time",
			Usage:      "(BETA) The UTC time of day (in HH:MM format, 24-hour clock) when calls are no longer allowed to start.",
			InnerField: "end_time",
		},
		&requestflag.InnerFlag[any]{
			Name:       "calling-window.start-time",
			Usage:      "(BETA) The UTC time of day (in HH:MM format, 24-hour clock) when calls are allowed to start.",
			InnerField: "start_time",
		},
	},
})

var outboundVoiceProfilesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves the details of an existing outbound voice profile.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleOutboundVoiceProfilesRetrieve,
	HideHelpCommand: true,
}

var outboundVoiceProfilesUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates an existing outbound voice profile.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "A user-supplied name to help with organization.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "billing-group-id",
			Usage:    "The ID of the billing group associated with the outbound proflile. Defaults to null (for no group assigned).",
			Default:  nil,
			BodyPath: "billing_group_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "call-recording",
			BodyPath: "call_recording",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "calling-window",
			Usage:    "(BETA) Specifies the time window and call limits for calls made using this outbound voice profile.",
			BodyPath: "calling_window",
		},
		&requestflag.Flag[any]{
			Name:     "concurrent-call-limit",
			Usage:    "Must be no more than your global concurrent call limit. Null means no limit.",
			BodyPath: "concurrent_call_limit",
		},
		&requestflag.Flag[string]{
			Name:     "daily-spend-limit",
			Usage:    "The maximum amount of usage charges, in USD, you want Telnyx to allow on this outbound voice profile in a day before disallowing new calls.",
			BodyPath: "daily_spend_limit",
		},
		&requestflag.Flag[bool]{
			Name:     "daily-spend-limit-enabled",
			Usage:    "Specifies whether to enforce the daily_spend_limit on this outbound voice profile.",
			Default:  false,
			BodyPath: "daily_spend_limit_enabled",
		},
		&requestflag.Flag[bool]{
			Name:     "enabled",
			Usage:    "Specifies whether the outbound voice profile can be used. Disabled profiles will result in outbound calls being blocked for the associated Connections.",
			Default:  true,
			BodyPath: "enabled",
		},
		&requestflag.Flag[float64]{
			Name:     "max-destination-rate",
			Usage:    "Maximum rate (price per minute) for a Destination to be allowed when making outbound calls.",
			BodyPath: "max_destination_rate",
		},
		&requestflag.Flag[string]{
			Name:     "service-plan",
			Usage:    "Indicates the coverage of the termination regions.",
			Default:  "global",
			BodyPath: "service_plan",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "traffic-type",
			Usage:    "Specifies the type of traffic allowed in this profile.",
			Default:  "conversational",
			BodyPath: "traffic_type",
		},
		&requestflag.Flag[string]{
			Name:     "usage-payment-method",
			Usage:    "Setting for how costs for outbound profile are calculated.",
			Default:  "rate-deck",
			BodyPath: "usage_payment_method",
		},
		&requestflag.Flag[[]string]{
			Name:     "whitelisted-destination",
			Usage:    "The list of destinations you want to be able to call using this outbound voice profile formatted in alpha2.",
			Default:  []string{"US", "CA"},
			BodyPath: "whitelisted_destinations",
		},
	},
	Action:          handleOutboundVoiceProfilesUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"call-recording": {
		&requestflag.InnerFlag[[]string]{
			Name:       "call-recording.call-recording-caller-phone-numbers",
			Usage:      "When call_recording_type is 'by_caller_phone_number', only outbound calls using one of these numbers will be recorded. Numbers must be specified in E164 format.",
			InnerField: "call_recording_caller_phone_numbers",
		},
		&requestflag.InnerFlag[string]{
			Name:       "call-recording.call-recording-channels",
			Usage:      "When using 'dual' channels, the final audio file will be a stereo recording with the first leg on channel A, and the rest on channel B.",
			InnerField: "call_recording_channels",
		},
		&requestflag.InnerFlag[string]{
			Name:       "call-recording.call-recording-format",
			Usage:      "The audio file format for calls being recorded.",
			InnerField: "call_recording_format",
		},
		&requestflag.InnerFlag[string]{
			Name:       "call-recording.call-recording-type",
			Usage:      "Specifies which calls are recorded.",
			InnerField: "call_recording_type",
		},
	},
	"calling-window": {
		&requestflag.InnerFlag[int64]{
			Name:       "calling-window.calls-per-cld",
			Usage:      "(BETA) The maximum number of calls that can be initiated to a single called party (CLD) within the calling window. A null value means no limit.",
			InnerField: "calls_per_cld",
		},
		&requestflag.InnerFlag[any]{
			Name:       "calling-window.end-time",
			Usage:      "(BETA) The UTC time of day (in HH:MM format, 24-hour clock) when calls are no longer allowed to start.",
			InnerField: "end_time",
		},
		&requestflag.InnerFlag[any]{
			Name:       "calling-window.start-time",
			Usage:      "(BETA) The UTC time of day (in HH:MM format, 24-hour clock) when calls are allowed to start.",
			InnerField: "start_time",
		},
	},
})

var outboundVoiceProfilesList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Get all outbound voice profiles belonging to the user that match the given\nfilters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[name][contains]",
			QueryPath: "filter",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "page",
			Usage:     "Consolidated page parameter (deepObject style). Originally: page[size], page[number]",
			QueryPath: "page",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the <code>-</code> prefix.<br/><br/>\nThat is: <ul>\n  <li>\n    <code>name</code>: sorts the result by the\n    <code>name</code> field in ascending order.\n  </li>\n\n  <li>\n    <code>-name</code>: sorts the result by the\n    <code>name</code> field in descending order.\n  </li>\n</ul> <br/>",
			Default:   "-created_at",
			QueryPath: "sort",
		},
	},
	Action:          handleOutboundVoiceProfilesList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.name",
			Usage:      "Name filtering operations",
			InnerField: "name",
		},
	},
	"page": {
		&requestflag.InnerFlag[int64]{
			Name:       "page.number",
			Usage:      "The page number to load.",
			InnerField: "number",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "page.size",
			Usage:      "The size of the page.",
			InnerField: "size",
		},
	},
})

var outboundVoiceProfilesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes an existing outbound voice profile.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleOutboundVoiceProfilesDelete,
	HideHelpCommand: true,
}

func handleOutboundVoiceProfilesCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.OutboundVoiceProfileNewParams{}

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
	_, err = client.OutboundVoiceProfiles.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "outbound-voice-profiles create", obj, format, transform)
}

func handleOutboundVoiceProfilesRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.OutboundVoiceProfiles.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "outbound-voice-profiles retrieve", obj, format, transform)
}

func handleOutboundVoiceProfilesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.OutboundVoiceProfileUpdateParams{}

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
	_, err = client.OutboundVoiceProfiles.Update(
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
	return ShowJSON(os.Stdout, "outbound-voice-profiles update", obj, format, transform)
}

func handleOutboundVoiceProfilesList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.OutboundVoiceProfileListParams{}

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
		_, err = client.OutboundVoiceProfiles.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "outbound-voice-profiles list", obj, format, transform)
	} else {
		iter := client.OutboundVoiceProfiles.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "outbound-voice-profiles list", iter, format, transform)
	}
}

func handleOutboundVoiceProfilesDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.OutboundVoiceProfiles.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "outbound-voice-profiles delete", obj, format, transform)
}
