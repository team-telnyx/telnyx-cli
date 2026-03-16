// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestQueuesCallsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"queues:calls", "retrieve",
			"--queue-name", "queue_name",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestQueuesCallsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"queues:calls", "update",
			"--queue-name", "queue_name",
			"--call-control-id", "call_control_id",
			"--keep-after-hangup=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("keep_after_hangup: true")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"queues:calls", "update",
			"--queue-name", "queue_name",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestQueuesCallsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"queues:calls", "list",
			"--max-items", "10",
			"--queue-name", "queue_name",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestQueuesCallsRemove(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"queues:calls", "remove",
			"--queue-name", "queue_name",
			"--call-control-id", "call_control_id",
		)
	})
}
