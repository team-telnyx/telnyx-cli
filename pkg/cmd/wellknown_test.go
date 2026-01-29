// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestWellKnownRetrieveAuthorizationServerMetadata(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"well-known", "retrieve-authorization-server-metadata",
	)
}

func TestWellKnownRetrieveProtectedResourceMetadata(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"well-known", "retrieve-protected-resource-metadata",
	)
}
