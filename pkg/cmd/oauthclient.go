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

var oauthClientsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new OAuth client",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "allowed-grant-type",
			Usage:    "List of allowed OAuth grant types",
			Required: true,
			BodyPath: "allowed_grant_types",
		},
		&requestflag.Flag[[]string]{
			Name:     "allowed-scope",
			Usage:    "List of allowed OAuth scopes",
			Required: true,
			BodyPath: "allowed_scopes",
		},
		&requestflag.Flag[string]{
			Name:     "client-type",
			Usage:    "OAuth client type",
			Required: true,
			BodyPath: "client_type",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The name of the OAuth client",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "logo-uri",
			Usage:    "URL of the client logo",
			BodyPath: "logo_uri",
		},
		&requestflag.Flag[string]{
			Name:     "policy-uri",
			Usage:    "URL of the client's privacy policy",
			BodyPath: "policy_uri",
		},
		&requestflag.Flag[[]string]{
			Name:     "redirect-uris",
			Usage:    "List of redirect URIs (required for authorization_code flow)",
			Default:  []string{},
			BodyPath: "redirect_uris",
		},
		&requestflag.Flag[bool]{
			Name:     "require-pkce",
			Usage:    "Whether PKCE (Proof Key for Code Exchange) is required for this client",
			Default:  false,
			BodyPath: "require_pkce",
		},
		&requestflag.Flag[string]{
			Name:     "tos-uri",
			Usage:    "URL of the client's terms of service",
			BodyPath: "tos_uri",
		},
	},
	Action:          handleOAuthClientsCreate,
	HideHelpCommand: true,
}

var oauthClientsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a single OAuth client by ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleOAuthClientsRetrieve,
	HideHelpCommand: true,
}

var oauthClientsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update an existing OAuth client",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[[]string]{
			Name:     "allowed-grant-type",
			Usage:    "List of allowed OAuth grant types",
			BodyPath: "allowed_grant_types",
		},
		&requestflag.Flag[[]string]{
			Name:     "allowed-scope",
			Usage:    "List of allowed OAuth scopes",
			BodyPath: "allowed_scopes",
		},
		&requestflag.Flag[string]{
			Name:     "logo-uri",
			Usage:    "URL of the client logo",
			BodyPath: "logo_uri",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The name of the OAuth client",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "policy-uri",
			Usage:    "URL of the client's privacy policy",
			BodyPath: "policy_uri",
		},
		&requestflag.Flag[[]string]{
			Name:     "redirect-uris",
			Usage:    "List of redirect URIs",
			BodyPath: "redirect_uris",
		},
		&requestflag.Flag[bool]{
			Name:     "require-pkce",
			Usage:    "Whether PKCE (Proof Key for Code Exchange) is required for this client",
			BodyPath: "require_pkce",
		},
		&requestflag.Flag[string]{
			Name:     "tos-uri",
			Usage:    "URL of the client's terms of service",
			BodyPath: "tos_uri",
		},
	},
	Action:          handleOAuthClientsUpdate,
	HideHelpCommand: true,
}

var oauthClientsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve a paginated list of OAuth clients for the authenticated user",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "filter-allowed-grant-types-contains",
			Usage:     "Filter by allowed grant type",
			QueryPath: "filter[allowed_grant_types][contains]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-client-id",
			Usage:     "Filter by client ID",
			QueryPath: "filter[client_id]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-client-type",
			Usage:     "Filter by client type",
			QueryPath: "filter[client_type]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-name",
			Usage:     "Filter by exact client name",
			QueryPath: "filter[name]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-name-contains",
			Usage:     "Filter by client name containing text",
			QueryPath: "filter[name][contains]",
		},
		&requestflag.Flag[bool]{
			Name:      "filter-verified",
			Usage:     "Filter by verification status",
			QueryPath: "filter[verified]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "Page number",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Number of results per page",
			Default:   20,
			QueryPath: "page[size]",
		},
	},
	Action:          handleOAuthClientsList,
	HideHelpCommand: true,
}

var oauthClientsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete an OAuth client",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleOAuthClientsDelete,
	HideHelpCommand: true,
}

func handleOAuthClientsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.OAuthClientNewParams{}

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
	_, err = client.OAuthClients.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "oauth-clients create", obj, format, transform)
}

func handleOAuthClientsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.OAuthClients.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "oauth-clients retrieve", obj, format, transform)
}

func handleOAuthClientsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.OAuthClientUpdateParams{}

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
	_, err = client.OAuthClients.Update(
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
	return ShowJSON(os.Stdout, "oauth-clients update", obj, format, transform)
}

func handleOAuthClientsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.OAuthClientListParams{}

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
		_, err = client.OAuthClients.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "oauth-clients list", obj, format, transform)
	} else {
		iter := client.OAuthClients.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "oauth-clients list", iter, format, transform)
	}
}

func handleOAuthClientsDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.OAuthClients.Delete(ctx, cmd.Value("id").(string), options...)
}
