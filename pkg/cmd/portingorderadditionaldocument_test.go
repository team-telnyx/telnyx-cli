// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestPortingOrdersAdditionalDocumentsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:additional-documents", "create",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--additional-document", "{document_id: 22771a52-c43a-4539-80db-9dd9ec36e237, document_type: loa}",
		"--additional-document", "{document_id: d91474e6-4ebc-4ec1-b379-c596eeb405d6, document_type: invoice}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(portingOrdersAdditionalDocumentsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:additional-documents", "create",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--additional-document.document-id", "22771a52-c43a-4539-80db-9dd9ec36e237",
		"--additional-document.document-type", "loa",
		"--additional-document.document-id", "d91474e6-4ebc-4ec1-b379-c596eeb405d6",
		"--additional-document.document-type", "invoice",
	)
}

func TestPortingOrdersAdditionalDocumentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:additional-documents", "list",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--filter", "{document_type: [loa]}",
		"--page-number", "0",
		"--page-size", "0",
		"--sort", "{value: created_at}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(portingOrdersAdditionalDocumentsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:additional-documents", "list",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--filter.document-type", "[loa]",
		"--page-number", "0",
		"--page-size", "0",
		"--sort.value", "created_at",
	)
}

func TestPortingOrdersAdditionalDocumentsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"porting-orders:additional-documents", "delete",
		"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--additional-document-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}
