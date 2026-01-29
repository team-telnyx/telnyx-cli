// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestWirelessBlocklistsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"wireless-blocklists", "create",
		"--name", "My Wireless Blocklist",
		"--type", "country",
		"--value", "CA",
		"--value", "US",
	)
}

func TestWirelessBlocklistsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"wireless-blocklists", "retrieve",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestWirelessBlocklistsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"wireless-blocklists", "update",
		"--name", "My Wireless Blocklist",
		"--type", "country",
		"--value", "CA",
		"--value", "US",
	)
}

func TestWirelessBlocklistsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"wireless-blocklists", "list",
		"--filter-name", "filter[name]",
		"--filter-type", "filter[type]",
		"--filter-values", "filter[values]",
		"--page-number", "1",
		"--page-size", "1",
	)
}

func TestWirelessBlocklistsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"wireless-blocklists", "delete",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}
