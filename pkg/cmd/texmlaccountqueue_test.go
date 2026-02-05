// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestTexmlAccountsQueuesCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml:accounts:queues", "create",
		"--account-sid", "account_sid",
		"--friendly-name", "Support Queue",
		"--max-size", "10",
	)
}

func TestTexmlAccountsQueuesRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml:accounts:queues", "retrieve",
		"--account-sid", "account_sid",
		"--queue-sid", "queue_sid",
	)
}

func TestTexmlAccountsQueuesUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml:accounts:queues", "update",
		"--account-sid", "account_sid",
		"--queue-sid", "queue_sid",
		"--max-size", "10",
	)
}

func TestTexmlAccountsQueuesList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml:accounts:queues", "list",
		"--account-sid", "account_sid",
		"--date-created", "DateCreated",
		"--date-updated", "DateUpdated",
		"--page", "0",
		"--page-size", "0",
		"--page-token", "PageToken",
	)
}

func TestTexmlAccountsQueuesDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"texml:accounts:queues", "delete",
		"--account-sid", "account_sid",
		"--queue-sid", "queue_sid",
	)
}
