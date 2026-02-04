// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestSiprecConnectorsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"siprec-connectors", "create",
		"--host", "siprec.telnyx.com",
		"--name", "my-siprec-connector",
		"--port", "5060",
		"--app-subdomain", "my-app",
	)
}

func TestSiprecConnectorsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"siprec-connectors", "retrieve",
		"--connector-name", "connector_name",
	)
}

func TestSiprecConnectorsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"siprec-connectors", "update",
		"--connector-name", "connector_name",
		"--host", "siprec.telnyx.com",
		"--name", "my-siprec-connector",
		"--port", "5060",
		"--app-subdomain", "my-app",
	)
}

func TestSiprecConnectorsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"siprec-connectors", "delete",
		"--connector-name", "connector_name",
	)
}
