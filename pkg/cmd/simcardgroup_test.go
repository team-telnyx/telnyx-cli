// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestSimCardGroupsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sim-card-groups", "create",
			"--name", "My Test Group",
			"--data-limit", "{amount: '2048.1', unit: MB}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(simCardGroupsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sim-card-groups", "create",
			"--name", "My Test Group",
			"--data-limit.amount", "2048.1",
			"--data-limit.unit", "MB",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: My Test Group\n" +
			"data_limit:\n" +
			"  amount: '2048.1'\n" +
			"  unit: MB\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"sim-card-groups", "create",
		)
	})
}

func TestSimCardGroupsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sim-card-groups", "retrieve",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--include-iccids=true",
		)
	})
}

func TestSimCardGroupsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sim-card-groups", "update",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--data-limit", "{amount: '2048.1', unit: MB}",
			"--name", "My Test Group",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(simCardGroupsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sim-card-groups", "update",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--data-limit.amount", "2048.1",
			"--data-limit.unit", "MB",
			"--name", "My Test Group",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"data_limit:\n" +
			"  amount: '2048.1'\n" +
			"  unit: MB\n" +
			"name: My Test Group\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"sim-card-groups", "update",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}

func TestSimCardGroupsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sim-card-groups", "list",
			"--max-items", "10",
			"--filter-name", "My Test Group",
			"--filter-private-wireless-gateway-id", "7606c6d3-ff7c-49c1-943d-68879e9d584d",
			"--filter-wireless-blocklist-id", "0f3f490e-c4d3-4cf5-838a-9970f10ee259",
			"--page-number", "1",
			"--page-size", "1",
		)
	})
}

func TestSimCardGroupsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sim-card-groups", "delete",
			"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		)
	})
}
