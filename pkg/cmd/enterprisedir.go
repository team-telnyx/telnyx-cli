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

var enterprisesDirCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a new DIR under the given enterprise. The DIR starts in `draft` status;\nit must be submitted (`POST .../submit`) and approved by Telnyx before any phone\nnumber can be attached.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "enterprise-id",
			Required:  true,
			PathParam: "enterprise_id",
		},
		&requestflag.Flag[string]{
			Name:     "authorizer-email",
			Usage:    "Contact email of the authorizer. Telnyx may send verification or infringement-notice email here; use a monitored mailbox.",
			Required: true,
			BodyPath: "authorizer_email",
		},
		&requestflag.Flag[string]{
			Name:     "authorizer-name",
			Usage:    "Name of the person at your enterprise who is authorizing this DIR registration. Must be a real individual (used for audit and trademark-claim contests).",
			Required: true,
			BodyPath: "authorizer_name",
		},
		&requestflag.Flag[[]string]{
			Name:     "call-reason",
			Usage:    "1–10 reasons your business calls customers. Validate phrasing against `POST /call_reasons/validate`.",
			Required: true,
			BodyPath: "call_reasons",
		},
		&requestflag.Flag[bool]{
			Name:     "certify-brand-is-accurate",
			Usage:    "Must be `true`.",
			Required: true,
			BodyPath: "certify_brand_is_accurate",
		},
		&requestflag.Flag[bool]{
			Name:     "certify-ip-ownership",
			Usage:    "Must be `true`. Confirms ownership of any logos/trademarks shown.",
			Required: true,
			BodyPath: "certify_ip_ownership",
		},
		&requestflag.Flag[bool]{
			Name:     "certify-no-shaft-content",
			Usage:    "Must be `true`. Confirms this DIR is not used for SHAFT content (Sex, Hate, Alcohol, Firearms, Tobacco) where prohibited.",
			Required: true,
			BodyPath: "certify_no_shaft_content",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			Usage:    "Name shown to call recipients. No emoji; not whitespace-only.",
			Required: true,
			BodyPath: "display_name",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "document",
			Usage:    "Supporting documents. Each `document_id` may appear at most once on a DIR.",
			BodyPath: "documents",
		},
		&requestflag.Flag[string]{
			Name:     "logo-url",
			Usage:    "Publicly accessible HTTPS URL (max 128 chars) to a 256x256 BMP logo (max 1 MB).",
			BodyPath: "logo_url",
		},
		&requestflag.Flag[bool]{
			Name:     "reselling",
			Usage:    "Set to true if your organization places calls on behalf of other enterprises (BPO/reseller).",
			Default:  false,
			BodyPath: "reselling",
		},
	},
	Action:          handleEnterprisesDirCreate,
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

var enterprisesDirList = cli.Command{
	Name:    "list",
	Usage:   "Return the DIRs (Display Identity Records) belonging to a single enterprise.\nPagination is JSON:API style (`page[number]`, `page[size]`, max 250). Supports\n`filter[]` query params: `filter[status]`, `filter[display_name][contains]`,\n`filter[call_reason][contains]`, plus the renewal-window filters\n`filter[expiring_at][gte]` / `filter[expiring_at][lte]` and the convenience\n`filter[expiring_within_days]` (mutually exclusive with the explicit gte/lte\nform). Sortable by `created_at`, `updated_at`, `display_name`, `status`,\n`submitted_at`, `verified_at`, `expiring_at` (prefix `-` for descending; default\n`-created_at`).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "enterprise-id",
			Required:  true,
			PathParam: "enterprise_id",
		},
		&requestflag.Flag[string]{
			Name:      "filter-call-reason-contains",
			Usage:     "Case-insensitive partial match on call reason.",
			QueryPath: "filter[call_reason][contains]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-display-name-contains",
			Usage:     "Case-insensitive partial match on display name.",
			QueryPath: "filter[display_name][contains]",
		},
		&requestflag.Flag[any]{
			Name:      "filter-expiring-at-gte",
			Usage:     "Return only DIRs whose `expiring_at` is at or after this ISO-8601 timestamp.",
			QueryPath: "filter[expiring_at][gte]",
		},
		&requestflag.Flag[any]{
			Name:      "filter-expiring-at-lte",
			Usage:     "Return only DIRs whose `expiring_at` is at or before this ISO-8601 timestamp.",
			QueryPath: "filter[expiring_at][lte]",
		},
		&requestflag.Flag[int64]{
			Name:      "filter-expiring-within-days",
			Usage:     "Convenience: returns DIRs whose `expiring_at` falls within the next N days (1–365). Equivalent to setting `filter[expiring_at][gte]=<now>` + `filter[expiring_at][lte]=<now+N>`. Mutually exclusive with the explicit `[gte]`/`[lte]` filters - combining returns 400.",
			QueryPath: "filter[expiring_within_days]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-status",
			Usage:     "Filter by DIR status.",
			QueryPath: "filter[status]",
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
			Name:      "sort",
			Usage:     "Sort field. Allowed: `created_at`, `updated_at`, `display_name`, `status`, `submitted_at`, `verified_at`, `expiring_at`. Prefix with `-` for descending. Default `-created_at`.",
			Default:   "-created_at",
			QueryPath: "sort",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleEnterprisesDirList,
	HideHelpCommand: true,
}

func handleEnterprisesDirCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EnterpriseDirNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Enterprises.Dir.New(
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
		Title:          "enterprises:dir create",
		Transform:      transform,
	})
}

func handleEnterprisesDirList(ctx context.Context, cmd *cli.Command) error {
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
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := telnyx.EnterpriseDirListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Enterprises.Dir.List(
			ctx,
			cmd.Value("enterprise-id").(string),
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
			Title:          "enterprises:dir list",
			Transform:      transform,
		})
	} else {
		iter := client.Enterprises.Dir.ListAutoPaging(
			ctx,
			cmd.Value("enterprise-id").(string),
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
			Title:          "enterprises:dir list",
			Transform:      transform,
		})
	}
}
