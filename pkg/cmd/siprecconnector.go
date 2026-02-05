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

var siprecConnectorsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a new SIPREC connector configuration.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "host",
			Usage:    "Hostname/IPv4 address of the SIPREC SRS.",
			Required: true,
			BodyPath: "host",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name for the SIPREC connector resource.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[int64]{
			Name:     "port",
			Usage:    "Port for the SIPREC SRS.",
			Required: true,
			BodyPath: "port",
		},
		&requestflag.Flag[string]{
			Name:     "app-subdomain",
			Usage:    "Subdomain to route the call when using Telnyx SRS (optional for non-Telnyx SRS).",
			BodyPath: "app_subdomain",
		},
	},
	Action:          handleSiprecConnectorsCreate,
	HideHelpCommand: true,
}

var siprecConnectorsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns details of a stored SIPREC connector.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "connector-name",
			Required: true,
		},
	},
	Action:          handleSiprecConnectorsRetrieve,
	HideHelpCommand: true,
}

var siprecConnectorsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates a stored SIPREC connector configuration.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "connector-name",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "host",
			Usage:    "Hostname/IPv4 address of the SIPREC SRS.",
			Required: true,
			BodyPath: "host",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name for the SIPREC connector resource.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[int64]{
			Name:     "port",
			Usage:    "Port for the SIPREC SRS.",
			Required: true,
			BodyPath: "port",
		},
		&requestflag.Flag[string]{
			Name:     "app-subdomain",
			Usage:    "Subdomain to route the call when using Telnyx SRS (optional for non-Telnyx SRS).",
			BodyPath: "app_subdomain",
		},
	},
	Action:          handleSiprecConnectorsUpdate,
	HideHelpCommand: true,
}

var siprecConnectorsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a stored SIPREC connector.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "connector-name",
			Required: true,
		},
	},
	Action:          handleSiprecConnectorsDelete,
	HideHelpCommand: true,
}

func handleSiprecConnectorsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SiprecConnectorNewParams{}

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
	_, err = client.SiprecConnectors.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "siprec-connectors create", obj, format, transform)
}

func handleSiprecConnectorsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("connector-name") && len(unusedArgs) > 0 {
		cmd.Set("connector-name", unusedArgs[0])
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
	_, err = client.SiprecConnectors.Get(ctx, cmd.Value("connector-name").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "siprec-connectors retrieve", obj, format, transform)
}

func handleSiprecConnectorsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("connector-name") && len(unusedArgs) > 0 {
		cmd.Set("connector-name", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.SiprecConnectorUpdateParams{}

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
	_, err = client.SiprecConnectors.Update(
		ctx,
		cmd.Value("connector-name").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "siprec-connectors update", obj, format, transform)
}

func handleSiprecConnectorsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("connector-name") && len(unusedArgs) > 0 {
		cmd.Set("connector-name", unusedArgs[0])
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

	return client.SiprecConnectors.Delete(ctx, cmd.Value("connector-name").(string), options...)
}
