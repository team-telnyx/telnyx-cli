// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestDialogflowConnectionsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"dialogflow-connections", "create",
		"--connection-id", "connection_id",
		"--service-account", "{type: bar, project_id: bar, private_key_id: bar, private_key: bar, client_email: bar, client_id: bar, auth_uri: bar, token_uri: bar, auth_provider_x509_cert_url: bar, client_x509_cert_url: bar}",
		"--conversation-profile-id", "a-VMHLWzTmKjiJw5S6O0-w",
		"--dialogflow-api", "cx",
		"--environment", "development",
		"--location", "global",
	)
}

func TestDialogflowConnectionsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"dialogflow-connections", "retrieve",
		"--connection-id", "connection_id",
	)
}

func TestDialogflowConnectionsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"dialogflow-connections", "update",
		"--connection-id", "connection_id",
		"--service-account", "{type: bar, project_id: bar, private_key_id: bar, private_key: bar, client_email: bar, client_id: bar, auth_uri: bar, token_uri: bar, auth_provider_x509_cert_url: bar, client_x509_cert_url: bar}",
		"--conversation-profile-id", "a-VMHLWzTmKjiJw5S6O0-w",
		"--dialogflow-api", "cx",
		"--environment", "development",
		"--location", "global",
	)
}

func TestDialogflowConnectionsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"dialogflow-connections", "delete",
		"--connection-id", "connection_id",
	)
}
