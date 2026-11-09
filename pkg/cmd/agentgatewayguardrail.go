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

var agentGatewaysGuardrailsCreate = cli.Command{
	Name:    "create",
	Usage:   "Adds a check the instance runs on model and tool requests. It runs on every\nrequest when `always_on`, and otherwise on the requests that name it in\n`guardrails`. The instance must be `ready`; otherwise the request answers `409`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
		&requestflag.Flag[[]string]{
			Name:     "mode",
			Usage:    "When the guardrail runs. Each type runs in some modes only.",
			Required: true,
			BodyPath: "modes",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "A name, unique on the instance. A request names it in its `guardrails` field.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "The kind of check. It cannot change after create.",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[bool]{
			Name:     "always-on",
			Usage:    "`true` runs the guardrail on every request, and a request cannot turn it off. `false` runs it on the requests that name it.",
			Default:  false,
			BodyPath: "always_on",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "settings",
			Usage:    "The type's settings. `content_filter`: `patterns` (prebuilt patterns by `name`, such as `us_ssn` or `email`), `regexes` (each a `name`, which tags masked matches, and a `pattern`), `keywords` (each a `keyword`), and `categories` (each a `name`), each entry with an `action` of `block` or `mask`. `tool_permission`: `rules` (`id`, `tool_name`, `tool_type`, `decision`, `argument_patterns`), a `default` decision, and `on_denied` (`block` or `rewrite`). `model_judge`: a `model` on the instance and weighted `criteria` (`name`, `description`, `weight`, the weights summing to 100). `custom_code`: `code` defining `apply_guardrail(inputs, request_data, input_type)`. `lakera`: `api_key`, `project_id`, `api_base`. `presidio`: `analyzer_url`, `anonymizer_url`, `entities` (entity name to `block` or `mask`), `language`. `bedrock`: `guardrail_id`, `version`, `region`, and `access_key_id` with `secret_access_key`, or `role_name`. `azure_content_safety`: `api_key`, `endpoint`, `severity` (0 to 7). Provider secrets are write-only.",
			BodyPath: "settings",
		},
	},
	Action:          handleAgentGatewaysGuardrailsCreate,
	HideHelpCommand: true,
}

var agentGatewaysGuardrailsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Changes a guardrail's `name`, `modes`, `always_on`, or `settings`. `settings` is\nreplaced whole, secrets included. The `type` cannot change. The instance must be\n`ready`; otherwise the request answers `409`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
		&requestflag.Flag[string]{
			Name:      "guardrail-id",
			Required:  true,
			PathParam: "guardrail_id",
		},
		&requestflag.Flag[bool]{
			Name:     "always-on",
			Usage:    "`true` runs the guardrail on every request, and a request cannot turn it off. `false` runs it on the requests that name it.",
			Default:  false,
			BodyPath: "always_on",
		},
		&requestflag.Flag[[]string]{
			Name:     "mode",
			Usage:    "When the guardrail runs. Each type runs in some modes only.",
			BodyPath: "modes",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "A name, unique on the instance. A request names it in its `guardrails` field.",
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "settings",
			Usage:    "The type's settings. `content_filter`: `patterns` (prebuilt patterns by name, such as `us_ssn` or `email`), `regexes`, `keywords`, and `categories` (such as `prompt_injection`), each with an `action` of `block` or `mask`. `tool_permission`: `rules` (`id`, `tool_name`, `tool_type`, `decision`, `argument_patterns`), a `default` decision, and `on_denied` (`block` or `rewrite`). `model_judge`: a `model` on the instance and weighted `criteria` (`name`, `description`, `weight`, the weights summing to 100). `custom_code`: `code` defining `apply_guardrail(inputs, request_data, input_type)`. `lakera`: `api_key`, `project_id`, `api_base`. `presidio`: `analyzer_url`, `anonymizer_url`, `entities` (entity name to `block` or `mask`), `language`. `bedrock`: `guardrail_id`, `version`, `region`, and `access_key_id` with `secret_access_key`, or `role_name`. `azure_content_safety`: `api_key`, `endpoint`, `severity` (0 to 7). Provider secrets are write-only.",
			BodyPath: "settings",
		},
	},
	Action:          handleAgentGatewaysGuardrailsUpdate,
	HideHelpCommand: true,
}

var agentGatewaysGuardrailsList = cli.Command{
	Name:    "list",
	Usage:   "Lists the instance's guardrails, without their secrets. The instance must be\n`ready`; otherwise the request answers `409`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
	},
	Action:          handleAgentGatewaysGuardrailsList,
	HideHelpCommand: true,
}

var agentGatewaysGuardrailsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Removes a guardrail and returns it as it was. The instance must be `ready`;\notherwise the request answers `409`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
		&requestflag.Flag[string]{
			Name:      "guardrail-id",
			Required:  true,
			PathParam: "guardrail_id",
		},
	},
	Action:          handleAgentGatewaysGuardrailsDelete,
	HideHelpCommand: true,
}

func handleAgentGatewaysGuardrailsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("instance-id") && len(unusedArgs) > 0 {
		cmd.Set("instance-id", unusedArgs[0])
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

	params := telnyx.AgentGatewayGuardrailNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AgentGateways.Guardrails.New(
		ctx,
		cmd.Value("instance-id").(string),
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
		Title:          "agent-gateways:guardrails create",
		Transform:      transform,
	})
}

func handleAgentGatewaysGuardrailsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("guardrail-id") && len(unusedArgs) > 0 {
		cmd.Set("guardrail-id", unusedArgs[0])
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

	params := telnyx.AgentGatewayGuardrailUpdateParams{
		InstanceID: cmd.Value("instance-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AgentGateways.Guardrails.Update(
		ctx,
		cmd.Value("guardrail-id").(string),
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
		Title:          "agent-gateways:guardrails update",
		Transform:      transform,
	})
}

func handleAgentGatewaysGuardrailsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("instance-id") && len(unusedArgs) > 0 {
		cmd.Set("instance-id", unusedArgs[0])
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
	_, err = client.AgentGateways.Guardrails.List(ctx, cmd.Value("instance-id").(string), options...)
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
		Title:          "agent-gateways:guardrails list",
		Transform:      transform,
	})
}

func handleAgentGatewaysGuardrailsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("guardrail-id") && len(unusedArgs) > 0 {
		cmd.Set("guardrail-id", unusedArgs[0])
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

	params := telnyx.AgentGatewayGuardrailDeleteParams{
		InstanceID: cmd.Value("instance-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AgentGateways.Guardrails.Delete(
		ctx,
		cmd.Value("guardrail-id").(string),
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
		Title:          "agent-gateways:guardrails delete",
		Transform:      transform,
	})
}
