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

var browserSessionsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates an isolated browser session with network access restricted to the\ndomains you allow.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "policy",
			Required: true,
			BodyPath: "policy",
		},
		&requestflag.Flag[string]{
			Name:     "idempotency-key",
			Usage:    "Scoped to the organization. The same key and identical normalized policy, limits, and metadata replay the existing session (HTTP 201); a different request returns 409 idempotency_conflict. Replays do not restart terminal sessions.",
			BodyPath: "idempotency_key",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "limits",
			BodyPath: "limits",
		},
		&requestflag.Flag[any]{
			Name:     "metadata",
			Usage:    "Arbitrary key-value metadata.",
			BodyPath: "metadata",
		},
	},
	Action:          handleBrowserSessionsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"policy": {
		&requestflag.InnerFlag[[]string]{
			Name:       "policy.allowed-domains",
			Usage:      "Domain patterns the browser can navigate to. Exact match (example.com) or wildcard (*.example.com). IP literals not allowed.",
			InnerField: "allowed_domains",
		},
	},
	"limits": {
		&requestflag.InnerFlag[int64]{
			Name:       "limits.idle-timeout-s",
			Usage:      "Idle timeout in seconds. Must not exceed lifetime_s.",
			InnerField: "idle_timeout_s",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "limits.lifetime-s",
			Usage:      "Max session duration in seconds (60-1800).",
			InnerField: "lifetime_s",
		},
	},
})

var browserSessionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns the current status of a browser session.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleBrowserSessionsRetrieve,
	HideHelpCommand: true,
}

var browserSessionsList = cli.Command{
	Name:    "list",
	Usage:   "Lists sessions for the authenticated org. Optional status filter.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     `Allowed values: "provisioning", "ready", "closing", "closed", "failed", "expired".`,
			QueryPath: "status",
		},
	},
	Action:          handleBrowserSessionsList,
	HideHelpCommand: true,
}

var browserSessionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Closes a ready session and releases its resources. An already closed, failed, or\nexpired session is returned unchanged; a provisioning session returns 409\nsession_not_ready.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleBrowserSessionsDelete,
	HideHelpCommand: true,
}

var browserSessionsRetrieveArtifact = cli.Command{
	Name:    "retrieve-artifact",
	Usage:   "Returns artifact metadata and a signed download URL. The download URL requires\nBearer auth.",
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
			PathParam: "artifactId",
		},
	},
	Action:          handleBrowserSessionsRetrieveArtifact,
	HideHelpCommand: true,
}

func handleBrowserSessionsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.BrowserSessionNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.BrowserSessions.New(ctx, params, options...)
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
		Title:          "browser-sessions create",
		Transform:      transform,
	})
}

func handleBrowserSessionsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.BrowserSessions.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "browser-sessions retrieve",
		Transform:      transform,
	})
}

func handleBrowserSessionsList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.BrowserSessionListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.BrowserSessions.List(ctx, params, options...)
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
		Title:          "browser-sessions list",
		Transform:      transform,
	})
}

func handleBrowserSessionsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.BrowserSessions.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "browser-sessions delete",
		Transform:      transform,
	})
}

func handleBrowserSessionsRetrieveArtifact(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.BrowserSessionGetArtifactParams{
		ID: cmd.Value("id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.BrowserSessions.GetArtifact(
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
		Title:          "browser-sessions retrieve-artifact",
		Transform:      transform,
	})
}
