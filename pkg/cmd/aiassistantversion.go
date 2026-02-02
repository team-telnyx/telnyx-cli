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

var aiAssistantsVersionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a specific version of an assistant by assistant_id and version_id",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "assistant-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "version-id",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:      "include-mcp-servers",
			QueryPath: "include_mcp_servers",
		},
	},
	Action:          handleAIAssistantsVersionsRetrieve,
	HideHelpCommand: true,
}

var aiAssistantsVersionsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates the configuration of a specific assistant version. Can not update main\nversion",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "assistant-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "version-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "dynamic-variables",
			Usage:    "Map of dynamic variables and their default values",
			BodyPath: "dynamic_variables",
		},
		&requestflag.Flag[string]{
			Name:     "dynamic-variables-webhook-url",
			Usage:    "If the dynamic_variables_webhook_url is set for the assistant, we will send a request at the start of the conversation. See our [guide](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables) for more information.",
			BodyPath: "dynamic_variables_webhook_url",
		},
		&requestflag.Flag[[]string]{
			Name:     "enabled-feature",
			BodyPath: "enabled_features",
		},
		&requestflag.Flag[string]{
			Name:     "greeting",
			Usage:    "Text that the assistant will use to start the conversation. This may be templated with [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables). Use an empty string to have the assistant wait for the user to speak first. Use the special value `<assistant-speaks-first-with-model-generated-message>` to have the assistant generate the greeting based on the system instructions.",
			BodyPath: "greeting",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "insight-settings",
			BodyPath: "insight_settings",
		},
		&requestflag.Flag[string]{
			Name:     "instructions",
			Usage:    "System instructions for the assistant. These may be templated with [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables)",
			BodyPath: "instructions",
		},
		&requestflag.Flag[string]{
			Name:     "llm-api-key-ref",
			Usage:    "This is only needed when using third-party inference providers. The `identifier` for an integration secret [/v2/integration_secrets](https://developers.telnyx.com/api-reference/integration-secrets/create-a-secret) that refers to your LLM provider's API key. Warning: Free plans are unlikely to work with this integration.",
			BodyPath: "llm_api_key_ref",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "messaging-settings",
			BodyPath: "messaging_settings",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			Usage:    "ID of the model to use. You can use the [Get models API](https://developers.telnyx.com/api-reference/chat/get-available-models) to see all of your available models,",
			BodyPath: "model",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "privacy-settings",
			BodyPath: "privacy_settings",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "telephony-settings",
			BodyPath: "telephony_settings",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "tool",
			Usage:    "The tools that the assistant can use. These may be templated with [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables)",
			BodyPath: "tools",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "transcription",
			BodyPath: "transcription",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "voice-settings",
			BodyPath: "voice_settings",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "widget-settings",
			Usage:    "Configuration settings for the assistant's web widget.",
			BodyPath: "widget_settings",
		},
	},
	Action:          handleAIAssistantsVersionsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"insight-settings": {
		&requestflag.InnerFlag[string]{
			Name:       "insight-settings.insight-group-id",
			Usage:      "Reference to an Insight Group. Insights in this group will be run automatically for all the assistant's conversations.",
			InnerField: "insight_group_id",
		},
	},
	"messaging-settings": {
		&requestflag.InnerFlag[int64]{
			Name:       "messaging-settings.conversation-inactivity-minutes",
			Usage:      "If more than this many minutes have passed since the last message, the assistant will start a new conversation instead of continuing the existing one.",
			InnerField: "conversation_inactivity_minutes",
		},
		&requestflag.InnerFlag[string]{
			Name:       "messaging-settings.default-messaging-profile-id",
			Usage:      "Default Messaging Profile used for messaging exchanges with your assistant. This will be created automatically on assistant creation.",
			InnerField: "default_messaging_profile_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "messaging-settings.delivery-status-webhook-url",
			Usage:      "The URL where webhooks related to delivery statused for assistant messages will be sent.",
			InnerField: "delivery_status_webhook_url",
		},
	},
	"privacy-settings": {
		&requestflag.InnerFlag[bool]{
			Name:       "privacy-settings.data-retention",
			Usage:      "If true, conversation history and insights will be stored. If false, they will not be stored. This in‑tool toggle governs solely the retention of conversation history and insights via the AI assistant. It has no effect on any separate recording, transcription, or storage configuration that you have set at the account, number, or application level. All such external settings remain in force regardless of your selection here.",
			InnerField: "data_retention",
		},
	},
	"telephony-settings": {
		&requestflag.InnerFlag[string]{
			Name:       "telephony-settings.default-texml-app-id",
			Usage:      "Default Texml App used for voice calls with your assistant. This will be created automatically on assistant creation.",
			InnerField: "default_texml_app_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "telephony-settings.noise-suppression",
			Usage:      "The noise suppression engine to use. Use 'disabled' to turn off noise suppression.",
			InnerField: "noise_suppression",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "telephony-settings.noise-suppression-config",
			Usage:      "Configuration for noise suppression. Only applicable when noise_suppression is 'deepfilternet'.",
			InnerField: "noise_suppression_config",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "telephony-settings.supports-unauthenticated-web-calls",
			Usage:      "When enabled, allows users to interact with your AI assistant directly from your website without requiring authentication. This is required for FE widgets that work with assistants that have telephony enabled.",
			InnerField: "supports_unauthenticated_web_calls",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "telephony-settings.time-limit-secs",
			Usage:      "Maximum duration in seconds for the AI assistant to participate on the call. When this limit is reached the assistant will be stopped. This limit does not apply to portions of a call without an active assistant (for instance, a call transferred to a human representative).",
			InnerField: "time_limit_secs",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "telephony-settings.user-idle-timeout-secs",
			Usage:      "Maximum duration in seconds of end user silence on the call. When this limit is reached the assistant will be stopped. This limit does not apply to portions of a call without an active assistant (for instance, a call transferred to a human representative).",
			InnerField: "user_idle_timeout_secs",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "telephony-settings.voicemail-detection",
			Usage:      "Configuration for voicemail detection (AMD - Answering Machine Detection) on outgoing calls. These settings only apply if AMD is enabled on the Dial command. See [TeXML Dial documentation](https://developers.telnyx.com/api-reference/texml-rest-commands/initiate-an-outbound-call) for enabling AMD. Recommended settings: MachineDetection=Enable, AsyncAmd=true, DetectionMode=Premium.",
			InnerField: "voicemail_detection",
		},
	},
	"transcription": {
		&requestflag.InnerFlag[string]{
			Name:       "transcription.language",
			Usage:      "The language of the audio to be transcribed. If not set, of if set to `auto`, the model will automatically detect the language.",
			InnerField: "language",
		},
		&requestflag.InnerFlag[string]{
			Name:       "transcription.model",
			Usage:      "The speech to text model to be used by the voice assistant. All the deepgram models are run on-premise.\n\n- `deepgram/flux` is optimized for turn-taking but is English-only.\n- `deepgram/nova-3` is multi-lingual with automatic language detection but slightly higher latency.",
			InnerField: "model",
		},
		&requestflag.InnerFlag[string]{
			Name:       "transcription.region",
			Usage:      "Region on third party cloud providers (currently Azure) if using one of their models",
			InnerField: "region",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "transcription.settings",
			InnerField: "settings",
		},
	},
	"voice-settings": {
		&requestflag.InnerFlag[string]{
			Name:       "voice-settings.voice",
			Usage:      "The voice to be used by the voice assistant. Check the full list of [available voices](https://developers.telnyx.com/api/call-control/list-text-to-speech-voices) via our voices API.\nTo use ElevenLabs, you must reference your ElevenLabs API key as an integration secret under the `api_key_ref` field. See [integration secrets documentation](https://developers.telnyx.com/api-reference/integration-secrets/create-a-secret) for details. For Telnyx voices, use `Telnyx.<model_id>.<voice_id>` (e.g. Telnyx.KokoroTTS.af_heart)",
			InnerField: "voice",
		},
		&requestflag.InnerFlag[string]{
			Name:       "voice-settings.api-key-ref",
			Usage:      "The `identifier` for an integration secret [/v2/integration_secrets](https://developers.telnyx.com/api-reference/integration-secrets/create-a-secret) that refers to your ElevenLabs API key. Warning: Free plans are unlikely to work with this integration.",
			InnerField: "api_key_ref",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "voice-settings.background-audio",
			Usage:      "Optional background audio to play on the call. Use a predefined media bed, or supply a looped MP3 URL. If a media URL is chosen in the portal, customers can preview it before saving.",
			InnerField: "background_audio",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "voice-settings.similarity-boost",
			Usage:      "Determines how closely the AI should adhere to the original voice when attempting to replicate it. Only applicable when using ElevenLabs.",
			InnerField: "similarity_boost",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "voice-settings.speed",
			Usage:      "Adjusts speech velocity. 1.0 is default speed; values less than 1.0 slow speech; values greater than 1.0 accelerate it. Only applicable when using ElevenLabs.",
			InnerField: "speed",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "voice-settings.style",
			Usage:      "Determines the style exaggeration of the voice. Amplifies speaker style but consumes additional resources when set above 0. Only applicable when using ElevenLabs.",
			InnerField: "style",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "voice-settings.temperature",
			Usage:      "Determines how stable the voice is and the randomness between each generation. Lower values create a broader emotional range; higher values produce more consistent, monotonous output. Only applicable when using ElevenLabs.",
			InnerField: "temperature",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "voice-settings.use-speaker-boost",
			Usage:      "Amplifies similarity to the original speaker voice. Increases computational load and latency slightly. Only applicable when using ElevenLabs.",
			InnerField: "use_speaker_boost",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "voice-settings.voice-speed",
			Usage:      "The speed of the voice in the range [0.25, 2.0]. 1.0 is deafult speed. Larger numbers make the voice faster, smaller numbers make it slower. This is only applicable for Telnyx Natural voices.",
			InnerField: "voice_speed",
		},
	},
	"widget-settings": {
		&requestflag.InnerFlag[string]{
			Name:       "widget-settings.agent-thinking-text",
			Usage:      "Text displayed while the agent is processing.",
			InnerField: "agent_thinking_text",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "widget-settings.audio-visualizer-config",
			InnerField: "audio_visualizer_config",
		},
		&requestflag.InnerFlag[string]{
			Name:       "widget-settings.default-state",
			Usage:      "The default state of the widget.",
			InnerField: "default_state",
		},
		&requestflag.InnerFlag[any]{
			Name:       "widget-settings.give-feedback-url",
			Usage:      "URL for users to give feedback.",
			InnerField: "give_feedback_url",
		},
		&requestflag.InnerFlag[any]{
			Name:       "widget-settings.logo-icon-url",
			Usage:      "URL to a custom logo icon for the widget.",
			InnerField: "logo_icon_url",
		},
		&requestflag.InnerFlag[string]{
			Name:       "widget-settings.position",
			Usage:      "The positioning style for the widget.",
			InnerField: "position",
		},
		&requestflag.InnerFlag[any]{
			Name:       "widget-settings.report-issue-url",
			Usage:      "URL for users to report issues.",
			InnerField: "report_issue_url",
		},
		&requestflag.InnerFlag[string]{
			Name:       "widget-settings.speak-to-interrupt-text",
			Usage:      "Text prompting users to speak to interrupt.",
			InnerField: "speak_to_interrupt_text",
		},
		&requestflag.InnerFlag[string]{
			Name:       "widget-settings.start-call-text",
			Usage:      "Custom text displayed on the start call button.",
			InnerField: "start_call_text",
		},
		&requestflag.InnerFlag[string]{
			Name:       "widget-settings.theme",
			Usage:      "The visual theme for the widget.",
			InnerField: "theme",
		},
		&requestflag.InnerFlag[any]{
			Name:       "widget-settings.view-history-url",
			Usage:      "URL to view conversation history.",
			InnerField: "view_history_url",
		},
	},
})

var aiAssistantsVersionsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieves all versions of a specific assistant with complete configuration and\nmetadata",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "assistant-id",
			Required: true,
		},
	},
	Action:          handleAIAssistantsVersionsList,
	HideHelpCommand: true,
}

var aiAssistantsVersionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Permanently removes a specific version of an assistant. Can not delete main\nversion",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "assistant-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "version-id",
			Required: true,
		},
	},
	Action:          handleAIAssistantsVersionsDelete,
	HideHelpCommand: true,
}

var aiAssistantsVersionsPromote = cli.Command{
	Name:    "promote",
	Usage:   "Promotes a specific version to be the main/current version of the assistant.\nThis will delete any existing canary deploy configuration and send all live\nproduction traffic to this version.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "assistant-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "version-id",
			Required: true,
		},
	},
	Action:          handleAIAssistantsVersionsPromote,
	HideHelpCommand: true,
}

func handleAIAssistantsVersionsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("version-id") && len(unusedArgs) > 0 {
		cmd.Set("version-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIAssistantVersionGetParams{
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
	_, err = client.AI.Assistants.Versions.Get(
		ctx,
		cmd.Value("version-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:assistants:versions retrieve", obj, format, transform)
}

func handleAIAssistantsVersionsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("version-id") && len(unusedArgs) > 0 {
		cmd.Set("version-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIAssistantVersionUpdateParams{
		AssistantID: cmd.Value("assistant-id").(string),
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
	_, err = client.AI.Assistants.Versions.Update(
		ctx,
		cmd.Value("version-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:assistants:versions update", obj, format, transform)
}

func handleAIAssistantsVersionsList(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("assistant-id") && len(unusedArgs) > 0 {
		cmd.Set("assistant-id", unusedArgs[0])
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
	_, err = client.AI.Assistants.Versions.List(ctx, cmd.Value("assistant-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:assistants:versions list", obj, format, transform)
}

func handleAIAssistantsVersionsDelete(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("version-id") && len(unusedArgs) > 0 {
		cmd.Set("version-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIAssistantVersionDeleteParams{
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

	return client.AI.Assistants.Versions.Delete(
		ctx,
		cmd.Value("version-id").(string),
		params,
		options...,
	)
}

func handleAIAssistantsVersionsPromote(ctx context.Context, cmd *cli.Command) error {
	client := telnyx.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("version-id") && len(unusedArgs) > 0 {
		cmd.Set("version-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := telnyx.AIAssistantVersionPromoteParams{
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
	_, err = client.AI.Assistants.Versions.Promote(
		ctx,
		cmd.Value("version-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:assistants:versions promote", obj, format, transform)
}
