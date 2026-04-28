// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/team-telnyx/telnyx-cli/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/urfave/cli/v3"
)

var termsOfServiceNumberReputationAgree = cli.Command{
	Name:            "agree",
	Usage:           "Accept the Terms of Service for the Number Reputation product. Must be called\nbefore using Number Reputation endpoints.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleTermsOfServiceNumberReputationAgree,
	HideHelpCommand: true,
}

func handleTermsOfServiceNumberReputationAgree(ctx context.Context, cmd *cli.Command) error {
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

	return client.TermsOfService.NumberReputation.Agree(ctx, options...)
}
