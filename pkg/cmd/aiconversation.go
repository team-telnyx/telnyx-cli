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

var aiConversationsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new AI Conversation.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Metadata associated with the conversation.",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
	},
	Action:          handleAIConversationsCreate,
	HideHelpCommand: true,
}

var aiConversationsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a specific AI conversation by its ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "conversation-id",
			Required: true,
		},
	},
	Action:          handleAIConversationsRetrieve,
	HideHelpCommand: true,
}

var aiConversationsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update metadata for a specific conversation.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "conversation-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Metadata associated with the conversation.",
			BodyPath: "metadata",
		},
	},
	Action:          handleAIConversationsUpdate,
	HideHelpCommand: true,
}

var aiConversationsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve a list of all AI conversations configured by the user. Supports\n[PostgREST-style query parameters](https://postgrest.org/en/stable/api.html#horizontal-filtering-rows)\nfor filtering. Examples are included for the standard metadata fields, but you\ncan filter on any field in the metadata JSON object. For example, to filter by a\ncustom field `metadata->custom_field`, use `metadata->custom_field=eq.value`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Filter by conversation ID (e.g. id=eq.123)",
			QueryPath: "id",
		},
		&requestflag.Flag[string]{
			Name:      "created-at",
			Usage:     "Filter by creation datetime (e.g., `created_at=gte.2025-01-01`)",
			QueryPath: "created_at",
		},
		&requestflag.Flag[string]{
			Name:      "last-message-at",
			Usage:     "Filter by last message datetime (e.g., `last_message_at=lte.2025-06-01`)",
			QueryPath: "last_message_at",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Limit the number of returned conversations (e.g., `limit=10`)",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "metadata-assistant-id",
			Usage:     "Filter by assistant ID (e.g., `metadata->assistant_id=eq.assistant-123`)",
			QueryPath: "metadata->assistant_id",
		},
		&requestflag.Flag[string]{
			Name:      "metadata-call-control-id",
			Usage:     "Filter by call control ID (e.g., `metadata->call_control_id=eq.v3:123`)",
			QueryPath: "metadata->call_control_id",
		},
		&requestflag.Flag[string]{
			Name:      "metadata-telnyx-agent-target",
			Usage:     "Filter by the phone number, SIP URI, or other identifier for the agent (e.g., `metadata->telnyx_agent_target=eq.+13128675309`)",
			QueryPath: "metadata->telnyx_agent_target",
		},
		&requestflag.Flag[string]{
			Name:      "metadata-telnyx-conversation-channel",
			Usage:     "Filter by conversation channel (e.g., `metadata->telnyx_conversation_channel=eq.phone_call`)",
			QueryPath: "metadata->telnyx_conversation_channel",
		},
		&requestflag.Flag[string]{
			Name:      "metadata-telnyx-end-user-target",
			Usage:     "Filter by the phone number, SIP URI, or other identifier for the end user (e.g., `metadata->telnyx_end_user_target=eq.+13128675309`)",
			QueryPath: "metadata->telnyx_end_user_target",
		},
		&requestflag.Flag[string]{
			Name:      "name",
			Usage:     "Filter by conversation Name (e.g. `name=like.Voice%`)",
			QueryPath: "name",
		},
		&requestflag.Flag[string]{
			Name:      "or",
			Usage:     "Apply OR conditions using PostgREST syntax (e.g., `or=(created_at.gte.2025-04-01,last_message_at.gte.2025-04-01)`)",
			QueryPath: "or",
		},
		&requestflag.Flag[string]{
			Name:      "order",
			Usage:     "Order the results by specific fields (e.g., `order=created_at.desc` or `order=last_message_at.asc`)",
			QueryPath: "order",
		},
	},
	Action:          handleAIConversationsList,
	HideHelpCommand: true,
}

var aiConversationsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a specific conversation by its ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "conversation-id",
			Required: true,
		},
	},
	Action:          handleAIConversationsDelete,
	HideHelpCommand: true,
}

var aiConversationsAddMessage = cli.Command{
	Name:    "add-message",
	Usage:   "Add a new message to the conversation. Used to insert a new messages to a\nconversation manually ( without using chat endpoint )",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "conversation-id",
			Usage:    "The ID of the conversation",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "role",
			Required: true,
			BodyPath: "role",
		},
		&requestflag.Flag[string]{
			Name:     "content",
			Default:  "",
			BodyPath: "content",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "sent-at",
			BodyPath: "sent_at",
		},
		&requestflag.Flag[string]{
			Name:     "tool-call-id",
			BodyPath: "tool_call_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "tool-call",
			BodyPath: "tool_calls",
		},
		&requestflag.Flag[any]{
			Name:     "tool-choice",
			BodyPath: "tool_choice",
		},
	},
	Action:          handleAIConversationsAddMessage,
	HideHelpCommand: true,
}

var aiConversationsRetrieveConversationsInsights = cli.Command{
	Name:    "retrieve-conversations-insights",
	Usage:   "Retrieve insights for a specific conversation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "conversation-id",
			Required: true,
		},
	},
	Action:          handleAIConversationsRetrieveConversationsInsights,
	HideHelpCommand: true,
}

func handleAIConversationsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIConversationNewParams{}

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
	_, err = client.AI.Conversations.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:conversations create", obj, format, transform)
}

func handleAIConversationsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("conversation-id") && len(unusedArgs) > 0 {
		cmd.Set("conversation-id", unusedArgs[0])
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
	_, err = client.AI.Conversations.Get(ctx, cmd.Value("conversation-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:conversations retrieve", obj, format, transform)
}

func handleAIConversationsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("conversation-id") && len(unusedArgs) > 0 {
		cmd.Set("conversation-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIConversationUpdateParams{}

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
	_, err = client.AI.Conversations.Update(
		ctx,
		cmd.Value("conversation-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:conversations update", obj, format, transform)
}

func handleAIConversationsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIConversationListParams{}

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
	_, err = client.AI.Conversations.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:conversations list", obj, format, transform)
}

func handleAIConversationsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("conversation-id") && len(unusedArgs) > 0 {
		cmd.Set("conversation-id", unusedArgs[0])
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

	return client.AI.Conversations.Delete(ctx, cmd.Value("conversation-id").(string), options...)
}

func handleAIConversationsAddMessage(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("conversation-id") && len(unusedArgs) > 0 {
		cmd.Set("conversation-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIConversationAddMessageParams{}

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

	return client.AI.Conversations.AddMessage(
		ctx,
		cmd.Value("conversation-id").(string),
		params,
		options...,
	)
}

func handleAIConversationsRetrieveConversationsInsights(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("conversation-id") && len(unusedArgs) > 0 {
		cmd.Set("conversation-id", unusedArgs[0])
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
	_, err = client.AI.Conversations.GetConversationsInsights(ctx, cmd.Value("conversation-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:conversations retrieve-conversations-insights", obj, format, transform)
}
