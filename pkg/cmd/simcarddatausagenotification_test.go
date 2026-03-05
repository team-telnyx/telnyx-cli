// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestSimCardDataUsageNotificationsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "sim-card-data-usage-notifications", "create",
			"--api-key", "string",
			"--sim-card-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--threshold", "{amount: '2048.1', unit: MB}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(simCardDataUsageNotificationsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "sim-card-data-usage-notifications", "create",
			"--api-key", "string",
			"--sim-card-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--threshold.amount", "2048.1",
			"--threshold.unit", "MB",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"sim_card_id: 6a09cdc3-8948-47f0-aa62-74ac943d6c58\n" +
			"threshold:\n" +
			"  amount: '2048.1'\n" +
			"  unit: MB\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "sim-card-data-usage-notifications", "create",
			"--api-key", "string",
		)
	})
}

func TestSimCardDataUsageNotificationsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "sim-card-data-usage-notifications", "retrieve",
			"--api-key", "string",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestSimCardDataUsageNotificationsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "sim-card-data-usage-notifications", "update",
			"--api-key", "string",
			"--sim-card-data-usage-notification-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--sim-card-id", "b34c1683-cd85-4493-b9a5-315eb4bc5e19",
			"--threshold", "{amount: '2048.0', unit: MB}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(simCardDataUsageNotificationsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "sim-card-data-usage-notifications", "update",
			"--api-key", "string",
			"--sim-card-data-usage-notification-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--sim-card-id", "b34c1683-cd85-4493-b9a5-315eb4bc5e19",
			"--threshold.amount", "2048.0",
			"--threshold.unit", "MB",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"sim_card_id: b34c1683-cd85-4493-b9a5-315eb4bc5e19\n" +
			"threshold:\n" +
			"  amount: '2048.0'\n" +
			"  unit: MB\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "sim-card-data-usage-notifications", "update",
			"--api-key", "string",
			"--sim-card-data-usage-notification-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestSimCardDataUsageNotificationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "sim-card-data-usage-notifications", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--filter-sim-card-id", "47a1c2b0-cc7b-4ab1-bb98-b33fb0fc61b9",
			"--page-number", "1",
			"--page-size", "1",
		)
	})
}

func TestSimCardDataUsageNotificationsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "sim-card-data-usage-notifications", "delete",
			"--api-key", "string",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}
