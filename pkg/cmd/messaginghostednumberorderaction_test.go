// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestMessagingHostedNumberOrdersActionsUploadFile(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"messaging-hosted-number-orders:actions", "upload-file",
			"--id", "id",
			"--bill", "Example data",
			"--loa", "Example data",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"bill: Example data\n" +
			"loa: Example data\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"messaging-hosted-number-orders:actions", "upload-file",
			"--id", "id",
		)
	})
}
