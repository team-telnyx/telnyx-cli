// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/telnyx-cli/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var wellKnownRetrieveAuthorizationServerMetadata = cli.Command{
	Name:            "retrieve-authorization-server-metadata",
	Usage:           "OAuth 2.0 Authorization Server Metadata (RFC 8414)",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleWellKnownRetrieveAuthorizationServerMetadata,
	HideHelpCommand: true,
}

var wellKnownRetrieveProtectedResourceMetadata = cli.Command{
	Name:            "retrieve-protected-resource-metadata",
	Usage:           "OAuth 2.0 Protected Resource Metadata for resource discovery",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleWellKnownRetrieveProtectedResourceMetadata,
	HideHelpCommand: true,
}

func handleWellKnownRetrieveAuthorizationServerMetadata(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.WellKnown.GetAuthorizationServerMetadata(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "well-known retrieve-authorization-server-metadata", obj, format, transform)
}

func handleWellKnownRetrieveProtectedResourceMetadata(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.WellKnown.GetProtectedResourceMetadata(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "well-known retrieve-protected-resource-metadata", obj, format, transform)
}
