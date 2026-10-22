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

var aiAvatarsCreate = cli.Command{
	Name:    "create",
	Usage:   "Uploads an image, validates its file type, size, safe decoding, and resolution,\nstores it in Telnyx-managed private storage, and creates an avatar owned by the\nauthenticated organization. Creating an avatar does not update an AI Assistant;\nuse the returned ID separately wherever the avatar is consumed. The avatar image\nis stored independently from the rendering model and background, which are\nselected where the avatar is consumed. Supported image types are JPEG, PNG, and\nWebP. Images may be up to 10 MiB, must be between 256 and 8192 pixels on each\nside, and may contain at most 40 megapixels. Each organization may have up to 20\ncustom avatars. For best results, use a 16:9 landscape image, preferably 1920 by\n1080 pixels and at least 1280 by 720 pixels. Show one centered, front-facing\nperson from the shoulders up, keep the full head and hair visible with space\naround it, and use a sharp, evenly lit image with unobstructed eyes and mouth.\nUse a transparent PNG or WebP when the avatar will be combined with a selectable\nbackground. These composition recommendations are not enforced in the first\nrelease; the first release does not validate the number of faces, pose, or image\ncontent.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "image-file",
			Usage:     "JPEG, PNG, or WebP image file. The maximum file size is 10 MiB.",
			Required:  true,
			BodyPath:  "image_file",
			FileInput: true,
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Customer-facing name of the custom avatar.",
			Required: true,
			BodyPath: "name",
		},
	},
	Action:          handleAIAvatarsCreate,
	HideHelpCommand: true,
}

var aiAvatarsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns a Telnyx stock avatar or a custom avatar owned by the authenticated\norganization. A custom avatar belonging to another organization is returned as\nnot found.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "avatar-id",
			Usage:     "Stable identifier returned by the Avatars API. Discover available stock IDs with GET /v2/ai/avatars and filter[avatar_type]=stock rather than maintaining a fixed list. Custom avatar IDs use the `avatar_` prefix and must belong to the authenticated organization. The service validates that the selected avatar is available.",
			Required:  true,
			PathParam: "avatar_id",
		},
	},
	Action:          handleAIAvatarsRetrieve,
	HideHelpCommand: true,
}

var aiAvatarsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates the customer-facing name of a custom avatar owned by the authenticated\norganization. Stock avatars cannot be modified. To replace the image, create a\nnew custom avatar so existing references keep their original content.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "avatar-id",
			Usage:     "Stable identifier returned by the Avatars API. Discover available stock IDs with GET /v2/ai/avatars and filter[avatar_type]=stock rather than maintaining a fixed list. Custom avatar IDs use the `avatar_` prefix and must belong to the authenticated organization. The service validates that the selected avatar is available.",
			Required:  true,
			PathParam: "avatar_id",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "New customer-facing name of the custom avatar.",
			Required: true,
			BodyPath: "name",
		},
	},
	Action:          handleAIAvatarsUpdate,
	HideHelpCommand: true,
}

var aiAvatarsList = cli.Command{
	Name:    "list",
	Usage:   "Returns a paginated list of the Telnyx stock avatars and custom avatars\navailable to the authenticated organization. Custom avatars owned by other\norganizations are never returned. The `image_url` value is suitable for\ndisplaying a thumbnail and may be refreshed by listing or retrieving the avatar\nagain. `suggested_voice_id` recommends a compatible Telnyx voice for stock\navatars and is an empty string for custom avatars; it does not change an AI\nAssistant's voice automatically.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "filter-avatar-type",
			Usage:     "Whether the avatar is provided by Telnyx or was uploaded by the authenticated organization.",
			QueryPath: "filter[avatar_type]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-id",
			Usage:     "Stable identifier returned by the Avatars API. Discover available stock IDs with GET /v2/ai/avatars and filter[avatar_type]=stock rather than maintaining a fixed list. Custom avatar IDs use the `avatar_` prefix and must belong to the authenticated organization. The service validates that the selected avatar is available.",
			QueryPath: "filter[id]",
		},
		&requestflag.Flag[string]{
			Name:      "filter-name",
			Usage:     "Return avatars whose names contain this value, using case-insensitive matching.",
			QueryPath: "filter[name]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "Page number to retrieve, starting at 1.",
			Default:   1,
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Number of avatars to return per page.",
			Default:   20,
			QueryPath: "page[size]",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleAIAvatarsList,
	HideHelpCommand: true,
}

func handleAIAvatarsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := telnyx.AIAvatarNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Avatars.New(ctx, params, options...)
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
		Title:          "ai:avatars create",
		Transform:      transform,
	})
}

func handleAIAvatarsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("avatar-id") && len(unusedArgs) > 0 {
		cmd.Set("avatar-id", unusedArgs[0])
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
	_, err = client.AI.Avatars.Get(ctx, telnyx.AvatarIDParam(cmd.Value("avatar-id").(string)), options...)
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
		Title:          "ai:avatars retrieve",
		Transform:      transform,
	})
}

func handleAIAvatarsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("avatar-id") && len(unusedArgs) > 0 {
		cmd.Set("avatar-id", unusedArgs[0])
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

	params := telnyx.AIAvatarUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AI.Avatars.Update(
		ctx,
		telnyx.AvatarIDParam(cmd.Value("avatar-id").(string)),
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
		Title:          "ai:avatars update",
		Transform:      transform,
	})
}

func handleAIAvatarsList(ctx context.Context, cmd *cli.Command) error {
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

	params := telnyx.AIAvatarListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.AI.Avatars.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "ai:avatars list",
			Transform:      transform,
		})
	} else {
		iter := client.AI.Avatars.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "ai:avatars list",
			Transform:      transform,
		})
	}
}
