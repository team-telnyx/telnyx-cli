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
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var dialogflowConnectionsCreate = cli.Command{
	Name:    "create",
	Usage:   "Save Dialogflow Credentiails to Telnyx, so it can be used with other Telnyx\nservices.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "connection-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "service-account",
			Usage:    "The JSON map to connect your Dialoglow account.",
			Required: true,
			BodyPath: "service_account",
		},
		&requestflag.Flag[string]{
			Name:     "conversation-profile-id",
			Usage:    "The id of a configured conversation profile on your Dialogflow account. (If you use Dialogflow CX, this param is required)",
			BodyPath: "conversation_profile_id",
		},
		&requestflag.Flag[string]{
			Name:     "dialogflow-api",
			Usage:    "Determine which Dialogflow will be used.",
			Default:  "es",
			BodyPath: "dialogflow_api",
		},
		&requestflag.Flag[string]{
			Name:     "environment",
			Usage:    "Which Dialogflow environment will be used.",
			BodyPath: "environment",
		},
		&requestflag.Flag[string]{
			Name:     "location",
			Usage:    "The region of your agent is. (If you use Dialogflow CX, this param is required)",
			BodyPath: "location",
		},
	},
	Action:          handleDialogflowConnectionsCreate,
	HideHelpCommand: true,
}

var dialogflowConnectionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Return details of the Dialogflow connection associated with the given\nCallControl connection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "connection-id",
			Required: true,
		},
	},
	Action:          handleDialogflowConnectionsRetrieve,
	HideHelpCommand: true,
}

var dialogflowConnectionsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates a stored Dialogflow Connection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "connection-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "service-account",
			Usage:    "The JSON map to connect your Dialoglow account.",
			Required: true,
			BodyPath: "service_account",
		},
		&requestflag.Flag[string]{
			Name:     "conversation-profile-id",
			Usage:    "The id of a configured conversation profile on your Dialogflow account. (If you use Dialogflow CX, this param is required)",
			BodyPath: "conversation_profile_id",
		},
		&requestflag.Flag[string]{
			Name:     "dialogflow-api",
			Usage:    "Determine which Dialogflow will be used.",
			Default:  "es",
			BodyPath: "dialogflow_api",
		},
		&requestflag.Flag[string]{
			Name:     "environment",
			Usage:    "Which Dialogflow environment will be used.",
			BodyPath: "environment",
		},
		&requestflag.Flag[string]{
			Name:     "location",
			Usage:    "The region of your agent is. (If you use Dialogflow CX, this param is required)",
			BodyPath: "location",
		},
	},
	Action:          handleDialogflowConnectionsUpdate,
	HideHelpCommand: true,
}

var dialogflowConnectionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a stored Dialogflow Connection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "connection-id",
			Required: true,
		},
	},
	Action:          handleDialogflowConnectionsDelete,
	HideHelpCommand: true,
}

func handleDialogflowConnectionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("connection-id") && len(unusedArgs) > 0 {
		cmd.Set("connection-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.DialogflowConnectionNewParams{}

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
	_, err = client.DialogflowConnections.New(
		ctx,
		cmd.Value("connection-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "dialogflow-connections create", obj, format, transform)
}

func handleDialogflowConnectionsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("connection-id") && len(unusedArgs) > 0 {
		cmd.Set("connection-id", unusedArgs[0])
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
	_, err = client.DialogflowConnections.Get(ctx, cmd.Value("connection-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "dialogflow-connections retrieve", obj, format, transform)
}

func handleDialogflowConnectionsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("connection-id") && len(unusedArgs) > 0 {
		cmd.Set("connection-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.DialogflowConnectionUpdateParams{}

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
	_, err = client.DialogflowConnections.Update(
		ctx,
		cmd.Value("connection-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "dialogflow-connections update", obj, format, transform)
}

func handleDialogflowConnectionsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("connection-id") && len(unusedArgs) > 0 {
		cmd.Set("connection-id", unusedArgs[0])
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

	return client.DialogflowConnections.Delete(ctx, cmd.Value("connection-id").(string), options...)
}
