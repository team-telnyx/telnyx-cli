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

var fqdnsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new FQDN object.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "connection-id",
			Usage:    "ID of the FQDN connection to which this IP should be attached.",
			Required: true,
			BodyPath: "connection_id",
		},
		&requestflag.Flag[string]{
			Name:     "dns-record-type",
			Usage:    "The DNS record type for the FQDN. For cases where a port is not set, the DNS record type must be 'srv'. For cases where a port is set, the DNS record type must be 'a'. If the DNS record type is 'a' and a port is not specified, 5060 will be used.",
			Required: true,
			BodyPath: "dns_record_type",
		},
		&requestflag.Flag[string]{
			Name:     "fqdn",
			Usage:    "FQDN represented by this resource.",
			Required: true,
			BodyPath: "fqdn",
		},
		&requestflag.Flag[any]{
			Name:     "port",
			Usage:    "Port to use when connecting to this FQDN.",
			Default:  5060,
			BodyPath: "port",
		},
	},
	Action:          handleFqdnsCreate,
	HideHelpCommand: true,
}

var fqdnsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Return the details regarding a specific FQDN.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleFqdnsRetrieve,
	HideHelpCommand: true,
}

var fqdnsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update the details of a specific FQDN.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "connection-id",
			Usage:    "ID of the FQDN connection to which this IP should be attached.",
			BodyPath: "connection_id",
		},
		&requestflag.Flag[string]{
			Name:     "dns-record-type",
			Usage:    "The DNS record type for the FQDN. For cases where a port is not set, the DNS record type must be 'srv'. For cases where a port is set, the DNS record type must be 'a'. If the DNS record type is 'a' and a port is not specified, 5060 will be used.",
			BodyPath: "dns_record_type",
		},
		&requestflag.Flag[string]{
			Name:     "fqdn",
			Usage:    "FQDN represented by this resource.",
			BodyPath: "fqdn",
		},
		&requestflag.Flag[any]{
			Name:     "port",
			Usage:    "Port to use when connecting to this FQDN.",
			Default:  5060,
			BodyPath: "port",
		},
	},
	Action:          handleFqdnsUpdate,
	HideHelpCommand: true,
}

var fqdnsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Get all FQDNs belonging to the user that match the given filters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[connection_id], filter[fqdn], filter[port], filter[dns_record_type]",
			QueryPath: "filter",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "page",
			Usage:     "Consolidated page parameter (deepObject style). Originally: page[size], page[number]",
			QueryPath: "page",
		},
	},
	Action:          handleFqdnsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.connection-id",
			Usage:      "ID of the FQDN connection to which the FQDN belongs.",
			InnerField: "connection_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.dns-record-type",
			Usage:      "DNS record type used by the FQDN.",
			InnerField: "dns_record_type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.fqdn",
			Usage:      "FQDN represented by the resource.",
			InnerField: "fqdn",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "filter.port",
			Usage:      "Port to use when connecting to the FQDN.",
			InnerField: "port",
		},
	},
	"page": {
		&requestflag.InnerFlag[int64]{
			Name:       "page.number",
			Usage:      "The page number to load",
			InnerField: "number",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "page.size",
			Usage:      "The size of the page",
			InnerField: "size",
		},
	},
})

var fqdnsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete an FQDN.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleFqdnsDelete,
	HideHelpCommand: true,
}

func handleFqdnsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.FqdnNewParams{}

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
	_, err = client.Fqdns.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "fqdns create", obj, format, transform)
}

func handleFqdnsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Fqdns.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "fqdns retrieve", obj, format, transform)
}

func handleFqdnsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.FqdnUpdateParams{}

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
	_, err = client.Fqdns.Update(
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
	return ShowJSON(os.Stdout, "fqdns update", obj, format, transform)
}

func handleFqdnsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.FqdnListParams{}

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
		_, err = client.Fqdns.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "fqdns list", obj, format, transform)
	} else {
		iter := client.Fqdns.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "fqdns list", iter, format, transform)
	}
}

func handleFqdnsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Fqdns.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "fqdns delete", obj, format, transform)
}
