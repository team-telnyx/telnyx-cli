// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestMobilePushCredentialsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"mobile-push-credentials", "create",
		"--create-mobile-push-credential-request", "{alias: LucyIosCredential, certificate: '-----BEGIN CERTIFICATE----- MIIGVDCCBTKCAQEAsNlRJVZn9ZvXcECQm65czs... -----END CERTIFICATE-----', private_key: '-----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAsNlRJVZn9ZvXcECQm65czs... -----END RSA PRIVATE KEY-----', type: ios}",
	)
}

func TestMobilePushCredentialsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"mobile-push-credentials", "retrieve",
		"--push-credential-id", "0ccc7b76-4df3-4bca-a05a-3da1ecc389f0",
	)
}

func TestMobilePushCredentialsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"mobile-push-credentials", "list",
		"--filter", "{alias: LucyCredential, type: ios}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(mobilePushCredentialsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"mobile-push-credentials", "list",
		"--filter.alias", "LucyCredential",
		"--filter.type", "ios",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestMobilePushCredentialsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"mobile-push-credentials", "delete",
		"--push-credential-id", "0ccc7b76-4df3-4bca-a05a-3da1ecc389f0",
	)
}
