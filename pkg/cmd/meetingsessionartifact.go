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

var meetingSessionsArtifactsCreate = cli.Command{
	Name:    "create",
	Usage:   "Requests asynchronous generation of one artifact: `summary`, `action_items`,\n`decisions`, `topics`, `open_questions`, or `custom`. Each request produces one\nartifact. `custom` is answered from a `prompt` you supply, which is required for\n`custom` and rejected on the five named types. Generation requires transcript\ncontent and configured inference and currently reads at most the first 10,000\nsegments, so exceptionally long transcripts may produce incomplete artifacts or\nfail model limits. **Not idempotent, and every call is billed**: each request is\na separate inference run, so a retry or a duplicate POST produces a second\nartifact and a second charge. Guard the call rather than relying on the service\nto collapse it. The automatic `summarize_on_end` attempt is billed on the same\nbasis.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "What to generate from the transcript. `custom` is answered from a `prompt` you supply; the five named types need none.",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "prompt",
			Usage:    "An open-ended request answered from the transcript. Required when `type` is `custom`, and rejected with 400 on any named type. Trimmed before storage and echoed back in artifact responses and the `artifact.completed` webhook.",
			BodyPath: "prompt",
		},
	},
	Action:          handleMeetingSessionsArtifactsCreate,
	HideHelpCommand: true,
}

var meetingSessionsArtifactsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a single meeting session artifact by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "artifact-id",
			Required:  true,
			PathParam: "artifact_id",
		},
	},
	Action:          handleMeetingSessionsArtifactsRetrieve,
	HideHelpCommand: true,
}

var meetingSessionsArtifactsList = cli.Command{
	Name:    "list",
	Usage:   "Returns a list of artifacts for a meeting session.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleMeetingSessionsArtifactsList,
	HideHelpCommand: true,
}

func handleMeetingSessionsArtifactsCreate(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := telnyx.MeetingSessionArtifactNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.MeetingSessions.Artifacts.New(
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "meeting-sessions:artifacts create",
		Transform:      transform,
	})
}

func handleMeetingSessionsArtifactsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("artifact-id") && len(unusedArgs) > 0 {
		cmd.Set("artifact-id", unusedArgs[0])
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

	params := telnyx.MeetingSessionArtifactGetParams{
		ID: cmd.Value("id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.MeetingSessions.Artifacts.Get(
		ctx,
		cmd.Value("artifact-id").(string),
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
		Title:          "meeting-sessions:artifacts retrieve",
		Transform:      transform,
	})
}

func handleMeetingSessionsArtifactsList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.MeetingSessions.Artifacts.List(ctx, cmd.Value("id").(string), options...)
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
		Title:          "meeting-sessions:artifacts list",
		Transform:      transform,
	})
}
