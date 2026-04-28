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

var messaging10dlcGetEnum = cli.Command{
	Name:    "get-enum",
	Usage:   "Get Enum",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "endpoint",
			Usage:    `Allowed values: "mno", "optionalAttributes", "usecase", "vertical", "altBusinessIdType", "brandIdentityStatus", "brandRelationship", "campaignStatus", "entityType", "extVettingProvider", "vettingStatus", "brandStatus", "operationStatus", "approvedPublicCompany", "stockExchange", "vettingClass".`,
			Required: true,
		},
	},
	Action:          handleMessaging10dlcGetEnum,
	HideHelpCommand: true,
}

func handleMessaging10dlcGetEnum(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("endpoint") && len(unusedArgs) > 0 {
		cmd.Set("endpoint", unusedArgs[0])
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
	_, err = client.Messaging10dlc.GetEnum(ctx, telnyx.Messaging10dlcGetEnumParamsEndpoint(cmd.Value("endpoint").(string)), options...)
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
		Title:          "messaging-10dlc get-enum",
		Transform:      transform,
	})
}
