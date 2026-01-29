// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestIPsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ips", "create",
		"--ip-address", "192.168.0.0",
		"--connection-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--port", "5060",
	)
}

func TestIPsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ips", "retrieve",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestIPsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ips", "update",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--ip-address", "192.168.0.0",
		"--connection-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--port", "5060",
	)
}

func TestIPsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ips", "list",
		"--filter", "{connection_id: connection_id, ip_address: 192.168.0.0, port: 5060}",
		"--page", "{number: 1, size: 1}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(ipsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"ips", "list",
		"--filter.connection-id", "connection_id",
		"--filter.ip-address", "192.168.0.0",
		"--filter.port", "5060",
		"--page.number", "1",
		"--page.size", "1",
	)
}

func TestIPsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ips", "delete",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}
