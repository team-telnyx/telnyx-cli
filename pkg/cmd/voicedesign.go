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

var voiceDesignsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a new voice design (version 1) when `voice_design_id` is omitted. When\n`voice_design_id` is provided, adds a new version to the existing design\ninstead. A design can have at most 50 versions.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "prompt",
			Usage:    "Natural language description of the voice style, e.g. 'Speak in a warm, friendly tone with a slight British accent'.",
			Required: true,
			BodyPath: "prompt",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "Sample text to synthesize for this voice design.",
			Required: true,
			BodyPath: "text",
		},
		&requestflag.Flag[string]{
			Name:     "language",
			Usage:    "Language for synthesis. Supported values: Auto, Chinese, English, Japanese, Korean, German, French, Russian, Portuguese, Spanish, Italian. Defaults to Auto.",
			Default:  "Auto",
			BodyPath: "language",
		},
		&requestflag.Flag[int64]{
			Name:     "max-new-tokens",
			Usage:    "Maximum number of tokens to generate. Default: 2048.",
			BodyPath: "max_new_tokens",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name for the voice design. Required when creating a new design (`voice_design_id` is not provided); ignored when adding a version. Cannot be a UUID.",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "provider",
			Usage:    "Voice synthesis provider. `telnyx` uses the Qwen3TTS model; `minimax` uses the Minimax speech models. Case-insensitive. Defaults to `telnyx`.",
			Default:  "telnyx",
			BodyPath: "provider",
		},
		&requestflag.Flag[float64]{
			Name:     "repetition-penalty",
			Usage:    "Repetition penalty to reduce repeated patterns in generated audio. Default: 1.05.",
			BodyPath: "repetition_penalty",
		},
		&requestflag.Flag[float64]{
			Name:     "temperature",
			Usage:    "Sampling temperature controlling randomness. Higher values produce more varied output. Default: 0.9.",
			BodyPath: "temperature",
		},
		&requestflag.Flag[int64]{
			Name:     "top-k",
			Usage:    "Top-k sampling parameter — limits the token vocabulary considered at each step. Default: 50.",
			BodyPath: "top_k",
		},
		&requestflag.Flag[float64]{
			Name:     "top-p",
			Usage:    "Top-p (nucleus) sampling parameter — cumulative probability cutoff for token selection. Default: 1.0.",
			BodyPath: "top_p",
		},
		&requestflag.Flag[string]{
			Name:     "voice-design-id",
			Usage:    "ID of an existing voice design to add a new version to. When provided, a new version is created instead of a new design.",
			BodyPath: "voice_design_id",
		},
	},
	Action:          handleVoiceDesignsCreate,
	HideHelpCommand: true,
}

var voiceDesignsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns the latest version of a voice design, or a specific version when\n`?version=N` is provided. The `id` parameter accepts either a UUID or the design\nname.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:      "version",
			Usage:     "Specific version number to retrieve. Defaults to the latest version.",
			QueryPath: "version",
		},
	},
	Action:          handleVoiceDesignsRetrieve,
	HideHelpCommand: true,
}

var voiceDesignsList = cli.Command{
	Name:    "list",
	Usage:   "Returns a paginated list of voice designs belonging to the authenticated\naccount.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "filter-name",
			Usage:     "Case-insensitive substring filter on the name field.",
			QueryPath: "filter[name]",
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
	Action:          handleVoiceDesignsList,
	HideHelpCommand: true,
}

var voiceDesignsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Permanently deletes a voice design and all of its versions. This action cannot\nbe undone.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleVoiceDesignsDelete,
	HideHelpCommand: true,
}

var voiceDesignsDeleteVersion = cli.Command{
	Name:    "delete-version",
	Usage:   "Permanently deletes a specific version of a voice design. The version number\nmust be a positive integer.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:     "version",
			Required: true,
		},
	},
	Action:          handleVoiceDesignsDeleteVersion,
	HideHelpCommand: true,
}

var voiceDesignsDownloadSample = cli.Command{
	Name:    "download-sample",
	Usage:   "Downloads the WAV audio sample for the voice design. Returns the latest\nversion's sample by default, or a specific version when `?version=N` is\nprovided. The `id` parameter accepts either a UUID or the design name.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:      "version",
			Usage:     "Specific version number to download the sample for. Defaults to the latest version.",
			QueryPath: "version",
		},
		&requestflag.Flag[string]{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "The file where the response contents will be stored. Use the value '-' to force output to stdout.",
		},
	},
	Action:          handleVoiceDesignsDownloadSample,
	HideHelpCommand: true,
}

var voiceDesignsRename = cli.Command{
	Name:    "rename",
	Usage:   "Updates the name of a voice design. All versions retain their other properties.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "New name for the voice design.",
			Required: true,
			BodyPath: "name",
		},
	},
	Action:          handleVoiceDesignsRename,
	HideHelpCommand: true,
}

func handleVoiceDesignsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.VoiceDesignNewParams{}

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
	_, err = client.VoiceDesigns.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "voice-designs create", obj, format, transform)
}

func handleVoiceDesignsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.VoiceDesignGetParams{}

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
	_, err = client.VoiceDesigns.Get(
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
	return ShowJSON(os.Stdout, "voice-designs retrieve", obj, format, transform)
}

func handleVoiceDesignsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.VoiceDesignListParams{}

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
		_, err = client.VoiceDesigns.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "voice-designs list", obj, format, transform)
	} else {
		iter := client.VoiceDesigns.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "voice-designs list", iter, format, transform, maxItems)
	}
}

func handleVoiceDesignsDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.VoiceDesigns.Delete(ctx, cmd.Value("id").(string), options...)
}

func handleVoiceDesignsDeleteVersion(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("version") && len(unusedArgs) > 0 {
		cmd.Set("version", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.VoiceDesignDeleteVersionParams{
		ID: cmd.Value("id").(string),
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

	return client.VoiceDesigns.DeleteVersion(
		ctx,
		cmd.Value("version").(int64),
		params,
		options...,
	)
}

func handleVoiceDesignsDownloadSample(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.VoiceDesignDownloadSampleParams{}

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

	response, err := client.VoiceDesigns.DownloadSample(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}
	message, err := writeBinaryResponse(response, os.Stdout, cmd.String("output"))
	if message != "" {
		fmt.Println(message)
	}
	return err
}

func handleVoiceDesignsRename(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.VoiceDesignRenameParams{}

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
	_, err = client.VoiceDesigns.Rename(
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
	return ShowJSON(os.Stdout, "voice-designs rename", obj, format, transform)
}
