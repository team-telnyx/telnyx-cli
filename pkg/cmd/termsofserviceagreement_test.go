// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestTermsOfServiceAgreementsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"terms-of-service:agreements", "retrieve",
			"--agreement-id", "550e8400-e29b-41d4-a716-446655440000",
		)
	})
}

func TestTermsOfServiceAgreementsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"terms-of-service:agreements", "list",
			"--max-items", "10",
			"--page-number", "1",
			"--page-size", "20",
			"--product-type", "branded_calling",
		)
	})
}
