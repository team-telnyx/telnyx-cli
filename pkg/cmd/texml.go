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

var texmlSecrets = cli.Command{
	Name:    "secrets",
	Usage:   "Create a TeXML secret which can be later used as a Dynamic Parameter for TeXML\nwhen using Mustache Templates in your TeXML. In your TeXML you will be able to\nuse your secret name, and this name will be replaced by the actual secret value\nwhen processing the TeXML on Telnyx side. The secrets are not visible in any\nlogs.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name used as a reference for the secret, if the name already exists within the account its value will be replaced",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "value",
			Usage:    "Secret value which will be used when rendering the TeXML template",
			Required: true,
			BodyPath: "value",
		},
	},
	Action:          handleTexmlSecrets,
	HideHelpCommand: true,
}

func handleTexmlSecrets(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TexmlSecretsParams{}

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
	_, err = client.Texml.Secrets(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "texml secrets", obj, format, transform)
}
