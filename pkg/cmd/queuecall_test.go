// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestQueuesCallsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"queues:calls", "retrieve",
		"--queue-name", "queue_name",
		"--call-control-id", "call_control_id",
	)
}

func TestQueuesCallsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"queues:calls", "update",
		"--queue-name", "queue_name",
		"--call-control-id", "call_control_id",
		"--keep-after-hangup=true",
	)
}

func TestQueuesCallsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"queues:calls", "list",
		"--queue-name", "queue_name",
		"--page", "{after: after, before: before, limit: 1}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(queuesCallsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"queues:calls", "list",
		"--queue-name", "queue_name",
		"--page.after", "after",
		"--page.before", "before",
		"--page.limit", "1",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestQueuesCallsRemove(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"queues:calls", "remove",
		"--queue-name", "queue_name",
		"--call-control-id", "call_control_id",
	)
}
