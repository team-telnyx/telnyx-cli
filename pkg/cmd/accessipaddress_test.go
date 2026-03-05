// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestAccessIPAddressCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "access-ip-address", "create",
			"--api-key", "string",
			"--ip-address", "ip_address",
			"--description", "description",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"ip_address: ip_address\n" +
			"description: description\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "access-ip-address", "create",
			"--api-key", "string",
		)
	})
}

func TestAccessIPAddressRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "access-ip-address", "retrieve",
			"--api-key", "string",
			"--access-ip-address-id", "access_ip_address_id",
		)
	})
}

func TestAccessIPAddressList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "access-ip-address", "list",
			"--api-key", "string",
			"--filter", "{created_at: '2019-12-27T18:11:19.117Z', ip_address: ip_address, ip_source: ip_source}",
			"--page-number", "0",
			"--page-size", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(accessIPAddressList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "access-ip-address", "list",
			"--api-key", "string",
			"--filter.created-at", "2019-12-27T18:11:19.117Z",
			"--filter.ip-address", "ip_address",
			"--filter.ip-source", "ip_source",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestAccessIPAddressDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "access-ip-address", "delete",
			"--api-key", "string",
			"--access-ip-address-id", "access_ip_address_id",
		)
	})
}
