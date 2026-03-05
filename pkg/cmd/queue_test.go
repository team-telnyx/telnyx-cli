// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestQueuesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "queues", "create",
			"--api-key", "string",
			"--queue-name", "tier_1_support",
			"--max-size", "100",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"queue_name: tier_1_support\n" +
			"max_size: 100\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "queues", "create",
			"--api-key", "string",
		)
	})
}

func TestQueuesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "queues", "retrieve",
			"--api-key", "string",
			"--queue-name", "queue_name",
		)
	})
}

func TestQueuesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "queues", "update",
			"--api-key", "string",
			"--queue-name", "queue_name",
			"--max-size", "200",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("max_size: 200")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "queues", "update",
			"--api-key", "string",
			"--queue-name", "queue_name",
		)
	})
}

func TestQueuesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "queues", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--page-number", "1",
			"--page-size", "1",
		)
	})
}

func TestQueuesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "queues", "delete",
			"--api-key", "string",
			"--queue-name", "queue_name",
		)
	})
}
