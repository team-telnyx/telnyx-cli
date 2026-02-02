// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestSimCardDataUsageNotificationsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-card-data-usage-notifications", "create",
		"--sim-card-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--threshold", "{amount: '2048.1', unit: MB}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(simCardDataUsageNotificationsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-card-data-usage-notifications", "create",
		"--sim-card-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--threshold.amount", "2048.1",
		"--threshold.unit", "MB",
	)
}

func TestSimCardDataUsageNotificationsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-card-data-usage-notifications", "retrieve",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestSimCardDataUsageNotificationsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-card-data-usage-notifications", "update",
		"--sim-card-data-usage-notification-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--sim-card-id", "b34c1683-cd85-4493-b9a5-315eb4bc5e19",
		"--threshold", "{amount: '2048.0', unit: MB}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(simCardDataUsageNotificationsUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-card-data-usage-notifications", "update",
		"--sim-card-data-usage-notification-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--sim-card-id", "b34c1683-cd85-4493-b9a5-315eb4bc5e19",
		"--threshold.amount", "2048.0",
		"--threshold.unit", "MB",
	)
}

func TestSimCardDataUsageNotificationsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-card-data-usage-notifications", "list",
		"--filter-sim-card-id", "47a1c2b0-cc7b-4ab1-bb98-b33fb0fc61b9",
		"--page-number", "1",
		"--page-size", "1",
	)
}

func TestSimCardDataUsageNotificationsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-card-data-usage-notifications", "delete",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}
