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
			t, "queues:calls", "retrieve",
			"--api-key", "string",
			"--queue-name", "queue_name",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestQueuesCallsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "queues:calls", "update",
			"--api-key", "string",
			"--queue-name", "queue_name",
			"--call-control-id", "call_control_id",
			"--keep-after-hangup=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("keep_after_hangup: true")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "queues:calls", "update",
			"--api-key", "string",
			"--queue-name", "queue_name",
			"--call-control-id", "call_control_id",
		)
	})
}

func TestQueuesCallsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "queues:calls", "list",
			"--api-key", "string",
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
			t, "queues:calls", "remove",
			"--api-key", "string",
			"--queue-name", "queue_name",
			"--call-control-id", "call_control_id",
		)
	})
}
