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

var actionsRegisterCreate = cli.Command{
	Name:    "create",
	Usage:   "Register the SIM cards associated with the provided registration codes to the\ncurrent user's account.<br/><br/> If <code>sim_card_group_id</code> is provided,\nthe SIM cards will be associated with that group. Otherwise, the default group\nfor the current user will be used.<br/><br/>",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "registration-code",
			Required: true,
			BodyPath: "registration_codes",
		},
		&requestflag.Flag[string]{
			Name:     "sim-card-group-id",
			Usage:    "The group SIMCardGroup identification. This attribute can be <code>null</code> when it's present in an associated resource.",
			BodyPath: "sim_card_group_id",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "Status on which the SIM card will be set after being successful registered.",
			Default:  "enabled",
			BodyPath: "status",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Searchable tags associated with the SIM card",
			BodyPath: "tags",
		},
	},
	Action:          handleActionsRegisterCreate,
	HideHelpCommand: true,
}

func handleActionsRegisterCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.ActionRegisterNewParams{}

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
	_, err = client.Actions.Register.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "actions:register create", obj, format, transform)
}
