// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIAssistantsToolsAdd(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:assistants:tools", "add",
			"--assistant-id", "assistant_id",
			"--tool-id", "tool_id",
		)
	})
}

func TestAIAssistantsToolsRemove(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:assistants:tools", "remove",
			"--assistant-id", "assistant_id",
			"--tool-id", "tool_id",
		)
	})
}

func TestAIAssistantsToolsTest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:assistants:tools", "test",
			"--assistant-id", "assistant_id",
			"--tool-id", "tool_id",
			"--arguments", "{order_id: bar}",
			"--dynamic-variables", "{customer_name: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"arguments:\n" +
			"  order_id: bar\n" +
			"dynamic_variables:\n" +
			"  customer_name: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:assistants:tools", "test",
			"--assistant-id", "assistant_id",
			"--tool-id", "tool_id",
		)
	})
}
