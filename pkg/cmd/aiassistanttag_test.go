// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIAssistantsTagsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:assistants:tags", "list",
			"--api-key", "string",
		)
	})
}

func TestAIAssistantsTagsAdd(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:assistants:tags", "add",
			"--api-key", "string",
			"--assistant-id", "assistant_id",
			"--tag", "tag",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("tag: tag")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "ai:assistants:tags", "add",
			"--api-key", "string",
			"--assistant-id", "assistant_id",
		)
	})
}

func TestAIAssistantsTagsRemove(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "ai:assistants:tags", "remove",
			"--api-key", "string",
			"--assistant-id", "assistant_id",
			"--tag", "tag",
		)
	})
}
