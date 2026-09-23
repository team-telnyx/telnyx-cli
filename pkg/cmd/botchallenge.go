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

var botChallengeCreate = cli.Command{
	Name:    "create",
	Usage:   "Generates a reverse-CAPTCHA challenge used to gate the bot signup flow. A random\nactive problem is selected from the pool; math problems are returned obfuscated\n(case randomization, symbol injection, spacing noise) with an unobfuscated\nrounding instruction appended, while string and binary problems are returned\nas-is. The response contains a single-use nonce, the problem text, and the\ncurrent terms-and-conditions and privacy-policy URLs, which must be echoed back\non the signup request. Challenges expire after a short window (10 minutes by\ndefault) and can only be answered once. This endpoint is public and\nunauthenticated.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "llm-model-name",
			Usage:    "Name of the LLM the client is using.",
			BodyPath: "llm_model_name",
		},
		&requestflag.Flag[string]{
			Name:     "llm-parameter-count",
			Usage:    "Parameter count of the client LLM.",
			BodyPath: "llm_parameter_count",
		},
		&requestflag.Flag[string]{
			Name:     "llm-quantization",
			Usage:    "Quantization of the client LLM.",
			BodyPath: "llm_quantization",
		},
	},
	Action:          handleBotChallengeCreate,
	HideHelpCommand: true,
}

func handleBotChallengeCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.BotChallengeNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.BotChallenge.New(ctx, params, options...)
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
		Title:          "bot-challenge create",
		Transform:      transform,
	})
}
