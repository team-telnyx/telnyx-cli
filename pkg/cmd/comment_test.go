// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestCommentsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"comments", "create",
		"--body", "Hi there, ....",
		"--comment-record-id", "8ffb3622-7c6b-4ccc-b65f-7a3dc0099576",
		"--comment-record-type", "sub_number_order",
	)
}

func TestCommentsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"comments", "retrieve",
		"--id", "id",
	)
}

func TestCommentsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"comments", "list",
		"--filter", "{comment_record_id: 8ffb3622-7c6b-4ccc-b65f-7a3dc0099576, comment_record_type: sub_number_order}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(commentsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"comments", "list",
		"--filter.comment-record-id", "8ffb3622-7c6b-4ccc-b65f-7a3dc0099576",
		"--filter.comment-record-type", "sub_number_order",
	)
}

func TestCommentsMarkAsRead(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"comments", "mark-as-read",
		"--id", "id",
	)
}
