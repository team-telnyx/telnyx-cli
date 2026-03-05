// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestIntegrationSecretsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "integration-secrets", "create",
			"--api-key", "string",
			"--identifier", "my_secret",
			"--type", "bearer",
			"--token", "my_secret_value",
			"--password", "password",
			"--username", "username",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"identifier: my_secret\n" +
			"type: bearer\n" +
			"token: my_secret_value\n" +
			"password: password\n" +
			"username: username\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "integration-secrets", "create",
			"--api-key", "string",
		)
	})
}

func TestIntegrationSecretsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "integration-secrets", "list",
			"--api-key", "string",
			"--filter", "{type: bearer}",
			"--page-number", "0",
			"--page-size", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(integrationSecretsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "integration-secrets", "list",
			"--api-key", "string",
			"--filter.type", "bearer",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestIntegrationSecretsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "integration-secrets", "delete",
			"--api-key", "string",
			"--id", "id",
		)
	})
}
