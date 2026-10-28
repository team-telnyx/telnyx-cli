// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAITestsTestFrameworkV1ResultsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:results", "retrieve",
			"--id", "0",
		)
	})
}

func TestAITestsTestFrameworkV1ResultsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:results", "list",
			"--agent-id", "0",
			"--page", "1",
			"--page-size", "1",
			"--scenarios", "scenarios",
			"--search", "search",
		)
	})
}

func TestAITestsTestFrameworkV1ResultsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:results", "delete",
			"--id", "0",
		)
	})
}

func TestAITestsTestFrameworkV1ResultsRerun(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:tests:test-framework:v1:results", "rerun",
			"--id", "0",
		)
	})
}
