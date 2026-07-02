// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestDocumentLinksList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"document-links", "list",
			"--max-items", "10",
			"--filter", "{linked_record_type: porting_order, linked_resource_id: 6a09cdc3-8948-47f0-aa62-74ac943d6c58}",
			"--page-number", "0",
			"--page-size", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(documentLinksList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"document-links", "list",
			"--max-items", "10",
			"--filter.linked-record-type", "porting_order",
			"--filter.linked-resource-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
			"--page-number", "0",
			"--page-size", "0",
		)
	})
}
