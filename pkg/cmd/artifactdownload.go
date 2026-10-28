// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/team-telnyx/telnyx-cli/internal/apiquery"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/urfave/cli/v3"
)

var artifactDownloadsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Downloads the artifact content. The signed URL (with expires and sig query\nparams) is obtained from the artifact metadata endpoint.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "artifact-id",
			Required:  true,
			PathParam: "artifactId",
		},
		&requestflag.Flag[string]{
			Name:      "expires",
			Required:  true,
			QueryPath: "expires",
		},
		&requestflag.Flag[string]{
			Name:      "sig",
			Required:  true,
			QueryPath: "sig",
		},
		&requestflag.Flag[string]{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "The file where the response contents will be stored. Use the value '-' to force output to stdout.",
		},
	},
	Action:          handleArtifactDownloadsRetrieve,
	HideHelpCommand: true,
}

func handleArtifactDownloadsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("artifact-id") && len(unusedArgs) > 0 {
		cmd.Set("artifact-id", unusedArgs[0])
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

	params := telnyx.ArtifactDownloadGetParams{}

	response, err := client.ArtifactDownloads.Get(
		ctx,
		cmd.Value("artifact-id").(string),
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
