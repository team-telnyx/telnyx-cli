// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestFqdnConnectionsFqdnAuthenticationList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"fqdn-connections:fqdn-authentication", "list",
			"--fqdn-connection-id", "fqdn_connection_id",
		)
	})
}

func TestFqdnConnectionsFqdnAuthenticationPatchAll(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"fqdn-connections:fqdn-authentication", "patch-all",
			"--fqdn-connection-id", "fqdn_connection_id",
			"--failover-url", "https://failover.example.com",
			"--fqdn-outbound-authentication", "ip-authentication",
			"--ip-authentication-method", "p-charge-info",
			"--password", "new_password",
			"--txt-name", "new_txt_name",
			"--txt-ttl", "300",
			"--txt-value", "new_txt_value",
			"--user-name", "newusername",
			"--webhook-url", "https://example.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"failover_url: https://failover.example.com\n" +
			"fqdn_outbound_authentication: ip-authentication\n" +
			"ip_authentication_method: p-charge-info\n" +
			"password: new_password\n" +
			"txt_name: new_txt_name\n" +
			"txt_ttl: 300\n" +
			"txt_value: new_txt_value\n" +
			"user_name: newusername\n" +
			"webhook_url: https://example.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"fqdn-connections:fqdn-authentication", "patch-all",
			"--fqdn-connection-id", "fqdn_connection_id",
		)
	})
}
