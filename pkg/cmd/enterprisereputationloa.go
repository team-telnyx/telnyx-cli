// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/telnyx-cli/internal/apiquery"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
	"github.com/stainless-sdks/telnyx-go/v4"
	"github.com/stainless-sdks/telnyx-go/v4/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var enterprisesReputationLoaUpdate = cli.Command{
	Name:    "update",
	Usage:   "Point the enterprise's reputation settings at a new signed LOA document. This\nresets LOA approval to `pending`; the new document must be approved before\nadditional phone numbers can be added.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "enterprise-id",
			Required:  true,
			PathParam: "enterprise_id",
		},
		&requestflag.Flag[string]{
			Name:     "loa-document-id",
			Usage:    "Id of the new signed LOA document (from the Telnyx Documents API). Changing it resets LOA approval; the new document must be approved before more numbers can be added.",
			Required: true,
			BodyPath: "loa_document_id",
		},
	},
	Action:          handleEnterprisesReputationLoaUpdate,
	HideHelpCommand: true,
}

var enterprisesReputationLoaRender = requestflag.WithInnerFlags(cli.Command{
	Name:    "render",
	Usage:   "Render the LOA for this enterprise as a PDF. The enterprise identity, address,\nand authorized-representative contact are taken from the enterprise record; the\noptional `agent` block is supplied only when a third-party partner manages the\nnumbers. The response is the PDF itself (unsigned unless a `signature` is\nprovided). Sign it and upload it to the Telnyx Documents API\n(`POST /v2/documents`, see https://developers.telnyx.com/api/documents) to\nobtain the `loa_document_id` required by `POST .../reputation`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "enterprise-id",
			Required:  true,
			PathParam: "enterprise_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "agent",
			Usage:    "Third-party reseller / partner managing the enterprise's phone numbers. Omit when the enterprise works directly with Telnyx.",
			BodyPath: "agent",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "signature",
			Usage:    "Optional signature embedded in the rendered PDF. When omitted the PDF is returned unsigned for the customer to sign and upload.",
			BodyPath: "signature",
		},
		&requestflag.Flag[string]{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "The file where the response contents will be stored. Use the value '-' to force output to stdout.",
		},
	},
	Action:          handleEnterprisesReputationLoaRender,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"agent": {
		&requestflag.InnerFlag[string]{
			Name:       "agent.administrative-area",
			InnerField: "administrative_area",
		},
		&requestflag.InnerFlag[string]{
			Name:       "agent.city",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "agent.contact-email",
			InnerField: "contact_email",
		},
		&requestflag.InnerFlag[string]{
			Name:       "agent.contact-name",
			InnerField: "contact_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "agent.contact-phone",
			InnerField: "contact_phone",
		},
		&requestflag.InnerFlag[string]{
			Name:       "agent.contact-title",
			InnerField: "contact_title",
		},
		&requestflag.InnerFlag[string]{
			Name:       "agent.country",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "agent.legal-name",
			InnerField: "legal_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "agent.postal-code",
			InnerField: "postal_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "agent.street-address",
			InnerField: "street_address",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "agent.dba",
			InnerField: "dba",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "agent.extended-address",
			InnerField: "extended_address",
		},
	},
	"signature": {
		&requestflag.InnerFlag[string]{
			Name:       "signature.image-base64",
			Usage:      "Base64-encoded signature image.",
			InnerField: "image_base64",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "signature.signer-name",
			InnerField: "signer_name",
		},
	},
})

func handleEnterprisesReputationLoaUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("enterprise-id") && len(unusedArgs) > 0 {
		cmd.Set("enterprise-id", unusedArgs[0])
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

	params := telnyx.EnterpriseReputationLoaUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Enterprises.Reputation.Loa.Update(
		ctx,
		cmd.Value("enterprise-id").(string),
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
		Title:          "enterprises:reputation:loa update",
		Transform:      transform,
	})
}

func handleEnterprisesReputationLoaRender(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("enterprise-id") && len(unusedArgs) > 0 {
		cmd.Set("enterprise-id", unusedArgs[0])
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

	params := telnyx.EnterpriseReputationLoaRenderParams{}

	response, err := client.Enterprises.Reputation.Loa.Render(
		ctx,
		cmd.Value("enterprise-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}
	message, err := writeBinaryResponse(response, os.Stdout, cmd.String("output"))
	if message != "" {
		fmt.Println(message)
	}
	return err
}
