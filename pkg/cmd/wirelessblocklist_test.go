// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestWirelessBlocklistsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "wireless-blocklists", "create",
			"--api-key", "string",
			"--name", "My Wireless Blocklist",
			"--type", "country",
			"--value", "CA",
			"--value", "US",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: My Wireless Blocklist\n" +
			"type: country\n" +
			"values:\n" +
			"  - CA\n" +
			"  - US\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "wireless-blocklists", "create",
			"--api-key", "string",
		)
	})
}

func TestWirelessBlocklistsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "wireless-blocklists", "retrieve",
			"--api-key", "string",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestWirelessBlocklistsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "wireless-blocklists", "update",
			"--api-key", "string",
			"--name", "My Wireless Blocklist",
			"--type", "country",
			"--value", "CA",
			"--value", "US",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: My Wireless Blocklist\n" +
			"type: country\n" +
			"values:\n" +
			"  - CA\n" +
			"  - US\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "wireless-blocklists", "update",
			"--api-key", "string",
		)
	})
}

func TestWirelessBlocklistsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "wireless-blocklists", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--filter-name", "filter[name]",
			"--filter-type", "filter[type]",
			"--filter-values", "filter[values]",
			"--page-number", "1",
			"--page-size", "1",
		)
	})
}

func TestWirelessBlocklistsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "wireless-blocklists", "delete",
			"--api-key", "string",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}
