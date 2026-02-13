// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestSimCardsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-cards", "retrieve",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--include-pin-puk-codes=true",
		"--include-sim-card-group=true",
	)
}

func TestSimCardsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
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

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(simCardsUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
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
}

func TestSimCardsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-cards", "list",
		"--filter", "{iccid: '89310410106543789301', msisdn: '+13109976224', status: [enabled], tags: [personal, customers, active-customers]}",
		"--filter-sim-card-group-id", "47a1c2b0-cc7b-4ab1-bb98-b33fb0fc61b9",
		"--include-sim-card-group=true",
		"--page-number", "0",
		"--page-size", "0",
		"--sort", "current_billing_period_consumed_data.amount",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(simCardsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-cards", "list",
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
}

func TestSimCardsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-cards", "delete",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--report-lost=true",
	)
}

func TestSimCardsGetActivationCode(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-cards", "get-activation-code",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestSimCardsGetDeviceDetails(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-cards", "get-device-details",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestSimCardsGetPublicIP(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-cards", "get-public-ip",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestSimCardsListWirelessConnectivityLogs(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"sim-cards", "list-wireless-connectivity-logs",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--page-number", "1",
		"--page-size", "1",
	)
}
