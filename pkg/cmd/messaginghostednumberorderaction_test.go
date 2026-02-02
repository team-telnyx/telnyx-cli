// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestMessagingHostedNumberOrdersActionsUploadFile(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"messaging-hosted-number-orders:actions", "upload-file",
		"--id", "id",
		"--bill", "",
		"--loa", "",
	)
}
