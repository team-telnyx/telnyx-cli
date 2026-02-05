// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestTelephonyCredentialsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"telephony-credentials", "create",
		"--connection-id", "1234567890",
		"--expires-at", "2018-02-02T22:25:27.521Z",
		"--name", "My-new-credential",
		"--tag", "some_tag",
	)
}

func TestTelephonyCredentialsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"telephony-credentials", "retrieve",
		"--id", "id",
	)
}

func TestTelephonyCredentialsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"telephony-credentials", "update",
		"--id", "id",
		"--connection-id", "987654321",
		"--expires-at", "2018-02-02T22:25:27.521Z",
		"--name", "My-new-updated-credential",
		"--tag", "some_tag",
	)
}

func TestTelephonyCredentialsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"telephony-credentials", "list",
		"--filter", "{name: name, resource_id: resource_id, sip_username: sip_username, status: status, tag: tag}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(telephonyCredentialsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"telephony-credentials", "list",
		"--filter.name", "name",
		"--filter.resource-id", "resource_id",
		"--filter.sip-username", "sip_username",
		"--filter.status", "status",
		"--filter.tag", "tag",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestTelephonyCredentialsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"telephony-credentials", "delete",
		"--id", "id",
	)
}
