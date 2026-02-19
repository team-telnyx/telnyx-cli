// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestAIChatCreateCompletion(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:chat", "create-completion",
		"--message", "{content: You are a friendly chatbot., role: system}",
		"--message", "{content: 'Hello, world!', role: user}",
		"--api-key-ref", "api_key_ref",
		"--best-of", "0",
		"--early-stopping=true",
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

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(aiChatCreateCompletion)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:chat", "create-completion",
		"--message.content", "You are a friendly chatbot.",
		"--message.role", "system",
		"--message.content", "Hello, world!",
		"--message.role", "user",
		"--api-key-ref", "api_key_ref",
		"--best-of", "0",
		"--early-stopping=true",
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
}
