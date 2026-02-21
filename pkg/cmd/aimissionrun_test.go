// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIMissionsRunsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:runs", "create",
		"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--input", "{foo: bar}",
		"--metadata", "{foo: bar}",
	)
}

func TestAIMissionsRunsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:runs", "retrieve",
		"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestAIMissionsRunsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:runs", "update",
		"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--error", "error",
		"--metadata", "{foo: bar}",
		"--result-payload", "{foo: bar}",
		"--result-summary", "result_summary",
		"--status", "pending",
	)
}

func TestAIMissionsRunsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:runs", "list",
		"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--page-number", "1",
		"--page-size", "1",
		"--status", "status",
	)
}

func TestAIMissionsRunsCancelRun(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:runs", "cancel-run",
		"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestAIMissionsRunsListRuns(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:runs", "list-runs",
		"--page-number", "1",
		"--page-size", "1",
		"--status", "status",
	)
}

func TestAIMissionsRunsPauseRun(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:runs", "pause-run",
		"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestAIMissionsRunsResumeRun(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:runs", "resume-run",
		"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}
