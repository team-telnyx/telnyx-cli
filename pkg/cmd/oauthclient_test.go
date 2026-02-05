// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestOAuthClientsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"oauth-clients", "create",
		"--allowed-grant-type", "client_credentials",
		"--allowed-scope", "admin",
		"--client-type", "public",
		"--name", "My OAuth client",
		"--logo-uri", "https://example.com",
		"--policy-uri", "https://example.com",
		"--redirect-uris", "https://example.com",
		"--require-pkce=true",
		"--tos-uri", "https://example.com",
	)
}

func TestOAuthClientsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"oauth-clients", "retrieve",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestOAuthClientsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"oauth-clients", "update",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--allowed-grant-type", "client_credentials",
		"--allowed-scope", "admin",
		"--logo-uri", "https://example.com",
		"--name", "name",
		"--policy-uri", "https://example.com",
		"--redirect-uris", "https://example.com",
		"--require-pkce=true",
		"--tos-uri", "https://example.com",
	)
}

func TestOAuthClientsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"oauth-clients", "list",
		"--filter-allowed-grant-types-contains", "client_credentials",
		"--filter-client-id", "filter[client_id]",
		"--filter-client-type", "confidential",
		"--filter-name", "filter[name]",
		"--filter-name-contains", "filter[name][contains]",
		"--filter-verified=true",
		"--page-number", "1",
		"--page-size", "1",
	)
}

func TestOAuthClientsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"oauth-clients", "delete",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}
