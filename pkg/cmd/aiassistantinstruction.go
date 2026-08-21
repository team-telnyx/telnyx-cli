// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/team-telnyx/telnyx-cli/internal/apiquery"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/urfave/cli/v3"
)

var aiAssistantsInstructionsEnhance = cli.Command{
	Name:    "enhance",
	Usage:   "Enhance an assistant's instructions using an LLM. The endpoint reads the\nassistant's current instructions and tools, then streams back improved\ninstructions as they are generated.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "assistant-id",
			Required:  true,
			PathParam: "assistant_id",
		},
		&requestflag.Flag[*string]{
			Name:     "enhancement-prompt",
			Usage:    "Optional guidance describing how the instructions should be enhanced. When provided, the LLM applies these requested changes in addition to fixing any identified issues.",
			BodyPath: "enhancement_prompt",
		},
		&requestflag.Flag[*string]{
			Name:     "instructions",
			Usage:    "The instructions to enhance. When omitted, the assistant's existing instructions are used.",
			BodyPath: "instructions",
		},
	},
	Action:          handleAIAssistantsInstructionsEnhance,
	HideHelpCommand: true,
}

func handleAIAssistantsInstructionsEnhance(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("assistant-id") && len(unusedArgs) > 0 {
		cmd.Set("assistant-id", unusedArgs[0])
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

	params := telnyx.AIAssistantInstructionEnhanceParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Assistants.Instructions.Enhance(
		ctx,
		cmd.Value("assistant-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(res)
	return err
}
