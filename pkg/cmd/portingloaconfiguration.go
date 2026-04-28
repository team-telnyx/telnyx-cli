// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/team-telnyx/telnyx-cli/internal/apiquery"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var portingLoaConfigurationsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a LOA configuration.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "address",
			Usage:    "The address of the company.",
			Required: true,
			BodyPath: "address",
		},
		&requestflag.Flag[string]{
			Name:     "company-name",
			Usage:    "The name of the company",
			Required: true,
			BodyPath: "company_name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "contact",
			Usage:    "The contact information of the company.",
			Required: true,
			BodyPath: "contact",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "logo",
			Usage:    "The logo of the LOA configuration",
			Required: true,
			BodyPath: "logo",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The name of the LOA configuration",
			Required: true,
			BodyPath: "name",
		},
	},
	Action:          handlePortingLoaConfigurationsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"address": {
		&requestflag.InnerFlag[string]{
			Name:       "address.city",
			Usage:      "The locality of the company",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.country-code",
			Usage:      "The country code of the company",
			InnerField: "country_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.state",
			Usage:      "The administrative area of the company",
			InnerField: "state",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.street-address",
			Usage:      "The street address of the company",
			InnerField: "street_address",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.zip-code",
			Usage:      "The postal code of the company",
			InnerField: "zip_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.extended-address",
			Usage:      "The extended address of the company",
			InnerField: "extended_address",
		},
	},
	"contact": {
		&requestflag.InnerFlag[string]{
			Name:       "contact.email",
			Usage:      "The email address of the contact",
			InnerField: "email",
		},
		&requestflag.InnerFlag[string]{
			Name:       "contact.phone-number",
			Usage:      "The phone number of the contact",
			InnerField: "phone_number",
		},
	},
	"logo": {
		&requestflag.InnerFlag[string]{
			Name:       "logo.document-id",
			Usage:      "The document identification",
			InnerField: "document_id",
		},
	},
})

var portingLoaConfigurationsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a specific LOA configuration.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handlePortingLoaConfigurationsRetrieve,
	HideHelpCommand: true,
}

var portingLoaConfigurationsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update a specific LOA configuration.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "address",
			Usage:    "The address of the company.",
			Required: true,
			BodyPath: "address",
		},
		&requestflag.Flag[string]{
			Name:     "company-name",
			Usage:    "The name of the company",
			Required: true,
			BodyPath: "company_name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "contact",
			Usage:    "The contact information of the company.",
			Required: true,
			BodyPath: "contact",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "logo",
			Usage:    "The logo of the LOA configuration",
			Required: true,
			BodyPath: "logo",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The name of the LOA configuration",
			Required: true,
			BodyPath: "name",
		},
	},
	Action:          handlePortingLoaConfigurationsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"address": {
		&requestflag.InnerFlag[string]{
			Name:       "address.city",
			Usage:      "The locality of the company",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.country-code",
			Usage:      "The country code of the company",
			InnerField: "country_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.state",
			Usage:      "The administrative area of the company",
			InnerField: "state",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.street-address",
			Usage:      "The street address of the company",
			InnerField: "street_address",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.zip-code",
			Usage:      "The postal code of the company",
			InnerField: "zip_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.extended-address",
			Usage:      "The extended address of the company",
			InnerField: "extended_address",
		},
	},
	"contact": {
		&requestflag.InnerFlag[string]{
			Name:       "contact.email",
			Usage:      "The email address of the contact",
			InnerField: "email",
		},
		&requestflag.InnerFlag[string]{
			Name:       "contact.phone-number",
			Usage:      "The phone number of the contact",
			InnerField: "phone_number",
		},
	},
	"logo": {
		&requestflag.InnerFlag[string]{
			Name:       "logo.document-id",
			Usage:      "The document identification",
			InnerField: "document_id",
		},
	},
})

var portingLoaConfigurationsList = cli.Command{
	Name:    "list",
	Usage:   "List the LOA configurations.",
	Suggest: true,
	Flags: []cli.Flag{
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
	Action:          handlePortingLoaConfigurationsList,
	HideHelpCommand: true,
}

var portingLoaConfigurationsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a specific LOA configuration.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handlePortingLoaConfigurationsDelete,
	HideHelpCommand: true,
}

var portingLoaConfigurationsPreview = requestflag.WithInnerFlags(cli.Command{
	Name:    "preview",
	Usage:   "Preview the LOA template that would be generated without need to create LOA\nconfiguration.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "address",
			Usage:    "The address of the company.",
			Required: true,
			BodyPath: "address",
		},
		&requestflag.Flag[string]{
			Name:     "company-name",
			Usage:    "The name of the company",
			Required: true,
			BodyPath: "company_name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "contact",
			Usage:    "The contact information of the company.",
			Required: true,
			BodyPath: "contact",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "logo",
			Usage:    "The logo of the LOA configuration",
			Required: true,
			BodyPath: "logo",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The name of the LOA configuration",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "The file where the response contents will be stored. Use the value '-' to force output to stdout.",
		},
	},
	Action:          handlePortingLoaConfigurationsPreview,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"address": {
		&requestflag.InnerFlag[string]{
			Name:       "address.city",
			Usage:      "The locality of the company",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.country-code",
			Usage:      "The country code of the company",
			InnerField: "country_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.state",
			Usage:      "The administrative area of the company",
			InnerField: "state",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.street-address",
			Usage:      "The street address of the company",
			InnerField: "street_address",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.zip-code",
			Usage:      "The postal code of the company",
			InnerField: "zip_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.extended-address",
			Usage:      "The extended address of the company",
			InnerField: "extended_address",
		},
	},
	"contact": {
		&requestflag.InnerFlag[string]{
			Name:       "contact.email",
			Usage:      "The email address of the contact",
			InnerField: "email",
		},
		&requestflag.InnerFlag[string]{
			Name:       "contact.phone-number",
			Usage:      "The phone number of the contact",
			InnerField: "phone_number",
		},
	},
	"logo": {
		&requestflag.InnerFlag[string]{
			Name:       "logo.document-id",
			Usage:      "The document identification",
			InnerField: "document_id",
		},
	},
})

var portingLoaConfigurationsPreview1 = cli.Command{
	Name:    "preview-1",
	Usage:   "Preview a specific LOA configuration.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "The file where the response contents will be stored. Use the value '-' to force output to stdout.",
		},
	},
	Action:          handlePortingLoaConfigurationsPreview1,
	HideHelpCommand: true,
}

func handlePortingLoaConfigurationsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortingLoaConfigurationNewParams{}

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
	_, err = client.Porting.LoaConfigurations.New(ctx, params, options...)
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
		Title:          "porting:loa-configurations create",
		Transform:      transform,
	})
}

func handlePortingLoaConfigurationsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Porting.LoaConfigurations.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "porting:loa-configurations retrieve",
		Transform:      transform,
	})
}

func handlePortingLoaConfigurationsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortingLoaConfigurationUpdateParams{}

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
	_, err = client.Porting.LoaConfigurations.Update(
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
		Title:          "porting:loa-configurations update",
		Transform:      transform,
	})
}

func handlePortingLoaConfigurationsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortingLoaConfigurationListParams{}

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
		_, err = client.Porting.LoaConfigurations.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "porting:loa-configurations list",
			Transform:      transform,
		})
	} else {
		iter := client.Porting.LoaConfigurations.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "porting:loa-configurations list",
			Transform:      transform,
		})
	}
}

func handlePortingLoaConfigurationsDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.Porting.LoaConfigurations.Delete(ctx, cmd.Value("id").(string), options...)
}

func handlePortingLoaConfigurationsPreview(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortingLoaConfigurationPreviewParams{}

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

	response, err := client.Porting.LoaConfigurations.Preview(ctx, params, options...)
	if err != nil {
		return err
	}
	message, err := writeBinaryResponse(response, os.Stdout, cmd.String("output"))
	if message != "" {
		fmt.Println(message)
	}
	return err
}

func handlePortingLoaConfigurationsPreview0(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortingLoaConfigurationPreview0Params{}

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

	response, err := client.Porting.LoaConfigurations.Preview0(ctx, params, options...)
	if err != nil {
		return err
	}
	message, err := writeBinaryResponse(response, os.Stdout, cmd.String("output"))
	if message != "" {
		fmt.Println(message)
	}
	return err
}

func handlePortingLoaConfigurationsPreview1(ctx context.Context, cmd *cli.Command) error {
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

	response, err := client.Porting.LoaConfigurations.Preview1(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}
	message, err := writeBinaryResponse(response, os.Stdout, cmd.String("output"))
	if message != "" {
		fmt.Println(message)
	}
	return err
}
