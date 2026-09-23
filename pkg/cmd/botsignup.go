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

var botSignupCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a freemium Telnyx account through the agentic signup flow. The request\nmust carry a valid answer to a previously issued bot challenge\n(`bot_challenge_nonce` and `bot_challenge_answer`), accept the terms of service,\nand echo the exact terms-and-conditions and privacy-policy URLs returned by the\nchallenge endpoint. When EU consent enforcement is enabled,\n`terms_of_service_eu` and `terms_and_conditions_eu_url` are also required. On\nsuccess a one-time sign-in (magic) link is emailed to the address provided; if\nthe email address belongs to an existing account, a sign-in link is sent instead\nof creating a duplicate account. `email` may only be omitted when\nplaceholder-email registration is enabled server-side. This endpoint is public\nand unauthenticated, gated by the freemium feature flags and per-country\navailability, and subject to per-IP and per-domain registration limits.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "bot-challenge-answer",
			Usage:    "Answer to the issued bot challenge.",
			Required: true,
			BodyPath: "bot_challenge_answer",
		},
		&requestflag.Flag[string]{
			Name:     "bot-challenge-nonce",
			Usage:    "Nonce from a previously issued bot challenge.",
			Required: true,
			BodyPath: "bot_challenge_nonce",
		},
		&requestflag.Flag[string]{
			Name:     "privacy-policy-url",
			Usage:    "Must exactly match the privacy-policy URL returned by the challenge endpoint.",
			Required: true,
			BodyPath: "privacy_policy_url",
		},
		&requestflag.Flag[string]{
			Name:     "terms-and-conditions-url",
			Usage:    "Must exactly match the terms-and-conditions URL returned by the challenge endpoint.",
			Required: true,
			BodyPath: "terms_and_conditions_url",
		},
		&requestflag.Flag[bool]{
			Name:     "terms-of-service",
			Usage:    "Must be true to accept the terms of service.",
			Required: true,
			BodyPath: "terms_of_service",
		},
		&requestflag.Flag[string]{
			Name:     "email",
			Usage:    "Email address for the new account. The magic link is sent here. May only be omitted when placeholder-email registration is enabled server-side.",
			BodyPath: "email",
		},
		&requestflag.Flag[string]{
			Name:     "terms-and-conditions-eu-url",
			Usage:    "EU terms-and-conditions URL. Required when EU consent enforcement is enabled.",
			BodyPath: "terms_and_conditions_eu_url",
		},
		&requestflag.Flag[bool]{
			Name:     "terms-of-service-eu",
			Usage:    "EU terms-of-service acceptance. Required when EU consent enforcement is enabled.",
			BodyPath: "terms_of_service_eu",
		},
	},
	Action:          handleBotSignupCreate,
	HideHelpCommand: true,
}

var botSignupResendMagicLink = cli.Command{
	Name:    "resend-magic-link",
	Usage:   "Resends the one-time sign-in (magic) link for an eligible bot signup account.\nEligibility (account exists, was registered through bot signup, is active, and\nhas not exceeded the resend limit or rate window) is evaluated server-side; the\nresponse is intentionally uniform and does not reveal whether the account exists\nor whether a link was actually sent. This endpoint is public and\nunauthenticated, gated by the freemium feature flags and per-country\navailability.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "email",
			Usage:    "Email address of the bot signup account to resend the magic link to.",
			Required: true,
			BodyPath: "email",
		},
	},
	Action:          handleBotSignupResendMagicLink,
	HideHelpCommand: true,
}

func handleBotSignupCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.BotSignupNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.BotSignup.New(ctx, params, options...)
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
		Title:          "bot-signup create",
		Transform:      transform,
	})
}

func handleBotSignupResendMagicLink(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.BotSignupResendMagicLinkParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.BotSignup.ResendMagicLink(ctx, params, options...)
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
		Title:          "bot-signup resend-magic-link",
		Transform:      transform,
	})
}
