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

var trafficPolicyProfilesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new traffic policy profile. At least one of `services`, `ip_ranges`, or\n`domains` must be provided.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "The type of the traffic policy profile.",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[[]string]{
			Name:     "domain",
			Usage:    "Array of domain names.",
			BodyPath: "domains",
		},
		&requestflag.Flag[[]string]{
			Name:     "ip-range",
			Usage:    "Array of IP ranges in CIDR notation.",
			BodyPath: "ip_ranges",
		},
		&requestflag.Flag[int64]{
			Name:     "limit-bw-kbps",
			Usage:    "Bandwidth limit in kbps. Must be 512 or 1024.",
			BodyPath: "limit_bw_kbps",
		},
		&requestflag.Flag[[]string]{
			Name:     "service",
			Usage:    "Array of PCEF service IDs to include in the profile.",
			BodyPath: "services",
		},
	},
	Action:          handleTrafficPolicyProfilesCreate,
	HideHelpCommand: true,
}

var trafficPolicyProfilesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns the details regarding a specific traffic policy profile.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleTrafficPolicyProfilesRetrieve,
	HideHelpCommand: true,
}

var trafficPolicyProfilesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates a traffic policy profile.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[[]string]{
			Name:     "domain",
			Usage:    "Array of domain names.",
			BodyPath: "domains",
		},
		&requestflag.Flag[[]string]{
			Name:     "ip-range",
			Usage:    "Array of IP ranges in CIDR notation.",
			BodyPath: "ip_ranges",
		},
		&requestflag.Flag[any]{
			Name:     "limit-bw-kbps",
			Usage:    "Bandwidth limit in kbps. Must be 512 or 1024, or null to remove.",
			BodyPath: "limit_bw_kbps",
		},
		&requestflag.Flag[[]string]{
			Name:     "service",
			Usage:    "Array of PCEF service IDs to include in the profile.",
			BodyPath: "services",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "The type of the traffic policy profile.",
			BodyPath: "type",
		},
	},
	Action:          handleTrafficPolicyProfilesUpdate,
	HideHelpCommand: true,
}

var trafficPolicyProfilesList = cli.Command{
	Name:    "list",
	Usage:   "Get all traffic policy profiles belonging to the user that match the given\nfilters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "filter-service",
			Usage:     "Filter by service ID.",
			QueryPath: "filter[service]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-type",
			Usage:     "Filter by traffic policy profile type.",
			QueryPath: "filter[type]",
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
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Sorts traffic policy profiles by the given field. Defaults to ascending order unless field is prefixed with a minus sign.",
			QueryPath: "sort",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleTrafficPolicyProfilesList,
	HideHelpCommand: true,
}

var trafficPolicyProfilesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes the traffic policy profile.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleTrafficPolicyProfilesDelete,
	HideHelpCommand: true,
}

var trafficPolicyProfilesListServices = cli.Command{
	Name:    "list-services",
	Usage:   "Get all available PCEF services that can be used in traffic policy profiles.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "filter-group",
			Usage:     "Filter services by group.",
			QueryPath: "filter[group]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-name",
			Usage:     "Filter services by name.",
			QueryPath: "filter[name]",
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
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleTrafficPolicyProfilesListServices,
	HideHelpCommand: true,
}

func handleTrafficPolicyProfilesCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TrafficPolicyProfileNewParams{}

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
	_, err = client.TrafficPolicyProfiles.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "traffic-policy-profiles create", obj, format, explicitFormat, transform)
}

func handleTrafficPolicyProfilesRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.TrafficPolicyProfiles.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "traffic-policy-profiles retrieve", obj, format, explicitFormat, transform)
}

func handleTrafficPolicyProfilesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TrafficPolicyProfileUpdateParams{}

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
	_, err = client.TrafficPolicyProfiles.Update(
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
	return ShowJSON(os.Stdout, os.Stderr, "traffic-policy-profiles update", obj, format, explicitFormat, transform)
}

func handleTrafficPolicyProfilesList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TrafficPolicyProfileListParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.TrafficPolicyProfiles.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, os.Stderr, "traffic-policy-profiles list", obj, format, explicitFormat, transform)
	} else {
		iter := client.TrafficPolicyProfiles.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, os.Stderr, "traffic-policy-profiles list", iter, format, explicitFormat, transform, maxItems)
	}
}

func handleTrafficPolicyProfilesDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.TrafficPolicyProfiles.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "traffic-policy-profiles delete", obj, format, explicitFormat, transform)
}

func handleTrafficPolicyProfilesListServices(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.TrafficPolicyProfileListServicesParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.TrafficPolicyProfiles.ListServices(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, os.Stderr, "traffic-policy-profiles list-services", obj, format, explicitFormat, transform)
	} else {
		iter := client.TrafficPolicyProfiles.ListServicesAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, os.Stderr, "traffic-policy-profiles list-services", iter, format, explicitFormat, transform, maxItems)
	}
}
