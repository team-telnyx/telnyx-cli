// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestPhoneNumberBlocksJobsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-number-blocks:jobs", "retrieve",
		"--id", "id",
	)
}

func TestPhoneNumberBlocksJobsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-number-blocks:jobs", "list",
		"--filter", "{status: in_progress, type: delete_phone_number_block}",
		"--page", "{number: 1, size: 1}",
		"--sort", "created_at",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(phoneNumberBlocksJobsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-number-blocks:jobs", "list",
		"--filter.status", "in_progress",
		"--filter.type", "delete_phone_number_block",
		"--page.number", "1",
		"--page.size", "1",
		"--sort", "created_at",
	)
}

func TestPhoneNumberBlocksJobsDeletePhoneNumberBlock(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"phone-number-blocks:jobs", "delete-phone-number-block",
		"--phone-number-block-id", "f3946371-7199-4261-9c3d-81a0d7935146",
	)
}
