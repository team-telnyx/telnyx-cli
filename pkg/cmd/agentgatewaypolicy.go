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

var agentGatewaysPoliciesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a named set of guardrails. It applies from the next request on, to the\nrequests that name it in `policies` and to those its\n[attachments](https://developers.telnyx.com/api-reference/agent-gateway-policy-attachments/attach-a-policy)\nmatch.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "A name, unique on the instance. A request names the policy in its `policies` field.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "base",
			Usage:    "A policy this one builds on, by name. The policy starts with the base's guardrails.",
			BodyPath: "base",
		},
		&requestflag.Flag[*string]{
			Name:     "description",
			Usage:    "What the policy is for.",
			BodyPath: "description",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "guardrails",
			Usage:    "The guardrails a policy adds to its base's, and the ones it removes from them, by name.",
			BodyPath: "guardrails",
		},
		&requestflag.Flag[*string]{
			Name:     "model",
			Usage:    "A condition on the model, by name or regular expression. The policy's guardrails run only on requests for a matching model. `null` matches every model.",
			BodyPath: "model",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "pipeline",
			Usage:    "Ordered steps that run in place of the policy's guardrails running independently.",
			BodyPath: "pipeline",
		},
	},
	Action:          handleAgentGatewaysPoliciesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"guardrails": {
		&requestflag.InnerFlag[[]string]{
			Name:       "guardrails.add",
			Usage:      "Guardrails the policy runs in addition to its base's.",
			InnerField: "add",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "guardrails.remove",
			Usage:      "Guardrails from its base that the policy does not run.",
			InnerField: "remove",
		},
	},
	"pipeline": {
		&requestflag.InnerFlag[string]{
			Name:       "pipeline.hook",
			Usage:      "`before_request` runs the steps on the request, and `after_request` on the model's response.",
			InnerField: "hook",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "pipeline.steps",
			InnerField: "steps",
		},
	},
})

var agentGatewaysPoliciesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns one policy, in its current version.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
		&requestflag.Flag[string]{
			Name:      "name",
			Required:  true,
			PathParam: "name",
		},
	},
	Action:          handleAgentGatewaysPoliciesRetrieve,
	HideHelpCommand: true,
}

var agentGatewaysPoliciesUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Changes a policy, which publishes a new version of it. A field left out stays as\nit is; `guardrails` and `pipeline` are replaced whole, and `model: null` removes\nthe condition.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
		&requestflag.Flag[string]{
			Name:      "name",
			Required:  true,
			PathParam: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "base",
			Usage:    "A policy this one builds on, by name. The policy starts with the base's guardrails.",
			BodyPath: "base",
		},
		&requestflag.Flag[*string]{
			Name:     "description",
			Usage:    "What the policy is for.",
			BodyPath: "description",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "guardrails",
			Usage:    "The guardrails a policy adds to its base's, and the ones it removes from them, by name.",
			BodyPath: "guardrails",
		},
		&requestflag.Flag[*string]{
			Name:     "model",
			Usage:    "A condition on the model, by name or regular expression. The policy's guardrails run only on requests for a matching model. `null` matches every model.",
			BodyPath: "model",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "pipeline",
			Usage:    "Ordered steps that run in place of the policy's guardrails running independently.",
			BodyPath: "pipeline",
		},
	},
	Action:          handleAgentGatewaysPoliciesUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"guardrails": {
		&requestflag.InnerFlag[[]string]{
			Name:       "guardrails.add",
			Usage:      "Guardrails the policy runs in addition to its base's.",
			InnerField: "add",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "guardrails.remove",
			Usage:      "Guardrails from its base that the policy does not run.",
			InnerField: "remove",
		},
	},
	"pipeline": {
		&requestflag.InnerFlag[string]{
			Name:       "pipeline.hook",
			Usage:      "`before_request` runs the steps on the request, and `after_request` on the model's response.",
			InnerField: "hook",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "pipeline.steps",
			InnerField: "steps",
		},
	},
})

var agentGatewaysPoliciesList = cli.Command{
	Name:    "list",
	Usage:   "Lists the instance's policies, each in its current version.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
	},
	Action:          handleAgentGatewaysPoliciesList,
	HideHelpCommand: true,
}

var agentGatewaysPoliciesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a policy and every version of it, and returns it as it was. A policy\nthat an attachment names, or that another policy builds on, cannot be deleted.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "instance-id",
			Required:  true,
			PathParam: "instance_id",
		},
		&requestflag.Flag[string]{
			Name:      "name",
			Required:  true,
			PathParam: "name",
		},
	},
	Action:          handleAgentGatewaysPoliciesDelete,
	HideHelpCommand: true,
}

func handleAgentGatewaysPoliciesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AgentGatewayPolicyNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AgentGateways.Policies.New(
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
		Title:          "agent-gateways:policies create",
		Transform:      transform,
	})
}

func handleAgentGatewaysPoliciesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("name") && len(unusedArgs) > 0 {
		cmd.Set("name", unusedArgs[0])
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

	params := telnyx.AgentGatewayPolicyGetParams{
		InstanceID: cmd.Value("instance-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AgentGateways.Policies.Get(
		ctx,
		cmd.Value("name").(string),
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
		Title:          "agent-gateways:policies retrieve",
		Transform:      transform,
	})
}

func handleAgentGatewaysPoliciesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("name") && len(unusedArgs) > 0 {
		cmd.Set("name", unusedArgs[0])
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

	params := telnyx.AgentGatewayPolicyUpdateParams{
		InstanceID: cmd.Value("instance-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AgentGateways.Policies.Update(
		ctx,
		cmd.Value("name").(string),
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
		Title:          "agent-gateways:policies update",
		Transform:      transform,
	})
}

func handleAgentGatewaysPoliciesList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.AgentGateways.Policies.List(ctx, cmd.Value("instance-id").(string), options...)
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
		Title:          "agent-gateways:policies list",
		Transform:      transform,
	})
}

func handleAgentGatewaysPoliciesDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("name") && len(unusedArgs) > 0 {
		cmd.Set("name", unusedArgs[0])
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

	params := telnyx.AgentGatewayPolicyDeleteParams{
		InstanceID: cmd.Value("instance-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AgentGateways.Policies.Delete(
		ctx,
		cmd.Value("name").(string),
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
		Title:          "agent-gateways:policies delete",
		Transform:      transform,
	})
}
