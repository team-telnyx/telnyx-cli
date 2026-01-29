// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestAIAssistantsTestsRunsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:assistants:tests:runs", "retrieve",
		"--test-id", "test_id",
		"--run-id", "run_id",
	)
}

func TestAIAssistantsTestsRunsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:assistants:tests:runs", "list",
		"--test-id", "test_id",
		"--page-number", "0",
		"--page-size", "0",
		"--status", "status",
	)
}

func TestAIAssistantsTestsRunsTrigger(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:assistants:tests:runs", "trigger",
		"--test-id", "test_id",
		"--destination-version-id", "123e4567-e89b-12d3-a456-426614174000",
	)
}
