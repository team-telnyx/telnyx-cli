// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

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

var dirUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Edit a DIR. DIRs in `draft`, `rejected`, `unsuccessful`, or `suspended` can be\nedited freely: PATCH is a pure edit, `status` is never changed, and you re-vet\nby calling `POST /v2/dir/{dir_id}/submit` explicitly. A `verified` DIR can also\nbe edited in place: a PATCH that changes any value returns the DIR to `draft`;\nthe currently approved identity keeps displaying, and the edited content goes\nlive only after you re-submit and the DIR is approved again. A PATCH that\nchanges nothing (an empty body or values identical to the current ones) leaves\nthe DIR `verified`, so idempotent retries are safe. Changing only\n`bpo_authorizations` or `webhook_url` is the exception: the DIR stays\n`verified`. Each BPO authorization is reviewed on its own instead. DIRs in any\nother status (`submitted`, `in_review`, `expired`, `infringement_claimed`,\n`permanently_rejected`) cannot be edited.",
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
		&requestflag.Flag[[]map[string]any]{
			Name:     "bpo-authorization",
			Usage:    "Optional. Replace this DIR's authorized BPO (Business Process Outsourcer) accounts with these, each with its signed Letter of Authorization. The supplied list replaces the current one: a BPO left out has its authorization removed, and a new BPO (or a changed Letter of Authorization) is created `pending` admin review. Send an empty list to clear all authorizations; omit the field to leave them unchanged. Editing this list does not re-vet the DIR. Maximum 10.",
			BodyPath: "bpo_authorizations",
		},
		&requestflag.Flag[[]string]{
			Name:     "call-reason",
			Usage:    "1–10 reasons your business calls customers. Validate phrasing against `POST /call_reasons/validate`.",
			BodyPath: "call_reasons",
		},
		&requestflag.Flag[bool]{
			Name:     "certify-brand-is-accurate",
			Usage:    "Certification that the DIR information is accurate. Must be `true` for the DIR to be submitted for vetting.",
			BodyPath: "certify_brand_is_accurate",
		},
		&requestflag.Flag[bool]{
			Name:     "certify-ip-ownership",
			Usage:    "Certification of ownership of any logos/trademarks shown. Must be `true` for the DIR to be submitted for vetting.",
			BodyPath: "certify_ip_ownership",
		},
		&requestflag.Flag[bool]{
			Name:     "certify-no-shaft-content",
			Usage:    "Certification that this DIR is not used for SHAFT content (Sex, Hate, Alcohol, Firearms, Tobacco) where prohibited. Must be `true` for the DIR to be submitted for vetting.",
			BodyPath: "certify_no_shaft_content",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			Usage:    "Name shown to call recipients. 1–35 characters, no emoji, not whitespace-only.",
			BodyPath: "display_name",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "document",
			Usage:    "Additional supporting documents to attach. Append-only: existing documents are never removed or replaced, and an empty or omitted list is a no-op. Each `document_id` may appear at most once on a DIR.",
			BodyPath: "documents",
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
		&requestflag.Flag[*string]{
			Name:     "webhook-url",
			Usage:    "Optional `https://` URL that receives webhook notifications when this DIR's compliance review completes. Send `null` to clear. Changing only this field on a `verified` DIR does not re-vet it. Maximum 2048 characters.",
			BodyPath: "webhook_url",
		},
	},
	Action:          handleDirUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"bpo-authorization": {
		&requestflag.InnerFlag[string]{
			Name:       "bpo-authorization.bpo-enterprise-id",
			Usage:      "Enterprise id of an approved BPO (Business Process Outsourcer) account on your organization to authorize for this DIR.",
			InnerField: "bpo_enterprise_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "bpo-authorization.loa-document-id",
			Usage:      "Id of the signed Letter of Authorization document (uploaded via the Telnyx Documents API) in which the Brand Owner authorizes this BPO.",
			InnerField: "loa_document_id",
		},
	},
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
			Usage:      "An optional note describing this document, for example what it proves.",
			InnerField: "description",
		},
	},
})

var dirList = cli.Command{
	Name:    "list",
	Usage:   "Returns every DIR (Display Identity Record) you own, across all of your\nenterprises, as a single list. Pagination is JSON:API style (`page[number]`,\n`page[size]`, max 250). Supports `filter[]` query params:\n`filter[enterprise_id]`, `filter[status]`, `filter[display_name][contains]`,\n`filter[call_reason][contains]`, plus the renewal-window filters\n`filter[expiring_at][gte]` / `filter[expiring_at][lte]`. Sortable by\n`created_at`, `updated_at`, `display_name`, `status` (prefix `-` for descending;\ndefault `-created_at`).",
	Suggest: true,
	Flags: []cli.Flag{
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
		&requestflag.Flag[string]{
			Name:      "filter-enterprise-id",
			Usage:     "Filter by enterprise ID.",
			QueryPath: "filter[enterprise_id]",
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
		&requestflag.Flag[string]{
			Name:      "filter-status",
			Usage:     "DIR lifecycle status.\n- `draft` - newly created; editable; not yet submitted.\n- `submitted` / `in_review` - Telnyx is reviewing.\n- `verified` - approved; phone numbers may be attached.\n- `rejected` - Telnyx rejected this submission; `rejection_reasons` is populated; customer can edit and resubmit.\n- `unsuccessful` - system-side error during processing; customer can edit and resubmit.\n- `suspended` - temporarily disabled (e.g. by an active infringement claim).\n- `expired` - verification expired; customer must resubmit.\n- `infringement_claimed` - a trademark/impersonation claim is open against this DIR.\n- `permanently_rejected` - terminal; cannot be resubmitted.\n- `delete_requested` - you have requested deletion; the DIR still exists and Telnyx is completing the removal (de-registration and cleanup). A verified DIR keeps serving its branded identity, and keeps billing, until the removal finishes.",
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
			Usage:     "Sort field. Allowed values: `created_at`, `updated_at`, `display_name`, `status`. Prefix with `-` for descending. Default `-created_at`.",
			Default:   "-created_at",
			QueryPath: "sort",
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
	Usage:   "Request deletion of a DIR. This does not remove the DIR on this call: it records\nthe request, moves the DIR to `delete_requested`, and Telnyx completes the\nremoval (de-registration and cleanup) shortly after. A verified DIR keeps\nserving its branded identity, and keeps billing, until the removal is executed.\nFailure modes: `400` if a child phone number is still attached or the DIR is\n`in_review` (wait for the review to finish), `409` if the DIR has an unresolved\ninfringement claim, `404` if the DIR is not yours.",
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

var dirBpoLoa = requestflag.WithInnerFlags(cli.Command{
	Name:    "bpo-loa",
	Usage:   "The Letter of Authorization in which a Brand Owner authorizes an approved BPO\n(Business Process Outsourcer) to place branded calls that display this DIR on\nthe owner's behalf. Both parties are read from the caller's account: the Brand\nOwner is the enterprise that owns the DIR, and the BPO is `bpo_enterprise_id`.\nNo business identity is accepted in the body.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "dir-id",
			Required:  true,
			PathParam: "dir_id",
		},
		&requestflag.Flag[string]{
			Name:     "bpo-enterprise-id",
			Usage:    "The approved BPO enterprise the Brand Owner is authorizing. Must be a BPO account on the caller's organization that has already been approved.",
			Required: true,
			BodyPath: "bpo_enterprise_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "signature",
			BodyPath: "signature",
		},
		&requestflag.Flag[string]{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "The file where the response contents will be stored. Use the value '-' to force output to stdout.",
		},
	},
	Action:          handleDirBpoLoa,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"signature": {
		&requestflag.InnerFlag[string]{
			Name:       "signature.image-base64",
			Usage:      "PNG image, base64-encoded.",
			InnerField: "image_base64",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "signature.signer-name",
			Usage:      "Optional. When absent the rendered PDF falls back to the enterprise contact's legal name.",
			InnerField: "signer_name",
		},
	},
})

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

var dirNewLoa = requestflag.WithInnerFlags(cli.Command{
	Name:    "new-loa",
	Usage:   "Generate a pre-filled Letter of Authorization (LOA) PDF for a DIR. Enterprise\nidentity (legal name, DBA, address, contact, website, tax id) and the DIR\ndisplay name are read server-side; the caller supplies the telephone numbers to\nauthorize, an optional Authorized Agent block, and an optional drawn signature.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "dir-id",
			Required:  true,
			PathParam: "dir_id",
		},
		&requestflag.Flag[[]string]{
			Name:     "phone-number",
			Usage:    "Telephone numbers to authorize on the DIR, in `+E164` format (`+` followed by 10-15 digits). Max 15 per request.",
			Required: true,
			BodyPath: "phone_numbers",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "agent",
			Usage:    "Third-party reseller / partner managing the enterprise's phone numbers on its behalf, such as a call center (BPO) acting for a client. Omit when the enterprise manages its own numbers.",
			BodyPath: "agent",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "signature",
			BodyPath: "signature",
		},
		&requestflag.Flag[string]{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "The file where the response contents will be stored. Use the value '-' to force output to stdout.",
		},
	},
	Action:          handleDirNewLoa,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"agent": {
		&requestflag.InnerFlag[string]{
			Name:       "agent.administrative-area",
			Usage:      "The state or province of the partner's address, as its code, for example IL or ON.",
			InnerField: "administrative_area",
		},
		&requestflag.InnerFlag[string]{
			Name:       "agent.city",
			Usage:      "The city of the partner's address.",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "agent.contact-email",
			Usage:      "The email address of the contact person at the partner.",
			InnerField: "contact_email",
		},
		&requestflag.InnerFlag[string]{
			Name:       "agent.contact-name",
			Usage:      "The name of a contact person at the partner.",
			InnerField: "contact_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "agent.contact-phone",
			Usage:      "The phone number of the contact person at the partner, in E.164 format, for example +13125550000.",
			InnerField: "contact_phone",
		},
		&requestflag.InnerFlag[string]{
			Name:       "agent.contact-title",
			Usage:      "The job title of the contact person at the partner.",
			InnerField: "contact_title",
		},
		&requestflag.InnerFlag[string]{
			Name:       "agent.country",
			Usage:      "The two-letter country code of the partner's address, for example US.",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "agent.legal-name",
			Usage:      "The legal name of the third-party partner or reseller managing these numbers on your behalf.",
			InnerField: "legal_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "agent.postal-code",
			Usage:      "The postal or ZIP code of the partner's address.",
			InnerField: "postal_code",
		},
		&requestflag.InnerFlag[string]{
			Name:       "agent.street-address",
			Usage:      "The street address of the partner, including the building number and street name.",
			InnerField: "street_address",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "agent.dba",
			Usage:      "The trade name (Doing Business As) the partner operates under, if different from its legal name. Leave blank if it does not apply.",
			InnerField: "dba",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "agent.extended-address",
			Usage:      "An optional second address line for the partner, such as a suite, unit, or floor. Leave blank if it does not apply.",
			InnerField: "extended_address",
		},
	},
	"signature": {
		&requestflag.InnerFlag[string]{
			Name:       "signature.image-base64",
			Usage:      "PNG image, base64-encoded.",
			InnerField: "image_base64",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "signature.signer-name",
			Usage:      "Optional. When absent the rendered PDF falls back to the enterprise contact's legal name.",
			InnerField: "signer_name",
		},
	},
})

var dirRetrieveBpoAuthorizations = cli.Command{
	Name:    "retrieve-bpo-authorizations",
	Usage:   "List the BPO (Business Process Outsourcer) accounts a Brand Owner has authorized\non this DIR, together with the review state of each authorization.",
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
	},
	Action:          handleDirRetrieveBpoAuthorizations,
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
			Usage:    "Check to certify that the brand no longer infringes anyone else's trademark or intellectual property.",
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
			Usage:    "The business name shown to call recipients, 1 to 35 characters, no emoji, not blank.",
			BodyPath: "display_name",
		},
		&requestflag.Flag[any]{
			Name:     "document",
			Usage:    "Append-only supporting documents to attach while resolving the claim (e.g. authorization or licensing proof).",
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
			Usage:                 "An optional note describing this document, for example what it proves.",
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Dir.Delete(ctx, cmd.Value("dir-id").(string), options...)
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
		Title:          "dir delete",
		Transform:      transform,
	})
}

func handleDirBpoLoa(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.DirBpoLoaParams{}

	response, err := client.Dir.BpoLoa(
		ctx,
		cmd.Value("dir-id").(string),
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

func handleDirNewLoa(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.DirNewLoaParams{}

	response, err := client.Dir.NewLoa(
		ctx,
		cmd.Value("dir-id").(string),
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

func handleDirRetrieveBpoAuthorizations(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.DirGetBpoAuthorizationsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Dir.GetBpoAuthorizations(
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
		Title:          "dir retrieve-bpo-authorizations",
		Transform:      transform,
	})
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
