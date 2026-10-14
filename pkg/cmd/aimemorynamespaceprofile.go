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

var aiMemoryNamespacesProfilesList = cli.Command{
	Name:    "list",
	Usage:   "Profiles are never created, only written to, so this lists the ones that hold a\nmemory. A profile whose first ingest is still running is not here yet. Ordered\nby memory count, largest first, so a profile written to while the listing is\npaged can move between pages and be repeated or missed.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "namespace",
			Required:  true,
			PathParam: "namespace",
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
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleAIMemoryNamespacesProfilesList,
	HideHelpCommand: true,
}

var aiMemoryNamespacesProfilesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete everything held about one profile. A 2xx means none of its memories are\nleft, and its summary goes with them. There is no undo.",
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
			Required:  true,
			PathParam: "profile_id",
		},
	},
	Action:          handleAIMemoryNamespacesProfilesDelete,
	HideHelpCommand: true,
}

var aiMemoryNamespacesProfilesIngest = cli.Command{
	Name:    "ingest",
	Usage:   "Store a session. Facts are extracted from whatever you send — the body is taken\nas any JSON value and stored whole, so a framework's own transcript shape works\nunchanged. `messages` of `role`/`content` is the conventional shape, not a\nrequirement. An empty object or a null body is refused. Carry a `session_id` to\nname the session: re-ingesting the same one replaces what it held. Omit it and a\nsession is opened and returned. Extraction runs asynchronously — poll the\nreturned operation.",
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
		&requestflag.Flag[map[string]any]{
			Name:     "body",
			Required: true,
			BodyRoot: true,
		},
		&requestflag.Flag[*string]{
			Name:      "session-id",
			Usage:     "Names the session. Re-ingesting the same session replaces what it held and keeps its `source_id`. Omit it to have one derived from the content and returned. No whitespace, control characters, or any of / \\ # ? %.",
			QueryPath: "session_id",
		},
	},
	Action:          handleAIMemoryNamespacesProfilesIngest,
	HideHelpCommand: true,
}

var aiMemoryNamespacesProfilesPreview = cli.Command{
	Name:    "preview",
	Usage:   "Try a candidate note on one profile and see how that profile's summary would\nread under it, without saving the note and without touching the stored summary.\nComposes the same source query the apply path\n(`PATCH /ai/memory/namespaces/{namespace}/settings`) would use, so a preview\ncannot drift from what saving the note would produce. Reads the candidate alone\n-- the namespace's current instructions and the profile's stored summary are\nneither read nor replaced. Bounded at 120 seconds: the summary is generated on\nthe spot rather than read from a stored one. Nothing about the request is\nstored.",
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
			Required:  true,
			PathParam: "profile_id",
		},
		&requestflag.Flag[*string]{
			Name:     "instructions",
			Usage:    "The candidate note to try. `null`, empty, or whitespace-only previews the neutral default -- the summary a profile returns to when instructions are cleared. Capped at 2000 characters, the same cap `PATCH /settings` applies on save; a longer note is refused with `422` before anything is generated.",
			BodyPath: "instructions",
		},
	},
	Action:          handleAIMemoryNamespacesProfilesPreview,
	HideHelpCommand: true,
}

var aiMemoryNamespacesProfilesRecall = cli.Command{
	Name:    "recall",
	Usage:   "Ranked memories for a question. Matching runs over the profile's memories and\nreturns them in rank order with a relevance `score`; the score is null where the\ndeployment's reranker is a passthrough, in which case order is the only signal.\nNo model runs in this path — recall returns facts, it does not compose an\nanswer.",
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
		&requestflag.Flag[string]{
			Name:     "query",
			Required: true,
			BodyPath: "query",
		},
		&requestflag.Flag[*int64]{
			Name:     "top-k",
			BodyPath: "top_k",
		},
	},
	Action:          handleAIMemoryNamespacesProfilesRecall,
	HideHelpCommand: true,
}

var aiMemoryNamespacesProfilesRemember = cli.Command{
	Name:    "remember",
	Usage:   "For a fact the agent has already distilled: `text` is stored as given, with\nnothing extracted from it. Send a transcript to `ingest` instead. Remembering\nthe same text again writes the same memory rather than a second copy of it, so a\nretry is safe. The write runs asynchronously -- poll the returned operation.",
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
		&requestflag.Flag[string]{
			Name:     "text",
			Required: true,
			BodyPath: "text",
		},
	},
	Action:          handleAIMemoryNamespacesProfilesRemember,
	HideHelpCommand: true,
}

var aiMemoryNamespacesProfilesRetrieveSummary = cli.Command{
	Name:    "retrieve-summary",
	Usage:   "The whole profile as one card, precomputed, with no query. Built for the start\nof a session, where there is no question to ask yet.",
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
			Required:  true,
			PathParam: "profile_id",
		},
	},
	Action:          handleAIMemoryNamespacesProfilesRetrieveSummary,
	HideHelpCommand: true,
}

func handleAIMemoryNamespacesProfilesList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("namespace") && len(unusedArgs) > 0 {
		cmd.Set("namespace", unusedArgs[0])
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

	params := telnyx.AIMemoryNamespaceProfileListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.AI.Memory.Namespaces.Profiles.List(
			ctx,
			cmd.Value("namespace").(string),
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
			Title:          "ai:memory:namespaces:profiles list",
			Transform:      transform,
		})
	} else {
		iter := client.AI.Memory.Namespaces.Profiles.ListAutoPaging(
			ctx,
			cmd.Value("namespace").(string),
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
			Title:          "ai:memory:namespaces:profiles list",
			Transform:      transform,
		})
	}
}

func handleAIMemoryNamespacesProfilesDelete(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AIMemoryNamespaceProfileDeleteParams{
		Namespace: cmd.Value("namespace").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Memory.Namespaces.Profiles.Delete(
		ctx,
		cmd.Value("profile-id").(string),
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
		Title:          "ai:memory:namespaces:profiles delete",
		Transform:      transform,
	})
}

func handleAIMemoryNamespacesProfilesIngest(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := telnyx.AIMemoryNamespaceProfileIngestParams{
		Namespace: cmd.Value("namespace").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Memory.Namespaces.Profiles.Ingest(
		ctx,
		cmd.Value("profile-id").(string),
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
		Title:          "ai:memory:namespaces:profiles ingest",
		Transform:      transform,
	})
}

func handleAIMemoryNamespacesProfilesPreview(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := telnyx.AIMemoryNamespaceProfilePreviewParams{
		Namespace: cmd.Value("namespace").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Memory.Namespaces.Profiles.Preview(
		ctx,
		cmd.Value("profile-id").(string),
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
		Title:          "ai:memory:namespaces:profiles preview",
		Transform:      transform,
	})
}

func handleAIMemoryNamespacesProfilesRecall(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := telnyx.AIMemoryNamespaceProfileRecallParams{
		Namespace: cmd.Value("namespace").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Memory.Namespaces.Profiles.Recall(
		ctx,
		cmd.Value("profile-id").(string),
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
		Title:          "ai:memory:namespaces:profiles recall",
		Transform:      transform,
	})
}

func handleAIMemoryNamespacesProfilesRemember(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := telnyx.AIMemoryNamespaceProfileRememberParams{
		Namespace: cmd.Value("namespace").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Memory.Namespaces.Profiles.Remember(
		ctx,
		cmd.Value("profile-id").(string),
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
		Title:          "ai:memory:namespaces:profiles remember",
		Transform:      transform,
	})
}

func handleAIMemoryNamespacesProfilesRetrieveSummary(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AIMemoryNamespaceProfileGetSummaryParams{
		Namespace: cmd.Value("namespace").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Memory.Namespaces.Profiles.GetSummary(
		ctx,
		cmd.Value("profile-id").(string),
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
		Title:          "ai:memory:namespaces:profiles retrieve-summary",
		Transform:      transform,
	})
}
