// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
)

func TestPortoutsCommentsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"portouts:comments", "create",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--body", "body",
	)
}

func TestPortoutsCommentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"portouts:comments", "list",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}
