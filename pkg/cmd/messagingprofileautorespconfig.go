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

var messagingProfilesAutorespConfigsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create auto-response setting",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "profile-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "country-code",
			Required: true,
			BodyPath: "country_code",
		},
		&requestflag.Flag[[]string]{
			Name:     "keyword",
			Required: true,
			BodyPath: "keywords",
		},
		&requestflag.Flag[string]{
			Name:     "op",
			Required: true,
			BodyPath: "op",
		},
		&requestflag.Flag[string]{
			Name:     "resp-text",
			BodyPath: "resp_text",
		},
	},
	Action:          handleMessagingProfilesAutorespConfigsCreate,
	HideHelpCommand: true,
}

var messagingProfilesAutorespConfigsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get Auto-Response Setting",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "profile-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "autoresp-cfg-id",
			Required: true,
		},
	},
	Action:          handleMessagingProfilesAutorespConfigsRetrieve,
	HideHelpCommand: true,
}

var messagingProfilesAutorespConfigsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update Auto-Response Setting",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "profile-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "autoresp-cfg-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "country-code",
			Required: true,
			BodyPath: "country_code",
		},
		&requestflag.Flag[[]string]{
			Name:     "keyword",
			Required: true,
			BodyPath: "keywords",
		},
		&requestflag.Flag[string]{
			Name:     "op",
			Required: true,
			BodyPath: "op",
		},
		&requestflag.Flag[string]{
			Name:     "resp-text",
			BodyPath: "resp_text",
		},
	},
	Action:          handleMessagingProfilesAutorespConfigsUpdate,
	HideHelpCommand: true,
}

var messagingProfilesAutorespConfigsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "List Auto-Response Settings",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "profile-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "country-code",
			QueryPath: "country_code",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "created-at",
			Usage:     "Consolidated created_at parameter (deepObject style). Originally: created_at[gte], created_at[lte]",
			QueryPath: "created_at",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "updated-at",
			Usage:     "Consolidated updated_at parameter (deepObject style). Originally: updated_at[gte], updated_at[lte]",
			QueryPath: "updated_at",
		},
	},
	Action:          handleMessagingProfilesAutorespConfigsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"created-at": {
		&requestflag.InnerFlag[string]{
			Name:       "created-at.gte",
			InnerField: "gte",
		},
		&requestflag.InnerFlag[string]{
			Name:       "created-at.lte",
			InnerField: "lte",
		},
	},
	"updated-at": {
		&requestflag.InnerFlag[string]{
			Name:       "updated-at.gte",
			InnerField: "gte",
		},
		&requestflag.InnerFlag[string]{
			Name:       "updated-at.lte",
			InnerField: "lte",
		},
	},
})

var messagingProfilesAutorespConfigsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete Auto-Response Setting",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "profile-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "autoresp-cfg-id",
			Required: true,
		},
	},
	Action:          handleMessagingProfilesAutorespConfigsDelete,
	HideHelpCommand: true,
}

func handleMessagingProfilesAutorespConfigsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("profile-id") && len(unusedArgs) > 0 {
		cmd.Set("profile-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MessagingProfileAutorespConfigNewParams{}

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
	_, err = client.MessagingProfiles.AutorespConfigs.New(
		ctx,
		cmd.Value("profile-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-profiles:autoresp-configs create", obj, format, transform)
}

func handleMessagingProfilesAutorespConfigsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("autoresp-cfg-id") && len(unusedArgs) > 0 {
		cmd.Set("autoresp-cfg-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MessagingProfileAutorespConfigGetParams{
		ProfileID: cmd.Value("profile-id").(string),
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
	_, err = client.MessagingProfiles.AutorespConfigs.Get(
		ctx,
		cmd.Value("autoresp-cfg-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-profiles:autoresp-configs retrieve", obj, format, transform)
}

func handleMessagingProfilesAutorespConfigsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("autoresp-cfg-id") && len(unusedArgs) > 0 {
		cmd.Set("autoresp-cfg-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MessagingProfileAutorespConfigUpdateParams{
		ProfileID: cmd.Value("profile-id").(string),
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
	_, err = client.MessagingProfiles.AutorespConfigs.Update(
		ctx,
		cmd.Value("autoresp-cfg-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-profiles:autoresp-configs update", obj, format, transform)
}

func handleMessagingProfilesAutorespConfigsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("profile-id") && len(unusedArgs) > 0 {
		cmd.Set("profile-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MessagingProfileAutorespConfigListParams{}

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
	_, err = client.MessagingProfiles.AutorespConfigs.List(
		ctx,
		cmd.Value("profile-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-profiles:autoresp-configs list", obj, format, transform)
}

func handleMessagingProfilesAutorespConfigsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("autoresp-cfg-id") && len(unusedArgs) > 0 {
		cmd.Set("autoresp-cfg-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.MessagingProfileAutorespConfigDeleteParams{
		ProfileID: cmd.Value("profile-id").(string),
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
	_, err = client.MessagingProfiles.AutorespConfigs.Delete(
		ctx,
		cmd.Value("autoresp-cfg-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "messaging-profiles:autoresp-configs delete", obj, format, transform)
}
