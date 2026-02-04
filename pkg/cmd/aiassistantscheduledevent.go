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

var aiAssistantsScheduledEventsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a scheduled event for an assistant",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "assistant-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "scheduled-at-fixed-datetime",
			Usage:    "The datetime at which the event should be scheduled. Formatted as ISO 8601.",
			Required: true,
			BodyPath: "scheduled_at_fixed_datetime",
		},
		&requestflag.Flag[string]{
			Name:     "telnyx-agent-target",
			Usage:    "The phone number, SIP URI, to schedule the call or text from.",
			Required: true,
			BodyPath: "telnyx_agent_target",
		},
		&requestflag.Flag[string]{
			Name:     "telnyx-conversation-channel",
			Required: true,
			BodyPath: "telnyx_conversation_channel",
		},
		&requestflag.Flag[string]{
			Name:     "telnyx-end-user-target",
			Usage:    "The phone number, SIP URI, to schedule the call or text to.",
			Required: true,
			BodyPath: "telnyx_end_user_target",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "conversation-metadata",
			Usage:    "Metadata associated with the conversation. Telnyx provides several pieces of metadata, but customers can also add their own.",
			BodyPath: "conversation_metadata",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "Required for sms scheduled events. The text to be sent to the end user.",
			BodyPath: "text",
		},
	},
	Action:          handleAIAssistantsScheduledEventsCreate,
	HideHelpCommand: true,
}

var aiAssistantsScheduledEventsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a scheduled event by event ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "assistant-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "event-id",
			Required: true,
		},
	},
	Action:          handleAIAssistantsScheduledEventsRetrieve,
	HideHelpCommand: true,
}

var aiAssistantsScheduledEventsList = cli.Command{
	Name:    "list",
	Usage:   "Get scheduled events for an assistant with pagination and filtering",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "assistant-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "conversation-channel",
			QueryPath: "conversation_channel",
		},
		&requestflag.Flag[any]{
			Name:      "from-date",
			QueryPath: "from_date",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			QueryPath: "page[number]",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page[size]",
		},
		&requestflag.Flag[any]{
			Name:      "to-date",
			QueryPath: "to_date",
		},
	},
	Action:          handleAIAssistantsScheduledEventsList,
	HideHelpCommand: true,
}

var aiAssistantsScheduledEventsDelete = cli.Command{
	Name:    "delete",
	Usage:   "If the event is pending, this will cancel the event. Otherwise, this will simply\nremove the record of the event.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "assistant-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "event-id",
			Required: true,
		},
	},
	Action:          handleAIAssistantsScheduledEventsDelete,
	HideHelpCommand: true,
}

func handleAIAssistantsScheduledEventsCreate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("assistant-id") && len(unusedArgs) > 0 {
		cmd.Set("assistant-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIAssistantScheduledEventNewParams{}

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
	_, err = client.AI.Assistants.ScheduledEvents.New(
		ctx,
		cmd.Value("assistant-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:assistants:scheduled-events create", obj, format, transform)
}

func handleAIAssistantsScheduledEventsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("event-id") && len(unusedArgs) > 0 {
		cmd.Set("event-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIAssistantScheduledEventGetParams{
		AssistantID: cmd.Value("assistant-id").(string),
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
	_, err = client.AI.Assistants.ScheduledEvents.Get(
		ctx,
		cmd.Value("event-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:assistants:scheduled-events retrieve", obj, format, transform)
}

func handleAIAssistantsScheduledEventsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("assistant-id") && len(unusedArgs) > 0 {
		cmd.Set("assistant-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIAssistantScheduledEventListParams{}

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
		_, err = client.AI.Assistants.ScheduledEvents.List(
			ctx,
			cmd.Value("assistant-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "ai:assistants:scheduled-events list", obj, format, transform)
	} else {
		iter := client.AI.Assistants.ScheduledEvents.ListAutoPaging(
			ctx,
			cmd.Value("assistant-id").(string),
			params,
			options...,
		)
		return ShowJSONIterator(os.Stdout, "ai:assistants:scheduled-events list", iter, format, transform)
	}
}

func handleAIAssistantsScheduledEventsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("event-id") && len(unusedArgs) > 0 {
		cmd.Set("event-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIAssistantScheduledEventDeleteParams{
		AssistantID: cmd.Value("assistant-id").(string),
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

	return client.AI.Assistants.ScheduledEvents.Delete(
		ctx,
		cmd.Value("event-id").(string),
		params,
		options...,
	)
}
