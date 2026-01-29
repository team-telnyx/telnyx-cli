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

var privateWirelessGatewaysCreate = cli.Command{
	Name:    "create",
	Usage:   "Asynchronously create a Private Wireless Gateway for SIM cards for a previously\ncreated network. This operation may take several minutes so you can check the\nPrivate Wireless Gateway status at the section Get a Private Wireless Gateway.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The private wireless gateway name.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "network-id",
			Usage:    "The identification of the related network resource.",
			Required: true,
			BodyPath: "network_id",
		},
		&requestflag.Flag[string]{
			Name:     "region-code",
			Usage:    "The code of the region where the private wireless gateway will be assigned. A list of available regions can be found at the regions endpoint",
			BodyPath: "region_code",
		},
	},
	Action:          handlePrivateWirelessGatewaysCreate,
	HideHelpCommand: true,
}

var privateWirelessGatewaysRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve information about a Private Wireless Gateway.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handlePrivateWirelessGatewaysRetrieve,
	HideHelpCommand: true,
}

var privateWirelessGatewaysList = cli.Command{
	Name:    "list",
	Usage:   "Get all Private Wireless Gateways belonging to the user.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "filter-created-at",
			Usage:     "Private Wireless Gateway resource creation date.",
			QueryPath: "filter[created_at]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-ip-range",
			Usage:     "The IP address range of the Private Wireless Gateway.",
			QueryPath: "filter[ip_range]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-name",
			Usage:     "The name of the Private Wireless Gateway.",
			QueryPath: "filter[name]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-region-code",
			Usage:     "The name of the region where the Private Wireless Gateway is deployed.",
			QueryPath: "filter[region_code]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-updated-at",
			Usage:     "When the Private Wireless Gateway was last updated.",
			QueryPath: "filter[updated_at]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "The page number to load.",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "The size of the page.",
			Default:   20,
			QueryPath: "page[size]",
		},
	},
	Action:          handlePrivateWirelessGatewaysList,
	HideHelpCommand: true,
}

var privateWirelessGatewaysDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes the Private Wireless Gateway.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handlePrivateWirelessGatewaysDelete,
	HideHelpCommand: true,
}

func handlePrivateWirelessGatewaysCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PrivateWirelessGatewayNewParams{}

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
	_, err = client.PrivateWirelessGateways.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "private-wireless-gateways create", obj, format, transform)
}

func handlePrivateWirelessGatewaysRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.PrivateWirelessGateways.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "private-wireless-gateways retrieve", obj, format, transform)
}

func handlePrivateWirelessGatewaysList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PrivateWirelessGatewayListParams{}

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
		_, err = client.PrivateWirelessGateways.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "private-wireless-gateways list", obj, format, transform)
	} else {
		iter := client.PrivateWirelessGateways.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "private-wireless-gateways list", iter, format, transform)
	}
}

func handlePrivateWirelessGatewaysDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.PrivateWirelessGateways.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "private-wireless-gateways delete", obj, format, transform)
}
