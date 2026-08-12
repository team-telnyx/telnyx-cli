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

var rcsAgentsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates an editable RCS agent draft under a brand. The `Idempotency-Key` is\nscoped to the authenticated organization. Reusing the key with the same request\nreturns the original agent, while reusing it with a different request returns a\nconflict.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "brand-id",
			Required: true,
			BodyPath: "brand_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "configuration",
			Required: true,
			BodyPath: "configuration",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			Required: true,
			BodyPath: "display_name",
		},
		&requestflag.Flag[string]{
			Name:     "use-case",
			Usage:    `Allowed values: "MULTI_USE", "PROMOTIONAL", "TRANSACTIONAL", "OTP".`,
			Required: true,
			BodyPath: "use_case",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			Required:   true,
			HeaderPath: "Idempotency-Key",
		},
		&requestflag.Flag[*string]{
			Name:     "hosting-region",
			BodyPath: "hosting_region",
		},
		&requestflag.Flag[*string]{
			Name:     "profile-id",
			Usage:    "A Messaging Profile owned by the authenticated organization. When omitted, the agent inherits the brand profile.",
			BodyPath: "profile_id",
		},
	},
	Action:          handleRcsAgentsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"configuration": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "configuration.basics",
			Usage:      "Basic agent identity and contact information. At least one complete phone, website, or email contact is required.",
			InnerField: "basics",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "configuration.campaign",
			InnerField: "campaign",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "configuration.testing",
			InnerField: "testing",
		},
	},
})

var rcsAgentsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves an RCS agent, section statuses, test devices, carrier approvals, and\nprovider capabilities.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleRcsAgentsRetrieve,
	HideHelpCommand: true,
}

var rcsAgentsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates one or more fields on an agent while its status is `CREATED`. Submitted\nagents cannot be changed through this endpoint.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "configuration",
			BodyPath: "configuration",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			BodyPath: "display_name",
		},
		&requestflag.Flag[string]{
			Name:     "hosting-region",
			BodyPath: "hosting_region",
		},
		&requestflag.Flag[string]{
			Name:     "profile-id",
			BodyPath: "profile_id",
		},
		&requestflag.Flag[string]{
			Name:     "use-case",
			Usage:    `Allowed values: "MULTI_USE", "PROMOTIONAL", "TRANSACTIONAL", "OTP".`,
			BodyPath: "use_case",
		},
	},
	Action:          handleRcsAgentsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"configuration": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "configuration.basics",
			Usage:      "Basic agent identity and contact information. At least one complete phone, website, or email contact is required.",
			InnerField: "basics",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "configuration.campaign",
			InnerField: "campaign",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "configuration.testing",
			InnerField: "testing",
		},
	},
})

var rcsAgentsList = cli.Command{
	Name:    "list",
	Usage:   "Lists RCS agents owned by the authenticated organization, optionally filtered by\nbrand.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "brand-id",
			Usage:     "Only return agents belonging to this brand.",
			QueryPath: "brand_id",
		},
	},
	Action:          handleRcsAgentsList,
	HideHelpCommand: true,
}

var rcsAgentsLaunch = requestflag.WithInnerFlags(cli.Command{
	Name:    "launch",
	Usage:   "Adds the campaign and testing configuration, then starts asynchronous carrier\nlaunch. Agent basics must already be submitted. Repeating a launch that is\nalready in progress returns the current agent without creating new work.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "campaign",
			Required: true,
			BodyPath: "campaign",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "testing",
			Required: true,
			BodyPath: "testing",
		},
	},
	Action:          handleRcsAgentsLaunch,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"testing": {
		&requestflag.InnerFlag[string]{
			Name:       "testing.test-url",
			Usage:      "A publicly accessible test video or evidence URL.",
			InnerField: "test_url",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "testing.additional-information",
			InnerField: "additional_information",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "testing.message-id",
			InnerField: "message_id",
		},
	},
})

var rcsAgentsRetrieveCarrierApprovals = cli.Command{
	Name:    "retrieve-carrier-approvals",
	Usage:   "Lists carrier approval records for an RCS agent. The provider may expose\nper-carrier, hub-level, or bot-level approval status.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleRcsAgentsRetrieveCarrierApprovals,
	HideHelpCommand: true,
}

var rcsAgentsSubmit = cli.Command{
	Name:    "submit",
	Usage:   "Starts asynchronous provider provisioning and submits the agent's basic\nconfiguration. The brand must be `VERIFIED`. Repeating this request for an\nin-progress agent returns its current state without creating new work.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleRcsAgentsSubmit,
	HideHelpCommand: true,
}

func handleRcsAgentsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.RcAgentNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Rcs.Agents.New(ctx, params, options...)
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
		Title:          "rcs:agents create",
		Transform:      transform,
	})
}

func handleRcsAgentsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Rcs.Agents.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "rcs:agents retrieve",
		Transform:      transform,
	})
}

func handleRcsAgentsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.RcAgentUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Rcs.Agents.Update(
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
		Title:          "rcs:agents update",
		Transform:      transform,
	})
}

func handleRcsAgentsList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.RcAgentListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Rcs.Agents.List(ctx, params, options...)
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
		Title:          "rcs:agents list",
		Transform:      transform,
	})
}

func handleRcsAgentsLaunch(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.RcAgentLaunchParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Rcs.Agents.Launch(
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
		Title:          "rcs:agents launch",
		Transform:      transform,
	})
}

func handleRcsAgentsRetrieveCarrierApprovals(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Rcs.Agents.GetCarrierApprovals(ctx, cmd.Value("id").(string), options...)
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
		Title:          "rcs:agents retrieve-carrier-approvals",
		Transform:      transform,
	})
}

func handleRcsAgentsSubmit(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Rcs.Agents.Submit(ctx, cmd.Value("id").(string), options...)
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
		Title:          "rcs:agents submit",
		Transform:      transform,
	})
}
