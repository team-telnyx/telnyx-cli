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

var authenticationProvidersCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates an authentication provider.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The name associated with the authentication provider.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "settings",
			Usage:    "The settings associated with the authentication provider.",
			Required: true,
			BodyPath: "settings",
		},
		&requestflag.Flag[string]{
			Name:     "short-name",
			Usage:    "The short name associated with the authentication provider. This must be unique and URL-friendly, as it's going to be part of the login URL.",
			Required: true,
			BodyPath: "short_name",
		},
		&requestflag.Flag[bool]{
			Name:     "active",
			Usage:    "The active status of the authentication provider",
			Default:  true,
			BodyPath: "active",
		},
		&requestflag.Flag[string]{
			Name:     "settings-url",
			Usage:    "The URL for the identity provider metadata file to populate the settings automatically. If the settings attribute is provided, that will be used instead.",
			BodyPath: "settings_url",
		},
	},
	Action:          handleAuthenticationProvidersCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"settings": {
		&requestflag.InnerFlag[string]{
			Name:       "settings.idp-cert-fingerprint",
			Usage:      "The certificate fingerprint for the identity provider (IdP)",
			InnerField: "idp_cert_fingerprint",
		},
		&requestflag.InnerFlag[string]{
			Name:       "settings.idp-entity-id",
			Usage:      "The Entity ID for the identity provider (IdP).",
			InnerField: "idp_entity_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "settings.idp-sso-target-url",
			Usage:      "The SSO target url for the identity provider (IdP).",
			InnerField: "idp_sso_target_url",
		},
		&requestflag.InnerFlag[string]{
			Name:       "settings.idp-cert-fingerprint-algorithm",
			Usage:      "The algorithm used to generate the identity provider's (IdP) certificate fingerprint",
			InnerField: "idp_cert_fingerprint_algorithm",
		},
	},
})

var authenticationProvidersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves the details of an existing authentication provider.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleAuthenticationProvidersRetrieve,
	HideHelpCommand: true,
}

var authenticationProvidersUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates settings of an existing authentication provider.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:     "active",
			Usage:    "The active status of the authentication provider",
			Default:  true,
			BodyPath: "active",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The name associated with the authentication provider.",
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "settings",
			Usage:    "The settings associated with the authentication provider.",
			BodyPath: "settings",
		},
		&requestflag.Flag[string]{
			Name:     "settings-url",
			Usage:    "The URL for the identity provider metadata file to populate the settings automatically. If the settings attribute is provided, that will be used instead.",
			BodyPath: "settings_url",
		},
		&requestflag.Flag[string]{
			Name:     "short-name",
			Usage:    "The short name associated with the authentication provider. This must be unique and URL-friendly, as it's going to be part of the login URL.",
			BodyPath: "short_name",
		},
	},
	Action:          handleAuthenticationProvidersUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"settings": {
		&requestflag.InnerFlag[string]{
			Name:       "settings.idp-cert-fingerprint",
			Usage:      "The certificate fingerprint for the identity provider (IdP)",
			InnerField: "idp_cert_fingerprint",
		},
		&requestflag.InnerFlag[string]{
			Name:       "settings.idp-entity-id",
			Usage:      "The Entity ID for the identity provider (IdP).",
			InnerField: "idp_entity_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "settings.idp-sso-target-url",
			Usage:      "The SSO target url for the identity provider (IdP).",
			InnerField: "idp_sso_target_url",
		},
		&requestflag.InnerFlag[string]{
			Name:       "settings.idp-cert-fingerprint-algorithm",
			Usage:      "The algorithm used to generate the identity provider's (IdP) certificate fingerprint",
			InnerField: "idp_cert_fingerprint_algorithm",
		},
	},
})

var authenticationProvidersList = cli.Command{
	Name:    "list",
	Usage:   "Returns a list of your SSO authentication providers.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "page-number",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page[size]",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the <code>-</code> prefix.<br/><br/>\nThat is: <ul>\n  <li>\n    <code>name</code>: sorts the result by the\n    <code>name</code> field in ascending order.\n  </li>\n  <li>\n    <code>-name</code>: sorts the result by the\n    <code>name</code> field in descending order.\n  </li>\n</ul><br/>If not given, results are sorted by <code>created_at</code> in descending order.",
			Default:   "-created_at",
			QueryPath: "sort",
		},
	},
	Action:          handleAuthenticationProvidersList,
	HideHelpCommand: true,
}

var authenticationProvidersDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes an existing authentication provider.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleAuthenticationProvidersDelete,
	HideHelpCommand: true,
}

func handleAuthenticationProvidersCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AuthenticationProviderNewParams{}

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
	_, err = client.AuthenticationProviders.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "authentication-providers create", obj, format, transform)
}

func handleAuthenticationProvidersRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.AuthenticationProviders.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "authentication-providers retrieve", obj, format, transform)
}

func handleAuthenticationProvidersUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AuthenticationProviderUpdateParams{}

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
	_, err = client.AuthenticationProviders.Update(
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
	return ShowJSON(os.Stdout, "authentication-providers update", obj, format, transform)
}

func handleAuthenticationProvidersList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AuthenticationProviderListParams{}

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
		_, err = client.AuthenticationProviders.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "authentication-providers list", obj, format, transform)
	} else {
		iter := client.AuthenticationProviders.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "authentication-providers list", iter, format, transform)
	}
}

func handleAuthenticationProvidersDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.AuthenticationProviders.Delete(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "authentication-providers delete", obj, format, transform)
}
