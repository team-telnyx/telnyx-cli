// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestGlobalIPsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ips", "create",
		"--description", "test interface",
		"--name", "test interface",
		"--ports", "{tcp: bar, udp: bar}",
	)
}

func TestGlobalIPsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ips", "retrieve",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestGlobalIPsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ips", "list",
		"--page", "{number: 1, size: 1}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(globalIPsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ips", "list",
		"--page.number", "1",
		"--page.size", "1",
	)
}

func TestGlobalIPsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ips", "delete",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}
