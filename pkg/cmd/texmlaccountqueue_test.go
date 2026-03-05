// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestTexmlAccountsQueuesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "texml:accounts:queues", "create",
			"--api-key", "string",
			"--account-sid", "account_sid",
			"--friendly-name", "Support Queue",
			"--max-size", "10",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"FriendlyName: Support Queue\n" +
			"MaxSize: 10\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "texml:accounts:queues", "create",
			"--api-key", "string",
			"--account-sid", "account_sid",
		)
	})
}

func TestTexmlAccountsQueuesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "texml:accounts:queues", "retrieve",
			"--api-key", "string",
			"--account-sid", "account_sid",
			"--queue-sid", "queue_sid",
		)
	})
}

func TestTexmlAccountsQueuesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "texml:accounts:queues", "update",
			"--api-key", "string",
			"--account-sid", "account_sid",
			"--queue-sid", "queue_sid",
			"--max-size", "10",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("MaxSize: 10")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "texml:accounts:queues", "update",
			"--api-key", "string",
			"--account-sid", "account_sid",
			"--queue-sid", "queue_sid",
		)
	})
}

func TestTexmlAccountsQueuesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "texml:accounts:queues", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--account-sid", "account_sid",
			"--date-created", "DateCreated",
			"--date-updated", "DateUpdated",
			"--page", "0",
			"--page-size", "0",
			"--page-token", "PageToken",
		)
	})
}

func TestTexmlAccountsQueuesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "texml:accounts:queues", "delete",
			"--api-key", "string",
			"--account-sid", "account_sid",
			"--queue-sid", "queue_sid",
		)
	})
}
