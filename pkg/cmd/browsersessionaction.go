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

var browserSessionsActionsCreate = cli.Command{
	Name:    "create",
	Usage:   "Executes a browser action on the session. Returns 200 if completed within\nwait_seconds, 202 if still in-flight.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Allowed values: "navigate".`,
			Default:  "navigate",
			Const:    true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			BodyPath: "url",
		},
		&requestflag.Flag[int64]{
			Name:      "wait-seconds",
			Usage:     "Maximum seconds the submission request waits for completion, not an execution timeout. If still queued or running, returns 202 without cancelling execution; poll the returned action ID.",
			Default:   10,
			QueryPath: "wait_seconds",
		},
		&requestflag.Flag[string]{
			Name:     "idempotency-key",
			BodyPath: "idempotency_key",
		},
		&requestflag.Flag[bool]{
			Name:     "include-observation",
			Default:  false,
			BodyPath: "include_observation",
		},
		&requestflag.Flag[string]{
			Name:     "wait-until",
			Usage:    `Allowed values: "load", "dom_stable", "none".`,
			Default:  "load",
			BodyPath: "wait_until",
		},
		&requestflag.Flag[string]{
			Name:     "observation-id",
			BodyPath: "observation_id",
		},
		&requestflag.Flag[string]{
			Name:     "ref",
			BodyPath: "ref",
		},
		&requestflag.Flag[int64]{
			Name:     "timeout-ms",
			Default:  5000,
			BodyPath: "timeout_ms",
		},
		&requestflag.Flag[string]{
			Name:     "wait",
			Usage:    `Allowed values: "no_wait", "dom_stable", "url_change".`,
			Default:  "dom_stable",
			BodyPath: "wait",
		},
		&requestflag.Flag[string]{
			Name:     "value",
			BodyPath: "value",
		},
		&requestflag.Flag[string]{
			Name:     "key",
			Usage:    `Allowed values: "Enter", "Tab", "Escape", "Backspace", "ArrowUp", "ArrowDown", "ArrowLeft", "ArrowRight", "PageUp", "PageDown", "Home", "End".`,
			BodyPath: "key",
		},
		&requestflag.Flag[int64]{
			Name:     "delta-x",
			Default:  0,
			BodyPath: "delta_x",
		},
		&requestflag.Flag[int64]{
			Name:     "delta-y",
			Default:  0,
			BodyPath: "delta_y",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "condition",
			BodyPath: "condition",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "action",
			BodyPath: "actions",
		},
	},
	Action:          handleBrowserSessionsActionsCreate,
	HideHelpCommand: true,
}

var browserSessionsActionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns the stored status and result of an action. Use this to poll actions that\nreturned 202; HTTP 200 can contain any action status. Stored observations\ncontain metadata with empty tree and refs, not the full snapshot; submit a fresh\nobserve action to obtain current refs.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "action-id",
			Required:  true,
			PathParam: "actionId",
		},
	},
	Action:          handleBrowserSessionsActionsRetrieve,
	HideHelpCommand: true,
}

func handleBrowserSessionsActionsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.BrowserSessionActionNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.BrowserSessions.Actions.New(
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
		Title:          "browser-sessions:actions create",
		Transform:      transform,
	})
}

func handleBrowserSessionsActionsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("action-id") && len(unusedArgs) > 0 {
		cmd.Set("action-id", unusedArgs[0])
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

	params := telnyx.BrowserSessionActionGetParams{
		ID: cmd.Value("id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.BrowserSessions.Actions.Get(
		ctx,
		cmd.Value("action-id").(string),
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
		Title:          "browser-sessions:actions retrieve",
		Transform:      transform,
	})
}
