// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestQueuesCallsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"queues:calls", "retrieve",
		"--queue-name", "queue_name",
		"--call-control-id", "call_control_id",
	)
}

func TestQueuesCallsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"queues:calls", "update",
		"--queue-name", "queue_name",
		"--call-control-id", "call_control_id",
		"--keep-after-hangup=true",
	)
}

func TestQueuesCallsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"queues:calls", "list",
		"--queue-name", "queue_name",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestQueuesCallsRemove(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"queues:calls", "remove",
		"--queue-name", "queue_name",
		"--call-control-id", "call_control_id",
	)
}
