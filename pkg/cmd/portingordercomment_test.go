// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestPortingOrdersCommentsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:comments", "create",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--body", "Please, let me know when the port completes",
	)
}

func TestPortingOrdersCommentsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:comments", "list",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--page", "{number: 1, size: 1}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(portingOrdersCommentsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:comments", "list",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--page.number", "1",
		"--page.size", "1",
	)
}
