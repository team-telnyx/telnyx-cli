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

var webSearchResearchCreate = cli.Command{
	Name:    "create",
	Usage:   "Starts a deep research task that runs multiple searches, reads sources, and\nsynthesizes an answer with citations.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "The research question or topic.",
			Required: true,
			BodyPath: "query",
		},
		&requestflag.Flag[bool]{
			Name:     "background",
			Usage:    "When `true`, the research runs asynchronously. The response returns a `task_id` immediately instead of waiting for the result. Poll `GET /web_search/research/{task_id}` to check status.",
			BodyPath: "background",
		},
		&requestflag.Flag[int64]{
			Name:     "max-sources",
			Usage:    "Maximum number of sources to use.",
			BodyPath: "max_sources",
		},
		&requestflag.Flag[string]{
			Name:     "research-effort",
			Usage:    "Research depth level. `lite` is fastest, `deep` is most thorough.",
			BodyPath: "research_effort",
		},
	},
	Action:          handleWebSearchResearchCreate,
	HideHelpCommand: true,
}

var webSearchResearchRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Polls the status of a previously started asynchronous research task. When the\nstatus is `completed`, the response includes the answer and citations. When the\nstatus is `failed`, the response includes an error message.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "task-id",
			Required:  true,
			PathParam: "task_id",
		},
	},
	Action:          handleWebSearchResearchRetrieve,
	HideHelpCommand: true,
}

func handleWebSearchResearchCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.WebSearchResearchNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.WebSearch.Research.New(ctx, params, options...)
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
		Title:          "web-search:research create",
		Transform:      transform,
	})
}

func handleWebSearchResearchRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("task-id") && len(unusedArgs) > 0 {
		cmd.Set("task-id", unusedArgs[0])
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
	_, err = client.WebSearch.Research.Get(ctx, cmd.Value("task-id").(string), options...)
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
		Title:          "web-search:research retrieve",
		Transform:      transform,
	})
}
