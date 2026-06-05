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

var dirRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns a single DIR by id. The enterprise is resolved server-side from the DIR\nid. Returns `404` if the DIR does not exist or is not yours.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "dir-id",
			Required:  true,
			PathParam: "dir_id",
		},
	},
	Action:          handleDirRetrieve,
	HideHelpCommand: true,
}

var dirUpdate = cli.Command{
	Name:    "update",
	Usage:   "Edit a DIR. Only DIRs in `draft`, `rejected`, `unsuccessful`, or `suspended` are\neditable. PATCH is a pure edit — `status` is never changed by this endpoint. To\nre-vet after editing, call `POST /v2/dir/{dir_id}/submit` explicitly.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "dir-id",
			Required:  true,
			PathParam: "dir_id",
		},
		&requestflag.Flag[string]{
			Name:     "authorizer-email",
			Usage:    "Contact email of the authorizer. Telnyx may send verification or infringement notices here.",
			BodyPath: "authorizer_email",
		},
		&requestflag.Flag[string]{
			Name:     "authorizer-name",
			Usage:    "Name of the person at your enterprise authorizing this DIR. Must be a real individual.",
			BodyPath: "authorizer_name",
		},
		&requestflag.Flag[[]string]{
			Name:     "call-reason",
			Usage:    "1–10 reasons your business calls customers. Validate phrasing against `POST /call_reasons/validate`.",
			BodyPath: "call_reasons",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			Usage:    "Name shown to call recipients. 1–35 characters, no emoji, not whitespace-only.",
			BodyPath: "display_name",
		},
		&requestflag.Flag[string]{
			Name:     "logo-url",
			Usage:    "Publicly accessible HTTPS URL (max 128 chars) to a 256x256 BMP logo (max 1 MB).",
			BodyPath: "logo_url",
		},
		&requestflag.Flag[bool]{
			Name:     "reselling",
			Usage:    "Set to true if your organization places calls on behalf of other enterprises (BPO/reseller). Updating this triggers re-vetting on next submit.",
			BodyPath: "reselling",
		},
	},
	Action:          handleDirUpdate,
	HideHelpCommand: true,
}

var dirList = cli.Command{
	Name:    "list",
	Usage:   "Convenience endpoint that returns every DIR you own without scoping to a\nspecific enterprise. Equivalent to calling\n`GET /v2/enterprises/{enterprise_id}/dir` for each enterprise and concatenating\nthe results, but server-side and paginated as a single list.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "enterprise-id",
			Usage:     "Restrict results to a single enterprise.",
			QueryPath: "enterprise_id",
		},
		&requestflag.Flag[any]{
			Name:      "filter-expiring-at-gte",
			Usage:     "Return only DIRs whose `expiring_at` is at or after this ISO-8601 timestamp. Pairs with the `[lte]` variant to build renewal-window dashboards.",
			QueryPath: "filter[expiring_at][gte]",
		},
		&requestflag.Flag[any]{
			Name:      "filter-expiring-at-lte",
			Usage:     "Return only DIRs whose `expiring_at` is at or before this ISO-8601 timestamp.",
			QueryPath: "filter[expiring_at][lte]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "1-based page number. Out-of-range values return an empty page with correct meta.",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Items per page. Maximum 250; values above are clamped to 250.",
			Default:   20,
			QueryPath: "page[size]",
		},
		&requestflag.Flag[string]{
			Name:      "search",
			Usage:     "Case-insensitive partial match on `display_name` or call reason.",
			QueryPath: "search",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Sort field. Allowed values: `created_at`, `updated_at`, `display_name`, `status`. Prefix with `-` for descending. Default `-created_at`.",
			Default:   "-created_at",
			QueryPath: "sort",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter by DIR status.",
			QueryPath: "status",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleDirList,
	HideHelpCommand: true,
}

var dirDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a DIR. Failure modes: `400` if a child phone number is in a non-deletable\nstatus, `409` if the DIR has an unresolved infringement claim, `404` if the DIR\nis not yours.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "dir-id",
			Required:  true,
			PathParam: "dir_id",
		},
	},
	Action:          handleDirDelete,
	HideHelpCommand: true,
}

var dirListDocumentTypes = cli.Command{
	Name:            "list-document-types",
	Usage:           "Reference list of `document_type` values accepted by\n`DirCreateRequest.documents[].document_type` and the infringement-contest\nendpoint. Each entry has a stable `short_name` (used in API calls) and a\ncustomer-facing description.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleDirListDocumentTypes,
	HideHelpCommand: true,
}

var dirListInfringementClaims = cli.Command{
	Name:    "list-infringement-claims",
	Usage:   "Return the trademark or copyright claims filed against this DIR. Each claim's\n`status` is `pending` (newly filed; DIR auto-suspended), `contested` (you have\nsubmitted contest evidence; awaiting resolution), or `resolved` (final).\nResolution outcomes: `upheld` (claim accepted; DIR stays\nsuspended/permanently_rejected), `rejected` (claim dismissed; DIR restored to\n`verified`), `modified` (partial outcome).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "dir-id",
			Required:  true,
			PathParam: "dir_id",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "1-based page number. Out-of-range values return an empty page with correct meta.",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Items per page. Maximum 250; values above are clamped to 250.",
			Default:   20,
			QueryPath: "page[size]",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleDirListInfringementClaims,
	HideHelpCommand: true,
}

var dirSubmit = cli.Command{
	Name:    "submit",
	Usage:   "Submit a DIR for vetting. Sends the DIR back through the vetting cycle from any\nnon-terminal status. When re-submitting from `suspended` or `expired`, the DIR's\nprevious Branded Calling registration is torn down transactionally and its phone\nnumbers flip back to `submitted`. When re-submitting from `verified`, the\nexisting registration stays live throughout the new vetting cycle.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "dir-id",
			Required:  true,
			PathParam: "dir_id",
		},
	},
	Action:          handleDirSubmit,
	HideHelpCommand: true,
}

var dirUpdateInfringement = requestflag.WithInnerFlags(cli.Command{
	Name:    "update-infringement",
	Usage:   "Push a fix for a DIR that is `suspended` with an open infringement claim back\ninto vetting. `POST /dir/{dir_id}/submit` is blocked while a claim is open, so\nthis is the customer-callable path to update the DIR's content and re-certify\nbefore Telnyx adjudicates the claim. All four certification booleans must be\n`true`. Optional content fields (`display_name`, `logo_url`, `call_reasons`,\n`documents`) update the DIR; documents are append-only.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "dir-id",
			Required:  true,
			PathParam: "dir_id",
		},
		&requestflag.Flag[bool]{
			Name:     "certify-brand-is-accurate",
			Usage:    "Must be `true`.",
			Required: true,
			BodyPath: "certify_brand_is_accurate",
		},
		&requestflag.Flag[bool]{
			Name:     "certify-ip-ownership",
			Usage:    "Must be `true`.",
			Required: true,
			BodyPath: "certify_ip_ownership",
		},
		&requestflag.Flag[bool]{
			Name:     "certify-no-infringement",
			Usage:    "Must be `true`.",
			Required: true,
			BodyPath: "certify_no_infringement",
		},
		&requestflag.Flag[bool]{
			Name:     "certify-no-shaft-content",
			Usage:    "Must be `true`.",
			Required: true,
			BodyPath: "certify_no_shaft_content",
		},
		&requestflag.Flag[string]{
			Name:     "infringement-resolution-notes",
			Usage:    "Explanation of how the infringement concern was addressed.",
			Required: true,
			BodyPath: "infringement_resolution_notes",
		},
		&requestflag.Flag[any]{
			Name:     "call-reason",
			BodyPath: "call_reasons",
		},
		&requestflag.Flag[*string]{
			Name:     "display-name",
			BodyPath: "display_name",
		},
		&requestflag.Flag[any]{
			Name:     "document",
			Usage:    "Append-only supporting documents.",
			BodyPath: "documents",
		},
		&requestflag.Flag[*string]{
			Name:     "logo-url",
			Usage:    "Publicly accessible HTTPS URL (max 128 chars) to a 256x256 BMP logo (max 1 MB).",
			BodyPath: "logo_url",
		},
	},
	Action:          handleDirUpdateInfringement,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"document": {
		&requestflag.InnerFlag[string]{
			Name:                  "document.document-id",
			Usage:                 "Id returned by the Telnyx Documents API after you upload the file (upload via `POST /v2/documents`; see https://developers.telnyx.com/api/documents).",
			InnerField:            "document_id",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[string]{
			Name:                  "document.document-type",
			Usage:                 "Type of supporting document. Pick the closest match to what the file actually contains; `other` triggers manual vetting and may slow approval. The matching short_name reference list is at `GET /v2/dir/document_types`.",
			InnerField:            "document_type",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[string]{
			Name:                  "document.description",
			InnerField:            "description",
			OuterIsArrayOfObjects: true,
		},
	},
})

func handleDirRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Dir.Get(ctx, cmd.Value("dir-id").(string), options...)
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
		Title:          "dir retrieve",
		Transform:      transform,
	})
}

func handleDirUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.DirUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Dir.Update(
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
		Title:          "dir update",
		Transform:      transform,
	})
}

func handleDirList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.DirListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Dir.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "dir list",
			Transform:      transform,
		})
	} else {
		iter := client.Dir.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "dir list",
			Transform:      transform,
		})
	}
}

func handleDirDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.Dir.Delete(ctx, cmd.Value("dir-id").(string), options...)
}

func handleDirListDocumentTypes(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Dir.ListDocumentTypes(ctx, options...)
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
		Title:          "dir list-document-types",
		Transform:      transform,
	})
}

func handleDirListInfringementClaims(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.DirListInfringementClaimsParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Dir.ListInfringementClaims(
			ctx,
			cmd.Value("dir-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "dir list-infringement-claims",
			Transform:      transform,
		})
	} else {
		iter := client.Dir.ListInfringementClaimsAutoPaging(
			ctx,
			cmd.Value("dir-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "dir list-infringement-claims",
			Transform:      transform,
		})
	}
}

func handleDirSubmit(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Dir.Submit(ctx, cmd.Value("dir-id").(string), options...)
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
		Title:          "dir submit",
		Transform:      transform,
	})
}

func handleDirUpdateInfringement(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.DirUpdateInfringementParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Dir.UpdateInfringement(
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
		Title:          "dir update-infringement",
		Transform:      transform,
	})
}
