// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestAIIntegrationsConnectionsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:integrations:connections", "retrieve",
		"--user-connection-id", "user_connection_id",
	)
}

func TestAIIntegrationsConnectionsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:integrations:connections", "list",
	)
}

func TestAIIntegrationsConnectionsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:integrations:connections", "delete",
		"--user-connection-id", "user_connection_id",
	)
}
