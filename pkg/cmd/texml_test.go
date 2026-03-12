// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestTexmlSecrets(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "texml", "secrets",
			"--api-key", "string",
			"--name", "My Secret Name",
			"--value", "My Secret Value",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: My Secret Name\n" +
			"value: My Secret Value\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "texml", "secrets",
			"--api-key", "string",
		)
	})
}
