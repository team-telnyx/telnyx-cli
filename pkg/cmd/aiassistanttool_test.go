// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIAssistantsToolsTest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:assistants:tools", "test",
		"--assistant-id", "assistant_id",
		"--tool-id", "tool_id",
		"--arguments", "{foo: bar}",
		"--dynamic-variables", "{foo: bar}",
	)
}
