// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/team-telnyx/telnyx-cli/internal/apiquery"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var networkRoutesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Adds a route to a Network. Traffic on the Network to an address in `cidr` goes\nvia `next_hop`. With a tunnel as the next hop, functions on the Network reach\nthe CIDR through the tunnel without a redeploy. A Network holds each CIDR once.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "cidr",
			Usage:    "The IPv4 subnet to route, such as a subnet behind the tunnel. It must be a canonical CIDR. Default, loopback, link-local, multicast, and shared address space (`100.64.0.0/10`) ranges are rejected. So are ranges that overlap the Tailscale Subnet Router ranges, `172.27.254.0/24` and `198.18.0.0/20`, or a subnet already on the Network.",
			Required: true,
			BodyPath: "cidr",
		},
		&requestflag.Flag[string]{
			Name:     "network-id",
			Usage:    "The id of the Network to add the route to.",
			Required: true,
			BodyPath: "network_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "next-hop",
			Usage:    "Where the Network sends traffic for `cidr`.",
			Required: true,
			BodyPath: "next_hop",
		},
	},
	Action:          handleNetworkRoutesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"next-hop": {
		&requestflag.InnerFlag[string]{
			Name:       "next-hop.id",
			Usage:      "The id of the tunnel.",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "next-hop.type",
			Usage:      "The type of the next hop.",
			InnerField: "type",
		},
	},
})

var networkRoutesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a route and its status.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleNetworkRoutesRetrieve,
	HideHelpCommand: true,
}

var networkRoutesList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Lists the routes of the account's Networks. A Network's list is its effective\nroute table: the routes added with `POST /network_routes` (`origin` `static`)\nand the routes that its attachments install (`origin` `attachment`). Lookup is\nlongest prefix within the Network.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[network_id], filter[origin], filter[next_hop.type], filter[cidr], filter[status]",
			Required:  true,
			QueryPath: "filter",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "page",
			Usage:     "Consolidated page parameter (deepObject style). Originally: page[after], page[limit]",
			QueryPath: "page",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Sorts the results by a field, ascending, or descending with a leading `-`. The default is `-created_at`.",
			QueryPath: "sort",
		},
	},
	Action:          handleNetworkRoutesList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.network-id",
			Usage:      "The id of the Network whose route table is returned. Required.",
			InnerField: "network_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.cidr",
			Usage:      "Only routes with this CIDR.",
			InnerField: "cidr",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.next-hop-type",
			Usage:      "Only routes whose next hop has this type.",
			InnerField: "next_hop.type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.origin",
			Usage:      "Only routes of this origin.",
			InnerField: "origin",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.status",
			Usage:      "Only routes with this status.",
			InnerField: "status",
		},
	},
	"page": {
		&requestflag.InnerFlag[string]{
			Name:       "page.after",
			Usage:      "The cursor from `meta.cursors.after` of the previous page.",
			InnerField: "after",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "page.limit",
			Usage:      "The number of results per page.",
			InnerField: "limit",
		},
	},
})

var networkRoutesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a route that was added with `POST /network_routes`. Traffic to its CIDR\nstops going via the next hop at once. Routes that an attachment installs\n(`origin` `attachment`) leave with the attachment and cannot be deleted here.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleNetworkRoutesDelete,
	HideHelpCommand: true,
}

func handleNetworkRoutesCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := telnyx.NetworkRouteNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.NetworkRoutes.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "network-routes create",
		Transform:      transform,
	})
}

func handleNetworkRoutesRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.NetworkRoutes.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "network-routes retrieve",
		Transform:      transform,
	})
}

func handleNetworkRoutesList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := telnyx.NetworkRouteListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.NetworkRoutes.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "network-routes list",
		Transform:      transform,
	})
}

func handleNetworkRoutesDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.NetworkRoutes.Delete(ctx, cmd.Value("id").(string), options...)
}
