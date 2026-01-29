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

var storageBucketsCreatePresignedURL = cli.Command{
	Name:    "create-presigned-url",
	Usage:   "Returns a timed and authenticated URL to download (GET) or upload (PUT) an\nobject. This is the equivalent to AWS S3’s “presigned” URL. Please note that\nTelnyx performs authentication differently from AWS S3 and you MUST NOT use the\npresign method of AWS s3api CLI or SDK to generate the presigned URL.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "bucket-name",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "object-name",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:     "ttl",
			Usage:    "The time to live of the token in seconds",
			BodyPath: "ttl",
		},
	},
	Action:          handleStorageBucketsCreatePresignedURL,
	HideHelpCommand: true,
}

func handleStorageBucketsCreatePresignedURL(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("object-name") && len(unusedArgs) > 0 {
		cmd.Set("object-name", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.StorageBucketNewPresignedURLParams{
		BucketName: cmd.Value("bucket-name").(string),
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Storage.Buckets.NewPresignedURL(
		ctx,
		cmd.Value("object-name").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "storage:buckets create-presigned-url", obj, format, transform)
}
