// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestOAuthRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "oauth", "retrieve",
			"--api-key", "string",
			"--consent-token", "consent_token",
		)
	})
}

func TestOAuthGrants(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "oauth", "grants",
			"--api-key", "string",
			"--allowed=true",
			"--consent-token", "consent_token",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"allowed: true\n" +
			"consent_token: consent_token\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "oauth", "grants",
			"--api-key", "string",
		)
	})
}

func TestOAuthIntrospect(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "oauth", "introspect",
			"--api-key", "string",
			"--client-id", "string",
			"--client-secret", "string",
			"--token", "token",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("token: token")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "oauth", "introspect",
			"--api-key", "string",
			"--client-id", "string",
			"--client-secret", "string",
		)
	})
}

func TestOAuthRegister(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "oauth", "register",
			"--api-key", "string",
			"--client-name", "My OAuth Application",
			"--grant-type", "authorization_code",
			"--logo-uri", "https://example.com",
			"--policy-uri", "https://example.com",
			"--redirect-uris", "https://example.com/callback",
			"--response-type", "string",
			"--scope", "admin",
			"--token-endpoint-auth-method", "none",
			"--tos-uri", "https://example.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"client_name: My OAuth Application\n" +
			"grant_types:\n" +
			"  - authorization_code\n" +
			"logo_uri: https://example.com\n" +
			"policy_uri: https://example.com\n" +
			"redirect_uris:\n" +
			"  - https://example.com/callback\n" +
			"response_types:\n" +
			"  - string\n" +
			"scope: admin\n" +
			"token_endpoint_auth_method: none\n" +
			"tos_uri: https://example.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "oauth", "register",
			"--api-key", "string",
		)
	})
}

func TestOAuthRetrieveAuthorize(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "oauth", "retrieve-authorize",
			"--api-key", "string",
			"--client-id", "client_id",
			"--redirect-uri", "https://example.com",
			"--response-type", "code",
			"--code-challenge", "code_challenge",
			"--code-challenge-method", "plain",
			"--scope", "scope",
			"--state", "state",
		)
	})
}

func TestOAuthRetrieveJwks(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "oauth", "retrieve-jwks",
			"--api-key", "string",
		)
	})
}

func TestOAuthToken(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "oauth", "token",
			"--api-key", "string",
			"--client-id", "string",
			"--client-secret", "string",
			"--grant-type", "client_credentials",
			"--client-id", "client_id",
			"--client-secret", "client_secret",
			"--code", "code",
			"--code-verifier", "code_verifier",
			"--redirect-uri", "https://example.com",
			"--refresh-token", "refresh_token",
			"--scope", "admin",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"grant_type: client_credentials\n" +
			"client_id: client_id\n" +
			"client_secret: client_secret\n" +
			"code: code\n" +
			"code_verifier: code_verifier\n" +
			"redirect_uri: https://example.com\n" +
			"refresh_token: refresh_token\n" +
			"scope: admin\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "oauth", "token",
			"--api-key", "string",
			"--client-id", "string",
			"--client-secret", "string",
		)
	})
}
