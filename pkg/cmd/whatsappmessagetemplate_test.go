// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestWhatsappMessageTemplatesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "whatsapp-message-templates", "retrieve",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestWhatsappMessageTemplatesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "whatsapp-message-templates", "update",
			"--api-key", "string",
			"--id", "id",
			"--category", "MARKETING",
			"--component", "{foo: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"category: MARKETING\n" +
			"components:\n" +
			"  - foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "whatsapp-message-templates", "update",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestWhatsappMessageTemplatesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "whatsapp-message-templates", "delete",
			"--api-key", "string",
			"--id", "id",
		)
	})
}
