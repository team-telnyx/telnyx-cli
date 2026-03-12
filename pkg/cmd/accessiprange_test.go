// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestAccessIPRangesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "access-ip-ranges", "create",
			"--api-key", "string",
			"--cidr-block", "cidr_block",
			"--description", "description",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"cidr_block: cidr_block\n" +
			"description: description\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "access-ip-ranges", "create",
			"--api-key", "string",
		)
	})
}

func TestAccessIPRangesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "access-ip-ranges", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--filter", "{cidr_block: string, created_at: '2019-12-27T18:11:19.117Z'}",
			"--page-number", "0",
			"--page-size", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(accessIPRangesList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "access-ip-ranges", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--filter.cidr-block", "string",
			"--filter.created-at", "2019-12-27T18:11:19.117Z",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestAccessIPRangesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "access-ip-ranges", "delete",
			"--api-key", "string",
			"--access-ip-range-id", "access_ip_range_id",
		)
	})
}
