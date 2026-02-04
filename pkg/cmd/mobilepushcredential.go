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

var mobilePushCredentialsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a new mobile push credential",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "create-mobile-push-credential-request",
			Required: true,
			BodyRoot: true,
		},
	},
	Action:          handleMobilePushCredentialsCreate,
	HideHelpCommand: true,
}

var mobilePushCredentialsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves mobile push credential based on the given `push_credential_id`",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "push-credential-id",
			Required: true,
		},
	},
	Action:          handleMobilePushCredentialsRetrieve,
	HideHelpCommand: true,
}

var mobilePushCredentialsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List mobile push credentials",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[type], filter[alias]",
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
	Action:          handleMobilePushCredentialsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.alias",
			Usage:      "Unique mobile push credential alias",
			InnerField: "alias",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.type",
			Usage:      "type of mobile push credentials",
			InnerField: "type",
		},
	},
})

var mobilePushCredentialsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a mobile push credential based on the given `push_credential_id`",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "push-credential-id",
			Required: true,
		},
	},
	Action:          handleMobilePushCredentialsDelete,
	HideHelpCommand: true,
}

func handleMobilePushCredentialsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MobilePushCredentialNewParams{}

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
	_, err = client.MobilePushCredentials.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "mobile-push-credentials create", obj, format, transform)
}

func handleMobilePushCredentialsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("push-credential-id") && len(unusedArgs) > 0 {
		cmd.Set("push-credential-id", unusedArgs[0])
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
	_, err = client.MobilePushCredentials.Get(ctx, cmd.Value("push-credential-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "mobile-push-credentials retrieve", obj, format, transform)
}

func handleMobilePushCredentialsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MobilePushCredentialListParams{}

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
		_, err = client.MobilePushCredentials.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "mobile-push-credentials list", obj, format, transform)
	} else {
		iter := client.MobilePushCredentials.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "mobile-push-credentials list", iter, format, transform)
	}
}

func handleMobilePushCredentialsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("push-credential-id") && len(unusedArgs) > 0 {
		cmd.Set("push-credential-id", unusedArgs[0])
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

	return client.MobilePushCredentials.Delete(ctx, cmd.Value("push-credential-id").(string), options...)
}
