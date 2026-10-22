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

var tailscaleSubnetRoutersCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a Tailscale Subnet Router on a Network and enrolls it in the user's\ntailnet with the auth key. The response returns the router with `status`\n`provisioning`. Poll `GET /tailscale_subnet_routers/{id}` until it is\n`provisioned`. Then approve its routes in the tailnet, unless an auto-approver\ncovers them.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "advertised-route",
			Usage:    "IPv4 subnets of the Network that the router advertises to the tailnet. Each entry is a canonical CIDR that the Network routes. Default, loopback, link-local, multicast, and Tailscale (`100.64.0.0/10`) ranges are rejected. So are ranges that overlap the router's own, `172.27.254.0/24` and `198.18.0.0/20`.",
			Required: true,
			BodyPath: "advertised_routes",
		},
		&requestflag.Flag[string]{
			Name:     "auth-key",
			Usage:    "A Tailscale auth key that enrolls the router in the user's tailnet: reusable, pre-approved, tagged, and not ephemeral. Telnyx stores it encrypted until enrollment succeeds, then discards it. It is never returned.",
			Required: true,
			BodyPath: "auth_key",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "A user specified name for the router.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "network-id",
			Usage:    "The id of the Network to attach the router to.",
			Required: true,
			BodyPath: "network_id",
		},
		&requestflag.Flag[string]{
			Name:     "region-code",
			Usage:    "The code of the region to create the router in. `GET /network_coverage` lists the regions whose `available_services` include `tailscale_subnet_router`.",
			Required: true,
			BodyPath: "region_code",
		},
	},
	Action:          handleTailscaleSubnetRoutersCreate,
	HideHelpCommand: true,
}

var tailscaleSubnetRoutersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns the details of a single Tailscale Subnet Router by its identifier,\nincluding its enrollment state.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleTailscaleSubnetRoutersRetrieve,
	HideHelpCommand: true,
}

var tailscaleSubnetRoutersUpdate = cli.Command{
	Name:    "update",
	Usage:   "Renames the router, changes its advertised routes, or retries a failed\nenrollment. Any change queues a new provisioning run and returns the router with\n`status` `provisioning`. The router keeps its machine identity on the tailnet.\nSend an empty object to retry a `failed` router with its stored auth key.\nInclude a new `auth_key` only while the router has never enrolled. The router's\nNetwork and region cannot change: create a new router instead.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[[]string]{
			Name:     "advertised-route",
			Usage:    "IPv4 subnets of the Network that the router advertises to the tailnet. Each entry is a canonical CIDR that the Network routes. Default, loopback, link-local, multicast, and Tailscale (`100.64.0.0/10`) ranges are rejected. So are ranges that overlap the router's own, `172.27.254.0/24` and `198.18.0.0/20`.",
			BodyPath: "advertised_routes",
		},
		&requestflag.Flag[string]{
			Name:     "auth-key",
			Usage:    "A replacement Tailscale auth key. Accepted only while the router has never enrolled, for example after a `failed` enrollment. Never returned.",
			BodyPath: "auth_key",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "A user specified name for the router.",
			BodyPath: "name",
		},
	},
	Action:          handleTailscaleSubnetRoutersUpdate,
	HideHelpCommand: true,
}

var tailscaleSubnetRoutersList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Returns a paginated list of the Tailscale Subnet Routers on the account, with\nsupport for filtering by Network.",
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
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleTailscaleSubnetRoutersList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.network-id",
			Usage:      "The id of the Network to filter on.",
			InnerField: "network_id",
		},
	},
})

var tailscaleSubnetRoutersDelete = cli.Command{
	Name:    "delete",
	Usage:   "Removes the router from the Network and logs it out of the tailnet. The response\nreturns the router with `status` `deleting`. The record disappears once removal\ncompletes. A machine entry can remain in the tailnet's admin console until it is\nremoved there.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleTailscaleSubnetRoutersDelete,
	HideHelpCommand: true,
}

func handleTailscaleSubnetRoutersCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.TailscaleSubnetRouterNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.TailscaleSubnetRouters.New(ctx, params, options...)
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
		Title:          "tailscale-subnet-routers create",
		Transform:      transform,
	})
}

func handleTailscaleSubnetRoutersRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.TailscaleSubnetRouters.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "tailscale-subnet-routers retrieve",
		Transform:      transform,
	})
}

func handleTailscaleSubnetRoutersUpdate(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := telnyx.TailscaleSubnetRouterUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.TailscaleSubnetRouters.Update(
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "tailscale-subnet-routers update",
		Transform:      transform,
	})
}

func handleTailscaleSubnetRoutersList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.TailscaleSubnetRouterListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.TailscaleSubnetRouters.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "tailscale-subnet-routers list",
			Transform:      transform,
		})
	} else {
		iter := client.TailscaleSubnetRouters.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "tailscale-subnet-routers list",
			Transform:      transform,
		})
	}
}

func handleTailscaleSubnetRoutersDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.TailscaleSubnetRouters.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "tailscale-subnet-routers delete",
		Transform:      transform,
	})
}
