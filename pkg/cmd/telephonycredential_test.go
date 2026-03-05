// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestTelephonyCredentialsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "telephony-credentials", "create",
			"--api-key", "string",
			"--connection-id", "1234567890",
			"--expires-at", "2018-02-02T22:25:27.521Z",
			"--name", "My-new-credential",
			"--tag", "some_tag",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"connection_id: '1234567890'\n" +
			"expires_at: '2018-02-02T22:25:27.521Z'\n" +
			"name: My-new-credential\n" +
			"tag: some_tag\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "telephony-credentials", "create",
			"--api-key", "string",
		)
	})
}

func TestTelephonyCredentialsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "telephony-credentials", "retrieve",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestTelephonyCredentialsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "telephony-credentials", "update",
			"--api-key", "string",
			"--id", "id",
			"--connection-id", "987654321",
			"--expires-at", "2018-02-02T22:25:27.521Z",
			"--name", "My-new-updated-credential",
			"--tag", "some_tag",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"connection_id: '987654321'\n" +
			"expires_at: '2018-02-02T22:25:27.521Z'\n" +
			"name: My-new-updated-credential\n" +
			"tag: some_tag\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "telephony-credentials", "update",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestTelephonyCredentialsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "telephony-credentials", "list",
			"--api-key", "string",
			"--filter", "{name: name, resource_id: resource_id, sip_username: sip_username, status: status, tag: tag}",
			"--page-number", "0",
			"--page-size", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(telephonyCredentialsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "telephony-credentials", "list",
			"--api-key", "string",
			"--filter.name", "name",
			"--filter.resource-id", "resource_id",
			"--filter.sip-username", "sip_username",
			"--filter.status", "status",
			"--filter.tag", "tag",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestTelephonyCredentialsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "telephony-credentials", "delete",
			"--api-key", "string",
			"--id", "id",
		)
	})
}
