// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestTrafficPolicyProfilesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"traffic-policy-profiles", "create",
			"--type", "whitelist",
			"--domain", "www.hbomax.com",
			"--domain", "netflix.com",
			"--ip-range", "10.64.0.0/24",
			"--ip-range", "10.64.0.0/25",
			"--limit-bw-kbps", "512",
			"--service", "service_123",
			"--service", "service_456",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"type: whitelist\n" +
			"domains:\n" +
			"  - www.hbomax.com\n" +
			"  - netflix.com\n" +
			"ip_ranges:\n" +
			"  - 10.64.0.0/24\n" +
			"  - 10.64.0.0/25\n" +
			"limit_bw_kbps: 512\n" +
			"services:\n" +
			"  - service_123\n" +
			"  - service_456\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"traffic-policy-profiles", "create",
		)
	})
}

func TestTrafficPolicyProfilesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"traffic-policy-profiles", "retrieve",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestTrafficPolicyProfilesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"traffic-policy-profiles", "update",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--domain", "netflix.com",
			"--ip-range", "10.64.0.0/24",
			"--limit-bw-kbps", "1024",
			"--service", "service_123",
			"--type", "whitelist",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"domains:\n" +
			"  - netflix.com\n" +
			"ip_ranges:\n" +
			"  - 10.64.0.0/24\n" +
			"limit_bw_kbps: 1024\n" +
			"services:\n" +
			"  - service_123\n" +
			"type: whitelist\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"traffic-policy-profiles", "update",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestTrafficPolicyProfilesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"traffic-policy-profiles", "list",
			"--max-items", "10",
			"--filter-service", "filter[service]",
			"--filter-type", "whitelist",
			"--page-number", "1",
			"--page-size", "1",
			"--sort", "service",
		)
	})
}

func TestTrafficPolicyProfilesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"traffic-policy-profiles", "delete",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestTrafficPolicyProfilesListServices(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"traffic-policy-profiles", "list-services",
			"--max-items", "10",
			"--filter-group", "filter[group]",
			"--filter-name", "filter[name]",
			"--page-number", "1",
			"--page-size", "1",
		)
	})
}
