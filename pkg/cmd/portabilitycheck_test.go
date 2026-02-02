// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestPortabilityChecksRun(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"portability-checks", "run",
		"--phone-number", "+13035550000",
		"--phone-number", "+13035550001",
		"--phone-number", "+13035550002",
	)
}
