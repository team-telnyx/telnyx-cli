// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestEnterprisesVerifyEmailCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises:verify-email", "create",
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
		)
	})
}

func TestEnterprisesVerifyEmailConfirm(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enterprises:verify-email", "confirm",
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
			"--code", "482915",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("code: '482915'")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"enterprises:verify-email", "confirm",
			"--enterprise-id", "4a6192a4-573d-446d-b3ce-aff9117272a6",
		)
	})
}
