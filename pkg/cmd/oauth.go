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

var oauthRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve details about an OAuth consent token",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "consent-token",
			Required: true,
		},
	},
	Action:          handleOAuthRetrieve,
	HideHelpCommand: true,
}

var oauthGrants = cli.Command{
	Name:    "grants",
	Usage:   "Create an OAuth authorization grant",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[bool]{
			Name:     "allowed",
			Usage:    "Whether the grant is allowed",
			Required: true,
			BodyPath: "allowed",
		},
		&requestflag.Flag[string]{
			Name:     "consent-token",
			Usage:    "Consent token",
			Required: true,
			BodyPath: "consent_token",
		},
	},
	Action:          handleOAuthGrants,
	HideHelpCommand: true,
}

var oauthIntrospect = cli.Command{
	Name:    "introspect",
	Usage:   "Introspect an OAuth access token to check its validity and metadata",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "token",
			Usage:    "The token to introspect",
			Required: true,
			BodyPath: "token",
		},
	},
	Action:          handleOAuthIntrospect,
	HideHelpCommand: true,
}

var oauthRegister = cli.Command{
	Name:    "register",
	Usage:   "Register a new OAuth client dynamically (RFC 7591)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "client-name",
			Usage:    "Human-readable string name of the client to be presented to the end-user",
			BodyPath: "client_name",
		},
		&requestflag.Flag[[]string]{
			Name:     "grant-type",
			Usage:    "Array of OAuth 2.0 grant type strings that the client may use",
			Default:  []string{"authorization_code"},
			BodyPath: "grant_types",
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
			Usage:    "Array of redirection URI strings for use in redirect-based flows",
			BodyPath: "redirect_uris",
		},
		&requestflag.Flag[[]string]{
			Name:     "response-type",
			Usage:    "Array of the OAuth 2.0 response type strings that the client may use",
			Default:  []string{"code"},
			BodyPath: "response_types",
		},
		&requestflag.Flag[string]{
			Name:     "scope",
			Usage:    "Space-separated string of scope values that the client may use",
			BodyPath: "scope",
		},
		&requestflag.Flag[string]{
			Name:     "token-endpoint-auth-method",
			Usage:    "Authentication method for the token endpoint",
			Default:  "client_secret_basic",
			BodyPath: "token_endpoint_auth_method",
		},
		&requestflag.Flag[string]{
			Name:     "tos-uri",
			Usage:    "URL of the client's terms of service",
			BodyPath: "tos_uri",
		},
	},
	Action:          handleOAuthRegister,
	HideHelpCommand: true,
}

var oauthRetrieveAuthorize = cli.Command{
	Name:    "retrieve-authorize",
	Usage:   "OAuth 2.0 authorization endpoint for the authorization code flow",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "client-id",
			Usage:     "OAuth client identifier",
			Required:  true,
			QueryPath: "client_id",
		},
		&requestflag.Flag[string]{
			Name:      "redirect-uri",
			Usage:     "Redirect URI",
			Required:  true,
			QueryPath: "redirect_uri",
		},
		&requestflag.Flag[string]{
			Name:      "response-type",
			Usage:     "OAuth response type",
			Required:  true,
			QueryPath: "response_type",
		},
		&requestflag.Flag[string]{
			Name:      "code-challenge",
			Usage:     "PKCE code challenge",
			QueryPath: "code_challenge",
		},
		&requestflag.Flag[string]{
			Name:      "code-challenge-method",
			Usage:     "PKCE code challenge method",
			QueryPath: "code_challenge_method",
		},
		&requestflag.Flag[string]{
			Name:      "scope",
			Usage:     "Space-separated list of requested scopes",
			QueryPath: "scope",
		},
		&requestflag.Flag[string]{
			Name:      "state",
			Usage:     "State parameter for CSRF protection",
			QueryPath: "state",
		},
	},
	Action:          handleOAuthRetrieveAuthorize,
	HideHelpCommand: true,
}

var oauthRetrieveJwks = cli.Command{
	Name:            "retrieve-jwks",
	Usage:           "Retrieve the JSON Web Key Set for token verification",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleOAuthRetrieveJwks,
	HideHelpCommand: true,
}

var oauthToken = cli.Command{
	Name:    "token",
	Usage:   "Exchange authorization code, client credentials, or refresh token for access\ntoken",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "grant-type",
			Usage:    "OAuth 2.0 grant type",
			Required: true,
			BodyPath: "grant_type",
		},
		&requestflag.Flag[string]{
			Name:     "client-id",
			Usage:    "OAuth client ID (if not using HTTP Basic auth)",
			BodyPath: "client_id",
		},
		&requestflag.Flag[string]{
			Name:     "client-secret",
			Usage:    "OAuth client secret (if not using HTTP Basic auth)",
			BodyPath: "client_secret",
		},
		&requestflag.Flag[string]{
			Name:     "code",
			Usage:    "Authorization code (for authorization_code flow)",
			BodyPath: "code",
		},
		&requestflag.Flag[string]{
			Name:     "code-verifier",
			Usage:    "PKCE code verifier (for authorization_code flow)",
			BodyPath: "code_verifier",
		},
		&requestflag.Flag[string]{
			Name:     "redirect-uri",
			Usage:    "Redirect URI (for authorization_code flow)",
			BodyPath: "redirect_uri",
		},
		&requestflag.Flag[string]{
			Name:     "refresh-token",
			Usage:    "Refresh token (for refresh_token flow)",
			BodyPath: "refresh_token",
		},
		&requestflag.Flag[string]{
			Name:     "scope",
			Usage:    "Space-separated list of requested scopes (for client_credentials)",
			BodyPath: "scope",
		},
	},
	Action:          handleOAuthToken,
	HideHelpCommand: true,
}

func handleOAuthRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("consent-token") && len(unusedArgs) > 0 {
		cmd.Set("consent-token", unusedArgs[0])
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
	_, err = client.OAuth.Get(ctx, cmd.Value("consent-token").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "oauth retrieve", obj, format, transform)
}

func handleOAuthGrants(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.OAuthGrantsParams{}

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
	_, err = client.OAuth.Grants(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "oauth grants", obj, format, transform)
}

func handleOAuthIntrospect(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.OAuthIntrospectParams{}

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
	_, err = client.OAuth.Introspect(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "oauth introspect", obj, format, transform)
}

func handleOAuthRegister(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.OAuthRegisterParams{}

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
	_, err = client.OAuth.Register(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "oauth register", obj, format, transform)
}

func handleOAuthRetrieveAuthorize(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.OAuthGetAuthorizeParams{}

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

	return client.OAuth.GetAuthorize(ctx, params, options...)
}

func handleOAuthRetrieveJwks(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.OAuth.GetJwks(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "oauth retrieve-jwks", obj, format, transform)
}

func handleOAuthToken(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.OAuthTokenParams{}

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
	_, err = client.OAuth.Token(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "oauth token", obj, format, transform)
}
