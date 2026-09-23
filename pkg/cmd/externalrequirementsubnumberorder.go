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

var externalRequirementsSubNumberOrdersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns the input fields an action requirement needs and the current requirement\naction for a sub number order. Action requirements are fulfilled by an external\nstep rather than by uploading documents. Australia mobile ID verification is\ncurrently the only action requirement. Once a verification link has been\ngenerated, it is returned in `requirement_action.value`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "regulatory-requirement-id",
			Required:  true,
			PathParam: "regulatory_requirement_id",
		},
		&requestflag.Flag[string]{
			Name:      "sub-number-order-id",
			Required:  true,
			PathParam: "sub_number_order_id",
		},
	},
	Action:          handleExternalRequirementsSubNumberOrdersRetrieve,
	HideHelpCommand: true,
}

var externalRequirementsSubNumberOrdersUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Submits the end user's details to the external verification provider and returns\nthe requirement action. Australia mobile ID verification is currently the only\naction requirement. It generates a unique Onfido verification link, returned in\n`requirement_action.value`, which you share with the end user. The end user's\n`first_name` and `last_name` must be nested inside a `requirement` object;\nsending them at the top level is rejected.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "regulatory-requirement-id",
			Required:  true,
			PathParam: "regulatory_requirement_id",
		},
		&requestflag.Flag[string]{
			Name:      "sub-number-order-id",
			Required:  true,
			PathParam: "sub_number_order_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "requirement",
			Usage:    "The end user's identity details for the action requirement. Australia mobile ID verification is currently the only action requirement. It requires `first_name` and `last_name`, the same fields the corresponding GET lists in `fields_required`.",
			Required: true,
			BodyPath: "requirement",
		},
	},
	Action:          handleExternalRequirementsSubNumberOrdersUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"requirement": {
		&requestflag.InnerFlag[string]{
			Name:       "requirement.first-name",
			Usage:      "The end user's first name.",
			InnerField: "first_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "requirement.last-name",
			Usage:      "The end user's last name.",
			InnerField: "last_name",
		},
	},
})

func handleExternalRequirementsSubNumberOrdersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("sub-number-order-id") && len(unusedArgs) > 0 {
		cmd.Set("sub-number-order-id", unusedArgs[0])
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

	params := telnyx.ExternalRequirementSubNumberOrderGetParams{
		RegulatoryRequirementID: cmd.Value("regulatory-requirement-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ExternalRequirements.SubNumberOrders.Get(
		ctx,
		cmd.Value("sub-number-order-id").(string),
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
		Title:          "external-requirements:sub-number-orders retrieve",
		Transform:      transform,
	})
}

func handleExternalRequirementsSubNumberOrdersUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("sub-number-order-id") && len(unusedArgs) > 0 {
		cmd.Set("sub-number-order-id", unusedArgs[0])
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

	params := telnyx.ExternalRequirementSubNumberOrderUpdateParams{
		RegulatoryRequirementID: cmd.Value("regulatory-requirement-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ExternalRequirements.SubNumberOrders.Update(
		ctx,
		cmd.Value("sub-number-order-id").(string),
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
		Title:          "external-requirements:sub-number-orders update",
		Transform:      transform,
	})
}
