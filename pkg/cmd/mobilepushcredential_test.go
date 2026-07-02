// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestMobilePushCredentialsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mobile-push-credentials", "create",
			"--create-mobile-push-credential-request", "{alias: LucyIosCredential, certificate: '-----BEGIN CERTIFICATE----- MIIGVDCCBTKCAQEAsNlRJVZn9ZvXcECQm65czs... -----END CERTIFICATE-----', private_key: '-----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAsNlRJVZn9ZvXcECQm65czs... -----END RSA PRIVATE KEY-----', type: ios}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"alias: LucyIosCredential\n" +
			"certificate: >-\n" +
			"  -----BEGIN CERTIFICATE----- MIIGVDCCBTKCAQEAsNlRJVZn9ZvXcECQm65czs... -----END\n" +
			"  CERTIFICATE-----\n" +
			"private_key: >-\n" +
			"  -----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAsNlRJVZn9ZvXcECQm65czs...\n" +
			"  -----END RSA PRIVATE KEY-----\n" +
			"type: ios\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"mobile-push-credentials", "create",
		)
	})
}

func TestMobilePushCredentialsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mobile-push-credentials", "retrieve",
			"--push-credential-id", "0ccc7b76-4df3-4bca-a05a-3da1ecc389f0",
		)
	})
}

func TestMobilePushCredentialsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mobile-push-credentials", "list",
			"--max-items", "10",
			"--filter", "{alias: LucyCredential, type: ios}",
			"--page-number", "0",
			"--page-size", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(mobilePushCredentialsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mobile-push-credentials", "list",
			"--max-items", "10",
			"--filter.alias", "LucyCredential",
			"--filter.type", "ios",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestMobilePushCredentialsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"mobile-push-credentials", "delete",
			"--push-credential-id", "0ccc7b76-4df3-4bca-a05a-3da1ecc389f0",
		)
	})
}
