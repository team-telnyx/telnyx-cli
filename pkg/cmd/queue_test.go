// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestQueuesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"queues", "create",
		"--api-key", "string",
		"--queue-name", "tier_1_support",
		"--max-size", "100",
	)
}

func TestQueuesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"queues", "retrieve",
		"--api-key", "string",
		"--queue-name", "queue_name",
	)
}

func TestQueuesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"queues", "update",
		"--api-key", "string",
		"--queue-name", "queue_name",
		"--max-size", "200",
	)
}

func TestQueuesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"queues", "list",
		"--api-key", "string",
		"--page-number", "1",
		"--page-size", "1",
	)
}

func TestQueuesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"queues", "delete",
		"--api-key", "string",
		"--queue-name", "queue_name",
	)
}
