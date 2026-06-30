// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestDirVerifyEmailConfirmCode(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dir:verify-email", "confirm-code",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
			"--code", "482915",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("code: '482915'")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"dir:verify-email", "confirm-code",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
		)
	})
}

func TestDirVerifyEmailSendCode(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dir:verify-email", "send-code",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
		)
	})
}

func TestDirVerifyEmailStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dir:verify-email", "status",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
		)
	})
}
