// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestSetiRetrieveBlackBoxTestResults(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"seti", "retrieve-black-box-test-results",
		"--filter", "{product: product}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(setiRetrieveBlackBoxTestResults)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"seti", "retrieve-black-box-test-results",
		"--filter.product", "product",
	)
}
