// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestAIChatCreateCompletion(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:chat", "create-completion",
			"--api-key", "string",
			"--message", "{content: You are a friendly chatbot., role: system}",
			"--message", "{content: 'Hello, world!', role: user}",
			"--api-key-ref", "api_key_ref",
			"--best-of", "0",
			"--early-stopping=true",
			"--enable-thinking=true",
			"--frequency-penalty", "0",
			"--guided-choice", "string",
			"--guided-json", "{foo: bar}",
			"--guided-regex", "guided_regex",
			"--length-penalty", "0",
			"--logprobs=true",
			"--max-tokens", "0",
			"--min-p", "0",
			"--model", "model",
			"--n", "0",
			"--presence-penalty", "0",
			"--response-format", "{type: text}",
			"--stream=true",
			"--temperature", "0",
			"--tool-choice", "none",
			"--tool", "{function: {name: name, description: description, parameters: {foo: bar}}, type: function}",
			"--top-logprobs", "0",
			"--top-p", "0",
			"--use-beam-search=true",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(aiChatCreateCompletion)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "ai:chat", "create-completion",
			"--api-key", "string",
			"--message.content", "You are a friendly chatbot.",
			"--message.role", "system",
			"--message.content", "Hello, world!",
			"--message.role", "user",
			"--api-key-ref", "api_key_ref",
			"--best-of", "0",
			"--early-stopping=true",
			"--enable-thinking=true",
			"--frequency-penalty", "0",
			"--guided-choice", "string",
			"--guided-json", "{foo: bar}",
			"--guided-regex", "guided_regex",
			"--length-penalty", "0",
			"--logprobs=true",
			"--max-tokens", "0",
			"--min-p", "0",
			"--model", "model",
			"--n", "0",
			"--presence-penalty", "0",
			"--response-format.type", "text",
			"--stream=true",
			"--temperature", "0",
			"--tool-choice", "none",
			"--tool", "{function: {name: name, description: description, parameters: {foo: bar}}, type: function}",
			"--top-logprobs", "0",
			"--top-p", "0",
			"--use-beam-search=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"messages:\n" +
			"  - content: You are a friendly chatbot.\n" +
			"    role: system\n" +
			"  - content: Hello, world!\n" +
			"    role: user\n" +
			"api_key_ref: api_key_ref\n" +
			"best_of: 0\n" +
			"early_stopping: true\n" +
			"enable_thinking: true\n" +
			"frequency_penalty: 0\n" +
			"guided_choice:\n" +
			"  - string\n" +
			"guided_json:\n" +
			"  foo: bar\n" +
			"guided_regex: guided_regex\n" +
			"length_penalty: 0\n" +
			"logprobs: true\n" +
			"max_tokens: 0\n" +
			"min_p: 0\n" +
			"model: model\n" +
			"'n': 0\n" +
			"presence_penalty: 0\n" +
			"response_format:\n" +
			"  type: text\n" +
			"stream: true\n" +
			"temperature: 0\n" +
			"tool_choice: none\n" +
			"tools:\n" +
			"  - function:\n" +
			"      name: name\n" +
			"      description: description\n" +
			"      parameters:\n" +
			"        foo: bar\n" +
			"    type: function\n" +
			"top_logprobs: 0\n" +
			"top_p: 0\n" +
			"use_beam_search: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "ai:chat", "create-completion",
			"--api-key", "string",
		)
	})
}
