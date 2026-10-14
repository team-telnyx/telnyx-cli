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

var aiMemoryNamespacesProfilesSourcesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "One source and its content, as it was stored: an ingested session's payload or a\nremembered fact. A source whose ingest is still queued answers 404 until it has\nbeen stored.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "namespace",
			Usage:     "The namespace. `default` exists for every organization.",
			Required:  true,
			PathParam: "namespace",
		},
		&requestflag.Flag[string]{
			Name:      "profile-id",
			Usage:     "The profile: your identifier for the user, caller or agent this memory is about.",
			Required:  true,
			PathParam: "profile_id",
		},
		&requestflag.Flag[string]{
			Name:      "source-id",
			Usage:     "Identifies one source within its profile: an ingested session, or one remembered fact. Returned by `ingest` and `remember` when the write is accepted. Re-ingesting a session keeps its source id.",
			Required:  true,
			PathParam: "source_id",
		},
	},
	Action:          handleAIMemoryNamespacesProfilesSourcesRetrieve,
	HideHelpCommand: true,
}

var aiMemoryNamespacesProfilesSourcesList = cli.Command{
	Name:    "list",
	Usage:   "Everything a profile has stored and extracts memories from: each ingested\nsession, and each remembered fact, which has no session. Content is not listed;\nread one source for it. A source whose ingest is still queued is not here yet.\nRe-ingesting a session moves it to the front, so a listing paged while sessions\nare written can repeat or miss one at a page boundary. A `session_id` narrows\nthe listing to the source that session was stored as: one source or none, and\nnone -- an empty page, not a 404 -- for a session never ingested, still queued,\nor another profile's.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "namespace",
			Required:  true,
			PathParam: "namespace",
		},
		&requestflag.Flag[string]{
			Name:      "profile-id",
			Required:  true,
			PathParam: "profile_id",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "The page to return, counting from 1. Bounded in depth: (page[number] - 1) * page[size] may be at most 10000.",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "How many results a page holds.",
			Default:   20,
			QueryPath: "page[size]",
		},
		&requestflag.Flag[*string]{
			Name:      "session-id",
			Usage:     "An ingested session, by the `session_id` it was ingested with. Narrows the request to the source that session was stored as.",
			QueryPath: "session_id",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleAIMemoryNamespacesProfilesSourcesList,
	HideHelpCommand: true,
}

var aiMemoryNamespacesProfilesSourcesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes one source -- an ingested session or a remembered fact -- together with\nthe memories derived from it. A memory derived from this source and others is\ndeleted too, and derived again from what remains in the background. It answers\nonly once the source is gone. A source that is not there -- never stored,\nanother profile's, or already deleted -- answers 404, so on a `502` or a `504`\nrepeat the identical request and read a 404 as done. An ingest of the same\nsession that is still queued is not cancelled, and stores the session again when\nit runs. Nothing here can be undone.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "namespace",
			Usage:     "The namespace. `default` exists for every organization.",
			Required:  true,
			PathParam: "namespace",
		},
		&requestflag.Flag[string]{
			Name:      "profile-id",
			Usage:     "The profile: your identifier for the user, caller or agent this memory is about.",
			Required:  true,
			PathParam: "profile_id",
		},
		&requestflag.Flag[string]{
			Name:      "source-id",
			Usage:     "Identifies one source within its profile: an ingested session, or one remembered fact. Returned by `ingest` and `remember` when the write is accepted. Re-ingesting a session keeps its source id.",
			Required:  true,
			PathParam: "source_id",
		},
	},
	Action:          handleAIMemoryNamespacesProfilesSourcesDelete,
	HideHelpCommand: true,
}

func handleAIMemoryNamespacesProfilesSourcesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("source-id") && len(unusedArgs) > 0 {
		cmd.Set("source-id", unusedArgs[0])
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

	params := telnyx.AIMemoryNamespaceProfileSourceGetParams{
		Namespace: cmd.Value("namespace").(string),
		ProfileID: cmd.Value("profile-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Memory.Namespaces.Profiles.Sources.Get(
		ctx,
		cmd.Value("source-id").(string),
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
		Title:          "ai:memory:namespaces:profiles:sources retrieve",
		Transform:      transform,
	})
}

func handleAIMemoryNamespacesProfilesSourcesList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("profile-id") && len(unusedArgs) > 0 {
		cmd.Set("profile-id", unusedArgs[0])
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

	params := telnyx.AIMemoryNamespaceProfileSourceListParams{
		Namespace: cmd.Value("namespace").(string),
	}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.AI.Memory.Namespaces.Profiles.Sources.List(
			ctx,
			cmd.Value("profile-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "ai:memory:namespaces:profiles:sources list",
			Transform:      transform,
		})
	} else {
		iter := client.AI.Memory.Namespaces.Profiles.Sources.ListAutoPaging(
			ctx,
			cmd.Value("profile-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "ai:memory:namespaces:profiles:sources list",
			Transform:      transform,
		})
	}
}

func handleAIMemoryNamespacesProfilesSourcesDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("source-id") && len(unusedArgs) > 0 {
		cmd.Set("source-id", unusedArgs[0])
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

	params := telnyx.AIMemoryNamespaceProfileSourceDeleteParams{
		Namespace: cmd.Value("namespace").(string),
		ProfileID: cmd.Value("profile-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Memory.Namespaces.Profiles.Sources.Delete(
		ctx,
		cmd.Value("source-id").(string),
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
		Title:          "ai:memory:namespaces:profiles:sources delete",
		Transform:      transform,
	})
}
