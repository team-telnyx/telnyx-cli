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

var voiceClonesCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a new voice clone by capturing the voice identity of an existing voice\ndesign. The clone can then be used for text-to-speech synthesis.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "gender",
			Usage:    "Gender of the voice clone.",
			Required: true,
			BodyPath: "gender",
		},
		&requestflag.Flag[string]{
			Name:     "language",
			Usage:    "ISO 639-1 language code for the clone (e.g. `en`, `fr`, `de`).",
			Required: true,
			BodyPath: "language",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name for the voice clone.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "voice-design-id",
			Usage:    "UUID of the source voice design to clone.",
			Required: true,
			BodyPath: "voice_design_id",
		},
		&requestflag.Flag[string]{
			Name:     "provider",
			Usage:    "Voice synthesis provider. Case-insensitive. Defaults to `telnyx`.",
			Default:  "telnyx",
			BodyPath: "provider",
		},
	},
	Action:          handleVoiceClonesCreate,
	HideHelpCommand: true,
}

var voiceClonesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates the name, language, or gender of a voice clone.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "New name for the voice clone.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "gender",
			Usage:    "Updated gender for the voice clone.",
			BodyPath: "gender",
		},
		&requestflag.Flag[string]{
			Name:     "language",
			Usage:    "Updated ISO 639-1 language code or `auto`.",
			BodyPath: "language",
		},
	},
	Action:          handleVoiceClonesUpdate,
	HideHelpCommand: true,
}

var voiceClonesList = cli.Command{
	Name:    "list",
	Usage:   "Returns a paginated list of voice clones belonging to the authenticated account.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "filter-name",
			Usage:     "Case-insensitive substring filter on the name field.",
			QueryPath: "filter[name]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-provider",
			Usage:     "Filter by voice synthesis provider. Case-insensitive.",
			QueryPath: "filter[provider]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "Page number for pagination (1-based).",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Number of results per page.",
			Default:   20,
			QueryPath: "page[size]",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Sort order. Prefix with `-` for descending. Defaults to `-created_at`.",
			Default:   "-created_at",
			QueryPath: "sort",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleVoiceClonesList,
	HideHelpCommand: true,
}

var voiceClonesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Permanently deletes a voice clone. This action cannot be undone.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleVoiceClonesDelete,
	HideHelpCommand: true,
}

var voiceClonesCreateFromUpload = cli.Command{
	Name:    "create-from-upload",
	Usage:   "Creates a new voice clone by uploading an audio file directly. Supported\nformats: WAV, MP3, FLAC, OGG, M4A. For best results, provide 5–10 seconds of\nclear speech. Maximum file size: 2MB.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "audio-file",
			Usage:     "Audio file to clone the voice from. Supported formats: WAV, MP3, FLAC, OGG, M4A. For best quality, provide 5–10 seconds of clear, uninterrupted speech. Maximum size: 5MB for Telnyx, 20MB for Minimax.",
			Required:  true,
			BodyPath:  "audio_file",
			FileInput: true,
		},
		&requestflag.Flag[string]{
			Name:     "language",
			Usage:    "ISO 639-1 language code (e.g. `en`, `fr`) or `auto` for automatic detection.",
			Required: true,
			BodyPath: "language",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name for the voice clone.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "gender",
			Usage:    "Gender of the voice clone.",
			BodyPath: "gender",
		},
		&requestflag.Flag[string]{
			Name:     "label",
			Usage:    "Optional custom label describing the voice style. If omitted, falls back to the source design's prompt text.",
			BodyPath: "label",
		},
		&requestflag.Flag[string]{
			Name:     "provider",
			Usage:    "Voice synthesis provider. Case-insensitive. Defaults to `telnyx`.",
			Default:  "telnyx",
			BodyPath: "provider",
		},
		&requestflag.Flag[string]{
			Name:     "ref-text",
			Usage:    "Optional transcript of the audio file. Providing this improves clone quality.",
			BodyPath: "ref_text",
		},
	},
	Action:          handleVoiceClonesCreateFromUpload,
	HideHelpCommand: true,
}

var voiceClonesDownloadSample = cli.Command{
	Name:    "download-sample",
	Usage:   "Downloads the WAV audio sample that was used to create the voice clone.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "The file where the response contents will be stored. Use the value '-' to force output to stdout.",
		},
	},
	Action:          handleVoiceClonesDownloadSample,
	HideHelpCommand: true,
}

func handleVoiceClonesCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.VoiceCloneNewParams{}

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
	_, err = client.VoiceClones.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "voice-clones create", obj, format, transform)
}

func handleVoiceClonesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.VoiceCloneUpdateParams{}

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
	_, err = client.VoiceClones.Update(
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
	return ShowJSON(os.Stdout, "voice-clones update", obj, format, transform)
}

func handleVoiceClonesList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.VoiceCloneListParams{}

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
		_, err = client.VoiceClones.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "voice-clones list", obj, format, transform)
	} else {
		iter := client.VoiceClones.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "voice-clones list", iter, format, transform, maxItems)
	}
}

func handleVoiceClonesDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.VoiceClones.Delete(ctx, cmd.Value("id").(string), options...)
}

func handleVoiceClonesCreateFromUpload(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.VoiceCloneNewFromUploadParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		MultipartFormEncoded,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.VoiceClones.NewFromUpload(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "voice-clones create-from-upload", obj, format, transform)
}

func handleVoiceClonesDownloadSample(ctx context.Context, cmd *cli.Command) error {
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

	response, err := client.VoiceClones.DownloadSample(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}
	message, err := writeBinaryResponse(response, cmd.String("output"))
	if message != "" {
		fmt.Println(message)
	}
	return err
}
