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

var recordingsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves the details of an existing call recording.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "recording-id",
			Required: true,
		},
	},
	Action:          handleRecordingsRetrieve,
	HideHelpCommand: true,
}

var recordingsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Returns a list of your call recordings.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[conference_id], filter[created_at][gte], filter[created_at][lte], filter[call_leg_id], filter[call_session_id], filter[from], filter[to], filter[connection_id], filter[sip_call_id]",
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
	},
	Action:          handleRecordingsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.call-leg-id",
			Usage:      "If present, recordings will be filtered to those with a matching call_leg_id.",
			InnerField: "call_leg_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.call-session-id",
			Usage:      "If present, recordings will be filtered to those with a matching call_session_id.",
			InnerField: "call_session_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.conference-id",
			Usage:      "Returns only recordings associated with a given conference.",
			InnerField: "conference_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.connection-id",
			Usage:      "If present, recordings will be filtered to those with a matching `connection_id` attribute (case-sensitive).",
			InnerField: "connection_id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.created-at",
			InnerField: "created_at",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.from",
			Usage:      "If present, recordings will be filtered to those with a matching `from` attribute (case-sensitive).",
			InnerField: "from",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.sip-call-id",
			Usage:      "If present, recordings will be filtered to those with a matching `sip_call_id` attribute. Matching is case-sensitive",
			InnerField: "sip_call_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.to",
			Usage:      "If present, recordings will be filtered to those with a matching `to` attribute (case-sensitive).",
			InnerField: "to",
		},
	},
})

var recordingsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Permanently deletes a call recording.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "recording-id",
			Required: true,
		},
	},
	Action:          handleRecordingsDelete,
	HideHelpCommand: true,
}

func handleRecordingsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("recording-id") && len(unusedArgs) > 0 {
		cmd.Set("recording-id", unusedArgs[0])
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
	_, err = client.Recordings.Get(ctx, cmd.Value("recording-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "recordings retrieve", obj, format, transform)
}

func handleRecordingsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.RecordingListParams{}

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
		_, err = client.Recordings.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "recordings list", obj, format, transform)
	} else {
		iter := client.Recordings.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "recordings list", iter, format, transform)
	}
}

func handleRecordingsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("recording-id") && len(unusedArgs) > 0 {
		cmd.Set("recording-id", unusedArgs[0])
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
	_, err = client.Recordings.Delete(ctx, cmd.Value("recording-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "recordings delete", obj, format, transform)
}
