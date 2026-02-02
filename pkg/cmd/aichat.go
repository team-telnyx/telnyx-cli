// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/telnyx-cli/internal/apiquery"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var aiChatCreateCompletion = requestflag.WithInnerFlags(cli.Command{
	Name:    "create-completion",
	Usage:   "Chat with a language model. This endpoint is consistent with the\n[OpenAI Chat Completions API](https://platform.openai.com/docs/api-reference/chat)\nand may be used with the OpenAI JS or Python SDK.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "message",
			Usage:    "A list of the previous chat messages for context.",
			Required: true,
			BodyPath: "messages",
		},
		&requestflag.Flag[string]{
			Name:     "api-key-ref",
			Usage:    "If you are using an external inference provider like xAI or OpenAI, this field allows you to pass along a reference to your API key. After creating an [integration secret](https://developers.telnyx.com/api-reference/integration-secrets/create-a-secret) for you API key, pass the secret's `identifier` in this field.",
			BodyPath: "api_key_ref",
		},
		&requestflag.Flag[int64]{
			Name:     "best-of",
			Usage:    "This is used with `use_beam_search` to determine how many candidate beams to explore.",
			BodyPath: "best_of",
		},
		&requestflag.Flag[bool]{
			Name:     "early-stopping",
			Usage:    "This is used with `use_beam_search`. If `true`, generation stops as soon as there are `best_of` complete candidates; if `false`, a heuristic is applied and the generation stops when is it very unlikely to find better candidates.",
			Default:  false,
			BodyPath: "early_stopping",
		},
		&requestflag.Flag[float64]{
			Name:     "frequency-penalty",
			Usage:    "Higher values will penalize the model from repeating the same output tokens.",
			Default:  0,
			BodyPath: "frequency_penalty",
		},
		&requestflag.Flag[[]string]{
			Name:     "guided-choice",
			Usage:    "If specified, the output will be exactly one of the choices.",
			BodyPath: "guided_choice",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "guided-json",
			Usage:    "Must be a valid JSON schema. If specified, the output will follow the JSON schema.",
			BodyPath: "guided_json",
		},
		&requestflag.Flag[string]{
			Name:     "guided-regex",
			Usage:    "If specified, the output will follow the regex pattern.",
			BodyPath: "guided_regex",
		},
		&requestflag.Flag[float64]{
			Name:     "length-penalty",
			Usage:    "This is used with `use_beam_search` to prefer shorter or longer completions.",
			Default:  1,
			BodyPath: "length_penalty",
		},
		&requestflag.Flag[bool]{
			Name:     "logprobs",
			Usage:    "Whether to return log probabilities of the output tokens or not. If true, returns the log probabilities of each output token returned in the `content` of `message`.",
			Default:  false,
			BodyPath: "logprobs",
		},
		&requestflag.Flag[int64]{
			Name:     "max-tokens",
			Usage:    "Maximum number of completion tokens the model should generate.",
			BodyPath: "max_tokens",
		},
		&requestflag.Flag[float64]{
			Name:     "min-p",
			Usage:    "This is an alternative to `top_p` that [many prefer](https://github.com/huggingface/transformers/issues/27670). Must be in [0, 1].",
			BodyPath: "min_p",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			Usage:    "The language model to chat with.",
			Default:  "meta-llama/Meta-Llama-3.1-8B-Instruct",
			BodyPath: "model",
		},
		&requestflag.Flag[float64]{
			Name:     "n",
			Usage:    "This will return multiple choices for you instead of a single chat completion.",
			BodyPath: "n",
		},
		&requestflag.Flag[float64]{
			Name:     "presence-penalty",
			Usage:    "Higher values will penalize the model from repeating the same output tokens.",
			Default:  0,
			BodyPath: "presence_penalty",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "response-format",
			Usage:    "Use this is you want to guarantee a JSON output without defining a schema. For control over the schema, use `guided_json`.",
			BodyPath: "response_format",
		},
		&requestflag.Flag[bool]{
			Name:     "stream",
			Usage:    "Whether or not to stream data-only server-sent events as they become available.",
			Default:  false,
			BodyPath: "stream",
		},
		&requestflag.Flag[float64]{
			Name:     "temperature",
			Usage:    `Adjusts the "creativity" of the model. Lower values make the model more deterministic and repetitive, while higher values make the model more random and creative.`,
			Default:  0.1,
			BodyPath: "temperature",
		},
		&requestflag.Flag[string]{
			Name:     "tool-choice",
			BodyPath: "tool_choice",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "tool",
			Usage:    "The `function` tool type follows the same schema as the [OpenAI Chat Completions API](https://platform.openai.com/docs/api-reference/chat). The `retrieval` tool type is unique to Telnyx. You may pass a list of [embedded storage buckets](https://developers.telnyx.com/api-reference/embeddings/embed-documents) for retrieval-augmented generation.",
			BodyPath: "tools",
		},
		&requestflag.Flag[int64]{
			Name:     "top-logprobs",
			Usage:    "This is used with `logprobs`. An integer between 0 and 20 specifying the number of most likely tokens to return at each token position, each with an associated log probability.",
			BodyPath: "top_logprobs",
		},
		&requestflag.Flag[float64]{
			Name:     "top-p",
			Usage:    "An alternative or complement to `temperature`. This adjusts how many of the top possibilities to consider.",
			BodyPath: "top_p",
		},
		&requestflag.Flag[bool]{
			Name:     "use-beam-search",
			Usage:    "Setting this to `true` will allow the model to [explore more completion options](https://huggingface.co/blog/how-to-generate#beam-search). This is not supported by OpenAI.",
			Default:  false,
			BodyPath: "use_beam_search",
		},
	},
	Action:          handleAIChatCreateCompletion,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"message": {
		&requestflag.InnerFlag[any]{
			Name:       "message.content",
			InnerField: "content",
		},
		&requestflag.InnerFlag[string]{
			Name:       "message.role",
			InnerField: "role",
		},
	},
	"response-format": {
		&requestflag.InnerFlag[string]{
			Name:       "response-format.type",
			InnerField: "type",
		},
	},
})

func handleAIChatCreateCompletion(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIChatNewCompletionParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Chat.NewCompletion(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:chat create-completion", obj, format, transform)
}
