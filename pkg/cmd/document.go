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

var documentsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a document.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleDocumentsRetrieve,
	HideHelpCommand: true,
}

var documentsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a document.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "document-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "customer-reference",
			Usage:    "Optional reference string for customer tracking.",
			BodyPath: "customer_reference",
		},
		&requestflag.Flag[string]{
			Name:     "filename",
			Usage:    "The filename of the document.",
			BodyPath: "filename",
		},
	},
	Action:          handleDocumentsUpdate,
	HideHelpCommand: true,
}

var documentsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List all documents ordered by created_at descending.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:      "filter",
			Usage:     "Consolidated filter parameter for documents (deepObject style). Originally: filter[filename][contains], filter[customer_reference][eq], filter[customer_reference][in][], filter[created_at][gt], filter[created_at][lt]",
			QueryPath: "filter",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "page",
			Usage:     "Consolidated page parameter (deepObject style). Originally: page[size], page[number]",
			QueryPath: "page",
		},
		&requestflag.Flag[[]string]{
			Name:      "sort",
			Usage:     "Consolidated sort parameter for documents (deepObject style). Originally: sort[]",
			QueryPath: "sort",
		},
	},
	Action:          handleDocumentsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.created-at",
			InnerField: "created_at",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.customer-reference",
			InnerField: "customer_reference",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filter.filename",
			InnerField: "filename",
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

var documentsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a document.<br /><br />A document can only be deleted if it's not linked\nto a service. If it is linked to a service, it must be unlinked prior to\ndeleting.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleDocumentsDelete,
	HideHelpCommand: true,
}

var documentsGenerateDownloadLink = cli.Command{
	Name:    "generate-download-link",
	Usage:   "Generates a temporary pre-signed URL that can be used to download the document\ndirectly from the storage backend without authentication.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleDocumentsGenerateDownloadLink,
	HideHelpCommand: true,
}

var documentsUpload = requestflag.WithInnerFlags(cli.Command{
	Name:    "upload",
	Usage:   "Upload a document.<br /><br />Uploaded files must be linked to a service within\n30 minutes or they will be automatically deleted.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "document",
			Required: true,
			BodyRoot: true,
		},
	},
	Action:          handleDocumentsUpload,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"document": {
		&requestflag.InnerFlag[string]{
			Name:       "document.customer-reference",
			Usage:      "A customer reference string for customer look ups.",
			InnerField: "customer_reference",
		},
		&requestflag.InnerFlag[string]{
			Name:       "document.file",
			Usage:      "Alternatively, instead of the URL you can provide the Base64 encoded contents of the file you are uploading.",
			InnerField: "file",
		},
		&requestflag.InnerFlag[string]{
			Name:       "document.filename",
			Usage:      "The filename of the document.",
			InnerField: "filename",
		},
		&requestflag.InnerFlag[string]{
			Name:       "document.url",
			Usage:      "If the file is already hosted publicly, you can provide a URL and have the documents service fetch it for you.",
			InnerField: "url",
		},
	},
})

var documentsUploadJson = requestflag.WithInnerFlags(cli.Command{
	Name:    "upload-json",
	Usage:   "Upload a document.<br /><br />Uploaded files must be linked to a service within\n30 minutes or they will be automatically deleted.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "document",
			Required: true,
			BodyRoot: true,
		},
	},
	Action:          handleDocumentsUploadJson,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"document": {
		&requestflag.InnerFlag[string]{
			Name:       "document.customer-reference",
			Usage:      "A customer reference string for customer look ups.",
			InnerField: "customer_reference",
		},
		&requestflag.InnerFlag[string]{
			Name:       "document.file",
			Usage:      "Alternatively, instead of the URL you can provide the Base64 encoded contents of the file you are uploading.",
			InnerField: "file",
		},
		&requestflag.InnerFlag[string]{
			Name:       "document.filename",
			Usage:      "The filename of the document.",
			InnerField: "filename",
		},
		&requestflag.InnerFlag[string]{
			Name:       "document.url",
			Usage:      "If the file is already hosted publicly, you can provide a URL and have the documents service fetch it for you.",
			InnerField: "url",
		},
	},
})

func handleDocumentsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Documents.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "documents retrieve", obj, format, transform)
}

func handleDocumentsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("document-id") && len(unusedArgs) > 0 {
		cmd.Set("document-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.DocumentUpdateParams{}

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
	_, err = client.Documents.Update(
		ctx,
		cmd.Value("document-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "documents update", obj, format, transform)
}

func handleDocumentsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.DocumentListParams{}

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
		_, err = client.Documents.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "documents list", obj, format, transform)
	} else {
		iter := client.Documents.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "documents list", iter, format, transform)
	}
}

func handleDocumentsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Documents.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "documents delete", obj, format, transform)
}

func handleDocumentsGenerateDownloadLink(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Documents.GenerateDownloadLink(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "documents generate-download-link", obj, format, transform)
}

func handleDocumentsUpload(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.DocumentUploadParams{}

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
	_, err = client.Documents.Upload(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "documents upload", obj, format, transform)
}

func handleDocumentsUploadJson(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.DocumentUploadJsonParams{}

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
	_, err = client.Documents.UploadJson(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "documents upload-json", obj, format, transform)
}
