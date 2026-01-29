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

var portingOrdersAdditionalDocumentsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a list of additional documents for a porting order.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "additional-document",
			BodyPath: "additional_documents",
		},
	},
	Action:          handlePortingOrdersAdditionalDocumentsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"additional-document": {
		&requestflag.InnerFlag[string]{
			Name:       "additional-document.document-id",
			Usage:      "The document identification",
			InnerField: "document_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "additional-document.document-type",
			Usage:      "The type of document being created.",
			InnerField: "document_type",
		},
	},
})

var portingOrdersAdditionalDocumentsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Returns a list of additional documents for a porting order.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter (deepObject style). Originally: filter[document_type]",
			QueryPath: "filter",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "page",
			Usage:     "Consolidated page parameter (deepObject style). Originally: page[size], page[number]",
			QueryPath: "page",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "sort",
			Usage:     "Consolidated sort parameter (deepObject style). Originally: sort[value]",
			QueryPath: "sort",
		},
	},
	Action:          handlePortingOrdersAdditionalDocumentsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[[]string]{
			Name:       "filter.document-type",
			Usage:      "Filter additional documents by a list of document types",
			InnerField: "document_type",
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
	"sort": {
		&requestflag.InnerFlag[string]{
			Name:       "sort.value",
			Usage:      "Specifies the sort order for results. If not given, results are sorted by created_at in descending order.",
			InnerField: "value",
		},
	},
})

var portingOrdersAdditionalDocumentsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes an additional document for a porting order.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "additional-document-id",
			Required: true,
		},
	},
	Action:          handlePortingOrdersAdditionalDocumentsDelete,
	HideHelpCommand: true,
}

func handlePortingOrdersAdditionalDocumentsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortingOrderAdditionalDocumentNewParams{}

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
	_, err = client.PortingOrders.AdditionalDocuments.New(
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
	return ShowJSON(os.Stdout, "porting-orders:additional-documents create", obj, format, transform)
}

func handlePortingOrdersAdditionalDocumentsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortingOrderAdditionalDocumentListParams{}

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
		_, err = client.PortingOrders.AdditionalDocuments.List(
			ctx,
			cmd.Value("id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "porting-orders:additional-documents list", obj, format, transform)
	} else {
		iter := client.PortingOrders.AdditionalDocuments.ListAutoPaging(
			ctx,
			cmd.Value("id").(string),
			params,
			options...,
		)
		return ShowJSONIterator(os.Stdout, "porting-orders:additional-documents list", iter, format, transform)
	}
}

func handlePortingOrdersAdditionalDocumentsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("additional-document-id") && len(unusedArgs) > 0 {
		cmd.Set("additional-document-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.PortingOrderAdditionalDocumentDeleteParams{
		ID: cmd.Value("id").(string),
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

	return client.PortingOrders.AdditionalDocuments.Delete(
		ctx,
		cmd.Value("additional-document-id").(string),
		params,
		options...,
	)
}
