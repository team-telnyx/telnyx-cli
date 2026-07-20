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

var aiAnthropicV1Messages = cli.Command{
	Name:    "messages",
	Usage:   "Send a message to a language model using the Anthropic Messages API format. This\nendpoint is compatible with the\n[Anthropic Messages API](https://docs.anthropic.com/en/api/messages) and may be\nused with the Anthropic JS or Python SDK by setting the base URL to\n`https://api.telnyx.com/v2/ai/anthropic`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "max-tokens",
			Usage:    "The maximum number of tokens to generate in the response.",
			Required: true,
			BodyPath: "max_tokens",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "message",
			Usage:    "The messages to send to the model, following the [Anthropic Messages API](https://docs.anthropic.com/en/api/messages) format.",
			Required: true,
			BodyPath: "messages",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			Usage:    "The model to use for generating the response, for example `zai-org/GLM-5.2` or another model available from the Telnyx models endpoint.",
			Required: true,
			BodyPath: "model",
		},
		&requestflag.Flag[string]{
			Name:     "api-key-ref",
			Usage:    "If you are using an external inference provider, this field allows you to pass along a reference to your API key. After creating an [integration secret](https://developers.telnyx.com/api-reference/integration-secrets/create-a-secret) for your API key, pass the secret's `identifier` in this field.",
			BodyPath: "api_key_ref",
		},
		&requestflag.Flag[string]{
			Name:     "billing-group-id",
			Usage:    "The billing group ID to associate with this request.",
			BodyPath: "billing_group_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "fallback-config",
			Usage:    "Configuration for model fallback behavior when the primary model is unavailable.",
			BodyPath: "fallback_config",
		},
		&requestflag.Flag[int64]{
			Name:     "max-retries",
			Usage:    "Maximum number of retries for the request.",
			BodyPath: "max_retries",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "mcp-server",
			Usage:    "List of MCP (Model Context Protocol) servers to make available to the model.",
			BodyPath: "mcp_servers",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "An object describing metadata about the request.",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "service-tier",
			Usage:    "Service tier for the request.",
			BodyPath: "service_tier",
		},
		&requestflag.Flag[[]string]{
			Name:     "stop-sequence",
			Usage:    "Custom sequences that will cause the model to stop generating.",
			BodyPath: "stop_sequences",
		},
		&requestflag.Flag[bool]{
			Name:     "stream",
			Usage:    "Whether to stream the response as Anthropic-format Server-Sent Events.",
			Default:  false,
			BodyPath: "stream",
		},
		&requestflag.Flag[any]{
			Name:     "system",
			Usage:    "System prompt. Can be a string or an array of content blocks following the Anthropic API format.",
			BodyPath: "system",
		},
		&requestflag.Flag[float64]{
			Name:     "temperature",
			Usage:    "Amount of randomness injected into the response. Ranges from 0 to 1.",
			BodyPath: "temperature",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "thinking",
			Usage:    "Extended thinking configuration for models that support it. Set `type` to `enabled` to turn on extended thinking.",
			BodyPath: "thinking",
		},
		&requestflag.Flag[float64]{
			Name:     "timeout",
			Usage:    "Request timeout in seconds.",
			Default:  300,
			BodyPath: "timeout",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "tool-choice",
			Usage:    "Controls how the model uses tools, following the Anthropic API format.",
			BodyPath: "tool_choice",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "tool",
			Usage:    "Definitions of tools that the model may use, following the Anthropic API format.",
			BodyPath: "tools",
		},
		&requestflag.Flag[int64]{
			Name:     "top-k",
			Usage:    "Top-k sampling parameter. Only sample from the top K options for each subsequent token.",
			BodyPath: "top_k",
		},
		&requestflag.Flag[float64]{
			Name:     "top-p",
			Usage:    "Nucleus sampling parameter. Use temperature or top_p, but not both.",
			BodyPath: "top_p",
		},
	},
	Action:          handleAIAnthropicV1Messages,
	HideHelpCommand: true,
}

func handleAIAnthropicV1Messages(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AIAnthropicV1MessagesParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Anthropic.V1.Messages(ctx, params, options...)
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
		Title:          "ai:anthropic:v1 messages",
		Transform:      transform,
	})
}
