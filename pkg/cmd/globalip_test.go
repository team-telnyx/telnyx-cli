// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestGlobalIPsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ips", "create",
		"--api-key", "string",
		"--description", "test interface",
		"--name", "test interface",
		"--ports", "{tcp: bar, udp: bar}",
	)
}

func TestGlobalIPsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ips", "retrieve",
		"--api-key", "string",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestGlobalIPsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ips", "list",
		"--api-key", "string",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestGlobalIPsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ips", "delete",
		"--api-key", "string",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}
