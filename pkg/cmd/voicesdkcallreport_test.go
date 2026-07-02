// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestVoiceSDKCallReportsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"voice-sdk-call-reports", "retrieve",
			"--call-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestVoiceSDKCallReportsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"voice-sdk-call-reports", "list",
			"--max-items", "10",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "-created_at",
		)
	})
}
