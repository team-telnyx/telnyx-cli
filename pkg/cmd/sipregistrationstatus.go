// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/stainless-sdks/telnyx-cli/internal/apiquery"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
	"github.com/stainless-sdks/telnyx-go/v4"
	"github.com/stainless-sdks/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var sipRegistrationStatusRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns the live SIP registration state of a UAC connection: whether it is\ncurrently registered, when it last registered, and the last response Telnyx\nreceived from the registrar. Only `uac_external_credential` is supported today.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "credential-type",
			Usage:     "The kind of credential to look up. `uac_external_credential` is keyed by `connection_id`; `telephony_credential` is keyed by `username`.",
			Required:  true,
			QueryPath: "credential_type",
		},
		&requestflag.Flag[string]{
			Name:      "connection-id",
			Usage:     "Identifier of the UAC connection to look up. Required when `credential_type` is `uac_external_credential`.",
			QueryPath: "connection_id",
		},
		&requestflag.Flag[string]{
			Name:      "username",
			Usage:     "SIP username of the telephony credential to look up. Required when `credential_type` is `telephony_credential`.",
			QueryPath: "username",
		},
	},
	Action:          handleSipRegistrationStatusRetrieve,
	HideHelpCommand: true,
}

func handleSipRegistrationStatusRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.SipRegistrationStatusGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.SipRegistrationStatus.Get(ctx, params, options...)
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
		Title:          "sip-registration-status retrieve",
		Transform:      transform,
	})
}
