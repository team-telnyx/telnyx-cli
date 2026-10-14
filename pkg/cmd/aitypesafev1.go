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

var aiTypesafeV1Systemone = cli.Command{
	Name:    "systemone",
	Usage:   "**Beta API.** Choose telnyx/decision-flash for the lowest cost and latency, or\ntelnyx/decision-pro for decisions that require long context, including inputs\nbeyond Jev’s 32k per-decision limit. Omitted model defaults to\ntelnyx/decision-flash.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "questions",
			Usage:    "Between 1 and 64 named questions. Each key identifies the corresponding answer.",
			Required: true,
			BodyPath: "questions",
		},
		&requestflag.Flag[any]{
			Name:     "state",
			Usage:    "Shared context evaluated by every question.",
			Required: true,
			BodyPath: "state",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			Usage:    "Public model alias. telnyx/decision-flash offers the lowest cost and latency; telnyx/decision-pro supports decisions that require long context, including inputs beyond Jev’s 32k per-decision limit. Applies to every question in the request. Other values are rejected.",
			Default:  "telnyx/decision-flash",
			BodyPath: "model",
		},
	},
	Action:          handleAITypesafeV1Systemone,
	HideHelpCommand: true,
}

func handleAITypesafeV1Systemone(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AITypesafeV1SystemoneParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Typesafe.V1.Systemone(ctx, params, options...)
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
		Title:          "ai:typesafe:v1 systemone",
		Transform:      transform,
	})
}
