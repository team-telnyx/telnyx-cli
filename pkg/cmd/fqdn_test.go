// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestFqdnsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"fqdns", "create",
		"--connection-id", "1516447646313612565",
		"--dns-record-type", "a",
		"--fqdn", "example.com",
		"--port", "8080",
	)
}

func TestFqdnsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"fqdns", "retrieve",
		"--id", "id",
	)
}

func TestFqdnsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"fqdns", "update",
		"--id", "id",
		"--connection-id", "1516447646313612565",
		"--dns-record-type", "a",
		"--fqdn", "example.com",
		"--port", "8080",
	)
}

func TestFqdnsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"fqdns", "list",
		"--filter", "{connection_id: connection_id, dns_record_type: a, fqdn: example.com, port: 5060}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(fqdnsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"fqdns", "list",
		"--filter.connection-id", "connection_id",
		"--filter.dns-record-type", "a",
		"--filter.fqdn", "example.com",
		"--filter.port", "5060",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestFqdnsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"fqdns", "delete",
		"--id", "id",
	)
}
