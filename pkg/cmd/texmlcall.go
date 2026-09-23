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

var texmlCallsCreate = cli.Command{
	Name:    "create",
	Usage:   "Initiate an outbound TeXML call using a TeXML application connection ID, not an\naccount SID. Request parameter names are case-sensitive. From and To are\nrequired; Texml supplies inline instructions and Url overrides the application\nXML request URL. When neither is supplied, the application configuration\nsupplies the instructions. The response is a flat call object without a data\nwrapper.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "connection-id",
			Required:  true,
			PathParam: "connection_id",
		},
		&requestflag.Flag[string]{
			Name:     "from",
			Usage:    "The E.164-formatted phone number or SIP URI to present as the caller.",
			Required: true,
			BodyPath: "From",
		},
		&requestflag.Flag[string]{
			Name:     "to",
			Usage:    "The E.164-formatted phone number or SIP URI to call.",
			Required: true,
			BodyPath: "To",
		},
		&requestflag.Flag[string]{
			Name:     "method",
			Usage:    "HTTP method used to retrieve TeXML instructions from Url.",
			BodyPath: "Method",
		},
		&requestflag.Flag[string]{
			Name:     "texml",
			Usage:    "Inline TeXML instructions to execute when the call is answered.",
			BodyPath: "Texml",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "The URL from which to retrieve TeXML instructions. Overrides the TeXML application XML request URL.",
			BodyPath: "Url",
		},
	},
	Action:          handleTexmlCallsCreate,
	HideHelpCommand: true,
}

func handleTexmlCallsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("connection-id") && len(unusedArgs) > 0 {
		cmd.Set("connection-id", unusedArgs[0])
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

	params := telnyx.TexmlCallNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Texml.Calls.New(
		ctx,
		cmd.Value("connection-id").(string),
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
		Title:          "texml:calls create",
		Transform:      transform,
	})
}
