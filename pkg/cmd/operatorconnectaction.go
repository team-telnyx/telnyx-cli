// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/team-telnyx/telnyx-cli/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var operatorConnectActionsRefresh = cli.Command{
	Name:            "refresh",
	Usage:           "This endpoint will make an asynchronous request to refresh the Operator Connect\nintegration with Microsoft Teams for the current user. This will create new\nexternal connections on the user's account if needed, and/or report the\nintegration results as\n[log messages](https://developers.telnyx.com/api-reference/external-connections/list-all-log-messages#list-all-log-messages).",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleOperatorConnectActionsRefresh,
	HideHelpCommand: true,
}

func handleOperatorConnectActionsRefresh(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.OperatorConnect.Actions.Refresh(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "operator-connect:actions refresh", obj, format, transform)
}
