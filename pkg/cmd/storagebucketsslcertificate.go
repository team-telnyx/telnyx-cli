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

var storageBucketsSslCertificateCreate = cli.Command{
	Name:    "create",
	Usage:   "Uploads an SSL certificate and its matching secret so that you can use Telnyx's\nstorage as your CDN.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "bucket-name",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "certificate",
			Usage:    "The SSL certificate file",
			BodyPath: "certificate",
		},
		&requestflag.Flag[string]{
			Name:     "private-key",
			Usage:    "The private key file",
			BodyPath: "private_key",
		},
	},
	Action:          handleStorageBucketsSslCertificateCreate,
	HideHelpCommand: true,
}

var storageBucketsSslCertificateRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns the stored certificate detail of a bucket, if applicable.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "bucket-name",
			Required: true,
		},
	},
	Action:          handleStorageBucketsSslCertificateRetrieve,
	HideHelpCommand: true,
}

var storageBucketsSslCertificateDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes an SSL certificate and its matching secret.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "bucket-name",
			Required: true,
		},
	},
	Action:          handleStorageBucketsSslCertificateDelete,
	HideHelpCommand: true,
}

func handleStorageBucketsSslCertificateCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("bucket-name") && len(unusedArgs) > 0 {
		cmd.Set("bucket-name", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.StorageBucketSslCertificateNewParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		MultipartFormEncoded,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Storage.Buckets.SslCertificate.New(
		ctx,
		cmd.Value("bucket-name").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "storage:buckets:ssl-certificate create", obj, format, transform)
}

func handleStorageBucketsSslCertificateRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("bucket-name") && len(unusedArgs) > 0 {
		cmd.Set("bucket-name", unusedArgs[0])
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
	_, err = client.Storage.Buckets.SslCertificate.Get(ctx, cmd.Value("bucket-name").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "storage:buckets:ssl-certificate retrieve", obj, format, transform)
}

func handleStorageBucketsSslCertificateDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("bucket-name") && len(unusedArgs) > 0 {
		cmd.Set("bucket-name", unusedArgs[0])
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
	_, err = client.Storage.Buckets.SslCertificate.Delete(ctx, cmd.Value("bucket-name").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "storage:buckets:ssl-certificate delete", obj, format, transform)
}
