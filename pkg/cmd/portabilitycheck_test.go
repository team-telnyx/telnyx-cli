// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestPortabilityChecksRun(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "portability-checks", "run",
			"--api-key", "string",
			"--phone-number", "+13035550000",
			"--phone-number", "+13035550001",
			"--phone-number", "+13035550002",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"phone_numbers:\n" +
			"  - '+13035550000'\n" +
			"  - '+13035550001'\n" +
			"  - '+13035550002'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "portability-checks", "run",
			"--api-key", "string",
		)
	})
}
