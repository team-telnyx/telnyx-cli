// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestConnectionsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"connections", "retrieve",
		"--id", "id",
	)
}

func TestConnectionsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"connections", "list",
		"--filter", "{connection_name: {contains: contains}, fqdn: fqdn, outbound_voice_profile_id: outbound_voice_profile_id}",
		"--page", "{number: 1, size: 1}",
		"--sort", "connection_name",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(connectionsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"connections", "list",
		"--filter.connection-name", "{contains: contains}",
		"--filter.fqdn", "fqdn",
		"--filter.outbound-voice-profile-id", "outbound_voice_profile_id",
		"--page.number", "1",
		"--page.size", "1",
		"--sort", "connection_name",
	)
}

func TestConnectionsListActiveCalls(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"connections", "list-active-calls",
		"--connection-id", "1293384261075731461",
		"--page", "{after: after, before: before, limit: 1}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(connectionsListActiveCalls)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"connections", "list-active-calls",
		"--connection-id", "1293384261075731461",
		"--page.after", "after",
		"--page.before", "before",
		"--page.limit", "1",
		"--page-number", "0",
		"--page-size", "0",
	)
}
