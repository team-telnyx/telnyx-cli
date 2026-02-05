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

var actionsPurchaseCreate = cli.Command{
	Name:    "create",
	Usage:   "Purchases and registers the specified amount of eSIMs to the current user's\naccount.<br/><br/> If <code>sim_card_group_id</code> is provided, the eSIMs will\nbe associated with that group. Otherwise, the default group for the current user\nwill be used.<br/><br/>",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "amount",
			Usage:    "The amount of eSIMs to be purchased.",
			Required: true,
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "product",
			Usage:    `Type of product to be purchased, specify "whitelabel" to use a custom SPN`,
			BodyPath: "product",
		},
		&requestflag.Flag[string]{
			Name:     "sim-card-group-id",
			Usage:    "The group SIMCardGroup identification. This attribute can be <code>null</code> when it's present in an associated resource.",
			BodyPath: "sim_card_group_id",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "Status on which the SIM cards will be set after being successfully registered.",
			Default:  "enabled",
			BodyPath: "status",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Searchable tags associated with the SIM cards",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "whitelabel-name",
			Usage:    "Service Provider Name (SPN) for the Whitelabel eSIM product. It will be displayed as the mobile service name by operating systems of smartphones. This parameter must only contain letters, numbers and whitespaces.",
			BodyPath: "whitelabel_name",
		},
	},
	Action:          handleActionsPurchaseCreate,
	HideHelpCommand: true,
}

func handleActionsPurchaseCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ActionPurchaseNewParams{}

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
	_, err = client.Actions.Purchase.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "actions:purchase create", obj, format, transform)
}
