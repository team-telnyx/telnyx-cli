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

var simCardOrderPreviewPreview = cli.Command{
	Name:    "preview",
	Usage:   "Preview SIM card order purchases.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "address-id",
			Usage:    "Uniquely identifies the address for the order.",
			Required: true,
			BodyPath: "address_id",
		},
		&requestflag.Flag[int64]{
			Name:     "quantity",
			Usage:    "The amount of SIM cards that the user would like to purchase in the SIM card order.",
			Required: true,
			BodyPath: "quantity",
		},
	},
	Action:          handleSimCardOrderPreviewPreview,
	HideHelpCommand: true,
}

func handleSimCardOrderPreviewPreview(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SimCardOrderPreviewPreviewParams{}

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
	_, err = client.SimCardOrderPreview.Preview(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "sim-card-order-preview preview", obj, format, transform)
}
