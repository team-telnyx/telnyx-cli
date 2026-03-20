// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestSimCardsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sim-cards", "retrieve",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--include-pin-puk-codes=true",
			"--include-sim-card-group=true",
		)
	})
}

func TestSimCardsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sim-cards", "update",
			"--sim-card-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--authorized-imei", "['106516771852751', '534051870479563', '508821468377961']",
			"--data-limit", "{amount: '2048.1', unit: MB}",
			"--sim-card-group-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--status", "{}",
			"--tag", "personal",
			"--tag", "customers",
			"--tag", "active-customers",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(simCardsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sim-cards", "update",
			"--sim-card-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--authorized-imei", "['106516771852751', '534051870479563', '508821468377961']",
			"--data-limit.amount", "2048.1",
			"--data-limit.unit", "MB",
			"--sim-card-group-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--tag", "personal",
			"--tag", "customers",
			"--tag", "active-customers",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"authorized_imeis:\n" +
			"  - '106516771852751'\n" +
			"  - '534051870479563'\n" +
			"  - '508821468377961'\n" +
			"data_limit:\n" +
			"  amount: '2048.1'\n" +
			"  unit: MB\n" +
			"sim_card_group_id: 6a09cdc3-8948-47f0-aa62-74ac943d6c58\n" +
			"status: {}\n" +
			"tags:\n" +
			"  - personal\n" +
			"  - customers\n" +
			"  - active-customers\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"sim-cards", "update",
			"--sim-card-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestSimCardsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sim-cards", "list",
			"--max-items", "10",
			"--filter", "{iccid: '89310410106543789301', msisdn: '+13109976224', status: [enabled], tags: [personal, customers, active-customers]}",
			"--filter-sim-card-group-id", "47a1c2b0-cc7b-4ab1-bb98-b33fb0fc61b9",
			"--include-sim-card-group=true",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "current_billing_period_consumed_data.amount",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(simCardsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sim-cards", "list",
			"--max-items", "10",
			"--filter.iccid", "89310410106543789301",
			"--filter.msisdn", "+13109976224",
			"--filter.status", "[enabled]",
			"--filter.tags", "[personal, customers, active-customers]",
			"--filter-sim-card-group-id", "47a1c2b0-cc7b-4ab1-bb98-b33fb0fc61b9",
			"--include-sim-card-group=true",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "current_billing_period_consumed_data.amount",
		)
	})
}

func TestSimCardsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sim-cards", "delete",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--report-lost=true",
		)
	})
}

func TestSimCardsGetActivationCode(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sim-cards", "get-activation-code",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestSimCardsGetDeviceDetails(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sim-cards", "get-device-details",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestSimCardsGetPublicIP(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sim-cards", "get-public-ip",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestSimCardsListWirelessConnectivityLogs(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sim-cards", "list-wireless-connectivity-logs",
			"--max-items", "10",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--page-number", "1",
			"--page-size", "1",
		)
	})
}
