// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIMissionsRunsEventsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:runs:events", "list",
		"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--agent-id", "agent_id",
		"--page-number", "1",
		"--page-size", "1",
		"--step-id", "step_id",
		"--type", "type",
	)
}

func TestAIMissionsRunsEventsGetEventDetails(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:runs:events", "get-event-details",
		"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--event-id", "event_id",
	)
}

func TestAIMissionsRunsEventsLog(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:missions:runs:events", "log",
		"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--summary", "summary",
		"--type", "status_change",
		"--agent-id", "agent_id",
		"--idempotency-key", "idempotency_key",
		"--payload", "{foo: bar}",
		"--step-id", "step_id",
	)
}
