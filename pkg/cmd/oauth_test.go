// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestOAuthRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"oauth", "retrieve",
		"--api-key", "string",
		"--consent-token", "consent_token",
	)
}

func TestOAuthGrants(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"oauth", "grants",
		"--api-key", "string",
		"--allowed=true",
		"--consent-token", "consent_token",
	)
}

func TestOAuthIntrospect(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"oauth", "introspect",
		"--api-key", "string",
		"--client-id", "string",
		"--client-secret", "string",
		"--token", "token",
	)
}

func TestOAuthRegister(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"oauth", "register",
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
}

func TestOAuthRetrieveAuthorize(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"oauth", "retrieve-authorize",
		"--api-key", "string",
		"--client-id", "client_id",
		"--redirect-uri", "https://example.com",
		"--response-type", "code",
		"--code-challenge", "code_challenge",
		"--code-challenge-method", "plain",
		"--scope", "scope",
		"--state", "state",
	)
}

func TestOAuthRetrieveJwks(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"oauth", "retrieve-jwks",
		"--api-key", "string",
	)
}

func TestOAuthToken(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"oauth", "token",
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
}
