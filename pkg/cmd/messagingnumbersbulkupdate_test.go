// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestMessagingNumbersBulkUpdatesCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-numbers-bulk-updates", "create",
		"--messaging-profile-id", "00000000-0000-0000-0000-000000000000",
		"--number", "+18880000000",
		"--number", "+18880000001",
		"--number", "+18880000002",
	)
}

func TestMessagingNumbersBulkUpdatesRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-numbers-bulk-updates", "retrieve",
		"--order-id", "order_id",
	)
}
