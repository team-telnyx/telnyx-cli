// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestBillingGroupsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "billing-groups", "create",
			"--api-key", "string",
			"--name", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("name: string")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "billing-groups", "create",
			"--api-key", "string",
		)
	})
}

func TestBillingGroupsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "billing-groups", "retrieve",
			"--api-key", "string",
			"--id", "f5586561-8ff0-4291-a0ac-84fe544797bd",
		)
	})
}

func TestBillingGroupsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "billing-groups", "update",
			"--api-key", "string",
			"--id", "f5586561-8ff0-4291-a0ac-84fe544797bd",
			"--name", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("name: string")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "billing-groups", "update",
			"--api-key", "string",
			"--id", "f5586561-8ff0-4291-a0ac-84fe544797bd",
		)
	})
}

func TestBillingGroupsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "billing-groups", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}

func TestBillingGroupsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "billing-groups", "delete",
			"--api-key", "string",
			"--id", "f5586561-8ff0-4291-a0ac-84fe544797bd",
		)
	})
}
