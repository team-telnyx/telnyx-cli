// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestWhatsappMessageTemplatesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"whatsapp:message-templates", "create",
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
			"whatsapp:message-templates", "create",
		)
	})
}

func TestWhatsappMessageTemplatesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"whatsapp:message-templates", "retrieve",
			"--id", "id",
		)
	})
}

func TestWhatsappMessageTemplatesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"whatsapp:message-templates", "update",
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
			t, pipeData,
			"--api-key", "string",
			"whatsapp:message-templates", "update",
			"--id", "id",
		)
	})
}

func TestWhatsappMessageTemplatesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"whatsapp:message-templates", "list",
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

func TestWhatsappMessageTemplatesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"whatsapp:message-templates", "delete",
			"--id", "id",
		)
	})
}
