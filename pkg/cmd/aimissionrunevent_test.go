// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIMissionsRunsEventsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:missions:runs:events", "list",
			"--max-items", "10",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--agent-id", "agent_id",
			"--page-number", "1",
			"--page-size", "1",
			"--step-id", "step_id",
			"--type", "type",
		)
	})
}

func TestAIMissionsRunsEventsGetEventDetails(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai:missions:runs:events", "get-event-details",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--event-id", "event_id",
		)
	})
}

func TestAIMissionsRunsEventsLog(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"summary: summary\n" +
			"type: status_change\n" +
			"agent_id: agent_id\n" +
			"idempotency_key: idempotency_key\n" +
			"payload:\n" +
			"  foo: bar\n" +
			"step_id: step_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai:missions:runs:events", "log",
			"--mission-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--run-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
