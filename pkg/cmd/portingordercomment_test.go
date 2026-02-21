// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestPortingOrdersCommentsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:comments", "create",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--body", "Please, let me know when the port completes",
	)
}

func TestPortingOrdersCommentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:comments", "list",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--page-number", "0",
		"--page-size", "0",
	)
}
