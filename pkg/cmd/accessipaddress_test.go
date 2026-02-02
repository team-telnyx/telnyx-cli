// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestAccessIPAddressCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"access-ip-address", "create",
		"--ip-address", "ip_address",
		"--description", "description",
	)
}

func TestAccessIPAddressRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"access-ip-address", "retrieve",
		"--access-ip-address-id", "access_ip_address_id",
	)
}

func TestAccessIPAddressList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"access-ip-address", "list",
		"--filter", "{created_at: '2019-12-27T18:11:19.117Z', ip_address: ip_address, ip_source: ip_source}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(accessIPAddressList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"access-ip-address", "list",
		"--filter.created-at", "2019-12-27T18:11:19.117Z",
		"--filter.ip-address", "ip_address",
		"--filter.ip-source", "ip_source",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestAccessIPAddressDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"access-ip-address", "delete",
		"--access-ip-address-id", "access_ip_address_id",
	)
}
