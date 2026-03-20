// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestWhatsappTemplatesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"whatsapp:templates", "create",
			"--category", "MARKETING",
			"--component", "{foo: bar}",
			"--language", "language",
			"--name", "name",
			"--waba-id", "waba_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"category: MARKETING\n" +
			"components:\n" +
			"  - foo: bar\n" +
			"language: language\n" +
			"name: name\n" +
			"waba_id: waba_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"whatsapp:templates", "create",
		)
	})
}

func TestWhatsappTemplatesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"whatsapp:templates", "list",
			"--max-items", "10",
			"--filter-category", "MARKETING",
			"--filter-search", "filter[search]",
			"--filter-status", "filter[status]",
			"--filter-waba-id", "filter[waba_id]",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}
