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

var aiFineTuningJobsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a new fine tuning job.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "model",
			Usage:    "The base model that is being fine-tuned.",
			Required: true,
			BodyPath: "model",
		},
		&requestflag.Flag[string]{
			Name:     "training-file",
			Usage:    "The storage bucket or object used for training.",
			Required: true,
			BodyPath: "training_file",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "hyperparameters",
			Usage:    "The hyperparameters used for the fine-tuning job.",
			BodyPath: "hyperparameters",
		},
		&requestflag.Flag[string]{
			Name:     "suffix",
			Usage:    "Optional suffix to append to the fine tuned model's name.",
			BodyPath: "suffix",
		},
	},
	Action:          handleAIFineTuningJobsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"hyperparameters": {
		&requestflag.InnerFlag[int64]{
			Name:       "hyperparameters.n-epochs",
			Usage:      "The number of epochs to train the model for. An epoch refers to one full cycle through the training dataset. 'auto' decides the optimal number of epochs based on the size of the dataset. If setting the number manually, we support any number between 1 and 50 epochs.",
			InnerField: "n_epochs",
		},
	},
})

var aiFineTuningJobsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a fine tuning job by `job_id`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "job-id",
			Required: true,
		},
	},
	Action:          handleAIFineTuningJobsRetrieve,
	HideHelpCommand: true,
}

var aiFineTuningJobsList = cli.Command{
	Name:            "list",
	Usage:           "Retrieve a list of all fine tuning jobs created by the user.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleAIFineTuningJobsList,
	HideHelpCommand: true,
}

var aiFineTuningJobsCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Cancel a fine tuning job.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "job-id",
			Required: true,
		},
	},
	Action:          handleAIFineTuningJobsCancel,
	HideHelpCommand: true,
}

func handleAIFineTuningJobsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIFineTuningJobNewParams{}

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
	_, err = client.AI.FineTuning.Jobs.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:fine-tuning:jobs create", obj, format, transform)
}

func handleAIFineTuningJobsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("job-id") && len(unusedArgs) > 0 {
		cmd.Set("job-id", unusedArgs[0])
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
	_, err = client.AI.FineTuning.Jobs.Get(ctx, cmd.Value("job-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:fine-tuning:jobs retrieve", obj, format, transform)
}

func handleAIFineTuningJobsList(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.FineTuning.Jobs.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:fine-tuning:jobs list", obj, format, transform)
}

func handleAIFineTuningJobsCancel(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("job-id") && len(unusedArgs) > 0 {
		cmd.Set("job-id", unusedArgs[0])
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
	_, err = client.AI.FineTuning.Jobs.Cancel(ctx, cmd.Value("job-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:fine-tuning:jobs cancel", obj, format, transform)
}
