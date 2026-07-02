// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestDialogflowConnectionsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dialogflow-connections", "create",
			"--connection-id", "connection_id",
			"--service-account", "{type: bar, project_id: bar, private_key_id: bar, private_key: bar, client_email: bar, client_id: bar, auth_uri: bar, token_uri: bar, auth_provider_x509_cert_url: bar, client_x509_cert_url: bar}",
			"--conversation-profile-id", "a-VMHLWzTmKjiJw5S6O0-w",
			"--dialogflow-api", "cx",
			"--environment", "development",
			"--location", "global",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"service_account:\n" +
			"  type: bar\n" +
			"  project_id: bar\n" +
			"  private_key_id: bar\n" +
			"  private_key: bar\n" +
			"  client_email: bar\n" +
			"  client_id: bar\n" +
			"  auth_uri: bar\n" +
			"  token_uri: bar\n" +
			"  auth_provider_x509_cert_url: bar\n" +
			"  client_x509_cert_url: bar\n" +
			"conversation_profile_id: a-VMHLWzTmKjiJw5S6O0-w\n" +
			"dialogflow_api: cx\n" +
			"environment: development\n" +
			"location: global\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"dialogflow-connections", "create",
			"--connection-id", "connection_id",
		)
	})
}

func TestDialogflowConnectionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dialogflow-connections", "retrieve",
			"--connection-id", "connection_id",
		)
	})
}

func TestDialogflowConnectionsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dialogflow-connections", "update",
			"--connection-id", "connection_id",
			"--service-account", "{type: bar, project_id: bar, private_key_id: bar, private_key: bar, client_email: bar, client_id: bar, auth_uri: bar, token_uri: bar, auth_provider_x509_cert_url: bar, client_x509_cert_url: bar}",
			"--conversation-profile-id", "a-VMHLWzTmKjiJw5S6O0-w",
			"--dialogflow-api", "cx",
			"--environment", "development",
			"--location", "global",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"service_account:\n" +
			"  type: bar\n" +
			"  project_id: bar\n" +
			"  private_key_id: bar\n" +
			"  private_key: bar\n" +
			"  client_email: bar\n" +
			"  client_id: bar\n" +
			"  auth_uri: bar\n" +
			"  token_uri: bar\n" +
			"  auth_provider_x509_cert_url: bar\n" +
			"  client_x509_cert_url: bar\n" +
			"conversation_profile_id: a-VMHLWzTmKjiJw5S6O0-w\n" +
			"dialogflow_api: cx\n" +
			"environment: development\n" +
			"location: global\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"dialogflow-connections", "update",
			"--connection-id", "connection_id",
		)
	})
}

func TestDialogflowConnectionsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dialogflow-connections", "delete",
			"--connection-id", "connection_id",
		)
	})
}
