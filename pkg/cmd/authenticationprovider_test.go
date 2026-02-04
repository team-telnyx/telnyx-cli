// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestAuthenticationProvidersCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"authentication-providers", "create",
		"--name", "Okta",
		"--settings", "{idp_cert_fingerprint: 13:38:C7:BB:C9:FF:4A:70:38:3A:E3:D9:5C:CD:DB:2E:50:1E:80:A7, idp_entity_id: https://myorg.myidp.com/saml/metadata, idp_sso_target_url: https://myorg.myidp.com/trust/saml2/http-post/sso, idp_cert_fingerprint_algorithm: sha256}",
		"--short-name", "myorg",
		"--active=true",
		"--settings-url", "https://myorg.myidp.com/saml/metadata",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(authenticationProvidersCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"authentication-providers", "create",
		"--name", "Okta",
		"--settings.idp-cert-fingerprint", "13:38:C7:BB:C9:FF:4A:70:38:3A:E3:D9:5C:CD:DB:2E:50:1E:80:A7",
		"--settings.idp-entity-id", "https://myorg.myidp.com/saml/metadata",
		"--settings.idp-sso-target-url", "https://myorg.myidp.com/trust/saml2/http-post/sso",
		"--settings.idp-cert-fingerprint-algorithm", "sha256",
		"--short-name", "myorg",
		"--active=true",
		"--settings-url", "https://myorg.myidp.com/saml/metadata",
	)
}

func TestAuthenticationProvidersRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"authentication-providers", "retrieve",
		"--id", "id",
	)
}

func TestAuthenticationProvidersUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"authentication-providers", "update",
		"--id", "id",
		"--active=true",
		"--name", "Okta",
		"--settings", "{idp_cert_fingerprint: 13:38:C7:BB:C9:FF:4A:70:38:3A:E3:D9:5C:CD:DB:2E:50:1E:80:A7, idp_entity_id: https://myorg.myidp.com/saml/metadata, idp_sso_target_url: https://myorg.myidp.com/trust/saml2/http-post/sso, idp_cert_fingerprint_algorithm: sha1}",
		"--settings-url", "https://myorg.myidp.com/saml/metadata",
		"--short-name", "myorg",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(authenticationProvidersUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"authentication-providers", "update",
		"--id", "id",
		"--active=true",
		"--name", "Okta",
		"--settings.idp-cert-fingerprint", "13:38:C7:BB:C9:FF:4A:70:38:3A:E3:D9:5C:CD:DB:2E:50:1E:80:A7",
		"--settings.idp-entity-id", "https://myorg.myidp.com/saml/metadata",
		"--settings.idp-sso-target-url", "https://myorg.myidp.com/trust/saml2/http-post/sso",
		"--settings.idp-cert-fingerprint-algorithm", "sha1",
		"--settings-url", "https://myorg.myidp.com/saml/metadata",
		"--short-name", "myorg",
	)
}

func TestAuthenticationProvidersList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"authentication-providers", "list",
		"--page-number", "0",
		"--page-size", "0",
		"--sort", "name",
	)
}

func TestAuthenticationProvidersDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"authentication-providers", "delete",
		"--id", "id",
	)
}
