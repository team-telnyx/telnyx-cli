// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestPrivateWirelessGatewaysCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"private-wireless-gateways", "create",
		"--name", "My private wireless gateway",
		"--network-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--region-code", "dc2",
	)
}

func TestPrivateWirelessGatewaysRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"private-wireless-gateways", "retrieve",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestPrivateWirelessGatewaysList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"private-wireless-gateways", "list",
		"--filter-created-at", "filter[created_at]",
		"--filter-ip-range", "filter[ip_range]",
		"--filter-name", "filter[name]",
		"--filter-region-code", "filter[region_code]",
		"--filter-updated-at", "filter[updated_at]",
		"--page-number", "1",
		"--page-size", "1",
	)
}

func TestPrivateWirelessGatewaysDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"private-wireless-gateways", "delete",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}
