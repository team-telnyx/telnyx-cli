// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestWhatsappBusinessAccountsSettingsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"whatsapp:business-accounts:settings", "retrieve",
			"--id", "id",
		)
	})
}

func TestWhatsappBusinessAccountsSettingsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"whatsapp:business-accounts:settings", "update",
			"--id", "id",
			"--name", "name",
			"--timezone", "America/New_York",
			"--webhook-enabled=true",
			"--webhook-event", "string",
			"--webhook-failover-url", "webhook_failover_url",
			"--webhook-url", "webhook_url",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"timezone: America/New_York\n" +
			"webhook_enabled: true\n" +
			"webhook_events:\n" +
			"  - string\n" +
			"webhook_failover_url: webhook_failover_url\n" +
			"webhook_url: webhook_url\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"whatsapp:business-accounts:settings", "update",
			"--id", "id",
		)
	})
}
