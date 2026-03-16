// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestSiprecConnectorsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"siprec-connectors", "create",
			"--host", "siprec.telnyx.com",
			"--name", "my-siprec-connector",
			"--port", "5060",
			"--app-subdomain", "my-app",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"host: siprec.telnyx.com\n" +
			"name: my-siprec-connector\n" +
			"port: 5060\n" +
			"app_subdomain: my-app\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"siprec-connectors", "create",
		)
	})
}

func TestSiprecConnectorsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"siprec-connectors", "retrieve",
			"--connector-name", "connector_name",
		)
	})
}

func TestSiprecConnectorsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"siprec-connectors", "update",
			"--connector-name", "connector_name",
			"--host", "siprec.telnyx.com",
			"--name", "my-siprec-connector",
			"--port", "5060",
			"--app-subdomain", "my-app",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"host: siprec.telnyx.com\n" +
			"name: my-siprec-connector\n" +
			"port: 5060\n" +
			"app_subdomain: my-app\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"siprec-connectors", "update",
			"--connector-name", "connector_name",
		)
	})
}

func TestSiprecConnectorsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"siprec-connectors", "delete",
			"--connector-name", "connector_name",
		)
	})
}
