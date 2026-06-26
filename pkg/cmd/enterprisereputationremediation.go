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

var enterprisesReputationRemediationCreate = cli.Command{
	Name:    "create",
	Usage:   "Submit a batch of phone numbers belonging to this enterprise for reputation\nremediation. The request is accepted asynchronously: this endpoint returns `202`\nwith the persisted request id, then the request transitions through processing\nstates until completion. Use the GET endpoints to poll status and per-number\nresults.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "enterprise-id",
			Required:  true,
			PathParam: "enterprise_id",
		},
		&requestflag.Flag[string]{
			Name:     "call-purpose",
			Usage:    "How the numbers are used (free text).",
			Required: true,
			BodyPath: "call_purpose",
		},
		&requestflag.Flag[[]string]{
			Name:     "phone-number",
			Usage:    "Phone numbers in E.164 format. Each must belong to this enterprise. Maximum 2,000 per request.",
			Required: true,
			BodyPath: "phone_numbers",
		},
		&requestflag.Flag[string]{
			Name:     "contact-email",
			Usage:    "Optional contact email for this remediation request.",
			BodyPath: "contact_email",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "Optional https:// URL for status notifications.",
			BodyPath: "webhook_url",
		},
	},
	Action:          handleEnterprisesReputationRemediationCreate,
	HideHelpCommand: true,
}

var enterprisesReputationRemediationRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve the full detail of a remediation request, including current status,\nper-number results (once available), and submission metadata.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "enterprise-id",
			Required:  true,
			PathParam: "enterprise_id",
		},
		&requestflag.Flag[string]{
			Name:      "remediation-id",
			Required:  true,
			PathParam: "remediation_id",
		},
	},
	Action:          handleEnterprisesReputationRemediationRetrieve,
	HideHelpCommand: true,
}

var enterprisesReputationRemediationList = cli.Command{
	Name:    "list",
	Usage:   "Paginated list of remediation requests for this enterprise. List items omit\nper-number results and webhook URLs to keep the response small; call GET by id\nfor full detail. Supports JSON:API pagination and optional filters on status and\ncreated-at range.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "enterprise-id",
			Required:  true,
			PathParam: "enterprise_id",
		},
		&requestflag.Flag[any]{
			Name:      "filter-created-at-gte",
			Usage:     "Only requests created on or after this timestamp (ISO 8601).",
			QueryPath: "filter[created_at][gte]",
		},
		&requestflag.Flag[any]{
			Name:      "filter-created-at-lte",
			Usage:     "Only requests created on or before this timestamp (ISO 8601).",
			QueryPath: "filter[created_at][lte]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-status",
			Usage:     "Customer-facing status of a remediation request.",
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
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleEnterprisesReputationRemediationList,
	HideHelpCommand: true,
}

func handleEnterprisesReputationRemediationCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EnterpriseReputationRemediationNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Enterprises.Reputation.Remediation.New(
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
		Title:          "enterprises:reputation:remediation create",
		Transform:      transform,
	})
}

func handleEnterprisesReputationRemediationRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("remediation-id") && len(unusedArgs) > 0 {
		cmd.Set("remediation-id", unusedArgs[0])
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

	params := telnyx.EnterpriseReputationRemediationGetParams{
		EnterpriseID: cmd.Value("enterprise-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Enterprises.Reputation.Remediation.Get(
		ctx,
		cmd.Value("remediation-id").(string),
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
		Title:          "enterprises:reputation:remediation retrieve",
		Transform:      transform,
	})
}

func handleEnterprisesReputationRemediationList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EnterpriseReputationRemediationListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Enterprises.Reputation.Remediation.List(
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
			Title:          "enterprises:reputation:remediation list",
			Transform:      transform,
		})
	} else {
		iter := client.Enterprises.Reputation.Remediation.ListAutoPaging(
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
			Title:          "enterprises:reputation:remediation list",
			Transform:      transform,
		})
	}
}
