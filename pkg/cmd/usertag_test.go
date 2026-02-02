// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestUserTagsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"user-tags", "list",
		"--filter", "{starts_with: my-tag}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(userTagsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"user-tags", "list",
		"--filter.starts-with", "my-tag",
	)
}
