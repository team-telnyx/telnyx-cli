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

var accessIPAddressCreate = cli.Command{
	Name:    "create",
	Usage:   "Create new Access IP Address",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "ip-address",
			Required: true,
			BodyPath: "ip_address",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
	},
	Action:          handleAccessIPAddressCreate,
	HideHelpCommand: true,
}

var accessIPAddressRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve an access IP address",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "access-ip-address-id",
			Required: true,
		},
	},
	Action:          handleAccessIPAddressRetrieve,
	HideHelpCommand: true,
}

var accessIPAddressList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List all Access IP Addresses",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[ip_source], filter[ip_address], filter[created_at]. Supports complex bracket operations for dynamic filtering.",
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
	Action:          handleAccessIPAddressList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[any]{
			Name:       "filter.created-at",
			Usage:      "Filter by exact creation date-time",
			InnerField: "created_at",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.ip-address",
			Usage:      "Filter by IP address",
			InnerField: "ip_address",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.ip-source",
			Usage:      "Filter by IP source",
			InnerField: "ip_source",
		},
	},
})

var accessIPAddressDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete access IP address",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "access-ip-address-id",
			Required: true,
		},
	},
	Action:          handleAccessIPAddressDelete,
	HideHelpCommand: true,
}

func handleAccessIPAddressCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AccessIPAddressNewParams{}

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
	_, err = client.AccessIPAddress.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "access-ip-address create", obj, format, transform)
}

func handleAccessIPAddressRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("access-ip-address-id") && len(unusedArgs) > 0 {
		cmd.Set("access-ip-address-id", unusedArgs[0])
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
	_, err = client.AccessIPAddress.Get(ctx, cmd.Value("access-ip-address-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "access-ip-address retrieve", obj, format, transform)
}

func handleAccessIPAddressList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AccessIPAddressListParams{}

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
		_, err = client.AccessIPAddress.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "access-ip-address list", obj, format, transform)
	} else {
		iter := client.AccessIPAddress.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "access-ip-address list", iter, format, transform)
	}
}

func handleAccessIPAddressDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("access-ip-address-id") && len(unusedArgs) > 0 {
		cmd.Set("access-ip-address-id", unusedArgs[0])
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
	_, err = client.AccessIPAddress.Delete(ctx, cmd.Value("access-ip-address-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "access-ip-address delete", obj, format, transform)
}
