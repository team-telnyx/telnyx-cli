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

var virtualCrossConnectsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new Virtual Cross Connect.<br /><br />For AWS and GCE, you have the\noption of creating the primary connection first and the secondary connection\nlater. You also have the option of disabling the primary and/or secondary\nconnections at any time and later re-enabling them. With Azure, you do not have\nthis option. Azure requires both the primary and secondary connections to be\ncreated at the same time and they can not be independantly disabled.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[float64]{
			Name:     "bandwidth-mbps",
			Usage:    "The desired throughput in Megabits per Second (Mbps) for your Virtual Cross Connect.<br /><br />The available bandwidths can be found using the /virtual_cross_connect_regions endpoint.",
			BodyPath: "bandwidth_mbps",
		},
		&requestflag.Flag[float64]{
			Name:     "bgp-asn",
			Usage:    "The Border Gateway Protocol (BGP) Autonomous System Number (ASN). If null, value will be assigned by Telnyx.",
			BodyPath: "bgp_asn",
		},
		&requestflag.Flag[string]{
			Name:     "cloud-provider",
			Usage:    "The Virtual Private Cloud with which you would like to establish a cross connect.",
			BodyPath: "cloud_provider",
		},
		&requestflag.Flag[string]{
			Name:     "cloud-provider-region",
			Usage:    "The region where your Virtual Private Cloud hosts are located.<br /><br />The available regions can be found using the /virtual_cross_connect_regions endpoint.",
			BodyPath: "cloud_provider_region",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "A user specified name for the interface.",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "network-id",
			Usage:    "The id of the network associated with the interface.",
			BodyPath: "network_id",
		},
		&requestflag.Flag[string]{
			Name:     "primary-bgp-key",
			Usage:    "The authentication key for BGP peer configuration.",
			BodyPath: "primary_bgp_key",
		},
		&requestflag.Flag[string]{
			Name:     "primary-cloud-account-id",
			Usage:    "The identifier for your Virtual Private Cloud. The number will be different based upon your Cloud provider.",
			BodyPath: "primary_cloud_account_id",
		},
		&requestflag.Flag[string]{
			Name:     "primary-cloud-ip",
			Usage:    "The IP address assigned for your side of the Virtual Cross Connect.<br /><br />If none is provided, one will be generated for you.<br /><br />This value should be null for GCE as Google will only inform you of your assigned IP once the connection has been accepted.",
			BodyPath: "primary_cloud_ip",
		},
		&requestflag.Flag[string]{
			Name:     "primary-telnyx-ip",
			Usage:    "The IP address assigned to the Telnyx side of the Virtual Cross Connect.<br /><br />If none is provided, one will be generated for you.<br /><br />This value should be null for GCE as Google will only inform you of your assigned IP once the connection has been accepted.",
			BodyPath: "primary_telnyx_ip",
		},
		&requestflag.Flag[string]{
			Name:     "secondary-bgp-key",
			Usage:    "The authentication key for BGP peer configuration.",
			BodyPath: "secondary_bgp_key",
		},
		&requestflag.Flag[string]{
			Name:     "secondary-cloud-account-id",
			Usage:    "The identifier for your Virtual Private Cloud. The number will be different based upon your Cloud provider.<br /><br />This attribute is only necessary for GCE.",
			BodyPath: "secondary_cloud_account_id",
		},
		&requestflag.Flag[string]{
			Name:     "secondary-cloud-ip",
			Usage:    "The IP address assigned for your side of the Virtual Cross Connect.<br /><br />If none is provided, one will be generated for you.<br /><br />This value should be null for GCE as Google will only inform you of your assigned IP once the connection has been accepted.",
			BodyPath: "secondary_cloud_ip",
		},
		&requestflag.Flag[string]{
			Name:     "secondary-telnyx-ip",
			Usage:    "The IP address assigned to the Telnyx side of the Virtual Cross Connect.<br /><br />If none is provided, one will be generated for you.<br /><br />This value should be null for GCE as Google will only inform you of your assigned IP once the connection has been accepted.",
			BodyPath: "secondary_telnyx_ip",
		},
	},
	Action:          handleVirtualCrossConnectsCreate,
	HideHelpCommand: true,
}

var virtualCrossConnectsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a Virtual Cross Connect.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleVirtualCrossConnectsRetrieve,
	HideHelpCommand: true,
}

var virtualCrossConnectsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update the Virtual Cross Connect.<br /><br />Cloud IPs can only be patched\nduring the `created` state, as GCE will only inform you of your generated IP\nonce the pending connection requested has been accepted. Once the Virtual Cross\nConnect has moved to `provisioning`, the IPs can no longer be\npatched.<br /><br />Once the Virtual Cross Connect has moved to `provisioned`\nand you are ready to enable routing, you can toggle the routing announcements to\n`true`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "primary-cloud-ip",
			Usage:    "The IP address assigned for your side of the Virtual Cross Connect.<br /><br />If none is provided, one will be generated for you.<br /><br />This value can not be patched once the VXC has bene provisioned.",
			BodyPath: "primary_cloud_ip",
		},
		&requestflag.Flag[bool]{
			Name:     "primary-enabled",
			Usage:    "Indicates whether the primary circuit is enabled. Setting this to `false` will disable the circuit.",
			BodyPath: "primary_enabled",
		},
		&requestflag.Flag[bool]{
			Name:     "primary-routing-announcement",
			Usage:    "Whether the primary BGP route is being announced.",
			BodyPath: "primary_routing_announcement",
		},
		&requestflag.Flag[string]{
			Name:     "secondary-cloud-ip",
			Usage:    "The IP address assigned for your side of the Virtual Cross Connect.<br /><br />If none is provided, one will be generated for you.<br /><br />This value can not be patched once the VXC has bene provisioned.",
			BodyPath: "secondary_cloud_ip",
		},
		&requestflag.Flag[bool]{
			Name:     "secondary-enabled",
			Usage:    "Indicates whether the secondary circuit is enabled. Setting this to `false` will disable the circuit.",
			BodyPath: "secondary_enabled",
		},
		&requestflag.Flag[bool]{
			Name:     "secondary-routing-announcement",
			Usage:    "Whether the secondary BGP route is being announced.",
			BodyPath: "secondary_routing_announcement",
		},
	},
	Action:          handleVirtualCrossConnectsUpdate,
	HideHelpCommand: true,
}

var virtualCrossConnectsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List all Virtual Cross Connects.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[network_id]",
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
	Action:          handleVirtualCrossConnectsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.network-id",
			Usage:      "The associated network id to filter on.",
			InnerField: "network_id",
		},
	},
})

var virtualCrossConnectsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a Virtual Cross Connect.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleVirtualCrossConnectsDelete,
	HideHelpCommand: true,
}

func handleVirtualCrossConnectsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.VirtualCrossConnectNewParams{}

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
	_, err = client.VirtualCrossConnects.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "virtual-cross-connects create", obj, format, transform)
}

func handleVirtualCrossConnectsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.VirtualCrossConnects.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "virtual-cross-connects retrieve", obj, format, transform)
}

func handleVirtualCrossConnectsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.VirtualCrossConnectUpdateParams{}

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
	_, err = client.VirtualCrossConnects.Update(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "virtual-cross-connects update", obj, format, transform)
}

func handleVirtualCrossConnectsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.VirtualCrossConnectListParams{}

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
		_, err = client.VirtualCrossConnects.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "virtual-cross-connects list", obj, format, transform)
	} else {
		iter := client.VirtualCrossConnects.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "virtual-cross-connects list", iter, format, transform)
	}
}

func handleVirtualCrossConnectsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.VirtualCrossConnects.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "virtual-cross-connects delete", obj, format, transform)
}
