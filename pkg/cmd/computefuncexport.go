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

var computeFuncsExportCreate = cli.Command{
	Name:    "create",
	Usage:   "Configures the external OTLP endpoint a function's runtime and/or invocation\nlogs are pushed to as they happen. This operation is a **full replace, not a\npatch**: `endpoint`, `headers`, `runtime_export_enabled`, and\n`invocation_export_enabled` are all required on every call — omitting any of\nthem is a 422, not \"keep the current value\". Headers are encrypted at rest and\nnever returned in any response.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "endpoint",
			Usage:    "HTTPS URL to push logs to",
			Required: true,
			BodyPath: "endpoint",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "headers",
			Usage:    `Headers attached to every export push, as key-value pairs (e.g. an auth token the collector expects). Required even when empty — {} means "no headers". Encrypted at rest; never returned.`,
			Required: true,
			BodyPath: "headers",
		},
		&requestflag.Flag[bool]{
			Name:     "invocation-export-enabled",
			Usage:    "Export invocation records (one per HTTP request) to this destination",
			Required: true,
			BodyPath: "invocation_export_enabled",
		},
		&requestflag.Flag[bool]{
			Name:     "runtime-export-enabled",
			Usage:    "Export runtime logs (function stdout/stderr) to this destination",
			Required: true,
			BodyPath: "runtime_export_enabled",
		},
	},
	Action:          handleComputeFuncsExportCreate,
	HideHelpCommand: true,
}

var computeFuncsExportList = cli.Command{
	Name:    "list",
	Usage:   "Returns the function's configured log export destination and which log types are\nexported. Headers are never returned. Returns 404 (error code 10005) when no\ndestination is configured for the function.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleComputeFuncsExportList,
	HideHelpCommand: true,
}

var computeFuncsExportDeleteAll = cli.Command{
	Name:    "delete-all",
	Usage:   "Stops exporting a function's logs and removes its destination configuration.\nIdempotent: deleting when nothing is configured succeeds.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleComputeFuncsExportDeleteAll,
	HideHelpCommand: true,
}

func handleComputeFuncsExportCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.ComputeFuncExportNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Compute.Funcs.Export.New(
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
		Title:          "compute:funcs:export create",
		Transform:      transform,
	})
}

func handleComputeFuncsExportList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Compute.Funcs.Export.List(ctx, cmd.Value("id").(string), options...)
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
		Title:          "compute:funcs:export list",
		Transform:      transform,
	})
}

func handleComputeFuncsExportDeleteAll(ctx context.Context, cmd *cli.Command) error {
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

	return client.Compute.Funcs.Export.DeleteAll(ctx, cmd.Value("id").(string), options...)
}
