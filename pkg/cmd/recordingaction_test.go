// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestRecordingsActionsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"recordings:actions", "delete",
		"--id", "428c31b6-7af4-4bcb-b7f5-5013ef9657c1",
		"--id", "428c31b6-7af4-4bcb-b7f5-5013ef9657c2",
	)
}
