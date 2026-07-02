// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestOAuthClientsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"allowed_grant_types:\n" +
			"  - client_credentials\n" +
			"allowed_scopes:\n" +
			"  - admin\n" +
			"client_type: public\n" +
			"name: My OAuth client\n" +
			"logo_uri: https://example.com\n" +
			"policy_uri: https://example.com\n" +
			"redirect_uris:\n" +
			"  - https://example.com\n" +
			"require_pkce: true\n" +
			"tos_uri: https://example.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"oauth-clients", "create",
		)
	})
}

func TestOAuthClientsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"oauth-clients", "retrieve",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestOAuthClientsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"allowed_grant_types:\n" +
			"  - client_credentials\n" +
			"allowed_scopes:\n" +
			"  - admin\n" +
			"logo_uri: https://example.com\n" +
			"name: name\n" +
			"policy_uri: https://example.com\n" +
			"redirect_uris:\n" +
			"  - https://example.com\n" +
			"require_pkce: true\n" +
			"tos_uri: https://example.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"oauth-clients", "update",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestOAuthClientsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"oauth-clients", "list",
			"--max-items", "10",
			"--filter-allowed-grant-types-contains", "client_credentials",
			"--filter-client-id", "filter[client_id]",
			"--filter-client-type", "confidential",
			"--filter-name", "filter[name]",
			"--filter-name-contains", "filter[name][contains]",
			"--filter-verified=true",
			"--page-number", "1",
			"--page-size", "1",
		)
	})
}

func TestOAuthClientsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"oauth-clients", "delete",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
