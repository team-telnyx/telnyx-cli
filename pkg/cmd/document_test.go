// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/team-telnyx/telnyx-cli/internal/mocktest"
	"github.com/team-telnyx/telnyx-cli/internal/requestflag"
)

func TestDocumentsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"documents", "retrieve",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestDocumentsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"documents", "update",
		"--document-id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		"--customer-reference", "MY REF 001",
		"--filename", "test-document.pdf",
	)
}

func TestDocumentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"documents", "list",
		"--filter", "{created_at: {gt: '2021-01-01T00:00:00Z', lt: '2021-04-09T22:25:27.521Z'}, customer_reference: {eq: MY REF 001, in: [REF001, REF002]}, filename: {contains: invoice}}",
		"--page-number", "0",
		"--page-size", "0",
		"--sort", "filename",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(documentsList)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"documents", "list",
		"--filter.created-at", "{gt: '2021-01-01T00:00:00Z', lt: '2021-04-09T22:25:27.521Z'}",
		"--filter.customer-reference", "{eq: MY REF 001, in: [REF001, REF002]}",
		"--filter.filename", "{contains: invoice}",
		"--page-number", "0",
		"--page-size", "0",
		"--sort", "filename",
	)
}

func TestDocumentsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"documents", "delete",
		"--id", "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
	)
}

func TestDocumentsGenerateDownloadLink(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"documents", "generate-download-link",
		"--id", "550e8400-e29b-41d4-a716-446655440000",
	)
}

func TestDocumentsUpload(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"documents", "upload",
		"--document", "{customer_reference: MY REF 001, file: ZXhhbXBsZSBvZiBlbmNvZGVkIGNvbnRlbnQ=, filename: test-document.pdf, url: https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(documentsUpload)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"documents", "upload",
		"--document.customer-reference", "MY REF 001",
		"--document.file", "ZXhhbXBsZSBvZiBlbmNvZGVkIGNvbnRlbnQ=",
		"--document.filename", "test-document.pdf",
		"--document.url", "https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf",
	)
}

func TestDocumentsUploadJson(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"documents", "upload-json",
		"--document", "{customer_reference: MY REF 001, file: ZXhhbXBsZSBvZiBlbmNvZGVkIGNvbnRlbnQ=, filename: test-document.pdf, url: https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(documentsUploadJson)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"documents", "upload-json",
		"--document.customer-reference", "MY REF 001",
		"--document.file", "ZXhhbXBsZSBvZiBlbmNvZGVkIGNvbnRlbnQ=",
		"--document.filename", "test-document.pdf",
		"--document.url", "https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf",
	)
}
