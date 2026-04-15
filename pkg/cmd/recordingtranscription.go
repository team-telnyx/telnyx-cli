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

var recordingTranscriptionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves the details of an existing recording transcription.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "recording-transcription-id",
			Required: true,
		},
	},
	Action:          handleRecordingTranscriptionsRetrieve,
	HideHelpCommand: true,
}

var recordingTranscriptionsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Returns a list of your recording transcriptions.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Filter recording transcriptions by various attributes.",
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
	Action:          handleRecordingTranscriptionsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.created-at",
			InnerField: "created_at",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.recording-id",
			Usage:      "If present, transcriptions will be filtered to those associated with the given recording.",
			InnerField: "recording_id",
		},
	},
})

var recordingTranscriptionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Permanently deletes a recording transcription.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "recording-transcription-id",
			Required: true,
		},
	},
	Action:          handleRecordingTranscriptionsDelete,
	HideHelpCommand: true,
}

func handleRecordingTranscriptionsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("recording-transcription-id") && len(unusedArgs) > 0 {
		cmd.Set("recording-transcription-id", unusedArgs[0])
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
	_, err = client.RecordingTranscriptions.Get(ctx, cmd.Value("recording-transcription-id").(string), options...)
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
		Title:          "recording-transcriptions retrieve",
		Transform:      transform,
	})
}

func handleRecordingTranscriptionsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.RecordingTranscriptionListParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.RecordingTranscriptions.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "recording-transcriptions list",
			Transform:      transform,
		})
	} else {
		iter := client.RecordingTranscriptions.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "recording-transcriptions list",
			Transform:      transform,
		})
	}
}

func handleRecordingTranscriptionsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("recording-transcription-id") && len(unusedArgs) > 0 {
		cmd.Set("recording-transcription-id", unusedArgs[0])
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
	_, err = client.RecordingTranscriptions.Delete(ctx, cmd.Value("recording-transcription-id").(string), options...)
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
		Title:          "recording-transcriptions delete",
		Transform:      transform,
	})
}
