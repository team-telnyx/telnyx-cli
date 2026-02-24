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

var alphanumericSenderIDsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new alphanumeric sender ID associated with a messaging profile.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "alphanumeric-sender-id",
			Usage:    "The alphanumeric sender ID string.",
			Required: true,
			BodyPath: "alphanumeric_sender_id",
		},
		&requestflag.Flag[string]{
			Name:     "messaging-profile-id",
			Usage:    "The messaging profile to associate the sender ID with.",
			Required: true,
			BodyPath: "messaging_profile_id",
		},
		&requestflag.Flag[string]{
			Name:     "us-long-code-fallback",
			Usage:    "A US long code number to use as fallback when sending to US destinations.",
			BodyPath: "us_long_code_fallback",
		},
	},
	Action:          handleAlphanumericSenderIDsCreate,
	HideHelpCommand: true,
}

var alphanumericSenderIDsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a specific alphanumeric sender ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleAlphanumericSenderIDsRetrieve,
	HideHelpCommand: true,
}

var alphanumericSenderIDsList = cli.Command{
	Name:    "list",
	Usage:   "List all alphanumeric sender IDs for the authenticated user.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "filter-messaging-profile-id",
			Usage:     "Filter by messaging profile ID.",
			QueryPath: "filter[messaging_profile_id]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "Page number.",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Page size.",
			Default:   20,
			QueryPath: "page[size]",
		},
	},
	Action:          handleAlphanumericSenderIDsList,
	HideHelpCommand: true,
}

var alphanumericSenderIDsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete an alphanumeric sender ID and disassociate it from its messaging profile.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleAlphanumericSenderIDsDelete,
	HideHelpCommand: true,
}

func handleAlphanumericSenderIDsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AlphanumericSenderIDNewParams{}

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
	_, err = client.AlphanumericSenderIDs.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "alphanumeric-sender-ids create", obj, format, transform)
}

func handleAlphanumericSenderIDsRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AlphanumericSenderIDs.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "alphanumeric-sender-ids retrieve", obj, format, transform)
}

func handleAlphanumericSenderIDsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AlphanumericSenderIDListParams{}

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
		_, err = client.AlphanumericSenderIDs.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "alphanumeric-sender-ids list", obj, format, transform)
	} else {
		iter := client.AlphanumericSenderIDs.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "alphanumeric-sender-ids list", iter, format, transform)
	}
}

func handleAlphanumericSenderIDsDelete(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AlphanumericSenderIDs.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "alphanumeric-sender-ids delete", obj, format, transform)
}
