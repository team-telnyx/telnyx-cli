// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/telnyx-cli/internal/mocktest"
	"github.com/stainless-sdks/telnyx-cli/internal/requestflag"
)

func TestDocumentLinksList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"document-links", "list",
		"--filter", "{linked_record_type: porting_order, linked_resource_id: 6a09cdc3-8948-47f0-aa62-74ac943d6c58}",
		"--page", "{number: 1, size: 1}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(documentLinksList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"document-links", "list",
		"--filter.linked-record-type", "porting_order",
		"--filter.linked-resource-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--page.number", "1",
		"--page.size", "1",
	)
}
