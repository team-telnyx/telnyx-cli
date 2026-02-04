// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestExternalConnectionsLogMessagesRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"external-connections:log-messages", "retrieve",
		"--id", "id",
	)
}

func TestExternalConnectionsLogMessagesList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"external-connections:log-messages", "list",
		"--filter", "{external_connection_id: 67ea7693-9cd5-4a68-8c76-abb3aa5bf5d2, telephone_number: {contains: '+123', eq: '+1234567890'}}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(externalConnectionsLogMessagesList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"external-connections:log-messages", "list",
		"--filter.external-connection-id", "67ea7693-9cd5-4a68-8c76-abb3aa5bf5d2",
		"--filter.telephone-number", "{contains: '+123', eq: '+1234567890'}",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestExternalConnectionsLogMessagesDismiss(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"external-connections:log-messages", "dismiss",
		"--id", "id",
	)
}
