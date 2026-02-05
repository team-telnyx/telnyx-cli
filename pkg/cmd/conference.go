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

var conferencesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a conference from an existing call leg using a `call_control_id` and a\nconference name. Upon creating the conference, the call will be automatically\nbridged to the conference. Conferences will expire after all participants have\nleft the conference or after 4 hours regardless of the number of active\nparticipants.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-control-id",
			Usage:    "Unique identifier and token for controlling the call",
			Required: true,
			BodyPath: "call_control_id",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the conference",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "beep-enabled",
			Usage:    "Whether a beep sound should be played when participants join and/or leave the conference.",
			Default:  "never",
			BodyPath: "beep_enabled",
		},
		&requestflag.Flag[string]{
			Name:     "client-state",
			Usage:    "Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. The client_state will be updated for the creator call leg and will be used for all webhooks related to the created conference.",
			BodyPath: "client_state",
		},
		&requestflag.Flag[bool]{
			Name:     "comfort-noise",
			Usage:    "Toggle background comfort noise.",
			Default:  true,
			BodyPath: "comfort_noise",
		},
		&requestflag.Flag[string]{
			Name:     "command-id",
			Usage:    "Use this field to avoid execution of duplicate commands. Telnyx will ignore subsequent commands with the same `command_id` as one that has already been executed.",
			BodyPath: "command_id",
		},
		&requestflag.Flag[int64]{
			Name:     "duration-minutes",
			Usage:    "Time length (minutes) after which the conference will end.",
			BodyPath: "duration_minutes",
		},
		&requestflag.Flag[string]{
			Name:     "hold-audio-url",
			Usage:    `The URL of a file to be played to participants joining the conference. The URL can point to either a WAV or MP3 file. hold_media_name and hold_audio_url cannot be used together in one request. Takes effect only when "start_conference_on_create" is set to "false".`,
			BodyPath: "hold_audio_url",
		},
		&requestflag.Flag[string]{
			Name:     "hold-media-name",
			Usage:    `The media_name of a file to be played to participants joining the conference. The media_name must point to a file previously uploaded to api.telnyx.com/v2/media by the same user/organization. The file must either be a WAV or MP3 file. Takes effect only when "start_conference_on_create" is set to "false".`,
			BodyPath: "hold_media_name",
		},
		&requestflag.Flag[int64]{
			Name:     "max-participants",
			Usage:    "The maximum number of active conference participants to allow. Must be between 2 and 800. Defaults to 250",
			BodyPath: "max_participants",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Sets the region where the conference data will be hosted. Defaults to the region defined in user's data locality settings (Europe or US).",
			BodyPath: "region",
		},
		&requestflag.Flag[bool]{
			Name:     "start-conference-on-create",
			Usage:    `Whether the conference should be started on creation. If the conference isn't started all participants that join are automatically put on hold. Defaults to "true".`,
			BodyPath: "start_conference_on_create",
		},
	},
	Action:          handleConferencesCreate,
	HideHelpCommand: true,
}

var conferencesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve an existing conference",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "region",
			Usage:     "Region where the conference data is located",
			QueryPath: "region",
		},
	},
	Action:          handleConferencesRetrieve,
	HideHelpCommand: true,
}

var conferencesList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Lists conferences. Conferences are created on demand, and will expire after all\nparticipants have left the conference or after 4 hours regardless of the number\nof active participants. Conferences are listed in descending order by\n`expires_at`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[application_name][contains], filter[outbound.outbound_voice_profile_id], filter[leg_id], filter[application_session_id], filter[connection_id], filter[product], filter[failed], filter[from], filter[to], filter[name], filter[type], filter[occurred_at][eq/gt/gte/lt/lte], filter[status]",
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
		&requestflag.Flag[string]{
			Name:      "region",
			Usage:     "Region where the conference data is located",
			QueryPath: "region",
		},
	},
	Action:          handleConferencesList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.application-name",
			Usage:      "Application name filters",
			InnerField: "application_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.application-session-id",
			Usage:      "The unique identifier of the call session. A session may include multiple call leg events.",
			InnerField: "application_session_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.connection-id",
			Usage:      "The unique identifier of the conection.",
			InnerField: "connection_id",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "filter.failed",
			Usage:      "Delivery failed or not.",
			InnerField: "failed",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.from",
			Usage:      "Filter by From number.",
			InnerField: "from",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.leg-id",
			Usage:      "The unique identifier of an individual call leg.",
			InnerField: "leg_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.name",
			Usage:      "If present, conferences will be filtered to those with a matching `name` attribute. Matching is case-sensitive",
			InnerField: "name",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.occurred-at",
			Usage:      "Event occurred_at filters",
			InnerField: "occurred_at",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.outbound-outbound-voice-profile-id",
			Usage:      "Identifies the associated outbound voice profile.",
			InnerField: "outbound.outbound_voice_profile_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.product",
			Usage:      "Filter by product.",
			InnerField: "product",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.status",
			Usage:      "If present, conferences will be filtered by status.",
			InnerField: "status",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.to",
			Usage:      "Filter by To number.",
			InnerField: "to",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.type",
			Usage:      "Event type",
			InnerField: "type",
		},
	},
})

var conferencesListParticipants = requestflag.WithInnerFlags(cli.Command{
	Name:    "list-participants",
	Usage:   "Lists conference participants",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "conference-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[muted], filter[on_hold], filter[whispering]",
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
		&requestflag.Flag[string]{
			Name:      "region",
			Usage:     "Region where the conference data is located",
			QueryPath: "region",
		},
	},
	Action:          handleConferencesListParticipants,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[bool]{
			Name:       "filter.muted",
			Usage:      "If present, participants will be filtered to those who are/are not muted",
			InnerField: "muted",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "filter.on-hold",
			Usage:      "If present, participants will be filtered to those who are/are not put on hold",
			InnerField: "on_hold",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "filter.whispering",
			Usage:      "If present, participants will be filtered to those who are whispering or are not",
			InnerField: "whispering",
		},
	},
})

func handleConferencesCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ConferenceNewParams{}

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
	_, err = client.Conferences.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "conferences create", obj, format, transform)
}

func handleConferencesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ConferenceGetParams{}

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
	_, err = client.Conferences.Get(
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
	return ShowJSON(os.Stdout, "conferences retrieve", obj, format, transform)
}

func handleConferencesList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ConferenceListParams{}

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
		_, err = client.Conferences.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "conferences list", obj, format, transform)
	} else {
		iter := client.Conferences.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "conferences list", iter, format, transform)
	}
}

func handleConferencesListParticipants(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("conference-id") && len(unusedArgs) > 0 {
		cmd.Set("conference-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ConferenceListParticipantsParams{}

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
		_, err = client.Conferences.ListParticipants(
			ctx,
			cmd.Value("conference-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "conferences list-participants", obj, format, transform)
	} else {
		iter := client.Conferences.ListParticipantsAutoPaging(
			ctx,
			cmd.Value("conference-id").(string),
			params,
			options...,
		)
		return ShowJSONIterator(os.Stdout, "conferences list-participants", iter, format, transform)
	}
}
