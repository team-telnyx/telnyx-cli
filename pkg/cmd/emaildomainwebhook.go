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

var emailDomainsWebhooksCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a webhook endpoint subscribed to a specific allowlist of event types.\nBoth `email.*` events (published by email-api) and `email_domain.*` events\n(published by this service) flow through the same webhooks.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "domain-id",
			Required:  true,
			PathParam: "domain_id",
		},
		&requestflag.Flag[[]string]{
			Name:     "event",
			Usage:    "At least one event type is required.",
			Required: true,
			BodyPath: "events",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "HTTPS endpoint to deliver subscribed events to.",
			Required: true,
			BodyPath: "url",
		},
	},
	Action:          handleEmailDomainsWebhooksCreate,
	HideHelpCommand: true,
}

var emailDomainsWebhooksRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a webhook",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "domain-id",
			Required:  true,
			PathParam: "domain_id",
		},
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleEmailDomainsWebhooksRetrieve,
	HideHelpCommand: true,
}

var emailDomainsWebhooksUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a webhook's URL and/or event subscription. A webhook is bound to its\ndomain — `domain_id` is not mutable.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "domain-id",
			Required:  true,
			PathParam: "domain_id",
		},
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[[]string]{
			Name:     "event",
			BodyPath: "events",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			BodyPath: "url",
		},
	},
	Action:          handleEmailDomainsWebhooksUpdate,
	HideHelpCommand: true,
}

var emailDomainsWebhooksList = cli.Command{
	Name:    "list",
	Usage:   "List webhooks for an email domain",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "domain-id",
			Required:  true,
			PathParam: "domain_id",
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
	Action:          handleEmailDomainsWebhooksList,
	HideHelpCommand: true,
}

var emailDomainsWebhooksDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a webhook",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "domain-id",
			Required:  true,
			PathParam: "domain_id",
		},
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleEmailDomainsWebhooksDelete,
	HideHelpCommand: true,
}

func handleEmailDomainsWebhooksCreate(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := telnyx.EmailDomainWebhookNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailDomains.Webhooks.New(
		ctx,
		cmd.Value("domain-id").(string),
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
		Title:          "email-domains:webhooks create",
		Transform:      transform,
	})
}

func handleEmailDomainsWebhooksRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailDomainWebhookGetParams{
		DomainID: cmd.Value("domain-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailDomains.Webhooks.Get(
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
		Title:          "email-domains:webhooks retrieve",
		Transform:      transform,
	})
}

func handleEmailDomainsWebhooksUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailDomainWebhookUpdateParams{
		DomainID: cmd.Value("domain-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailDomains.Webhooks.Update(
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
		Title:          "email-domains:webhooks update",
		Transform:      transform,
	})
}

func handleEmailDomainsWebhooksList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailDomainWebhookListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.EmailDomains.Webhooks.List(
			ctx,
			cmd.Value("domain-id").(string),
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
			Title:          "email-domains:webhooks list",
			Transform:      transform,
		})
	} else {
		iter := client.EmailDomains.Webhooks.ListAutoPaging(
			ctx,
			cmd.Value("domain-id").(string),
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
			Title:          "email-domains:webhooks list",
			Transform:      transform,
		})
	}
}

func handleEmailDomainsWebhooksDelete(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.EmailDomainWebhookDeleteParams{
		DomainID: cmd.Value("domain-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EmailDomains.Webhooks.Delete(
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
		Title:          "email-domains:webhooks delete",
		Transform:      transform,
	})
}
