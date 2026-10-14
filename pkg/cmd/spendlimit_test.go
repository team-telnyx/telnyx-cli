// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestSpendLimitsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"spend-limits", "create",
			"--amount", "100",
			"--product", "inference",
			"--period", "daily",
			"--reason", "Team budget",
			"--unlimited=false",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"amount: 100\n" +
			"product: inference\n" +
			"period: daily\n" +
			"reason: Team budget\n" +
			"unlimited: false\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"spend-limits", "create",
		)
	})
}

func TestSpendLimitsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"spend-limits", "update",
			"--product", "inference",
			"--amount", "100",
			"--period", "daily",
			"--reason", "Raised for the product launch",
			"--unlimited=false",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"amount: 100\n" +
			"reason: Raised for the product launch\n" +
			"unlimited: false\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"spend-limits", "update",
			"--product", "inference",
			"--period", "daily",
		)
	})
}

func TestSpendLimitsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"spend-limits", "list",
		)
	})
}

func TestSpendLimitsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"spend-limits", "delete",
			"--product", "inference",
			"--period", "daily",
			"--reason", "reason",
		)
	})
}
