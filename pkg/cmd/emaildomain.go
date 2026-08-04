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

var emailDomainsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create an email domain",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "domain",
			Required: true,
			BodyPath: "domain",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "dmarc-policy",
			Usage:    "DMARC policy for a sending domain. Drives the recommended _dmarc.<domain> TXT record. DMARC is advisory and never blocks sending. When omitted or null, the domain uses the advisory default (v=DMARC1; p=none; rua=mailto:dmarc@telnyx.com).",
			BodyPath: "dmarc_policy",
		},
		&requestflag.Flag[bool]{
			Name:     "inbound-enabled",
			Usage:    "Enable inbound routing for this domain",
			Default:  false,
			BodyPath: "inbound_enabled",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "tracking",
			BodyPath: "tracking",
		},
	},
	Action:          handleEmailDomainsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"dmarc-policy": {
		&requestflag.InnerFlag[string]{
			Name:       "dmarc-policy.p",
			Usage:      "Policy applied to messages that fail alignment.",
			InnerField: "p",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "dmarc-policy.pct",
			Usage:      "Percentage of messages the policy applies to. Omitted from the record when 100.",
			InnerField: "pct",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "dmarc-policy.rua",
			Usage:      "URI for aggregate reports. Defaults to the Telnyx address when absent; null omits it.",
			InnerField: "rua",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "dmarc-policy.sp",
			Usage:      "Policy for subdomains. Omitted from the record when null.",
			InnerField: "sp",
		},
	},
	"tracking": {
		&requestflag.InnerFlag[bool]{
			Name:       "tracking.click-tracking",
			Usage:      "Rewrite HTML links through a tracking redirect to record click events.",
			InnerField: "click_tracking",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "tracking.open-tracking",
			Usage:      "Inject a tracking pixel into HTML messages to record open events.",
			InnerField: "open_tracking",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "tracking.unsubscribe-tracking",
			Usage:      "Add RFC 8058 List-Unsubscribe headers with a signed one-click unsubscribe URL. Enabled by default; Gmail/Yahoo bulk-sender rules require one-click unsubscribe support.",
			InnerField: "unsubscribe_tracking",
		},
	},
})

var emailDomainsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Shared (`type: shared`) Telnyx-managed domains are included/readable for every\naccount, in addition to the account's own custom domains.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleEmailDomainsRetrieve,
	HideHelpCommand: true,
}

var emailDomainsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update an email domain",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "dmarc-policy",
			Usage:    "DMARC policy for a sending domain. Drives the recommended _dmarc.<domain> TXT record. DMARC is advisory and never blocks sending. When omitted or null, the domain uses the advisory default (v=DMARC1; p=none; rua=mailto:dmarc@telnyx.com).",
			BodyPath: "dmarc_policy",
		},
		&requestflag.Flag[bool]{
			Name:     "inbound-enabled",
			Usage:    "Enable or disable inbound routing for this domain",
			BodyPath: "inbound_enabled",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "tracking",
			BodyPath: "tracking",
		},
	},
	Action:          handleEmailDomainsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"dmarc-policy": {
		&requestflag.InnerFlag[string]{
			Name:       "dmarc-policy.p",
			Usage:      "Policy applied to messages that fail alignment.",
			InnerField: "p",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "dmarc-policy.pct",
			Usage:      "Percentage of messages the policy applies to. Omitted from the record when 100.",
			InnerField: "pct",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "dmarc-policy.rua",
			Usage:      "URI for aggregate reports. Defaults to the Telnyx address when absent; null omits it.",
			InnerField: "rua",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "dmarc-policy.sp",
			Usage:      "Policy for subdomains. Omitted from the record when null.",
			InnerField: "sp",
		},
	},
	"tracking": {
		&requestflag.InnerFlag[bool]{
			Name:       "tracking.click-tracking",
			Usage:      "Rewrite HTML links through a tracking redirect to record click events.",
			InnerField: "click_tracking",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "tracking.open-tracking",
			Usage:      "Inject a tracking pixel into HTML messages to record open events.",
			InnerField: "open_tracking",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "tracking.unsubscribe-tracking",
			Usage:      "Add RFC 8058 List-Unsubscribe headers with a signed one-click unsubscribe URL. Enabled by default; Gmail/Yahoo bulk-sender rules require one-click unsubscribe support.",
			InnerField: "unsubscribe_tracking",
		},
	},
})

var emailDomainsList = cli.Command{
	Name:    "list",
	Usage:   "Shared (`type: shared`) Telnyx-managed domains are included/readable for every\naccount, in addition to the account's own custom domains.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "filter-domain",
			Usage:     "Partial match on domain name (case-insensitive)",
			QueryPath: "filter[domain]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-profile-id",
			Usage:     "Filter by profile UUID",
			QueryPath: "filter[profile_id]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-status",
			Usage:     `Allowed values: "pending", "verifying", "verified", "failed", "degraded", "suspended".`,
			QueryPath: "filter[status]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-type",
			Usage:     `Allowed values: "custom", "shared", "shared_inbound".`,
			QueryPath: "filter[type]",
		},
		&requestflag.Flag[bool]{
			Name:      "filter-usable-for-inbound",
			QueryPath: "filter[usable_for_inbound]",
		},
		&requestflag.Flag[bool]{
			Name:      "filter-usable-for-sending",
			QueryPath: "filter[usable_for_sending]",
		},
		&requestflag.Flag[string]{
			Name:      "page-after",
			Usage:     "Cursor for records after the provided value (cursor pagination)",
			QueryPath: "page[after]",
		},
		&requestflag.Flag[string]{
			Name:      "page-before",
			Usage:     "Cursor for records before the provided value (cursor pagination)",
			QueryPath: "page[before]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "Page number to return (offset pagination)",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Number of records per page",
			Default:   25,
			QueryPath: "page[size]",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Field to sort by. Prefix with `-` for descending order.",
			QueryPath: "sort",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleEmailDomainsList,
	HideHelpCommand: true,
}

var emailDomainsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete an email domain",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[bool]{
			Name:      "force",
			Usage:     "Required as true when deleting verified domains",
			Default:   false,
			QueryPath: "force",
		},
	},
	Action:          handleEmailDomainsDelete,
	HideHelpCommand: true,
}

var emailDomainsRetrieveDNSRecords = cli.Command{
	Name:    "retrieve-dns-records",
	Usage:   "List DNS records for an email domain",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "domain-id",
			Required:  true,
			PathParam: "domain_id",
		},
	},
	Action:          handleEmailDomainsRetrieveDNSRecords,
	HideHelpCommand: true,
}

var emailDomainsRetrieveHealth = cli.Command{
	Name:    "retrieve-health",
	Usage:   "Returns a summary of domain health including verification status and usability.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleEmailDomainsRetrieveHealth,
	HideHelpCommand: true,
}

var emailDomainsVerify = cli.Command{
	Name:    "verify",
	Usage:   "Verify DNS records for an email domain",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "domain-id",
			Required:  true,
			PathParam: "domain_id",
		},
	},
	Action:          handleEmailDomainsVerify,
	HideHelpCommand: true,
}

func handleEmailDomainsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := telnyx.EmailDomainNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailDomains.New(ctx, params, options...)
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
		Title:          "email-domains create",
		Transform:      transform,
	})
}

func handleEmailDomainsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.EmailDomains.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "email-domains retrieve",
		Transform:      transform,
	})
}

func handleEmailDomainsUpdate(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := telnyx.EmailDomainUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailDomains.Update(
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
		Title:          "email-domains update",
		Transform:      transform,
	})
}

func handleEmailDomainsList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailDomainListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.EmailDomains.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "email-domains list",
			Transform:      transform,
		})
	} else {
		iter := client.EmailDomains.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "email-domains list",
			Transform:      transform,
		})
	}
}

func handleEmailDomainsDelete(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailDomainDeleteParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailDomains.Delete(
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
		Title:          "email-domains delete",
		Transform:      transform,
	})
}

func handleEmailDomainsRetrieveDNSRecords(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("domain-id") && len(unusedArgs) > 0 {
		cmd.Set("domain-id", unusedArgs[0])
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
	_, err = client.EmailDomains.GetDNSRecords(ctx, cmd.Value("domain-id").(string), options...)
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
		Title:          "email-domains retrieve-dns-records",
		Transform:      transform,
	})
}

func handleEmailDomainsRetrieveHealth(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.EmailDomains.GetHealth(ctx, cmd.Value("id").(string), options...)
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
		Title:          "email-domains retrieve-health",
		Transform:      transform,
	})
}

func handleEmailDomainsVerify(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("domain-id") && len(unusedArgs) > 0 {
		cmd.Set("domain-id", unusedArgs[0])
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
	_, err = client.EmailDomains.Verify(ctx, cmd.Value("domain-id").(string), options...)
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
		Title:          "email-domains verify",
		Transform:      transform,
	})
}
