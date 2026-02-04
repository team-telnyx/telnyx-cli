// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIAssistantsTestsTestSuitesRunsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:assistants:tests:test-suites:runs", "list",
		"--suite-name", "suite_name",
		"--page-number", "0",
		"--page-size", "0",
		"--status", "status",
		"--test-suite-run-id", "test_suite_run_id",
	)
}

func TestAIAssistantsTestsTestSuitesRunsTrigger(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:assistants:tests:test-suites:runs", "trigger",
		"--suite-name", "suite_name",
		"--destination-version-id", "123e4567-e89b-12d3-a456-426614174000",
	)
}
