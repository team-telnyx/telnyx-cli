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

var infringementClaimsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a single claim by id. Returns `404` if the claim does not exist or is\nnot against a DIR you own.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "claim-id",
			Required:  true,
			PathParam: "claim_id",
		},
	},
	Action:          handleInfringementClaimsRetrieve,
	HideHelpCommand: true,
}

var infringementClaimsContest = requestflag.WithInnerFlags(cli.Command{
	Name:    "contest",
	Usage:   "Submit a written response and supporting documents disputing the claim. The\nfirst call moves the claim from `pending` to `contested`; subsequent calls\nappend supplementary evidence without changing status. The `documents[]` you\nattach are aggregated across rounds in the claim's `contest_documents` field.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "claim-id",
			Required:  true,
			PathParam: "claim_id",
		},
		&requestflag.Flag[string]{
			Name:     "contest-notes",
			Usage:    "Customer's response to the claim. 10–2000 characters.",
			Required: true,
			BodyPath: "contest_notes",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "document",
			Usage:    "Up to 20 supporting documents per submission. `document_id` must be unique within this submission. Documents are aggregated into the claim's `contest_documents` across all submissions.",
			BodyPath: "documents",
		},
	},
	Action:          handleInfringementClaimsContest,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"document": {
		&requestflag.InnerFlag[string]{
			Name:       "document.document-id",
			Usage:      "Id returned by the Telnyx Documents API after you upload the file (upload via `POST /v2/documents`; see https://developers.telnyx.com/api/documents).",
			InnerField: "document_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "document.document-type",
			Usage:      "Type of supporting document. Pick the closest match to what the file actually contains; `other` triggers manual vetting and may slow approval. The matching short_name reference list is at `GET /v2/dir/document_types`.",
			InnerField: "document_type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "document.description",
			InnerField: "description",
		},
	},
})

func handleInfringementClaimsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("claim-id") && len(unusedArgs) > 0 {
		cmd.Set("claim-id", unusedArgs[0])
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
	_, err = client.InfringementClaims.Get(ctx, cmd.Value("claim-id").(string), options...)
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
		Title:          "infringement-claims retrieve",
		Transform:      transform,
	})
}

func handleInfringementClaimsContest(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("claim-id") && len(unusedArgs) > 0 {
		cmd.Set("claim-id", unusedArgs[0])
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

	params := telnyx.InfringementClaimContestParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.InfringementClaims.Contest(
		ctx,
		cmd.Value("claim-id").(string),
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
		Title:          "infringement-claims contest",
		Transform:      transform,
	})
}
