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

var aiMemoryNamespacesProfilesMemoriesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "One memory by its id, as `recall` and the listing return it, together with what\nit came from. A fact names its `source_id`: read it with\n`GET .../sources/{source_id}` to see what was stored. A memory derived from\nother memories names them in `derived_from` instead; read each of those to reach\nits source.",
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
			Name:      "memory-id",
			Usage:     "A memory's id, as `recall` and the listing return it.",
			Required:  true,
			PathParam: "memory_id",
		},
	},
	Action:          handleAIMemoryNamespacesProfilesMemoriesRetrieve,
	HideHelpCommand: true,
}

var aiMemoryNamespacesProfilesMemoriesList = cli.Command{
	Name:    "list",
	Usage:   "Everything stored under one profile, unranked -- ask `recall` for the memories\nthat answer a question. A profile that holds nothing is an empty page rather\nthan a 404: profiles exist by being written to. Each memory names the\n`source_id` it was extracted from, or null for a memory derived from other\nmemories -- which can read almost the same as the fact it restates. A\n`source_id` narrows the listing to the memories extracted from that source, and\na `session_id` to those extracted from the session, which is the same thing\nnamed another way; pass one or the other. Neither is everything the source led\nto: a memory derived from several sources belongs to no single one and appears\nonly in the unfiltered listing. A memory written while the listing is paged\nshifts the pages after it, so an entry can be repeated or missed at a page\nboundary.",
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
		&requestflag.Flag[*string]{
			Name:      "source-id",
			Usage:     "Narrows the listing to the memories extracted from one source, a remembered fact as well as a session. Pass this or `session_id`, not both.",
			QueryPath: "source_id",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleAIMemoryNamespacesProfilesMemoriesList,
	HideHelpCommand: true,
}

func handleAIMemoryNamespacesProfilesMemoriesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("memory-id") && len(unusedArgs) > 0 {
		cmd.Set("memory-id", unusedArgs[0])
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

	params := telnyx.AIMemoryNamespaceProfileMemoryGetParams{
		Namespace: cmd.Value("namespace").(string),
		ProfileID: cmd.Value("profile-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Memory.Namespaces.Profiles.Memories.Get(
		ctx,
		cmd.Value("memory-id").(string),
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
		Title:          "ai:memory:namespaces:profiles:memories retrieve",
		Transform:      transform,
	})
}

func handleAIMemoryNamespacesProfilesMemoriesList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AIMemoryNamespaceProfileMemoryListParams{
		Namespace: cmd.Value("namespace").(string),
	}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.AI.Memory.Namespaces.Profiles.Memories.List(
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
			Title:          "ai:memory:namespaces:profiles:memories list",
			Transform:      transform,
		})
	} else {
		iter := client.AI.Memory.Namespaces.Profiles.Memories.ListAutoPaging(
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
			Title:          "ai:memory:namespaces:profiles:memories list",
			Transform:      transform,
		})
	}
}
