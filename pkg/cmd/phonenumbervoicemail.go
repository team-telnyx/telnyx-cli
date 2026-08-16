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

var phoneNumbersVoicemailCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create voicemail settings for a phone number. You can also configure a custom\ngreeting by setting the `greeting` object: use `mode` `custom_greeting` together\nwith a `media_name` that points to an audio file uploaded through the Media\nStorage API, or `mode` `default` to use the standard system greeting.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "phone-number-id",
			Required:  true,
			PathParam: "phone_number_id",
		},
		&requestflag.Flag[bool]{
			Name:     "enabled",
			Usage:    "Whether voicemail is enabled.",
			BodyPath: "enabled",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "greeting",
			Usage:    "Controls the greeting a caller hears before leaving a voicemail. Set `mode` to `default` to play the standard system greeting, or to `custom_greeting` to play your own audio. When `mode` is `custom_greeting`, `media_name` is required and must reference an audio file already uploaded to your account through the Media Storage API.",
			BodyPath: "greeting",
		},
		&requestflag.Flag[string]{
			Name:     "pin",
			Usage:    "The pin used for voicemail",
			BodyPath: "pin",
		},
	},
	Action:          handlePhoneNumbersVoicemailCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"greeting": {
		&requestflag.InnerFlag[*string]{
			Name:       "greeting.media-name",
			Usage:      "The name of the media file to play as the greeting. Required when `mode` is `custom_greeting`; ignored when `mode` is `default`. The value must match the `media_name` of a file you previously uploaded with the Media Storage API (`POST /v2/media`).",
			InnerField: "media_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "greeting.mode",
			Usage:      "The greeting mode. `default` plays the standard system greeting. `custom_greeting` plays the audio referenced by `media_name`.",
			InnerField: "mode",
		},
	},
})

var phoneNumbersVoicemailRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns the voicemail settings for a phone number",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "phone-number-id",
			Required:  true,
			PathParam: "phone_number_id",
		},
	},
	Action:          handlePhoneNumbersVoicemailRetrieve,
	HideHelpCommand: true,
}

var phoneNumbersVoicemailUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update voicemail settings for a phone number. You can also configure a custom\ngreeting by setting the `greeting` object: use `mode` `custom_greeting` together\nwith a `media_name` that points to an audio file uploaded through the Media\nStorage API, or `mode` `default` to use the standard system greeting.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "phone-number-id",
			Required:  true,
			PathParam: "phone_number_id",
		},
		&requestflag.Flag[bool]{
			Name:     "enabled",
			Usage:    "Whether voicemail is enabled.",
			BodyPath: "enabled",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "greeting",
			Usage:    "Controls the greeting a caller hears before leaving a voicemail. Set `mode` to `default` to play the standard system greeting, or to `custom_greeting` to play your own audio. When `mode` is `custom_greeting`, `media_name` is required and must reference an audio file already uploaded to your account through the Media Storage API.",
			BodyPath: "greeting",
		},
		&requestflag.Flag[string]{
			Name:     "pin",
			Usage:    "The pin used for voicemail",
			BodyPath: "pin",
		},
	},
	Action:          handlePhoneNumbersVoicemailUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"greeting": {
		&requestflag.InnerFlag[*string]{
			Name:       "greeting.media-name",
			Usage:      "The name of the media file to play as the greeting. Required when `mode` is `custom_greeting`; ignored when `mode` is `default`. The value must match the `media_name` of a file you previously uploaded with the Media Storage API (`POST /v2/media`).",
			InnerField: "media_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "greeting.mode",
			Usage:      "The greeting mode. `default` plays the standard system greeting. `custom_greeting` plays the audio referenced by `media_name`.",
			InnerField: "mode",
		},
	},
})

func handlePhoneNumbersVoicemailCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("phone-number-id") && len(unusedArgs) > 0 {
		cmd.Set("phone-number-id", unusedArgs[0])
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

	params := telnyx.PhoneNumberVoicemailNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.PhoneNumbers.Voicemail.New(
		ctx,
		cmd.Value("phone-number-id").(string),
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
		Title:          "phone-numbers:voicemail create",
		Transform:      transform,
	})
}

func handlePhoneNumbersVoicemailRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("phone-number-id") && len(unusedArgs) > 0 {
		cmd.Set("phone-number-id", unusedArgs[0])
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
	_, err = client.PhoneNumbers.Voicemail.Get(ctx, cmd.Value("phone-number-id").(string), options...)
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
		Title:          "phone-numbers:voicemail retrieve",
		Transform:      transform,
	})
}

func handlePhoneNumbersVoicemailUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("phone-number-id") && len(unusedArgs) > 0 {
		cmd.Set("phone-number-id", unusedArgs[0])
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

	params := telnyx.PhoneNumberVoicemailUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.PhoneNumbers.Voicemail.Update(
		ctx,
		cmd.Value("phone-number-id").(string),
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
		Title:          "phone-numbers:voicemail update",
		Transform:      transform,
	})
}
