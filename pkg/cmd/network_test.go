// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestNetworksCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"networks", "create",
		"--name", "test network",
	)
}

func TestNetworksRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"networks", "retrieve",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestNetworksUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"networks", "update",
		"--network-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--name", "test network",
	)
}

func TestNetworksList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"networks", "list",
		"--filter", "{name: test network}",
		"--page", "{number: 1, size: 1}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(networksList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"networks", "list",
		"--filter.name", "test network",
		"--page.number", "1",
		"--page.size", "1",
	)
}

func TestNetworksDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"networks", "delete",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestNetworksListInterfaces(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"networks", "list-interfaces",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--filter", "{name: test interface, type: wireguard_interface}",
		"--page", "{number: 1, size: 1}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(networksListInterfaces)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"networks", "list-interfaces",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--filter.name", "test interface",
		"--filter.type", "wireguard_interface",
		"--page.number", "1",
		"--page.size", "1",
	)
}
