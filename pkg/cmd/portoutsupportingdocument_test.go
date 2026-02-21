// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestPortoutsSupportingDocumentsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"portouts:supporting-documents", "create",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--document", "{document_id: 0ccc7b54-4df3-4bca-a65a-3da1ecc777f0, type: loa}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(portoutsSupportingDocumentsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"portouts:supporting-documents", "create",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--document.document-id", "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		"--document.type", "loa",
	)
}

func TestPortoutsSupportingDocumentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"portouts:supporting-documents", "list",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}
