// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestCommentsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"comments", "create",
			"--body", "Hi there, ....",
			"--comment-record-id", "8ffb3622-7c6b-4ccc-b65f-7a3dc0099576",
			"--comment-record-type", "sub_number_order",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"body: Hi there, ....\n" +
			"comment_record_id: 8ffb3622-7c6b-4ccc-b65f-7a3dc0099576\n" +
			"comment_record_type: sub_number_order\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"comments", "create",
		)
	})
}

func TestCommentsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"comments", "retrieve",
			"--id", "id",
		)
	})
}

func TestCommentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"comments", "list",
			"--filter", "{comment_record_id: 8ffb3622-7c6b-4ccc-b65f-7a3dc0099576, comment_record_type: sub_number_order}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(commentsList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"comments", "list",
			"--filter.comment-record-id", "8ffb3622-7c6b-4ccc-b65f-7a3dc0099576",
			"--filter.comment-record-type", "sub_number_order",
		)
	})
}

func TestCommentsMarkAsRead(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"comments", "mark-as-read",
			"--id", "id",
		)
	})
}
