// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestSipRegistrationStatusRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sip-registration-status", "retrieve",
			"--credential-type", "uac_external_credential",
			"--connection-id", "connection_id",
			"--username", "username",
		)
	})
}
