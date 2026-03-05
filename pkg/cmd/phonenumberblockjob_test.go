// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestPhoneNumberBlocksJobsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "phone-number-blocks:jobs", "retrieve",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestPhoneNumberBlocksJobsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "phone-number-blocks:jobs", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--filter", "{status: in_progress, type: delete_phone_number_block}",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "created_at",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(phoneNumberBlocksJobsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "phone-number-blocks:jobs", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--filter.status", "in_progress",
			"--filter.type", "delete_phone_number_block",
			"--page-number", "0",
			"--page-size", "0",
			"--sort", "created_at",
		)
	})
}

func TestPhoneNumberBlocksJobsDeletePhoneNumberBlock(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "phone-number-blocks:jobs", "delete-phone-number-block",
			"--api-key", "string",
			"--phone-number-block-id", "f3946371-7199-4261-9c3d-81a0d7935146",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("phone_number_block_id: f3946371-7199-4261-9c3d-81a0d7935146")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "phone-number-blocks:jobs", "delete-phone-number-block",
			"--api-key", "string",
		)
	})
}
