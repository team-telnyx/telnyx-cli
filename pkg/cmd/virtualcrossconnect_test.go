// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestVirtualCrossConnectsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "virtual-cross-connects", "create",
			"--api-key", "string",
			"--region-code", "ashburn-va",
			"--bandwidth-mbps", "50",
			"--bgp-asn", "1234",
			"--cloud-provider", "aws",
			"--cloud-provider-region", "us-east-1",
			"--name", "test interface",
			"--network-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--primary-bgp-key", "yFV4wEPtPVPfDUGLWiyQzwga",
			"--primary-cloud-account-id", "123456789012",
			"--primary-cloud-ip", "169.254.0.2",
			"--primary-telnyx-ip", "169.254.0.1",
			"--secondary-bgp-key", "ge1lONeK9RcA83uuWaw9DvZy",
			"--secondary-cloud-account-id", "",
			"--secondary-cloud-ip", "169.254.0.4",
			"--secondary-telnyx-ip", "169.254.0.3",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"region_code: ashburn-va\n" +
			"bandwidth_mbps: 50\n" +
			"bgp_asn: 1234\n" +
			"cloud_provider: aws\n" +
			"cloud_provider_region: us-east-1\n" +
			"name: test interface\n" +
			"network_id: 6a09cdc3-8948-47f0-aa62-74ac943d6c58\n" +
			"primary_bgp_key: yFV4wEPtPVPfDUGLWiyQzwga\n" +
			"primary_cloud_account_id: '123456789012'\n" +
			"primary_cloud_ip: 169.254.0.2\n" +
			"primary_telnyx_ip: 169.254.0.1\n" +
			"secondary_bgp_key: ge1lONeK9RcA83uuWaw9DvZy\n" +
			"secondary_cloud_account_id: ''\n" +
			"secondary_cloud_ip: 169.254.0.4\n" +
			"secondary_telnyx_ip: 169.254.0.3\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "virtual-cross-connects", "create",
			"--api-key", "string",
		)
	})
}

func TestVirtualCrossConnectsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "virtual-cross-connects", "retrieve",
			"--api-key", "string",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestVirtualCrossConnectsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "virtual-cross-connects", "update",
			"--api-key", "string",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--primary-cloud-ip", "169.254.0.2",
			"--primary-enabled=true",
			"--primary-routing-announcement=false",
			"--secondary-cloud-ip", "169.254.0.4",
			"--secondary-enabled=true",
			"--secondary-routing-announcement=false",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"primary_cloud_ip: 169.254.0.2\n" +
			"primary_enabled: true\n" +
			"primary_routing_announcement: false\n" +
			"secondary_cloud_ip: 169.254.0.4\n" +
			"secondary_enabled: true\n" +
			"secondary_routing_announcement: false\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "virtual-cross-connects", "update",
			"--api-key", "string",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestVirtualCrossConnectsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "virtual-cross-connects", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--filter", "{network_id: 6a09cdc3-8948-47f0-aa62-74ac943d6c58}",
			"--page-number", "0",
			"--page-size", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(virtualCrossConnectsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "virtual-cross-connects", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--filter.network-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestVirtualCrossConnectsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "virtual-cross-connects", "delete",
			"--api-key", "string",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}
