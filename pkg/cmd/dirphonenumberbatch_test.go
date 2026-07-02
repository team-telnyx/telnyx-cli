// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
)

func TestDirPhoneNumberBatchesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dir:phone-number-batches", "retrieve",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
			"--batch-id", "0a4b1f5e-2f12-4c0c-9a98-9b3a7d8b8e62",
		)
	})
}

func TestDirPhoneNumberBatchesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"dir:phone-number-batches", "list",
			"--max-items", "10",
			"--dir-id", "16635d38-75a6-4481-82e8-69af60e05011",
			"--filter-status", "submitted",
			"--page-number", "1",
			"--page-size", "20",
		)
	})
}
