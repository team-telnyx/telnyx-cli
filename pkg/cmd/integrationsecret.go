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

var integrationSecretsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new secret with an associated identifier that can be used to securely\nintegrate with other services.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "identifier",
			Usage:    "The unique identifier of the secret.",
			Required: true,
			BodyPath: "identifier",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "The type of secret.",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "token",
			Usage:    "The token for the secret. Required for bearer type secrets, ignored otherwise.",
			BodyPath: "token",
		},
		&requestflag.Flag[string]{
			Name:     "password",
			Usage:    "The password for the secret. Required for basic type secrets, ignored otherwise.",
			BodyPath: "password",
		},
		&requestflag.Flag[string]{
			Name:     "username",
			Usage:    "The username for the secret. Required for basic type secrets, ignored otherwise.",
			BodyPath: "username",
		},
	},
	Action:          handleIntegrationSecretsCreate,
	HideHelpCommand: true,
}

var integrationSecretsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Retrieve a list of all integration secrets configured by the user.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[type]",
			QueryPath: "filter",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page[size]",
		},
	},
	Action:          handleIntegrationSecretsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.type",
			InnerField: "type",
		},
	},
})

var integrationSecretsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete an integration secret given its ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleIntegrationSecretsDelete,
	HideHelpCommand: true,
}

func handleIntegrationSecretsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.IntegrationSecretNewParams{}

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
	_, err = client.IntegrationSecrets.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "integration-secrets create", obj, format, transform)
}

func handleIntegrationSecretsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.IntegrationSecretListParams{}

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

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.IntegrationSecrets.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "integration-secrets list", obj, format, transform)
	} else {
		iter := client.IntegrationSecrets.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "integration-secrets list", iter, format, transform)
	}
}

func handleIntegrationSecretsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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

	return client.IntegrationSecrets.Delete(ctx, cmd.Value("id").(string), options...)
}
