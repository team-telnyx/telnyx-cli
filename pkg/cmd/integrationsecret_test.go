// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestIntegrationSecretsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"integration-secrets", "create",
		"--identifier", "my_secret",
		"--type", "bearer",
		"--token", "my_secret_value",
		"--password", "password",
		"--username", "username",
	)
}

func TestIntegrationSecretsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"integration-secrets", "list",
		"--filter", "{type: bearer}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(integrationSecretsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"integration-secrets", "list",
		"--filter.type", "bearer",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestIntegrationSecretsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"integration-secrets", "delete",
		"--id", "id",
	)
}
