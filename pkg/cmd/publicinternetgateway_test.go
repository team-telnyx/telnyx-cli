// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestPublicInternetGatewaysCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"public-internet-gateways", "create",
		"--name", "test interface",
		"--network-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--region-code", "ashburn-va",
	)
}

func TestPublicInternetGatewaysRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"public-internet-gateways", "retrieve",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestPublicInternetGatewaysList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"public-internet-gateways", "list",
		"--filter", "{network_id: 6a09cdc3-8948-47f0-aa62-74ac943d6c58}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(publicInternetGatewaysList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"public-internet-gateways", "list",
		"--filter.network-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestPublicInternetGatewaysDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"public-internet-gateways", "delete",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}
