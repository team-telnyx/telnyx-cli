// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestGlobalIPHealthChecksCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ip-health-checks", "create",
		"--global-ip-id", "a836125b-20b6-452e-9c03-2653f09c7ed7",
		"--health-check-params", "{path: bar, port: bar}",
		"--health-check-type", "http_status_2xx",
	)
}

func TestGlobalIPHealthChecksRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ip-health-checks", "retrieve",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestGlobalIPHealthChecksList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ip-health-checks", "list",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestGlobalIPHealthChecksDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"global-ip-health-checks", "delete",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}
