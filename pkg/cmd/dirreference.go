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

var dirReferencesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Partially update one reference, addressed by the DIR id plus the reference's\ntype (business or financial) and slot.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "dir-id",
			Required:  true,
			PathParam: "dir_id",
		},
		&requestflag.Flag[string]{
			Name:      "ref-type",
			Usage:     `Allowed values: "business", "financial".`,
			Required:  true,
			PathParam: "ref_type",
		},
		&requestflag.Flag[int64]{
			Name:      "slot",
			Required:  true,
			PathParam: "slot",
		},
		&requestflag.Flag[string]{
			Name:     "email",
			Usage:    "Reference contact email address.",
			BodyPath: "email",
		},
		&requestflag.Flag[string]{
			Name:     "full-name",
			Usage:    "Full name of the reference contact.",
			BodyPath: "full_name",
		},
		&requestflag.Flag[*string]{
			Name:     "job-title",
			Usage:    "Job title of the reference contact.",
			BodyPath: "job_title",
		},
		&requestflag.Flag[*string]{
			Name:     "organization",
			Usage:    "Organization the reference contact belongs to.",
			BodyPath: "organization",
		},
		&requestflag.Flag[string]{
			Name:     "phone-e164",
			Usage:    "Reference phone number in E.164 format.",
			BodyPath: "phone_e164",
		},
		&requestflag.Flag[*string]{
			Name:     "relationship-to-registrant",
			Usage:    "How the reference contact is related to the registering business.",
			BodyPath: "relationship_to_registrant",
		},
		&requestflag.Flag[string]{
			Name:     "timezone",
			Usage:    "IANA timezone id for the reference.",
			BodyPath: "timezone",
		},
	},
	Action:          handleDirReferencesUpdate,
	HideHelpCommand: true,
}

var dirReferencesList = cli.Command{
	Name:    "list",
	Usage:   "List the business and financial references submitted for a DIR.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "dir-id",
			Required:  true,
			PathParam: "dir_id",
		},
	},
	Action:          handleDirReferencesList,
	HideHelpCommand: true,
}

var dirReferencesSubmit = requestflag.WithInnerFlags(cli.Command{
	Name:    "submit",
	Usage:   "Submit the two business references and one financial reference for a DIR.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "dir-id",
			Required:  true,
			PathParam: "dir_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "business-reference",
			Usage:    "Exactly two business references.",
			Required: true,
			BodyPath: "business_references",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "financial-reference",
			Usage:    "One reference supplied at submit. The reference type is implied by the field that carries it (business_references vs financial_reference).",
			Required: true,
			BodyPath: "financial_reference",
		},
	},
	Action:          handleDirReferencesSubmit,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"business-reference": {
		&requestflag.InnerFlag[string]{
			Name:       "business-reference.email",
			Usage:      "Reference contact email address. Required: the reference is emailed scheduling and dial-in notices.",
			InnerField: "email",
		},
		&requestflag.InnerFlag[string]{
			Name:       "business-reference.full-name",
			Usage:      "Full name of the reference contact.",
			InnerField: "full_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "business-reference.phone-e164",
			Usage:      "Reference phone number in E.164 format, e.g. +14155550123.",
			InnerField: "phone_e164",
		},
		&requestflag.InnerFlag[string]{
			Name:       "business-reference.timezone",
			Usage:      "IANA timezone id for the reference (e.g. America/New_York). Required: calls are only placed within the reference's local 8am-9pm window.",
			InnerField: "timezone",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "business-reference.job-title",
			Usage:      "Job title of the reference contact.",
			InnerField: "job_title",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "business-reference.organization",
			Usage:      "Organization the reference contact belongs to.",
			InnerField: "organization",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "business-reference.relationship-to-registrant",
			Usage:      "How the reference contact is related to the registering business.",
			InnerField: "relationship_to_registrant",
		},
	},
	"financial-reference": {
		&requestflag.InnerFlag[string]{
			Name:       "financial-reference.email",
			Usage:      "Reference contact email address. Required: the reference is emailed scheduling and dial-in notices.",
			InnerField: "email",
		},
		&requestflag.InnerFlag[string]{
			Name:       "financial-reference.full-name",
			Usage:      "Full name of the reference contact.",
			InnerField: "full_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "financial-reference.phone-e164",
			Usage:      "Reference phone number in E.164 format, e.g. +14155550123.",
			InnerField: "phone_e164",
		},
		&requestflag.InnerFlag[string]{
			Name:       "financial-reference.timezone",
			Usage:      "IANA timezone id for the reference (e.g. America/New_York). Required: calls are only placed within the reference's local 8am-9pm window.",
			InnerField: "timezone",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "financial-reference.job-title",
			Usage:      "Job title of the reference contact.",
			InnerField: "job_title",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "financial-reference.organization",
			Usage:      "Organization the reference contact belongs to.",
			InnerField: "organization",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "financial-reference.relationship-to-registrant",
			Usage:      "How the reference contact is related to the registering business.",
			InnerField: "relationship_to_registrant",
		},
	},
})

func handleDirReferencesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("slot") && len(unusedArgs) > 0 {
		cmd.Set("slot", unusedArgs[0])
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

	params := telnyx.DirReferenceUpdateParams{
		DirID:   cmd.Value("dir-id").(string),
		RefType: telnyx.DirReferenceUpdateParamsRefType(cmd.Value("ref-type").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Dir.References.Update(
		ctx,
		cmd.Value("slot").(int64),
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
		Title:          "dir:references update",
		Transform:      transform,
	})
}

func handleDirReferencesList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("dir-id") && len(unusedArgs) > 0 {
		cmd.Set("dir-id", unusedArgs[0])
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
	_, err = client.Dir.References.List(ctx, cmd.Value("dir-id").(string), options...)
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
		Title:          "dir:references list",
		Transform:      transform,
	})
}

func handleDirReferencesSubmit(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("dir-id") && len(unusedArgs) > 0 {
		cmd.Set("dir-id", unusedArgs[0])
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

	params := telnyx.DirReferenceSubmitParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Dir.References.Submit(
		ctx,
		cmd.Value("dir-id").(string),
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
		Title:          "dir:references submit",
		Transform:      transform,
	})
}
