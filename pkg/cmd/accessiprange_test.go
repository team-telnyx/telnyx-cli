// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestAccessIPRangesCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"access-ip-ranges", "create",
		"--cidr-block", "cidr_block",
		"--description", "description",
	)
}

func TestAccessIPRangesList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"access-ip-ranges", "list",
		"--filter", "{cidr_block: string, created_at: '2019-12-27T18:11:19.117Z'}",
		"--page-number", "0",
		"--page-size", "0",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(accessIPRangesList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"access-ip-ranges", "list",
		"--filter.cidr-block", "string",
		"--filter.created-at", "2019-12-27T18:11:19.117Z",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestAccessIPRangesDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"access-ip-ranges", "delete",
		"--access-ip-range-id", "access_ip_range_id",
	)
}
